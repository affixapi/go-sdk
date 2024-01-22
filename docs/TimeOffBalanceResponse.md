# TimeOffBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployeeId** | **string** | The Affix-assigned id of the individual | 
**RemoteId** | **string** | the remote system-assigned id of the individual | 
**Balance** | **float32** |  | 
**Used** | **float32** |  | 
**PolicyName** | **NullableString** | The name of the policy, as assigned by the remote system | 
**PolicyType** | **NullableString** |  | 
**RemoteCreatedAt** | **NullableString** |  | 
**RemoteModifiedAt** | **NullableString** |  | 

## Methods

### NewTimeOffBalanceResponse

`func NewTimeOffBalanceResponse(employeeId string, remoteId string, balance float32, used float32, policyName NullableString, policyType NullableString, remoteCreatedAt NullableString, remoteModifiedAt NullableString, ) *TimeOffBalanceResponse`

NewTimeOffBalanceResponse instantiates a new TimeOffBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOffBalanceResponseWithDefaults

`func NewTimeOffBalanceResponseWithDefaults() *TimeOffBalanceResponse`

NewTimeOffBalanceResponseWithDefaults instantiates a new TimeOffBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployeeId

`func (o *TimeOffBalanceResponse) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *TimeOffBalanceResponse) GetEmployeeIdOk() (*string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId

`func (o *TimeOffBalanceResponse) SetEmployeeId(v string)`

SetEmployeeId sets EmployeeId field to given value.


### GetRemoteId

`func (o *TimeOffBalanceResponse) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *TimeOffBalanceResponse) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *TimeOffBalanceResponse) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetBalance

`func (o *TimeOffBalanceResponse) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *TimeOffBalanceResponse) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *TimeOffBalanceResponse) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetUsed

`func (o *TimeOffBalanceResponse) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *TimeOffBalanceResponse) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *TimeOffBalanceResponse) SetUsed(v float32)`

SetUsed sets Used field to given value.


### GetPolicyName

`func (o *TimeOffBalanceResponse) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *TimeOffBalanceResponse) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *TimeOffBalanceResponse) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### SetPolicyNameNil

`func (o *TimeOffBalanceResponse) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *TimeOffBalanceResponse) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetPolicyType

`func (o *TimeOffBalanceResponse) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *TimeOffBalanceResponse) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *TimeOffBalanceResponse) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.


### SetPolicyTypeNil

`func (o *TimeOffBalanceResponse) SetPolicyTypeNil(b bool)`

 SetPolicyTypeNil sets the value for PolicyType to be an explicit nil

### UnsetPolicyType
`func (o *TimeOffBalanceResponse) UnsetPolicyType()`

UnsetPolicyType ensures that no value is present for PolicyType, not even an explicit nil
### GetRemoteCreatedAt

`func (o *TimeOffBalanceResponse) GetRemoteCreatedAt() string`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *TimeOffBalanceResponse) GetRemoteCreatedAtOk() (*string, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *TimeOffBalanceResponse) SetRemoteCreatedAt(v string)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.


### SetRemoteCreatedAtNil

`func (o *TimeOffBalanceResponse) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *TimeOffBalanceResponse) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetRemoteModifiedAt

`func (o *TimeOffBalanceResponse) GetRemoteModifiedAt() string`

GetRemoteModifiedAt returns the RemoteModifiedAt field if non-nil, zero value otherwise.

### GetRemoteModifiedAtOk

`func (o *TimeOffBalanceResponse) GetRemoteModifiedAtOk() (*string, bool)`

GetRemoteModifiedAtOk returns a tuple with the RemoteModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteModifiedAt

`func (o *TimeOffBalanceResponse) SetRemoteModifiedAt(v string)`

SetRemoteModifiedAt sets RemoteModifiedAt field to given value.


### SetRemoteModifiedAtNil

`func (o *TimeOffBalanceResponse) SetRemoteModifiedAtNil(b bool)`

 SetRemoteModifiedAtNil sets the value for RemoteModifiedAt to be an explicit nil

### UnsetRemoteModifiedAt
`func (o *TimeOffBalanceResponse) UnsetRemoteModifiedAt()`

UnsetRemoteModifiedAt ensures that no value is present for RemoteModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


