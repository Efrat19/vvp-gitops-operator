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

// PersistentVolumeClaimTemplate is used to produce PersistentVolumeClaim objects as part of an EphemeralVolumeSource.
//+kubebuilder:object:generate=true
type V1PersistentVolumeClaimTemplate struct {
	Metadata *V1ObjectMeta                `json:"metadata,omitempty"`
	Spec     *V1PersistentVolumeClaimSpec `json:"spec"`
}
