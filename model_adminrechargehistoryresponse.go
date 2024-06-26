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

// checks if the Adminrechargehistoryresponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Adminrechargehistoryresponse{}

// Adminrechargehistoryresponse struct for Adminrechargehistoryresponse
type Adminrechargehistoryresponse struct {
	Message *string `json:"message,omitempty"`
	Status *bool `json:"status,omitempty"`
	Data []Adminrechargehistoryfields `json:"data,omitempty"`
}

// NewAdminrechargehistoryresponse instantiates a new Adminrechargehistoryresponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminrechargehistoryresponse() *Adminrechargehistoryresponse {
	this := Adminrechargehistoryresponse{}
	return &this
}

// NewAdminrechargehistoryresponseWithDefaults instantiates a new Adminrechargehistoryresponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminrechargehistoryresponseWithDefaults() *Adminrechargehistoryresponse {
	this := Adminrechargehistoryresponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Adminrechargehistoryresponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Adminrechargehistoryresponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Adminrechargehistoryresponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Adminrechargehistoryresponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Adminrechargehistoryresponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Adminrechargehistoryresponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Adminrechargehistoryresponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *Adminrechargehistoryresponse) SetStatus(v bool) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Adminrechargehistoryresponse) GetData() []Adminrechargehistoryfields {
	if o == nil || IsNil(o.Data) {
		var ret []Adminrechargehistoryfields
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Adminrechargehistoryresponse) GetDataOk() ([]Adminrechargehistoryfields, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Adminrechargehistoryresponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Adminrechargehistoryfields and assigns it to the Data field.
func (o *Adminrechargehistoryresponse) SetData(v []Adminrechargehistoryfields) {
	o.Data = v
}

func (o Adminrechargehistoryresponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Adminrechargehistoryresponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAdminrechargehistoryresponse struct {
	value *Adminrechargehistoryresponse
	isSet bool
}

func (v NullableAdminrechargehistoryresponse) Get() *Adminrechargehistoryresponse {
	return v.value
}

func (v *NullableAdminrechargehistoryresponse) Set(val *Adminrechargehistoryresponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminrechargehistoryresponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminrechargehistoryresponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminrechargehistoryresponse(val *Adminrechargehistoryresponse) *NullableAdminrechargehistoryresponse {
	return &NullableAdminrechargehistoryresponse{value: val, isSet: true}
}

func (v NullableAdminrechargehistoryresponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminrechargehistoryresponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


