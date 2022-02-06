package vvp_client

import (
	"context"
	"errors"
	"net/http"

	appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/api/v1alpha1"
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
)

type SessionClustersService struct {
	client *appmanager_apis.APIClient
}

func (c SessionClustersService) ResourceExistsInVVP(d *appmanagervvpv1alpha1.SessionCluster) (error, bool) {
	ctx := context.Background()
	_, response, err := c.client.SessionClusterResourceApi.GetSessionClusterUsingGET(ctx, d.Spec.Metadata.Namespace, d.Spec.Metadata.Id)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil, false
		}
		return err, false
	}
	// savepoint exists
	return err, true
}

func (c SessionClustersService) DeleteExternalResources(d *appmanagervvpv1alpha1.SessionCluster, force bool) error {
	ctx := context.Background()
    _, response, err := c.client.SessionClusterResourceApi.DeleteSessionClusterUsingDELETE(ctx, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
	if err != nil {
		if response.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return err
}

func (c SessionClustersService) CreateExternalResources(d *appmanagervvpv1alpha1.SessionCluster) error {
	ctx := context.Background()
	sc := c.vvpScFromK8sSc(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.SessionClusterResourceApi.CreateSessionClusterUsingPOST(ctx, d.Spec.Metadata.Namespace, *sc )
	return err
}

func (c SessionClustersService) validateName(d *appmanagervvpv1alpha1.SessionCluster) error {
	if d.Spec.Metadata.Id != d.Name {
		msg := "Validation failed: SavePoint name must match SessionCluster.spec.metadata.name"
		return errors.New(msg)
	}
	return nil
}


func (c SessionClustersService) vvpScFromK8sSc(d *appmanagervvpv1alpha1.SessionCluster) *appmanager_apis.SessionCluster {
	sc := &appmanager_apis.SessionCluster{
		ApiVersion: "v1",
		Kind:       "SessionCluster",
		Metadata:   &d.Spec.Metadata,
		Spec:       &d.Spec.Spec,
	}
	return sc
}


func (c SessionClustersService) UpdateExternalResources(d *appmanagervvpv1alpha1.SessionCluster) error {
	ctx := context.Background()
	sc := c.vvpScFromK8sSc(d)
	if err := c.validateName(d); err != nil {
		return err
	}
	_, _, err := c.client.SessionClusterResourceApi.UpdateSessionClusterUsingPATCH(ctx, *sc, d.Spec.Metadata.Name, d.Spec.Metadata.Namespace)
	return err
}
