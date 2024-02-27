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

// checks if the PricebookModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricebookModel{}

// PricebookModel struct for PricebookModel
type PricebookModel struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Value *float32 `json:"value,omitempty"`
	OriginalValue *float32 `json:"original_value,omitempty"`
	DiscountApplied *bool `json:"discount_applied,omitempty"`
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime *time.Time `json:"end_time,omitempty"`
}

// NewPricebookModel instantiates a new PricebookModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricebookModel() *PricebookModel {
	this := PricebookModel{}
	return &this
}

// NewPricebookModelWithDefaults instantiates a new PricebookModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricebookModelWithDefaults() *PricebookModel {
	this := PricebookModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PricebookModel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookModel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PricebookModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PricebookModel) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PricebookModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PricebookModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PricebookModel) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PricebookModel) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookModel) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PricebookModel) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *PricebookModel) SetValue(v float32) {
	o.Value = &v
}

// GetOriginalValue returns the OriginalValue field value if set, zero value otherwise.
func (o *PricebookModel) GetOriginalValue() float32 {
	if o == nil || IsNil(o.OriginalValue) {
		var ret float32
		return ret
	}
	return *o.OriginalValue
}

// GetOriginalValueOk returns a tuple with the OriginalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookModel) GetOriginalValueOk() (*float32, bool) {
	if o == nil || IsNil(o.OriginalValue) {
		return nil, false
	}
	return o.OriginalValue, true
}

// HasOriginalValue returns a boolean if a field has been set.
func (o *PricebookModel) HasOriginalValue() bool {
	if o != nil && !IsNil(o.OriginalValue) {
		return true
	}

	return false
}

// SetOriginalValue gets a reference to the given float32 and assigns it to the OriginalValue field.
func (o *PricebookModel) SetOriginalValue(v float32) {
	o.OriginalValue = &v
}

// GetDiscountApplied returns the DiscountApplied field value if set, zero value otherwise.
func (o *PricebookModel) GetDiscountApplied() bool {
	if o == nil || IsNil(o.DiscountApplied) {
		var ret bool
		return ret
	}
	return *o.DiscountApplied
}

// GetDiscountAppliedOk returns a tuple with the DiscountApplied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookModel) GetDiscountAppliedOk() (*bool, bool) {
	if o == nil || IsNil(o.DiscountApplied) {
		return nil, false
	}
	return o.DiscountApplied, true
}

// HasDiscountApplied returns a boolean if a field has been set.
func (o *PricebookModel) HasDiscountApplied() bool {
	if o != nil && !IsNil(o.DiscountApplied) {
		return true
	}

	return false
}

// SetDiscountApplied gets a reference to the given bool and assigns it to the DiscountApplied field.
func (o *PricebookModel) SetDiscountApplied(v bool) {
	o.DiscountApplied = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *PricebookModel) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookModel) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *PricebookModel) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *PricebookModel) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *PricebookModel) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookModel) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *PricebookModel) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *PricebookModel) SetEndTime(v time.Time) {
	o.EndTime = &v
}

func (o PricebookModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricebookModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.OriginalValue) {
		toSerialize["original_value"] = o.OriginalValue
	}
	if !IsNil(o.DiscountApplied) {
		toSerialize["discount_applied"] = o.DiscountApplied
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	return toSerialize, nil
}

type NullablePricebookModel struct {
	value *PricebookModel
	isSet bool
}

func (v NullablePricebookModel) Get() *PricebookModel {
	return v.value
}

func (v *NullablePricebookModel) Set(val *PricebookModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePricebookModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePricebookModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricebookModel(val *PricebookModel) *NullablePricebookModel {
	return &NullablePricebookModel{value: val, isSet: true}
}

func (v NullablePricebookModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricebookModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


