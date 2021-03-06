package vvp_client

import (
	"context"
	"errors"
	"net/http"

	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/apis/appmanager.vvp/v1alpha1"
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"

	// "github.com/davecgh/go-spew/spew"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type DeploymentsService struct {
	client *appmanager_apis.APIClient
}

func (c DeploymentsService) ResourceExistsInVVP(d *appmanagervvpv1alpha1.Deployment) (error, bool) {
	ctx := context.Background()
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

func (c DeploymentsService) DeleteExternalResources(d *appmanagervvpv1alpha1.Deployment) error {
	ctx := context.Background()
	if err := c.cancelStateForDeletion(d); err != nil {
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

func (c DeploymentsService) CreateExternalResources(d *appmanagervvpv1alpha1.Deployment) error {
	ctx := context.Background()
	deployment := c.vvpDeplomentFromK8sDeployment(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.DeploymentResourceApi.CreateDeploymentUsingPOST(ctx, *deployment, CommunityEditionNamespace)
	return err
}

func (c DeploymentsService) UpdateExternalResources(d *appmanagervvpv1alpha1.Deployment) error {
	ctx := context.Background()
	deployment := c.vvpDeplomentFromK8sDeployment(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.DeploymentResourceApi.UpdateDeploymentUsingPATCH(ctx, *deployment, d.Spec.Metadata.Name, CommunityEditionNamespace)
	return err
}

func (c DeploymentsService) GetStatus(d *appmanagervvpv1alpha1.Deployment) (*appmanager_apis.DeploymentStatus, error) {
	ctx := context.Background()
	deployment, _, err := c.client.DeploymentResourceApi.GetDeploymentUsingGET(ctx, d.Spec.Metadata.Name, CommunityEditionNamespace)
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

func (c DeploymentsService) cancelStateForDeletion(d *appmanagervvpv1alpha1.Deployment) error {
	cancelledState := "CANCELLED"
	status, err := c.GetStatus(d)
	if err != nil {
		return err
	}
	if status.State != cancelledState {
		d.Spec.Spec.State = cancelledState
		err := c.UpdateExternalResources(d)
		return err
	}
	return nil
}
