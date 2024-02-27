# AdminOrganizationsSummaryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Organizations** | Pointer to [**[]AdminOrganizationSummaryFields**](AdminOrganizationSummaryFields.md) |  | [optional] 

## Methods

### NewAdminOrganizationsSummaryResponseModel

`func NewAdminOrganizationsSummaryResponseModel() *AdminOrganizationsSummaryResponseModel`

NewAdminOrganizationsSummaryResponseModel instantiates a new AdminOrganizationsSummaryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminOrganizationsSummaryResponseModelWithDefaults

`func NewAdminOrganizationsSummaryResponseModelWithDefaults() *AdminOrganizationsSummaryResponseModel`

NewAdminOrganizationsSummaryResponseModelWithDefaults instantiates a new AdminOrganizationsSummaryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminOrganizationsSummaryResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminOrganizationsSummaryResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminOrganizationsSummaryResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminOrganizationsSummaryResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AdminOrganizationsSummaryResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminOrganizationsSummaryResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminOrganizationsSummaryResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AdminOrganizationsSummaryResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrganizations

`func (o *AdminOrganizationsSummaryResponseModel) GetOrganizations() []AdminOrganizationSummaryFields`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *AdminOrganizationsSummaryResponseModel) GetOrganizationsOk() (*[]AdminOrganizationSummaryFields, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *AdminOrganizationsSummaryResponseModel) SetOrganizations(v []AdminOrganizationSummaryFields)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *AdminOrganizationsSummaryResponseModel) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


