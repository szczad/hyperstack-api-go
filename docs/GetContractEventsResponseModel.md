# GetContractEventsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ContractEvents** | Pointer to [**[]AdminContractEventFields**](AdminContractEventFields.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetContractEventsResponseModel

`func NewGetContractEventsResponseModel() *GetContractEventsResponseModel`

NewGetContractEventsResponseModel instantiates a new GetContractEventsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContractEventsResponseModelWithDefaults

`func NewGetContractEventsResponseModelWithDefaults() *GetContractEventsResponseModel`

NewGetContractEventsResponseModelWithDefaults instantiates a new GetContractEventsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetContractEventsResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetContractEventsResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetContractEventsResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetContractEventsResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *GetContractEventsResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetContractEventsResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetContractEventsResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetContractEventsResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetContractEvents

`func (o *GetContractEventsResponseModel) GetContractEvents() []AdminContractEventFields`

GetContractEvents returns the ContractEvents field if non-nil, zero value otherwise.

### GetContractEventsOk

`func (o *GetContractEventsResponseModel) GetContractEventsOk() (*[]AdminContractEventFields, bool)`

GetContractEventsOk returns a tuple with the ContractEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractEvents

`func (o *GetContractEventsResponseModel) SetContractEvents(v []AdminContractEventFields)`

SetContractEvents sets ContractEvents field to given value.

### HasContractEvents

`func (o *GetContractEventsResponseModel) HasContractEvents() bool`

HasContractEvents returns a boolean if a field has been set.

### GetCount

`func (o *GetContractEventsResponseModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetContractEventsResponseModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetContractEventsResponseModel) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetContractEventsResponseModel) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


