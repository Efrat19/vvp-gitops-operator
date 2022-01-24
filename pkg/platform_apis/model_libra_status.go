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

type LibraStatus struct {
	LastActionTime time.Time `json:"lastActionTime,omitempty"`
	Message string `json:"message,omitempty"`
	Metrics string `json:"metrics,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
}