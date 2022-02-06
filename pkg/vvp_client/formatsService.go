package vvp_client

import (
	"context"
	"errors"
	"net/http"

	platformvvpv1alpha1 "efrat19.io/vvp-gitops-operator/apis/platform.vvp/v1alpha1"
	platform_apis "efrat19.io/vvp-gitops-operator/pkg/platform_apis"
)

type FormatsService struct {
	client *platform_apis.APIClient
}

func (c FormatsService) ResourceExistsInVVP(d *platformvvpv1alpha1.Format) (error, bool) {
	ctx := context.Background()
	_, response, err := c.client.ConnectorControllerApi.GetFormatUsingGET(ctx, d.Spec.Name, CommunityEditionNamespace)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil, false
		}
		return err, false
	}
	// Format exists
	return err, true
}

func (c FormatsService) DeleteExternalResources(d *platformvvpv1alpha1.Format, force bool) error {
	ctx := context.Background()
	_, response, err := c.client.ConnectorControllerApi.DeleteFormatUsingDELETE(ctx, CommunityEditionNamespace, d.Spec.Name)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return err
}

func (c FormatsService) CreateExternalResources(d *platformvvpv1alpha1.Format) error {
	ctx := context.Background()
	at := c.vvpAtFromK8sAt(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.ConnectorControllerApi.CreateFormatUsingPOST(ctx, *at, CommunityEditionNamespace)
	return err
}

func (c FormatsService) validateName(d *platformvvpv1alpha1.Format) error {
	if d.Spec.Name != d.Name {
		msg := "Validation failed: Format name must match Format.spec.name"
		return errors.New(msg)
	}
	return nil
}

func (c FormatsService) vvpAtFromK8sAt(d *platformvvpv1alpha1.Format) *platform_apis.Format {
	at := &platform_apis.Format{
		Dependencies: d.Spec.Dependencies,
		Name:         d.Spec.Name,
		Packaged:     d.Spec.Packaged,
		Properties:   d.Spec.Properties,
		Sink:         d.Spec.Sink,
		Source:       d.Spec.Source,
		Type_:        d.Spec.Type_,
	}
	return at
}

func (c FormatsService) UpdateExternalResources(d *platformvvpv1alpha1.Format) error {
	ctx := context.Background()
	at := c.vvpAtFromK8sAt(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.ConnectorControllerApi.UpdateFormatUsingPUT(ctx, *at, d.Spec.Name, CommunityEditionNamespace)
	return err
}
