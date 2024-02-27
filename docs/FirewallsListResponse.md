# FirewallsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Firewalls** | Pointer to [**[]FirewallDetailFields**](FirewallDetailFields.md) |  | [optional] 

## Methods

### NewFirewallsListResponse

`func NewFirewallsListResponse() *FirewallsListResponse`

NewFirewallsListResponse instantiates a new FirewallsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallsListResponseWithDefaults

`func NewFirewallsListResponseWithDefaults() *FirewallsListResponse`

NewFirewallsListResponseWithDefaults instantiates a new FirewallsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FirewallsListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallsListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallsListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirewallsListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *FirewallsListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FirewallsListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FirewallsListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FirewallsListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFirewalls

`func (o *FirewallsListResponse) GetFirewalls() []FirewallDetailFields`

GetFirewalls returns the Firewalls field if non-nil, zero value otherwise.

### GetFirewallsOk

`func (o *FirewallsListResponse) GetFirewallsOk() (*[]FirewallDetailFields, bool)`

GetFirewallsOk returns a tuple with the Firewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewalls

`func (o *FirewallsListResponse) SetFirewalls(v []FirewallDetailFields)`

SetFirewalls sets Firewalls field to given value.

### HasFirewalls

`func (o *FirewallsListResponse) HasFirewalls() bool`

HasFirewalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


