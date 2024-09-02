# PayslipResponseReimbursements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Amount** | **float32** | if USD/EUR/GBP, in cent | 

## Methods

### NewPayslipResponseReimbursements

`func NewPayslipResponseReimbursements(name string, amount float32, ) *PayslipResponseReimbursements`

NewPayslipResponseReimbursements instantiates a new PayslipResponseReimbursements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayslipResponseReimbursementsWithDefaults

`func NewPayslipResponseReimbursementsWithDefaults() *PayslipResponseReimbursements`

NewPayslipResponseReimbursementsWithDefaults instantiates a new PayslipResponseReimbursements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PayslipResponseReimbursements) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PayslipResponseReimbursements) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PayslipResponseReimbursements) SetName(v string)`

SetName sets Name field to given value.


### GetAmount

`func (o *PayslipResponseReimbursements) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayslipResponseReimbursements) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayslipResponseReimbursements) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


