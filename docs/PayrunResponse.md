# PayrunResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Affix-assigned id of the individual | 
**RemoteId** | **string** | the remote system-assigned id of the payrun | 
**RunState** | **string** |  | 
**RunType** | **NullableString** |  | 
**StartDate** | **NullableString** |  | 
**EndDate** | **NullableString** |  | 
**CheckDate** | **NullableString** |  | 
**FieldMappings** | **map[string]interface{}** |  | 

## Methods

### NewPayrunResponse

`func NewPayrunResponse(id string, remoteId string, runState string, runType NullableString, startDate NullableString, endDate NullableString, checkDate NullableString, fieldMappings map[string]interface{}, ) *PayrunResponse`

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


### GetRunState

`func (o *PayrunResponse) GetRunState() string`

GetRunState returns the RunState field if non-nil, zero value otherwise.

### GetRunStateOk

`func (o *PayrunResponse) GetRunStateOk() (*string, bool)`

GetRunStateOk returns a tuple with the RunState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunState

`func (o *PayrunResponse) SetRunState(v string)`

SetRunState sets RunState field to given value.


### GetRunType

`func (o *PayrunResponse) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *PayrunResponse) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *PayrunResponse) SetRunType(v string)`

SetRunType sets RunType field to given value.


### SetRunTypeNil

`func (o *PayrunResponse) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *PayrunResponse) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
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
### GetCheckDate

`func (o *PayrunResponse) GetCheckDate() string`

GetCheckDate returns the CheckDate field if non-nil, zero value otherwise.

### GetCheckDateOk

`func (o *PayrunResponse) GetCheckDateOk() (*string, bool)`

GetCheckDateOk returns a tuple with the CheckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckDate

`func (o *PayrunResponse) SetCheckDate(v string)`

SetCheckDate sets CheckDate field to given value.


### SetCheckDateNil

`func (o *PayrunResponse) SetCheckDateNil(b bool)`

 SetCheckDateNil sets the value for CheckDate to be an explicit nil

### UnsetCheckDate
`func (o *PayrunResponse) UnsetCheckDate()`

UnsetCheckDate ensures that no value is present for CheckDate, not even an explicit nil
### GetFieldMappings

`func (o *PayrunResponse) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *PayrunResponse) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *PayrunResponse) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.


### SetFieldMappingsNil

`func (o *PayrunResponse) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *PayrunResponse) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


