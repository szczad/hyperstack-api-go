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

// checks if the PaymentInitiateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentInitiateResponse{}

// PaymentInitiateResponse struct for PaymentInitiateResponse
type PaymentInitiateResponse struct {
	Message *string `json:"message,omitempty"`
	Status *bool `json:"status,omitempty"`
	Data *PaymentInitiateFields `json:"data,omitempty"`
}

// NewPaymentInitiateResponse instantiates a new PaymentInitiateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiateResponse() *PaymentInitiateResponse {
	this := PaymentInitiateResponse{}
	return &this
}

// NewPaymentInitiateResponseWithDefaults instantiates a new PaymentInitiateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiateResponseWithDefaults() *PaymentInitiateResponse {
	this := PaymentInitiateResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PaymentInitiateResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiateResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PaymentInitiateResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PaymentInitiateResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentInitiateResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiateResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentInitiateResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *PaymentInitiateResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PaymentInitiateResponse) GetData() PaymentInitiateFields {
	if o == nil || IsNil(o.Data) {
		var ret PaymentInitiateFields
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiateResponse) GetDataOk() (*PaymentInitiateFields, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PaymentInitiateResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PaymentInitiateFields and assigns it to the Data field.
func (o *PaymentInitiateResponse) SetData(v PaymentInitiateFields) {
	o.Data = &v
}

func (o PaymentInitiateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentInitiateResponse) ToMap() (map[string]interface{}, error) {
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

type NullablePaymentInitiateResponse struct {
	value *PaymentInitiateResponse
	isSet bool
}

func (v NullablePaymentInitiateResponse) Get() *PaymentInitiateResponse {
	return v.value
}

func (v *NullablePaymentInitiateResponse) Set(val *PaymentInitiateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiateResponse(val *PaymentInitiateResponse) *NullablePaymentInitiateResponse {
	return &NullablePaymentInitiateResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

