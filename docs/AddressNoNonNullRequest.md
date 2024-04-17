# AddressNoNonNullRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreetAddress** | **NullableString** |  | 
**Locality** | **NullableString** |  | 
**AdministrativeArea** | **NullableString** | The administrative area of the address. If US or CA, the two-letter state or province abbreviation. Else, the province / administrative area; such as, &#x60;Dublin 2&#x60; or &#x60;County Cork&#x60;  | 
**Country** | **string** | The ISO-3166-2 two-letter abbreviation of the country. Reference https://en.wikipedia.org/wiki/ISO_3166-2 for more details  | 
**PostCode** | **NullableString** |  | 

## Methods

### NewAddressNoNonNullRequest

`func NewAddressNoNonNullRequest(streetAddress NullableString, locality NullableString, administrativeArea NullableString, country string, postCode NullableString, ) *AddressNoNonNullRequest`

NewAddressNoNonNullRequest instantiates a new AddressNoNonNullRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressNoNonNullRequestWithDefaults

`func NewAddressNoNonNullRequestWithDefaults() *AddressNoNonNullRequest`

NewAddressNoNonNullRequestWithDefaults instantiates a new AddressNoNonNullRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreetAddress

`func (o *AddressNoNonNullRequest) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *AddressNoNonNullRequest) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *AddressNoNonNullRequest) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.


### SetStreetAddressNil

`func (o *AddressNoNonNullRequest) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *AddressNoNonNullRequest) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetLocality

`func (o *AddressNoNonNullRequest) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *AddressNoNonNullRequest) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *AddressNoNonNullRequest) SetLocality(v string)`

SetLocality sets Locality field to given value.


### SetLocalityNil

`func (o *AddressNoNonNullRequest) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *AddressNoNonNullRequest) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil
### GetAdministrativeArea

`func (o *AddressNoNonNullRequest) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *AddressNoNonNullRequest) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *AddressNoNonNullRequest) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.


### SetAdministrativeAreaNil

`func (o *AddressNoNonNullRequest) SetAdministrativeAreaNil(b bool)`

 SetAdministrativeAreaNil sets the value for AdministrativeArea to be an explicit nil

### UnsetAdministrativeArea
`func (o *AddressNoNonNullRequest) UnsetAdministrativeArea()`

UnsetAdministrativeArea ensures that no value is present for AdministrativeArea, not even an explicit nil
### GetCountry

`func (o *AddressNoNonNullRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressNoNonNullRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressNoNonNullRequest) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetPostCode

`func (o *AddressNoNonNullRequest) GetPostCode() string`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *AddressNoNonNullRequest) GetPostCodeOk() (*string, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *AddressNoNonNullRequest) SetPostCode(v string)`

SetPostCode sets PostCode field to given value.


### SetPostCodeNil

`func (o *AddressNoNonNullRequest) SetPostCodeNil(b bool)`

 SetPostCodeNil sets the value for PostCode to be an explicit nil

### UnsetPostCode
`func (o *AddressNoNonNullRequest) UnsetPostCode()`

UnsetPostCode ensures that no value is present for PostCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


