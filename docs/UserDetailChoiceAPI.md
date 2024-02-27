# \UserDetailChoiceAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveDefaultFlavorsAndImagesForUser**](UserDetailChoiceAPI.md#RetrieveDefaultFlavorsAndImagesForUser) | **Get** /core/user/resources/defaults | Retrieve default flavors and images for user



## RetrieveDefaultFlavorsAndImagesForUser

> UserDefaultChoicesForUserResponse RetrieveDefaultFlavorsAndImagesForUser(ctx).Execute()

Retrieve default flavors and images for user

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
	resp, r, err := apiClient.UserDetailChoiceAPI.RetrieveDefaultFlavorsAndImagesForUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDetailChoiceAPI.RetrieveDefaultFlavorsAndImagesForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDefaultFlavorsAndImagesForUser`: UserDefaultChoicesForUserResponse
	fmt.Fprintf(os.Stdout, "Response from `UserDetailChoiceAPI.RetrieveDefaultFlavorsAndImagesForUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDefaultFlavorsAndImagesForUserRequest struct via the builder pattern


### Return type

[**UserDefaultChoicesForUserResponse**](UserDefaultChoicesForUserResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

