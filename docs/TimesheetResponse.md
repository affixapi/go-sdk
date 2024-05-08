# TimesheetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Affix-assigned id of the time off entry | 
**RemoteId** | **string** | the remote system-assigned id of the time off entry | 
**EmployeeId** | **string** | the Affix-assigned id of the individual | 
**EmployeeRemoteId** | **string** | the remote system-assigned id of the individual | 
**StartTime** | **NullableTime** |  | 
**EndTime** | **NullableTime** |  | 
**HoursWorked** | **float32** |  | 
**RemoteCreatedAt** | **NullableString** |  | 
**RemoteModifiedAt** | **NullableString** |  | 

## Methods

### NewTimesheetResponse

`func NewTimesheetResponse(id string, remoteId string, employeeId string, employeeRemoteId string, startTime NullableTime, endTime NullableTime, hoursWorked float32, remoteCreatedAt NullableString, remoteModifiedAt NullableString, ) *TimesheetResponse`

NewTimesheetResponse instantiates a new TimesheetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimesheetResponseWithDefaults

`func NewTimesheetResponseWithDefaults() *TimesheetResponse`

NewTimesheetResponseWithDefaults instantiates a new TimesheetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimesheetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimesheetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimesheetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetRemoteId

`func (o *TimesheetResponse) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *TimesheetResponse) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *TimesheetResponse) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetEmployeeId

`func (o *TimesheetResponse) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *TimesheetResponse) GetEmployeeIdOk() (*string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId

`func (o *TimesheetResponse) SetEmployeeId(v string)`

SetEmployeeId sets EmployeeId field to given value.


### GetEmployeeRemoteId

`func (o *TimesheetResponse) GetEmployeeRemoteId() string`

GetEmployeeRemoteId returns the EmployeeRemoteId field if non-nil, zero value otherwise.

### GetEmployeeRemoteIdOk

`func (o *TimesheetResponse) GetEmployeeRemoteIdOk() (*string, bool)`

GetEmployeeRemoteIdOk returns a tuple with the EmployeeRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeRemoteId

`func (o *TimesheetResponse) SetEmployeeRemoteId(v string)`

SetEmployeeRemoteId sets EmployeeRemoteId field to given value.


### GetStartTime

`func (o *TimesheetResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TimesheetResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TimesheetResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### SetStartTimeNil

`func (o *TimesheetResponse) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *TimesheetResponse) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *TimesheetResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TimesheetResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TimesheetResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### SetEndTimeNil

`func (o *TimesheetResponse) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *TimesheetResponse) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetHoursWorked

`func (o *TimesheetResponse) GetHoursWorked() float32`

GetHoursWorked returns the HoursWorked field if non-nil, zero value otherwise.

### GetHoursWorkedOk

`func (o *TimesheetResponse) GetHoursWorkedOk() (*float32, bool)`

GetHoursWorkedOk returns a tuple with the HoursWorked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursWorked

`func (o *TimesheetResponse) SetHoursWorked(v float32)`

SetHoursWorked sets HoursWorked field to given value.


### GetRemoteCreatedAt

`func (o *TimesheetResponse) GetRemoteCreatedAt() string`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *TimesheetResponse) GetRemoteCreatedAtOk() (*string, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *TimesheetResponse) SetRemoteCreatedAt(v string)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.


### SetRemoteCreatedAtNil

`func (o *TimesheetResponse) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *TimesheetResponse) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetRemoteModifiedAt

`func (o *TimesheetResponse) GetRemoteModifiedAt() string`

GetRemoteModifiedAt returns the RemoteModifiedAt field if non-nil, zero value otherwise.

### GetRemoteModifiedAtOk

`func (o *TimesheetResponse) GetRemoteModifiedAtOk() (*string, bool)`

GetRemoteModifiedAtOk returns a tuple with the RemoteModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteModifiedAt

`func (o *TimesheetResponse) SetRemoteModifiedAt(v string)`

SetRemoteModifiedAt sets RemoteModifiedAt field to given value.


### SetRemoteModifiedAtNil

`func (o *TimesheetResponse) SetRemoteModifiedAtNil(b bool)`

 SetRemoteModifiedAtNil sets the value for RemoteModifiedAt to be an explicit nil

### UnsetRemoteModifiedAt
`func (o *TimesheetResponse) UnsetRemoteModifiedAt()`

UnsetRemoteModifiedAt ensures that no value is present for RemoteModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


