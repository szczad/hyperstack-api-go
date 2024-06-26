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
	"bytes"
	"fmt"
)

// checks if the AdminOrganizationSummaryFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminOrganizationSummaryFields{}

// AdminOrganizationSummaryFields struct for AdminOrganizationSummaryFields
type AdminOrganizationSummaryFields struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type _AdminOrganizationSummaryFields AdminOrganizationSummaryFields

// NewAdminOrganizationSummaryFields instantiates a new AdminOrganizationSummaryFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminOrganizationSummaryFields(id int32, name string) *AdminOrganizationSummaryFields {
	this := AdminOrganizationSummaryFields{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAdminOrganizationSummaryFieldsWithDefaults instantiates a new AdminOrganizationSummaryFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminOrganizationSummaryFieldsWithDefaults() *AdminOrganizationSummaryFields {
	this := AdminOrganizationSummaryFields{}
	return &this
}

// GetId returns the Id field value
func (o *AdminOrganizationSummaryFields) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AdminOrganizationSummaryFields) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdminOrganizationSummaryFields) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AdminOrganizationSummaryFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdminOrganizationSummaryFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdminOrganizationSummaryFields) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AdminOrganizationSummaryFields) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminOrganizationSummaryFields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AdminOrganizationSummaryFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AdminOrganizationSummaryFields) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o AdminOrganizationSummaryFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminOrganizationSummaryFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

func (o *AdminOrganizationSummaryFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAdminOrganizationSummaryFields := _AdminOrganizationSummaryFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAdminOrganizationSummaryFields)

	if err != nil {
		return err
	}

	*o = AdminOrganizationSummaryFields(varAdminOrganizationSummaryFields)

	return err
}

type NullableAdminOrganizationSummaryFields struct {
	value *AdminOrganizationSummaryFields
	isSet bool
}

func (v NullableAdminOrganizationSummaryFields) Get() *AdminOrganizationSummaryFields {
	return v.value
}

func (v *NullableAdminOrganizationSummaryFields) Set(val *AdminOrganizationSummaryFields) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminOrganizationSummaryFields) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminOrganizationSummaryFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminOrganizationSummaryFields(val *AdminOrganizationSummaryFields) *NullableAdminOrganizationSummaryFields {
	return &NullableAdminOrganizationSummaryFields{value: val, isSet: true}
}

func (v NullableAdminOrganizationSummaryFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminOrganizationSummaryFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


