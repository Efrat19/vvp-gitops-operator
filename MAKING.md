```bash
# Create Project:
kubebuilder init --domain efrat19.io --repo efrat19.io/vvp-gitops-operator  

# Create CRDs:
kubebuilder create api --group appmanager.vvp --version v1alpha1 --kind Deployment
kubebuilder create api --group appmanager.vvp --version v1alpha1 --kind DeploymentTarget
kubebuilder create api --group appmanager.vvp --version v1alpha1 --kind Savepoint
kubebuilder create api --group appmanager.vvp --version v1alpha1 --kind SecretValue
kubebuilder create api --group appmanager.vvp --version v1alpha1 --kind SessionCluster

kubebuilder edit --multigroup=true

kubebuilder create api --group platform.vvp --version v1alpha1 --kind ApiTokens
kubebuilder create api --group platform.vvp --version v1alpha1 --kind CatalogConnectors
kubebuilder create api --group platform.vvp --version v1alpha1 --kind Connectors
kubebuilder create api --group platform.vvp --version v1alpha1 --kind Formats
kubebuilder create api --group platform.vvp --version v1alpha1 --kind SqlScripts
kubebuilder create api --group platform.vvp --version v1alpha1 --kind UdfArtifacts

# Create VVP Client for ververica platform 2.6.1
mkdir -p pkg/appmanager_apis
mkdir -p pkg/platform_apis
docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate  -i http://host.docker.internal:8080/swagger.json -l go -o /local/pkg/platform_apis
docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate  -i http://host.docker.internal:8080/api/swagger.json -l go -o /local/pkg/appmanager_apis

# Gen k9s files
python3 pre-gen.py
make install 
kustomize build config/crd > k.yaml
k create k.yaml
make build
make run ENABLE_WEBHOOKS=false

mockgen -destination=mocks/vvp_client/client.go efrat19.io/vvp-gitops-operator/pkg/vvp_client VvpClient 
```