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
type CatalogConnector struct {
	Dependencies []string   `json:"dependencies,omitempty"`
	Name         string     `json:"name,omitempty"`
	Packaged     bool       `json:"packaged,omitempty"`
	Properties   []Property `json:"properties,omitempty"`
	ReadOnly     bool       `json:"readOnly,omitempty"`
	Type_        string     `json:"type,omitempty"`
}
