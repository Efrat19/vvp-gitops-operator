# V1RbdVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | **string** | Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd | [optional] [default to null]
**Image** | **string** | The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [default to null]
**Keyring** | **string** | Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [optional] [default to null]
**Monitors** | **[]string** | A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [default to null]
**Pool** | **string** | The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [optional] [default to null]
**ReadOnly** | **bool** | ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [optional] [default to null]
**SecretRef** | [***V1LocalObjectReference**](V1LocalObjectReference.md) |  | [optional] [default to null]
**User** | **string** | The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


