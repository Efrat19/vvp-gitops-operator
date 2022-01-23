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

// A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
type V1NodeSelectorTerm struct {
	// A list of node selector requirements by node's labels.
	MatchExpressions []V1NodeSelectorRequirement `json:"matchExpressions,omitempty"`
	// A list of node selector requirements by node's fields.
	MatchFields []V1NodeSelectorRequirement `json:"matchFields,omitempty"`
}
