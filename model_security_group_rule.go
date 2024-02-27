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

// checks if the SecurityGroupRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupRule{}

// SecurityGroupRule struct for SecurityGroupRule
type SecurityGroupRule struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	SecurityRule *SecurityGroupRuleFields `json:"security_rule,omitempty"`
}

// NewSecurityGroupRule instantiates a new SecurityGroupRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupRule() *SecurityGroupRule {
	this := SecurityGroupRule{}
	return &this
}

// NewSecurityGroupRuleWithDefaults instantiates a new SecurityGroupRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupRuleWithDefaults() *SecurityGroupRule {
	this := SecurityGroupRule{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *SecurityGroupRule) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SecurityGroupRule) SetMessage(v string) {
	o.Message = &v
}

// GetSecurityRule returns the SecurityRule field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetSecurityRule() SecurityGroupRuleFields {
	if o == nil || IsNil(o.SecurityRule) {
		var ret SecurityGroupRuleFields
		return ret
	}
	return *o.SecurityRule
}

// GetSecurityRuleOk returns a tuple with the SecurityRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetSecurityRuleOk() (*SecurityGroupRuleFields, bool) {
	if o == nil || IsNil(o.SecurityRule) {
		return nil, false
	}
	return o.SecurityRule, true
}

// HasSecurityRule returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasSecurityRule() bool {
	if o != nil && !IsNil(o.SecurityRule) {
		return true
	}

	return false
}

// SetSecurityRule gets a reference to the given SecurityGroupRuleFields and assigns it to the SecurityRule field.
func (o *SecurityGroupRule) SetSecurityRule(v SecurityGroupRuleFields) {
	o.SecurityRule = &v
}

func (o SecurityGroupRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityGroupRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.SecurityRule) {
		toSerialize["security_rule"] = o.SecurityRule
	}
	return toSerialize, nil
}

type NullableSecurityGroupRule struct {
	value *SecurityGroupRule
	isSet bool
}

func (v NullableSecurityGroupRule) Get() *SecurityGroupRule {
	return v.value
}

func (v *NullableSecurityGroupRule) Set(val *SecurityGroupRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupRule(val *SecurityGroupRule) *NullableSecurityGroupRule {
	return &NullableSecurityGroupRule{value: val, isSet: true}
}

func (v NullableSecurityGroupRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

