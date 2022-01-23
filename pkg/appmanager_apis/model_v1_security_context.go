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

// SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext.  When both are set, the values in SecurityContext take precedence.
type V1SecurityContext struct {
	// AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation,omitempty"`
	Capabilities *V1Capabilities `json:"capabilities,omitempty"`
	// Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.
	Privileged bool `json:"privileged,omitempty"`
	// procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled.
	ProcMount string `json:"procMount,omitempty"`
	// Whether this container has a read-only root filesystem. Default is false.
	ReadOnlyRootFilesystem bool `json:"readOnlyRootFilesystem,omitempty"`
	// The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsGroup int64 `json:"runAsGroup,omitempty"`
	// Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsNonRoot bool `json:"runAsNonRoot,omitempty"`
	// The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsUser int64 `json:"runAsUser,omitempty"`
	SeLinuxOptions *V1SeLinuxOptions `json:"seLinuxOptions,omitempty"`
	SeccompProfile *V1SeccompProfile `json:"seccompProfile,omitempty"`
	WindowsOptions *V1WindowsSecurityContextOptions `json:"windowsOptions,omitempty"`
}
