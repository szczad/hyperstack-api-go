# AdminOrganizationResourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**AdminOrganizationResources**](AdminOrganizationResources.md) |  | [optional] 

## Methods

### NewAdminOrganizationResourceResponse

`func NewAdminOrganizationResourceResponse() *AdminOrganizationResourceResponse`

NewAdminOrganizationResourceResponse instantiates a new AdminOrganizationResourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminOrganizationResourceResponseWithDefaults

`func NewAdminOrganizationResourceResponseWithDefaults() *AdminOrganizationResourceResponse`

NewAdminOrganizationResourceResponseWithDefaults instantiates a new AdminOrganizationResourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminOrganizationResourceResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminOrganizationResourceResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminOrganizationResourceResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminOrganizationResourceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AdminOrganizationResourceResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminOrganizationResourceResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminOrganizationResourceResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AdminOrganizationResourceResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrganization

`func (o *AdminOrganizationResourceResponse) GetOrganization() AdminOrganizationResources`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AdminOrganizationResourceResponse) GetOrganizationOk() (*AdminOrganizationResources, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AdminOrganizationResourceResponse) SetOrganization(v AdminOrganizationResources)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AdminOrganizationResourceResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


