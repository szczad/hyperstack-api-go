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

// checks if the Adminrechargehistoryfields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Adminrechargehistoryfields{}

// Adminrechargehistoryfields struct for Adminrechargehistoryfields
type Adminrechargehistoryfields struct {
	Credit *float32 `json:"credit,omitempty"`
	Threshold *float32 `json:"threshold,omitempty"`
	PrevBalance *float32 `json:"prev_balance,omitempty"`
	CurrBalance *float32 `json:"curr_balance,omitempty"`
	Medium *string `json:"medium,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewAdminrechargehistoryfields instantiates a new Adminrechargehistoryfields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminrechargehistoryfields() *Adminrechargehistoryfields {
	this := Adminrechargehistoryfields{}
	return &this
}

// NewAdminrechargehistoryfieldsWithDefaults instantiates a new Adminrechargehistoryfields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminrechargehistoryfieldsWithDefaults() *Adminrechargehistoryfields {
	this := Adminrechargehistoryfields{}
	return &this
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *Adminrechargehistoryfields) GetCredit() float32 {
	if o == nil || IsNil(o.Credit) {
		var ret float32
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Adminrechargehistoryfields) GetCreditOk() (*float32, bool) {
	if o == nil || IsNil(o.Credit) {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *Adminrechargehistoryfields) HasCredit() bool {
	if o != nil && !IsNil(o.Credit) {
		return true
	}

	return false
}

// SetCredit gets a reference to the given float32 and assigns it to the Credit field.
func (o *Adminrechargehistoryfields) SetCredit(v float32) {
	o.Credit = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *Adminrechargehistoryfields) GetThreshold() float32 {
	if o == nil || IsNil(o.Threshold) {
		var ret float32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Adminrechargehistoryfields) GetThresholdOk() (*float32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *Adminrechargehistoryfields) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float32 and assigns it to the Threshold field.
func (o *Adminrechargehistoryfields) SetThreshold(v float32) {
	o.Threshold = &v
}

// GetPrevBalance returns the PrevBalance field value if set, zero value otherwise.
func (o *Adminrechargehistoryfields) GetPrevBalance() float32 {
	if o == nil || IsNil(o.PrevBalance) {
		var ret float32
		return ret
	}
	return *o.PrevBalance
}

// GetPrevBalanceOk returns a tuple with the PrevBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Adminrechargehistoryfields) GetPrevBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.PrevBalance) {
		return nil, false
	}
	return o.PrevBalance, true
}

// HasPrevBalance returns a boolean if a field has been set.
func (o *Adminrechargehistoryfields) HasPrevBalance() bool {
	if o != nil && !IsNil(o.PrevBalance) {
		return true
	}

	return false
}

// SetPrevBalance gets a reference to the given float32 and assigns it to the PrevBalance field.
func (o *Adminrechargehistoryfields) SetPrevBalance(v float32) {
	o.PrevBalance = &v
}

// GetCurrBalance returns the CurrBalance field value if set, zero value otherwise.
func (o *Adminrechargehistoryfields) GetCurrBalance() float32 {
	if o == nil || IsNil(o.CurrBalance) {
		var ret float32
		return ret
	}
	return *o.CurrBalance
}

// GetCurrBalanceOk returns a tuple with the CurrBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Adminrechargehistoryfields) GetCurrBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.CurrBalance) {
		return nil, false
	}
	return o.CurrBalance, true
}

// HasCurrBalance returns a boolean if a field has been set.
func (o *Adminrechargehistoryfields) HasCurrBalance() bool {
	if o != nil && !IsNil(o.CurrBalance) {
		return true
	}

	return false
}

// SetCurrBalance gets a reference to the given float32 and assigns it to the CurrBalance field.
func (o *Adminrechargehistoryfields) SetCurrBalance(v float32) {
	o.CurrBalance = &v
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *Adminrechargehistoryfields) GetMedium() string {
	if o == nil || IsNil(o.Medium) {
		var ret string
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Adminrechargehistoryfields) GetMediumOk() (*string, bool) {
	if o == nil || IsNil(o.Medium) {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *Adminrechargehistoryfields) HasMedium() bool {
	if o != nil && !IsNil(o.Medium) {
		return true
	}

	return false
}

// SetMedium gets a reference to the given string and assigns it to the Medium field.
func (o *Adminrechargehistoryfields) SetMedium(v string) {
	o.Medium = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Adminrechargehistoryfields) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Adminrechargehistoryfields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Adminrechargehistoryfields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Adminrechargehistoryfields) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o Adminrechargehistoryfields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Adminrechargehistoryfields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credit) {
		toSerialize["credit"] = o.Credit
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.PrevBalance) {
		toSerialize["prev_balance"] = o.PrevBalance
	}
	if !IsNil(o.CurrBalance) {
		toSerialize["curr_balance"] = o.CurrBalance
	}
	if !IsNil(o.Medium) {
		toSerialize["medium"] = o.Medium
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableAdminrechargehistoryfields struct {
	value *Adminrechargehistoryfields
	isSet bool
}

func (v NullableAdminrechargehistoryfields) Get() *Adminrechargehistoryfields {
	return v.value
}

func (v *NullableAdminrechargehistoryfields) Set(val *Adminrechargehistoryfields) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminrechargehistoryfields) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminrechargehistoryfields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminrechargehistoryfields(val *Adminrechargehistoryfields) *NullableAdminrechargehistoryfields {
	return &NullableAdminrechargehistoryfields{value: val, isSet: true}
}

func (v NullableAdminrechargehistoryfields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminrechargehistoryfields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

