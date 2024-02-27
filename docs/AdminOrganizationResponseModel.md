# AdminOrganizationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Organizations** | Pointer to [**OrganizationFields**](OrganizationFields.md) |  | [optional] 

## Methods

### NewAdminOrganizationResponseModel

`func NewAdminOrganizationResponseModel() *AdminOrganizationResponseModel`

NewAdminOrganizationResponseModel instantiates a new AdminOrganizationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminOrganizationResponseModelWithDefaults

`func NewAdminOrganizationResponseModelWithDefaults() *AdminOrganizationResponseModel`

NewAdminOrganizationResponseModelWithDefaults instantiates a new AdminOrganizationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminOrganizationResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminOrganizationResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminOrganizationResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminOrganizationResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AdminOrganizationResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminOrganizationResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminOrganizationResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AdminOrganizationResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrganizations

`func (o *AdminOrganizationResponseModel) GetOrganizations() OrganizationFields`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *AdminOrganizationResponseModel) GetOrganizationsOk() (*OrganizationFields, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *AdminOrganizationResponseModel) SetOrganizations(v OrganizationFields)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *AdminOrganizationResponseModel) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


