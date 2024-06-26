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

// checks if the GetContractEventsResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContractEventsResponseModel{}

// GetContractEventsResponseModel struct for GetContractEventsResponseModel
type GetContractEventsResponseModel struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	ContractEvents []AdminContractEventFields `json:"contract_events,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// NewGetContractEventsResponseModel instantiates a new GetContractEventsResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContractEventsResponseModel() *GetContractEventsResponseModel {
	this := GetContractEventsResponseModel{}
	return &this
}

// NewGetContractEventsResponseModelWithDefaults instantiates a new GetContractEventsResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContractEventsResponseModelWithDefaults() *GetContractEventsResponseModel {
	this := GetContractEventsResponseModel{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetContractEventsResponseModel) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContractEventsResponseModel) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetContractEventsResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *GetContractEventsResponseModel) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetContractEventsResponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContractEventsResponseModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetContractEventsResponseModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetContractEventsResponseModel) SetMessage(v string) {
	o.Message = &v
}

// GetContractEvents returns the ContractEvents field value if set, zero value otherwise.
func (o *GetContractEventsResponseModel) GetContractEvents() []AdminContractEventFields {
	if o == nil || IsNil(o.ContractEvents) {
		var ret []AdminContractEventFields
		return ret
	}
	return o.ContractEvents
}

// GetContractEventsOk returns a tuple with the ContractEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContractEventsResponseModel) GetContractEventsOk() ([]AdminContractEventFields, bool) {
	if o == nil || IsNil(o.ContractEvents) {
		return nil, false
	}
	return o.ContractEvents, true
}

// HasContractEvents returns a boolean if a field has been set.
func (o *GetContractEventsResponseModel) HasContractEvents() bool {
	if o != nil && !IsNil(o.ContractEvents) {
		return true
	}

	return false
}

// SetContractEvents gets a reference to the given []AdminContractEventFields and assigns it to the ContractEvents field.
func (o *GetContractEventsResponseModel) SetContractEvents(v []AdminContractEventFields) {
	o.ContractEvents = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetContractEventsResponseModel) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContractEventsResponseModel) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetContractEventsResponseModel) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetContractEventsResponseModel) SetCount(v int32) {
	o.Count = &v
}

func (o GetContractEventsResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContractEventsResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.ContractEvents) {
		toSerialize["contract_events"] = o.ContractEvents
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableGetContractEventsResponseModel struct {
	value *GetContractEventsResponseModel
	isSet bool
}

func (v NullableGetContractEventsResponseModel) Get() *GetContractEventsResponseModel {
	return v.value
}

func (v *NullableGetContractEventsResponseModel) Set(val *GetContractEventsResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContractEventsResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContractEventsResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContractEventsResponseModel(val *GetContractEventsResponseModel) *NullableGetContractEventsResponseModel {
	return &NullableGetContractEventsResponseModel{value: val, isSet: true}
}

func (v NullableGetContractEventsResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContractEventsResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


