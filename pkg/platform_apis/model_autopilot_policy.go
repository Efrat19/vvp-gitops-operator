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

type AutopilotPolicy struct {
	AutoscalerConfig *AutoscalerConfig `json:"autoscalerConfig,omitempty"`
	LibraConfig *LibraConfig `json:"libraConfig,omitempty"`
	Mode string `json:"mode,omitempty"`
	ModifyTime time.Time `json:"modifyTime,omitempty"`
	Name string `json:"name,omitempty"`
}
