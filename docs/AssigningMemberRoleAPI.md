# \AssigningMemberRoleAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRBACRoles**](AssigningMemberRoleAPI.md#AssignRBACRoles) | **Put** /auth/users/{user_id}/assign-roles | Assign RBAC Roles
[**RemoveRoleFromAUser**](AssigningMemberRoleAPI.md#RemoveRoleFromAUser) | **Delete** /auth/users/{user_id}/roles | Remove role from a user



## AssignRBACRoles

> RbacRoleDetailResponseModel AssignRBACRoles(ctx, userId).Payload(payload).Execute()

Assign RBAC Roles

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
	userId := int32(56) // int32 | 
	payload := *openapiclient.NewAssignRbacRolePayload(int32(123)) // AssignRbacRolePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssigningMemberRoleAPI.AssignRBACRoles(context.Background(), userId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssigningMemberRoleAPI.AssignRBACRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignRBACRoles`: RbacRoleDetailResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AssigningMemberRoleAPI.AssignRBACRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignRBACRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AssignRbacRolePayload**](AssignRbacRolePayload.md) |  | 

### Return type

[**RbacRoleDetailResponseModel**](RbacRoleDetailResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRoleFromAUser

> CommonResponseModel RemoveRoleFromAUser(ctx, userId).Execute()

Remove role from a user

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
	userId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssigningMemberRoleAPI.RemoveRoleFromAUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssigningMemberRoleAPI.RemoveRoleFromAUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveRoleFromAUser`: CommonResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AssigningMemberRoleAPI.RemoveRoleFromAUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleFromAUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonResponseModel**](CommonResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

