# EmploymentNoNullEnumRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobTitle** | **NullableString** |  | 
**PayRate** | **NullableFloat32** |  | 
**PayPeriod** | **NullableString** |  | 
**PayFrequency** | **NullableString** |  | 
**EmploymentType** | **NullableString** |  | 
**Currency** | [**NullableCurrencyNotNullRequest**](CurrencyNotNullRequest.md) |  | 
**EffectiveDate** | **NullableString** |  | 

## Methods

### NewEmploymentNoNullEnumRequest

`func NewEmploymentNoNullEnumRequest(jobTitle NullableString, payRate NullableFloat32, payPeriod NullableString, payFrequency NullableString, employmentType NullableString, currency NullableCurrencyNotNullRequest, effectiveDate NullableString, ) *EmploymentNoNullEnumRequest`

NewEmploymentNoNullEnumRequest instantiates a new EmploymentNoNullEnumRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmploymentNoNullEnumRequestWithDefaults

`func NewEmploymentNoNullEnumRequestWithDefaults() *EmploymentNoNullEnumRequest`

NewEmploymentNoNullEnumRequestWithDefaults instantiates a new EmploymentNoNullEnumRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobTitle

`func (o *EmploymentNoNullEnumRequest) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *EmploymentNoNullEnumRequest) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *EmploymentNoNullEnumRequest) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.


### SetJobTitleNil

`func (o *EmploymentNoNullEnumRequest) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *EmploymentNoNullEnumRequest) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetPayRate

`func (o *EmploymentNoNullEnumRequest) GetPayRate() float32`

GetPayRate returns the PayRate field if non-nil, zero value otherwise.

### GetPayRateOk

`func (o *EmploymentNoNullEnumRequest) GetPayRateOk() (*float32, bool)`

GetPayRateOk returns a tuple with the PayRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayRate

`func (o *EmploymentNoNullEnumRequest) SetPayRate(v float32)`

SetPayRate sets PayRate field to given value.


### SetPayRateNil

`func (o *EmploymentNoNullEnumRequest) SetPayRateNil(b bool)`

 SetPayRateNil sets the value for PayRate to be an explicit nil

### UnsetPayRate
`func (o *EmploymentNoNullEnumRequest) UnsetPayRate()`

UnsetPayRate ensures that no value is present for PayRate, not even an explicit nil
### GetPayPeriod

`func (o *EmploymentNoNullEnumRequest) GetPayPeriod() string`

GetPayPeriod returns the PayPeriod field if non-nil, zero value otherwise.

### GetPayPeriodOk

`func (o *EmploymentNoNullEnumRequest) GetPayPeriodOk() (*string, bool)`

GetPayPeriodOk returns a tuple with the PayPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriod

`func (o *EmploymentNoNullEnumRequest) SetPayPeriod(v string)`

SetPayPeriod sets PayPeriod field to given value.


### SetPayPeriodNil

`func (o *EmploymentNoNullEnumRequest) SetPayPeriodNil(b bool)`

 SetPayPeriodNil sets the value for PayPeriod to be an explicit nil

### UnsetPayPeriod
`func (o *EmploymentNoNullEnumRequest) UnsetPayPeriod()`

UnsetPayPeriod ensures that no value is present for PayPeriod, not even an explicit nil
### GetPayFrequency

`func (o *EmploymentNoNullEnumRequest) GetPayFrequency() string`

GetPayFrequency returns the PayFrequency field if non-nil, zero value otherwise.

### GetPayFrequencyOk

`func (o *EmploymentNoNullEnumRequest) GetPayFrequencyOk() (*string, bool)`

GetPayFrequencyOk returns a tuple with the PayFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayFrequency

`func (o *EmploymentNoNullEnumRequest) SetPayFrequency(v string)`

SetPayFrequency sets PayFrequency field to given value.


### SetPayFrequencyNil

`func (o *EmploymentNoNullEnumRequest) SetPayFrequencyNil(b bool)`

 SetPayFrequencyNil sets the value for PayFrequency to be an explicit nil

### UnsetPayFrequency
`func (o *EmploymentNoNullEnumRequest) UnsetPayFrequency()`

UnsetPayFrequency ensures that no value is present for PayFrequency, not even an explicit nil
### GetEmploymentType

`func (o *EmploymentNoNullEnumRequest) GetEmploymentType() string`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *EmploymentNoNullEnumRequest) GetEmploymentTypeOk() (*string, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *EmploymentNoNullEnumRequest) SetEmploymentType(v string)`

SetEmploymentType sets EmploymentType field to given value.


### SetEmploymentTypeNil

`func (o *EmploymentNoNullEnumRequest) SetEmploymentTypeNil(b bool)`

 SetEmploymentTypeNil sets the value for EmploymentType to be an explicit nil

### UnsetEmploymentType
`func (o *EmploymentNoNullEnumRequest) UnsetEmploymentType()`

UnsetEmploymentType ensures that no value is present for EmploymentType, not even an explicit nil
### GetCurrency

`func (o *EmploymentNoNullEnumRequest) GetCurrency() CurrencyNotNullRequest`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EmploymentNoNullEnumRequest) GetCurrencyOk() (*CurrencyNotNullRequest, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EmploymentNoNullEnumRequest) SetCurrency(v CurrencyNotNullRequest)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *EmploymentNoNullEnumRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *EmploymentNoNullEnumRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetEffectiveDate

`func (o *EmploymentNoNullEnumRequest) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *EmploymentNoNullEnumRequest) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *EmploymentNoNullEnumRequest) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### SetEffectiveDateNil

`func (o *EmploymentNoNullEnumRequest) SetEffectiveDateNil(b bool)`

 SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil

### UnsetEffectiveDate
`func (o *EmploymentNoNullEnumRequest) UnsetEffectiveDate()`

UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


