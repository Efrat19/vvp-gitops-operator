package vvp_client

import (
	// appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/api/v1alpha1"
	"context"
	"errors"
	"fmt"
	"net/http"

	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
	"sigs.k8s.io/controller-runtime/pkg/log"

	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	// "github.com/gophercloud/gophercloud/openstack/cdn/v1/base"
)

type VvpClient interface {
	ProbeServer() error
	MatchServerVersion() error
	Deployments() DeploymentsService
	DeploymentTargets() DeploymentTargetsService
	SavePoints() SavePointsService
	SecretValues() SecretValuesService
	SessionClusters() SessionClustersService
}

type vvpClient struct {
	appManagerClient   *appmanager_apis.APIClient
	DeploymentsService *DeploymentsService
	DeploymentTargetsService *DeploymentTargetsService
	SavePointsService *SavePointsService
	SecretValuesService *SecretValuesService
	SessionClustersService *SessionClustersService
}

func NewClient() VvpClient {
	return &vvpClient{
		appManagerClient: NewAppManagerClient(),
	}
}

func (v *vvpClient) Deployments() DeploymentsService {
	if v.DeploymentsService == nil {
		v.DeploymentsService = &DeploymentsService{client: v.appManagerClient}
	}
	return *v.DeploymentsService
}

func (v *vvpClient) DeploymentTargets() DeploymentTargetsService {
	if &v.DeploymentTargetsService == nil {
		v.DeploymentTargetsService = &DeploymentTargetsService{client: v.appManagerClient}
	}
	return *v.DeploymentTargetsService
}

func (v *vvpClient) SavePoints() SavePointsService {
	if &v.SavePointsService == nil {
		v.SavePointsService = &SavePointsService{client: v.appManagerClient}
	}
	return *v.SavePointsService
}

func (v *vvpClient) SecretValues() SecretValuesService {
	if &v.SecretValuesService == nil {
		v.SecretValuesService = &SecretValuesService{client: v.appManagerClient}
	}
	return *v.SecretValuesService
}

func (v *vvpClient) SessionClusters() SessionClustersService {
	if &v.SessionClustersService == nil {
		v.SessionClustersService = &SessionClustersService{client: v.appManagerClient}
	}
	return *v.SessionClustersService
}


func NewAppManagerClient() *appmanager_apis.APIClient {
	basePath := getEnv("VVP_URL", "http://vvp.data.yad2.io")
	cfg := &appmanager_apis.Configuration{
		BasePath:      basePath,
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	return appmanager_apis.NewAPIClient(cfg)
}

func (v *vvpClient) ProbeServer() error {
	ctx := context.Background()
	_, response, err := v.appManagerClient.StatusResourceApi.GetStatusUsingGET(ctx)
	if response.StatusCode != http.StatusOK {
		log := log.FromContext(ctx)
		log.Error(err, "error failed to connect")
		return v.ConnectionFailedError()
	}
	return nil
}

func (v *vvpClient) MatchServerVersion() error {
	ctx := context.Background()
	supportedVersion := "2.6.1"
	var si appmanager_apis.SystemInformation
	si, _, err := v.appManagerClient.StatusResourceApi.GetSystemInfoUsingGET(ctx)
	if err != nil {
		return v.ConnectionFailedError()
	}
	serverVersion := si.Status.RevisionInformation.BuildVersion
	if serverVersion != supportedVersion {
		return errors.New(fmt.Sprintf("Error Version Mismatch - vvp server version is %s but the operator only supports %s", serverVersion, supportedVersion))
	}
	return nil
}

func (v *vvpClient) ConnectionFailedError() error {
	return NewRetryableError(errors.New(fmt.Sprintf("Failed to connect to vvp server at %s, retrying", v.appManagerClient.GetBasePath())))

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
