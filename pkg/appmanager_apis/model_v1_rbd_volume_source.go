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

// Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
//+kubebuilder:object:generate=true
type V1RbdVolumeSource struct {
	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	FsType string `json:"fsType,omitempty"`
	// The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Image string `json:"image"`
	// Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Keyring string `json:"keyring,omitempty"`
	// A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Monitors []string `json:"monitors"`
	// The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Pool string `json:"pool,omitempty"`
	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	ReadOnly  bool                    `json:"readOnly,omitempty"`
	SecretRef *V1LocalObjectReference `json:"secretRef,omitempty"`
	// The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	User string `json:"user,omitempty"`
}
