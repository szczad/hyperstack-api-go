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

// checks if the Excludebillingpostpayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Excludebillingpostpayload{}

// Excludebillingpostpayload struct for Excludebillingpostpayload
type Excludebillingpostpayload struct {
	// The ID of the resource which is being excluded from billing.
	ResourceId int32 `json:"resource_id"`
	// `true` excludes the resource from billing while `false` does not.
	Exclude bool `json:"exclude"`
}

type _Excludebillingpostpayload Excludebillingpostpayload

// NewExcludebillingpostpayload instantiates a new Excludebillingpostpayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExcludebillingpostpayload(resourceId int32, exclude bool) *Excludebillingpostpayload {
	this := Excludebillingpostpayload{}
	this.ResourceId = resourceId
	this.Exclude = exclude
	return &this
}

// NewExcludebillingpostpayloadWithDefaults instantiates a new Excludebillingpostpayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExcludebillingpostpayloadWithDefaults() *Excludebillingpostpayload {
	this := Excludebillingpostpayload{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *Excludebillingpostpayload) GetResourceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *Excludebillingpostpayload) GetResourceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *Excludebillingpostpayload) SetResourceId(v int32) {
	o.ResourceId = v
}

// GetExclude returns the Exclude field value
func (o *Excludebillingpostpayload) GetExclude() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value
// and a boolean to check if the value has been set.
func (o *Excludebillingpostpayload) GetExcludeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exclude, true
}

// SetExclude sets field value
func (o *Excludebillingpostpayload) SetExclude(v bool) {
	o.Exclude = v
}

func (o Excludebillingpostpayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Excludebillingpostpayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["exclude"] = o.Exclude
	return toSerialize, nil
}

func (o *Excludebillingpostpayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource_id",
		"exclude",
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

	varExcludebillingpostpayload := _Excludebillingpostpayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExcludebillingpostpayload)

	if err != nil {
		return err
	}

	*o = Excludebillingpostpayload(varExcludebillingpostpayload)

	return err
}

type NullableExcludebillingpostpayload struct {
	value *Excludebillingpostpayload
	isSet bool
}

func (v NullableExcludebillingpostpayload) Get() *Excludebillingpostpayload {
	return v.value
}

func (v *NullableExcludebillingpostpayload) Set(val *Excludebillingpostpayload) {
	v.value = val
	v.isSet = true
}

func (v NullableExcludebillingpostpayload) IsSet() bool {
	return v.isSet
}

func (v *NullableExcludebillingpostpayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExcludebillingpostpayload(val *Excludebillingpostpayload) *NullableExcludebillingpostpayload {
	return &NullableExcludebillingpostpayload{value: val, isSet: true}
}

func (v NullableExcludebillingpostpayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExcludebillingpostpayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


