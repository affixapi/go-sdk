# \DeveloperApi

All URIs are relative to *https://api.affixapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeveloperEmployees20230301**](DeveloperApi.md#DeveloperEmployees20230301) | **Get** /2023-03-01/developer/employees | Employees
[**DeveloperIdentity20230301**](DeveloperApi.md#DeveloperIdentity20230301) | **Get** /2023-03-01/developer/identity | Identity
[**DeveloperPayruns20230301**](DeveloperApi.md#DeveloperPayruns20230301) | **Get** /2023-03-01/developer/payruns | Payruns
[**DeveloperPayslips20230320**](DeveloperApi.md#DeveloperPayslips20230320) | **Get** /2023-03-01/developer/payruns/{payrun_id} | Get payslips of a payrun (the payslips/pay stubs/check stubs + detail)



## DeveloperEmployees20230301

> []Employee20230301Response DeveloperEmployees20230301(ctx).Execute()

Employees



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
    resp, r, err := api_client.DeveloperApi.DeveloperEmployees20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperApi.DeveloperEmployees20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeveloperEmployees20230301`: []Employee20230301Response
    fmt.Fprintf(os.Stdout, "Response from `DeveloperApi.DeveloperEmployees20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeveloperEmployees20230301Request struct via the builder pattern


### Return type

[**[]Employee20230301Response**](Employee20230301Response.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeveloperIdentity20230301

> IdentityResponse DeveloperIdentity20230301(ctx).Execute()

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
    resp, r, err := api_client.DeveloperApi.DeveloperIdentity20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperApi.DeveloperIdentity20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeveloperIdentity20230301`: IdentityResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperApi.DeveloperIdentity20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeveloperIdentity20230301Request struct via the builder pattern


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


## DeveloperPayruns20230301

> []PayrunResponse DeveloperPayruns20230301(ctx).StartDate(startDate).EndDate(endDate).Execute()

Payruns



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
    resp, r, err := api_client.DeveloperApi.DeveloperPayruns20230301(context.Background()).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperApi.DeveloperPayruns20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeveloperPayruns20230301`: []PayrunResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperApi.DeveloperPayruns20230301`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeveloperPayruns20230301Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | The start date of the search period | 
 **endDate** | **string** | The end date of the search period | 

### Return type

[**[]PayrunResponse**](PayrunResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeveloperPayslips20230320

> []PayslipResponse DeveloperPayslips20230320(ctx, payrunId).Execute()

Get payslips of a payrun (the payslips/pay stubs/check stubs + detail)



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
    payrunId := "payrunId_example" // string | The id of the payrun.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeveloperApi.DeveloperPayslips20230320(context.Background(), payrunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperApi.DeveloperPayslips20230320``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeveloperPayslips20230320`: []PayslipResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperApi.DeveloperPayslips20230320`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payrunId** | **string** | The id of the payrun. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeveloperPayslips20230320Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PayslipResponse**](PayslipResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

