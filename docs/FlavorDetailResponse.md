# FlavorDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Flavor** | Pointer to [**AdminFlavorDetailFields**](AdminFlavorDetailFields.md) |  | [optional] 

## Methods

### NewFlavorDetailResponse

`func NewFlavorDetailResponse() *FlavorDetailResponse`

NewFlavorDetailResponse instantiates a new FlavorDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorDetailResponseWithDefaults

`func NewFlavorDetailResponseWithDefaults() *FlavorDetailResponse`

NewFlavorDetailResponseWithDefaults instantiates a new FlavorDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FlavorDetailResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlavorDetailResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlavorDetailResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlavorDetailResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *FlavorDetailResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FlavorDetailResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FlavorDetailResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FlavorDetailResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFlavor

`func (o *FlavorDetailResponse) GetFlavor() AdminFlavorDetailFields`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *FlavorDetailResponse) GetFlavorOk() (*AdminFlavorDetailFields, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *FlavorDetailResponse) SetFlavor(v AdminFlavorDetailFields)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *FlavorDetailResponse) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


