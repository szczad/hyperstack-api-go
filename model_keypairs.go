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

// checks if the Keypairs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Keypairs{}

// Keypairs struct for Keypairs
type Keypairs struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	Keypairs []KeypairFields `json:"Keypairs,omitempty"`
}

// NewKeypairs instantiates a new Keypairs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeypairs() *Keypairs {
	this := Keypairs{}
	return &this
}

// NewKeypairsWithDefaults instantiates a new Keypairs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeypairsWithDefaults() *Keypairs {
	this := Keypairs{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Keypairs) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keypairs) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Keypairs) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *Keypairs) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Keypairs) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keypairs) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Keypairs) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Keypairs) SetMessage(v string) {
	o.Message = &v
}

// GetKeypairs returns the Keypairs field value if set, zero value otherwise.
func (o *Keypairs) GetKeypairs() []KeypairFields {
	if o == nil || IsNil(o.Keypairs) {
		var ret []KeypairFields
		return ret
	}
	return o.Keypairs
}

// GetKeypairsOk returns a tuple with the Keypairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keypairs) GetKeypairsOk() ([]KeypairFields, bool) {
	if o == nil || IsNil(o.Keypairs) {
		return nil, false
	}
	return o.Keypairs, true
}

// HasKeypairs returns a boolean if a field has been set.
func (o *Keypairs) HasKeypairs() bool {
	if o != nil && !IsNil(o.Keypairs) {
		return true
	}

	return false
}

// SetKeypairs gets a reference to the given []KeypairFields and assigns it to the Keypairs field.
func (o *Keypairs) SetKeypairs(v []KeypairFields) {
	o.Keypairs = v
}

func (o Keypairs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Keypairs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Keypairs) {
		toSerialize["Keypairs"] = o.Keypairs
	}
	return toSerialize, nil
}

type NullableKeypairs struct {
	value *Keypairs
	isSet bool
}

func (v NullableKeypairs) Get() *Keypairs {
	return v.value
}

func (v *NullableKeypairs) Set(val *Keypairs) {
	v.value = val
	v.isSet = true
}

func (v NullableKeypairs) IsSet() bool {
	return v.isSet
}

func (v *NullableKeypairs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeypairs(val *Keypairs) *NullableKeypairs {
	return &NullableKeypairs{value: val, isSet: true}
}

func (v NullableKeypairs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeypairs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

