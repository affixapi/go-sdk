# PayslipResponseTaxes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Amount** | **float32** | if USD/EUR/GBP, in cent | 
**EmployerTax** | **NullableBool** |  | 

## Methods

### NewPayslipResponseTaxes

`func NewPayslipResponseTaxes(name string, amount float32, employerTax NullableBool, ) *PayslipResponseTaxes`

NewPayslipResponseTaxes instantiates a new PayslipResponseTaxes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayslipResponseTaxesWithDefaults

`func NewPayslipResponseTaxesWithDefaults() *PayslipResponseTaxes`

NewPayslipResponseTaxesWithDefaults instantiates a new PayslipResponseTaxes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PayslipResponseTaxes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PayslipResponseTaxes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PayslipResponseTaxes) SetName(v string)`

SetName sets Name field to given value.


### GetAmount

`func (o *PayslipResponseTaxes) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayslipResponseTaxes) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayslipResponseTaxes) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetEmployerTax

`func (o *PayslipResponseTaxes) GetEmployerTax() bool`

GetEmployerTax returns the EmployerTax field if non-nil, zero value otherwise.

### GetEmployerTaxOk

`func (o *PayslipResponseTaxes) GetEmployerTaxOk() (*bool, bool)`

GetEmployerTaxOk returns a tuple with the EmployerTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerTax

`func (o *PayslipResponseTaxes) SetEmployerTax(v bool)`

SetEmployerTax sets EmployerTax field to given value.


### SetEmployerTaxNil

`func (o *PayslipResponseTaxes) SetEmployerTaxNil(b bool)`

 SetEmployerTaxNil sets the value for EmployerTax to be an explicit nil

### UnsetEmployerTax
`func (o *PayslipResponseTaxes) UnsetEmployerTax()`

UnsetEmployerTax ensures that no value is present for EmployerTax, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


