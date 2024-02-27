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
	"bytes"
	"fmt"
)

// checks if the CreateDiscountsPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDiscountsPayload{}

// CreateDiscountsPayload struct for CreateDiscountsPayload
type CreateDiscountsPayload struct {
	Customers []CustomerPayload `json:"customers"`
	DiscountResources []ResourcePayload `json:"discount_resources"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate *time.Time `json:"end_date,omitempty"`
	DiscountStatus string `json:"discount_status"`
}

type _CreateDiscountsPayload CreateDiscountsPayload

// NewCreateDiscountsPayload instantiates a new CreateDiscountsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDiscountsPayload(customers []CustomerPayload, discountResources []ResourcePayload, discountStatus string) *CreateDiscountsPayload {
	this := CreateDiscountsPayload{}
	this.Customers = customers
	this.DiscountResources = discountResources
	this.DiscountStatus = discountStatus
	return &this
}

// NewCreateDiscountsPayloadWithDefaults instantiates a new CreateDiscountsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDiscountsPayloadWithDefaults() *CreateDiscountsPayload {
	this := CreateDiscountsPayload{}
	return &this
}

// GetCustomers returns the Customers field value
func (o *CreateDiscountsPayload) GetCustomers() []CustomerPayload {
	if o == nil {
		var ret []CustomerPayload
		return ret
	}

	return o.Customers
}

// GetCustomersOk returns a tuple with the Customers field value
// and a boolean to check if the value has been set.
func (o *CreateDiscountsPayload) GetCustomersOk() ([]CustomerPayload, bool) {
	if o == nil {
		return nil, false
	}
	return o.Customers, true
}

// SetCustomers sets field value
func (o *CreateDiscountsPayload) SetCustomers(v []CustomerPayload) {
	o.Customers = v
}

// GetDiscountResources returns the DiscountResources field value
func (o *CreateDiscountsPayload) GetDiscountResources() []ResourcePayload {
	if o == nil {
		var ret []ResourcePayload
		return ret
	}

	return o.DiscountResources
}

// GetDiscountResourcesOk returns a tuple with the DiscountResources field value
// and a boolean to check if the value has been set.
func (o *CreateDiscountsPayload) GetDiscountResourcesOk() ([]ResourcePayload, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountResources, true
}

// SetDiscountResources sets field value
func (o *CreateDiscountsPayload) SetDiscountResources(v []ResourcePayload) {
	o.DiscountResources = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CreateDiscountsPayload) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountsPayload) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CreateDiscountsPayload) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *CreateDiscountsPayload) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CreateDiscountsPayload) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountsPayload) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CreateDiscountsPayload) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *CreateDiscountsPayload) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetDiscountStatus returns the DiscountStatus field value
func (o *CreateDiscountsPayload) GetDiscountStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscountStatus
}

// GetDiscountStatusOk returns a tuple with the DiscountStatus field value
// and a boolean to check if the value has been set.
func (o *CreateDiscountsPayload) GetDiscountStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountStatus, true
}

// SetDiscountStatus sets field value
func (o *CreateDiscountsPayload) SetDiscountStatus(v string) {
	o.DiscountStatus = v
}

func (o CreateDiscountsPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDiscountsPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customers"] = o.Customers
	toSerialize["discount_resources"] = o.DiscountResources
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	toSerialize["discount_status"] = o.DiscountStatus
	return toSerialize, nil
}

func (o *CreateDiscountsPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customers",
		"discount_resources",
		"discount_status",
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

	varCreateDiscountsPayload := _CreateDiscountsPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDiscountsPayload)

	if err != nil {
		return err
	}

	*o = CreateDiscountsPayload(varCreateDiscountsPayload)

	return err
}

type NullableCreateDiscountsPayload struct {
	value *CreateDiscountsPayload
	isSet bool
}

func (v NullableCreateDiscountsPayload) Get() *CreateDiscountsPayload {
	return v.value
}

func (v *NullableCreateDiscountsPayload) Set(val *CreateDiscountsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDiscountsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDiscountsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDiscountsPayload(val *CreateDiscountsPayload) *NullableCreateDiscountsPayload {
	return &NullableCreateDiscountsPayload{value: val, isSet: true}
}

func (v NullableCreateDiscountsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDiscountsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


