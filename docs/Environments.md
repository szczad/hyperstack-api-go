# Environments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Environments** | Pointer to [**[]EnvironmentFields**](EnvironmentFields.md) |  | [optional] 

## Methods

### NewEnvironments

`func NewEnvironments() *Environments`

NewEnvironments instantiates a new Environments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentsWithDefaults

`func NewEnvironmentsWithDefaults() *Environments`

NewEnvironmentsWithDefaults instantiates a new Environments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Environments) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Environments) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Environments) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Environments) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Environments) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Environments) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Environments) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Environments) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEnvironments

`func (o *Environments) GetEnvironments() []EnvironmentFields`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *Environments) GetEnvironmentsOk() (*[]EnvironmentFields, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *Environments) SetEnvironments(v []EnvironmentFields)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *Environments) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


