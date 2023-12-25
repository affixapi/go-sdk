# PayslipResponseContributions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Amount** | **float32** | if USD/EUR/GBP, in cent | 

## Methods

### NewPayslipResponseContributions

`func NewPayslipResponseContributions(name string, amount float32, ) *PayslipResponseContributions`

NewPayslipResponseContributions instantiates a new PayslipResponseContributions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayslipResponseContributionsWithDefaults

`func NewPayslipResponseContributionsWithDefaults() *PayslipResponseContributions`

NewPayslipResponseContributionsWithDefaults instantiates a new PayslipResponseContributions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PayslipResponseContributions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PayslipResponseContributions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PayslipResponseContributions) SetName(v string)`

SetName sets Name field to given value.


### GetAmount

`func (o *PayslipResponseContributions) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayslipResponseContributions) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayslipResponseContributions) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


