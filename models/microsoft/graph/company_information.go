package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CompanyInformation struct {
    Entity
    // 
    address *PostalAddressType;
    // 
    currencyCode *string;
    // 
    currentFiscalYearStartDate *string;
    // 
    displayName *string;
    // 
    email *string;
    // 
    faxNumber *string;
    // 
    industry *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    phoneNumber *string;
    // 
    picture []byte;
    // 
    taxRegistrationNumber *string;
    // 
    website *string;
}
// Instantiates a new companyInformation and sets the default values.
func NewCompanyInformation()(*CompanyInformation) {
    m := &CompanyInformation{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the address property value. 
func (m *CompanyInformation) GetAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// Gets the currencyCode property value. 
func (m *CompanyInformation) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// Gets the currentFiscalYearStartDate property value. 
func (m *CompanyInformation) GetCurrentFiscalYearStartDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currentFiscalYearStartDate
    }
}
// Gets the displayName property value. 
func (m *CompanyInformation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the email property value. 
func (m *CompanyInformation) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the faxNumber property value. 
func (m *CompanyInformation) GetFaxNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.faxNumber
    }
}
// Gets the industry property value. 
func (m *CompanyInformation) GetIndustry()(*string) {
    if m == nil {
        return nil
    } else {
        return m.industry
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *CompanyInformation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the phoneNumber property value. 
func (m *CompanyInformation) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// Gets the picture property value. 
func (m *CompanyInformation) GetPicture()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.picture
    }
}
// Gets the taxRegistrationNumber property value. 
func (m *CompanyInformation) GetTaxRegistrationNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxRegistrationNumber
    }
}
// Gets the website property value. 
func (m *CompanyInformation) GetWebsite()(*string) {
    if m == nil {
        return nil
    } else {
        return m.website
    }
}
// The deserialization information for the current model
func (m *CompanyInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(*PostalAddressType))
        }
        return nil
    }
    res["currencyCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyCode(val)
        }
        return nil
    }
    res["currentFiscalYearStartDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentFiscalYearStartDate(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["faxNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaxNumber(val)
        }
        return nil
    }
    res["industry"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndustry(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["picture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPicture(val)
        }
        return nil
    }
    res["taxRegistrationNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxRegistrationNumber(val)
        }
        return nil
    }
    res["website"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebsite(val)
        }
        return nil
    }
    return res
}
func (m *CompanyInformation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CompanyInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("currencyCode", m.GetCurrencyCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("currentFiscalYearStartDate", m.GetCurrentFiscalYearStartDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("faxNumber", m.GetFaxNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("industry", m.GetIndustry())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("picture", m.GetPicture())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taxRegistrationNumber", m.GetTaxRegistrationNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("website", m.GetWebsite())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the address property value. 
// Parameters:
//  - value : Value to set for the address property.
func (m *CompanyInformation) SetAddress(value *PostalAddressType)() {
    m.address = value
}
// Sets the currencyCode property value. 
// Parameters:
//  - value : Value to set for the currencyCode property.
func (m *CompanyInformation) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// Sets the currentFiscalYearStartDate property value. 
// Parameters:
//  - value : Value to set for the currentFiscalYearStartDate property.
func (m *CompanyInformation) SetCurrentFiscalYearStartDate(value *string)() {
    m.currentFiscalYearStartDate = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CompanyInformation) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the email property value. 
// Parameters:
//  - value : Value to set for the email property.
func (m *CompanyInformation) SetEmail(value *string)() {
    m.email = value
}
// Sets the faxNumber property value. 
// Parameters:
//  - value : Value to set for the faxNumber property.
func (m *CompanyInformation) SetFaxNumber(value *string)() {
    m.faxNumber = value
}
// Sets the industry property value. 
// Parameters:
//  - value : Value to set for the industry property.
func (m *CompanyInformation) SetIndustry(value *string)() {
    m.industry = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *CompanyInformation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the phoneNumber property value. 
// Parameters:
//  - value : Value to set for the phoneNumber property.
func (m *CompanyInformation) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// Sets the picture property value. 
// Parameters:
//  - value : Value to set for the picture property.
func (m *CompanyInformation) SetPicture(value []byte)() {
    m.picture = value
}
// Sets the taxRegistrationNumber property value. 
// Parameters:
//  - value : Value to set for the taxRegistrationNumber property.
func (m *CompanyInformation) SetTaxRegistrationNumber(value *string)() {
    m.taxRegistrationNumber = value
}
// Sets the website property value. 
// Parameters:
//  - value : Value to set for the website property.
func (m *CompanyInformation) SetWebsite(value *string)() {
    m.website = value
}
