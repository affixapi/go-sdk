# PayslipResponseDeductions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayrunId** | **string** |  | 
**Name** | **string** |  | 
**Amount** | **float32** | if USD/EUR/GBP, in cent | 

## Methods

### NewPayslipResponseDeductions

`func NewPayslipResponseDeductions(payrunId string, name string, amount float32, ) *PayslipResponseDeductions`

NewPayslipResponseDeductions instantiates a new PayslipResponseDeductions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayslipResponseDeductionsWithDefaults

`func NewPayslipResponseDeductionsWithDefaults() *PayslipResponseDeductions`

NewPayslipResponseDeductionsWithDefaults instantiates a new PayslipResponseDeductions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayrunId

`func (o *PayslipResponseDeductions) GetPayrunId() string`

GetPayrunId returns the PayrunId field if non-nil, zero value otherwise.

### GetPayrunIdOk

`func (o *PayslipResponseDeductions) GetPayrunIdOk() (*string, bool)`

GetPayrunIdOk returns a tuple with the PayrunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrunId

`func (o *PayslipResponseDeductions) SetPayrunId(v string)`

SetPayrunId sets PayrunId field to given value.


### GetName

`func (o *PayslipResponseDeductions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PayslipResponseDeductions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PayslipResponseDeductions) SetName(v string)`

SetName sets Name field to given value.


### GetAmount

`func (o *PayslipResponseDeductions) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayslipResponseDeductions) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayslipResponseDeductions) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


