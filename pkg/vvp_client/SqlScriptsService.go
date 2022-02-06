package vvp_client

import (
	"context"
	"errors"
	"net/http"

	platformvvpv1alpha1 "efrat19.io/vvp-gitops-operator/apis/platform.vvp/v1alpha1"
	platform_apis "efrat19.io/vvp-gitops-operator/pkg/platform_apis"
)

type SqlScriptsService struct {
	client *platform_apis.APIClient
}

func (c SqlScriptsService) ResourceExistsInVVP(d *platformvvpv1alpha1.SqlScript) (error, bool) {
	ctx := context.Background()
	_, response, err := c.client.SqlControllerApi.GetSqlScriptUsingGET(ctx, CommunityEditionNamespace, d.Spec.Name)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil, false
		}
		return err, false
	}
	// SqlScript exists
	return err, true
}

func (c SqlScriptsService) DeleteExternalResources(d *platformvvpv1alpha1.SqlScript, force bool) error {
	ctx := context.Background()
	_, response, err := c.client.SqlControllerApi.DeleteSqlScriptUsingDELETE(ctx, CommunityEditionNamespace, d.Spec.Name)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return err
}

func (c SqlScriptsService) CreateExternalResources(d *platformvvpv1alpha1.SqlScript) error {
	ctx := context.Background()
	at := c.vvpAtFromK8sAt(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.SqlControllerApi.CreateSqlScriptUsingPOST(ctx, CommunityEditionNamespace, *at)
	return err
}

func (c SqlScriptsService) validateName(d *platformvvpv1alpha1.SqlScript) error {
	if d.Spec.Name != d.Name {
		msg := "Validation failed: SqlScript name must match SqlScript.spec.name"
		return errors.New(msg)
	}
	return nil
}

func (c SqlScriptsService) vvpAtFromK8sAt(d *platformvvpv1alpha1.SqlScript) *platform_apis.SqlScript {
	at := &platform_apis.SqlScript{
		CreateTime:  d.Spec.CreateTime,
		Description: d.Spec.Description,
		DisplayName: d.Spec.DisplayName,
		Name:        d.Spec.Name,
		Script:      d.Spec.Script,
		UpdateTime:  d.Spec.UpdateTime,
	}
	return at
}

func (c SqlScriptsService) UpdateExternalResources(d *platformvvpv1alpha1.SqlScript) error {
	ctx := context.Background()
	at := c.vvpAtFromK8sAt(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.SqlControllerApi.UpdateSqlScriptUsingPUT(ctx, CommunityEditionNamespace, *at, d.Spec.Name)
	return err
}
