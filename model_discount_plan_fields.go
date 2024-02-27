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

// checks if the DiscountPlanFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscountPlanFields{}

// DiscountPlanFields struct for DiscountPlanFields
type DiscountPlanFields struct {
	Id *int32 `json:"id,omitempty"`
	ResourceId *int32 `json:"resource_id,omitempty"`
	ResourceName *string `json:"resource_name,omitempty"`
	ResourceCount *int32 `json:"resource_count,omitempty"`
	DiscountType *string `json:"discount_type,omitempty"`
	DiscountCode *string `json:"discount_code,omitempty"`
	DiscountPercent *float32 `json:"discount_percent,omitempty"`
	DiscountAmount *float32 `json:"discount_amount,omitempty"`
	DiscountStatus *string `json:"discount_status,omitempty"`
}

// NewDiscountPlanFields instantiates a new DiscountPlanFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscountPlanFields() *DiscountPlanFields {
	this := DiscountPlanFields{}
	return &this
}

// NewDiscountPlanFieldsWithDefaults instantiates a new DiscountPlanFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountPlanFieldsWithDefaults() *DiscountPlanFields {
	this := DiscountPlanFields{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiscountPlanFields) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountPlanFields) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiscountPlanFields) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DiscountPlanFields) SetId(v int32) {
	o.Id = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *DiscountPlanFields) GetResourceId() int32 {
	if o == nil || IsNil(o.ResourceId) {
		var ret int32
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountPlanFields) GetResourceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *DiscountPlanFields) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given int32 and assigns it to the ResourceId field.
func (o *DiscountPlanFields) SetResourceId(v int32) {
	o.ResourceId = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *DiscountPlanFields) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountPlanFields) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *DiscountPlanFields) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *DiscountPlanFields) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourceCount returns the ResourceCount field value if set, zero value otherwise.
func (o *DiscountPlanFields) GetResourceCount() int32 {
	if o == nil || IsNil(o.ResourceCount) {
		var ret int32
		return ret
	}
	return *o.ResourceCount
}

// GetResourceCountOk returns a tuple with the ResourceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountPlanFields) GetResourceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ResourceCount) {
		return nil, false
	}
	return o.ResourceCount, true
}

// HasResourceCount returns a boolean if a field has been set.
func (o *DiscountPlanFields) HasResourceCount() bool {
	if o != nil && !IsNil(o.ResourceCount) {
		return true
	}

	return false
}

// SetResourceCount gets a reference to the given int32 and assigns it to the ResourceCount field.
func (o *DiscountPlanFields) SetResourceCount(v int32) {
	o.ResourceCount = &v
}

// GetDiscountType returns the DiscountType field value if set, zero value otherwise.
func (o *DiscountPlanFields) GetDiscountType() string {
	if o == nil || IsNil(o.DiscountType) {
		var ret string
		return ret
	}
	return *o.DiscountType
}

// GetDiscountTypeOk returns a tuple with the DiscountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountPlanFields) GetDiscountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountType) {
		return nil, false
	}
	return o.DiscountType, true
}

// HasDiscountType returns a boolean if a field has been set.
func (o *DiscountPlanFields) HasDiscountType() bool {
	if o != nil && !IsNil(o.DiscountType) {
		return true
	}

	return false
}

// SetDiscountType gets a reference to the given string and assigns it to the DiscountType field.
func (o *DiscountPlanFields) SetDiscountType(v string) {
	o.DiscountType = &v
}

// GetDiscountCode returns the DiscountCode field value if set, zero value otherwise.
func (o *DiscountPlanFields) GetDiscountCode() string {
	if o == nil || IsNil(o.DiscountCode) {
		var ret string
		return ret
	}
	return *o.DiscountCode
}

// GetDiscountCodeOk returns a tuple with the DiscountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountPlanFields) GetDiscountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountCode) {
		return nil, false
	}
	return o.DiscountCode, true
}

// HasDiscountCode returns a boolean if a field has been set.
func (o *DiscountPlanFields) HasDiscountCode() bool {
	if o != nil && !IsNil(o.DiscountCode) {
		return true
	}

	return false
}

// SetDiscountCode gets a reference to the given string and assigns it to the DiscountCode field.
func (o *DiscountPlanFields) SetDiscountCode(v string) {
	o.DiscountCode = &v
}

// GetDiscountPercent returns the DiscountPercent field value if set, zero value otherwise.
func (o *DiscountPlanFields) GetDiscountPercent() float32 {
	if o == nil || IsNil(o.DiscountPercent) {
		var ret float32
		return ret
	}
	return *o.DiscountPercent
}

// GetDiscountPercentOk returns a tuple with the DiscountPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountPlanFields) GetDiscountPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountPercent) {
		return nil, false
	}
	return o.DiscountPercent, true
}

// HasDiscountPercent returns a boolean if a field has been set.
func (o *DiscountPlanFields) HasDiscountPercent() bool {
	if o != nil && !IsNil(o.DiscountPercent) {
		return true
	}

	return false
}

// SetDiscountPercent gets a reference to the given float32 and assigns it to the DiscountPercent field.
func (o *DiscountPlanFields) SetDiscountPercent(v float32) {
	o.DiscountPercent = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *DiscountPlanFields) GetDiscountAmount() float32 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret float32
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountPlanFields) GetDiscountAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *DiscountPlanFields) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given float32 and assigns it to the DiscountAmount field.
func (o *DiscountPlanFields) SetDiscountAmount(v float32) {
	o.DiscountAmount = &v
}

// GetDiscountStatus returns the DiscountStatus field value if set, zero value otherwise.
func (o *DiscountPlanFields) GetDiscountStatus() string {
	if o == nil || IsNil(o.DiscountStatus) {
		var ret string
		return ret
	}
	return *o.DiscountStatus
}

// GetDiscountStatusOk returns a tuple with the DiscountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountPlanFields) GetDiscountStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountStatus) {
		return nil, false
	}
	return o.DiscountStatus, true
}

// HasDiscountStatus returns a boolean if a field has been set.
func (o *DiscountPlanFields) HasDiscountStatus() bool {
	if o != nil && !IsNil(o.DiscountStatus) {
		return true
	}

	return false
}

// SetDiscountStatus gets a reference to the given string and assigns it to the DiscountStatus field.
func (o *DiscountPlanFields) SetDiscountStatus(v string) {
	o.DiscountStatus = &v
}

func (o DiscountPlanFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscountPlanFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ResourceId) {
		toSerialize["resource_id"] = o.ResourceId
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resource_name"] = o.ResourceName
	}
	if !IsNil(o.ResourceCount) {
		toSerialize["resource_count"] = o.ResourceCount
	}
	if !IsNil(o.DiscountType) {
		toSerialize["discount_type"] = o.DiscountType
	}
	if !IsNil(o.DiscountCode) {
		toSerialize["discount_code"] = o.DiscountCode
	}
	if !IsNil(o.DiscountPercent) {
		toSerialize["discount_percent"] = o.DiscountPercent
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discount_amount"] = o.DiscountAmount
	}
	if !IsNil(o.DiscountStatus) {
		toSerialize["discount_status"] = o.DiscountStatus
	}
	return toSerialize, nil
}

type NullableDiscountPlanFields struct {
	value *DiscountPlanFields
	isSet bool
}

func (v NullableDiscountPlanFields) Get() *DiscountPlanFields {
	return v.value
}

func (v *NullableDiscountPlanFields) Set(val *DiscountPlanFields) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountPlanFields) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountPlanFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountPlanFields(val *DiscountPlanFields) *NullableDiscountPlanFields {
	return &NullableDiscountPlanFields{value: val, isSet: true}
}

func (v NullableDiscountPlanFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscountPlanFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


