# AdminImageAdminResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Image** | Pointer to [**AdminImageAdminResponseImage**](AdminImageAdminResponseImage.md) |  | [optional] 

## Methods

### NewAdminImageAdminResponse

`func NewAdminImageAdminResponse() *AdminImageAdminResponse`

NewAdminImageAdminResponse instantiates a new AdminImageAdminResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminImageAdminResponseWithDefaults

`func NewAdminImageAdminResponseWithDefaults() *AdminImageAdminResponse`

NewAdminImageAdminResponseWithDefaults instantiates a new AdminImageAdminResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminImageAdminResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminImageAdminResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminImageAdminResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminImageAdminResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AdminImageAdminResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminImageAdminResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminImageAdminResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AdminImageAdminResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetImage

`func (o *AdminImageAdminResponse) GetImage() AdminImageAdminResponseImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AdminImageAdminResponse) GetImageOk() (*AdminImageAdminResponseImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AdminImageAdminResponse) SetImage(v AdminImageAdminResponseImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *AdminImageAdminResponse) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


