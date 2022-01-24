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

type ResourceQuotaSpec struct {
	Limits           *ResourceConsumption `json:"limits,omitempty"`
	ToleratedOveruse *ResourceConsumption `json:"toleratedOveruse,omitempty"`
	Type_            string               `json:"type,omitempty"`
}
