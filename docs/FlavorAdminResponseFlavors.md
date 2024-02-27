# FlavorAdminResponseFlavors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OpenstackId** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **float32** |  | [optional] 
**Disk** | Pointer to **int32** |  | [optional] 
**Ephemeral** | Pointer to **int32** |  | [optional] 
**Gpu** | Pointer to **string** |  | [optional] 
**GpuCount** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**IsCustom** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Organizations** | **[]int32** |  | 
**Flavors** | Pointer to **[]string** |  | [optional] 
**Projects** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFlavorAdminResponseFlavors

`func NewFlavorAdminResponseFlavors(organizations []int32, ) *FlavorAdminResponseFlavors`

NewFlavorAdminResponseFlavors instantiates a new FlavorAdminResponseFlavors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorAdminResponseFlavorsWithDefaults

`func NewFlavorAdminResponseFlavorsWithDefaults() *FlavorAdminResponseFlavors`

NewFlavorAdminResponseFlavorsWithDefaults instantiates a new FlavorAdminResponseFlavors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlavorAdminResponseFlavors) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlavorAdminResponseFlavors) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlavorAdminResponseFlavors) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FlavorAdminResponseFlavors) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FlavorAdminResponseFlavors) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlavorAdminResponseFlavors) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlavorAdminResponseFlavors) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlavorAdminResponseFlavors) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FlavorAdminResponseFlavors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlavorAdminResponseFlavors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlavorAdminResponseFlavors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlavorAdminResponseFlavors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOpenstackId

`func (o *FlavorAdminResponseFlavors) GetOpenstackId() string`

GetOpenstackId returns the OpenstackId field if non-nil, zero value otherwise.

### GetOpenstackIdOk

`func (o *FlavorAdminResponseFlavors) GetOpenstackIdOk() (*string, bool)`

GetOpenstackIdOk returns a tuple with the OpenstackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackId

`func (o *FlavorAdminResponseFlavors) SetOpenstackId(v string)`

SetOpenstackId sets OpenstackId field to given value.

### HasOpenstackId

`func (o *FlavorAdminResponseFlavors) HasOpenstackId() bool`

HasOpenstackId returns a boolean if a field has been set.

### GetRegionId

`func (o *FlavorAdminResponseFlavors) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *FlavorAdminResponseFlavors) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *FlavorAdminResponseFlavors) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *FlavorAdminResponseFlavors) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetRegion

`func (o *FlavorAdminResponseFlavors) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FlavorAdminResponseFlavors) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FlavorAdminResponseFlavors) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FlavorAdminResponseFlavors) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCpu

`func (o *FlavorAdminResponseFlavors) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *FlavorAdminResponseFlavors) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *FlavorAdminResponseFlavors) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *FlavorAdminResponseFlavors) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *FlavorAdminResponseFlavors) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *FlavorAdminResponseFlavors) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *FlavorAdminResponseFlavors) SetRam(v float32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *FlavorAdminResponseFlavors) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetDisk

`func (o *FlavorAdminResponseFlavors) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *FlavorAdminResponseFlavors) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *FlavorAdminResponseFlavors) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *FlavorAdminResponseFlavors) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetEphemeral

`func (o *FlavorAdminResponseFlavors) GetEphemeral() int32`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *FlavorAdminResponseFlavors) GetEphemeralOk() (*int32, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *FlavorAdminResponseFlavors) SetEphemeral(v int32)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *FlavorAdminResponseFlavors) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetGpu

`func (o *FlavorAdminResponseFlavors) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *FlavorAdminResponseFlavors) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *FlavorAdminResponseFlavors) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *FlavorAdminResponseFlavors) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetGpuCount

`func (o *FlavorAdminResponseFlavors) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *FlavorAdminResponseFlavors) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *FlavorAdminResponseFlavors) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *FlavorAdminResponseFlavors) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetOrder

`func (o *FlavorAdminResponseFlavors) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FlavorAdminResponseFlavors) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FlavorAdminResponseFlavors) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FlavorAdminResponseFlavors) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetIsPublic

`func (o *FlavorAdminResponseFlavors) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *FlavorAdminResponseFlavors) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *FlavorAdminResponseFlavors) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *FlavorAdminResponseFlavors) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetIsCustom

`func (o *FlavorAdminResponseFlavors) GetIsCustom() bool`

GetIsCustom returns the IsCustom field if non-nil, zero value otherwise.

### GetIsCustomOk

`func (o *FlavorAdminResponseFlavors) GetIsCustomOk() (*bool, bool)`

GetIsCustomOk returns a tuple with the IsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustom

`func (o *FlavorAdminResponseFlavors) SetIsCustom(v bool)`

SetIsCustom sets IsCustom field to given value.

### HasIsCustom

`func (o *FlavorAdminResponseFlavors) HasIsCustom() bool`

HasIsCustom returns a boolean if a field has been set.

### GetStatus

`func (o *FlavorAdminResponseFlavors) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlavorAdminResponseFlavors) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlavorAdminResponseFlavors) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlavorAdminResponseFlavors) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrganizations

`func (o *FlavorAdminResponseFlavors) GetOrganizations() []int32`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *FlavorAdminResponseFlavors) GetOrganizationsOk() (*[]int32, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *FlavorAdminResponseFlavors) SetOrganizations(v []int32)`

SetOrganizations sets Organizations field to given value.


### GetFlavors

`func (o *FlavorAdminResponseFlavors) GetFlavors() []string`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *FlavorAdminResponseFlavors) GetFlavorsOk() (*[]string, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *FlavorAdminResponseFlavors) SetFlavors(v []string)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *FlavorAdminResponseFlavors) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### GetProjects

`func (o *FlavorAdminResponseFlavors) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *FlavorAdminResponseFlavors) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *FlavorAdminResponseFlavors) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *FlavorAdminResponseFlavors) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FlavorAdminResponseFlavors) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlavorAdminResponseFlavors) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlavorAdminResponseFlavors) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FlavorAdminResponseFlavors) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


