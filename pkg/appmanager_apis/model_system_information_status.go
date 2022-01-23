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

type SystemInformationStatus struct {
	JvmVersion string `json:"jvmVersion,omitempty"`
	ResourceQuota *ResourceQuota `json:"resourceQuota,omitempty"`
	RevisionInformation *RevisionInformation `json:"revisionInformation,omitempty"`
}
