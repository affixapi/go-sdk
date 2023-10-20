# PayslipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Affix-assigned id of the payslip | 
**RemoteId** | **string** | the remote system-assigned id of the payrun | 
**EmployeeId** | **string** |  | 
**PayrunId** | **string** |  | 
**Currency** | **string** |  | 
**GrossPay** | **float32** | if USD/EUR/GBP, in cent | 
**NetPay** | **float32** | if USD/EUR/GBP, in cent | 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**CheckDate** | **string** |  | 
**Earnings** | [**[]PayslipResponseEarnings**](PayslipResponseEarnings.md) |  | 
**Deductions** | [**[]PayslipResponseDeductions**](PayslipResponseDeductions.md) |  | 
**Taxes** | [**[]PayslipResponseTaxes**](PayslipResponseTaxes.md) |  | 
**FieldMappings** | **map[string]interface{}** |  | 

## Methods

### NewPayslipResponse

`func NewPayslipResponse(id string, remoteId string, employeeId string, payrunId string, currency string, grossPay float32, netPay float32, startDate string, endDate string, checkDate string, earnings []PayslipResponseEarnings, deductions []PayslipResponseDeductions, taxes []PayslipResponseTaxes, fieldMappings map[string]interface{}, ) *PayslipResponse`

NewPayslipResponse instantiates a new PayslipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayslipResponseWithDefaults

`func NewPayslipResponseWithDefaults() *PayslipResponse`

NewPayslipResponseWithDefaults instantiates a new PayslipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PayslipResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayslipResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayslipResponse) SetId(v string)`

SetId sets Id field to given value.


### GetRemoteId

`func (o *PayslipResponse) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayslipResponse) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayslipResponse) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetEmployeeId

`func (o *PayslipResponse) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *PayslipResponse) GetEmployeeIdOk() (*string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId

`func (o *PayslipResponse) SetEmployeeId(v string)`

SetEmployeeId sets EmployeeId field to given value.


### GetPayrunId

`func (o *PayslipResponse) GetPayrunId() string`

GetPayrunId returns the PayrunId field if non-nil, zero value otherwise.

### GetPayrunIdOk

`func (o *PayslipResponse) GetPayrunIdOk() (*string, bool)`

GetPayrunIdOk returns a tuple with the PayrunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrunId

`func (o *PayslipResponse) SetPayrunId(v string)`

SetPayrunId sets PayrunId field to given value.


### GetCurrency

`func (o *PayslipResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PayslipResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PayslipResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetGrossPay

`func (o *PayslipResponse) GetGrossPay() float32`

GetGrossPay returns the GrossPay field if non-nil, zero value otherwise.

### GetGrossPayOk

`func (o *PayslipResponse) GetGrossPayOk() (*float32, bool)`

GetGrossPayOk returns a tuple with the GrossPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossPay

`func (o *PayslipResponse) SetGrossPay(v float32)`

SetGrossPay sets GrossPay field to given value.


### GetNetPay

`func (o *PayslipResponse) GetNetPay() float32`

GetNetPay returns the NetPay field if non-nil, zero value otherwise.

### GetNetPayOk

`func (o *PayslipResponse) GetNetPayOk() (*float32, bool)`

GetNetPayOk returns a tuple with the NetPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPay

`func (o *PayslipResponse) SetNetPay(v float32)`

SetNetPay sets NetPay field to given value.


### GetStartDate

`func (o *PayslipResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PayslipResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PayslipResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *PayslipResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PayslipResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PayslipResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetCheckDate

`func (o *PayslipResponse) GetCheckDate() string`

GetCheckDate returns the CheckDate field if non-nil, zero value otherwise.

### GetCheckDateOk

`func (o *PayslipResponse) GetCheckDateOk() (*string, bool)`

GetCheckDateOk returns a tuple with the CheckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckDate

`func (o *PayslipResponse) SetCheckDate(v string)`

SetCheckDate sets CheckDate field to given value.


### GetEarnings

`func (o *PayslipResponse) GetEarnings() []PayslipResponseEarnings`

GetEarnings returns the Earnings field if non-nil, zero value otherwise.

### GetEarningsOk

`func (o *PayslipResponse) GetEarningsOk() (*[]PayslipResponseEarnings, bool)`

GetEarningsOk returns a tuple with the Earnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnings

`func (o *PayslipResponse) SetEarnings(v []PayslipResponseEarnings)`

SetEarnings sets Earnings field to given value.


### GetDeductions

`func (o *PayslipResponse) GetDeductions() []PayslipResponseDeductions`

GetDeductions returns the Deductions field if non-nil, zero value otherwise.

### GetDeductionsOk

`func (o *PayslipResponse) GetDeductionsOk() (*[]PayslipResponseDeductions, bool)`

GetDeductionsOk returns a tuple with the Deductions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductions

`func (o *PayslipResponse) SetDeductions(v []PayslipResponseDeductions)`

SetDeductions sets Deductions field to given value.


### GetTaxes

`func (o *PayslipResponse) GetTaxes() []PayslipResponseTaxes`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *PayslipResponse) GetTaxesOk() (*[]PayslipResponseTaxes, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *PayslipResponse) SetTaxes(v []PayslipResponseTaxes)`

SetTaxes sets Taxes field to given value.


### GetFieldMappings

`func (o *PayslipResponse) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *PayslipResponse) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *PayslipResponse) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.


### SetFieldMappingsNil

`func (o *PayslipResponse) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *PayslipResponse) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


