package vvp_client

import (
	"context"
	"errors"
	"net/http"

	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/apis/appmanager.vvp/v1alpha1"
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
)

type DeploymentTargetsService struct {
	client *appmanager_apis.APIClient
}

func (c DeploymentTargetsService) ResourceExistsInVVP(d *appmanagervvpv1alpha1.DeploymentTarget) (error, bool) {
	ctx := context.Background()
	_, response, err := c.client.DeploymentTargetResourceApi.GetDeploymentTargetsUsingGET(ctx, CommunityEditionNamespace)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil, false
		}
		return err, false
	}
	// deployment exists
	return err, true
}

func (c DeploymentTargetsService) DeleteExternalResources(d *appmanagervvpv1alpha1.DeploymentTarget, force bool) error {
	ctx := context.Background()
	_, response, err := c.client.DeploymentTargetResourceApi.DeleteDeploymentTargetUsingDELETE(ctx, d.Spec.Metadata.Name, CommunityEditionNamespace)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return err
}

func (c DeploymentTargetsService) CreateExternalResources(d *appmanagervvpv1alpha1.DeploymentTarget) error {
	ctx := context.Background()
	dt := c.vvpDtFromK8sDt(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.DeploymentTargetResourceApi.CreateDeploymentTargetUsingPOST(ctx, *dt, CommunityEditionNamespace)
	return err
}

func (c DeploymentTargetsService) UpdateExternalResources(d *appmanagervvpv1alpha1.DeploymentTarget) error {
	ctx := context.Background()
	dt := c.vvpDtFromK8sDt(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.DeploymentTargetResourceApi.UpsertDeploymentTargetUsingPUT(ctx, *dt, d.Spec.Metadata.Name, CommunityEditionNamespace)
	return err
}

func (c DeploymentTargetsService) validateName(d *appmanagervvpv1alpha1.DeploymentTarget) error {
	if d.Spec.Metadata.Name != d.Name {
		msg := "Validation failed: DeploymentTarget name must match deployment.spec.metadata.name"
		return errors.New(msg)
	}
	return nil
}

func (c DeploymentTargetsService) vvpDtFromK8sDt(d *appmanagervvpv1alpha1.DeploymentTarget) *appmanager_apis.DeploymentTarget {
	dt := &appmanager_apis.DeploymentTarget{
		ApiVersion: "v1",
		Kind:       "DeploymentTarget",
		Metadata:   &d.Spec.Metadata,
		Spec:       &d.Spec.Spec,
	}
	return dt
}
