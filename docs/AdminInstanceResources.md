# AdminInstanceResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OpenstackId** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Flavor** | Pointer to [**AdminFlavorResource**](AdminFlavorResource.md) |  | [optional] 
**ImageId** | Pointer to **int32** |  | [optional] 
**VolumeId** | Pointer to **int32** |  | [optional] 
**KeypairName** | Pointer to **string** |  | [optional] 
**FixedIp** | Pointer to **string** |  | [optional] 
**FloatingIp** | Pointer to **string** |  | [optional] 
**ContractId** | Pointer to **int32** |  | [optional] 
**FloatingIpStatus** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAdminInstanceResources

`func NewAdminInstanceResources() *AdminInstanceResources`

NewAdminInstanceResources instantiates a new AdminInstanceResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminInstanceResourcesWithDefaults

`func NewAdminInstanceResourcesWithDefaults() *AdminInstanceResources`

NewAdminInstanceResourcesWithDefaults instantiates a new AdminInstanceResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminInstanceResources) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminInstanceResources) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminInstanceResources) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdminInstanceResources) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AdminInstanceResources) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminInstanceResources) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminInstanceResources) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminInstanceResources) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpenstackId

`func (o *AdminInstanceResources) GetOpenstackId() string`

GetOpenstackId returns the OpenstackId field if non-nil, zero value otherwise.

### GetOpenstackIdOk

`func (o *AdminInstanceResources) GetOpenstackIdOk() (*string, bool)`

GetOpenstackIdOk returns a tuple with the OpenstackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackId

`func (o *AdminInstanceResources) SetOpenstackId(v string)`

SetOpenstackId sets OpenstackId field to given value.

### HasOpenstackId

`func (o *AdminInstanceResources) HasOpenstackId() bool`

HasOpenstackId returns a boolean if a field has been set.

### GetHost

`func (o *AdminInstanceResources) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AdminInstanceResources) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AdminInstanceResources) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AdminInstanceResources) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetStatus

`func (o *AdminInstanceResources) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminInstanceResources) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminInstanceResources) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminInstanceResources) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFlavor

`func (o *AdminInstanceResources) GetFlavor() AdminFlavorResource`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *AdminInstanceResources) GetFlavorOk() (*AdminFlavorResource, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *AdminInstanceResources) SetFlavor(v AdminFlavorResource)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *AdminInstanceResources) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetImageId

`func (o *AdminInstanceResources) GetImageId() int32`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AdminInstanceResources) GetImageIdOk() (*int32, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AdminInstanceResources) SetImageId(v int32)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *AdminInstanceResources) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetVolumeId

`func (o *AdminInstanceResources) GetVolumeId() int32`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *AdminInstanceResources) GetVolumeIdOk() (*int32, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *AdminInstanceResources) SetVolumeId(v int32)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *AdminInstanceResources) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetKeypairName

`func (o *AdminInstanceResources) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *AdminInstanceResources) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *AdminInstanceResources) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.

### HasKeypairName

`func (o *AdminInstanceResources) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### GetFixedIp

`func (o *AdminInstanceResources) GetFixedIp() string`

GetFixedIp returns the FixedIp field if non-nil, zero value otherwise.

### GetFixedIpOk

`func (o *AdminInstanceResources) GetFixedIpOk() (*string, bool)`

GetFixedIpOk returns a tuple with the FixedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIp

`func (o *AdminInstanceResources) SetFixedIp(v string)`

SetFixedIp sets FixedIp field to given value.

### HasFixedIp

`func (o *AdminInstanceResources) HasFixedIp() bool`

HasFixedIp returns a boolean if a field has been set.

### GetFloatingIp

`func (o *AdminInstanceResources) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *AdminInstanceResources) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *AdminInstanceResources) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *AdminInstanceResources) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.

### GetContractId

`func (o *AdminInstanceResources) GetContractId() int32`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *AdminInstanceResources) GetContractIdOk() (*int32, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *AdminInstanceResources) SetContractId(v int32)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *AdminInstanceResources) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetFloatingIpStatus

`func (o *AdminInstanceResources) GetFloatingIpStatus() string`

GetFloatingIpStatus returns the FloatingIpStatus field if non-nil, zero value otherwise.

### GetFloatingIpStatusOk

`func (o *AdminInstanceResources) GetFloatingIpStatusOk() (*string, bool)`

GetFloatingIpStatusOk returns a tuple with the FloatingIpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpStatus

`func (o *AdminInstanceResources) SetFloatingIpStatus(v string)`

SetFloatingIpStatus sets FloatingIpStatus field to given value.

### HasFloatingIpStatus

`func (o *AdminInstanceResources) HasFloatingIpStatus() bool`

HasFloatingIpStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminInstanceResources) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminInstanceResources) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminInstanceResources) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminInstanceResources) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


