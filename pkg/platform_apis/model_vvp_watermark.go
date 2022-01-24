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

type VvpWatermark struct {
	TimeColumn string `json:"timeColumn,omitempty"`
	WatermarkExpression string `json:"watermarkExpression,omitempty"`
	WatermarkType string `json:"watermarkType,omitempty"`
}