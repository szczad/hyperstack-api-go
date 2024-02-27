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

// checks if the ResourcePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourcePayload{}

// ResourcePayload struct for ResourcePayload
type ResourcePayload struct {
	ResourceId int32 `json:"resource_id"`
	DiscountPercent *float32 `json:"discount_percent,omitempty"`
	DiscountType string `json:"discount_type"`
	DiscountAmount *float32 `json:"discount_amount,omitempty"`
	ResourceCount *int32 `json:"resource_count,omitempty"`
}

type _ResourcePayload ResourcePayload

// NewResourcePayload instantiates a new ResourcePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcePayload(resourceId int32, discountType string) *ResourcePayload {
	this := ResourcePayload{}
	this.ResourceId = resourceId
	this.DiscountType = discountType
	return &this
}

// NewResourcePayloadWithDefaults instantiates a new ResourcePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcePayloadWithDefaults() *ResourcePayload {
	this := ResourcePayload{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *ResourcePayload) GetResourceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ResourcePayload) GetResourceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *ResourcePayload) SetResourceId(v int32) {
	o.ResourceId = v
}

// GetDiscountPercent returns the DiscountPercent field value if set, zero value otherwise.
func (o *ResourcePayload) GetDiscountPercent() float32 {
	if o == nil || IsNil(o.DiscountPercent) {
		var ret float32
		return ret
	}
	return *o.DiscountPercent
}

// GetDiscountPercentOk returns a tuple with the DiscountPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePayload) GetDiscountPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountPercent) {
		return nil, false
	}
	return o.DiscountPercent, true
}

// HasDiscountPercent returns a boolean if a field has been set.
func (o *ResourcePayload) HasDiscountPercent() bool {
	if o != nil && !IsNil(o.DiscountPercent) {
		return true
	}

	return false
}

// SetDiscountPercent gets a reference to the given float32 and assigns it to the DiscountPercent field.
func (o *ResourcePayload) SetDiscountPercent(v float32) {
	o.DiscountPercent = &v
}

// GetDiscountType returns the DiscountType field value
func (o *ResourcePayload) GetDiscountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscountType
}

// GetDiscountTypeOk returns a tuple with the DiscountType field value
// and a boolean to check if the value has been set.
func (o *ResourcePayload) GetDiscountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountType, true
}

// SetDiscountType sets field value
func (o *ResourcePayload) SetDiscountType(v string) {
	o.DiscountType = v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *ResourcePayload) GetDiscountAmount() float32 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret float32
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePayload) GetDiscountAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *ResourcePayload) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given float32 and assigns it to the DiscountAmount field.
func (o *ResourcePayload) SetDiscountAmount(v float32) {
	o.DiscountAmount = &v
}

// GetResourceCount returns the ResourceCount field value if set, zero value otherwise.
func (o *ResourcePayload) GetResourceCount() int32 {
	if o == nil || IsNil(o.ResourceCount) {
		var ret int32
		return ret
	}
	return *o.ResourceCount
}

// GetResourceCountOk returns a tuple with the ResourceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePayload) GetResourceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ResourceCount) {
		return nil, false
	}
	return o.ResourceCount, true
}

// HasResourceCount returns a boolean if a field has been set.
func (o *ResourcePayload) HasResourceCount() bool {
	if o != nil && !IsNil(o.ResourceCount) {
		return true
	}

	return false
}

// SetResourceCount gets a reference to the given int32 and assigns it to the ResourceCount field.
func (o *ResourcePayload) SetResourceCount(v int32) {
	o.ResourceCount = &v
}

func (o ResourcePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource_id"] = o.ResourceId
	if !IsNil(o.DiscountPercent) {
		toSerialize["discount_percent"] = o.DiscountPercent
	}
	toSerialize["discount_type"] = o.DiscountType
	if !IsNil(o.DiscountAmount) {
		toSerialize["discount_amount"] = o.DiscountAmount
	}
	if !IsNil(o.ResourceCount) {
		toSerialize["resource_count"] = o.ResourceCount
	}
	return toSerialize, nil
}

func (o *ResourcePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource_id",
		"discount_type",
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

	varResourcePayload := _ResourcePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourcePayload)

	if err != nil {
		return err
	}

	*o = ResourcePayload(varResourcePayload)

	return err
}

type NullableResourcePayload struct {
	value *ResourcePayload
	isSet bool
}

func (v NullableResourcePayload) Get() *ResourcePayload {
	return v.value
}

func (v *NullableResourcePayload) Set(val *ResourcePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcePayload(val *ResourcePayload) *NullableResourcePayload {
	return &NullableResourcePayload{value: val, isSet: true}
}

func (v NullableResourcePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


