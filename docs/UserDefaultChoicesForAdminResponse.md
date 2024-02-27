# UserDefaultChoicesForAdminResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**UserDefaultChoices** | Pointer to [**[]UserDefaultChoiceForAdminFields**](UserDefaultChoiceForAdminFields.md) |  | [optional] 

## Methods

### NewUserDefaultChoicesForAdminResponse

`func NewUserDefaultChoicesForAdminResponse() *UserDefaultChoicesForAdminResponse`

NewUserDefaultChoicesForAdminResponse instantiates a new UserDefaultChoicesForAdminResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefaultChoicesForAdminResponseWithDefaults

`func NewUserDefaultChoicesForAdminResponseWithDefaults() *UserDefaultChoicesForAdminResponse`

NewUserDefaultChoicesForAdminResponseWithDefaults instantiates a new UserDefaultChoicesForAdminResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserDefaultChoicesForAdminResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserDefaultChoicesForAdminResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserDefaultChoicesForAdminResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserDefaultChoicesForAdminResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *UserDefaultChoicesForAdminResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserDefaultChoicesForAdminResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserDefaultChoicesForAdminResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserDefaultChoicesForAdminResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetUserDefaultChoices

`func (o *UserDefaultChoicesForAdminResponse) GetUserDefaultChoices() []UserDefaultChoiceForAdminFields`

GetUserDefaultChoices returns the UserDefaultChoices field if non-nil, zero value otherwise.

### GetUserDefaultChoicesOk

`func (o *UserDefaultChoicesForAdminResponse) GetUserDefaultChoicesOk() (*[]UserDefaultChoiceForAdminFields, bool)`

GetUserDefaultChoicesOk returns a tuple with the UserDefaultChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefaultChoices

`func (o *UserDefaultChoicesForAdminResponse) SetUserDefaultChoices(v []UserDefaultChoiceForAdminFields)`

SetUserDefaultChoices sets UserDefaultChoices field to given value.

### HasUserDefaultChoices

`func (o *UserDefaultChoicesForAdminResponse) HasUserDefaultChoices() bool`

HasUserDefaultChoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


