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
type ReferencedTable struct {
	CatalogName  string    `json:"catalogName,omitempty"`
	DatabaseName string    `json:"databaseName,omitempty"`
	Table        *VvpTable `json:"table,omitempty"`
	Temporary    bool      `json:"temporary,omitempty"`
}
