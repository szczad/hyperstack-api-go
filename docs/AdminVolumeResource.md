# AdminVolumeResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OpenstackId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Bootable** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAdminVolumeResource

`func NewAdminVolumeResource() *AdminVolumeResource`

NewAdminVolumeResource instantiates a new AdminVolumeResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminVolumeResourceWithDefaults

`func NewAdminVolumeResourceWithDefaults() *AdminVolumeResource`

NewAdminVolumeResourceWithDefaults instantiates a new AdminVolumeResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminVolumeResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminVolumeResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminVolumeResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdminVolumeResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AdminVolumeResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminVolumeResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminVolumeResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminVolumeResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpenstackId

`func (o *AdminVolumeResource) GetOpenstackId() string`

GetOpenstackId returns the OpenstackId field if non-nil, zero value otherwise.

### GetOpenstackIdOk

`func (o *AdminVolumeResource) GetOpenstackIdOk() (*string, bool)`

GetOpenstackIdOk returns a tuple with the OpenstackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackId

`func (o *AdminVolumeResource) SetOpenstackId(v string)`

SetOpenstackId sets OpenstackId field to given value.

### HasOpenstackId

`func (o *AdminVolumeResource) HasOpenstackId() bool`

HasOpenstackId returns a boolean if a field has been set.

### GetType

`func (o *AdminVolumeResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdminVolumeResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdminVolumeResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdminVolumeResource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *AdminVolumeResource) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AdminVolumeResource) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AdminVolumeResource) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AdminVolumeResource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *AdminVolumeResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminVolumeResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminVolumeResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminVolumeResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBootable

`func (o *AdminVolumeResource) GetBootable() bool`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *AdminVolumeResource) GetBootableOk() (*bool, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *AdminVolumeResource) SetBootable(v bool)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *AdminVolumeResource) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminVolumeResource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminVolumeResource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminVolumeResource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminVolumeResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


