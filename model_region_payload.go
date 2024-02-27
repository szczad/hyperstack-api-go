/*
Infrahub-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RegionPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionPayload{}

// RegionPayload struct for RegionPayload
type RegionPayload struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
}

type _RegionPayload RegionPayload

// NewRegionPayload instantiates a new RegionPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionPayload(name string) *RegionPayload {
	this := RegionPayload{}
	this.Name = name
	return &this
}

// NewRegionPayloadWithDefaults instantiates a new RegionPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionPayloadWithDefaults() *RegionPayload {
	this := RegionPayload{}
	return &this
}

// GetName returns the Name field value
func (o *RegionPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RegionPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RegionPayload) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RegionPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RegionPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RegionPayload) SetDescription(v string) {
	o.Description = &v
}

func (o RegionPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *RegionPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRegionPayload := _RegionPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegionPayload)

	if err != nil {
		return err
	}

	*o = RegionPayload(varRegionPayload)

	return err
}

type NullableRegionPayload struct {
	value *RegionPayload
	isSet bool
}

func (v NullableRegionPayload) Get() *RegionPayload {
	return v.value
}

func (v *NullableRegionPayload) Set(val *RegionPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionPayload(val *RegionPayload) *NullableRegionPayload {
	return &NullableRegionPayload{value: val, isSet: true}
}

func (v NullableRegionPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


