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

type Pods struct {
	Annotations      map[string]string      `json:"annotations,omitempty"`
	Labels           map[string]string      `json:"labels,omitempty"`
	NodeSelector     map[string]string      `json:"nodeSelector,omitempty"`
	SecurityContext  *JsonNode              `json:"securityContext,omitempty"`
	Affinity         *JsonNode              `json:"affinity,omitempty"`
	Tolerations      []JsonNode             `json:"tolerations,omitempty"`
	VolumeMounts     []VolumeAndMount       `json:"volumeMounts,omitempty"`
	EnvVars          []EnvVar               `json:"envVars,omitempty"`
	ImagePullSecrets []LocalObjectReference `json:"imagePullSecrets,omitempty"`
}
