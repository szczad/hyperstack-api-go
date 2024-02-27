# FlavorAdminResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Flavor** | Pointer to [**FlavorAdminResponseFlavors**](FlavorAdminResponseFlavors.md) |  | [optional] 

## Methods

### NewFlavorAdminResponse

`func NewFlavorAdminResponse() *FlavorAdminResponse`

NewFlavorAdminResponse instantiates a new FlavorAdminResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorAdminResponseWithDefaults

`func NewFlavorAdminResponseWithDefaults() *FlavorAdminResponse`

NewFlavorAdminResponseWithDefaults instantiates a new FlavorAdminResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FlavorAdminResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlavorAdminResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlavorAdminResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlavorAdminResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *FlavorAdminResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FlavorAdminResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FlavorAdminResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FlavorAdminResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFlavor

`func (o *FlavorAdminResponse) GetFlavor() FlavorAdminResponseFlavors`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *FlavorAdminResponse) GetFlavorOk() (*FlavorAdminResponseFlavors, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *FlavorAdminResponse) SetFlavor(v FlavorAdminResponseFlavors)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *FlavorAdminResponse) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


