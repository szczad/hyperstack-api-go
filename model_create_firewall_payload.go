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

// checks if the CreateFirewallPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFirewallPayload{}

// CreateFirewallPayload struct for CreateFirewallPayload
type CreateFirewallPayload struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	EnvironmentId int32 `json:"environment_id"`
}

type _CreateFirewallPayload CreateFirewallPayload

// NewCreateFirewallPayload instantiates a new CreateFirewallPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFirewallPayload(name string, environmentId int32) *CreateFirewallPayload {
	this := CreateFirewallPayload{}
	this.Name = name
	this.EnvironmentId = environmentId
	return &this
}

// NewCreateFirewallPayloadWithDefaults instantiates a new CreateFirewallPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFirewallPayloadWithDefaults() *CreateFirewallPayload {
	this := CreateFirewallPayload{}
	return &this
}

// GetName returns the Name field value
func (o *CreateFirewallPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFirewallPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateFirewallPayload) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateFirewallPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirewallPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateFirewallPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateFirewallPayload) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *CreateFirewallPayload) GetEnvironmentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *CreateFirewallPayload) GetEnvironmentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *CreateFirewallPayload) SetEnvironmentId(v int32) {
	o.EnvironmentId = v
}

func (o CreateFirewallPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFirewallPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["environment_id"] = o.EnvironmentId
	return toSerialize, nil
}

func (o *CreateFirewallPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"environment_id",
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

	varCreateFirewallPayload := _CreateFirewallPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateFirewallPayload)

	if err != nil {
		return err
	}

	*o = CreateFirewallPayload(varCreateFirewallPayload)

	return err
}

type NullableCreateFirewallPayload struct {
	value *CreateFirewallPayload
	isSet bool
}

func (v NullableCreateFirewallPayload) Get() *CreateFirewallPayload {
	return v.value
}

func (v *NullableCreateFirewallPayload) Set(val *CreateFirewallPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFirewallPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFirewallPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFirewallPayload(val *CreateFirewallPayload) *NullableCreateFirewallPayload {
	return &NullableCreateFirewallPayload{value: val, isSet: true}
}

func (v NullableCreateFirewallPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFirewallPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

