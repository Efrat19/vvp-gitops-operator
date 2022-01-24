/*
 * Ververica Platform API
 *
 * The Ververica Platform APIs, excluding Application Manager.
 *
 * API version: 2.6.1
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type Namespace struct {
	CreateTime                time.Time     `json:"createTime,omitempty"`
	LifecyclePhase            string        `json:"lifecyclePhase,omitempty"`
	Name                      string        `json:"name,omitempty"`
	PreviewSessionClusterName string        `json:"previewSessionClusterName,omitempty"`
	RoleBindings              []RoleBinding `json:"roleBindings,omitempty"`
}
