# InstancesSummaryAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to [**[]InstancesSummaryFields**](InstancesSummaryFields.md) |  | [optional] 

## Methods

### NewInstancesSummaryAdmin

`func NewInstancesSummaryAdmin() *InstancesSummaryAdmin`

NewInstancesSummaryAdmin instantiates a new InstancesSummaryAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesSummaryAdminWithDefaults

`func NewInstancesSummaryAdminWithDefaults() *InstancesSummaryAdmin`

NewInstancesSummaryAdminWithDefaults instantiates a new InstancesSummaryAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InstancesSummaryAdmin) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstancesSummaryAdmin) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstancesSummaryAdmin) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstancesSummaryAdmin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *InstancesSummaryAdmin) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InstancesSummaryAdmin) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InstancesSummaryAdmin) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InstancesSummaryAdmin) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetInstances

`func (o *InstancesSummaryAdmin) GetInstances() []InstancesSummaryFields`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *InstancesSummaryAdmin) GetInstancesOk() (*[]InstancesSummaryFields, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *InstancesSummaryAdmin) SetInstances(v []InstancesSummaryFields)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *InstancesSummaryAdmin) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


