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

package platformvvp

import (
	"context"
	"efrat19.io/vvp-gitops-operator/pkg/vvp_client"
	"errors"
	"fmt"
	"time"

	// "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/polling"
	// "github.com/davecgh/go-spew/spew"
	platformvvpv1alpha1 "efrat19.io/vvp-gitops-operator/apis/platform.vvp/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ConnectorReconciler reconciles a Connector object
type ConnectorReconciler struct {
	client.Client
	Scheme    *runtime.Scheme
	vvpClient vvp_client.VvpClient
}

//+kubebuilder:rbac:groups=platform.vvp.efrat19.io,resources=connectors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=platform.vvp.efrat19.io,resources=connectors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=platform.vvp.efrat19.io,resources=connectors/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Connector object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *ConnectorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	var sp platformvvpv1alpha1.Connector
	if err := r.Get(ctx, req.NamespacedName, &sp); err != nil {
		log.Error(err, "unable to get Connector")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if err := r.vvpClient.MatchServerVersion(); err != nil {
		return r.handleOutOfSyncError(sp, err)
	}
	if err := r.handleConnectorFinalizers(sp); err != nil {
		return r.handleOutOfSyncError(sp, err)
	}
	// if the Connector needs to be deleted
	if !sp.ObjectMeta.DeletionTimestamp.IsZero() {
		if err := r.handleConnectorDeletion(sp); err != nil {
			return r.handleOutOfSyncError(sp, err)
		}
		return ctrl.Result{}, nil
	}
	if err := r.handleConnectorCreationIfNeeded(&sp); err != nil {
		return r.handleOutOfSyncError(sp, err)
	}
	if err := r.updateConnectorSpecInVVP(sp); err != nil {
		return r.handleOutOfSyncError(sp, err)
	}
	if err := r.setStatus(sp, vvp_client.InSyncState); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *ConnectorReconciler) handleConnectorCreationIfNeeded(sp *platformvvpv1alpha1.Connector) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	// Create Connector if not exists
	err, ConnectorExists := r.vvpClient.Connectors().ResourceExistsInVVP(sp)
	if err != nil {
		log.Error(err, "unable to check whether vvp Connector exists")
		return err
	}
	if !ConnectorExists {
		log.Info(fmt.Sprintf("Connector %s doesnt exist in vvp, attempting to create\n", sp.Spec.Name))
		if err := r.vvpClient.Connectors().CreateExternalResources(sp); err != nil {
			log.Error(err, "unable to create vvp Connector")
		}
	}
	return nil
}

func (r *ConnectorReconciler) handleConnectorFinalizers(sp platformvvpv1alpha1.Connector) error {
	ctx := context.Background()
	// name of our custom finalizer
	log := log.FromContext(ctx)

	// examine DeletionTimestamp to determine if object is under deletion
	if sp.ObjectMeta.DeletionTimestamp.IsZero() {
		log.Info(fmt.Sprintf("Attaching finalizers to Connector %s\n", sp.Spec.Name))
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// registering our finalizer.
		if !controllerutil.ContainsFinalizer(&sp, platformFinalizer) {
			controllerutil.AddFinalizer(&sp, platformFinalizer)
			if err := r.Update(ctx, &sp); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *ConnectorReconciler) handleConnectorDeletion(sp platformvvpv1alpha1.Connector) error {
	ctx := context.Background()
	// name of our custom finalizer
	log := log.FromContext(ctx)

	// The object is being deleted
	log.Info(fmt.Sprintf("Deleting Connector %s\n", sp.Spec.Name))
	if controllerutil.ContainsFinalizer(&sp, platformFinalizer) {
		// our finalizer is present, so lets handle any external spendency
		forceDelete := false
		if err := r.vvpClient.Connectors().DeleteExternalResources(&sp, forceDelete); err != nil {
			log.Error(err, fmt.Sprintf("Failed to delete Connector %s in vvp, retrying...\n", sp.Spec.Name))
			// if fail to delete the external spendency here, return with error
			// so that it can be retried
			return vvp_client.NewRetryableError(err)
		}

		// remove our finalizer from the list and update it.
		controllerutil.RemoveFinalizer(&sp, platformFinalizer)
		if err := r.Update(ctx, &sp); err != nil {
			log.Error(err, fmt.Sprintf("Failed to remove Connector %s finalizers\n", sp.Spec.Name))
			return err
		}
	}
	// Stop reconciliation as the item is being deleted
	return nil
}

func (r *ConnectorReconciler) handleOutOfSyncError(sp platformvvpv1alpha1.Connector, err error) (ctrl.Result, error) {
	if updateErr := r.setStatus(sp, vvp_client.FormatOutOfSync(err)); updateErr != nil {
		return ctrl.Result{}, updateErr
	}
	if errors.Is(err, vvp_client.ErrRetryable) {
		return ctrl.Result{RequeueAfter: time.Second * 30, Requeue: true}, err
	}
	return ctrl.Result{}, err
}

func (r *ConnectorReconciler) setStatus(sp platformvvpv1alpha1.Connector, syncState string) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	log.Info(fmt.Sprintf("Updating status for Connector %s \n", sp.Spec.Name))
	sp.Status.State = syncState
	sp.Status.LastSync = metav1.Now()
	if err := r.Status().Update(ctx, &sp); err != nil {
		log.Error(err, "unable to update k8s Connector status")
		return err
	}
	return nil
}

func (r *ConnectorReconciler) updateConnectorSpecInVVP(sp platformvvpv1alpha1.Connector) error {
	ctx := context.Background()
	log := log.FromContext(ctx)
	// update vvp Connector spec from k8s
	log.Info(fmt.Sprintf("Updating spec for Connector %s \n", sp.Spec.Name))
	if err := r.vvpClient.Connectors().UpdateExternalResources(&sp); err != nil {
		log.Error(err, "unable to update vvp Connector spec")
		return nil
	}
	return nil

}

// SetupWithManager sets up the controller with the Manager.
func (r *ConnectorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.vvpClient = vvp_client.NewClient()
	return ctrl.NewControllerManagedBy(mgr).
		For(&platformvvpv1alpha1.Connector{}).
		Complete(r)
}
