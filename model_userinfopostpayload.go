/*
Infrahub-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Userinfopostpayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Userinfopostpayload{}

// Userinfopostpayload struct for Userinfopostpayload
type Userinfopostpayload struct {
	Business bool `json:"business"`
	Name *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	VatNumber *string `json:"vat_number,omitempty"`
	Phone *string `json:"phone,omitempty"`
	BillingAddress1 *string `json:"billing_address1,omitempty"`
	BillingAddress2 *string `json:"billing_address2,omitempty"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
	State *string `json:"state,omitempty"`
}

type _Userinfopostpayload Userinfopostpayload

// NewUserinfopostpayload instantiates a new Userinfopostpayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserinfopostpayload(business bool, zipCode string, country string) *Userinfopostpayload {
	this := Userinfopostpayload{}
	this.Business = business
	this.ZipCode = zipCode
	this.Country = country
	return &this
}

// NewUserinfopostpayloadWithDefaults instantiates a new Userinfopostpayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserinfopostpayloadWithDefaults() *Userinfopostpayload {
	this := Userinfopostpayload{}
	return &this
}

// GetBusiness returns the Business field value
func (o *Userinfopostpayload) GetBusiness() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Business
}

// GetBusinessOk returns a tuple with the Business field value
// and a boolean to check if the value has been set.
func (o *Userinfopostpayload) GetBusinessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Business, true
}

// SetBusiness sets field value
func (o *Userinfopostpayload) SetBusiness(v bool) {
	o.Business = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Userinfopostpayload) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userinfopostpayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Userinfopostpayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Userinfopostpayload) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Userinfopostpayload) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userinfopostpayload) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Userinfopostpayload) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Userinfopostpayload) SetEmail(v string) {
	o.Email = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *Userinfopostpayload) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userinfopostpayload) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *Userinfopostpayload) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *Userinfopostpayload) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *Userinfopostpayload) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userinfopostpayload) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *Userinfopostpayload) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *Userinfopostpayload) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Userinfopostpayload) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userinfopostpayload) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Userinfopostpayload) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *Userinfopostpayload) SetPhone(v string) {
	o.Phone = &v
}

// GetBillingAddress1 returns the BillingAddress1 field value if set, zero value otherwise.
func (o *Userinfopostpayload) GetBillingAddress1() string {
	if o == nil || IsNil(o.BillingAddress1) {
		var ret string
		return ret
	}
	return *o.BillingAddress1
}

// GetBillingAddress1Ok returns a tuple with the BillingAddress1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userinfopostpayload) GetBillingAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.BillingAddress1) {
		return nil, false
	}
	return o.BillingAddress1, true
}

// HasBillingAddress1 returns a boolean if a field has been set.
func (o *Userinfopostpayload) HasBillingAddress1() bool {
	if o != nil && !IsNil(o.BillingAddress1) {
		return true
	}

	return false
}

// SetBillingAddress1 gets a reference to the given string and assigns it to the BillingAddress1 field.
func (o *Userinfopostpayload) SetBillingAddress1(v string) {
	o.BillingAddress1 = &v
}

// GetBillingAddress2 returns the BillingAddress2 field value if set, zero value otherwise.
func (o *Userinfopostpayload) GetBillingAddress2() string {
	if o == nil || IsNil(o.BillingAddress2) {
		var ret string
		return ret
	}
	return *o.BillingAddress2
}

// GetBillingAddress2Ok returns a tuple with the BillingAddress2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userinfopostpayload) GetBillingAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.BillingAddress2) {
		return nil, false
	}
	return o.BillingAddress2, true
}

// HasBillingAddress2 returns a boolean if a field has been set.
func (o *Userinfopostpayload) HasBillingAddress2() bool {
	if o != nil && !IsNil(o.BillingAddress2) {
		return true
	}

	return false
}

// SetBillingAddress2 gets a reference to the given string and assigns it to the BillingAddress2 field.
func (o *Userinfopostpayload) SetBillingAddress2(v string) {
	o.BillingAddress2 = &v
}

// GetZipCode returns the ZipCode field value
func (o *Userinfopostpayload) GetZipCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value
// and a boolean to check if the value has been set.
func (o *Userinfopostpayload) GetZipCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZipCode, true
}

// SetZipCode sets field value
func (o *Userinfopostpayload) SetZipCode(v string) {
	o.ZipCode = v
}

// GetCountry returns the Country field value
func (o *Userinfopostpayload) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *Userinfopostpayload) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *Userinfopostpayload) SetCountry(v string) {
	o.Country = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Userinfopostpayload) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userinfopostpayload) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Userinfopostpayload) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Userinfopostpayload) SetState(v string) {
	o.State = &v
}

func (o Userinfopostpayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Userinfopostpayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business"] = o.Business
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.CompanyName) {
		toSerialize["company_name"] = o.CompanyName
	}
	if !IsNil(o.VatNumber) {
		toSerialize["vat_number"] = o.VatNumber
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.BillingAddress1) {
		toSerialize["billing_address1"] = o.BillingAddress1
	}
	if !IsNil(o.BillingAddress2) {
		toSerialize["billing_address2"] = o.BillingAddress2
	}
	toSerialize["zip_code"] = o.ZipCode
	toSerialize["country"] = o.Country
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

func (o *Userinfopostpayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business",
		"zip_code",
		"country",
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

	varUserinfopostpayload := _Userinfopostpayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserinfopostpayload)

	if err != nil {
		return err
	}

	*o = Userinfopostpayload(varUserinfopostpayload)

	return err
}

type NullableUserinfopostpayload struct {
	value *Userinfopostpayload
	isSet bool
}

func (v NullableUserinfopostpayload) Get() *Userinfopostpayload {
	return v.value
}

func (v *NullableUserinfopostpayload) Set(val *Userinfopostpayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUserinfopostpayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUserinfopostpayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserinfopostpayload(val *Userinfopostpayload) *NullableUserinfopostpayload {
	return &NullableUserinfopostpayload{value: val, isSet: true}
}

func (v NullableUserinfopostpayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserinfopostpayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


