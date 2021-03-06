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

// VolumeMount describes a mounting of a Volume within a container.
//+kubebuilder:object:generate=true
type V1VolumeMount struct {
	// Path within the container at which the volume should be mounted.  Must not contain ':'.
	MountPath string `json:"mountPath"`
	// mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.
	MountPropagation string `json:"mountPropagation,omitempty"`
	// This must match the Name of a Volume.
	Name string `json:"name"`
	// Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
	ReadOnly bool `json:"readOnly,omitempty"`
	// Path within the volume from which the container's volume should be mounted. Defaults to \"\" (volume's root).
	SubPath string `json:"subPath,omitempty"`
	// Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to \"\" (volume's root). SubPathExpr and SubPath are mutually exclusive.
	SubPathExpr string `json:"subPathExpr,omitempty"`
}
