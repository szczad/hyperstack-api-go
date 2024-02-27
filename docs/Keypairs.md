# Keypairs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Keypairs** | Pointer to [**[]KeypairFields**](KeypairFields.md) |  | [optional] 

## Methods

### NewKeypairs

`func NewKeypairs() *Keypairs`

NewKeypairs instantiates a new Keypairs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeypairsWithDefaults

`func NewKeypairsWithDefaults() *Keypairs`

NewKeypairsWithDefaults instantiates a new Keypairs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Keypairs) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Keypairs) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Keypairs) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Keypairs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Keypairs) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Keypairs) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Keypairs) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Keypairs) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetKeypairs

`func (o *Keypairs) GetKeypairs() []KeypairFields`

GetKeypairs returns the Keypairs field if non-nil, zero value otherwise.

### GetKeypairsOk

`func (o *Keypairs) GetKeypairsOk() (*[]KeypairFields, bool)`

GetKeypairsOk returns a tuple with the Keypairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairs

`func (o *Keypairs) SetKeypairs(v []KeypairFields)`

SetKeypairs sets Keypairs field to given value.

### HasKeypairs

`func (o *Keypairs) HasKeypairs() bool`

HasKeypairs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


