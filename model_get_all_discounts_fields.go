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

// checks if the GetAllDiscountsFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllDiscountsFields{}

// GetAllDiscountsFields struct for GetAllDiscountsFields
type GetAllDiscountsFields struct {
	OrgId *int32 `json:"org_id,omitempty"`
	OrgName *string `json:"org_name,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate *time.Time `json:"end_date,omitempty"`
	DiscountStatus *string `json:"discount_status,omitempty"`
	DiscountResources []DiscountResourceFields `json:"discount_resources,omitempty"`
}

// NewGetAllDiscountsFields instantiates a new GetAllDiscountsFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllDiscountsFields() *GetAllDiscountsFields {
	this := GetAllDiscountsFields{}
	return &this
}

// NewGetAllDiscountsFieldsWithDefaults instantiates a new GetAllDiscountsFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllDiscountsFieldsWithDefaults() *GetAllDiscountsFields {
	this := GetAllDiscountsFields{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *GetAllDiscountsFields) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDiscountsFields) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *GetAllDiscountsFields) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *GetAllDiscountsFields) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *GetAllDiscountsFields) GetOrgName() string {
	if o == nil || IsNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDiscountsFields) GetOrgNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrgName) {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *GetAllDiscountsFields) HasOrgName() bool {
	if o != nil && !IsNil(o.OrgName) {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *GetAllDiscountsFields) SetOrgName(v string) {
	o.OrgName = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetAllDiscountsFields) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDiscountsFields) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetAllDiscountsFields) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *GetAllDiscountsFields) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetAllDiscountsFields) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDiscountsFields) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetAllDiscountsFields) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *GetAllDiscountsFields) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetDiscountStatus returns the DiscountStatus field value if set, zero value otherwise.
func (o *GetAllDiscountsFields) GetDiscountStatus() string {
	if o == nil || IsNil(o.DiscountStatus) {
		var ret string
		return ret
	}
	return *o.DiscountStatus
}

// GetDiscountStatusOk returns a tuple with the DiscountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDiscountsFields) GetDiscountStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountStatus) {
		return nil, false
	}
	return o.DiscountStatus, true
}

// HasDiscountStatus returns a boolean if a field has been set.
func (o *GetAllDiscountsFields) HasDiscountStatus() bool {
	if o != nil && !IsNil(o.DiscountStatus) {
		return true
	}

	return false
}

// SetDiscountStatus gets a reference to the given string and assigns it to the DiscountStatus field.
func (o *GetAllDiscountsFields) SetDiscountStatus(v string) {
	o.DiscountStatus = &v
}

// GetDiscountResources returns the DiscountResources field value if set, zero value otherwise.
func (o *GetAllDiscountsFields) GetDiscountResources() []DiscountResourceFields {
	if o == nil || IsNil(o.DiscountResources) {
		var ret []DiscountResourceFields
		return ret
	}
	return o.DiscountResources
}

// GetDiscountResourcesOk returns a tuple with the DiscountResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDiscountsFields) GetDiscountResourcesOk() ([]DiscountResourceFields, bool) {
	if o == nil || IsNil(o.DiscountResources) {
		return nil, false
	}
	return o.DiscountResources, true
}

// HasDiscountResources returns a boolean if a field has been set.
func (o *GetAllDiscountsFields) HasDiscountResources() bool {
	if o != nil && !IsNil(o.DiscountResources) {
		return true
	}

	return false
}

// SetDiscountResources gets a reference to the given []DiscountResourceFields and assigns it to the DiscountResources field.
func (o *GetAllDiscountsFields) SetDiscountResources(v []DiscountResourceFields) {
	o.DiscountResources = v
}

func (o GetAllDiscountsFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllDiscountsFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.OrgName) {
		toSerialize["org_name"] = o.OrgName
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.DiscountStatus) {
		toSerialize["discount_status"] = o.DiscountStatus
	}
	if !IsNil(o.DiscountResources) {
		toSerialize["discount_resources"] = o.DiscountResources
	}
	return toSerialize, nil
}

type NullableGetAllDiscountsFields struct {
	value *GetAllDiscountsFields
	isSet bool
}

func (v NullableGetAllDiscountsFields) Get() *GetAllDiscountsFields {
	return v.value
}

func (v *NullableGetAllDiscountsFields) Set(val *GetAllDiscountsFields) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllDiscountsFields) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllDiscountsFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllDiscountsFields(val *GetAllDiscountsFields) *NullableGetAllDiscountsFields {
	return &NullableGetAllDiscountsFields{value: val, isSet: true}
}

func (v NullableGetAllDiscountsFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllDiscountsFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


