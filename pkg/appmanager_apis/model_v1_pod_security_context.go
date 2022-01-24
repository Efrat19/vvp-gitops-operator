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

// PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext.  Field values of container.securityContext take precedence over field values of PodSecurityContext.
type V1PodSecurityContext struct {
	// A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume.
	FsGroup int64 `json:"fsGroup,omitempty"`
	// fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are \"OnRootMismatch\" and \"Always\". If not specified, \"Always\" is used.
	FsGroupChangePolicy string `json:"fsGroupChangePolicy,omitempty"`
	// The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
	RunAsGroup int64 `json:"runAsGroup,omitempty"`
	// Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsNonRoot bool `json:"runAsNonRoot,omitempty"`
	// The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
	RunAsUser      int64             `json:"runAsUser,omitempty"`
	SeLinuxOptions *V1SeLinuxOptions `json:"seLinuxOptions,omitempty"`
	SeccompProfile *V1SeccompProfile `json:"seccompProfile,omitempty"`
	// A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container.
	SupplementalGroups []int64 `json:"supplementalGroups,omitempty"`
	// Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch.
	Sysctls        []V1Sysctl                       `json:"sysctls,omitempty"`
	WindowsOptions *V1WindowsSecurityContextOptions `json:"windowsOptions,omitempty"`
}
