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

// SavepointReconciler reconciles a Savepoint object
type SavepointReconciler struct {
	client.Client
	Scheme    *runtime.Scheme
	vvpClient vvp_client.VvpClient
}

//+kubebuilder:rbac:groups=appmanager.vvp.efrat19.io,resources=savepoints,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=appmanager.vvp.efrat19.io,resources=savepoints/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=appmanager.vvp.efrat19.io,resources=savepoints/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Savepoint object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *SavepointReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	var sp appmanagervvpv1alpha1.Savepoint
	if err := r.Get(ctx, req.NamespacedName, &sp); err != nil {
		log.Error(err, "unable to get Savepoint")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if err := r.vvpClient.ProbeServer(); err != nil {
		return r.handleOutOfSyncError(sp, err)
	}
	// if the Savepoint needs to be deleted
	if !sp.ObjectMeta.DeletionTimestamp.IsZero() {
		if err := r.handleSavepointDeletion(sp); err != nil {
			return r.handleOutOfSyncError(sp, err)
		}
		return ctrl.Result{}, nil
	}
	if err := r.handleSavepointCreationIfNeeded(&sp); err != nil {
		return r.handleOutOfSyncError(sp, err)
	}
	if err := r.updateSavepointSpecInVVP(sp); err != nil {
		return r.handleOutOfSyncError(sp, err)
	}
	if err := r.setStatus(sp, vvp_client.InSyncState); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *SavepointReconciler) handleSavepointCreationIfNeeded(sp *appmanagervvpv1alpha1.Savepoint) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	// Create Savepoint if not exists
	err, SavepointExists := r.vvpClient.SavePoints().ResourceExistsInVVP(sp)
	if err != nil {
		log.Error(err, "unable to check whether vvp Savepoint exists")
		return err
	}
	if err := r.attachFinalizers(*sp); err != nil {
		log.Error(err, "failed to attach finalizers")
		return err
	}
	if !SavepointExists {
		log.Info(fmt.Sprintf("Savepoint %s doesnt exist in vvp, attempting to create\n", sp.Spec.Metadata.Id))
		if err := r.vvpClient.SavePoints().CreateExternalResources(sp); err != nil {
			log.Error(err, "unable to create vvp Savepoint")
			r.detachFinalizers(*sp)
			return err
		}
	}
	return nil
}

func (r *SavepointReconciler) attachFinalizers(sp appmanagervvpv1alpha1.Savepoint) error {
	ctx := context.Background()
	// name of our custom finalizer
	log := log.FromContext(ctx)

	// examine DeletionTimestamp to determine if object is under deletion
	if sp.ObjectMeta.DeletionTimestamp.IsZero() {
		log.Info(fmt.Sprintf("Attaching finalizers to Savepoint %s\n", sp.Spec.Metadata.Id))
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// registering our finalizer.
		if !controllerutil.ContainsFinalizer(&sp, appmanagerFinalizer) {
			controllerutil.AddFinalizer(&sp, appmanagerFinalizer)
			if err := r.Update(ctx, &sp); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *SavepointReconciler) detachFinalizers(sp appmanagervvpv1alpha1.Savepoint) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	if controllerutil.ContainsFinalizer(&sp, appmanagerFinalizer) {
		controllerutil.RemoveFinalizer(&sp, appmanagerFinalizer)
		if err := r.Update(ctx, &sp); err != nil {
			log.Error(err, fmt.Sprintf("Failed to remove Savepoint %s finalizers\n", sp.Spec.Metadata.Id))
			return err
		}
		return nil
	}
	return nil
}

func (r *SavepointReconciler) handleSavepointDeletion(sp appmanagervvpv1alpha1.Savepoint) error {
	ctx := context.Background()
	// name of our custom finalizer
	log := log.FromContext(ctx)
	err, SavepointExists := r.vvpClient.SavePoints().ResourceExistsInVVP(&sp)
	if err != nil {
		log.Error(err, "unable to check whether vvp Savepoint exists")
		return err
	}
	if !SavepointExists {
		log.Info(fmt.Sprintf("Savepoint %s doesnt exist in vvp, skipping deletion\n", sp.Spec.Metadata.Id))
		return nil
	}
	// The object is being deleted
	log.Info(fmt.Sprintf("Deleting Savepoint %s\n", sp.Spec.Metadata.Id))
	if controllerutil.ContainsFinalizer(&sp, appmanagerFinalizer) {
		// our finalizer is present, so lets handle any external spendency
		forceDelete := false
		if err := r.vvpClient.SavePoints().DeleteExternalResources(&sp, forceDelete); err != nil {
			log.Error(err, fmt.Sprintf("Failed to delete Savepoint %s in vvp, retrying...\n", sp.Spec.Metadata.Id))
			// if fail to delete the external spendency here, return with error
			// so that it can be retried
			return vvp_client.NewRetryableError(err)
		}

		// remove our finalizer from the list and update it.
		r.detachFinalizers(sp)
	}
	// Stop reconciliation as the item is being deleted
	return nil
}

func (r *SavepointReconciler) handleOutOfSyncError(sp appmanagervvpv1alpha1.Savepoint, err error) (ctrl.Result, error) {

	if updateErr := r.setStatus(sp, vvp_client.FormatOutOfSync(err)); updateErr != nil {
		return ctrl.Result{}, updateErr
	}
	if errors.Is(err, vvp_client.ErrRetryable) {
		return ctrl.Result{RequeueAfter: time.Second * 30, Requeue: true}, err
	}
	return ctrl.Result{}, err
}

func (r *SavepointReconciler) setStatus(sp appmanagervvpv1alpha1.Savepoint, syncState string) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	log.Info(fmt.Sprintf("Updating status for Savepoint %s \n", sp.Spec.Metadata.Id))
	sp.Status.State = syncState
	sp.Status.LastSync = metav1.Now()
	if err := r.Status().Update(ctx, &sp); err != nil {
		log.Error(err, "unable to update k8s Savepoint status")
		return err
	}
	return nil
}

func (r *SavepointReconciler) updateSavepointSpecInVVP(sp appmanagervvpv1alpha1.Savepoint) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	// update vvp Savepoint spec from k8s
	log.Info(fmt.Sprintf("Updating spec for Savepoint %s \n", sp.Spec.Metadata.Id))
	if err := r.vvpClient.SavePoints().UpdateExternalResources(&sp); err != nil {
		log.Error(err, "unable to update vvp Savepoint spec")
		return nil
	}
	return nil

}

// SetupWithManager sets up the controller with the Manager.
func (r *SavepointReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.vvpClient = vvp_client.NewClient()
	return ctrl.NewControllerManagedBy(mgr).
		For(&appmanagervvpv1alpha1.Savepoint{}).
		Complete(r)
}
