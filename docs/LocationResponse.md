# LocationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreetAddress** | **NullableString** |  | 
**Locality** | **NullableString** |  | 
**AdministrativeArea** | **NullableString** | The administrative area of the address. If US or CA, the two-letter state abbreviation  | 
**Country** | **NullableString** | The two-letter abbreviation of the country  | 
**PostCode** | **NullableString** |  | 

## Methods

### NewLocationResponse

`func NewLocationResponse(streetAddress NullableString, locality NullableString, administrativeArea NullableString, country NullableString, postCode NullableString, ) *LocationResponse`

NewLocationResponse instantiates a new LocationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationResponseWithDefaults

`func NewLocationResponseWithDefaults() *LocationResponse`

NewLocationResponseWithDefaults instantiates a new LocationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreetAddress

`func (o *LocationResponse) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *LocationResponse) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *LocationResponse) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.


### SetStreetAddressNil

`func (o *LocationResponse) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *LocationResponse) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetLocality

`func (o *LocationResponse) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *LocationResponse) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *LocationResponse) SetLocality(v string)`

SetLocality sets Locality field to given value.


### SetLocalityNil

`func (o *LocationResponse) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *LocationResponse) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil
### GetAdministrativeArea

`func (o *LocationResponse) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *LocationResponse) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *LocationResponse) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.


### SetAdministrativeAreaNil

`func (o *LocationResponse) SetAdministrativeAreaNil(b bool)`

 SetAdministrativeAreaNil sets the value for AdministrativeArea to be an explicit nil

### UnsetAdministrativeArea
`func (o *LocationResponse) UnsetAdministrativeArea()`

UnsetAdministrativeArea ensures that no value is present for AdministrativeArea, not even an explicit nil
### GetCountry

`func (o *LocationResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LocationResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LocationResponse) SetCountry(v string)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *LocationResponse) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *LocationResponse) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPostCode

`func (o *LocationResponse) GetPostCode() string`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *LocationResponse) GetPostCodeOk() (*string, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *LocationResponse) SetPostCode(v string)`

SetPostCode sets PostCode field to given value.


### SetPostCodeNil

`func (o *LocationResponse) SetPostCodeNil(b bool)`

 SetPostCodeNil sets the value for PostCode to be an explicit nil

### UnsetPostCode
`func (o *LocationResponse) UnsetPostCode()`

UnsetPostCode ensures that no value is present for PostCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


