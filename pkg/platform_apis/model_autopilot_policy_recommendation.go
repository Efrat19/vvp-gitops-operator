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

type AutopilotPolicyRecommendation struct {
	DeploymentPatch string `json:"deploymentPatch,omitempty"`
	Description string `json:"description,omitempty"`
	LastUpdateTime time.Time `json:"lastUpdateTime,omitempty"`
}