# CreateEmployeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployeeNumber** | Pointer to **NullableString** |  | [optional] 
**FirstName** | **string** | the first name of the individual | 
**LastName** | **string** | the last name of the individual | 
**DisplayFullName** | Pointer to **NullableString** |  | [optional] 
**Nationality** | Pointer to **NullableString** |  | [optional] 
**JobTitle** | Pointer to **NullableString** |  | [optional] 
**WorkEmail** | Pointer to **NullableString** | the work email of the individual | [optional] 
**PersonalEmail** | Pointer to **NullableString** | the personal email of the individual | [optional] 
**MobilePhoneNumber** | Pointer to **NullableString** | +1234567890 | [optional] 
**TaxId** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**Ethnicity** | Pointer to **NullableString** |  | [optional] 
**MaritalStatus** | Pointer to **NullableString** |  | [optional] 
**DateOfBirth** | Pointer to **NullableString** |  | [optional] 
**EmploymentStatus** | Pointer to **NullableString** |  | [optional] 
**EmploymentType** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**TerminationDate** | Pointer to **NullableString** |  | [optional] 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**HomeLocation** | Pointer to [**NullableCreateEmployeeRequestHomeLocation**](CreateEmployeeRequestHomeLocation.md) |  | [optional] 
**WorkLocation** | Pointer to [**NullableCreateEmployeeRequestWorkLocation**](CreateEmployeeRequestWorkLocation.md) |  | [optional] 
**Manager** | Pointer to [**NullableCreateEmployeeRequestManager**](CreateEmployeeRequestManager.md) |  | [optional] 
**BankAccount** | Pointer to [**NullableCreateEmployeeRequestBankAccount**](CreateEmployeeRequestBankAccount.md) |  | [optional] 
**Employments** | Pointer to [**[]EmploymentNoNullEnumRequest**](EmploymentNoNullEnumRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Groups** | Pointer to [**[]CreateEmployeeRequestGroups**](CreateEmployeeRequestGroups.md) |  | [optional] 
**Company** | Pointer to [**NullableCreateEmployeeRequestCompany**](CreateEmployeeRequestCompany.md) |  | [optional] 

## Methods

### NewCreateEmployeeRequest

`func NewCreateEmployeeRequest(firstName string, lastName string, ) *CreateEmployeeRequest`

NewCreateEmployeeRequest instantiates a new CreateEmployeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmployeeRequestWithDefaults

`func NewCreateEmployeeRequestWithDefaults() *CreateEmployeeRequest`

NewCreateEmployeeRequestWithDefaults instantiates a new CreateEmployeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployeeNumber

`func (o *CreateEmployeeRequest) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *CreateEmployeeRequest) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *CreateEmployeeRequest) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *CreateEmployeeRequest) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### SetEmployeeNumberNil

`func (o *CreateEmployeeRequest) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *CreateEmployeeRequest) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetFirstName

`func (o *CreateEmployeeRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateEmployeeRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CreateEmployeeRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *CreateEmployeeRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateEmployeeRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CreateEmployeeRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetDisplayFullName

`func (o *CreateEmployeeRequest) GetDisplayFullName() string`

GetDisplayFullName returns the DisplayFullName field if non-nil, zero value otherwise.

### GetDisplayFullNameOk

`func (o *CreateEmployeeRequest) GetDisplayFullNameOk() (*string, bool)`

GetDisplayFullNameOk returns a tuple with the DisplayFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayFullName

`func (o *CreateEmployeeRequest) SetDisplayFullName(v string)`

SetDisplayFullName sets DisplayFullName field to given value.

### HasDisplayFullName

`func (o *CreateEmployeeRequest) HasDisplayFullName() bool`

HasDisplayFullName returns a boolean if a field has been set.

### SetDisplayFullNameNil

`func (o *CreateEmployeeRequest) SetDisplayFullNameNil(b bool)`

 SetDisplayFullNameNil sets the value for DisplayFullName to be an explicit nil

### UnsetDisplayFullName
`func (o *CreateEmployeeRequest) UnsetDisplayFullName()`

UnsetDisplayFullName ensures that no value is present for DisplayFullName, not even an explicit nil
### GetNationality

`func (o *CreateEmployeeRequest) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *CreateEmployeeRequest) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *CreateEmployeeRequest) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *CreateEmployeeRequest) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *CreateEmployeeRequest) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *CreateEmployeeRequest) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetJobTitle

`func (o *CreateEmployeeRequest) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *CreateEmployeeRequest) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *CreateEmployeeRequest) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *CreateEmployeeRequest) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *CreateEmployeeRequest) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *CreateEmployeeRequest) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetWorkEmail

`func (o *CreateEmployeeRequest) GetWorkEmail() string`

GetWorkEmail returns the WorkEmail field if non-nil, zero value otherwise.

### GetWorkEmailOk

`func (o *CreateEmployeeRequest) GetWorkEmailOk() (*string, bool)`

GetWorkEmailOk returns a tuple with the WorkEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkEmail

`func (o *CreateEmployeeRequest) SetWorkEmail(v string)`

SetWorkEmail sets WorkEmail field to given value.

### HasWorkEmail

`func (o *CreateEmployeeRequest) HasWorkEmail() bool`

HasWorkEmail returns a boolean if a field has been set.

### SetWorkEmailNil

`func (o *CreateEmployeeRequest) SetWorkEmailNil(b bool)`

 SetWorkEmailNil sets the value for WorkEmail to be an explicit nil

### UnsetWorkEmail
`func (o *CreateEmployeeRequest) UnsetWorkEmail()`

UnsetWorkEmail ensures that no value is present for WorkEmail, not even an explicit nil
### GetPersonalEmail

`func (o *CreateEmployeeRequest) GetPersonalEmail() string`

GetPersonalEmail returns the PersonalEmail field if non-nil, zero value otherwise.

### GetPersonalEmailOk

`func (o *CreateEmployeeRequest) GetPersonalEmailOk() (*string, bool)`

GetPersonalEmailOk returns a tuple with the PersonalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalEmail

`func (o *CreateEmployeeRequest) SetPersonalEmail(v string)`

SetPersonalEmail sets PersonalEmail field to given value.

### HasPersonalEmail

`func (o *CreateEmployeeRequest) HasPersonalEmail() bool`

HasPersonalEmail returns a boolean if a field has been set.

### SetPersonalEmailNil

`func (o *CreateEmployeeRequest) SetPersonalEmailNil(b bool)`

 SetPersonalEmailNil sets the value for PersonalEmail to be an explicit nil

### UnsetPersonalEmail
`func (o *CreateEmployeeRequest) UnsetPersonalEmail()`

UnsetPersonalEmail ensures that no value is present for PersonalEmail, not even an explicit nil
### GetMobilePhoneNumber

`func (o *CreateEmployeeRequest) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *CreateEmployeeRequest) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *CreateEmployeeRequest) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *CreateEmployeeRequest) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### SetMobilePhoneNumberNil

`func (o *CreateEmployeeRequest) SetMobilePhoneNumberNil(b bool)`

 SetMobilePhoneNumberNil sets the value for MobilePhoneNumber to be an explicit nil

### UnsetMobilePhoneNumber
`func (o *CreateEmployeeRequest) UnsetMobilePhoneNumber()`

UnsetMobilePhoneNumber ensures that no value is present for MobilePhoneNumber, not even an explicit nil
### GetTaxId

`func (o *CreateEmployeeRequest) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CreateEmployeeRequest) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CreateEmployeeRequest) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CreateEmployeeRequest) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *CreateEmployeeRequest) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *CreateEmployeeRequest) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetGender

`func (o *CreateEmployeeRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CreateEmployeeRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CreateEmployeeRequest) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CreateEmployeeRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *CreateEmployeeRequest) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *CreateEmployeeRequest) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetEthnicity

`func (o *CreateEmployeeRequest) GetEthnicity() string`

GetEthnicity returns the Ethnicity field if non-nil, zero value otherwise.

### GetEthnicityOk

`func (o *CreateEmployeeRequest) GetEthnicityOk() (*string, bool)`

GetEthnicityOk returns a tuple with the Ethnicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthnicity

`func (o *CreateEmployeeRequest) SetEthnicity(v string)`

SetEthnicity sets Ethnicity field to given value.

### HasEthnicity

`func (o *CreateEmployeeRequest) HasEthnicity() bool`

HasEthnicity returns a boolean if a field has been set.

### SetEthnicityNil

`func (o *CreateEmployeeRequest) SetEthnicityNil(b bool)`

 SetEthnicityNil sets the value for Ethnicity to be an explicit nil

### UnsetEthnicity
`func (o *CreateEmployeeRequest) UnsetEthnicity()`

UnsetEthnicity ensures that no value is present for Ethnicity, not even an explicit nil
### GetMaritalStatus

`func (o *CreateEmployeeRequest) GetMaritalStatus() string`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *CreateEmployeeRequest) GetMaritalStatusOk() (*string, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *CreateEmployeeRequest) SetMaritalStatus(v string)`

SetMaritalStatus sets MaritalStatus field to given value.

### HasMaritalStatus

`func (o *CreateEmployeeRequest) HasMaritalStatus() bool`

HasMaritalStatus returns a boolean if a field has been set.

### SetMaritalStatusNil

`func (o *CreateEmployeeRequest) SetMaritalStatusNil(b bool)`

 SetMaritalStatusNil sets the value for MaritalStatus to be an explicit nil

### UnsetMaritalStatus
`func (o *CreateEmployeeRequest) UnsetMaritalStatus()`

UnsetMaritalStatus ensures that no value is present for MaritalStatus, not even an explicit nil
### GetDateOfBirth

`func (o *CreateEmployeeRequest) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CreateEmployeeRequest) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CreateEmployeeRequest) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *CreateEmployeeRequest) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *CreateEmployeeRequest) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *CreateEmployeeRequest) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetEmploymentStatus

`func (o *CreateEmployeeRequest) GetEmploymentStatus() string`

GetEmploymentStatus returns the EmploymentStatus field if non-nil, zero value otherwise.

### GetEmploymentStatusOk

`func (o *CreateEmployeeRequest) GetEmploymentStatusOk() (*string, bool)`

GetEmploymentStatusOk returns a tuple with the EmploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStatus

`func (o *CreateEmployeeRequest) SetEmploymentStatus(v string)`

SetEmploymentStatus sets EmploymentStatus field to given value.

### HasEmploymentStatus

`func (o *CreateEmployeeRequest) HasEmploymentStatus() bool`

HasEmploymentStatus returns a boolean if a field has been set.

### SetEmploymentStatusNil

`func (o *CreateEmployeeRequest) SetEmploymentStatusNil(b bool)`

 SetEmploymentStatusNil sets the value for EmploymentStatus to be an explicit nil

### UnsetEmploymentStatus
`func (o *CreateEmployeeRequest) UnsetEmploymentStatus()`

UnsetEmploymentStatus ensures that no value is present for EmploymentStatus, not even an explicit nil
### GetEmploymentType

`func (o *CreateEmployeeRequest) GetEmploymentType() string`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *CreateEmployeeRequest) GetEmploymentTypeOk() (*string, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *CreateEmployeeRequest) SetEmploymentType(v string)`

SetEmploymentType sets EmploymentType field to given value.

### HasEmploymentType

`func (o *CreateEmployeeRequest) HasEmploymentType() bool`

HasEmploymentType returns a boolean if a field has been set.

### SetEmploymentTypeNil

`func (o *CreateEmployeeRequest) SetEmploymentTypeNil(b bool)`

 SetEmploymentTypeNil sets the value for EmploymentType to be an explicit nil

### UnsetEmploymentType
`func (o *CreateEmployeeRequest) UnsetEmploymentType()`

UnsetEmploymentType ensures that no value is present for EmploymentType, not even an explicit nil
### GetStartDate

`func (o *CreateEmployeeRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateEmployeeRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateEmployeeRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateEmployeeRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *CreateEmployeeRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *CreateEmployeeRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetTerminationDate

`func (o *CreateEmployeeRequest) GetTerminationDate() string`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *CreateEmployeeRequest) GetTerminationDateOk() (*string, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *CreateEmployeeRequest) SetTerminationDate(v string)`

SetTerminationDate sets TerminationDate field to given value.

### HasTerminationDate

`func (o *CreateEmployeeRequest) HasTerminationDate() bool`

HasTerminationDate returns a boolean if a field has been set.

### SetTerminationDateNil

`func (o *CreateEmployeeRequest) SetTerminationDateNil(b bool)`

 SetTerminationDateNil sets the value for TerminationDate to be an explicit nil

### UnsetTerminationDate
`func (o *CreateEmployeeRequest) UnsetTerminationDate()`

UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
### GetAvatar

`func (o *CreateEmployeeRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CreateEmployeeRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CreateEmployeeRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CreateEmployeeRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *CreateEmployeeRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *CreateEmployeeRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetHomeLocation

`func (o *CreateEmployeeRequest) GetHomeLocation() CreateEmployeeRequestHomeLocation`

GetHomeLocation returns the HomeLocation field if non-nil, zero value otherwise.

### GetHomeLocationOk

`func (o *CreateEmployeeRequest) GetHomeLocationOk() (*CreateEmployeeRequestHomeLocation, bool)`

GetHomeLocationOk returns a tuple with the HomeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeLocation

`func (o *CreateEmployeeRequest) SetHomeLocation(v CreateEmployeeRequestHomeLocation)`

SetHomeLocation sets HomeLocation field to given value.

### HasHomeLocation

`func (o *CreateEmployeeRequest) HasHomeLocation() bool`

HasHomeLocation returns a boolean if a field has been set.

### SetHomeLocationNil

`func (o *CreateEmployeeRequest) SetHomeLocationNil(b bool)`

 SetHomeLocationNil sets the value for HomeLocation to be an explicit nil

### UnsetHomeLocation
`func (o *CreateEmployeeRequest) UnsetHomeLocation()`

UnsetHomeLocation ensures that no value is present for HomeLocation, not even an explicit nil
### GetWorkLocation

`func (o *CreateEmployeeRequest) GetWorkLocation() CreateEmployeeRequestWorkLocation`

GetWorkLocation returns the WorkLocation field if non-nil, zero value otherwise.

### GetWorkLocationOk

`func (o *CreateEmployeeRequest) GetWorkLocationOk() (*CreateEmployeeRequestWorkLocation, bool)`

GetWorkLocationOk returns a tuple with the WorkLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLocation

`func (o *CreateEmployeeRequest) SetWorkLocation(v CreateEmployeeRequestWorkLocation)`

SetWorkLocation sets WorkLocation field to given value.

### HasWorkLocation

`func (o *CreateEmployeeRequest) HasWorkLocation() bool`

HasWorkLocation returns a boolean if a field has been set.

### SetWorkLocationNil

`func (o *CreateEmployeeRequest) SetWorkLocationNil(b bool)`

 SetWorkLocationNil sets the value for WorkLocation to be an explicit nil

### UnsetWorkLocation
`func (o *CreateEmployeeRequest) UnsetWorkLocation()`

UnsetWorkLocation ensures that no value is present for WorkLocation, not even an explicit nil
### GetManager

`func (o *CreateEmployeeRequest) GetManager() CreateEmployeeRequestManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *CreateEmployeeRequest) GetManagerOk() (*CreateEmployeeRequestManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *CreateEmployeeRequest) SetManager(v CreateEmployeeRequestManager)`

SetManager sets Manager field to given value.

### HasManager

`func (o *CreateEmployeeRequest) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *CreateEmployeeRequest) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *CreateEmployeeRequest) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetBankAccount

`func (o *CreateEmployeeRequest) GetBankAccount() CreateEmployeeRequestBankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *CreateEmployeeRequest) GetBankAccountOk() (*CreateEmployeeRequestBankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *CreateEmployeeRequest) SetBankAccount(v CreateEmployeeRequestBankAccount)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *CreateEmployeeRequest) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### SetBankAccountNil

`func (o *CreateEmployeeRequest) SetBankAccountNil(b bool)`

 SetBankAccountNil sets the value for BankAccount to be an explicit nil

### UnsetBankAccount
`func (o *CreateEmployeeRequest) UnsetBankAccount()`

UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
### GetEmployments

`func (o *CreateEmployeeRequest) GetEmployments() []EmploymentNoNullEnumRequest`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *CreateEmployeeRequest) GetEmploymentsOk() (*[]EmploymentNoNullEnumRequest, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *CreateEmployeeRequest) SetEmployments(v []EmploymentNoNullEnumRequest)`

SetEmployments sets Employments field to given value.

### HasEmployments

`func (o *CreateEmployeeRequest) HasEmployments() bool`

HasEmployments returns a boolean if a field has been set.

### SetEmploymentsNil

`func (o *CreateEmployeeRequest) SetEmploymentsNil(b bool)`

 SetEmploymentsNil sets the value for Employments to be an explicit nil

### UnsetEmployments
`func (o *CreateEmployeeRequest) UnsetEmployments()`

UnsetEmployments ensures that no value is present for Employments, not even an explicit nil
### GetCustomFields

`func (o *CreateEmployeeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateEmployeeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateEmployeeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateEmployeeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CreateEmployeeRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CreateEmployeeRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetGroups

`func (o *CreateEmployeeRequest) GetGroups() []CreateEmployeeRequestGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CreateEmployeeRequest) GetGroupsOk() (*[]CreateEmployeeRequestGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CreateEmployeeRequest) SetGroups(v []CreateEmployeeRequestGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *CreateEmployeeRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *CreateEmployeeRequest) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *CreateEmployeeRequest) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetCompany

`func (o *CreateEmployeeRequest) GetCompany() CreateEmployeeRequestCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CreateEmployeeRequest) GetCompanyOk() (*CreateEmployeeRequestCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CreateEmployeeRequest) SetCompany(v CreateEmployeeRequestCompany)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CreateEmployeeRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *CreateEmployeeRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *CreateEmployeeRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


