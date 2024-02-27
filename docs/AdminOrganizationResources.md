# AdminOrganizationResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Environments** | Pointer to [**[]AdminEnvrionmentResources**](AdminEnvrionmentResources.md) |  | [optional] 

## Methods

### NewAdminOrganizationResources

`func NewAdminOrganizationResources() *AdminOrganizationResources`

NewAdminOrganizationResources instantiates a new AdminOrganizationResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminOrganizationResourcesWithDefaults

`func NewAdminOrganizationResourcesWithDefaults() *AdminOrganizationResources`

NewAdminOrganizationResourcesWithDefaults instantiates a new AdminOrganizationResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *AdminOrganizationResources) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AdminOrganizationResources) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AdminOrganizationResources) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AdminOrganizationResources) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetEnvironments

`func (o *AdminOrganizationResources) GetEnvironments() []AdminEnvrionmentResources`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *AdminOrganizationResources) GetEnvironmentsOk() (*[]AdminEnvrionmentResources, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *AdminOrganizationResources) SetEnvironments(v []AdminEnvrionmentResources)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *AdminOrganizationResources) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


