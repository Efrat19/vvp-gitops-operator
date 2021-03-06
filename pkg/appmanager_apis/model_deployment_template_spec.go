/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.6.1
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

//+kubebuilder:object:generate=true
type DeploymentTemplateSpec struct {
	Artifact             *Artifact               `json:"artifact,omitempty"`
	FlinkConfiguration   map[string]string       `json:"flinkConfiguration,omitempty"`
	Kubernetes           *KubernetesOptions      `json:"kubernetes,omitempty"`
	Logging              *Logging                `json:"logging,omitempty"`
	NumberOfTaskManagers int32                   `json:"numberOfTaskManagers,omitempty"`
	Parallelism          int32                   `json:"parallelism,omitempty"`
	Resources            map[string]ResourceSpec `json:"resources,omitempty"`
}
