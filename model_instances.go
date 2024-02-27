/*
Infrahub-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"encoding/json"
)

// checks if the Instances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Instances{}

// Instances struct for Instances
type Instances struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	Instances []InstanceFields `json:"instances,omitempty"`
}

// NewInstances instantiates a new Instances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstances() *Instances {
	this := Instances{}
	return &this
}

// NewInstancesWithDefaults instantiates a new Instances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstancesWithDefaults() *Instances {
	this := Instances{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Instances) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instances) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Instances) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *Instances) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Instances) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instances) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Instances) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Instances) SetMessage(v string) {
	o.Message = &v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *Instances) GetInstances() []InstanceFields {
	if o == nil || IsNil(o.Instances) {
		var ret []InstanceFields
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instances) GetInstancesOk() ([]InstanceFields, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *Instances) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []InstanceFields and assigns it to the Instances field.
func (o *Instances) SetInstances(v []InstanceFields) {
	o.Instances = v
}

func (o Instances) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Instances) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	return toSerialize, nil
}

type NullableInstances struct {
	value *Instances
	isSet bool
}

func (v NullableInstances) Get() *Instances {
	return v.value
}

func (v *NullableInstances) Set(val *Instances) {
	v.value = val
	v.isSet = true
}

func (v NullableInstances) IsSet() bool {
	return v.isSet
}

func (v *NullableInstances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstances(val *Instances) *NullableInstances {
	return &NullableInstances{value: val, isSet: true}
}

func (v NullableInstances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


