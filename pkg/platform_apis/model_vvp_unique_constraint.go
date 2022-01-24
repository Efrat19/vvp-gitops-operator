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

type VvpUniqueConstraint struct {
	Column []string `json:"column,omitempty"`
	Enforced bool `json:"enforced,omitempty"`
	Name string `json:"name,omitempty"`
	Type_ string `json:"type,omitempty"`
}