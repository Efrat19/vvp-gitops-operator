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

// Adds and removes POSIX capabilities from running containers.
//+kubebuilder:object:generate=true
type V1Capabilities struct {
	// Added capabilities
	Add []string `json:"add,omitempty"`
	// Removed capabilities
	Drop []string `json:"drop,omitempty"`
}
