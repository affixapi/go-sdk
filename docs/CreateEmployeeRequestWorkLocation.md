# CreateEmployeeRequestWorkLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**LocationRequest**](LocationRequest.md) |  | 
**Type** | **NullableString** |  | 
**Name** | **NullableString** |  | 

## Methods

### NewCreateEmployeeRequestWorkLocation

`func NewCreateEmployeeRequestWorkLocation(location LocationRequest, type_ NullableString, name NullableString, ) *CreateEmployeeRequestWorkLocation`

NewCreateEmployeeRequestWorkLocation instantiates a new CreateEmployeeRequestWorkLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmployeeRequestWorkLocationWithDefaults

`func NewCreateEmployeeRequestWorkLocationWithDefaults() *CreateEmployeeRequestWorkLocation`

NewCreateEmployeeRequestWorkLocationWithDefaults instantiates a new CreateEmployeeRequestWorkLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *CreateEmployeeRequestWorkLocation) GetLocation() LocationRequest`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateEmployeeRequestWorkLocation) GetLocationOk() (*LocationRequest, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateEmployeeRequestWorkLocation) SetLocation(v LocationRequest)`

SetLocation sets Location field to given value.


### GetType

`func (o *CreateEmployeeRequestWorkLocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateEmployeeRequestWorkLocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateEmployeeRequestWorkLocation) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CreateEmployeeRequestWorkLocation) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateEmployeeRequestWorkLocation) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *CreateEmployeeRequestWorkLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEmployeeRequestWorkLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEmployeeRequestWorkLocation) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateEmployeeRequestWorkLocation) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateEmployeeRequestWorkLocation) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


