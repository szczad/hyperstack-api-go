# AdminContractFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationPolicy** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**GpuName** | Pointer to **string** |  | [optional] 
**Percent** | Pointer to **float32** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**DiscountType** | Pointer to **string** |  | [optional] 
**ResourceCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAdminContractFields

`func NewAdminContractFields() *AdminContractFields`

NewAdminContractFields instantiates a new AdminContractFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminContractFieldsWithDefaults

`func NewAdminContractFieldsWithDefaults() *AdminContractFields`

NewAdminContractFieldsWithDefaults instantiates a new AdminContractFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminContractFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminContractFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminContractFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdminContractFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgName

`func (o *AdminContractFields) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *AdminContractFields) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *AdminContractFields) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *AdminContractFields) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetDescription

`func (o *AdminContractFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminContractFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminContractFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminContractFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartDate

`func (o *AdminContractFields) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AdminContractFields) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AdminContractFields) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AdminContractFields) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *AdminContractFields) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AdminContractFields) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AdminContractFields) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AdminContractFields) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExpirationPolicy

`func (o *AdminContractFields) GetExpirationPolicy() int32`

GetExpirationPolicy returns the ExpirationPolicy field if non-nil, zero value otherwise.

### GetExpirationPolicyOk

`func (o *AdminContractFields) GetExpirationPolicyOk() (*int32, bool)`

GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationPolicy

`func (o *AdminContractFields) SetExpirationPolicy(v int32)`

SetExpirationPolicy sets ExpirationPolicy field to given value.

### HasExpirationPolicy

`func (o *AdminContractFields) HasExpirationPolicy() bool`

HasExpirationPolicy returns a boolean if a field has been set.

### GetStatus

`func (o *AdminContractFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminContractFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminContractFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminContractFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGpuName

`func (o *AdminContractFields) GetGpuName() string`

GetGpuName returns the GpuName field if non-nil, zero value otherwise.

### GetGpuNameOk

`func (o *AdminContractFields) GetGpuNameOk() (*string, bool)`

GetGpuNameOk returns a tuple with the GpuName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuName

`func (o *AdminContractFields) SetGpuName(v string)`

SetGpuName sets GpuName field to given value.

### HasGpuName

`func (o *AdminContractFields) HasGpuName() bool`

HasGpuName returns a boolean if a field has been set.

### GetPercent

`func (o *AdminContractFields) GetPercent() float32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *AdminContractFields) GetPercentOk() (*float32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *AdminContractFields) SetPercent(v float32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *AdminContractFields) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetAmount

`func (o *AdminContractFields) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AdminContractFields) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AdminContractFields) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AdminContractFields) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDiscountType

`func (o *AdminContractFields) GetDiscountType() string`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *AdminContractFields) GetDiscountTypeOk() (*string, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *AdminContractFields) SetDiscountType(v string)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *AdminContractFields) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetResourceCount

`func (o *AdminContractFields) GetResourceCount() int32`

GetResourceCount returns the ResourceCount field if non-nil, zero value otherwise.

### GetResourceCountOk

`func (o *AdminContractFields) GetResourceCountOk() (*int32, bool)`

GetResourceCountOk returns a tuple with the ResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCount

`func (o *AdminContractFields) SetResourceCount(v int32)`

SetResourceCount sets ResourceCount field to given value.

### HasResourceCount

`func (o *AdminContractFields) HasResourceCount() bool`

HasResourceCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminContractFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminContractFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminContractFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminContractFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


