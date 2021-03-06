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

package appmanagervvp

import (
	"context"
	"errors"
	"fmt"
	"time"

	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/apis/appmanager.vvp/v1alpha1"

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

const (
	appmanagerFinalizer string = "appmanager.vvp.efrat19.io/finalizer"
)

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
	if err := r.vvpClient.ProbeServer(); err != nil {
		return r.handleOutOfSyncError(dep, err)
	}
	// if the deployment needs to be deleted
	if !dep.ObjectMeta.DeletionTimestamp.IsZero() {
		if err := r.handleDeploymentDeletion(dep); err != nil {
			return r.handleOutOfSyncError(dep, err)
		}
		return ctrl.Result{}, nil
	}
	if err := r.handleDeploymentCreationIfNeeded(&dep); err != nil {
		return r.handleOutOfSyncError(dep, err)
	}
	if err := r.updateDeploymentSpecInVVP(dep); err != nil {
		return r.handleOutOfSyncError(dep, err)
	}
	if err := r.setStatus(dep, vvp_client.InSyncState); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *DeploymentReconciler) handleDeploymentCreationIfNeeded(dep *appmanagervvpv1alpha1.Deployment) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	// Create deployment if not exists
	err, deploymentExists := r.vvpClient.Deployments().ResourceExistsInVVP(dep)
	if err != nil {
		log.Error(err, "unable to check whether vvp deployment exists")
		return err
	}
	if err := r.attachFinalizers(*dep); err != nil {
		log.Error(err, "failed to attach finalizers")
		return err
	}
	if !deploymentExists {
		log.Info(fmt.Sprintf("Deployment %s doesnt exist in vvp, attempting to create\n", dep.Spec.Metadata.Name))
		if err := r.vvpClient.Deployments().CreateExternalResources(dep); err != nil {
			log.Error(err, "unable to create vvp deployment")
			r.detachFinalizers(*dep)
			return err
		}
	}
	return nil
}

func (r *DeploymentReconciler) attachFinalizers(dep appmanagervvpv1alpha1.Deployment) error {
	ctx := context.Background()
	// name of our custom finalizer
	log := log.FromContext(ctx)

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
	}
	return nil
}

func (r *DeploymentReconciler) detachFinalizers(dep appmanagervvpv1alpha1.Deployment) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	if controllerutil.ContainsFinalizer(&dep, appmanagerFinalizer) {
		controllerutil.RemoveFinalizer(&dep, appmanagerFinalizer)
		if err := r.Update(ctx, &dep); err != nil {
			log.Error(err, fmt.Sprintf("Failed to remove deployment %s finalizers\n", dep.Spec.Metadata.Name))
			return err
		}
		return nil
	}
	return nil
}

func (r *DeploymentReconciler) handleDeploymentDeletion(dep appmanagervvpv1alpha1.Deployment) error {
	ctx := context.Background()
	// name of our custom finalizer
	log := log.FromContext(ctx)
	err, deploymentExists := r.vvpClient.Deployments().ResourceExistsInVVP(&dep)
	if err != nil {
		log.Error(err, "unable to check whether vvp deployment exists")
		return err
	}
	if !deploymentExists {
		log.Info(fmt.Sprintf("deployment %s doesnt exist in vvp, skipping deletion\n", dep.Spec.Metadata.Name))
		return nil
	}
	// The object is being deleted
	log.Info(fmt.Sprintf("Deleting deployment %s\n", dep.Spec.Metadata.Name))
	if controllerutil.ContainsFinalizer(&dep, appmanagerFinalizer) {
		// our finalizer is present, so lets handle any external dependency
		if err := r.vvpClient.Deployments().DeleteExternalResources(&dep); err != nil {
			log.Error(err, fmt.Sprintf("Failed to delete deployment %s in vvp, retrying...\n", dep.Spec.Metadata.Name))
			// if fail to delete the external dependency here, return with error
			// so that it can be retried
			return vvp_client.NewRetryableError(err)
		}

		// remove our finalizer from the list and update it.
		return r.detachFinalizers(dep)
	}
	// Stop reconciliation as the item is being deleted
	return nil
}

func (r *DeploymentReconciler) handleOutOfSyncError(dep appmanagervvpv1alpha1.Deployment, err error) (ctrl.Result, error) {
	if updateErr := r.setStatus(dep, vvp_client.FormatOutOfSync(err)); updateErr != nil {
		return ctrl.Result{}, updateErr
	}
	if errors.Is(err, vvp_client.ErrRetryable) {
		return ctrl.Result{RequeueAfter: time.Second * 30, Requeue: true}, err
	}
	return ctrl.Result{}, err
}

func (r *DeploymentReconciler) setStatus(dep appmanagervvpv1alpha1.Deployment, syncState string) error {
	ctx := context.Background()
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

func (r *DeploymentReconciler) updateDeploymentSpecInVVP(dep appmanagervvpv1alpha1.Deployment) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	// update vvp deployment spec from k8s
	log.Info(fmt.Sprintf("Updating spec for deployment %s \n", dep.Spec.Metadata.Name))
	if err := r.vvpClient.Deployments().UpdateExternalResources(&dep); err != nil {
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
