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

// Represents a source location of a volume to mount, managed by an external CSI driver
type V1CsiVolumeSource struct {
	// Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.
	Driver string `json:"driver"`
	// Filesystem type to mount. Ex. \"ext4\", \"xfs\", \"ntfs\". If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.
	FsType string `json:"fsType,omitempty"`
	NodePublishSecretRef *V1LocalObjectReference `json:"nodePublishSecretRef,omitempty"`
	// Specifies a read-only configuration for the volume. Defaults to false (read/write).
	ReadOnly bool `json:"readOnly,omitempty"`
	// VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.
	VolumeAttributes map[string]string `json:"volumeAttributes,omitempty"`
}
