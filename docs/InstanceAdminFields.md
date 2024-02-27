# InstanceAdminFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to [**InstanceEnvironmentFields**](InstanceEnvironmentFields.md) |  | [optional] 
**Image** | Pointer to [**InstanceImageFields**](InstanceImageFields.md) |  | [optional] 
**Flavor** | Pointer to [**InstanceFlavorFields**](InstanceFlavorFields.md) |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Keypair** | Pointer to [**InstanceKeypairFields**](InstanceKeypairFields.md) |  | [optional] 
**VolumeAttachments** | Pointer to [**[]VolumeAttachmentFields**](VolumeAttachmentFields.md) |  | [optional] 
**SecurityRules** | Pointer to [**[]SecurityRulesFieldsforInstance**](SecurityRulesFieldsforInstance.md) |  | [optional] 
**PowerState** | Pointer to **string** |  | [optional] 
**VmState** | Pointer to **string** |  | [optional] 
**FixedIp** | Pointer to **string** |  | [optional] 
**FloatingIp** | Pointer to **string** |  | [optional] 
**FloatingIpStatus** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**ContractId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**OpenstackId** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 

## Methods

### NewInstanceAdminFields

`func NewInstanceAdminFields() *InstanceAdminFields`

NewInstanceAdminFields instantiates a new InstanceAdminFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceAdminFieldsWithDefaults

`func NewInstanceAdminFieldsWithDefaults() *InstanceAdminFields`

NewInstanceAdminFieldsWithDefaults instantiates a new InstanceAdminFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceAdminFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceAdminFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceAdminFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceAdminFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InstanceAdminFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceAdminFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceAdminFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceAdminFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *InstanceAdminFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceAdminFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceAdminFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstanceAdminFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnvironment

`func (o *InstanceAdminFields) GetEnvironment() InstanceEnvironmentFields`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *InstanceAdminFields) GetEnvironmentOk() (*InstanceEnvironmentFields, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *InstanceAdminFields) SetEnvironment(v InstanceEnvironmentFields)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *InstanceAdminFields) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetImage

`func (o *InstanceAdminFields) GetImage() InstanceImageFields`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *InstanceAdminFields) GetImageOk() (*InstanceImageFields, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *InstanceAdminFields) SetImage(v InstanceImageFields)`

SetImage sets Image field to given value.

### HasImage

`func (o *InstanceAdminFields) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetFlavor

`func (o *InstanceAdminFields) GetFlavor() InstanceFlavorFields`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *InstanceAdminFields) GetFlavorOk() (*InstanceFlavorFields, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *InstanceAdminFields) SetFlavor(v InstanceFlavorFields)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *InstanceAdminFields) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetOs

`func (o *InstanceAdminFields) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *InstanceAdminFields) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *InstanceAdminFields) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *InstanceAdminFields) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetKeypair

`func (o *InstanceAdminFields) GetKeypair() InstanceKeypairFields`

GetKeypair returns the Keypair field if non-nil, zero value otherwise.

### GetKeypairOk

`func (o *InstanceAdminFields) GetKeypairOk() (*InstanceKeypairFields, bool)`

GetKeypairOk returns a tuple with the Keypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypair

`func (o *InstanceAdminFields) SetKeypair(v InstanceKeypairFields)`

SetKeypair sets Keypair field to given value.

### HasKeypair

`func (o *InstanceAdminFields) HasKeypair() bool`

HasKeypair returns a boolean if a field has been set.

### GetVolumeAttachments

`func (o *InstanceAdminFields) GetVolumeAttachments() []VolumeAttachmentFields`

GetVolumeAttachments returns the VolumeAttachments field if non-nil, zero value otherwise.

### GetVolumeAttachmentsOk

`func (o *InstanceAdminFields) GetVolumeAttachmentsOk() (*[]VolumeAttachmentFields, bool)`

GetVolumeAttachmentsOk returns a tuple with the VolumeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAttachments

`func (o *InstanceAdminFields) SetVolumeAttachments(v []VolumeAttachmentFields)`

SetVolumeAttachments sets VolumeAttachments field to given value.

### HasVolumeAttachments

`func (o *InstanceAdminFields) HasVolumeAttachments() bool`

HasVolumeAttachments returns a boolean if a field has been set.

### GetSecurityRules

`func (o *InstanceAdminFields) GetSecurityRules() []SecurityRulesFieldsforInstance`

GetSecurityRules returns the SecurityRules field if non-nil, zero value otherwise.

### GetSecurityRulesOk

`func (o *InstanceAdminFields) GetSecurityRulesOk() (*[]SecurityRulesFieldsforInstance, bool)`

GetSecurityRulesOk returns a tuple with the SecurityRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityRules

`func (o *InstanceAdminFields) SetSecurityRules(v []SecurityRulesFieldsforInstance)`

SetSecurityRules sets SecurityRules field to given value.

### HasSecurityRules

`func (o *InstanceAdminFields) HasSecurityRules() bool`

HasSecurityRules returns a boolean if a field has been set.

### GetPowerState

`func (o *InstanceAdminFields) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *InstanceAdminFields) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *InstanceAdminFields) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *InstanceAdminFields) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetVmState

`func (o *InstanceAdminFields) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *InstanceAdminFields) GetVmStateOk() (*string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *InstanceAdminFields) SetVmState(v string)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *InstanceAdminFields) HasVmState() bool`

HasVmState returns a boolean if a field has been set.

### GetFixedIp

`func (o *InstanceAdminFields) GetFixedIp() string`

GetFixedIp returns the FixedIp field if non-nil, zero value otherwise.

### GetFixedIpOk

`func (o *InstanceAdminFields) GetFixedIpOk() (*string, bool)`

GetFixedIpOk returns a tuple with the FixedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIp

`func (o *InstanceAdminFields) SetFixedIp(v string)`

SetFixedIp sets FixedIp field to given value.

### HasFixedIp

`func (o *InstanceAdminFields) HasFixedIp() bool`

HasFixedIp returns a boolean if a field has been set.

### GetFloatingIp

`func (o *InstanceAdminFields) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *InstanceAdminFields) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *InstanceAdminFields) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *InstanceAdminFields) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.

### GetFloatingIpStatus

`func (o *InstanceAdminFields) GetFloatingIpStatus() string`

GetFloatingIpStatus returns the FloatingIpStatus field if non-nil, zero value otherwise.

### GetFloatingIpStatusOk

`func (o *InstanceAdminFields) GetFloatingIpStatusOk() (*string, bool)`

GetFloatingIpStatusOk returns a tuple with the FloatingIpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpStatus

`func (o *InstanceAdminFields) SetFloatingIpStatus(v string)`

SetFloatingIpStatus sets FloatingIpStatus field to given value.

### HasFloatingIpStatus

`func (o *InstanceAdminFields) HasFloatingIpStatus() bool`

HasFloatingIpStatus returns a boolean if a field has been set.

### GetLocked

`func (o *InstanceAdminFields) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *InstanceAdminFields) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *InstanceAdminFields) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *InstanceAdminFields) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetContractId

`func (o *InstanceAdminFields) GetContractId() int32`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *InstanceAdminFields) GetContractIdOk() (*int32, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *InstanceAdminFields) SetContractId(v int32)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *InstanceAdminFields) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InstanceAdminFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceAdminFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceAdminFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InstanceAdminFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceAdminFields) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceAdminFields) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceAdminFields) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceAdminFields) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOpenstackId

`func (o *InstanceAdminFields) GetOpenstackId() string`

GetOpenstackId returns the OpenstackId field if non-nil, zero value otherwise.

### GetOpenstackIdOk

`func (o *InstanceAdminFields) GetOpenstackIdOk() (*string, bool)`

GetOpenstackIdOk returns a tuple with the OpenstackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackId

`func (o *InstanceAdminFields) SetOpenstackId(v string)`

SetOpenstackId sets OpenstackId field to given value.

### HasOpenstackId

`func (o *InstanceAdminFields) HasOpenstackId() bool`

HasOpenstackId returns a boolean if a field has been set.

### GetHost

`func (o *InstanceAdminFields) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InstanceAdminFields) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InstanceAdminFields) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *InstanceAdminFields) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


