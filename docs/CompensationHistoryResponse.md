# CompensationHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayRate** | **NullableFloat32** |  | 
**PayPeriod** | **NullableString** |  | 
**PayFrequency** | **NullableString** |  | 
**EmploymentType** | **NullableString** |  | 
**Currency** | [**NullableCurrencyResponse**](CurrencyResponse.md) |  | 
**EffectiveDate** | **NullableString** |  | 
**Notes** | **NullableString** |  | 

## Methods

### NewCompensationHistoryResponse

`func NewCompensationHistoryResponse(payRate NullableFloat32, payPeriod NullableString, payFrequency NullableString, employmentType NullableString, currency NullableCurrencyResponse, effectiveDate NullableString, notes NullableString, ) *CompensationHistoryResponse`

NewCompensationHistoryResponse instantiates a new CompensationHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompensationHistoryResponseWithDefaults

`func NewCompensationHistoryResponseWithDefaults() *CompensationHistoryResponse`

NewCompensationHistoryResponseWithDefaults instantiates a new CompensationHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayRate

`func (o *CompensationHistoryResponse) GetPayRate() float32`

GetPayRate returns the PayRate field if non-nil, zero value otherwise.

### GetPayRateOk

`func (o *CompensationHistoryResponse) GetPayRateOk() (*float32, bool)`

GetPayRateOk returns a tuple with the PayRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayRate

`func (o *CompensationHistoryResponse) SetPayRate(v float32)`

SetPayRate sets PayRate field to given value.


### SetPayRateNil

`func (o *CompensationHistoryResponse) SetPayRateNil(b bool)`

 SetPayRateNil sets the value for PayRate to be an explicit nil

### UnsetPayRate
`func (o *CompensationHistoryResponse) UnsetPayRate()`

UnsetPayRate ensures that no value is present for PayRate, not even an explicit nil
### GetPayPeriod

`func (o *CompensationHistoryResponse) GetPayPeriod() string`

GetPayPeriod returns the PayPeriod field if non-nil, zero value otherwise.

### GetPayPeriodOk

`func (o *CompensationHistoryResponse) GetPayPeriodOk() (*string, bool)`

GetPayPeriodOk returns a tuple with the PayPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriod

`func (o *CompensationHistoryResponse) SetPayPeriod(v string)`

SetPayPeriod sets PayPeriod field to given value.


### SetPayPeriodNil

`func (o *CompensationHistoryResponse) SetPayPeriodNil(b bool)`

 SetPayPeriodNil sets the value for PayPeriod to be an explicit nil

### UnsetPayPeriod
`func (o *CompensationHistoryResponse) UnsetPayPeriod()`

UnsetPayPeriod ensures that no value is present for PayPeriod, not even an explicit nil
### GetPayFrequency

`func (o *CompensationHistoryResponse) GetPayFrequency() string`

GetPayFrequency returns the PayFrequency field if non-nil, zero value otherwise.

### GetPayFrequencyOk

`func (o *CompensationHistoryResponse) GetPayFrequencyOk() (*string, bool)`

GetPayFrequencyOk returns a tuple with the PayFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayFrequency

`func (o *CompensationHistoryResponse) SetPayFrequency(v string)`

SetPayFrequency sets PayFrequency field to given value.


### SetPayFrequencyNil

`func (o *CompensationHistoryResponse) SetPayFrequencyNil(b bool)`

 SetPayFrequencyNil sets the value for PayFrequency to be an explicit nil

### UnsetPayFrequency
`func (o *CompensationHistoryResponse) UnsetPayFrequency()`

UnsetPayFrequency ensures that no value is present for PayFrequency, not even an explicit nil
### GetEmploymentType

`func (o *CompensationHistoryResponse) GetEmploymentType() string`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *CompensationHistoryResponse) GetEmploymentTypeOk() (*string, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *CompensationHistoryResponse) SetEmploymentType(v string)`

SetEmploymentType sets EmploymentType field to given value.


### SetEmploymentTypeNil

`func (o *CompensationHistoryResponse) SetEmploymentTypeNil(b bool)`

 SetEmploymentTypeNil sets the value for EmploymentType to be an explicit nil

### UnsetEmploymentType
`func (o *CompensationHistoryResponse) UnsetEmploymentType()`

UnsetEmploymentType ensures that no value is present for EmploymentType, not even an explicit nil
### GetCurrency

`func (o *CompensationHistoryResponse) GetCurrency() CurrencyResponse`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CompensationHistoryResponse) GetCurrencyOk() (*CurrencyResponse, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CompensationHistoryResponse) SetCurrency(v CurrencyResponse)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *CompensationHistoryResponse) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CompensationHistoryResponse) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetEffectiveDate

`func (o *CompensationHistoryResponse) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *CompensationHistoryResponse) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *CompensationHistoryResponse) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### SetEffectiveDateNil

`func (o *CompensationHistoryResponse) SetEffectiveDateNil(b bool)`

 SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil

### UnsetEffectiveDate
`func (o *CompensationHistoryResponse) UnsetEffectiveDate()`

UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil
### GetNotes

`func (o *CompensationHistoryResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CompensationHistoryResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CompensationHistoryResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.


### SetNotesNil

`func (o *CompensationHistoryResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *CompensationHistoryResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


