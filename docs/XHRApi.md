# \XHRApi

All URIs are relative to *https://api.affixapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**XhrCompanies20230301**](XHRApi.md#XhrCompanies20230301) | **Get** /2023-03-01/xhr/company | Company
[**XhrCreateEmployee20230301**](XHRApi.md#XhrCreateEmployee20230301) | **Post** /2023-03-01/xhr/employee | Create employee
[**XhrEmployees20230301**](XHRApi.md#XhrEmployees20230301) | **Get** /2023-03-01/xhr/employees | Employees
[**XhrGroups20230301**](XHRApi.md#XhrGroups20230301) | **Get** /2023-03-01/xhr/groups | Groups
[**XhrIdentity20230301**](XHRApi.md#XhrIdentity20230301) | **Get** /2023-03-01/xhr/identity | Identity
[**XhrPayruns20230301**](XHRApi.md#XhrPayruns20230301) | **Get** /2023-03-01/xhr/payruns | Payruns
[**XhrPayslips20230301**](XHRApi.md#XhrPayslips20230301) | **Get** /2023-03-01/xhr/payruns/{payrun_id} | Payslips
[**XhrTimeOffBalances20230301**](XHRApi.md#XhrTimeOffBalances20230301) | **Get** /2023-03-01/xhr/time-off-balances | Time off balances
[**XhrTimeOffEntries20230301**](XHRApi.md#XhrTimeOffEntries20230301) | **Get** /2023-03-01/xhr/time-off-entries | Time off entries
[**XhrTimesheets20230301**](XHRApi.md#XhrTimesheets20230301) | **Get** /2023-03-01/xhr/timesheets | Timesheets
[**XhrWorkLocations20230301**](XHRApi.md#XhrWorkLocations20230301) | **Get** /2023-03-01/xhr/work-locations | Work locations



## XhrCompanies20230301

> []CompanyResponse XhrCompanies20230301(ctx).Execute()

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
    resp, r, err := api_client.XHRApi.XhrCompanies20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XHRApi.XhrCompanies20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XhrCompanies20230301`: []CompanyResponse
    fmt.Fprintf(os.Stdout, "Response from `XHRApi.XhrCompanies20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiXhrCompanies20230301Request struct via the builder pattern


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


## XhrCreateEmployee20230301

> EmployeeResponse XhrCreateEmployee20230301(ctx).CreateEmployeeRequest(createEmployeeRequest).Execute()

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
    resp, r, err := api_client.XHRApi.XhrCreateEmployee20230301(context.Background()).CreateEmployeeRequest(createEmployeeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XHRApi.XhrCreateEmployee20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XhrCreateEmployee20230301`: EmployeeResponse
    fmt.Fprintf(os.Stdout, "Response from `XHRApi.XhrCreateEmployee20230301`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXhrCreateEmployee20230301Request struct via the builder pattern


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


## XhrEmployees20230301

> []EmployeeResponse XhrEmployees20230301(ctx).Execute()

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
    resp, r, err := api_client.XHRApi.XhrEmployees20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XHRApi.XhrEmployees20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XhrEmployees20230301`: []EmployeeResponse
    fmt.Fprintf(os.Stdout, "Response from `XHRApi.XhrEmployees20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiXhrEmployees20230301Request struct via the builder pattern


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


## XhrGroups20230301

> []GroupResponse XhrGroups20230301(ctx).Execute()

Groups



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
    resp, r, err := api_client.XHRApi.XhrGroups20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XHRApi.XhrGroups20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XhrGroups20230301`: []GroupResponse
    fmt.Fprintf(os.Stdout, "Response from `XHRApi.XhrGroups20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiXhrGroups20230301Request struct via the builder pattern


### Return type

[**[]GroupResponse**](GroupResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XhrIdentity20230301

> IdentityResponse XhrIdentity20230301(ctx).Execute()

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
    resp, r, err := api_client.XHRApi.XhrIdentity20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XHRApi.XhrIdentity20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XhrIdentity20230301`: IdentityResponse
    fmt.Fprintf(os.Stdout, "Response from `XHRApi.XhrIdentity20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiXhrIdentity20230301Request struct via the builder pattern


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


## XhrPayruns20230301

> []PayrunResponse XhrPayruns20230301(ctx).StartDate(startDate).EndDate(endDate).Execute()

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
    resp, r, err := api_client.XHRApi.XhrPayruns20230301(context.Background()).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XHRApi.XhrPayruns20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XhrPayruns20230301`: []PayrunResponse
    fmt.Fprintf(os.Stdout, "Response from `XHRApi.XhrPayruns20230301`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXhrPayruns20230301Request struct via the builder pattern


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


## XhrPayslips20230301

> []PayslipResponse XhrPayslips20230301(ctx, payrunId).Execute()

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
    resp, r, err := api_client.XHRApi.XhrPayslips20230301(context.Background(), payrunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XHRApi.XhrPayslips20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XhrPayslips20230301`: []PayslipResponse
    fmt.Fprintf(os.Stdout, "Response from `XHRApi.XhrPayslips20230301`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payrunId** | **string** | The id of the payrun. | 

### Other Parameters

Other parameters are passed through a pointer to a apiXhrPayslips20230301Request struct via the builder pattern


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


## XhrTimeOffBalances20230301

> []TimeOffBalanceResponse XhrTimeOffBalances20230301(ctx).Execute()

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
    resp, r, err := api_client.XHRApi.XhrTimeOffBalances20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XHRApi.XhrTimeOffBalances20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XhrTimeOffBalances20230301`: []TimeOffBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `XHRApi.XhrTimeOffBalances20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiXhrTimeOffBalances20230301Request struct via the builder pattern


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


## XhrTimeOffEntries20230301

> []TimeOffEntryResponse XhrTimeOffEntries20230301(ctx).Execute()

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
    resp, r, err := api_client.XHRApi.XhrTimeOffEntries20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XHRApi.XhrTimeOffEntries20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XhrTimeOffEntries20230301`: []TimeOffEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `XHRApi.XhrTimeOffEntries20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiXhrTimeOffEntries20230301Request struct via the builder pattern


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


## XhrTimesheets20230301

> []TimesheetResponse XhrTimesheets20230301(ctx).Execute()

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
    resp, r, err := api_client.XHRApi.XhrTimesheets20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XHRApi.XhrTimesheets20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XhrTimesheets20230301`: []TimesheetResponse
    fmt.Fprintf(os.Stdout, "Response from `XHRApi.XhrTimesheets20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiXhrTimesheets20230301Request struct via the builder pattern


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


## XhrWorkLocations20230301

> []LocationResponse XhrWorkLocations20230301(ctx).Execute()

Work locations



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
    resp, r, err := api_client.XHRApi.XhrWorkLocations20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XHRApi.XhrWorkLocations20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XhrWorkLocations20230301`: []LocationResponse
    fmt.Fprintf(os.Stdout, "Response from `XHRApi.XhrWorkLocations20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiXhrWorkLocations20230301Request struct via the builder pattern


### Return type

[**[]LocationResponse**](LocationResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
