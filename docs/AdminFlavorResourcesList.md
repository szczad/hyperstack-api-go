# AdminFlavorResourcesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Flavors** | Pointer to [**[]AdminFlavorResource**](AdminFlavorResource.md) |  | [optional] 

## Methods

### NewAdminFlavorResourcesList

`func NewAdminFlavorResourcesList() *AdminFlavorResourcesList`

NewAdminFlavorResourcesList instantiates a new AdminFlavorResourcesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminFlavorResourcesListWithDefaults

`func NewAdminFlavorResourcesListWithDefaults() *AdminFlavorResourcesList`

NewAdminFlavorResourcesListWithDefaults instantiates a new AdminFlavorResourcesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminFlavorResourcesList) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminFlavorResourcesList) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminFlavorResourcesList) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminFlavorResourcesList) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AdminFlavorResourcesList) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminFlavorResourcesList) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminFlavorResourcesList) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AdminFlavorResourcesList) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFlavors

`func (o *AdminFlavorResourcesList) GetFlavors() []AdminFlavorResource`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *AdminFlavorResourcesList) GetFlavorsOk() (*[]AdminFlavorResource, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *AdminFlavorResourcesList) SetFlavors(v []AdminFlavorResource)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *AdminFlavorResourcesList) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


