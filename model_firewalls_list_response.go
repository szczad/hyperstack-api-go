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

// checks if the FirewallsListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallsListResponse{}

// FirewallsListResponse struct for FirewallsListResponse
type FirewallsListResponse struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	Firewalls []FirewallDetailFields `json:"firewalls,omitempty"`
}

// NewFirewallsListResponse instantiates a new FirewallsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallsListResponse() *FirewallsListResponse {
	this := FirewallsListResponse{}
	return &this
}

// NewFirewallsListResponseWithDefaults instantiates a new FirewallsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallsListResponseWithDefaults() *FirewallsListResponse {
	this := FirewallsListResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FirewallsListResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallsListResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FirewallsListResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *FirewallsListResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FirewallsListResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallsListResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FirewallsListResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FirewallsListResponse) SetMessage(v string) {
	o.Message = &v
}

// GetFirewalls returns the Firewalls field value if set, zero value otherwise.
func (o *FirewallsListResponse) GetFirewalls() []FirewallDetailFields {
	if o == nil || IsNil(o.Firewalls) {
		var ret []FirewallDetailFields
		return ret
	}
	return o.Firewalls
}

// GetFirewallsOk returns a tuple with the Firewalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallsListResponse) GetFirewallsOk() ([]FirewallDetailFields, bool) {
	if o == nil || IsNil(o.Firewalls) {
		return nil, false
	}
	return o.Firewalls, true
}

// HasFirewalls returns a boolean if a field has been set.
func (o *FirewallsListResponse) HasFirewalls() bool {
	if o != nil && !IsNil(o.Firewalls) {
		return true
	}

	return false
}

// SetFirewalls gets a reference to the given []FirewallDetailFields and assigns it to the Firewalls field.
func (o *FirewallsListResponse) SetFirewalls(v []FirewallDetailFields) {
	o.Firewalls = v
}

func (o FirewallsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallsListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Firewalls) {
		toSerialize["firewalls"] = o.Firewalls
	}
	return toSerialize, nil
}

type NullableFirewallsListResponse struct {
	value *FirewallsListResponse
	isSet bool
}

func (v NullableFirewallsListResponse) Get() *FirewallsListResponse {
	return v.value
}

func (v *NullableFirewallsListResponse) Set(val *FirewallsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallsListResponse(val *FirewallsListResponse) *NullableFirewallsListResponse {
	return &NullableFirewallsListResponse{value: val, isSet: true}
}

func (v NullableFirewallsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

