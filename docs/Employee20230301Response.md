# Employee20230301Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Affix-assigned id of the individual | 
**RemoteId** | **string** | the remote system-assigned id of the individual | 
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
**MaritalStatus** | **NullableString** |  | 
**DateOfBirth** | **NullableString** |  | 
**EmploymentStatus** | **NullableString** |  | 
**EmploymentType** | **NullableString** |  | 
**StartDate** | **NullableString** |  | 
**RemoteCreatedAt** | **NullableString** |  | 
**TerminationDate** | **NullableString** |  | 
**Avatar** | **NullableString** |  | 
**HomeLocation** | [**Employee20230301ResponseHomeLocation**](Employee20230301ResponseHomeLocation.md) |  | 
**WorkLocation** | [**NullableEmployee20230301ResponseWorkLocation**](Employee20230301ResponseWorkLocation.md) |  | 
**Manager** | [**NullableEmployee20230301ResponseManager**](Employee20230301ResponseManager.md) |  | 
**BankAccount** | [**NullableEmployee20230301ResponseBankAccount**](Employee20230301ResponseBankAccount.md) |  | 
**Employments** | [**[]EmploymentResponse**](EmploymentResponse.md) |  | 
**CustomFields** | **map[string]interface{}** |  | 
**Groups** | [**[]Employee20230301ResponseGroups**](Employee20230301ResponseGroups.md) |  | 
**Company** | [**NullableEmployee20230301ResponseCompany**](Employee20230301ResponseCompany.md) |  | 

## Methods

### NewEmployee20230301Response

`func NewEmployee20230301Response(id string, remoteId string, employeeNumber NullableString, firstName string, lastName string, displayFullName NullableString, nationality NullableString, jobTitle NullableString, workEmail NullableString, personalEmail NullableString, mobilePhoneNumber NullableString, taxId NullableString, gender NullableString, ethnicity NullableString, maritalStatus NullableString, dateOfBirth NullableString, employmentStatus NullableString, employmentType NullableString, startDate NullableString, remoteCreatedAt NullableString, terminationDate NullableString, avatar NullableString, homeLocation Employee20230301ResponseHomeLocation, workLocation NullableEmployee20230301ResponseWorkLocation, manager NullableEmployee20230301ResponseManager, bankAccount NullableEmployee20230301ResponseBankAccount, employments []EmploymentResponse, customFields map[string]interface{}, groups []Employee20230301ResponseGroups, company NullableEmployee20230301ResponseCompany, ) *Employee20230301Response`

NewEmployee20230301Response instantiates a new Employee20230301Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployee20230301ResponseWithDefaults

`func NewEmployee20230301ResponseWithDefaults() *Employee20230301Response`

NewEmployee20230301ResponseWithDefaults instantiates a new Employee20230301Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Employee20230301Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Employee20230301Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Employee20230301Response) SetId(v string)`

SetId sets Id field to given value.


### GetRemoteId

`func (o *Employee20230301Response) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Employee20230301Response) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Employee20230301Response) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetEmployeeNumber

`func (o *Employee20230301Response) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *Employee20230301Response) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *Employee20230301Response) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.


### SetEmployeeNumberNil

`func (o *Employee20230301Response) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *Employee20230301Response) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetFirstName

`func (o *Employee20230301Response) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Employee20230301Response) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Employee20230301Response) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *Employee20230301Response) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Employee20230301Response) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Employee20230301Response) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetDisplayFullName

`func (o *Employee20230301Response) GetDisplayFullName() string`

GetDisplayFullName returns the DisplayFullName field if non-nil, zero value otherwise.

### GetDisplayFullNameOk

`func (o *Employee20230301Response) GetDisplayFullNameOk() (*string, bool)`

GetDisplayFullNameOk returns a tuple with the DisplayFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayFullName

`func (o *Employee20230301Response) SetDisplayFullName(v string)`

SetDisplayFullName sets DisplayFullName field to given value.


### SetDisplayFullNameNil

`func (o *Employee20230301Response) SetDisplayFullNameNil(b bool)`

 SetDisplayFullNameNil sets the value for DisplayFullName to be an explicit nil

### UnsetDisplayFullName
`func (o *Employee20230301Response) UnsetDisplayFullName()`

UnsetDisplayFullName ensures that no value is present for DisplayFullName, not even an explicit nil
### GetNationality

`func (o *Employee20230301Response) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *Employee20230301Response) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *Employee20230301Response) SetNationality(v string)`

SetNationality sets Nationality field to given value.


### SetNationalityNil

`func (o *Employee20230301Response) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *Employee20230301Response) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetJobTitle

`func (o *Employee20230301Response) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Employee20230301Response) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Employee20230301Response) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.


### SetJobTitleNil

`func (o *Employee20230301Response) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *Employee20230301Response) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetWorkEmail

`func (o *Employee20230301Response) GetWorkEmail() string`

GetWorkEmail returns the WorkEmail field if non-nil, zero value otherwise.

### GetWorkEmailOk

`func (o *Employee20230301Response) GetWorkEmailOk() (*string, bool)`

GetWorkEmailOk returns a tuple with the WorkEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkEmail

`func (o *Employee20230301Response) SetWorkEmail(v string)`

SetWorkEmail sets WorkEmail field to given value.


### SetWorkEmailNil

`func (o *Employee20230301Response) SetWorkEmailNil(b bool)`

 SetWorkEmailNil sets the value for WorkEmail to be an explicit nil

### UnsetWorkEmail
`func (o *Employee20230301Response) UnsetWorkEmail()`

UnsetWorkEmail ensures that no value is present for WorkEmail, not even an explicit nil
### GetPersonalEmail

`func (o *Employee20230301Response) GetPersonalEmail() string`

GetPersonalEmail returns the PersonalEmail field if non-nil, zero value otherwise.

### GetPersonalEmailOk

`func (o *Employee20230301Response) GetPersonalEmailOk() (*string, bool)`

GetPersonalEmailOk returns a tuple with the PersonalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalEmail

`func (o *Employee20230301Response) SetPersonalEmail(v string)`

SetPersonalEmail sets PersonalEmail field to given value.


### SetPersonalEmailNil

`func (o *Employee20230301Response) SetPersonalEmailNil(b bool)`

 SetPersonalEmailNil sets the value for PersonalEmail to be an explicit nil

### UnsetPersonalEmail
`func (o *Employee20230301Response) UnsetPersonalEmail()`

UnsetPersonalEmail ensures that no value is present for PersonalEmail, not even an explicit nil
### GetMobilePhoneNumber

`func (o *Employee20230301Response) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *Employee20230301Response) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *Employee20230301Response) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.


### SetMobilePhoneNumberNil

`func (o *Employee20230301Response) SetMobilePhoneNumberNil(b bool)`

 SetMobilePhoneNumberNil sets the value for MobilePhoneNumber to be an explicit nil

### UnsetMobilePhoneNumber
`func (o *Employee20230301Response) UnsetMobilePhoneNumber()`

UnsetMobilePhoneNumber ensures that no value is present for MobilePhoneNumber, not even an explicit nil
### GetTaxId

`func (o *Employee20230301Response) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *Employee20230301Response) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *Employee20230301Response) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.


### SetTaxIdNil

`func (o *Employee20230301Response) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *Employee20230301Response) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetGender

`func (o *Employee20230301Response) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Employee20230301Response) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Employee20230301Response) SetGender(v string)`

SetGender sets Gender field to given value.


### SetGenderNil

`func (o *Employee20230301Response) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *Employee20230301Response) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetEthnicity

`func (o *Employee20230301Response) GetEthnicity() string`

GetEthnicity returns the Ethnicity field if non-nil, zero value otherwise.

### GetEthnicityOk

`func (o *Employee20230301Response) GetEthnicityOk() (*string, bool)`

GetEthnicityOk returns a tuple with the Ethnicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthnicity

`func (o *Employee20230301Response) SetEthnicity(v string)`

SetEthnicity sets Ethnicity field to given value.


### SetEthnicityNil

`func (o *Employee20230301Response) SetEthnicityNil(b bool)`

 SetEthnicityNil sets the value for Ethnicity to be an explicit nil

### UnsetEthnicity
`func (o *Employee20230301Response) UnsetEthnicity()`

UnsetEthnicity ensures that no value is present for Ethnicity, not even an explicit nil
### GetMaritalStatus

`func (o *Employee20230301Response) GetMaritalStatus() string`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *Employee20230301Response) GetMaritalStatusOk() (*string, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *Employee20230301Response) SetMaritalStatus(v string)`

SetMaritalStatus sets MaritalStatus field to given value.


### SetMaritalStatusNil

`func (o *Employee20230301Response) SetMaritalStatusNil(b bool)`

 SetMaritalStatusNil sets the value for MaritalStatus to be an explicit nil

### UnsetMaritalStatus
`func (o *Employee20230301Response) UnsetMaritalStatus()`

UnsetMaritalStatus ensures that no value is present for MaritalStatus, not even an explicit nil
### GetDateOfBirth

`func (o *Employee20230301Response) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Employee20230301Response) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *Employee20230301Response) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### SetDateOfBirthNil

`func (o *Employee20230301Response) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *Employee20230301Response) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetEmploymentStatus

`func (o *Employee20230301Response) GetEmploymentStatus() string`

GetEmploymentStatus returns the EmploymentStatus field if non-nil, zero value otherwise.

### GetEmploymentStatusOk

`func (o *Employee20230301Response) GetEmploymentStatusOk() (*string, bool)`

GetEmploymentStatusOk returns a tuple with the EmploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStatus

`func (o *Employee20230301Response) SetEmploymentStatus(v string)`

SetEmploymentStatus sets EmploymentStatus field to given value.


### SetEmploymentStatusNil

`func (o *Employee20230301Response) SetEmploymentStatusNil(b bool)`

 SetEmploymentStatusNil sets the value for EmploymentStatus to be an explicit nil

### UnsetEmploymentStatus
`func (o *Employee20230301Response) UnsetEmploymentStatus()`

UnsetEmploymentStatus ensures that no value is present for EmploymentStatus, not even an explicit nil
### GetEmploymentType

`func (o *Employee20230301Response) GetEmploymentType() string`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *Employee20230301Response) GetEmploymentTypeOk() (*string, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *Employee20230301Response) SetEmploymentType(v string)`

SetEmploymentType sets EmploymentType field to given value.


### SetEmploymentTypeNil

`func (o *Employee20230301Response) SetEmploymentTypeNil(b bool)`

 SetEmploymentTypeNil sets the value for EmploymentType to be an explicit nil

### UnsetEmploymentType
`func (o *Employee20230301Response) UnsetEmploymentType()`

UnsetEmploymentType ensures that no value is present for EmploymentType, not even an explicit nil
### GetStartDate

`func (o *Employee20230301Response) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Employee20230301Response) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Employee20230301Response) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### SetStartDateNil

`func (o *Employee20230301Response) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Employee20230301Response) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetRemoteCreatedAt

`func (o *Employee20230301Response) GetRemoteCreatedAt() string`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *Employee20230301Response) GetRemoteCreatedAtOk() (*string, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *Employee20230301Response) SetRemoteCreatedAt(v string)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.


### SetRemoteCreatedAtNil

`func (o *Employee20230301Response) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *Employee20230301Response) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetTerminationDate

`func (o *Employee20230301Response) GetTerminationDate() string`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *Employee20230301Response) GetTerminationDateOk() (*string, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *Employee20230301Response) SetTerminationDate(v string)`

SetTerminationDate sets TerminationDate field to given value.


### SetTerminationDateNil

`func (o *Employee20230301Response) SetTerminationDateNil(b bool)`

 SetTerminationDateNil sets the value for TerminationDate to be an explicit nil

### UnsetTerminationDate
`func (o *Employee20230301Response) UnsetTerminationDate()`

UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
### GetAvatar

`func (o *Employee20230301Response) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Employee20230301Response) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Employee20230301Response) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *Employee20230301Response) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *Employee20230301Response) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetHomeLocation

`func (o *Employee20230301Response) GetHomeLocation() Employee20230301ResponseHomeLocation`

GetHomeLocation returns the HomeLocation field if non-nil, zero value otherwise.

### GetHomeLocationOk

`func (o *Employee20230301Response) GetHomeLocationOk() (*Employee20230301ResponseHomeLocation, bool)`

GetHomeLocationOk returns a tuple with the HomeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeLocation

`func (o *Employee20230301Response) SetHomeLocation(v Employee20230301ResponseHomeLocation)`

SetHomeLocation sets HomeLocation field to given value.


### GetWorkLocation

`func (o *Employee20230301Response) GetWorkLocation() Employee20230301ResponseWorkLocation`

GetWorkLocation returns the WorkLocation field if non-nil, zero value otherwise.

### GetWorkLocationOk

`func (o *Employee20230301Response) GetWorkLocationOk() (*Employee20230301ResponseWorkLocation, bool)`

GetWorkLocationOk returns a tuple with the WorkLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLocation

`func (o *Employee20230301Response) SetWorkLocation(v Employee20230301ResponseWorkLocation)`

SetWorkLocation sets WorkLocation field to given value.


### SetWorkLocationNil

`func (o *Employee20230301Response) SetWorkLocationNil(b bool)`

 SetWorkLocationNil sets the value for WorkLocation to be an explicit nil

### UnsetWorkLocation
`func (o *Employee20230301Response) UnsetWorkLocation()`

UnsetWorkLocation ensures that no value is present for WorkLocation, not even an explicit nil
### GetManager

`func (o *Employee20230301Response) GetManager() Employee20230301ResponseManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *Employee20230301Response) GetManagerOk() (*Employee20230301ResponseManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *Employee20230301Response) SetManager(v Employee20230301ResponseManager)`

SetManager sets Manager field to given value.


### SetManagerNil

`func (o *Employee20230301Response) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *Employee20230301Response) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetBankAccount

`func (o *Employee20230301Response) GetBankAccount() Employee20230301ResponseBankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *Employee20230301Response) GetBankAccountOk() (*Employee20230301ResponseBankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *Employee20230301Response) SetBankAccount(v Employee20230301ResponseBankAccount)`

SetBankAccount sets BankAccount field to given value.


### SetBankAccountNil

`func (o *Employee20230301Response) SetBankAccountNil(b bool)`

 SetBankAccountNil sets the value for BankAccount to be an explicit nil

### UnsetBankAccount
`func (o *Employee20230301Response) UnsetBankAccount()`

UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
### GetEmployments

`func (o *Employee20230301Response) GetEmployments() []EmploymentResponse`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *Employee20230301Response) GetEmploymentsOk() (*[]EmploymentResponse, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *Employee20230301Response) SetEmployments(v []EmploymentResponse)`

SetEmployments sets Employments field to given value.


### SetEmploymentsNil

`func (o *Employee20230301Response) SetEmploymentsNil(b bool)`

 SetEmploymentsNil sets the value for Employments to be an explicit nil

### UnsetEmployments
`func (o *Employee20230301Response) UnsetEmployments()`

UnsetEmployments ensures that no value is present for Employments, not even an explicit nil
### GetCustomFields

`func (o *Employee20230301Response) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Employee20230301Response) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Employee20230301Response) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.


### SetCustomFieldsNil

`func (o *Employee20230301Response) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *Employee20230301Response) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetGroups

`func (o *Employee20230301Response) GetGroups() []Employee20230301ResponseGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Employee20230301Response) GetGroupsOk() (*[]Employee20230301ResponseGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Employee20230301Response) SetGroups(v []Employee20230301ResponseGroups)`

SetGroups sets Groups field to given value.


### SetGroupsNil

`func (o *Employee20230301Response) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *Employee20230301Response) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetCompany

`func (o *Employee20230301Response) GetCompany() Employee20230301ResponseCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Employee20230301Response) GetCompanyOk() (*Employee20230301ResponseCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Employee20230301Response) SetCompany(v Employee20230301ResponseCompany)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *Employee20230301Response) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Employee20230301Response) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


