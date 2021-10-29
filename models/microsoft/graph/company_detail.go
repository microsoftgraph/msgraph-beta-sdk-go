package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new companyDetail and sets the default values.
func NewCompanyDetail()(*CompanyDetail) {
    m := &CompanyDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompanyDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the address property value. Address of the company.
func (m *CompanyDetail) GetAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// Gets the department property value. Department Name within a company.
func (m *CompanyDetail) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// Gets the displayName property value. Company name.
func (m *CompanyDetail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the officeLocation property value. Office Location of the person referred to.
func (m *CompanyDetail) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// Gets the pronunciation property value. Pronunciation guide for the company name.
func (m *CompanyDetail) GetPronunciation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pronunciation
    }
}
// Gets the webUrl property value. Link to the company home page.
func (m *CompanyDetail) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *CompanyDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        m.SetAddress(val.(*PhysicalAddress))
        return nil
    }
    res["department"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDepartment(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["officeLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOfficeLocation(val)
        return nil
    }
    res["pronunciation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPronunciation(val)
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *CompanyDetail) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CompanyDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the address property value. Address of the company.
// Parameters:
//  - value : Value to set for the address property.
func (m *CompanyDetail) SetAddress(value *PhysicalAddress)() {
    m.address = value
}
// Sets the department property value. Department Name within a company.
// Parameters:
//  - value : Value to set for the department property.
func (m *CompanyDetail) SetDepartment(value *string)() {
    m.department = value
}
// Sets the displayName property value. Company name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CompanyDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the officeLocation property value. Office Location of the person referred to.
// Parameters:
//  - value : Value to set for the officeLocation property.
func (m *CompanyDetail) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// Sets the pronunciation property value. Pronunciation guide for the company name.
// Parameters:
//  - value : Value to set for the pronunciation property.
func (m *CompanyDetail) SetPronunciation(value *string)() {
    m.pronunciation = value
}
// Sets the webUrl property value. Link to the company home page.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *CompanyDetail) SetWebUrl(value *string)() {
    m.webUrl = value
}
