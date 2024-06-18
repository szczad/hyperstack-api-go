# \FirewallsAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRulesToFirewall**](FirewallsAPI.md#AddRulesToFirewall) | **Post** /core/firewalls/{firewall_id}/firewall-rules | Add Rules to Firewall
[**CreateFirewall**](FirewallsAPI.md#CreateFirewall) | **Post** /core/firewalls | Create Firewall
[**DeleteFirewall**](FirewallsAPI.md#DeleteFirewall) | **Delete** /core/firewalls/{id} | Delete Firewall
[**DeleteFirewallRulesFromFirewall**](FirewallsAPI.md#DeleteFirewallRulesFromFirewall) | **Delete** /core/firewalls/{firewall_id}/firewall-rules/{firewall_rule_id} | Delete Firewall Rules from Firewall
[**DetailsOfFirewallByID**](FirewallsAPI.md#DetailsOfFirewallByID) | **Get** /core/firewalls/{id} | Details of Firewall by ID
[**RetrieveFirewalls**](FirewallsAPI.md#RetrieveFirewalls) | **Get** /core/firewalls | List firewalls



## AddRulesToFirewall

> FirewallRule AddRulesToFirewall(ctx, firewallId).Payload(payload).Execute()

Add Rules to Firewall

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
	firewallId := int32(56) // int32 | 
	payload := *openapiclient.NewCreateFirewallRulePayload("Direction_example", "Protocol_example", "Ethertype_example", "RemoteIpPrefix_example") // CreateFirewallRulePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsAPI.AddRulesToFirewall(context.Background(), firewallId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.AddRulesToFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRulesToFirewall`: FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.AddRulesToFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRulesToFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**CreateFirewallRulePayload**](CreateFirewallRulePayload.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFirewall

> FirewallResponse CreateFirewall(ctx).Payload(payload).Execute()

Create Firewall

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
	payload := *openapiclient.NewCreateFirewallPayload("Name_example", int32(123)) // CreateFirewallPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsAPI.CreateFirewall(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.CreateFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirewall`: FirewallResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.CreateFirewall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateFirewallPayload**](CreateFirewallPayload.md) |  | 

### Return type

[**FirewallResponse**](FirewallResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewall

> ResponseModel DeleteFirewall(ctx, id).Execute()

Delete Firewall

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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsAPI.DeleteFirewall(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.DeleteFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewall`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.DeleteFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewallRulesFromFirewall

> ResponseModel DeleteFirewallRulesFromFirewall(ctx, firewallId, firewallRuleId).Execute()

Delete Firewall Rules from Firewall

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
	firewallId := int32(56) // int32 | 
	firewallRuleId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsAPI.DeleteFirewallRulesFromFirewall(context.Background(), firewallId, firewallRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.DeleteFirewallRulesFromFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewallRulesFromFirewall`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.DeleteFirewallRulesFromFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** |  | 
**firewallRuleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRulesFromFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetailsOfFirewallByID

> FirewallDetailResponse DetailsOfFirewallByID(ctx, id).Execute()

Details of Firewall by ID

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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsAPI.DetailsOfFirewallByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.DetailsOfFirewallByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetailsOfFirewallByID`: FirewallDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.DetailsOfFirewallByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetailsOfFirewallByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirewallDetailResponse**](FirewallDetailResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFirewalls

> FirewallsListResponse RetrieveFirewalls(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

List firewalls

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
	pageSize := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsAPI.RetrieveFirewalls(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.RetrieveFirewalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveFirewalls`: FirewallsListResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.RetrieveFirewalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFirewallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**FirewallsListResponse**](FirewallsListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

