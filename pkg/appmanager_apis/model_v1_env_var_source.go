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

// EnvVarSource represents a source for the value of an EnvVar.
type V1EnvVarSource struct {
	ConfigMapKeyRef  *V1ConfigMapKeySelector  `json:"configMapKeyRef,omitempty"`
	FieldRef         *V1ObjectFieldSelector   `json:"fieldRef,omitempty"`
	ResourceFieldRef *V1ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
	SecretKeyRef     *V1SecretKeySelector     `json:"secretKeyRef,omitempty"`
}
