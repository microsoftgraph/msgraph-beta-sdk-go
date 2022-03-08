package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CompanyInformationable 
type CompanyInformationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddress()(PostalAddressTypeable)
    GetCurrencyCode()(*string)
    GetCurrentFiscalYearStartDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetFaxNumber()(*string)
    GetIndustry()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPhoneNumber()(*string)
    GetPicture()([]byte)
    GetTaxRegistrationNumber()(*string)
    GetWebsite()(*string)
    SetAddress(value PostalAddressTypeable)()
    SetCurrencyCode(value *string)()
    SetCurrentFiscalYearStartDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetFaxNumber(value *string)()
    SetIndustry(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPhoneNumber(value *string)()
    SetPicture(value []byte)()
    SetTaxRegistrationNumber(value *string)()
    SetWebsite(value *string)()
}
