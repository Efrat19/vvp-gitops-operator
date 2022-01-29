package vvp_connector

import (
	"context"
	"net/http"

	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/api/v1alpha1"
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
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
	deployment := appmanager_apis.Deployment{
		ApiVersion: "v1",
		Kind:       "Deployment",
		Metadata:   &d.Spec.Metadata,
		Spec:       &d.Spec.Spec,
		Status:     &d.Spec.Status,
	}
	// spew.Dump(deployment)
	_, _, err := c.client.DeploymentResourceApi.CreateDeploymentUsingPOST(ctx, deployment, d.Spec.Metadata.Namespace)
	// spew.Dump(response)
	return err
}

func (c *VvpConnector) UpdateExternalResources(d *appmanagervvpv1alpha1.Deployment) error {
	ctx := context.Background()
	deployment := &appmanager_apis.Deployment{
		ApiVersion: "v1",
		Kind:       "Deployment",
		Metadata:   &d.Spec.Metadata,
		Spec:       &d.Spec.Spec,
		Status:     &d.Spec.Status,
	}
	_, _, err := c.client.DeploymentResourceApi.UpdateDeploymentUsingPATCH(ctx, *deployment, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
	return err
}

func (c *VvpConnector) GetStatus(d *appmanagervvpv1alpha1.Deployment) (*appmanager_apis.DeploymentStatus, error) {
	ctx := context.Background()
	deployment, _, err := c.client.DeploymentResourceApi.GetDeploymentUsingGET(ctx, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
	return deployment.Status, err
}
