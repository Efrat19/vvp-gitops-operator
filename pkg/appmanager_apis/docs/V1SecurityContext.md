# V1SecurityContext

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPrivilegeEscalation** | **bool** | AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN | [optional] [default to null]
**Capabilities** | [***V1Capabilities**](V1Capabilities.md) |  | [optional] [default to null]
**Privileged** | **bool** | Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. | [optional] [default to null]
**ProcMount** | **string** | procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. | [optional] [default to null]
**ReadOnlyRootFilesystem** | **bool** | Whether this container has a read-only root filesystem. Default is false. | [optional] [default to null]
**RunAsGroup** | **int64** | The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. | [optional] [default to null]
**RunAsNonRoot** | **bool** | Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. | [optional] [default to null]
**RunAsUser** | **int64** | The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. | [optional] [default to null]
**SeLinuxOptions** | [***V1SeLinuxOptions**](V1SELinuxOptions.md) |  | [optional] [default to null]
**SeccompProfile** | [***V1SeccompProfile**](V1SeccompProfile.md) |  | [optional] [default to null]
**WindowsOptions** | [***V1WindowsSecurityContextOptions**](V1WindowsSecurityContextOptions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


