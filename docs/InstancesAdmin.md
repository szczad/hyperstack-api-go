# InstancesAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to [**[]InstanceAdminFields**](InstanceAdminFields.md) |  | [optional] 
**InstanceCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewInstancesAdmin

`func NewInstancesAdmin() *InstancesAdmin`

NewInstancesAdmin instantiates a new InstancesAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesAdminWithDefaults

`func NewInstancesAdminWithDefaults() *InstancesAdmin`

NewInstancesAdminWithDefaults instantiates a new InstancesAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InstancesAdmin) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstancesAdmin) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstancesAdmin) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstancesAdmin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *InstancesAdmin) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InstancesAdmin) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InstancesAdmin) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InstancesAdmin) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetInstances

`func (o *InstancesAdmin) GetInstances() []InstanceAdminFields`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *InstancesAdmin) GetInstancesOk() (*[]InstanceAdminFields, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *InstancesAdmin) SetInstances(v []InstanceAdminFields)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *InstancesAdmin) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetInstanceCount

`func (o *InstancesAdmin) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *InstancesAdmin) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *InstancesAdmin) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *InstancesAdmin) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


