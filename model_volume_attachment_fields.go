/*
Infrahub-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"encoding/json"
	"time"
)

// checks if the VolumeAttachmentFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeAttachmentFields{}

// VolumeAttachmentFields struct for VolumeAttachmentFields
type VolumeAttachmentFields struct {
	Volume *VolumeFieldsforInstance `json:"volume,omitempty"`
	Status *string `json:"status,omitempty"`
	Device *string `json:"device,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewVolumeAttachmentFields instantiates a new VolumeAttachmentFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeAttachmentFields() *VolumeAttachmentFields {
	this := VolumeAttachmentFields{}
	return &this
}

// NewVolumeAttachmentFieldsWithDefaults instantiates a new VolumeAttachmentFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeAttachmentFieldsWithDefaults() *VolumeAttachmentFields {
	this := VolumeAttachmentFields{}
	return &this
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *VolumeAttachmentFields) GetVolume() VolumeFieldsforInstance {
	if o == nil || IsNil(o.Volume) {
		var ret VolumeFieldsforInstance
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeAttachmentFields) GetVolumeOk() (*VolumeFieldsforInstance, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *VolumeAttachmentFields) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given VolumeFieldsforInstance and assigns it to the Volume field.
func (o *VolumeAttachmentFields) SetVolume(v VolumeFieldsforInstance) {
	o.Volume = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VolumeAttachmentFields) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeAttachmentFields) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VolumeAttachmentFields) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VolumeAttachmentFields) SetStatus(v string) {
	o.Status = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *VolumeAttachmentFields) GetDevice() string {
	if o == nil || IsNil(o.Device) {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeAttachmentFields) GetDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *VolumeAttachmentFields) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *VolumeAttachmentFields) SetDevice(v string) {
	o.Device = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VolumeAttachmentFields) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeAttachmentFields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VolumeAttachmentFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VolumeAttachmentFields) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o VolumeAttachmentFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeAttachmentFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableVolumeAttachmentFields struct {
	value *VolumeAttachmentFields
	isSet bool
}

func (v NullableVolumeAttachmentFields) Get() *VolumeAttachmentFields {
	return v.value
}

func (v *NullableVolumeAttachmentFields) Set(val *VolumeAttachmentFields) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeAttachmentFields) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeAttachmentFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeAttachmentFields(val *VolumeAttachmentFields) *NullableVolumeAttachmentFields {
	return &NullableVolumeAttachmentFields{value: val, isSet: true}
}

func (v NullableVolumeAttachmentFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeAttachmentFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

