package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CompanyDetail 
type CompanyDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Address of the company.
    address *PhysicalAddress;
    // Department Name within a company.
    department *string;
    // Company name.
    displayName *string;
    // Office Location of the person referred to.
    officeLocation *string;
    // Pronunciation guide for the company name.
    pronunciation *string;
    // Link to the company home page.
    webUrl *string;
}
// NewCompanyDetail instantiates a new companyDetail and sets the default values.
func NewCompanyDetail()(*CompanyDetail) {
    m := &CompanyDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompanyDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddress gets the address property value. Address of the company.
func (m *CompanyDetail) GetAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetDepartment gets the department property value. Department Name within a company.
func (m *CompanyDetail) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// GetDisplayName gets the displayName property value. Company name.
func (m *CompanyDetail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetOfficeLocation gets the officeLocation property value. Office Location of the person referred to.
func (m *CompanyDetail) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// GetPronunciation gets the pronunciation property value. Pronunciation guide for the company name.
func (m *CompanyDetail) GetPronunciation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pronunciation
    }
}
// GetWebUrl gets the webUrl property value. Link to the company home page.
func (m *CompanyDetail) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CompanyDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(*PhysicalAddress))
        }
        return nil
    }
    res["department"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val)
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
    res["officeLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeLocation(val)
        }
        return nil
    }
    res["pronunciation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPronunciation(val)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
func (m *CompanyDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CompanyDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("department", m.GetDepartment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pronunciation", m.GetPronunciation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompanyDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddress sets the address property value. Address of the company.
func (m *CompanyDetail) SetAddress(value *PhysicalAddress)() {
    if m != nil {
        m.address = value
    }
}
// SetDepartment sets the department property value. Department Name within a company.
func (m *CompanyDetail) SetDepartment(value *string)() {
    if m != nil {
        m.department = value
    }
}
// SetDisplayName sets the displayName property value. Company name.
func (m *CompanyDetail) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetOfficeLocation sets the officeLocation property value. Office Location of the person referred to.
func (m *CompanyDetail) SetOfficeLocation(value *string)() {
    if m != nil {
        m.officeLocation = value
    }
}
// SetPronunciation sets the pronunciation property value. Pronunciation guide for the company name.
func (m *CompanyDetail) SetPronunciation(value *string)() {
    if m != nil {
        m.pronunciation = value
    }
}
// SetWebUrl sets the webUrl property value. Link to the company home page.
func (m *CompanyDetail) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
