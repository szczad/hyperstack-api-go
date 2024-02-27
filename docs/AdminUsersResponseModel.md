# AdminUsersResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]AdminUserFields**](AdminUserFields.md) |  | [optional] 

## Methods

### NewAdminUsersResponseModel

`func NewAdminUsersResponseModel() *AdminUsersResponseModel`

NewAdminUsersResponseModel instantiates a new AdminUsersResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUsersResponseModelWithDefaults

`func NewAdminUsersResponseModelWithDefaults() *AdminUsersResponseModel`

NewAdminUsersResponseModelWithDefaults instantiates a new AdminUsersResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminUsersResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminUsersResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminUsersResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminUsersResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AdminUsersResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminUsersResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminUsersResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AdminUsersResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetUsers

`func (o *AdminUsersResponseModel) GetUsers() []AdminUserFields`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AdminUsersResponseModel) GetUsersOk() (*[]AdminUserFields, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AdminUsersResponseModel) SetUsers(v []AdminUserFields)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AdminUsersResponseModel) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


