# AdminAddUpdateImageOrganizationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**OpenstackId** | **string** |  | 
**RegionId** | **int32** |  | 
**Type** | **string** |  | 
**Version** | **string** |  | 
**Size** | **int32** |  | 
**IsPublic** | **bool** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Organizations** | **[]int32** |  | 

## Methods

### NewAdminAddUpdateImageOrganizationPayload

`func NewAdminAddUpdateImageOrganizationPayload(name string, openstackId string, regionId int32, type_ string, version string, size int32, isPublic bool, organizations []int32, ) *AdminAddUpdateImageOrganizationPayload`

NewAdminAddUpdateImageOrganizationPayload instantiates a new AdminAddUpdateImageOrganizationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAddUpdateImageOrganizationPayloadWithDefaults

`func NewAdminAddUpdateImageOrganizationPayloadWithDefaults() *AdminAddUpdateImageOrganizationPayload`

NewAdminAddUpdateImageOrganizationPayloadWithDefaults instantiates a new AdminAddUpdateImageOrganizationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdminAddUpdateImageOrganizationPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminAddUpdateImageOrganizationPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminAddUpdateImageOrganizationPayload) SetName(v string)`

SetName sets Name field to given value.


### GetOpenstackId

`func (o *AdminAddUpdateImageOrganizationPayload) GetOpenstackId() string`

GetOpenstackId returns the OpenstackId field if non-nil, zero value otherwise.

### GetOpenstackIdOk

`func (o *AdminAddUpdateImageOrganizationPayload) GetOpenstackIdOk() (*string, bool)`

GetOpenstackIdOk returns a tuple with the OpenstackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackId

`func (o *AdminAddUpdateImageOrganizationPayload) SetOpenstackId(v string)`

SetOpenstackId sets OpenstackId field to given value.


### GetRegionId

`func (o *AdminAddUpdateImageOrganizationPayload) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AdminAddUpdateImageOrganizationPayload) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AdminAddUpdateImageOrganizationPayload) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.


### GetType

`func (o *AdminAddUpdateImageOrganizationPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdminAddUpdateImageOrganizationPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdminAddUpdateImageOrganizationPayload) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *AdminAddUpdateImageOrganizationPayload) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdminAddUpdateImageOrganizationPayload) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdminAddUpdateImageOrganizationPayload) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetSize

`func (o *AdminAddUpdateImageOrganizationPayload) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AdminAddUpdateImageOrganizationPayload) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AdminAddUpdateImageOrganizationPayload) SetSize(v int32)`

SetSize sets Size field to given value.


### GetIsPublic

`func (o *AdminAddUpdateImageOrganizationPayload) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AdminAddUpdateImageOrganizationPayload) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AdminAddUpdateImageOrganizationPayload) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetDescription

`func (o *AdminAddUpdateImageOrganizationPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminAddUpdateImageOrganizationPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminAddUpdateImageOrganizationPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminAddUpdateImageOrganizationPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *AdminAddUpdateImageOrganizationPayload) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AdminAddUpdateImageOrganizationPayload) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AdminAddUpdateImageOrganizationPayload) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AdminAddUpdateImageOrganizationPayload) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOrganizations

`func (o *AdminAddUpdateImageOrganizationPayload) GetOrganizations() []int32`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *AdminAddUpdateImageOrganizationPayload) GetOrganizationsOk() (*[]int32, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *AdminAddUpdateImageOrganizationPayload) SetOrganizations(v []int32)`

SetOrganizations sets Organizations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


