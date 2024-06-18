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

// checks if the Creditrequests type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Creditrequests{}

// Creditrequests struct for Creditrequests
type Creditrequests struct {
	AdminUserId *int32 `json:"admin_user_id,omitempty"`
	Status *string `json:"status,omitempty"`
	Amount *float32 `json:"amount,omitempty"`
	Reason *string `json:"reason,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewCreditrequests instantiates a new Creditrequests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditrequests() *Creditrequests {
	this := Creditrequests{}
	return &this
}

// NewCreditrequestsWithDefaults instantiates a new Creditrequests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditrequestsWithDefaults() *Creditrequests {
	this := Creditrequests{}
	return &this
}

// GetAdminUserId returns the AdminUserId field value if set, zero value otherwise.
func (o *Creditrequests) GetAdminUserId() int32 {
	if o == nil || IsNil(o.AdminUserId) {
		var ret int32
		return ret
	}
	return *o.AdminUserId
}

// GetAdminUserIdOk returns a tuple with the AdminUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creditrequests) GetAdminUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AdminUserId) {
		return nil, false
	}
	return o.AdminUserId, true
}

// HasAdminUserId returns a boolean if a field has been set.
func (o *Creditrequests) HasAdminUserId() bool {
	if o != nil && !IsNil(o.AdminUserId) {
		return true
	}

	return false
}

// SetAdminUserId gets a reference to the given int32 and assigns it to the AdminUserId field.
func (o *Creditrequests) SetAdminUserId(v int32) {
	o.AdminUserId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Creditrequests) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creditrequests) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Creditrequests) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Creditrequests) SetStatus(v string) {
	o.Status = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Creditrequests) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creditrequests) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Creditrequests) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *Creditrequests) SetAmount(v float32) {
	o.Amount = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Creditrequests) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creditrequests) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Creditrequests) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Creditrequests) SetReason(v string) {
	o.Reason = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Creditrequests) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Creditrequests) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Creditrequests) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Creditrequests) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Creditrequests) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Creditrequests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminUserId) {
		toSerialize["admin_user_id"] = o.AdminUserId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableCreditrequests struct {
	value *Creditrequests
	isSet bool
}

func (v NullableCreditrequests) Get() *Creditrequests {
	return v.value
}

func (v *NullableCreditrequests) Set(val *Creditrequests) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditrequests) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditrequests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditrequests(val *Creditrequests) *NullableCreditrequests {
	return &NullableCreditrequests{value: val, isSet: true}
}

func (v NullableCreditrequests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditrequests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


