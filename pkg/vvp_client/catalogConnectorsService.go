package vvp_client

import (
	"context"
	"errors"
	"net/http"

	platformvvpv1alpha1 "efrat19.io/vvp-gitops-operator/apis/platform.vvp/v1alpha1"
	platform_apis "efrat19.io/vvp-gitops-operator/pkg/platform_apis"
)

type CatalogConnectorsService struct {
	client *platform_apis.APIClient
}

func (c CatalogConnectorsService) ResourceExistsInVVP(d *platformvvpv1alpha1.CatalogConnector) (error, bool) {
	ctx := context.Background()
	_, response, err := c.client.ConnectorControllerApi.GetCatalogConnectorUsingGET(ctx, d.Spec.Name, CommunityEditionNamespace)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil, false
		}
		return err, false
	}
	// CatalogConnector exists
	return err, true
}

func (c CatalogConnectorsService) DeleteExternalResources(d *platformvvpv1alpha1.CatalogConnector, force bool) error {
	ctx := context.Background()
	_, response, err := c.client.ConnectorControllerApi.DeleteCatalogConnectorUsingDELETE(ctx, CommunityEditionNamespace, d.Spec.Name)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return err
}

func (c CatalogConnectorsService) CreateExternalResources(d *platformvvpv1alpha1.CatalogConnector) error {
	ctx := context.Background()
	at := c.vvpAtFromK8sAt(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.ConnectorControllerApi.CreateCatalogConnectorUsingPOST(ctx, *at, CommunityEditionNamespace)
	return err
}

func (c CatalogConnectorsService) validateName(d *platformvvpv1alpha1.CatalogConnector) error {
	if d.Spec.Name != d.Name {
		msg := "Validation failed: CatalogConnector name must match CatalogConnector.spec.name"
		return errors.New(msg)
	}
	return nil
}

func (c CatalogConnectorsService) vvpAtFromK8sAt(d *platformvvpv1alpha1.CatalogConnector) *platform_apis.CatalogConnector {
	at := &platform_apis.CatalogConnector{
		Dependencies: d.Spec.Dependencies,
		Name:         d.Spec.Name,
		Packaged:     d.Spec.Packaged,
		Properties:   d.Spec.Properties,
		ReadOnly:     d.Spec.ReadOnly,
		Type_:        d.Spec.Type_,
	}
	return at
}

func (c CatalogConnectorsService) UpdateExternalResources(d *platformvvpv1alpha1.CatalogConnector) error {
	ctx := context.Background()
	at := c.vvpAtFromK8sAt(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.ConnectorControllerApi.UpdateCatalogConnectorUsingPUT(ctx, *at, d.Spec.Name, CommunityEditionNamespace)
	return err
}
