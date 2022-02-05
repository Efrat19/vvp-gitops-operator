package vvp_client

// import (
// 	"context"
// 	"errors"
// 	"net/http"

// 	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/api/v1alpha1"
// 	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
// 	"sigs.k8s.io/controller-runtime/pkg/log"
// )

// type DeploymentTargetsService struct {
// 	client *appmanager_apis.APIClient
// }

// func (c DeploymentTargetsService) ResourceExistsInVVP(d *appmanagervvpv1alpha1.DeploymentTarget) (error, bool) {
// 	ctx := context.Background()
// 	_, response, err := c.client.DeploymentTargetResourceApi.GetDeploymentTargetsUsingGET(ctx, d.Spec.Metadata.Namespace)
// 	if err != nil {
// 		if response.StatusCode == http.StatusNotFound {
// 			return nil, false
// 		}
// 		return err, false
// 	}
// 	// deployment exists
// 	return err, true
// }

// func (c DeploymentTargetsService) DeleteExternalResources(d *appmanagervvpv1alpha1.DeploymentTarget) error {
// 	ctx := context.Background()
// 	if err := c.cancelStateForDeletion(d); err != nil {
// 		log := log.FromContext(ctx)
// 		log.Error(err, "Failed to cancel deployment for deletion")
// 		return err
// 	}
// 	_, response, err := c.client.DeploymentTargetResourceApi.DeleteDeploymentTargetsUsingDELETE(ctx, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
// 	if err != nil {
// 		if response.StatusCode == http.StatusNotFound {
// 			return nil
// 		}
// 		return err
// 	}
// 	return err
// }

// func (c DeploymentTargetsService) CreateExternalResources(d *appmanagervvpv1alpha1.DeploymentTargets) error {
// 	ctx := context.Background()
// 	deployment := c.vvpDeplomentFromK8sDeploymentDefault(d)
// 	if err := c.validateName(d); err != nil {
// 		return err
// 	}
// 	_, _, err := c.client.DeploymentTargetResourceApi.CreateDeploymentTargetsUsingPOST(ctx, *deployment, d.Spec.Metadata.Namespace)
// 	return err
// }

// func (c DeploymentTargetsService) UpdateExternalResources(d *appmanagervvpv1alpha1.DeploymentTarget) error {
// 	ctx := context.Background()
// 	deployment := c.vvpDeplomentFromK8sDeploymentDefault(d)
// 	if err := c.validateName(d); err != nil {
// 		return err
// 	}
// 	_, _, err := c.client.DeploymentTargetsResourceApi.UpdateDeploymentTargetsUsingPATCH(ctx, *deployment, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
// 	return err
// }

// func (c DeploymentTargetsService) GetStatus(d *appmanagervvpv1alpha1.DeploymentTarget) (*appmanager_apis.DeploymentTargetsStatus, error) {
// 	ctx := context.Background()
// 	deployment, _, err := c.client.DeploymentTargetsResourceApi.GetDeploymentTargetsUsingGET(ctx, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
// 	return deployment.Status, err
// }

// func (c DeploymentTargetsService) validateName(d *appmanagervvpv1alpha1.DeploymentTargets) error {
// 	if d.Spec.Metadata.Name != d.Name {
// 		msg := "Validation failed: DeploymentDefault name must match deployment.spec.metadata.name"
// 		return errors.New(msg)
// 	}
// 	return nil
// }

// func (c DeploymentTargetsService) vvpDeplomentFromK8sDeploymentDefault(d *appmanagervvpv1alpha1.DeploymentTargets) *appmanager_apis.DeploymentTargets {
// 	deployment := &appmanager_apis.DeploymentTargets{
// 		ApiVersion: "v1",
// 		Kind:       "DeploymentDefault",
// 		Metadata:   &d.Spec.Metadata,
// 		Spec:       &d.Spec.Spec,
// 		Status:     &d.Spec.Status,
// 	}
// 	return deployment
// }

// func (c DeploymentTargetsService) cancelStateForDeletion(d *appmanagervvpv1alpha1.DeploymentTargets) error {
// 	cancelledState := "CANCELLED"
// 	status, err := c.GetStatus(d)
// 	if err != nil {
// 		return err
// 	}
// 	if status.State != cancelledState {
// 		d.Spec.Status.State = cancelledState
// 		err := c.UpdateExternalResources(d)
// 		return err
// 	}
// 	return nil
// }
