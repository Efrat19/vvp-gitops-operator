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

// Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
//+kubebuilder:object:generate=true
type V1QuobyteVolumeSource struct {
	// Group to map volume access to Default is no group
	Group string `json:"group,omitempty"`
	// ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
	ReadOnly bool `json:"readOnly,omitempty"`
	// Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes
	Registry string `json:"registry"`
	// Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin
	Tenant string `json:"tenant,omitempty"`
	// User to map volume access to Defaults to serivceaccount user
	User string `json:"user,omitempty"`
	// Volume is a string that references an already created Quobyte volume by name.
	Volume string `json:"volume"`
}
