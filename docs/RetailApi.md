# \RetailApi

All URIs are relative to *https://api.affixapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrder20230320**](RetailApi.md#GetOrder20230320) | **Get** /2023-03-20/retail/orders/{order_id} | Get an Order
[**Identity20230320**](RetailApi.md#Identity20230320) | **Get** /2023-03-20/retail/identity | Identity
[**Orders20230320**](RetailApi.md#Orders20230320) | **Get** /2023-03-20/retail/orders | Orders



## GetOrder20230320

> Order20230320Response GetOrder20230320(ctx, orderId).Execute()

Get an Order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderId := "orderId_example" // string | The id of the order.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RetailApi.GetOrder20230320(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetailApi.GetOrder20230320``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrder20230320`: Order20230320Response
    fmt.Fprintf(os.Stdout, "Response from `RetailApi.GetOrder20230320`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The id of the order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrder20230320Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Order20230320Response**](Order20230320Response.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Identity20230320

> IdentityResponse Identity20230320(ctx).Execute()

Identity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RetailApi.Identity20230320(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetailApi.Identity20230320``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Identity20230320`: IdentityResponse
    fmt.Fprintf(os.Stdout, "Response from `RetailApi.Identity20230320`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIdentity20230320Request struct via the builder pattern


### Return type

[**IdentityResponse**](IdentityResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Orders20230320

> []map[string]interface{} Orders20230320(ctx).StartDate(startDate).EndDate(endDate).Execute()

Orders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    startDate := time.Now() // string | The start date of the search period
    endDate := time.Now() // string | The end date of the search period

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RetailApi.Orders20230320(context.Background()).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetailApi.Orders20230320``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Orders20230320`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RetailApi.Orders20230320`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrders20230320Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | The start date of the search period | 
 **endDate** | **string** | The end date of the search period | 

### Return type

**[]map[string]interface{}**

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

