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

// checks if the CreateContarctFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContarctFields{}

// CreateContarctFields struct for CreateContarctFields
type CreateContarctFields struct {
	Id *int32 `json:"id,omitempty"`
	OrgId *int32 `json:"org_id,omitempty"`
	Description *string `json:"description,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate *time.Time `json:"end_date,omitempty"`
	ExpirationPolicy *int32 `json:"expiration_policy,omitempty"`
	DiscountPlans []DiscountPlanFields `json:"discount_plans,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewCreateContarctFields instantiates a new CreateContarctFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContarctFields() *CreateContarctFields {
	this := CreateContarctFields{}
	return &this
}

// NewCreateContarctFieldsWithDefaults instantiates a new CreateContarctFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContarctFieldsWithDefaults() *CreateContarctFields {
	this := CreateContarctFields{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateContarctFields) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContarctFields) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateContarctFields) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CreateContarctFields) SetId(v int32) {
	o.Id = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *CreateContarctFields) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContarctFields) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *CreateContarctFields) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *CreateContarctFields) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateContarctFields) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContarctFields) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateContarctFields) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateContarctFields) SetDescription(v string) {
	o.Description = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CreateContarctFields) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContarctFields) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CreateContarctFields) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *CreateContarctFields) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CreateContarctFields) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContarctFields) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CreateContarctFields) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *CreateContarctFields) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetExpirationPolicy returns the ExpirationPolicy field value if set, zero value otherwise.
func (o *CreateContarctFields) GetExpirationPolicy() int32 {
	if o == nil || IsNil(o.ExpirationPolicy) {
		var ret int32
		return ret
	}
	return *o.ExpirationPolicy
}

// GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContarctFields) GetExpirationPolicyOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpirationPolicy) {
		return nil, false
	}
	return o.ExpirationPolicy, true
}

// HasExpirationPolicy returns a boolean if a field has been set.
func (o *CreateContarctFields) HasExpirationPolicy() bool {
	if o != nil && !IsNil(o.ExpirationPolicy) {
		return true
	}

	return false
}

// SetExpirationPolicy gets a reference to the given int32 and assigns it to the ExpirationPolicy field.
func (o *CreateContarctFields) SetExpirationPolicy(v int32) {
	o.ExpirationPolicy = &v
}

// GetDiscountPlans returns the DiscountPlans field value if set, zero value otherwise.
func (o *CreateContarctFields) GetDiscountPlans() []DiscountPlanFields {
	if o == nil || IsNil(o.DiscountPlans) {
		var ret []DiscountPlanFields
		return ret
	}
	return o.DiscountPlans
}

// GetDiscountPlansOk returns a tuple with the DiscountPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContarctFields) GetDiscountPlansOk() ([]DiscountPlanFields, bool) {
	if o == nil || IsNil(o.DiscountPlans) {
		return nil, false
	}
	return o.DiscountPlans, true
}

// HasDiscountPlans returns a boolean if a field has been set.
func (o *CreateContarctFields) HasDiscountPlans() bool {
	if o != nil && !IsNil(o.DiscountPlans) {
		return true
	}

	return false
}

// SetDiscountPlans gets a reference to the given []DiscountPlanFields and assigns it to the DiscountPlans field.
func (o *CreateContarctFields) SetDiscountPlans(v []DiscountPlanFields) {
	o.DiscountPlans = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateContarctFields) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContarctFields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateContarctFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CreateContarctFields) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o CreateContarctFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContarctFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.ExpirationPolicy) {
		toSerialize["expiration_policy"] = o.ExpirationPolicy
	}
	if !IsNil(o.DiscountPlans) {
		toSerialize["discount_plans"] = o.DiscountPlans
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableCreateContarctFields struct {
	value *CreateContarctFields
	isSet bool
}

func (v NullableCreateContarctFields) Get() *CreateContarctFields {
	return v.value
}

func (v *NullableCreateContarctFields) Set(val *CreateContarctFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContarctFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContarctFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContarctFields(val *CreateContarctFields) *NullableCreateContarctFields {
	return &NullableCreateContarctFields{value: val, isSet: true}
}

func (v NullableCreateContarctFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContarctFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


