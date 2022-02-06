package vvp_client

import (
	"context"
	"errors"
	"net/http"

	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/api/v1alpha1"
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
)

type SecretValuesService struct {
	client *appmanager_apis.APIClient
}

func (c SecretValuesService) ResourceExistsInVVP(d *appmanagervvpv1alpha1.SecretValue) (error, bool) {
	ctx := context.Background()
	_, response, err := c.client.SecretValueResourceApi.GetSecretValueUsingGET(ctx, d.Spec.Metadata.Namespace, d.Spec.Metadata.Id)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil, false
		}
		return err, false
	}
	// secretvalue exists
	return err, true
}

func (c SecretValuesService) DeleteExternalResources(d *appmanagervvpv1alpha1.SecretValue, force bool) error {
	ctx := context.Background()
    _, response, err := c.client.SecretValueResourceApi.DeleteSecretValueUsingDELETE(ctx, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return err
}

func (c SecretValuesService) CreateExternalResources(d *appmanagervvpv1alpha1.SecretValue) error {
	ctx := context.Background()
	sv := c.vvpSvFromK8sSv(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.SecretValueResourceApi.CreateSecretValueUsingPOST(ctx, d.Spec.Metadata.Namespace, *sv )
	return err
}

func (c SecretValuesService) validateName(d *appmanagervvpv1alpha1.SecretValue) error {
	if d.Spec.Metadata.Id != d.Name {
		msg := "Validation failed: SavePoint name must match secretvalue.spec.metadata.name"
		return errors.New(msg)
	}
	return nil
}


func (c SecretValuesService) vvpSvFromK8sSv(d *appmanagervvpv1alpha1.SecretValue) *appmanager_apis.SecretValue {
	sv := &appmanager_apis.SecretValue{
		ApiVersion: "v1",
		Kind:       "SecretValue",
		Metadata:   &d.Spec.Metadata,
		Spec:       &d.Spec.Spec,
	}
	return sv
}


func (c SecretValuesService) UpdateExternalResources(d *appmanagervvpv1alpha1.SecretValue) error {
	ctx := context.Background()
	sv := c.vvpSvFromK8sSv(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.SecretValueResourceApi.UpdateSecretValueUsingPATCH(ctx, *sv, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
	return err
}
