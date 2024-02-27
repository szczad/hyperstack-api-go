# CreateContarctFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationPolicy** | Pointer to **int32** |  | [optional] 
**DiscountPlans** | Pointer to [**[]DiscountPlanFields**](DiscountPlanFields.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateContarctFields

`func NewCreateContarctFields() *CreateContarctFields`

NewCreateContarctFields instantiates a new CreateContarctFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContarctFieldsWithDefaults

`func NewCreateContarctFieldsWithDefaults() *CreateContarctFields`

NewCreateContarctFieldsWithDefaults instantiates a new CreateContarctFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateContarctFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateContarctFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateContarctFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateContarctFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *CreateContarctFields) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CreateContarctFields) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CreateContarctFields) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *CreateContarctFields) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateContarctFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateContarctFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateContarctFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateContarctFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateContarctFields) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateContarctFields) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateContarctFields) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateContarctFields) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CreateContarctFields) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateContarctFields) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateContarctFields) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateContarctFields) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExpirationPolicy

`func (o *CreateContarctFields) GetExpirationPolicy() int32`

GetExpirationPolicy returns the ExpirationPolicy field if non-nil, zero value otherwise.

### GetExpirationPolicyOk

`func (o *CreateContarctFields) GetExpirationPolicyOk() (*int32, bool)`

GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationPolicy

`func (o *CreateContarctFields) SetExpirationPolicy(v int32)`

SetExpirationPolicy sets ExpirationPolicy field to given value.

### HasExpirationPolicy

`func (o *CreateContarctFields) HasExpirationPolicy() bool`

HasExpirationPolicy returns a boolean if a field has been set.

### GetDiscountPlans

`func (o *CreateContarctFields) GetDiscountPlans() []DiscountPlanFields`

GetDiscountPlans returns the DiscountPlans field if non-nil, zero value otherwise.

### GetDiscountPlansOk

`func (o *CreateContarctFields) GetDiscountPlansOk() (*[]DiscountPlanFields, bool)`

GetDiscountPlansOk returns a tuple with the DiscountPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPlans

`func (o *CreateContarctFields) SetDiscountPlans(v []DiscountPlanFields)`

SetDiscountPlans sets DiscountPlans field to given value.

### HasDiscountPlans

`func (o *CreateContarctFields) HasDiscountPlans() bool`

HasDiscountPlans returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateContarctFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateContarctFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateContarctFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateContarctFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


