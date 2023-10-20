# \HrisApi

All URIs are relative to *https://api.affixapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HrisEmployees20230301**](HrisApi.md#HrisEmployees20230301) | **Get** /2023-03-01/hris/employees | Employees
[**HrisIdentity20230301**](HrisApi.md#HrisIdentity20230301) | **Get** /2023-03-01/hris/identity | Identity



## HrisEmployees20230301

> []Employee20230301Response HrisEmployees20230301(ctx).Execute()

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
    resp, r, err := api_client.HrisApi.HrisEmployees20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HrisApi.HrisEmployees20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HrisEmployees20230301`: []Employee20230301Response
    fmt.Fprintf(os.Stdout, "Response from `HrisApi.HrisEmployees20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHrisEmployees20230301Request struct via the builder pattern


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


## HrisIdentity20230301

> IdentityResponse HrisIdentity20230301(ctx).Execute()

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
    resp, r, err := api_client.HrisApi.HrisIdentity20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HrisApi.HrisIdentity20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HrisIdentity20230301`: IdentityResponse
    fmt.Fprintf(os.Stdout, "Response from `HrisApi.HrisIdentity20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHrisIdentity20230301Request struct via the builder pattern


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

