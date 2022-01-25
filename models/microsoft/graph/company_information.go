package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CompanyInformation 
type CompanyInformation struct {
    Entity
    // 
    address *PostalAddressType;
    // 
    currencyCode *string;
    // 
    currentFiscalYearStartDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
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
// NewCompanyInformation instantiates a new companyInformation and sets the default values.
func NewCompanyInformation()(*CompanyInformation) {
    m := &CompanyInformation{
        Entity: *NewEntity(),
    }
    return m
}
// GetAddress gets the address property value. 
func (m *CompanyInformation) GetAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetCurrencyCode gets the currencyCode property value. 
func (m *CompanyInformation) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// GetCurrentFiscalYearStartDate gets the currentFiscalYearStartDate property value. 
func (m *CompanyInformation) GetCurrentFiscalYearStartDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.currentFiscalYearStartDate
    }
}
// GetDisplayName gets the displayName property value. 
func (m *CompanyInformation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEmail gets the email property value. 
func (m *CompanyInformation) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetFaxNumber gets the faxNumber property value. 
func (m *CompanyInformation) GetFaxNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.faxNumber
    }
}
// GetIndustry gets the industry property value. 
func (m *CompanyInformation) GetIndustry()(*string) {
    if m == nil {
        return nil
    } else {
        return m.industry
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *CompanyInformation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPhoneNumber gets the phoneNumber property value. 
func (m *CompanyInformation) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// GetPicture gets the picture property value. 
func (m *CompanyInformation) GetPicture()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.picture
    }
}
// GetTaxRegistrationNumber gets the taxRegistrationNumber property value. 
func (m *CompanyInformation) GetTaxRegistrationNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.taxRegistrationNumber
    }
}
// GetWebsite gets the website property value. 
func (m *CompanyInformation) GetWebsite()(*string) {
    if m == nil {
        return nil
    } else {
        return m.website
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
        val, err := n.GetDateOnlyValue()
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
// Serialize serializes information the current object
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
        err = writer.WriteDateOnlyValue("currentFiscalYearStartDate", m.GetCurrentFiscalYearStartDate())
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
// SetAddress sets the address property value. 
func (m *CompanyInformation) SetAddress(value *PostalAddressType)() {
    if m != nil {
        m.address = value
    }
}
// SetCurrencyCode sets the currencyCode property value. 
func (m *CompanyInformation) SetCurrencyCode(value *string)() {
    if m != nil {
        m.currencyCode = value
    }
}
// SetCurrentFiscalYearStartDate sets the currentFiscalYearStartDate property value. 
func (m *CompanyInformation) SetCurrentFiscalYearStartDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.currentFiscalYearStartDate = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *CompanyInformation) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEmail sets the email property value. 
func (m *CompanyInformation) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetFaxNumber sets the faxNumber property value. 
func (m *CompanyInformation) SetFaxNumber(value *string)() {
    if m != nil {
        m.faxNumber = value
    }
}
// SetIndustry sets the industry property value. 
func (m *CompanyInformation) SetIndustry(value *string)() {
    if m != nil {
        m.industry = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *CompanyInformation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPhoneNumber sets the phoneNumber property value. 
func (m *CompanyInformation) SetPhoneNumber(value *string)() {
    if m != nil {
        m.phoneNumber = value
    }
}
// SetPicture sets the picture property value. 
func (m *CompanyInformation) SetPicture(value []byte)() {
    if m != nil {
        m.picture = value
    }
}
// SetTaxRegistrationNumber sets the taxRegistrationNumber property value. 
func (m *CompanyInformation) SetTaxRegistrationNumber(value *string)() {
    if m != nil {
        m.taxRegistrationNumber = value
    }
}
// SetWebsite sets the website property value. 
func (m *CompanyInformation) SetWebsite(value *string)() {
    if m != nil {
        m.website = value
    }
}
