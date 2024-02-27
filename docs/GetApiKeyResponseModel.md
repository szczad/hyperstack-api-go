# GetApiKeyResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to [**ApiKeyFields**](ApiKeyFields.md) |  | [optional] 

## Methods

### NewGetApiKeyResponseModel

`func NewGetApiKeyResponseModel() *GetApiKeyResponseModel`

NewGetApiKeyResponseModel instantiates a new GetApiKeyResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiKeyResponseModelWithDefaults

`func NewGetApiKeyResponseModelWithDefaults() *GetApiKeyResponseModel`

NewGetApiKeyResponseModelWithDefaults instantiates a new GetApiKeyResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetApiKeyResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetApiKeyResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetApiKeyResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetApiKeyResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *GetApiKeyResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetApiKeyResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetApiKeyResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetApiKeyResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetApiKey

`func (o *GetApiKeyResponseModel) GetApiKey() ApiKeyFields`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GetApiKeyResponseModel) GetApiKeyOk() (*ApiKeyFields, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GetApiKeyResponseModel) SetApiKey(v ApiKeyFields)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GetApiKeyResponseModel) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


