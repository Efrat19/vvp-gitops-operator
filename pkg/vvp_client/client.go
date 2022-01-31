package vvp_client

import (
	// appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/api/v1alpha1"
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	// "github.com/gophercloud/gophercloud/openstack/cdn/v1/base"
)

type VvpClient interface {
	Deployments() deploymentsService
}

type vvpClient struct {
	deploymentsService
}

func NewClient() VvpClient {
	return &vvpClient{}
}

func (v *vvpClient) Deployments() deploymentsService {
	if &v.deploymentsService == nil {
		v.deploymentsService = deploymentsService{client: initAppManagerClient()}
	}
	return v.deploymentsService
}

func initAppManagerClient() *appmanager_apis.APIClient {
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
