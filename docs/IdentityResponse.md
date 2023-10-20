# IdentityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the individual for the respective account, if known | 
**Email** | **string** | The email of the individual for the respective account, if known | 
**PhoneNumber** | **NullableString** | The phone number of the individual for the respective account, if known. Nullable for tokens created prior to 2023-03-05  | 

## Methods

### NewIdentityResponse

`func NewIdentityResponse(name string, email string, phoneNumber NullableString, ) *IdentityResponse`

NewIdentityResponse instantiates a new IdentityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityResponseWithDefaults

`func NewIdentityResponseWithDefaults() *IdentityResponse`

NewIdentityResponseWithDefaults instantiates a new IdentityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdentityResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityResponse) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *IdentityResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentityResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentityResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhoneNumber

`func (o *IdentityResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IdentityResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IdentityResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### SetPhoneNumberNil

`func (o *IdentityResponse) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *IdentityResponse) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


