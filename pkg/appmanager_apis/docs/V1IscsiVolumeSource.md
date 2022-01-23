# V1IscsiVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChapAuthDiscovery** | **bool** | whether support iSCSI Discovery CHAP authentication | [optional] [default to null]
**ChapAuthSession** | **bool** | whether support iSCSI Session CHAP authentication | [optional] [default to null]
**FsType** | **string** | Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi | [optional] [default to null]
**InitiatorName** | **string** | Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface &lt;target portal&gt;:&lt;volume name&gt; will be created for the connection. | [optional] [default to null]
**Iqn** | **string** | Target iSCSI Qualified Name. | [default to null]
**IscsiInterface** | **string** | iSCSI Interface Name that uses an iSCSI transport. Defaults to &#39;default&#39; (tcp). | [optional] [default to null]
**Lun** | **int32** | iSCSI Target Lun number. | [default to null]
**Portals** | **[]string** | iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). | [optional] [default to null]
**ReadOnly** | **bool** | ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. | [optional] [default to null]
**SecretRef** | [***V1LocalObjectReference**](V1LocalObjectReference.md) |  | [optional] [default to null]
**TargetPortal** | **string** | iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


