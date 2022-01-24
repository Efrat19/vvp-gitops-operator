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

type DeploymentMetadata struct {
	Annotations     map[string]string `json:"annotations,omitempty"`
	CreatedAt       time.Time         `json:"createdAt,omitempty"`
	DisplayName     string            `json:"displayName,omitempty"`
	Id              string            `json:"id,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
	ModifiedAt      time.Time         `json:"modifiedAt,omitempty"`
	Name            string            `json:"name,omitempty"`
	Namespace       string            `json:"namespace,omitempty"`
	ResourceVersion int32             `json:"resourceVersion,omitempty"`
}
