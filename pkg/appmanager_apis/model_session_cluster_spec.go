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

type SessionClusterSpec struct {
	State string `json:"state,omitempty"`
	DeploymentTargetName string `json:"deploymentTargetName,omitempty"`
	FlinkVersion string `json:"flinkVersion,omitempty"`
	FlinkImageRegistry string `json:"flinkImageRegistry,omitempty"`
	FlinkImageRepository string `json:"flinkImageRepository,omitempty"`
	FlinkImageTag string `json:"flinkImageTag,omitempty"`
	NumberOfTaskManagers int32 `json:"numberOfTaskManagers,omitempty"`
	Resources map[string]ResourceSpec `json:"resources,omitempty"`
	FlinkConfiguration map[string]string `json:"flinkConfiguration,omitempty"`
	Logging *Logging `json:"logging,omitempty"`
	Kubernetes *KubernetesOptions `json:"kubernetes,omitempty"`
}