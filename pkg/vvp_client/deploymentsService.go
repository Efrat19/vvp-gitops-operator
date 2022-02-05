package vvp_client

import (
	"context"
	"errors"
	"net/http"

	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/api/v1alpha1"
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
	// "github.com/davecgh/go-spew/spew"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type DeploymentsService struct {
	client *appmanager_apis.APIClient
}

func (c DeploymentsService) ResourceExistsInVVP(ctx context.Context, d *appmanagervvpv1alpha1.Deployment) (error, bool) {
	_, response, err := c.client.DeploymentResourceApi.GetDeploymentUsingGET(ctx, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil, false
		}
		return err, false
	}
	// deployment exists
	return err, true
}

func (c DeploymentsService) DeleteExternalResources(ctx context.Context, d *appmanagervvpv1alpha1.Deployment) error {
	if err := c.cancelStateForDeletion(ctx, d); err != nil {
		log := log.FromContext(ctx)
		log.Error(err, "Failed to cancel deployment for deletion")
		return err
	}
	_, response, err := c.client.DeploymentResourceApi.DeleteDeploymentUsingDELETE(ctx, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return err
}

func (c DeploymentsService) CreateExternalResources(ctx context.Context, d *appmanagervvpv1alpha1.Deployment) error {
	deployment := c.vvpDeplomentFromK8sDeployment(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.DeploymentResourceApi.CreateDeploymentUsingPOST(ctx, *deployment, d.Spec.Metadata.Namespace)
	return err
}

func (c DeploymentsService) UpdateExternalResources(ctx context.Context, d *appmanagervvpv1alpha1.Deployment) error {
	deployment := c.vvpDeplomentFromK8sDeployment(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.DeploymentResourceApi.UpdateDeploymentUsingPATCH(ctx, *deployment, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
	return err
}

func (c DeploymentsService) GetStatus(ctx context.Context, d *appmanagervvpv1alpha1.Deployment) (*appmanager_apis.DeploymentStatus, error) {
	deployment, _, err := c.client.DeploymentResourceApi.GetDeploymentUsingGET(ctx, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
	return deployment.Status, err
}

func (c DeploymentsService) validateName(d *appmanagervvpv1alpha1.Deployment) error {
	if d.Spec.Metadata.Name != d.Name {
		msg := "Validation failed: Deployment name must match deployment.spec.metadata.name"
		return errors.New(msg)
	}
	return nil
}

func (c DeploymentsService) vvpDeplomentFromK8sDeployment(d *appmanagervvpv1alpha1.Deployment) *appmanager_apis.Deployment {
	deployment := &appmanager_apis.Deployment{
		ApiVersion: "v1",
		Kind:       "Deployment",
		Metadata:   &d.Spec.Metadata,
		Spec:       &d.Spec.Spec,
	}
	return deployment
}

func (c DeploymentsService) cancelStateForDeletion(ctx context.Context, d *appmanagervvpv1alpha1.Deployment) error {
	cancelledState := "CANCELLED"
	status, err := c.GetStatus(ctx, d)
	if err != nil {
		return err
	}
	if status.State != cancelledState {
		d.Spec.Spec.State = cancelledState
		err := c.UpdateExternalResources(ctx, d)
		return err
	}
	return nil
}
