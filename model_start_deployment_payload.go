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

// checks if the StartDeploymentPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartDeploymentPayload{}

// StartDeploymentPayload struct for StartDeploymentPayload
type StartDeploymentPayload struct {
	Name string `json:"name"`
	Description string `json:"description"`
	TemplateId int32 `json:"template_id"`
	Variables *map[string]string `json:"variables,omitempty"`
}

type _StartDeploymentPayload StartDeploymentPayload

// NewStartDeploymentPayload instantiates a new StartDeploymentPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartDeploymentPayload(name string, description string, templateId int32) *StartDeploymentPayload {
	this := StartDeploymentPayload{}
	this.Name = name
	this.Description = description
	this.TemplateId = templateId
	return &this
}

// NewStartDeploymentPayloadWithDefaults instantiates a new StartDeploymentPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartDeploymentPayloadWithDefaults() *StartDeploymentPayload {
	this := StartDeploymentPayload{}
	return &this
}

// GetName returns the Name field value
func (o *StartDeploymentPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StartDeploymentPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StartDeploymentPayload) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *StartDeploymentPayload) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *StartDeploymentPayload) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *StartDeploymentPayload) SetDescription(v string) {
	o.Description = v
}

// GetTemplateId returns the TemplateId field value
func (o *StartDeploymentPayload) GetTemplateId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *StartDeploymentPayload) GetTemplateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *StartDeploymentPayload) SetTemplateId(v int32) {
	o.TemplateId = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *StartDeploymentPayload) GetVariables() map[string]string {
	if o == nil || IsNil(o.Variables) {
		var ret map[string]string
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartDeploymentPayload) GetVariablesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *StartDeploymentPayload) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]string and assigns it to the Variables field.
func (o *StartDeploymentPayload) SetVariables(v map[string]string) {
	o.Variables = &v
}

func (o StartDeploymentPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartDeploymentPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["template_id"] = o.TemplateId
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

func (o *StartDeploymentPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"template_id",
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

	varStartDeploymentPayload := _StartDeploymentPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStartDeploymentPayload)

	if err != nil {
		return err
	}

	*o = StartDeploymentPayload(varStartDeploymentPayload)

	return err
}

type NullableStartDeploymentPayload struct {
	value *StartDeploymentPayload
	isSet bool
}

func (v NullableStartDeploymentPayload) Get() *StartDeploymentPayload {
	return v.value
}

func (v *NullableStartDeploymentPayload) Set(val *StartDeploymentPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableStartDeploymentPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableStartDeploymentPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartDeploymentPayload(val *StartDeploymentPayload) *NullableStartDeploymentPayload {
	return &NullableStartDeploymentPayload{value: val, isSet: true}
}

func (v NullableStartDeploymentPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartDeploymentPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

