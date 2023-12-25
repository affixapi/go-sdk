# CreateEmployeeRequestCompany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Location** | [**NullableCreateEmployeeRequestCompanyLocation**](CreateEmployeeRequestCompanyLocation.md) |  | 

## Methods

### NewCreateEmployeeRequestCompany

`func NewCreateEmployeeRequestCompany(name NullableString, location NullableCreateEmployeeRequestCompanyLocation, ) *CreateEmployeeRequestCompany`

NewCreateEmployeeRequestCompany instantiates a new CreateEmployeeRequestCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmployeeRequestCompanyWithDefaults

`func NewCreateEmployeeRequestCompanyWithDefaults() *CreateEmployeeRequestCompany`

NewCreateEmployeeRequestCompanyWithDefaults instantiates a new CreateEmployeeRequestCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEmployeeRequestCompany) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEmployeeRequestCompany) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEmployeeRequestCompany) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateEmployeeRequestCompany) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateEmployeeRequestCompany) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLocation

`func (o *CreateEmployeeRequestCompany) GetLocation() CreateEmployeeRequestCompanyLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateEmployeeRequestCompany) GetLocationOk() (*CreateEmployeeRequestCompanyLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateEmployeeRequestCompany) SetLocation(v CreateEmployeeRequestCompanyLocation)`

SetLocation sets Location field to given value.


### SetLocationNil

`func (o *CreateEmployeeRequestCompany) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *CreateEmployeeRequestCompany) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


