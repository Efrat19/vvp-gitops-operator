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

// Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
type V1FlockerVolumeSource struct {
	// Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated
	DatasetName string `json:"datasetName,omitempty"`
	// UUID of the dataset. This is unique identifier of a Flocker dataset
	DatasetUUID string `json:"datasetUUID,omitempty"`
}
