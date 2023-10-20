# EmploymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobTitle** | **NullableString** |  | 
**PayRate** | **NullableFloat32** |  | 
**PayPeriod** | **NullableString** |  | 
**PayFrequency** | **NullableString** |  | 
**EmploymentType** | **NullableString** |  | 
**Currency** | [**CurrencyResponse**](CurrencyResponse.md) |  | 
**EffectiveDate** | **NullableString** |  | 

## Methods

### NewEmploymentResponse

`func NewEmploymentResponse(jobTitle NullableString, payRate NullableFloat32, payPeriod NullableString, payFrequency NullableString, employmentType NullableString, currency CurrencyResponse, effectiveDate NullableString, ) *EmploymentResponse`

NewEmploymentResponse instantiates a new EmploymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmploymentResponseWithDefaults

`func NewEmploymentResponseWithDefaults() *EmploymentResponse`

NewEmploymentResponseWithDefaults instantiates a new EmploymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobTitle

`func (o *EmploymentResponse) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *EmploymentResponse) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *EmploymentResponse) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.


### SetJobTitleNil

`func (o *EmploymentResponse) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *EmploymentResponse) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetPayRate

`func (o *EmploymentResponse) GetPayRate() float32`

GetPayRate returns the PayRate field if non-nil, zero value otherwise.

### GetPayRateOk

`func (o *EmploymentResponse) GetPayRateOk() (*float32, bool)`

GetPayRateOk returns a tuple with the PayRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayRate

`func (o *EmploymentResponse) SetPayRate(v float32)`

SetPayRate sets PayRate field to given value.


### SetPayRateNil

`func (o *EmploymentResponse) SetPayRateNil(b bool)`

 SetPayRateNil sets the value for PayRate to be an explicit nil

### UnsetPayRate
`func (o *EmploymentResponse) UnsetPayRate()`

UnsetPayRate ensures that no value is present for PayRate, not even an explicit nil
### GetPayPeriod

`func (o *EmploymentResponse) GetPayPeriod() string`

GetPayPeriod returns the PayPeriod field if non-nil, zero value otherwise.

### GetPayPeriodOk

`func (o *EmploymentResponse) GetPayPeriodOk() (*string, bool)`

GetPayPeriodOk returns a tuple with the PayPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriod

`func (o *EmploymentResponse) SetPayPeriod(v string)`

SetPayPeriod sets PayPeriod field to given value.


### SetPayPeriodNil

`func (o *EmploymentResponse) SetPayPeriodNil(b bool)`

 SetPayPeriodNil sets the value for PayPeriod to be an explicit nil

### UnsetPayPeriod
`func (o *EmploymentResponse) UnsetPayPeriod()`

UnsetPayPeriod ensures that no value is present for PayPeriod, not even an explicit nil
### GetPayFrequency

`func (o *EmploymentResponse) GetPayFrequency() string`

GetPayFrequency returns the PayFrequency field if non-nil, zero value otherwise.

### GetPayFrequencyOk

`func (o *EmploymentResponse) GetPayFrequencyOk() (*string, bool)`

GetPayFrequencyOk returns a tuple with the PayFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayFrequency

`func (o *EmploymentResponse) SetPayFrequency(v string)`

SetPayFrequency sets PayFrequency field to given value.


### SetPayFrequencyNil

`func (o *EmploymentResponse) SetPayFrequencyNil(b bool)`

 SetPayFrequencyNil sets the value for PayFrequency to be an explicit nil

### UnsetPayFrequency
`func (o *EmploymentResponse) UnsetPayFrequency()`

UnsetPayFrequency ensures that no value is present for PayFrequency, not even an explicit nil
### GetEmploymentType

`func (o *EmploymentResponse) GetEmploymentType() string`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *EmploymentResponse) GetEmploymentTypeOk() (*string, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *EmploymentResponse) SetEmploymentType(v string)`

SetEmploymentType sets EmploymentType field to given value.


### SetEmploymentTypeNil

`func (o *EmploymentResponse) SetEmploymentTypeNil(b bool)`

 SetEmploymentTypeNil sets the value for EmploymentType to be an explicit nil

### UnsetEmploymentType
`func (o *EmploymentResponse) UnsetEmploymentType()`

UnsetEmploymentType ensures that no value is present for EmploymentType, not even an explicit nil
### GetCurrency

`func (o *EmploymentResponse) GetCurrency() CurrencyResponse`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EmploymentResponse) GetCurrencyOk() (*CurrencyResponse, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EmploymentResponse) SetCurrency(v CurrencyResponse)`

SetCurrency sets Currency field to given value.


### GetEffectiveDate

`func (o *EmploymentResponse) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *EmploymentResponse) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *EmploymentResponse) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### SetEffectiveDateNil

`func (o *EmploymentResponse) SetEffectiveDateNil(b bool)`

 SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil

### UnsetEffectiveDate
`func (o *EmploymentResponse) UnsetEffectiveDate()`

UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


