# AdminOrganizationsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Organizations** | Pointer to [**[]OrganizationFields**](OrganizationFields.md) |  | [optional] 
**OrganizationCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewAdminOrganizationsResponseModel

`func NewAdminOrganizationsResponseModel() *AdminOrganizationsResponseModel`

NewAdminOrganizationsResponseModel instantiates a new AdminOrganizationsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminOrganizationsResponseModelWithDefaults

`func NewAdminOrganizationsResponseModelWithDefaults() *AdminOrganizationsResponseModel`

NewAdminOrganizationsResponseModelWithDefaults instantiates a new AdminOrganizationsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminOrganizationsResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminOrganizationsResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminOrganizationsResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminOrganizationsResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AdminOrganizationsResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminOrganizationsResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminOrganizationsResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AdminOrganizationsResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrganizations

`func (o *AdminOrganizationsResponseModel) GetOrganizations() []OrganizationFields`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *AdminOrganizationsResponseModel) GetOrganizationsOk() (*[]OrganizationFields, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *AdminOrganizationsResponseModel) SetOrganizations(v []OrganizationFields)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *AdminOrganizationsResponseModel) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetOrganizationCount

`func (o *AdminOrganizationsResponseModel) GetOrganizationCount() int32`

GetOrganizationCount returns the OrganizationCount field if non-nil, zero value otherwise.

### GetOrganizationCountOk

`func (o *AdminOrganizationsResponseModel) GetOrganizationCountOk() (*int32, bool)`

GetOrganizationCountOk returns a tuple with the OrganizationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCount

`func (o *AdminOrganizationsResponseModel) SetOrganizationCount(v int32)`

SetOrganizationCount sets OrganizationCount field to given value.

### HasOrganizationCount

`func (o *AdminOrganizationsResponseModel) HasOrganizationCount() bool`

HasOrganizationCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


