package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Employeeable 
type Employeeable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddress()(PostalAddressTypeable)
    GetBirthDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetEmploymentDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetGivenName()(*string)
    GetJobTitle()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMiddleName()(*string)
    GetMobilePhone()(*string)
    GetNumber()(*string)
    GetPersonalEmail()(*string)
    GetPhoneNumber()(*string)
    GetPicture()([]Pictureable)
    GetStatisticsGroupCode()(*string)
    GetStatus()(*string)
    GetSurname()(*string)
    GetTerminationDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    SetAddress(value PostalAddressTypeable)()
    SetBirthDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetEmploymentDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetGivenName(value *string)()
    SetJobTitle(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMiddleName(value *string)()
    SetMobilePhone(value *string)()
    SetNumber(value *string)()
    SetPersonalEmail(value *string)()
    SetPhoneNumber(value *string)()
    SetPicture(value []Pictureable)()
    SetStatisticsGroupCode(value *string)()
    SetStatus(value *string)()
    SetSurname(value *string)()
    SetTerminationDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
}
