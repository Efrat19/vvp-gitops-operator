package vvp_client

import (
	// appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/api/v1alpha1"
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	// "github.com/gophercloud/gophercloud/openstack/cdn/v1/base"
)

type VvpClient interface {
	Deployments() DeploymentsService
	// DeploymentTargets() DeploymentTargetsService
	// Events() EventsService
	// Jobs() JobsService
	// SavePoints() SavePointsService
	// SecretValues() SecretValuesService
	// SessionClusters() SessionClustersService
}

type vvpClient struct {
	appManagerClient   *appmanager_apis.APIClient
	DeploymentsService *DeploymentsService
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

// func (v *vvpClient) DeploymentTargets() DeploymentTargetsService {
// 	if &v.DeploymentTargetsService == nil {
// 		v.DeploymentTargetsService = DeploymentTargetsService{client: v.appManagerClient}
// 	}
// 	return v.DeploymentTargetsService
// }

// func (v *vvpClient) SavePoints() SavePointsService {
// 	if &v.SavePointsService == nil {
// 		v.SavePointsService = SavePointsService{client: v.appManagerClient}
// 	}
// 	return v.SavePointsService
// }

// func (v *vvpClient) SecretValues() SecretValuesService {
// 	if &v.SecretValuesService == nil {
// 		v.SecretValuesService = SecretValuesService{client: v.appManagerClient}
// 	}
// 	return v.SecretValuesService
// }

// func (v *vvpClient) SessionClusters() SessionClustersService {
// 	if &v.SessionClustersService == nil {
// 		v.SessionClustersService = SessionClustersService{client: v.appManagerClient}
// 	}
// 	return v.SessionClustersService
// }

func NewAppManagerClient() *appmanager_apis.APIClient {
	basePath := getEnv("VVP_URL", "http://localhost:8080")
	cfg := &appmanager_apis.Configuration{
		BasePath:      basePath,
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	return appmanager_apis.NewAPIClient(cfg)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
