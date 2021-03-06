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

// PodTemplateSpec describes the data a pod should have when created from a template
//+kubebuilder:object:generate=true
type V1PodTemplateSpec struct {
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
	Spec     *V1PodSpec    `json:"spec,omitempty"`
}
