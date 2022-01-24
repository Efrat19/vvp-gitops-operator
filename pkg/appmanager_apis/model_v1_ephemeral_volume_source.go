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

// Represents an ephemeral volume that is handled by a normal storage driver.
//+kubebuilder:object:generate=true
type V1EphemeralVolumeSource struct {
	// Specifies a read-only configuration for the volume. Defaults to false (read/write).
	ReadOnly            bool                             `json:"readOnly,omitempty"`
	VolumeClaimTemplate *V1PersistentVolumeClaimTemplate `json:"volumeClaimTemplate,omitempty"`
}
