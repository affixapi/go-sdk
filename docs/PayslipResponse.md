# PayslipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The Affix-assigned id of the payslip | 
**RemoteId** | **NullableString** | the remote system-assigned id of the payrun | 
**EmployeeId** | **string** |  | 
**EmployeeRemoteId** | **string** |  | 
**PayrunId** | **string** |  | 
**PayrunRemoteId** | **string** |  | 
**PayrunType** | [**NullablePayrunTypeResponse**](PayrunTypeResponse.md) |  | 
**Currency** | [**NullableCurrencyNotNullResponse**](CurrencyNotNullResponse.md) |  | 
**GrossPay** | **NullableFloat32** | if USD/EUR/GBP, in cent | 
**NetPay** | **NullableFloat32** | if USD/EUR/GBP, in cent | 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**PaymentDate** | **string** |  | 
**Earnings** | [**[]PayslipResponseEarnings**](PayslipResponseEarnings.md) |  | 
**Contributions** | [**[]PayslipResponseContributions**](PayslipResponseContributions.md) | Items paid by the employer that are not included in gross pay, such as employer-paid portion of private health insurance  | 
**Deductions** | [**[]PayslipResponseDeductions**](PayslipResponseDeductions.md) |  | 
**Taxes** | [**[]PayslipResponseTaxes**](PayslipResponseTaxes.md) |  | 
**Reimbursements** | [**[]PayslipResponseReimbursements**](PayslipResponseReimbursements.md) |  | 

## Methods

### NewPayslipResponse

`func NewPayslipResponse(id NullableString, remoteId NullableString, employeeId string, employeeRemoteId string, payrunId string, payrunRemoteId string, payrunType NullablePayrunTypeResponse, currency NullableCurrencyNotNullResponse, grossPay NullableFloat32, netPay NullableFloat32, startDate string, endDate string, paymentDate string, earnings []PayslipResponseEarnings, contributions []PayslipResponseContributions, deductions []PayslipResponseDeductions, taxes []PayslipResponseTaxes, reimbursements []PayslipResponseReimbursements, ) *PayslipResponse`

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


### SetIdNil

`func (o *PayslipResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PayslipResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
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


### SetRemoteIdNil

`func (o *PayslipResponse) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *PayslipResponse) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
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


### GetEmployeeRemoteId

`func (o *PayslipResponse) GetEmployeeRemoteId() string`

GetEmployeeRemoteId returns the EmployeeRemoteId field if non-nil, zero value otherwise.

### GetEmployeeRemoteIdOk

`func (o *PayslipResponse) GetEmployeeRemoteIdOk() (*string, bool)`

GetEmployeeRemoteIdOk returns a tuple with the EmployeeRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeRemoteId

`func (o *PayslipResponse) SetEmployeeRemoteId(v string)`

SetEmployeeRemoteId sets EmployeeRemoteId field to given value.


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


### GetPayrunRemoteId

`func (o *PayslipResponse) GetPayrunRemoteId() string`

GetPayrunRemoteId returns the PayrunRemoteId field if non-nil, zero value otherwise.

### GetPayrunRemoteIdOk

`func (o *PayslipResponse) GetPayrunRemoteIdOk() (*string, bool)`

GetPayrunRemoteIdOk returns a tuple with the PayrunRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrunRemoteId

`func (o *PayslipResponse) SetPayrunRemoteId(v string)`

SetPayrunRemoteId sets PayrunRemoteId field to given value.


### GetPayrunType

`func (o *PayslipResponse) GetPayrunType() PayrunTypeResponse`

GetPayrunType returns the PayrunType field if non-nil, zero value otherwise.

### GetPayrunTypeOk

`func (o *PayslipResponse) GetPayrunTypeOk() (*PayrunTypeResponse, bool)`

GetPayrunTypeOk returns a tuple with the PayrunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrunType

`func (o *PayslipResponse) SetPayrunType(v PayrunTypeResponse)`

SetPayrunType sets PayrunType field to given value.


### SetPayrunTypeNil

`func (o *PayslipResponse) SetPayrunTypeNil(b bool)`

 SetPayrunTypeNil sets the value for PayrunType to be an explicit nil

### UnsetPayrunType
`func (o *PayslipResponse) UnsetPayrunType()`

UnsetPayrunType ensures that no value is present for PayrunType, not even an explicit nil
### GetCurrency

`func (o *PayslipResponse) GetCurrency() CurrencyNotNullResponse`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PayslipResponse) GetCurrencyOk() (*CurrencyNotNullResponse, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PayslipResponse) SetCurrency(v CurrencyNotNullResponse)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *PayslipResponse) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PayslipResponse) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
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


### SetGrossPayNil

`func (o *PayslipResponse) SetGrossPayNil(b bool)`

 SetGrossPayNil sets the value for GrossPay to be an explicit nil

### UnsetGrossPay
`func (o *PayslipResponse) UnsetGrossPay()`

UnsetGrossPay ensures that no value is present for GrossPay, not even an explicit nil
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


### SetNetPayNil

`func (o *PayslipResponse) SetNetPayNil(b bool)`

 SetNetPayNil sets the value for NetPay to be an explicit nil

### UnsetNetPay
`func (o *PayslipResponse) UnsetNetPay()`

UnsetNetPay ensures that no value is present for NetPay, not even an explicit nil
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


### GetPaymentDate

`func (o *PayslipResponse) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *PayslipResponse) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *PayslipResponse) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.


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


### SetEarningsNil

`func (o *PayslipResponse) SetEarningsNil(b bool)`

 SetEarningsNil sets the value for Earnings to be an explicit nil

### UnsetEarnings
`func (o *PayslipResponse) UnsetEarnings()`

UnsetEarnings ensures that no value is present for Earnings, not even an explicit nil
### GetContributions

`func (o *PayslipResponse) GetContributions() []PayslipResponseContributions`

GetContributions returns the Contributions field if non-nil, zero value otherwise.

### GetContributionsOk

`func (o *PayslipResponse) GetContributionsOk() (*[]PayslipResponseContributions, bool)`

GetContributionsOk returns a tuple with the Contributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributions

`func (o *PayslipResponse) SetContributions(v []PayslipResponseContributions)`

SetContributions sets Contributions field to given value.


### SetContributionsNil

`func (o *PayslipResponse) SetContributionsNil(b bool)`

 SetContributionsNil sets the value for Contributions to be an explicit nil

### UnsetContributions
`func (o *PayslipResponse) UnsetContributions()`

UnsetContributions ensures that no value is present for Contributions, not even an explicit nil
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


### SetDeductionsNil

`func (o *PayslipResponse) SetDeductionsNil(b bool)`

 SetDeductionsNil sets the value for Deductions to be an explicit nil

### UnsetDeductions
`func (o *PayslipResponse) UnsetDeductions()`

UnsetDeductions ensures that no value is present for Deductions, not even an explicit nil
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


### SetTaxesNil

`func (o *PayslipResponse) SetTaxesNil(b bool)`

 SetTaxesNil sets the value for Taxes to be an explicit nil

### UnsetTaxes
`func (o *PayslipResponse) UnsetTaxes()`

UnsetTaxes ensures that no value is present for Taxes, not even an explicit nil
### GetReimbursements

`func (o *PayslipResponse) GetReimbursements() []PayslipResponseReimbursements`

GetReimbursements returns the Reimbursements field if non-nil, zero value otherwise.

### GetReimbursementsOk

`func (o *PayslipResponse) GetReimbursementsOk() (*[]PayslipResponseReimbursements, bool)`

GetReimbursementsOk returns a tuple with the Reimbursements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReimbursements

`func (o *PayslipResponse) SetReimbursements(v []PayslipResponseReimbursements)`

SetReimbursements sets Reimbursements field to given value.


### SetReimbursementsNil

`func (o *PayslipResponse) SetReimbursementsNil(b bool)`

 SetReimbursementsNil sets the value for Reimbursements to be an explicit nil

### UnsetReimbursements
`func (o *PayslipResponse) UnsetReimbursements()`

UnsetReimbursements ensures that no value is present for Reimbursements, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


