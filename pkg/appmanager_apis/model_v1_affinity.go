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

// Affinity is a group of affinity scheduling rules.
//+kubebuilder:object:generate=true
type V1Affinity struct {
	NodeAffinity    *V1NodeAffinity    `json:"nodeAffinity,omitempty"`
	PodAffinity     *V1PodAffinity     `json:"podAffinity,omitempty"`
	PodAntiAffinity *V1PodAntiAffinity `json:"podAntiAffinity,omitempty"`
}
