# TimeOffEntryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Affix-assigned id of the time off entry | 
**RemoteId** | **string** | the remote system-assigned id of the time off entry | 
**EmployeeId** | **string** | the Affix-assigned id of the individual | 
**RemoteEmployeeId** | **string** | the remote system-assigned id of the individual | 
**StartDate** | **NullableString** |  | 
**EndDate** | **NullableString** |  | 
**Amount** | **float32** |  | 
**Unit** | **string** |  | 
**Status** | **NullableString** |  | 
**EmployeeNote** | **NullableString** |  | 
**RequestType** | **NullableString** |  | 
**RemoteCreatedAt** | **NullableString** |  | 
**RemoteModifiedAt** | **NullableString** |  | 

## Methods

### NewTimeOffEntryResponse

`func NewTimeOffEntryResponse(id string, remoteId string, employeeId string, remoteEmployeeId string, startDate NullableString, endDate NullableString, amount float32, unit string, status NullableString, employeeNote NullableString, requestType NullableString, remoteCreatedAt NullableString, remoteModifiedAt NullableString, ) *TimeOffEntryResponse`

NewTimeOffEntryResponse instantiates a new TimeOffEntryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOffEntryResponseWithDefaults

`func NewTimeOffEntryResponseWithDefaults() *TimeOffEntryResponse`

NewTimeOffEntryResponseWithDefaults instantiates a new TimeOffEntryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeOffEntryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeOffEntryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeOffEntryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetRemoteId

`func (o *TimeOffEntryResponse) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *TimeOffEntryResponse) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *TimeOffEntryResponse) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetEmployeeId

`func (o *TimeOffEntryResponse) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *TimeOffEntryResponse) GetEmployeeIdOk() (*string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId

`func (o *TimeOffEntryResponse) SetEmployeeId(v string)`

SetEmployeeId sets EmployeeId field to given value.


### GetRemoteEmployeeId

`func (o *TimeOffEntryResponse) GetRemoteEmployeeId() string`

GetRemoteEmployeeId returns the RemoteEmployeeId field if non-nil, zero value otherwise.

### GetRemoteEmployeeIdOk

`func (o *TimeOffEntryResponse) GetRemoteEmployeeIdOk() (*string, bool)`

GetRemoteEmployeeIdOk returns a tuple with the RemoteEmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEmployeeId

`func (o *TimeOffEntryResponse) SetRemoteEmployeeId(v string)`

SetRemoteEmployeeId sets RemoteEmployeeId field to given value.


### GetStartDate

`func (o *TimeOffEntryResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TimeOffEntryResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TimeOffEntryResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### SetStartDateNil

`func (o *TimeOffEntryResponse) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *TimeOffEntryResponse) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *TimeOffEntryResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TimeOffEntryResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TimeOffEntryResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *TimeOffEntryResponse) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *TimeOffEntryResponse) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetAmount

`func (o *TimeOffEntryResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TimeOffEntryResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TimeOffEntryResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *TimeOffEntryResponse) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TimeOffEntryResponse) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TimeOffEntryResponse) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetStatus

`func (o *TimeOffEntryResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TimeOffEntryResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TimeOffEntryResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *TimeOffEntryResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TimeOffEntryResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetEmployeeNote

`func (o *TimeOffEntryResponse) GetEmployeeNote() string`

GetEmployeeNote returns the EmployeeNote field if non-nil, zero value otherwise.

### GetEmployeeNoteOk

`func (o *TimeOffEntryResponse) GetEmployeeNoteOk() (*string, bool)`

GetEmployeeNoteOk returns a tuple with the EmployeeNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNote

`func (o *TimeOffEntryResponse) SetEmployeeNote(v string)`

SetEmployeeNote sets EmployeeNote field to given value.


### SetEmployeeNoteNil

`func (o *TimeOffEntryResponse) SetEmployeeNoteNil(b bool)`

 SetEmployeeNoteNil sets the value for EmployeeNote to be an explicit nil

### UnsetEmployeeNote
`func (o *TimeOffEntryResponse) UnsetEmployeeNote()`

UnsetEmployeeNote ensures that no value is present for EmployeeNote, not even an explicit nil
### GetRequestType

`func (o *TimeOffEntryResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *TimeOffEntryResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *TimeOffEntryResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### SetRequestTypeNil

`func (o *TimeOffEntryResponse) SetRequestTypeNil(b bool)`

 SetRequestTypeNil sets the value for RequestType to be an explicit nil

### UnsetRequestType
`func (o *TimeOffEntryResponse) UnsetRequestType()`

UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
### GetRemoteCreatedAt

`func (o *TimeOffEntryResponse) GetRemoteCreatedAt() string`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *TimeOffEntryResponse) GetRemoteCreatedAtOk() (*string, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *TimeOffEntryResponse) SetRemoteCreatedAt(v string)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.


### SetRemoteCreatedAtNil

`func (o *TimeOffEntryResponse) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *TimeOffEntryResponse) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetRemoteModifiedAt

`func (o *TimeOffEntryResponse) GetRemoteModifiedAt() string`

GetRemoteModifiedAt returns the RemoteModifiedAt field if non-nil, zero value otherwise.

### GetRemoteModifiedAtOk

`func (o *TimeOffEntryResponse) GetRemoteModifiedAtOk() (*string, bool)`

GetRemoteModifiedAtOk returns a tuple with the RemoteModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteModifiedAt

`func (o *TimeOffEntryResponse) SetRemoteModifiedAt(v string)`

SetRemoteModifiedAt sets RemoteModifiedAt field to given value.


### SetRemoteModifiedAtNil

`func (o *TimeOffEntryResponse) SetRemoteModifiedAtNil(b bool)`

 SetRemoteModifiedAtNil sets the value for RemoteModifiedAt to be an explicit nil

### UnsetRemoteModifiedAt
`func (o *TimeOffEntryResponse) UnsetRemoteModifiedAt()`

UnsetRemoteModifiedAt ensures that no value is present for RemoteModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


