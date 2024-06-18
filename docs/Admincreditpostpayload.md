# Admincreditpostpayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credit** | Pointer to **float32** | The credit value in dollars. | [optional] 
**Reason** | Pointer to **string** | Reason for the recharge. | [optional] 

## Methods

### NewAdmincreditpostpayload

`func NewAdmincreditpostpayload() *Admincreditpostpayload`

NewAdmincreditpostpayload instantiates a new Admincreditpostpayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmincreditpostpayloadWithDefaults

`func NewAdmincreditpostpayloadWithDefaults() *Admincreditpostpayload`

NewAdmincreditpostpayloadWithDefaults instantiates a new Admincreditpostpayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredit

`func (o *Admincreditpostpayload) GetCredit() float32`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *Admincreditpostpayload) GetCreditOk() (*float32, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *Admincreditpostpayload) SetCredit(v float32)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *Admincreditpostpayload) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetReason

`func (o *Admincreditpostpayload) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Admincreditpostpayload) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Admincreditpostpayload) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Admincreditpostpayload) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


