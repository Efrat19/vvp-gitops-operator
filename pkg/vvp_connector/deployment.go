package vvp_connector

import (
	"context"
	"errors"
	"net/http"

	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/api/v1alpha1"
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (c *VvpConnector) DeploymentExistsInVVP(d *appmanagervvpv1alpha1.Deployment) (error, bool) {
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

func (c *VvpConnector) DeleteExternalResources(d *appmanagervvpv1alpha1.Deployment) error {
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

func (c *VvpConnector) CreateExternalResources(d *appmanagervvpv1alpha1.Deployment) error {
	ctx := context.Background()
	deployment := c.vvpDeplomentFromK8sDeployment(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.DeploymentResourceApi.CreateDeploymentUsingPOST(ctx, *deployment, d.Spec.Metadata.Namespace)
	return err
}

func (c *VvpConnector) UpdateExternalResources(d *appmanagervvpv1alpha1.Deployment) error {
	ctx := context.Background()
	deployment := c.vvpDeplomentFromK8sDeployment(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.DeploymentResourceApi.UpdateDeploymentUsingPATCH(ctx, *deployment, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
	return err
}

func (c *VvpConnector) GetStatus(d *appmanagervvpv1alpha1.Deployment) (*appmanager_apis.DeploymentStatus, error) {
	ctx := context.Background()
	deployment, _, err := c.client.DeploymentResourceApi.GetDeploymentUsingGET(ctx, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
	return deployment.Status, err
}

func (c *VvpConnector) validateName(d *appmanagervvpv1alpha1.Deployment) error {
	if d.Spec.Metadata.Name != d.Name {
		msg := "Validation failed: Deployment name must match deployment.spec.metadata.name"
		return errors.New(msg)
	}
	return nil
}

func (c *VvpConnector) vvpDeplomentFromK8sDeployment(d *appmanagervvpv1alpha1.Deployment) *appmanager_apis.Deployment {
	deployment := &appmanager_apis.Deployment{
		ApiVersion: "v1",
		Kind:       "Deployment",
		Metadata:   &d.Spec.Metadata,
		Spec:       &d.Spec.Spec,
		Status:     &d.Spec.Status,
	}
	return deployment
}

func (c *VvpConnector) cancelStateForDeletion(d *appmanagervvpv1alpha1.Deployment) error {
	cancelledState := "CANCELLED"
	status, err := c.GetStatus(d)
	if err != nil {
		return err
	}
	if status.State != cancelledState {
		d.Spec.Status.State = cancelledState
		err := c.UpdateExternalResources(d)
		return err
	}
	return nil
}
