package vvp_client

import (
	"context"
	"errors"
	"net/http"

	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/apis/appmanager.vvp/v1alpha1"
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
	"github.com/antihax/optional"
)

type SavePointsService struct {
	client *appmanager_apis.APIClient
}

func (c SavePointsService) ResourceExistsInVVP(d *appmanagervvpv1alpha1.Savepoint) (error, bool) {
	ctx := context.Background()
	_, response, err := c.client.SavepointResourceApi.GetSavepointUsingGET(ctx, CommunityEditionNamespace, d.Spec.Metadata.Id)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil, false
		}
		return err, false
	}
	// savepoint exists
	return err, true
}

func (c SavePointsService) DeleteExternalResources(d *appmanagervvpv1alpha1.Savepoint, force bool) error {
	ctx := context.Background()
	options := &appmanager_apis.SavepointResourceApiDeleteSavepointUsingDELETEOpts{
		Force: optional.NewBool(force),
	}
	response, err := c.client.SavepointResourceApi.DeleteSavepointUsingDELETE(ctx, CommunityEditionNamespace, d.Spec.Metadata.Id, options)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return err
}

func (c SavePointsService) CreateExternalResources(d *appmanagervvpv1alpha1.Savepoint) error {
	ctx := context.Background()
	sp := c.vvpSpFromK8sSp(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.SavepointResourceApi.CreateSavepointUsingPOST(ctx, CommunityEditionNamespace, *sp)
	return err
}

func (c SavePointsService) validateName(d *appmanagervvpv1alpha1.Savepoint) error {
	if d.Spec.Metadata.Id != d.Name {
		msg := "Validation failed: SavePoint name must match savepoint.spec.metadata.id"
		return errors.New(msg)
	}
	return nil
}

func (c SavePointsService) vvpSpFromK8sSp(d *appmanagervvpv1alpha1.Savepoint) *appmanager_apis.Savepoint {
	sp := &appmanager_apis.Savepoint{
		ApiVersion: "v1",
		Kind:       "Savepoint",
		Metadata:   &d.Spec.Metadata,
		Spec:       &d.Spec.Spec,
	}
	return sp
}
