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

// checks if the CreateGPU type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGPU{}

// CreateGPU struct for CreateGPU
type CreateGPU struct {
	Name string `json:"name"`
	Regions []string `json:"regions,omitempty"`
	// A valid JSON string.
	ExampleMetadata *string `json:"example_metadata,omitempty"`
}

type _CreateGPU CreateGPU

// NewCreateGPU instantiates a new CreateGPU object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGPU(name string) *CreateGPU {
	this := CreateGPU{}
	this.Name = name
	return &this
}

// NewCreateGPUWithDefaults instantiates a new CreateGPU object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGPUWithDefaults() *CreateGPU {
	this := CreateGPU{}
	return &this
}

// GetName returns the Name field value
func (o *CreateGPU) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGPU) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGPU) SetName(v string) {
	o.Name = v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *CreateGPU) GetRegions() []string {
	if o == nil || IsNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGPU) GetRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *CreateGPU) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *CreateGPU) SetRegions(v []string) {
	o.Regions = v
}

// GetExampleMetadata returns the ExampleMetadata field value if set, zero value otherwise.
func (o *CreateGPU) GetExampleMetadata() string {
	if o == nil || IsNil(o.ExampleMetadata) {
		var ret string
		return ret
	}
	return *o.ExampleMetadata
}

// GetExampleMetadataOk returns a tuple with the ExampleMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGPU) GetExampleMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.ExampleMetadata) {
		return nil, false
	}
	return o.ExampleMetadata, true
}

// HasExampleMetadata returns a boolean if a field has been set.
func (o *CreateGPU) HasExampleMetadata() bool {
	if o != nil && !IsNil(o.ExampleMetadata) {
		return true
	}

	return false
}

// SetExampleMetadata gets a reference to the given string and assigns it to the ExampleMetadata field.
func (o *CreateGPU) SetExampleMetadata(v string) {
	o.ExampleMetadata = &v
}

func (o CreateGPU) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGPU) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !IsNil(o.ExampleMetadata) {
		toSerialize["example_metadata"] = o.ExampleMetadata
	}
	return toSerialize, nil
}

func (o *CreateGPU) UnmarshalJSON(data []byte) (err error) {
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

	varCreateGPU := _CreateGPU{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateGPU)

	if err != nil {
		return err
	}

	*o = CreateGPU(varCreateGPU)

	return err
}

type NullableCreateGPU struct {
	value *CreateGPU
	isSet bool
}

func (v NullableCreateGPU) Get() *CreateGPU {
	return v.value
}

func (v *NullableCreateGPU) Set(val *CreateGPU) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGPU) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGPU) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGPU(val *CreateGPU) *NullableCreateGPU {
	return &NullableCreateGPU{value: val, isSet: true}
}

func (v NullableCreateGPU) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGPU) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


