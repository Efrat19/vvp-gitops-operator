package vvp_client

import (
	"context"
	"errors"
	"net/http"

	platformvvpv1alpha1 "efrat19.io/vvp-gitops-operator/apis/platform.vvp/v1alpha1"
	platform_apis "efrat19.io/vvp-gitops-operator/pkg/platform_apis"
	"github.com/antihax/optional"
)

type UdfArtifactsService struct {
	client *platform_apis.APIClient
}

func (c UdfArtifactsService) ResourceExistsInVVP(d *platformvvpv1alpha1.UdfArtifact, requireFunctionNames bool) (error, bool) {
	ctx := context.Background()
	options := &platform_apis.UdfArtifactControllerApiGetUdfArtifactUsingGETOpts{
		RequireFunctionNames: optional.NewBool(requireFunctionNames),
	}
	_, response, err := c.client.UdfArtifactControllerApi.GetUdfArtifactUsingGET(ctx, CommunityEditionNamespace, d.Spec.Name, options)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil, false
		}
		return err, false
	}
	// UdfArtifact exists
	return err, true
}

func (c UdfArtifactsService) DeleteExternalResources(d *platformvvpv1alpha1.UdfArtifact, force bool) error {
	ctx := context.Background()
	_, response, err := c.client.UdfArtifactControllerApi.DeleteUdfArtifactUsingDELETE(ctx, CommunityEditionNamespace, d.Spec.Name)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return err
}

func (c UdfArtifactsService) CreateExternalResources(d *platformvvpv1alpha1.UdfArtifact) error {
	ctx := context.Background()
	at := c.vvpAtFromK8sAt(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.UdfArtifactControllerApi.CreateUdfArtifactUsingPOST(ctx, CommunityEditionNamespace, *at)
	return err
}

func (c UdfArtifactsService) validateName(d *platformvvpv1alpha1.UdfArtifact) error {
	if d.Spec.Name != d.Name {
		msg := "Validation failed: UdfArtifact name must match UdfArtifact.spec.name"
		return errors.New(msg)
	}
	return nil
}

func (c UdfArtifactsService) vvpAtFromK8sAt(d *platformvvpv1alpha1.UdfArtifact) *platform_apis.UdfArtifact {
	at := &platform_apis.UdfArtifact{
		CreateTime:    d.Spec.CreateTime,
		JarUpdateTime: d.Spec.JarUpdateTime,
		JarUrl:        d.Spec.JarUrl,
		Name:          d.Spec.Name,
		UdfClasses:    d.Spec.UdfClasses,
		UpdateTime:    d.Spec.UpdateTime,
	}
	return at
}

func (c UdfArtifactsService) UpdateExternalResources(d *platformvvpv1alpha1.UdfArtifact) error {
	ctx := context.Background()
	at := c.vvpAtFromK8sAt(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.UdfArtifactControllerApi.UpdateUdfArtifactUsingPUT(ctx, CommunityEditionNamespace, *at, d.Spec.Name)
	return err
}
