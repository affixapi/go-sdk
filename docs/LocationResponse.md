# LocationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The Affix-assigned id of the individual | 
**RemoteId** | **NullableString** | the remote system-assigned id of the individual | 
**Name** | **NullableString** | System assigned description of the location | 
**Type** | **NullableString** | The location&#39;s type. In cases where there is no clear mapping, the original value passed through will be returned.  | 
**Address** | [**NullableAddressResponse**](AddressResponse.md) |  | 

## Methods

### NewLocationResponse

`func NewLocationResponse(id NullableString, remoteId NullableString, name NullableString, type_ NullableString, address NullableAddressResponse, ) *LocationResponse`

NewLocationResponse instantiates a new LocationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationResponseWithDefaults

`func NewLocationResponseWithDefaults() *LocationResponse`

NewLocationResponseWithDefaults instantiates a new LocationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationResponse) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *LocationResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LocationResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRemoteId

`func (o *LocationResponse) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *LocationResponse) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *LocationResponse) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### SetRemoteIdNil

`func (o *LocationResponse) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *LocationResponse) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetName

`func (o *LocationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationResponse) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *LocationResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LocationResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *LocationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocationResponse) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *LocationResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LocationResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAddress

`func (o *LocationResponse) GetAddress() AddressResponse`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LocationResponse) GetAddressOk() (*AddressResponse, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LocationResponse) SetAddress(v AddressResponse)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *LocationResponse) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *LocationResponse) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


