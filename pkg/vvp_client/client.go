package vvp_client

import (
	"context"
	"errors"
	"fmt"

	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
	platform_apis "efrat19.io/vvp-gitops-operator/pkg/platform_apis"
	"os"
)

const (
	supportedVersion string = "2.6.1"
)

type VvpClient interface {
	ProbeServer() error
	Deployments() DeploymentsService
	DeploymentTargets() DeploymentTargetsService
	SavePoints() SavePointsService
	SecretValues() SecretValuesService
	SessionClusters() SessionClustersService
	CatalogConnectors() CatalogConnectorsService
	Connectors() ConnectorsService
	Formats() FormatsService
	SqlScripts() SqlScriptsService
	UdfArtifacts() UdfArtifactsService
}

type vvpClient struct {
	appManagerClient         *appmanager_apis.APIClient
	platformClient           *platform_apis.APIClient
	DeploymentsService       *DeploymentsService
	DeploymentTargetsService *DeploymentTargetsService
	SavePointsService        *SavePointsService
	SecretValuesService      *SecretValuesService
	SessionClustersService   *SessionClustersService
	CatalogConnectorsService *CatalogConnectorsService
	ConnectorsService        *ConnectorsService
	FormatsService           *FormatsService
	SqlScriptsService        *SqlScriptsService
	UdfArtifactsService      *UdfArtifactsService
}

func NewClient() VvpClient {
	return &vvpClient{
		appManagerClient: NewAppManagerClient(),
		platformClient:   NewPlatformClient(),
	}
}

func (v *vvpClient) Deployments() DeploymentsService {
	if v.DeploymentsService == nil {
		v.DeploymentsService = &DeploymentsService{client: v.appManagerClient}
	}
	return *v.DeploymentsService
}

func (v *vvpClient) DeploymentTargets() DeploymentTargetsService {
	if v.DeploymentTargetsService == nil {
		v.DeploymentTargetsService = &DeploymentTargetsService{client: v.appManagerClient}
	}
	return *v.DeploymentTargetsService
}

func (v *vvpClient) SavePoints() SavePointsService {
	if v.SavePointsService == nil {
		v.SavePointsService = &SavePointsService{client: v.appManagerClient}
	}
	return *v.SavePointsService
}

func (v *vvpClient) SecretValues() SecretValuesService {
	if v.SecretValuesService == nil {
		v.SecretValuesService = &SecretValuesService{client: v.appManagerClient}
	}
	return *v.SecretValuesService
}

func (v *vvpClient) SessionClusters() SessionClustersService {
	if v.SessionClustersService == nil {
		v.SessionClustersService = &SessionClustersService{client: v.appManagerClient}
	}
	return *v.SessionClustersService
}

func (v *vvpClient) CatalogConnectors() CatalogConnectorsService {
	if v.CatalogConnectorsService == nil {
		v.CatalogConnectorsService = &CatalogConnectorsService{client: v.platformClient}
	}
	return *v.CatalogConnectorsService
}

func (v *vvpClient) Connectors() ConnectorsService {
	if v.ConnectorsService == nil {
		v.ConnectorsService = &ConnectorsService{client: v.platformClient}
	}
	return *v.ConnectorsService
}

func (v *vvpClient) Formats() FormatsService {
	if v.FormatsService == nil {
		v.FormatsService = &FormatsService{client: v.platformClient}
	}
	return *v.FormatsService
}

func (v *vvpClient) SqlScripts() SqlScriptsService {
	if v.SqlScriptsService == nil {
		v.SqlScriptsService = &SqlScriptsService{client: v.platformClient}
	}
	return *v.SqlScriptsService
}

func (v *vvpClient) UdfArtifacts() UdfArtifactsService {
	if v.UdfArtifactsService == nil {
		v.UdfArtifactsService = &UdfArtifactsService{client: v.platformClient}
	}
	return *v.UdfArtifactsService
}

func NewAppManagerClient() *appmanager_apis.APIClient {
	basePath := getEnv("VVP_URL", "http://localhost:8080")
	cfg := &appmanager_apis.Configuration{
		BasePath:      basePath,
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	return appmanager_apis.NewAPIClient(cfg)
}

func NewPlatformClient() *platform_apis.APIClient {
	basePath := getEnv("VVP_URL", "http://localhost:8080")
	cfg := &platform_apis.Configuration{
		BasePath:      basePath,
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	return platform_apis.NewAPIClient(cfg)
}

func (v *vvpClient) ProbeServer() error {
	ctx := context.Background()
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
