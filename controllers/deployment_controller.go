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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/api/v1alpha1"
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
)

// DeploymentReconciler reconciles a Deployment object
type DeploymentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
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
	cfg := appmanager_apis.NewConfiguration()
	apiClient := appmanager_apis.NewAPIClient(cfg)
	var deployment *appmanagervvpv1alpha1.Deployment

	if err := r.Get(ctx, req.NamespacedName, deployment); err != nil {
		log.Error(err, "unable to fetch deployment")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	// name of our custom finalizer
	appmanagerFinalizer := "appmanager.vvp.efrat19.io/finalizer"

	// examine DeletionTimestamp to determine if object is under deletion
	if deployment.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// registering our finalizer.
		if !controllerutil.ContainsFinalizer(deployment, appmanagerFinalizer) {
			controllerutil.AddFinalizer(deployment, appmanagerFinalizer)
			if err := r.Update(ctx, deployment); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		// The object is being deleted
		if controllerutil.ContainsFinalizer(deployment, appmanagerFinalizer) {
			// our finalizer is present, so lets handle any external dependency
			if err := r.deleteExternalResources(deployment, apiClient); err != nil {
				// if fail to delete the external dependency here, return with error
				// so that it can be retried
				return ctrl.Result{}, err
			}

			// remove our finalizer from the list and update it.
			controllerutil.RemoveFinalizer(deployment, appmanagerFinalizer)
			if err := r.Update(ctx, deployment); err != nil {
				return ctrl.Result{}, err
			}
		}

		// Stop reconciliation as the item is being deleted
		return ctrl.Result{}, nil
	}
	// Create deployment if not exists
	err, deploymentExists := r.deploymentExistsInVVP(deployment, apiClient)
	if err != nil {
		log.Error(err, "unable to check whether vvp deployment exists")
		return ctrl.Result{}, nil
	}
	if !deploymentExists {
		if err := r.createExternalResources(deployment, apiClient); err != nil {
			log.Error(err, "unable to create vvp deployment")
		}
	}
	// get k8s deployment status from vvp
	status, err := r.getStatus(deployment, apiClient)
	if err != nil {
		log.Error(err, "unable to get k8s deployment status")
		return ctrl.Result{}, nil
	}
	// update k8s deployment status from vvp
	deployment.Status.Phase = status
	if err := r.Status().Update(ctx, deployment); err != nil {
		log.Error(err, "unable to update k8s deployment status")
		return ctrl.Result{}, err
	}

	// update vvp deployment spec from vvp
	if err := r.updateExternalResources(deployment, apiClient); err != nil {
		log.Error(err, "unable to update vvp deployment spec")
		return ctrl.Result{}, nil
	}
	return ctrl.Result{}, nil
}

func (r *DeploymentReconciler) deploymentExistsInVVP(d *appmanagervvpv1alpha1.Deployment, c *appmanager_apis.APIClient) (error, bool) {
	ctx := context.Background()
	_, response, err := c.DeploymentResourceApi.GetDeploymentUsingGET(ctx, d.Spec.Metadata.Id, d.Spec.Metadata.Namespace)
	if err != nil {
		if response.StatusCode == 404 {
			// deployment doesn't exists
			return nil, false
		}
		// other error occurred
		return err, false
	}
	// deployment exists
	return err, true
}

func (r *DeploymentReconciler) deleteExternalResources(d *appmanagervvpv1alpha1.Deployment, c *appmanager_apis.APIClient) error {
	ctx := context.Background()
	_, _, err := c.DeploymentResourceApi.DeleteDeploymentUsingDELETE(ctx, d.Spec.Metadata.Id, d.Spec.Metadata.Namespace)
	return err
}

func (r *DeploymentReconciler) createExternalResources(d *appmanagervvpv1alpha1.Deployment, c *appmanager_apis.APIClient) error {
	ctx := context.Background()
	deployment := &appmanager_apis.Deployment{
		ApiVersion: "v1",
		Kind:       "Deployment",
		Metadata:   &d.Spec.Metadata,
		Spec:       &d.Spec.Spec,
		Status:     &d.Spec.Status,
	}
	_, _, err := c.DeploymentResourceApi.CreateDeploymentUsingPOST(ctx, *deployment, d.Spec.Metadata.Namespace)
	return err
}

func (r *DeploymentReconciler) updateExternalResources(d *appmanagervvpv1alpha1.Deployment, c *appmanager_apis.APIClient) error {
	ctx := context.Background()
	deployment := &appmanager_apis.Deployment{
		ApiVersion: "v1",
		Kind:       "Deployment",
		Metadata:   &d.Spec.Metadata,
		Spec:       &d.Spec.Spec,
		Status:     &d.Spec.Status,
	}
	_, _, err := c.DeploymentResourceApi.UpdateDeploymentUsingPATCH(ctx, *deployment, d.Spec.Metadata.Id, d.Spec.Metadata.Namespace)
	return err
}

func (r *DeploymentReconciler) getStatus(d *appmanagervvpv1alpha1.Deployment, c *appmanager_apis.APIClient) (string, error) {
	ctx := context.Background()
	deployment, _, err := c.DeploymentResourceApi.GetDeploymentUsingGET(ctx, d.Spec.Metadata.Id, d.Spec.Metadata.Namespace)
	return deployment.Status.State, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *DeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appmanagervvpv1alpha1.Deployment{}).
		Complete(r)
}
