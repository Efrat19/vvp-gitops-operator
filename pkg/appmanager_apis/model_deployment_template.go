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

type DeploymentTemplate struct {
	Metadata *DeploymentTemplateMetadata `json:"metadata,omitempty"`
	Spec *DeploymentTemplateSpec `json:"spec,omitempty"`
}
