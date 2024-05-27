# PayrunResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Affix-assigned id of the individual | 
**RemoteId** | **string** | the remote system-assigned id of the payrun | 
**State** | **NullableString** |  | 
**Type** | [**NullablePayrunTypeResponse**](PayrunTypeResponse.md) |  | 
**StartDate** | **NullableString** | Payrun period start date | 
**EndDate** | **NullableString** | Payrun period end date | 
**PaymentDate** | **NullableString** | Payment date / check date | 

## Methods

### NewPayrunResponse

`func NewPayrunResponse(id string, remoteId string, state NullableString, type_ NullablePayrunTypeResponse, startDate NullableString, endDate NullableString, paymentDate NullableString, ) *PayrunResponse`

NewPayrunResponse instantiates a new PayrunResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayrunResponseWithDefaults

`func NewPayrunResponseWithDefaults() *PayrunResponse`

NewPayrunResponseWithDefaults instantiates a new PayrunResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PayrunResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayrunResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayrunResponse) SetId(v string)`

SetId sets Id field to given value.


### GetRemoteId

`func (o *PayrunResponse) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayrunResponse) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayrunResponse) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetState

`func (o *PayrunResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PayrunResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PayrunResponse) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *PayrunResponse) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *PayrunResponse) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetType

`func (o *PayrunResponse) GetType() PayrunTypeResponse`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PayrunResponse) GetTypeOk() (*PayrunTypeResponse, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PayrunResponse) SetType(v PayrunTypeResponse)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PayrunResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PayrunResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStartDate

`func (o *PayrunResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PayrunResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PayrunResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### SetStartDateNil

`func (o *PayrunResponse) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *PayrunResponse) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *PayrunResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PayrunResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PayrunResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *PayrunResponse) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *PayrunResponse) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetPaymentDate

`func (o *PayrunResponse) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *PayrunResponse) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *PayrunResponse) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.


### SetPaymentDateNil

`func (o *PayrunResponse) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *PayrunResponse) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


