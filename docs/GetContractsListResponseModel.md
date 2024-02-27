# GetContractsListResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Contracts** | Pointer to [**[]AdminContractFields**](AdminContractFields.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetContractsListResponseModel

`func NewGetContractsListResponseModel() *GetContractsListResponseModel`

NewGetContractsListResponseModel instantiates a new GetContractsListResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContractsListResponseModelWithDefaults

`func NewGetContractsListResponseModelWithDefaults() *GetContractsListResponseModel`

NewGetContractsListResponseModelWithDefaults instantiates a new GetContractsListResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetContractsListResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetContractsListResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetContractsListResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetContractsListResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *GetContractsListResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetContractsListResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetContractsListResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetContractsListResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetContracts

`func (o *GetContractsListResponseModel) GetContracts() []AdminContractFields`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *GetContractsListResponseModel) GetContractsOk() (*[]AdminContractFields, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *GetContractsListResponseModel) SetContracts(v []AdminContractFields)`

SetContracts sets Contracts field to given value.

### HasContracts

`func (o *GetContractsListResponseModel) HasContracts() bool`

HasContracts returns a boolean if a field has been set.

### GetCount

`func (o *GetContractsListResponseModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetContractsListResponseModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetContractsListResponseModel) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetContractsListResponseModel) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


