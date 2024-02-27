# AdminUserFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Username** | **string** |  | 
**Email** | **string** |  | 
**OrgId** | **int32** |  | 
**OrgName** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastLogin** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAdminUserFields

`func NewAdminUserFields(id int32, username string, email string, orgId int32, orgName string, ) *AdminUserFields`

NewAdminUserFields instantiates a new AdminUserFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserFieldsWithDefaults

`func NewAdminUserFieldsWithDefaults() *AdminUserFields`

NewAdminUserFieldsWithDefaults instantiates a new AdminUserFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminUserFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminUserFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminUserFields) SetId(v int32)`

SetId sets Id field to given value.


### GetUsername

`func (o *AdminUserFields) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AdminUserFields) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AdminUserFields) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *AdminUserFields) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminUserFields) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminUserFields) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrgId

`func (o *AdminUserFields) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AdminUserFields) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AdminUserFields) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetOrgName

`func (o *AdminUserFields) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *AdminUserFields) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *AdminUserFields) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetCreatedAt

`func (o *AdminUserFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminUserFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminUserFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminUserFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastLogin

`func (o *AdminUserFields) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *AdminUserFields) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *AdminUserFields) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *AdminUserFields) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


