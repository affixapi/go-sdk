# \DeveloperApi

All URIs are relative to *https://api.affixapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeveloperCompanies20230301**](DeveloperApi.md#DeveloperCompanies20230301) | **Get** /2023-03-01/developer/company | Company
[**DeveloperCreateEmployee20230301**](DeveloperApi.md#DeveloperCreateEmployee20230301) | **Post** /2023-03-01/developer/employee | Create employee
[**DeveloperEmployees20230301**](DeveloperApi.md#DeveloperEmployees20230301) | **Get** /2023-03-01/developer/employees | Employees
[**DeveloperIdentity20230301**](DeveloperApi.md#DeveloperIdentity20230301) | **Get** /2023-03-01/developer/identity | Identity
[**DeveloperPayruns20230301**](DeveloperApi.md#DeveloperPayruns20230301) | **Get** /2023-03-01/developer/payruns | Payruns
[**DeveloperPayslips20230301**](DeveloperApi.md#DeveloperPayslips20230301) | **Get** /2023-03-01/developer/payruns/{payrun_id} | Payslips
[**DeveloperTimeOffBalances20230301**](DeveloperApi.md#DeveloperTimeOffBalances20230301) | **Get** /2023-03-01/developer/time-off-balances | Time off balances
[**DeveloperTimeOffEntries20230301**](DeveloperApi.md#DeveloperTimeOffEntries20230301) | **Get** /2023-03-01/developer/time-off-entries | Time off entries
[**DeveloperTimesheets20230301**](DeveloperApi.md#DeveloperTimesheets20230301) | **Get** /2023-03-01/developer/timesheets | Timesheets



## DeveloperCompanies20230301

> []CompanyResponse DeveloperCompanies20230301(ctx).Execute()

Company



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
    resp, r, err := api_client.DeveloperApi.DeveloperCompanies20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperApi.DeveloperCompanies20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeveloperCompanies20230301`: []CompanyResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperApi.DeveloperCompanies20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeveloperCompanies20230301Request struct via the builder pattern


### Return type

[**[]CompanyResponse**](CompanyResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeveloperCreateEmployee20230301

> EmployeeResponse DeveloperCreateEmployee20230301(ctx).CreateEmployeeRequest(createEmployeeRequest).Execute()

Create employee



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
    createEmployeeRequest := *openapiclient.NewCreateEmployeeRequest("Greg", "Hirsch") // CreateEmployeeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeveloperApi.DeveloperCreateEmployee20230301(context.Background()).CreateEmployeeRequest(createEmployeeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperApi.DeveloperCreateEmployee20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeveloperCreateEmployee20230301`: EmployeeResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperApi.DeveloperCreateEmployee20230301`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeveloperCreateEmployee20230301Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEmployeeRequest** | [**CreateEmployeeRequest**](CreateEmployeeRequest.md) |  | 

### Return type

[**EmployeeResponse**](EmployeeResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeveloperEmployees20230301

> []EmployeeResponse DeveloperEmployees20230301(ctx).Execute()

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
    // response from `DeveloperEmployees20230301`: []EmployeeResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperApi.DeveloperEmployees20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeveloperEmployees20230301Request struct via the builder pattern


### Return type

[**[]EmployeeResponse**](EmployeeResponse.md)

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


## DeveloperPayslips20230301

> []PayslipResponse DeveloperPayslips20230301(ctx, payrunId).Execute()

Payslips



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
    resp, r, err := api_client.DeveloperApi.DeveloperPayslips20230301(context.Background(), payrunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperApi.DeveloperPayslips20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeveloperPayslips20230301`: []PayslipResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperApi.DeveloperPayslips20230301`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payrunId** | **string** | The id of the payrun. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeveloperPayslips20230301Request struct via the builder pattern


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


## DeveloperTimeOffBalances20230301

> []TimeOffBalanceResponse DeveloperTimeOffBalances20230301(ctx).Execute()

Time off balances



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
    resp, r, err := api_client.DeveloperApi.DeveloperTimeOffBalances20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperApi.DeveloperTimeOffBalances20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeveloperTimeOffBalances20230301`: []TimeOffBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperApi.DeveloperTimeOffBalances20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeveloperTimeOffBalances20230301Request struct via the builder pattern


### Return type

[**[]TimeOffBalanceResponse**](TimeOffBalanceResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeveloperTimeOffEntries20230301

> []TimeOffEntryResponse DeveloperTimeOffEntries20230301(ctx).Execute()

Time off entries



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
    resp, r, err := api_client.DeveloperApi.DeveloperTimeOffEntries20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperApi.DeveloperTimeOffEntries20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeveloperTimeOffEntries20230301`: []TimeOffEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperApi.DeveloperTimeOffEntries20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeveloperTimeOffEntries20230301Request struct via the builder pattern


### Return type

[**[]TimeOffEntryResponse**](TimeOffEntryResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeveloperTimesheets20230301

> []TimesheetResponse DeveloperTimesheets20230301(ctx).Execute()

Timesheets



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
    resp, r, err := api_client.DeveloperApi.DeveloperTimesheets20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperApi.DeveloperTimesheets20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeveloperTimesheets20230301`: []TimesheetResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperApi.DeveloperTimesheets20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeveloperTimesheets20230301Request struct via the builder pattern


### Return type

[**[]TimesheetResponse**](TimesheetResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

