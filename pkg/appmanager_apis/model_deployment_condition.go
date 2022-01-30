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
	time "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:generate=true
type DeploymentCondition struct {
	LastTransitionTime time.Time `json:"lastTransitionTime,omitempty"`
	LastUpdateTime     time.Time `json:"lastUpdateTime,omitempty"`
	Message            string    `json:"message,omitempty"`
	Reason             string    `json:"reason,omitempty"`
	Status             string    `json:"status,omitempty"`
	Type_              string    `json:"type,omitempty"`
}
