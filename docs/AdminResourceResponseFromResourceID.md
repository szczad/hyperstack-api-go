# AdminResourceResponseFromResourceID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Resources** | Pointer to [**AdminResource**](AdminResource.md) |  | [optional] 

## Methods

### NewAdminResourceResponseFromResourceID

`func NewAdminResourceResponseFromResourceID() *AdminResourceResponseFromResourceID`

NewAdminResourceResponseFromResourceID instantiates a new AdminResourceResponseFromResourceID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminResourceResponseFromResourceIDWithDefaults

`func NewAdminResourceResponseFromResourceIDWithDefaults() *AdminResourceResponseFromResourceID`

NewAdminResourceResponseFromResourceIDWithDefaults instantiates a new AdminResourceResponseFromResourceID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminResourceResponseFromResourceID) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminResourceResponseFromResourceID) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminResourceResponseFromResourceID) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminResourceResponseFromResourceID) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AdminResourceResponseFromResourceID) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminResourceResponseFromResourceID) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminResourceResponseFromResourceID) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AdminResourceResponseFromResourceID) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AdminResourceResponseFromResourceID) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AdminResourceResponseFromResourceID) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AdminResourceResponseFromResourceID) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AdminResourceResponseFromResourceID) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetResources

`func (o *AdminResourceResponseFromResourceID) GetResources() AdminResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AdminResourceResponseFromResourceID) GetResourcesOk() (*AdminResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AdminResourceResponseFromResourceID) SetResources(v AdminResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AdminResourceResponseFromResourceID) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


