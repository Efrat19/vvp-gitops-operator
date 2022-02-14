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

// DeploymentTargetReconciler reconciles a DeploymentTarget object
type DeploymentTargetReconciler struct {
	client.Client
	Scheme    *runtime.Scheme
	vvpClient vvp_client.VvpClient
}

//+kubebuilder:rbac:groups=appmanager.vvp.efrat19.io,resources=deploymentTargets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=appmanager.vvp.efrat19.io,resources=deploymentTargets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=appmanager.vvp.efrat19.io,resources=deploymentTargets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DeploymentTarget object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *DeploymentTargetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	var dep appmanagervvpv1alpha1.DeploymentTarget
	if err := r.Get(ctx, req.NamespacedName, &dep); err != nil {
		log.Error(err, "unable to get deploymentTarget")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if err := r.vvpClient.ProbeServer(); err != nil {
		return r.handleOutOfSyncError(dep, err)
	}
	// if the deploymentTarget needs to be deleted
	if !dep.ObjectMeta.DeletionTimestamp.IsZero() {
		if err := r.handleDeploymentTargetDeletion(dep); err != nil {
			return r.handleOutOfSyncError(dep, err)
		}
		return ctrl.Result{}, nil
	}
	if err := r.handleDeploymentTargetCreationIfNeeded(&dep); err != nil {
		return r.handleOutOfSyncError(dep, err)
	}
	if err := r.updateDeploymentTargetSpecInVVP(dep); err != nil {
		return r.handleOutOfSyncError(dep, err)
	}
	if err := r.setStatus(dep, vvp_client.InSyncState); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *DeploymentTargetReconciler) handleDeploymentTargetCreationIfNeeded(dep *appmanagervvpv1alpha1.DeploymentTarget) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	// Create deploymentTarget if not exists
	err, deploymentTargetExists := r.vvpClient.DeploymentTargets().ResourceExistsInVVP(dep)
	if err != nil {
		log.Error(err, "unable to check whether vvp deploymentTarget exists")
		return err
	}
	if err := r.attachFinalizers(*dep); err != nil {
		log.Error(err, "failed to attach finalizers")
		return err
	}
	if !deploymentTargetExists {
		log.Info(fmt.Sprintf("DeploymentTarget %s doesnt exist in vvp, attempting to create\n", dep.Spec.Metadata.Name))
		if err := r.vvpClient.DeploymentTargets().CreateExternalResources(dep); err != nil {
			log.Error(err, "unable to create vvp deploymentTarget")
			r.detachFinalizers(*dep)
			return err
		}
	}
	return nil
}

func (r *DeploymentTargetReconciler) attachFinalizers(dep appmanagervvpv1alpha1.DeploymentTarget) error {
	ctx := context.Background()
	// name of our custom finalizer
	log := log.FromContext(ctx)

	// examine DeletionTimestamp to determine if object is under deletion
	if dep.ObjectMeta.DeletionTimestamp.IsZero() {
		log.Info(fmt.Sprintf("Attaching finalizers to deploymentTarget %s\n", dep.Spec.Metadata.Name))
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

func (r *DeploymentTargetReconciler) detachFinalizers(dep appmanagervvpv1alpha1.DeploymentTarget) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	if controllerutil.ContainsFinalizer(&dep, appmanagerFinalizer) {
		controllerutil.RemoveFinalizer(&dep, appmanagerFinalizer)
		if err := r.Update(ctx, &dep); err != nil {
			log.Error(err, fmt.Sprintf("Failed to remove deploymentTarget %s finalizers\n", dep.Spec.Metadata.Name))
			return err
		}
		return nil
	}
	return nil
}

func (r *DeploymentTargetReconciler) handleDeploymentTargetDeletion(dep appmanagervvpv1alpha1.DeploymentTarget) error {
	ctx := context.Background()
	// name of our custom finalizer
	log := log.FromContext(ctx)
	err, deploymentTargetExists := r.vvpClient.DeploymentTargets().ResourceExistsInVVP(&dep)
	if err != nil {
		log.Error(err, "unable to check whether vvp deploymentTarget exists")
		return err
	}
	if !deploymentTargetExists {
		log.Info(fmt.Sprintf("deploymentTarget %s doesnt exist in vvp, skipping deletion\n", dep.Spec.Metadata.Name))
		return nil
	}
	// The object is being deleted
	log.Info(fmt.Sprintf("Deleting deploymentTarget %s\n", dep.Spec.Metadata.Name))
	if controllerutil.ContainsFinalizer(&dep, appmanagerFinalizer) {
		// our finalizer is present, so lets handle any external dependency
		forceDelete := false
		if err := r.vvpClient.DeploymentTargets().DeleteExternalResources(&dep, forceDelete); err != nil {
			log.Error(err, fmt.Sprintf("Failed to delete deploymentTarget %s in vvp, retrying...\n", dep.Spec.Metadata.Name))
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

func (r *DeploymentTargetReconciler) handleOutOfSyncError(dep appmanagervvpv1alpha1.DeploymentTarget, err error) (ctrl.Result, error) {
	if updateErr := r.setStatus(dep, vvp_client.FormatOutOfSync(err)); updateErr != nil {
		return ctrl.Result{}, updateErr
	}
	if errors.Is(err, vvp_client.ErrRetryable) {
		return ctrl.Result{RequeueAfter: time.Second * 30, Requeue: true}, err
	}
	return ctrl.Result{}, err
}

func (r *DeploymentTargetReconciler) setStatus(dep appmanagervvpv1alpha1.DeploymentTarget, syncState string) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	log.Info(fmt.Sprintf("Updating status for deploymentTarget %s \n", dep.Spec.Metadata.Name))
	dep.Status.State = syncState
	dep.Status.LastSync = metav1.Now()
	if err := r.Status().Update(ctx, &dep); err != nil {
		log.Error(err, "unable to update k8s deploymentTarget status")
		return err
	}
	return nil
}

func (r *DeploymentTargetReconciler) updateDeploymentTargetSpecInVVP(dep appmanagervvpv1alpha1.DeploymentTarget) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	// update vvp deploymentTarget spec from k8s
	log.Info(fmt.Sprintf("Updating spec for deploymentTarget %s \n", dep.Spec.Metadata.Name))
	if err := r.vvpClient.DeploymentTargets().UpdateExternalResources(&dep); err != nil {
		log.Error(err, "unable to update vvp deploymentTarget spec")
		return nil
	}
	return nil

}

// SetupWithManager sets up the controller with the Manager.
func (r *DeploymentTargetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.vvpClient = vvp_client.NewClient()
	return ctrl.NewControllerManagedBy(mgr).
		For(&appmanagervvpv1alpha1.DeploymentTarget{}).
		Complete(r)
}
