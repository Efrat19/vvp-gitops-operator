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

type ApiToken struct {
	CreateTime time.Time `json:"createTime,omitempty"`
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
	Secret string `json:"secret,omitempty"`
}
