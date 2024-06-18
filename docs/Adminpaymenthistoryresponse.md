# Adminpaymenthistoryresponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**[]Adminpaymenthistoryfields**](Adminpaymenthistoryfields.md) |  | [optional] 

## Methods

### NewAdminpaymenthistoryresponse

`func NewAdminpaymenthistoryresponse() *Adminpaymenthistoryresponse`

NewAdminpaymenthistoryresponse instantiates a new Adminpaymenthistoryresponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminpaymenthistoryresponseWithDefaults

`func NewAdminpaymenthistoryresponseWithDefaults() *Adminpaymenthistoryresponse`

NewAdminpaymenthistoryresponseWithDefaults instantiates a new Adminpaymenthistoryresponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Adminpaymenthistoryresponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Adminpaymenthistoryresponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Adminpaymenthistoryresponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Adminpaymenthistoryresponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *Adminpaymenthistoryresponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Adminpaymenthistoryresponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Adminpaymenthistoryresponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Adminpaymenthistoryresponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *Adminpaymenthistoryresponse) GetData() []Adminpaymenthistoryfields`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Adminpaymenthistoryresponse) GetDataOk() (*[]Adminpaymenthistoryfields, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Adminpaymenthistoryresponse) SetData(v []Adminpaymenthistoryfields)`

SetData sets Data field to given value.

### HasData

`func (o *Adminpaymenthistoryresponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


