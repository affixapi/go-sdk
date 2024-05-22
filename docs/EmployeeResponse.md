# EmployeeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Affix-assigned id of the individual | [readonly] 
**RemoteId** | **string** | the remote system-assigned id of the individual | [readonly] 
**EmployeeNumber** | **NullableString** |  | 
**FirstName** | **string** | the first name of the individual | 
**LastName** | **string** | the last name of the individual | 
**DisplayFullName** | **NullableString** |  | 
**Nationality** | **NullableString** |  | 
**JobTitle** | **NullableString** |  | 
**WorkEmail** | **NullableString** | the work email of the individual | 
**PersonalEmail** | **NullableString** | the personal email of the individual | 
**MobilePhoneNumber** | **NullableString** | +1234567890 | 
**TaxId** | **NullableString** |  | 
**Gender** | **NullableString** |  | 
**Ethnicity** | **NullableString** |  | 
**MaritalStatus** | **NullableString** | &#x60;other&#x60; option can include co-habitating, civil partnership, separated, widowed, etc  | 
**DateOfBirth** | **NullableString** |  | 
**EmploymentStatus** | [**NullableEmploymentStatusResponse**](EmploymentStatusResponse.md) |  | 
**EmploymentType** | **NullableString** |  | 
**StartDate** | **NullableString** |  | 
**RemoteCreatedAt** | **NullableString** |  | [readonly] 
**TerminationDate** | **NullableString** |  | 
**Avatar** | **NullableString** |  | 
**HomeLocation** | [**NullableAddressResponse**](AddressResponse.md) |  | 
**WorkLocation** | [**NullableLocationResponse**](LocationResponse.md) |  | 
**Manager** | [**NullableEmployeeResponseManager**](EmployeeResponseManager.md) |  | 
**BankAccount** | [**NullableCreateEmployeeRequestBankAccount**](CreateEmployeeRequestBankAccount.md) |  | 
**EmploymentHistory** | [**[]EmploymentHistoryResponse**](EmploymentHistoryResponse.md) |  | 
**CompensationHistory** | [**[]CompensationHistoryResponse**](CompensationHistoryResponse.md) |  | 
**CustomFields** | **map[string]interface{}** |  | 
**Groups** | [**[]GroupResponse**](GroupResponse.md) |  | 
**Dependents** | [**[]CreateEmployeeRequestDependents**](CreateEmployeeRequestDependents.md) |  | 
**EmergencyContacts** | [**[]CreateEmployeeRequestEmergencyContacts**](CreateEmployeeRequestEmergencyContacts.md) |  | 

## Methods

### NewEmployeeResponse

`func NewEmployeeResponse(id string, remoteId string, employeeNumber NullableString, firstName string, lastName string, displayFullName NullableString, nationality NullableString, jobTitle NullableString, workEmail NullableString, personalEmail NullableString, mobilePhoneNumber NullableString, taxId NullableString, gender NullableString, ethnicity NullableString, maritalStatus NullableString, dateOfBirth NullableString, employmentStatus NullableEmploymentStatusResponse, employmentType NullableString, startDate NullableString, remoteCreatedAt NullableString, terminationDate NullableString, avatar NullableString, homeLocation NullableAddressResponse, workLocation NullableLocationResponse, manager NullableEmployeeResponseManager, bankAccount NullableCreateEmployeeRequestBankAccount, employmentHistory []EmploymentHistoryResponse, compensationHistory []CompensationHistoryResponse, customFields map[string]interface{}, groups []GroupResponse, dependents []CreateEmployeeRequestDependents, emergencyContacts []CreateEmployeeRequestEmergencyContacts, ) *EmployeeResponse`

NewEmployeeResponse instantiates a new EmployeeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeResponseWithDefaults

`func NewEmployeeResponseWithDefaults() *EmployeeResponse`

NewEmployeeResponseWithDefaults instantiates a new EmployeeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmployeeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployeeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployeeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetRemoteId

`func (o *EmployeeResponse) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *EmployeeResponse) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *EmployeeResponse) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetEmployeeNumber

`func (o *EmployeeResponse) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *EmployeeResponse) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *EmployeeResponse) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.


### SetEmployeeNumberNil

`func (o *EmployeeResponse) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *EmployeeResponse) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetFirstName

`func (o *EmployeeResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EmployeeResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EmployeeResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *EmployeeResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EmployeeResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EmployeeResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetDisplayFullName

`func (o *EmployeeResponse) GetDisplayFullName() string`

GetDisplayFullName returns the DisplayFullName field if non-nil, zero value otherwise.

### GetDisplayFullNameOk

`func (o *EmployeeResponse) GetDisplayFullNameOk() (*string, bool)`

GetDisplayFullNameOk returns a tuple with the DisplayFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayFullName

`func (o *EmployeeResponse) SetDisplayFullName(v string)`

SetDisplayFullName sets DisplayFullName field to given value.


### SetDisplayFullNameNil

`func (o *EmployeeResponse) SetDisplayFullNameNil(b bool)`

 SetDisplayFullNameNil sets the value for DisplayFullName to be an explicit nil

### UnsetDisplayFullName
`func (o *EmployeeResponse) UnsetDisplayFullName()`

UnsetDisplayFullName ensures that no value is present for DisplayFullName, not even an explicit nil
### GetNationality

`func (o *EmployeeResponse) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *EmployeeResponse) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *EmployeeResponse) SetNationality(v string)`

SetNationality sets Nationality field to given value.


### SetNationalityNil

`func (o *EmployeeResponse) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *EmployeeResponse) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetJobTitle

`func (o *EmployeeResponse) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *EmployeeResponse) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *EmployeeResponse) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.


### SetJobTitleNil

`func (o *EmployeeResponse) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *EmployeeResponse) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetWorkEmail

`func (o *EmployeeResponse) GetWorkEmail() string`

GetWorkEmail returns the WorkEmail field if non-nil, zero value otherwise.

### GetWorkEmailOk

`func (o *EmployeeResponse) GetWorkEmailOk() (*string, bool)`

GetWorkEmailOk returns a tuple with the WorkEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkEmail

`func (o *EmployeeResponse) SetWorkEmail(v string)`

SetWorkEmail sets WorkEmail field to given value.


### SetWorkEmailNil

`func (o *EmployeeResponse) SetWorkEmailNil(b bool)`

 SetWorkEmailNil sets the value for WorkEmail to be an explicit nil

### UnsetWorkEmail
`func (o *EmployeeResponse) UnsetWorkEmail()`

UnsetWorkEmail ensures that no value is present for WorkEmail, not even an explicit nil
### GetPersonalEmail

`func (o *EmployeeResponse) GetPersonalEmail() string`

GetPersonalEmail returns the PersonalEmail field if non-nil, zero value otherwise.

### GetPersonalEmailOk

`func (o *EmployeeResponse) GetPersonalEmailOk() (*string, bool)`

GetPersonalEmailOk returns a tuple with the PersonalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalEmail

`func (o *EmployeeResponse) SetPersonalEmail(v string)`

SetPersonalEmail sets PersonalEmail field to given value.


### SetPersonalEmailNil

`func (o *EmployeeResponse) SetPersonalEmailNil(b bool)`

 SetPersonalEmailNil sets the value for PersonalEmail to be an explicit nil

### UnsetPersonalEmail
`func (o *EmployeeResponse) UnsetPersonalEmail()`

UnsetPersonalEmail ensures that no value is present for PersonalEmail, not even an explicit nil
### GetMobilePhoneNumber

`func (o *EmployeeResponse) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *EmployeeResponse) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *EmployeeResponse) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.


### SetMobilePhoneNumberNil

`func (o *EmployeeResponse) SetMobilePhoneNumberNil(b bool)`

 SetMobilePhoneNumberNil sets the value for MobilePhoneNumber to be an explicit nil

### UnsetMobilePhoneNumber
`func (o *EmployeeResponse) UnsetMobilePhoneNumber()`

UnsetMobilePhoneNumber ensures that no value is present for MobilePhoneNumber, not even an explicit nil
### GetTaxId

`func (o *EmployeeResponse) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *EmployeeResponse) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *EmployeeResponse) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.


### SetTaxIdNil

`func (o *EmployeeResponse) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *EmployeeResponse) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetGender

`func (o *EmployeeResponse) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *EmployeeResponse) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *EmployeeResponse) SetGender(v string)`

SetGender sets Gender field to given value.


### SetGenderNil

`func (o *EmployeeResponse) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *EmployeeResponse) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetEthnicity

`func (o *EmployeeResponse) GetEthnicity() string`

GetEthnicity returns the Ethnicity field if non-nil, zero value otherwise.

### GetEthnicityOk

`func (o *EmployeeResponse) GetEthnicityOk() (*string, bool)`

GetEthnicityOk returns a tuple with the Ethnicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthnicity

`func (o *EmployeeResponse) SetEthnicity(v string)`

SetEthnicity sets Ethnicity field to given value.


### SetEthnicityNil

`func (o *EmployeeResponse) SetEthnicityNil(b bool)`

 SetEthnicityNil sets the value for Ethnicity to be an explicit nil

### UnsetEthnicity
`func (o *EmployeeResponse) UnsetEthnicity()`

UnsetEthnicity ensures that no value is present for Ethnicity, not even an explicit nil
### GetMaritalStatus

`func (o *EmployeeResponse) GetMaritalStatus() string`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *EmployeeResponse) GetMaritalStatusOk() (*string, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *EmployeeResponse) SetMaritalStatus(v string)`

SetMaritalStatus sets MaritalStatus field to given value.


### SetMaritalStatusNil

`func (o *EmployeeResponse) SetMaritalStatusNil(b bool)`

 SetMaritalStatusNil sets the value for MaritalStatus to be an explicit nil

### UnsetMaritalStatus
`func (o *EmployeeResponse) UnsetMaritalStatus()`

UnsetMaritalStatus ensures that no value is present for MaritalStatus, not even an explicit nil
### GetDateOfBirth

`func (o *EmployeeResponse) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *EmployeeResponse) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *EmployeeResponse) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### SetDateOfBirthNil

`func (o *EmployeeResponse) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *EmployeeResponse) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetEmploymentStatus

`func (o *EmployeeResponse) GetEmploymentStatus() EmploymentStatusResponse`

GetEmploymentStatus returns the EmploymentStatus field if non-nil, zero value otherwise.

### GetEmploymentStatusOk

`func (o *EmployeeResponse) GetEmploymentStatusOk() (*EmploymentStatusResponse, bool)`

GetEmploymentStatusOk returns a tuple with the EmploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStatus

`func (o *EmployeeResponse) SetEmploymentStatus(v EmploymentStatusResponse)`

SetEmploymentStatus sets EmploymentStatus field to given value.


### SetEmploymentStatusNil

`func (o *EmployeeResponse) SetEmploymentStatusNil(b bool)`

 SetEmploymentStatusNil sets the value for EmploymentStatus to be an explicit nil

### UnsetEmploymentStatus
`func (o *EmployeeResponse) UnsetEmploymentStatus()`

UnsetEmploymentStatus ensures that no value is present for EmploymentStatus, not even an explicit nil
### GetEmploymentType

`func (o *EmployeeResponse) GetEmploymentType() string`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *EmployeeResponse) GetEmploymentTypeOk() (*string, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *EmployeeResponse) SetEmploymentType(v string)`

SetEmploymentType sets EmploymentType field to given value.


### SetEmploymentTypeNil

`func (o *EmployeeResponse) SetEmploymentTypeNil(b bool)`

 SetEmploymentTypeNil sets the value for EmploymentType to be an explicit nil

### UnsetEmploymentType
`func (o *EmployeeResponse) UnsetEmploymentType()`

UnsetEmploymentType ensures that no value is present for EmploymentType, not even an explicit nil
### GetStartDate

`func (o *EmployeeResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EmployeeResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EmployeeResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### SetStartDateNil

`func (o *EmployeeResponse) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *EmployeeResponse) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetRemoteCreatedAt

`func (o *EmployeeResponse) GetRemoteCreatedAt() string`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *EmployeeResponse) GetRemoteCreatedAtOk() (*string, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *EmployeeResponse) SetRemoteCreatedAt(v string)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.


### SetRemoteCreatedAtNil

`func (o *EmployeeResponse) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *EmployeeResponse) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetTerminationDate

`func (o *EmployeeResponse) GetTerminationDate() string`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *EmployeeResponse) GetTerminationDateOk() (*string, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *EmployeeResponse) SetTerminationDate(v string)`

SetTerminationDate sets TerminationDate field to given value.


### SetTerminationDateNil

`func (o *EmployeeResponse) SetTerminationDateNil(b bool)`

 SetTerminationDateNil sets the value for TerminationDate to be an explicit nil

### UnsetTerminationDate
`func (o *EmployeeResponse) UnsetTerminationDate()`

UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
### GetAvatar

`func (o *EmployeeResponse) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *EmployeeResponse) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *EmployeeResponse) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *EmployeeResponse) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *EmployeeResponse) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetHomeLocation

`func (o *EmployeeResponse) GetHomeLocation() AddressResponse`

GetHomeLocation returns the HomeLocation field if non-nil, zero value otherwise.

### GetHomeLocationOk

`func (o *EmployeeResponse) GetHomeLocationOk() (*AddressResponse, bool)`

GetHomeLocationOk returns a tuple with the HomeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeLocation

`func (o *EmployeeResponse) SetHomeLocation(v AddressResponse)`

SetHomeLocation sets HomeLocation field to given value.


### SetHomeLocationNil

`func (o *EmployeeResponse) SetHomeLocationNil(b bool)`

 SetHomeLocationNil sets the value for HomeLocation to be an explicit nil

### UnsetHomeLocation
`func (o *EmployeeResponse) UnsetHomeLocation()`

UnsetHomeLocation ensures that no value is present for HomeLocation, not even an explicit nil
### GetWorkLocation

`func (o *EmployeeResponse) GetWorkLocation() LocationResponse`

GetWorkLocation returns the WorkLocation field if non-nil, zero value otherwise.

### GetWorkLocationOk

`func (o *EmployeeResponse) GetWorkLocationOk() (*LocationResponse, bool)`

GetWorkLocationOk returns a tuple with the WorkLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLocation

`func (o *EmployeeResponse) SetWorkLocation(v LocationResponse)`

SetWorkLocation sets WorkLocation field to given value.


### SetWorkLocationNil

`func (o *EmployeeResponse) SetWorkLocationNil(b bool)`

 SetWorkLocationNil sets the value for WorkLocation to be an explicit nil

### UnsetWorkLocation
`func (o *EmployeeResponse) UnsetWorkLocation()`

UnsetWorkLocation ensures that no value is present for WorkLocation, not even an explicit nil
### GetManager

`func (o *EmployeeResponse) GetManager() EmployeeResponseManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *EmployeeResponse) GetManagerOk() (*EmployeeResponseManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *EmployeeResponse) SetManager(v EmployeeResponseManager)`

SetManager sets Manager field to given value.


### SetManagerNil

`func (o *EmployeeResponse) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *EmployeeResponse) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetBankAccount

`func (o *EmployeeResponse) GetBankAccount() CreateEmployeeRequestBankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *EmployeeResponse) GetBankAccountOk() (*CreateEmployeeRequestBankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *EmployeeResponse) SetBankAccount(v CreateEmployeeRequestBankAccount)`

SetBankAccount sets BankAccount field to given value.


### SetBankAccountNil

`func (o *EmployeeResponse) SetBankAccountNil(b bool)`

 SetBankAccountNil sets the value for BankAccount to be an explicit nil

### UnsetBankAccount
`func (o *EmployeeResponse) UnsetBankAccount()`

UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
### GetEmploymentHistory

`func (o *EmployeeResponse) GetEmploymentHistory() []EmploymentHistoryResponse`

GetEmploymentHistory returns the EmploymentHistory field if non-nil, zero value otherwise.

### GetEmploymentHistoryOk

`func (o *EmployeeResponse) GetEmploymentHistoryOk() (*[]EmploymentHistoryResponse, bool)`

GetEmploymentHistoryOk returns a tuple with the EmploymentHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentHistory

`func (o *EmployeeResponse) SetEmploymentHistory(v []EmploymentHistoryResponse)`

SetEmploymentHistory sets EmploymentHistory field to given value.


### SetEmploymentHistoryNil

`func (o *EmployeeResponse) SetEmploymentHistoryNil(b bool)`

 SetEmploymentHistoryNil sets the value for EmploymentHistory to be an explicit nil

### UnsetEmploymentHistory
`func (o *EmployeeResponse) UnsetEmploymentHistory()`

UnsetEmploymentHistory ensures that no value is present for EmploymentHistory, not even an explicit nil
### GetCompensationHistory

`func (o *EmployeeResponse) GetCompensationHistory() []CompensationHistoryResponse`

GetCompensationHistory returns the CompensationHistory field if non-nil, zero value otherwise.

### GetCompensationHistoryOk

`func (o *EmployeeResponse) GetCompensationHistoryOk() (*[]CompensationHistoryResponse, bool)`

GetCompensationHistoryOk returns a tuple with the CompensationHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompensationHistory

`func (o *EmployeeResponse) SetCompensationHistory(v []CompensationHistoryResponse)`

SetCompensationHistory sets CompensationHistory field to given value.


### SetCompensationHistoryNil

`func (o *EmployeeResponse) SetCompensationHistoryNil(b bool)`

 SetCompensationHistoryNil sets the value for CompensationHistory to be an explicit nil

### UnsetCompensationHistory
`func (o *EmployeeResponse) UnsetCompensationHistory()`

UnsetCompensationHistory ensures that no value is present for CompensationHistory, not even an explicit nil
### GetCustomFields

`func (o *EmployeeResponse) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *EmployeeResponse) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *EmployeeResponse) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.


### SetCustomFieldsNil

`func (o *EmployeeResponse) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *EmployeeResponse) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetGroups

`func (o *EmployeeResponse) GetGroups() []GroupResponse`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EmployeeResponse) GetGroupsOk() (*[]GroupResponse, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EmployeeResponse) SetGroups(v []GroupResponse)`

SetGroups sets Groups field to given value.


### SetGroupsNil

`func (o *EmployeeResponse) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *EmployeeResponse) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetDependents

`func (o *EmployeeResponse) GetDependents() []CreateEmployeeRequestDependents`

GetDependents returns the Dependents field if non-nil, zero value otherwise.

### GetDependentsOk

`func (o *EmployeeResponse) GetDependentsOk() (*[]CreateEmployeeRequestDependents, bool)`

GetDependentsOk returns a tuple with the Dependents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependents

`func (o *EmployeeResponse) SetDependents(v []CreateEmployeeRequestDependents)`

SetDependents sets Dependents field to given value.


### SetDependentsNil

`func (o *EmployeeResponse) SetDependentsNil(b bool)`

 SetDependentsNil sets the value for Dependents to be an explicit nil

### UnsetDependents
`func (o *EmployeeResponse) UnsetDependents()`

UnsetDependents ensures that no value is present for Dependents, not even an explicit nil
### GetEmergencyContacts

`func (o *EmployeeResponse) GetEmergencyContacts() []CreateEmployeeRequestEmergencyContacts`

GetEmergencyContacts returns the EmergencyContacts field if non-nil, zero value otherwise.

### GetEmergencyContactsOk

`func (o *EmployeeResponse) GetEmergencyContactsOk() (*[]CreateEmployeeRequestEmergencyContacts, bool)`

GetEmergencyContactsOk returns a tuple with the EmergencyContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyContacts

`func (o *EmployeeResponse) SetEmergencyContacts(v []CreateEmployeeRequestEmergencyContacts)`

SetEmergencyContacts sets EmergencyContacts field to given value.


### SetEmergencyContactsNil

`func (o *EmployeeResponse) SetEmergencyContactsNil(b bool)`

 SetEmergencyContactsNil sets the value for EmergencyContacts to be an explicit nil

### UnsetEmergencyContacts
`func (o *EmployeeResponse) UnsetEmergencyContacts()`

UnsetEmergencyContacts ensures that no value is present for EmergencyContacts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


