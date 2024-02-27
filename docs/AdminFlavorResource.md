# AdminFlavorResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OpenstackId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **float32** |  | [optional] 
**Disk** | Pointer to **int32** |  | [optional] 
**Gpu** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**GpuCount** | Pointer to **int32** |  | [optional] 
**StockAvailable** | Pointer to **bool** |  | [optional] 
**Nodes** | Pointer to [**[]AdminNodeResource**](AdminNodeResource.md) |  | [optional] 
**Vms** | Pointer to **[]int32** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**IsCustom** | Pointer to **bool** |  | [optional] 
**Ephemeral** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAdminFlavorResource

`func NewAdminFlavorResource() *AdminFlavorResource`

NewAdminFlavorResource instantiates a new AdminFlavorResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminFlavorResourceWithDefaults

`func NewAdminFlavorResourceWithDefaults() *AdminFlavorResource`

NewAdminFlavorResourceWithDefaults instantiates a new AdminFlavorResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminFlavorResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminFlavorResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminFlavorResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdminFlavorResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOpenstackId

`func (o *AdminFlavorResource) GetOpenstackId() string`

GetOpenstackId returns the OpenstackId field if non-nil, zero value otherwise.

### GetOpenstackIdOk

`func (o *AdminFlavorResource) GetOpenstackIdOk() (*string, bool)`

GetOpenstackIdOk returns a tuple with the OpenstackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackId

`func (o *AdminFlavorResource) SetOpenstackId(v string)`

SetOpenstackId sets OpenstackId field to given value.

### HasOpenstackId

`func (o *AdminFlavorResource) HasOpenstackId() bool`

HasOpenstackId returns a boolean if a field has been set.

### GetName

`func (o *AdminFlavorResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminFlavorResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminFlavorResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminFlavorResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegionName

`func (o *AdminFlavorResource) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *AdminFlavorResource) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *AdminFlavorResource) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *AdminFlavorResource) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetCpu

`func (o *AdminFlavorResource) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *AdminFlavorResource) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *AdminFlavorResource) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *AdminFlavorResource) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *AdminFlavorResource) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *AdminFlavorResource) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *AdminFlavorResource) SetRam(v float32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *AdminFlavorResource) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetDisk

`func (o *AdminFlavorResource) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *AdminFlavorResource) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *AdminFlavorResource) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *AdminFlavorResource) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetGpu

`func (o *AdminFlavorResource) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *AdminFlavorResource) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *AdminFlavorResource) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *AdminFlavorResource) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetStatus

`func (o *AdminFlavorResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminFlavorResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminFlavorResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminFlavorResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGpuCount

`func (o *AdminFlavorResource) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *AdminFlavorResource) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *AdminFlavorResource) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *AdminFlavorResource) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetStockAvailable

`func (o *AdminFlavorResource) GetStockAvailable() bool`

GetStockAvailable returns the StockAvailable field if non-nil, zero value otherwise.

### GetStockAvailableOk

`func (o *AdminFlavorResource) GetStockAvailableOk() (*bool, bool)`

GetStockAvailableOk returns a tuple with the StockAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockAvailable

`func (o *AdminFlavorResource) SetStockAvailable(v bool)`

SetStockAvailable sets StockAvailable field to given value.

### HasStockAvailable

`func (o *AdminFlavorResource) HasStockAvailable() bool`

HasStockAvailable returns a boolean if a field has been set.

### GetNodes

`func (o *AdminFlavorResource) GetNodes() []AdminNodeResource`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *AdminFlavorResource) GetNodesOk() (*[]AdminNodeResource, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *AdminFlavorResource) SetNodes(v []AdminNodeResource)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *AdminFlavorResource) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetVms

`func (o *AdminFlavorResource) GetVms() []int32`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *AdminFlavorResource) GetVmsOk() (*[]int32, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *AdminFlavorResource) SetVms(v []int32)`

SetVms sets Vms field to given value.

### HasVms

`func (o *AdminFlavorResource) HasVms() bool`

HasVms returns a boolean if a field has been set.

### GetIsPublic

`func (o *AdminFlavorResource) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AdminFlavorResource) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AdminFlavorResource) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *AdminFlavorResource) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetIsCustom

`func (o *AdminFlavorResource) GetIsCustom() bool`

GetIsCustom returns the IsCustom field if non-nil, zero value otherwise.

### GetIsCustomOk

`func (o *AdminFlavorResource) GetIsCustomOk() (*bool, bool)`

GetIsCustomOk returns a tuple with the IsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustom

`func (o *AdminFlavorResource) SetIsCustom(v bool)`

SetIsCustom sets IsCustom field to given value.

### HasIsCustom

`func (o *AdminFlavorResource) HasIsCustom() bool`

HasIsCustom returns a boolean if a field has been set.

### GetEphemeral

`func (o *AdminFlavorResource) GetEphemeral() int32`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *AdminFlavorResource) GetEphemeralOk() (*int32, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *AdminFlavorResource) SetEphemeral(v int32)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *AdminFlavorResource) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetDescription

`func (o *AdminFlavorResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminFlavorResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminFlavorResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminFlavorResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminFlavorResource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminFlavorResource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminFlavorResource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminFlavorResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


