# \OrganizationAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationInfo**](OrganizationAPI.md#GetOrganizationInfo) | **Get** /auth/organizations | Organization Info
[**RemoveAMemberFromOrganization**](OrganizationAPI.md#RemoveAMemberFromOrganization) | **Post** /auth/organizations/remove-member | Remove a member from organization
[**UpdateOrganizationInfo**](OrganizationAPI.md#UpdateOrganizationInfo) | **Put** /auth/organizations/update | Update organization info



## GetOrganizationInfo

> GetOrganizationResponseModel GetOrganizationInfo(ctx).Execute()

Organization Info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/szczad/hyperstack-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetOrganizationInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganizationInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInfo`: GetOrganizationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetOrganizationInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInfoRequest struct via the builder pattern


### Return type

[**GetOrganizationResponseModel**](GetOrganizationResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAMemberFromOrganization

> RemoveMemberFromOrganizationResponseModel RemoveAMemberFromOrganization(ctx).Payload(payload).Execute()

Remove a member from organization

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/szczad/hyperstack-api-go"
)

func main() {
	payload := *openapiclient.NewRemoveMemberPayload("Email_example") // RemoveMemberPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.RemoveAMemberFromOrganization(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveAMemberFromOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAMemberFromOrganization`: RemoveMemberFromOrganizationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.RemoveAMemberFromOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAMemberFromOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**RemoveMemberPayload**](RemoveMemberPayload.md) |  | 

### Return type

[**RemoveMemberFromOrganizationResponseModel**](RemoveMemberFromOrganizationResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationInfo

> UpdateOrganizationResponseModel UpdateOrganizationInfo(ctx).Payload(payload).Execute()

Update organization info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/szczad/hyperstack-api-go"
)

func main() {
	payload := *openapiclient.NewUpdateOrganizationPayload("Name_example") // UpdateOrganizationPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateOrganizationInfo(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateOrganizationInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationInfo`: UpdateOrganizationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateOrganizationInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**UpdateOrganizationPayload**](UpdateOrganizationPayload.md) |  | 

### Return type

[**UpdateOrganizationResponseModel**](UpdateOrganizationResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

