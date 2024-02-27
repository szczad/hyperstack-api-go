# AdminImageListAdminResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]AdminImageAdminFields**](AdminImageAdminFields.md) |  | [optional] 

## Methods

### NewAdminImageListAdminResponse

`func NewAdminImageListAdminResponse() *AdminImageListAdminResponse`

NewAdminImageListAdminResponse instantiates a new AdminImageListAdminResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminImageListAdminResponseWithDefaults

`func NewAdminImageListAdminResponseWithDefaults() *AdminImageListAdminResponse`

NewAdminImageListAdminResponseWithDefaults instantiates a new AdminImageListAdminResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminImageListAdminResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminImageListAdminResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminImageListAdminResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminImageListAdminResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AdminImageListAdminResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminImageListAdminResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminImageListAdminResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AdminImageListAdminResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *AdminImageListAdminResponse) GetData() []AdminImageAdminFields`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdminImageListAdminResponse) GetDataOk() (*[]AdminImageAdminFields, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdminImageListAdminResponse) SetData(v []AdminImageAdminFields)`

SetData sets Data field to given value.

### HasData

`func (o *AdminImageListAdminResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


