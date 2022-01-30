package vvp_connector

import (
	"os"

	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
	// "github.com/gophercloud/gophercloud/openstack/cdn/v1/base"
)

// DeploymentReconciler reconciles a Deployment object
type VvpConnector struct {
	client *appmanager_apis.APIClient
}

func initClient() *appmanager_apis.APIClient {
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
