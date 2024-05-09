# CompensationHistoryNoNonNullRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayRate** | **float32** |  | 
**PayPeriod** | **string** |  | 
**PayFrequency** | **string** |  | 
**EmploymentType** | **string** |  | 
**Currency** | [**NullableCurrencyRequest**](CurrencyRequest.md) |  | 
**EffectiveDate** | **string** |  | 

## Methods

### NewCompensationHistoryNoNonNullRequest

`func NewCompensationHistoryNoNonNullRequest(payRate float32, payPeriod string, payFrequency string, employmentType string, currency NullableCurrencyRequest, effectiveDate string, ) *CompensationHistoryNoNonNullRequest`

NewCompensationHistoryNoNonNullRequest instantiates a new CompensationHistoryNoNonNullRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompensationHistoryNoNonNullRequestWithDefaults

`func NewCompensationHistoryNoNonNullRequestWithDefaults() *CompensationHistoryNoNonNullRequest`

NewCompensationHistoryNoNonNullRequestWithDefaults instantiates a new CompensationHistoryNoNonNullRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayRate

`func (o *CompensationHistoryNoNonNullRequest) GetPayRate() float32`

GetPayRate returns the PayRate field if non-nil, zero value otherwise.

### GetPayRateOk

`func (o *CompensationHistoryNoNonNullRequest) GetPayRateOk() (*float32, bool)`

GetPayRateOk returns a tuple with the PayRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayRate

`func (o *CompensationHistoryNoNonNullRequest) SetPayRate(v float32)`

SetPayRate sets PayRate field to given value.


### GetPayPeriod

`func (o *CompensationHistoryNoNonNullRequest) GetPayPeriod() string`

GetPayPeriod returns the PayPeriod field if non-nil, zero value otherwise.

### GetPayPeriodOk

`func (o *CompensationHistoryNoNonNullRequest) GetPayPeriodOk() (*string, bool)`

GetPayPeriodOk returns a tuple with the PayPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriod

`func (o *CompensationHistoryNoNonNullRequest) SetPayPeriod(v string)`

SetPayPeriod sets PayPeriod field to given value.


### GetPayFrequency

`func (o *CompensationHistoryNoNonNullRequest) GetPayFrequency() string`

GetPayFrequency returns the PayFrequency field if non-nil, zero value otherwise.

### GetPayFrequencyOk

`func (o *CompensationHistoryNoNonNullRequest) GetPayFrequencyOk() (*string, bool)`

GetPayFrequencyOk returns a tuple with the PayFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayFrequency

`func (o *CompensationHistoryNoNonNullRequest) SetPayFrequency(v string)`

SetPayFrequency sets PayFrequency field to given value.


### GetEmploymentType

`func (o *CompensationHistoryNoNonNullRequest) GetEmploymentType() string`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *CompensationHistoryNoNonNullRequest) GetEmploymentTypeOk() (*string, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *CompensationHistoryNoNonNullRequest) SetEmploymentType(v string)`

SetEmploymentType sets EmploymentType field to given value.


### GetCurrency

`func (o *CompensationHistoryNoNonNullRequest) GetCurrency() CurrencyRequest`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CompensationHistoryNoNonNullRequest) GetCurrencyOk() (*CurrencyRequest, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CompensationHistoryNoNonNullRequest) SetCurrency(v CurrencyRequest)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *CompensationHistoryNoNonNullRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CompensationHistoryNoNonNullRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetEffectiveDate

`func (o *CompensationHistoryNoNonNullRequest) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *CompensationHistoryNoNonNullRequest) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *CompensationHistoryNoNonNullRequest) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


