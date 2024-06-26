# \CustomerContractAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DetailsOfContractByIDForCustomer**](CustomerContractAPI.md#DetailsOfContractByIDForCustomer) | **Get** /pricebook/contracts/{contract_id} | Retrieve Contract Details
[**GetCustomerContract**](CustomerContractAPI.md#GetCustomerContract) | **Get** /pricebook/contracts | List Contracts



## DetailsOfContractByIDForCustomer

> CustomerContractDetailResponseModel DetailsOfContractByIDForCustomer(ctx, contractId).Execute()

Retrieve Contract Details

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
	contractId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerContractAPI.DetailsOfContractByIDForCustomer(context.Background(), contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerContractAPI.DetailsOfContractByIDForCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetailsOfContractByIDForCustomer`: CustomerContractDetailResponseModel
	fmt.Fprintf(os.Stdout, "Response from `CustomerContractAPI.DetailsOfContractByIDForCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetailsOfContractByIDForCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerContractDetailResponseModel**](CustomerContractDetailResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerContract

> GetCustomerContractsListResponseModel GetCustomerContract(ctx).Page(page).PerPage(perPage).Execute()

List Contracts



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
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerContractAPI.GetCustomerContract(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerContractAPI.GetCustomerContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerContract`: GetCustomerContractsListResponseModel
	fmt.Fprintf(os.Stdout, "Response from `CustomerContractAPI.GetCustomerContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**GetCustomerContractsListResponseModel**](GetCustomerContractsListResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

