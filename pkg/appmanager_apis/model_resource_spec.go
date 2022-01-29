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

import (
	"k8s.io/apimachinery/pkg/api/resource"
)

//+kubebuilder:object:generate=true
type ResourceSpec struct {
	Cpu    resource.Quantity `json:"cpu,omitempty"`
	Memory *string           `json:"memory,omitempty"`
}
