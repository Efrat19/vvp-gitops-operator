package vvp_client

import (
	"context"
	"errors"
	"net/http"

	platformvvpv1alpha1 "efrat19.io/vvp-gitops-operator/apis/platform.vvp/v1alpha1"
	platform_apis "efrat19.io/vvp-gitops-operator/pkg/platform_apis"
)

type ConnectorsService struct {
	client *platform_apis.APIClient
}

func (c ConnectorsService) ResourceExistsInVVP(d *platformvvpv1alpha1.Connector) (error, bool) {
	ctx := context.Background()
	_, response, err := c.client.ConnectorControllerApi.GetConnectorUsingGET(ctx, d.Spec.Name, CommunityEditionNamespace)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil, false
		}
		return err, false
	}
	// Connector exists
	return err, true
}

func (c ConnectorsService) DeleteExternalResources(d *platformvvpv1alpha1.Connector, force bool) error {
	ctx := context.Background()
	_, response, err := c.client.ConnectorControllerApi.DeleteConnectorUsingDELETE(ctx, CommunityEditionNamespace, d.Spec.Name)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return err
}

func (c ConnectorsService) CreateExternalResources(d *platformvvpv1alpha1.Connector) error {
	ctx := context.Background()
	at := c.vvpAtFromK8sAt(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.ConnectorControllerApi.CreateConnectorUsingPOST(ctx, *at, CommunityEditionNamespace)
	return err
}

func (c ConnectorsService) validateName(d *platformvvpv1alpha1.Connector) error {
	if d.Spec.Name != d.Name {
		msg := "Validation failed: Connector name must match Connector.spec.name"
		return errors.New(msg)
	}
	return nil
}

func (c ConnectorsService) vvpAtFromK8sAt(d *platformvvpv1alpha1.Connector) *platform_apis.Connector {
	at := &platform_apis.Connector{
		Dependencies:     d.Spec.Dependencies,
		Lookup:           d.Spec.Lookup,
		Name:             d.Spec.Name,
		Packaged:         d.Spec.Packaged,
		Properties:       d.Spec.Properties,
		Sink:             d.Spec.Sink,
		Source:           d.Spec.Source,
		SupportedFormats: d.Spec.SupportedFormats,
		Type_:            d.Spec.Type_,
	}
	return at
}

func (c ConnectorsService) UpdateExternalResources(d *platformvvpv1alpha1.Connector) error {
	ctx := context.Background()
	at := c.vvpAtFromK8sAt(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.ConnectorControllerApi.UpdateConnectorUsingPUT(ctx, *at, d.Spec.Name, CommunityEditionNamespace)
	return err
}
