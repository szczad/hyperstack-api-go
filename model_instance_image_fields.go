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

// checks if the InstanceImageFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceImageFields{}

// InstanceImageFields struct for InstanceImageFields
type InstanceImageFields struct {
	Name *string `json:"name,omitempty"`
}

// NewInstanceImageFields instantiates a new InstanceImageFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceImageFields() *InstanceImageFields {
	this := InstanceImageFields{}
	return &this
}

// NewInstanceImageFieldsWithDefaults instantiates a new InstanceImageFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceImageFieldsWithDefaults() *InstanceImageFields {
	this := InstanceImageFields{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceImageFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceImageFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceImageFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceImageFields) SetName(v string) {
	o.Name = &v
}

func (o InstanceImageFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceImageFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableInstanceImageFields struct {
	value *InstanceImageFields
	isSet bool
}

func (v NullableInstanceImageFields) Get() *InstanceImageFields {
	return v.value
}

func (v *NullableInstanceImageFields) Set(val *InstanceImageFields) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceImageFields) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceImageFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceImageFields(val *InstanceImageFields) *NullableInstanceImageFields {
	return &NullableInstanceImageFields{value: val, isSet: true}
}

func (v NullableInstanceImageFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceImageFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


