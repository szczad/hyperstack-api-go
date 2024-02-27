# Volumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]VolumeFields**](VolumeFields.md) |  | [optional] 

## Methods

### NewVolumes

`func NewVolumes() *Volumes`

NewVolumes instantiates a new Volumes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumesWithDefaults

`func NewVolumesWithDefaults() *Volumes`

NewVolumesWithDefaults instantiates a new Volumes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Volumes) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Volumes) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Volumes) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Volumes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Volumes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Volumes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Volumes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Volumes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetVolumes

`func (o *Volumes) GetVolumes() []VolumeFields`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *Volumes) GetVolumesOk() (*[]VolumeFields, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *Volumes) SetVolumes(v []VolumeFields)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *Volumes) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


