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

// Projection that may be projected along with other supported volume types
//+kubebuilder:object:generate=true
type V1VolumeProjection struct {
	ConfigMap           *V1ConfigMapProjection           `json:"configMap,omitempty"`
	DownwardAPI         *V1DownwardApiProjection         `json:"downwardAPI,omitempty"`
	Secret              *V1SecretProjection              `json:"secret,omitempty"`
	ServiceAccountToken *V1ServiceAccountTokenProjection `json:"serviceAccountToken,omitempty"`
}
