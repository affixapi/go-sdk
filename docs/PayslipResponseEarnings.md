# PayslipResponseEarnings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Amount** | **float32** | if USD/EUR/GBP, in cent | 
**Hours** | **NullableFloat32** | hours, if applicable | 

## Methods

### NewPayslipResponseEarnings

`func NewPayslipResponseEarnings(name string, amount float32, hours NullableFloat32, ) *PayslipResponseEarnings`

NewPayslipResponseEarnings instantiates a new PayslipResponseEarnings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayslipResponseEarningsWithDefaults

`func NewPayslipResponseEarningsWithDefaults() *PayslipResponseEarnings`

NewPayslipResponseEarningsWithDefaults instantiates a new PayslipResponseEarnings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PayslipResponseEarnings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PayslipResponseEarnings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PayslipResponseEarnings) SetName(v string)`

SetName sets Name field to given value.


### GetAmount

`func (o *PayslipResponseEarnings) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayslipResponseEarnings) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayslipResponseEarnings) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetHours

`func (o *PayslipResponseEarnings) GetHours() float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *PayslipResponseEarnings) GetHoursOk() (*float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *PayslipResponseEarnings) SetHours(v float32)`

SetHours sets Hours field to given value.


### SetHoursNil

`func (o *PayslipResponseEarnings) SetHoursNil(b bool)`

 SetHoursNil sets the value for Hours to be an explicit nil

### UnsetHours
`func (o *PayslipResponseEarnings) UnsetHours()`

UnsetHours ensures that no value is present for Hours, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


