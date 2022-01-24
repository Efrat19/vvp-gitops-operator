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

type JobStatus struct {
	Failure      *Failure          `json:"failure,omitempty"`
	SinkTables   []JobTable        `json:"sinkTables,omitempty"`
	SourceTables []JobTable        `json:"sourceTables,omitempty"`
	Started      *JobStatusStarted `json:"started,omitempty"`
	State        string            `json:"state,omitempty"`
}
