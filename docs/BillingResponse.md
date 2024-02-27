# BillingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**CalculatedBills** | Pointer to [**[]OrganizationObjectResponse**](OrganizationObjectResponse.md) |  | [optional] 

## Methods

### NewBillingResponse

`func NewBillingResponse() *BillingResponse`

NewBillingResponse instantiates a new BillingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingResponseWithDefaults

`func NewBillingResponseWithDefaults() *BillingResponse`

NewBillingResponseWithDefaults instantiates a new BillingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BillingResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BillingResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCalculatedBills

`func (o *BillingResponse) GetCalculatedBills() []OrganizationObjectResponse`

GetCalculatedBills returns the CalculatedBills field if non-nil, zero value otherwise.

### GetCalculatedBillsOk

`func (o *BillingResponse) GetCalculatedBillsOk() (*[]OrganizationObjectResponse, bool)`

GetCalculatedBillsOk returns a tuple with the CalculatedBills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedBills

`func (o *BillingResponse) SetCalculatedBills(v []OrganizationObjectResponse)`

SetCalculatedBills sets CalculatedBills field to given value.

### HasCalculatedBills

`func (o *BillingResponse) HasCalculatedBills() bool`

HasCalculatedBills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


