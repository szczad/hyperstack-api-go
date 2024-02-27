# AdminImageAdminFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OpenstackId** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to [**[]LableResonse**](LableResonse.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAdminImageAdminFields

`func NewAdminImageAdminFields() *AdminImageAdminFields`

NewAdminImageAdminFields instantiates a new AdminImageAdminFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminImageAdminFieldsWithDefaults

`func NewAdminImageAdminFieldsWithDefaults() *AdminImageAdminFields`

NewAdminImageAdminFieldsWithDefaults instantiates a new AdminImageAdminFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminImageAdminFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminImageAdminFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminImageAdminFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdminImageAdminFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AdminImageAdminFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminImageAdminFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminImageAdminFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminImageAdminFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpenstackId

`func (o *AdminImageAdminFields) GetOpenstackId() string`

GetOpenstackId returns the OpenstackId field if non-nil, zero value otherwise.

### GetOpenstackIdOk

`func (o *AdminImageAdminFields) GetOpenstackIdOk() (*string, bool)`

GetOpenstackIdOk returns a tuple with the OpenstackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackId

`func (o *AdminImageAdminFields) SetOpenstackId(v string)`

SetOpenstackId sets OpenstackId field to given value.

### HasOpenstackId

`func (o *AdminImageAdminFields) HasOpenstackId() bool`

HasOpenstackId returns a boolean if a field has been set.

### GetRegionId

`func (o *AdminImageAdminFields) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AdminImageAdminFields) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AdminImageAdminFields) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *AdminImageAdminFields) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetRegionName

`func (o *AdminImageAdminFields) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *AdminImageAdminFields) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *AdminImageAdminFields) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *AdminImageAdminFields) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetType

`func (o *AdminImageAdminFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdminImageAdminFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdminImageAdminFields) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdminImageAdminFields) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *AdminImageAdminFields) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdminImageAdminFields) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdminImageAdminFields) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdminImageAdminFields) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSize

`func (o *AdminImageAdminFields) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AdminImageAdminFields) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AdminImageAdminFields) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AdminImageAdminFields) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetIsPublic

`func (o *AdminImageAdminFields) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AdminImageAdminFields) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AdminImageAdminFields) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *AdminImageAdminFields) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetDescription

`func (o *AdminImageAdminFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminImageAdminFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminImageAdminFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminImageAdminFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *AdminImageAdminFields) GetLabels() []LableResonse`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AdminImageAdminFields) GetLabelsOk() (*[]LableResonse, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AdminImageAdminFields) SetLabels(v []LableResonse)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AdminImageAdminFields) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminImageAdminFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminImageAdminFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminImageAdminFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminImageAdminFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


