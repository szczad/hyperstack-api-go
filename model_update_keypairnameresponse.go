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

// checks if the UpdateKeypairnameresponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateKeypairnameresponse{}

// UpdateKeypairnameresponse struct for UpdateKeypairnameresponse
type UpdateKeypairnameresponse struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	Keypair *KeypairFields `json:"keypair,omitempty"`
}

// NewUpdateKeypairnameresponse instantiates a new UpdateKeypairnameresponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateKeypairnameresponse() *UpdateKeypairnameresponse {
	this := UpdateKeypairnameresponse{}
	return &this
}

// NewUpdateKeypairnameresponseWithDefaults instantiates a new UpdateKeypairnameresponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateKeypairnameresponseWithDefaults() *UpdateKeypairnameresponse {
	this := UpdateKeypairnameresponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateKeypairnameresponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateKeypairnameresponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateKeypairnameresponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *UpdateKeypairnameresponse) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateKeypairnameresponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateKeypairnameresponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateKeypairnameresponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateKeypairnameresponse) SetMessage(v string) {
	o.Message = &v
}

// GetKeypair returns the Keypair field value if set, zero value otherwise.
func (o *UpdateKeypairnameresponse) GetKeypair() KeypairFields {
	if o == nil || IsNil(o.Keypair) {
		var ret KeypairFields
		return ret
	}
	return *o.Keypair
}

// GetKeypairOk returns a tuple with the Keypair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateKeypairnameresponse) GetKeypairOk() (*KeypairFields, bool) {
	if o == nil || IsNil(o.Keypair) {
		return nil, false
	}
	return o.Keypair, true
}

// HasKeypair returns a boolean if a field has been set.
func (o *UpdateKeypairnameresponse) HasKeypair() bool {
	if o != nil && !IsNil(o.Keypair) {
		return true
	}

	return false
}

// SetKeypair gets a reference to the given KeypairFields and assigns it to the Keypair field.
func (o *UpdateKeypairnameresponse) SetKeypair(v KeypairFields) {
	o.Keypair = &v
}

func (o UpdateKeypairnameresponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateKeypairnameresponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Keypair) {
		toSerialize["keypair"] = o.Keypair
	}
	return toSerialize, nil
}

type NullableUpdateKeypairnameresponse struct {
	value *UpdateKeypairnameresponse
	isSet bool
}

func (v NullableUpdateKeypairnameresponse) Get() *UpdateKeypairnameresponse {
	return v.value
}

func (v *NullableUpdateKeypairnameresponse) Set(val *UpdateKeypairnameresponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateKeypairnameresponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateKeypairnameresponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateKeypairnameresponse(val *UpdateKeypairnameresponse) *NullableUpdateKeypairnameresponse {
	return &NullableUpdateKeypairnameresponse{value: val, isSet: true}
}

func (v NullableUpdateKeypairnameresponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateKeypairnameresponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


