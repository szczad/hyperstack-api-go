# AdminCreateContractResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Contract** | Pointer to [**CreateContarctFields**](CreateContarctFields.md) |  | [optional] 

## Methods

### NewAdminCreateContractResponseModel

`func NewAdminCreateContractResponseModel() *AdminCreateContractResponseModel`

NewAdminCreateContractResponseModel instantiates a new AdminCreateContractResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminCreateContractResponseModelWithDefaults

`func NewAdminCreateContractResponseModelWithDefaults() *AdminCreateContractResponseModel`

NewAdminCreateContractResponseModelWithDefaults instantiates a new AdminCreateContractResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminCreateContractResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminCreateContractResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminCreateContractResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminCreateContractResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AdminCreateContractResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminCreateContractResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminCreateContractResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AdminCreateContractResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetContract

`func (o *AdminCreateContractResponseModel) GetContract() CreateContarctFields`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *AdminCreateContractResponseModel) GetContractOk() (*CreateContarctFields, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *AdminCreateContractResponseModel) SetContract(v CreateContarctFields)`

SetContract sets Contract field to given value.

### HasContract

`func (o *AdminCreateContractResponseModel) HasContract() bool`

HasContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


