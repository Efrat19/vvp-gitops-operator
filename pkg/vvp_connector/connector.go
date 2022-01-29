package vvp_connector

import (
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
)

// DeploymentReconciler reconciles a Deployment object
type VvpConnector struct {
	client *appmanager_apis.APIClient
}

func NewConnector() (*VvpConnector, error) {
	vc := &VvpConnector{}
	if err := vc.initClient(); err != nil {
		return nil, err
	}
	return vc, nil
}
func (c *VvpConnector) initClient() error {
	cfg := &appmanager_apis.Configuration{
		BasePath:      "http://localhost:8080",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	c.client = appmanager_apis.NewAPIClient(cfg)
	return nil
}
