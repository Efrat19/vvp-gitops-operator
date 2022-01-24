/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.6.1
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

//+kubebuilder:object:generate=true
type JobTableColumn struct {
	Name  string `json:"name,omitempty"`
	Type_ string `json:"type,omitempty"`
}
