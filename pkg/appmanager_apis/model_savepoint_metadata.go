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
	"time"
)

type SavepointMetadata struct {
	Id              string            `json:"id,omitempty"`
	Namespace       string            `json:"namespace,omitempty"`
	CreatedAt       time.Time         `json:"createdAt,omitempty"`
	ModifiedAt      time.Time         `json:"modifiedAt,omitempty"`
	DeploymentId    string            `json:"deploymentId,omitempty"`
	JobId           string            `json:"jobId,omitempty"`
	Origin          string            `json:"origin,omitempty"`
	Type_           string            `json:"type,omitempty"`
	Annotations     map[string]string `json:"annotations,omitempty"`
	ResourceVersion int32             `json:"resourceVersion,omitempty"`
}
