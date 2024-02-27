# AdminOrganizationSummaryFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAdminOrganizationSummaryFields

`func NewAdminOrganizationSummaryFields(id int32, name string, ) *AdminOrganizationSummaryFields`

NewAdminOrganizationSummaryFields instantiates a new AdminOrganizationSummaryFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminOrganizationSummaryFieldsWithDefaults

`func NewAdminOrganizationSummaryFieldsWithDefaults() *AdminOrganizationSummaryFields`

NewAdminOrganizationSummaryFieldsWithDefaults instantiates a new AdminOrganizationSummaryFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminOrganizationSummaryFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminOrganizationSummaryFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminOrganizationSummaryFields) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AdminOrganizationSummaryFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminOrganizationSummaryFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminOrganizationSummaryFields) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *AdminOrganizationSummaryFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminOrganizationSummaryFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminOrganizationSummaryFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminOrganizationSummaryFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


