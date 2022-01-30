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
type JobStatusStarted struct {
	FlinkJobId               string    `json:"flinkJobId,omitempty"`
	LastUpdateTime           time.Time `json:"lastUpdateTime,omitempty"`
	ObservedFlinkJobRestarts int32     `json:"observedFlinkJobRestarts,omitempty"`
	ObservedFlinkJobStatus   string    `json:"observedFlinkJobStatus,omitempty"`
	StartedAt                time.Time `json:"startedAt,omitempty"`
}
