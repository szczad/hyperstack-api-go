# GetContractDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Contract** | Pointer to [**AdminGetContractDetailFields**](AdminGetContractDetailFields.md) |  | [optional] 
**DiscountPlans** | Pointer to [**[]DiscountPlanFields**](DiscountPlanFields.md) |  | [optional] 

## Methods

### NewGetContractDetailResponseModel

`func NewGetContractDetailResponseModel() *GetContractDetailResponseModel`

NewGetContractDetailResponseModel instantiates a new GetContractDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContractDetailResponseModelWithDefaults

`func NewGetContractDetailResponseModelWithDefaults() *GetContractDetailResponseModel`

NewGetContractDetailResponseModelWithDefaults instantiates a new GetContractDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetContractDetailResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetContractDetailResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetContractDetailResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetContractDetailResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *GetContractDetailResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetContractDetailResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetContractDetailResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetContractDetailResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetContract

`func (o *GetContractDetailResponseModel) GetContract() AdminGetContractDetailFields`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *GetContractDetailResponseModel) GetContractOk() (*AdminGetContractDetailFields, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *GetContractDetailResponseModel) SetContract(v AdminGetContractDetailFields)`

SetContract sets Contract field to given value.

### HasContract

`func (o *GetContractDetailResponseModel) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetDiscountPlans

`func (o *GetContractDetailResponseModel) GetDiscountPlans() []DiscountPlanFields`

GetDiscountPlans returns the DiscountPlans field if non-nil, zero value otherwise.

### GetDiscountPlansOk

`func (o *GetContractDetailResponseModel) GetDiscountPlansOk() (*[]DiscountPlanFields, bool)`

GetDiscountPlansOk returns a tuple with the DiscountPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPlans

`func (o *GetContractDetailResponseModel) SetDiscountPlans(v []DiscountPlanFields)`

SetDiscountPlans sets DiscountPlans field to given value.

### HasDiscountPlans

`func (o *GetContractDetailResponseModel) HasDiscountPlans() bool`

HasDiscountPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


