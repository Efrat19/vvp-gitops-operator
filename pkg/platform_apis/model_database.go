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

//+kubebuilder:object:generate=true
type Database struct {
	Comment    string            `json:"comment,omitempty"`
	Name       string            `json:"name,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}
