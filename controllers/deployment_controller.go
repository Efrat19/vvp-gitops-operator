/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"errors"
	"fmt"
	"time"

	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/api/v1alpha1"
	"efrat19.io/vvp-gitops-operator/pkg/vvp_client"

	// "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/polling"
	// "github.com/davecgh/go-spew/spew"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// DeploymentReconciler reconciles a Deployment object
type DeploymentReconciler struct {
	client.Client
	Scheme    *runtime.Scheme
	vvpClient vvp_client.VvpClient
}

//+kubebuilder:rbac:groups=appmanager.vvp.efrat19.io,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=appmanager.vvp.efrat19.io,resources=deployments/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=appmanager.vvp.efrat19.io,resources=deployments/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Deployment object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *DeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	var dep appmanagervvpv1alpha1.Deployment
	if err := r.Get(ctx, req.NamespacedName, &dep); err != nil {
		log.Error(err, "unable to get deployment")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if err := r.vvpClient.ProbeServer(ctx); err != nil {
		return r.handleOutOfSyncError(ctx, dep, err)
	}
	if err := r.handleDeploymentDeletionIfNeeded(ctx, dep); err != nil {
		return r.handleOutOfSyncError(ctx, dep, err)
	}
	if err := r.handleDeploymentCreationIfNeeded(ctx, &dep); err != nil {
		return r.handleOutOfSyncError(ctx, dep, err)
	}
	if err := r.updateDeploymentSpecInVVP(ctx, dep); err != nil {
		return r.handleOutOfSyncError(ctx, dep, err)
	}
	if err := r.setStatus(ctx, dep, vvp_client.InSyncState); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *DeploymentReconciler) handleDeploymentCreationIfNeeded(ctx context.Context, dep *appmanagervvpv1alpha1.Deployment) error {
	log := log.FromContext(ctx)
	// Create deployment if not exists
	err, deploymentExists := r.vvpClient.Deployments().ResourceExistsInVVP(ctx, dep)
	if err != nil {
		log.Error(err, "unable to check whether vvp deployment exists")
		return nil
	}
	if !deploymentExists {
		log.Info(fmt.Sprintf("Deployment %s doesnt exist in vvp, attempting to create\n", dep.Spec.Metadata.Name))
		if err := r.vvpClient.Deployments().CreateExternalResources(ctx, dep); err != nil {
			log.Error(err, "unable to create vvp deployment")
		}
	}
	return nil
}

func (r *DeploymentReconciler) handleDeploymentDeletionIfNeeded(ctx context.Context, dep appmanagervvpv1alpha1.Deployment) error {
	// name of our custom finalizer
	log := log.FromContext(ctx)
	appmanagerFinalizer := "appmanager.vvp.efrat19.io/finalizer"

	// examine DeletionTimestamp to determine if object is under deletion
	if dep.ObjectMeta.DeletionTimestamp.IsZero() {
		log.Info(fmt.Sprintf("Attaching finalizers to deployment %s\n", dep.Spec.Metadata.Name))
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// registering our finalizer.
		if !controllerutil.ContainsFinalizer(&dep, appmanagerFinalizer) {
			controllerutil.AddFinalizer(&dep, appmanagerFinalizer)
			if err := r.Update(ctx, &dep); err != nil {
				return err
			}
		}
	} else {
		// The object is being deleted
		log.Info(fmt.Sprintf("Deleting deployment %s\n", dep.Spec.Metadata.Name))
		if controllerutil.ContainsFinalizer(&dep, appmanagerFinalizer) {
			// our finalizer is present, so lets handle any external dependency
			if err := r.vvpClient.Deployments().DeleteExternalResources(ctx, &dep); err != nil {
				log.Error(err, fmt.Sprintf("Failed to delete deployment %s in vvp, retrying...\n", dep.Spec.Metadata.Name))
				// if fail to delete the external dependency here, return with error
				// so that it can be retried
				return vvp_client.NewRetryableError(err)
			}

			// remove our finalizer from the list and update it.
			controllerutil.RemoveFinalizer(&dep, appmanagerFinalizer)
			if err := r.Update(ctx, &dep); err != nil {
				log.Error(err, fmt.Sprintf("Failed to remove deployment %s finalizers\n", dep.Spec.Metadata.Name))
				return err
			}
		}
		// Stop reconciliation as the item is being deleted
		return nil
	}
	return nil
}

func (r *DeploymentReconciler) handleOutOfSyncError(ctx context.Context, dep appmanagervvpv1alpha1.Deployment, err error) (ctrl.Result, error) {
	if updateErr := r.setStatus(ctx, dep, vvp_client.FormatOutOfSync(err)); updateErr != nil {
		return ctrl.Result{}, updateErr
	}
	if errors.Is(err, vvp_client.ErrRetryable) {
		return ctrl.Result{RequeueAfter: time.Second * 30, Requeue: true}, err
	}
	return ctrl.Result{}, err
}

func (r *DeploymentReconciler) setStatus(ctx context.Context, dep appmanagervvpv1alpha1.Deployment, syncState string) error {
	log := log.FromContext(ctx)
	log.Info(fmt.Sprintf("Updating status for deployment %s \n", dep.Spec.Metadata.Name))
	dep.Status.State = syncState
	dep.Status.LastSync = metav1.Now()
	if err := r.Status().Update(ctx, &dep); err != nil {
		log.Error(err, "unable to update k8s deployment status")
		return err
	}
	return nil
}

func (r *DeploymentReconciler) updateDeploymentSpecInVVP(ctx context.Context, dep appmanagervvpv1alpha1.Deployment) error {
	log := log.FromContext(ctx)
	// update vvp deployment spec from k8s
	log.Info(fmt.Sprintf("Updating spec for deployment %s \n", dep.Spec.Metadata.Name))
	if err := r.vvpClient.Deployments().UpdateExternalResources(ctx, &dep); err != nil {
		log.Error(err, "unable to update vvp deployment spec")
		return nil
	}
	return nil

}

// SetupWithManager sets up the controller with the Manager.
func (r *DeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.vvpClient = vvp_client.NewClient()
	return ctrl.NewControllerManagedBy(mgr).
		For(&appmanagervvpv1alpha1.Deployment{}).
		Complete(r)
}
