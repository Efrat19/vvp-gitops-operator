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
	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/api/v1alpha1"
	"efrat19.io/vvp-gitops-operator/pkg/vvp_connector"
	"fmt"
	// "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/polling"
	// "github.com/davecgh/go-spew/spew"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// DeploymentReconciler reconciles a Deployment object
type DeploymentReconciler struct {
	client.Client
	Scheme       *runtime.Scheme
	vvpConnector *vvp_connector.VvpConnector
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
		log.Error(err, "unebale to get deployment")
		// log.Error(err, fmt.Sprintf("xxxxxxx %v", reflect.TypeOf(dep)))
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	// name of our custom finalizer
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
				return ctrl.Result{}, err
			}
		}
	} else {
		// The object is being deleted
		if controllerutil.ContainsFinalizer(&dep, appmanagerFinalizer) {
			// our finalizer is present, so lets handle any external dependency
			if err := r.vvpConnector.DeleteExternalResources(&dep); err != nil {
				log.Error(err, fmt.Sprintf("Failed to delete deployment %s in vvp\n", dep.Spec.Metadata.Name))
				// if fail to delete the external dependency here, return with error
				// so that it can be retried
				return ctrl.Result{}, err
			}

			// remove our finalizer from the list and update it.
			controllerutil.RemoveFinalizer(&dep, appmanagerFinalizer)
			if err := r.Update(ctx, &dep); err != nil {
				log.Error(err, fmt.Sprintf("Failed to remove deployment %s finalizers\n", dep.Spec.Metadata.Name))
				return ctrl.Result{}, err
			}
		}

		log.Info(fmt.Sprintf("Deployment %s deleted\n", dep.Spec.Metadata.Name))
		// Stop reconciliation as the item is being deleted
		return ctrl.Result{}, nil
	}
	// Create deployment if not exists
	err, deploymentExists := r.vvpConnector.DeploymentExistsInVVP(&dep)
	if err != nil {
		log.Error(err, "unable to check whether vvp deployment exists")
		return ctrl.Result{}, nil
	}
	if !deploymentExists {
		log.Info(fmt.Sprintf("Deployment %s doesnt exist in vvp, attempting to create\n", dep.Spec.Metadata.Name))
		if err := r.vvpConnector.CreateExternalResources(&dep); err != nil {
			log.Error(err, "unable to create vvp deployment")
		}
	}
	r.updateDeploymentStatus(ctx, &dep)
	// update vvp deployment spec from vvp
	log.Info(fmt.Sprintf("Updating spec for deployment %s \n", dep.Spec.Metadata.Name))
	if err := r.vvpConnector.UpdateExternalResources(&dep); err != nil {
		log.Error(err, "unable to update vvp deployment spec")
		return ctrl.Result{}, nil
	}
	return ctrl.Result{}, nil
}

func (r *DeploymentReconciler) updateDeploymentStatus(ctx context.Context, dep *appmanagervvpv1alpha1.Deployment) (ctrl.Result, error) {
	// update k8s deployment status from vvp
	log := log.FromContext(ctx)
	// get k8s deployment status from vvp
	log.Info(fmt.Sprintf("Getting status for deployment %s \n", dep.Spec.Metadata.Name))
	status, err := r.vvpConnector.GetStatus(dep)
	if err != nil {
		log.Error(err, "unable to get k8s deployment status")
		return ctrl.Result{}, nil
	}
	// update k8s deployment status from vvp
	log.Info(fmt.Sprintf("Updating status for deployment %s \n", dep.Spec.Metadata.Name))
	dep.Status.State = status.State
	dep.Status.Running = status.Running
	if err := r.Status().Update(ctx, dep); err != nil {
		log.Error(err, "unable to update k8s deployment status")
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.vvpConnector, _ = vvp_connector.NewConnector()
	return ctrl.NewControllerManagedBy(mgr).
		For(&appmanagervvpv1alpha1.Deployment{}).
		Complete(r)
}
