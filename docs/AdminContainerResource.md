# AdminContainerResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Flavor** | Pointer to [**AdminFlavorResource**](AdminFlavorResource.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**FixedIp** | Pointer to **string** |  | [optional] 
**FloatingIp** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAdminContainerResource

`func NewAdminContainerResource() *AdminContainerResource`

NewAdminContainerResource instantiates a new AdminContainerResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminContainerResourceWithDefaults

`func NewAdminContainerResourceWithDefaults() *AdminContainerResource`

NewAdminContainerResourceWithDefaults instantiates a new AdminContainerResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminContainerResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminContainerResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminContainerResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdminContainerResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AdminContainerResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminContainerResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminContainerResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminContainerResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlavor

`func (o *AdminContainerResource) GetFlavor() AdminFlavorResource`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *AdminContainerResource) GetFlavorOk() (*AdminFlavorResource, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *AdminContainerResource) SetFlavor(v AdminFlavorResource)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *AdminContainerResource) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetStatus

`func (o *AdminContainerResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminContainerResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminContainerResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminContainerResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImage

`func (o *AdminContainerResource) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AdminContainerResource) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AdminContainerResource) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AdminContainerResource) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetFixedIp

`func (o *AdminContainerResource) GetFixedIp() string`

GetFixedIp returns the FixedIp field if non-nil, zero value otherwise.

### GetFixedIpOk

`func (o *AdminContainerResource) GetFixedIpOk() (*string, bool)`

GetFixedIpOk returns a tuple with the FixedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIp

`func (o *AdminContainerResource) SetFixedIp(v string)`

SetFixedIp sets FixedIp field to given value.

### HasFixedIp

`func (o *AdminContainerResource) HasFixedIp() bool`

HasFixedIp returns a boolean if a field has been set.

### GetFloatingIp

`func (o *AdminContainerResource) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *AdminContainerResource) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *AdminContainerResource) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *AdminContainerResource) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminContainerResource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminContainerResource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminContainerResource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminContainerResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


