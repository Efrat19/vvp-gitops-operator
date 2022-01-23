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

// EnvFromSource represents the source of a set of ConfigMaps
type V1EnvFromSource struct {
	ConfigMapRef *V1ConfigMapEnvSource `json:"configMapRef,omitempty"`
	// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	Prefix string `json:"prefix,omitempty"`
	SecretRef *V1SecretEnvSource `json:"secretRef,omitempty"`
}
