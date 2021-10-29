package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PositionDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Detail about the company or employer.
    company *CompanyDetail;
    // Description of the position in question.
    description *string;
    // When the position ended.
    endMonthYear *string;
    // The title held when in that position.
    jobTitle *string;
    // The role the position entailed.
    role *string;
    // The start month and year of the position.
    startMonthYear *string;
    // Short summary of the position.
    summary *string;
}
// Instantiates a new positionDetail and sets the default values.
func NewPositionDetail()(*PositionDetail) {
    m := &PositionDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PositionDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the company property value. Detail about the company or employer.
func (m *PositionDetail) GetCompany()(*CompanyDetail) {
    if m == nil {
        return nil
    } else {
        return m.company
    }
}
// Gets the description property value. Description of the position in question.
func (m *PositionDetail) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the endMonthYear property value. When the position ended.
func (m *PositionDetail) GetEndMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endMonthYear
    }
}
// Gets the jobTitle property value. The title held when in that position.
func (m *PositionDetail) GetJobTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobTitle
    }
}
// Gets the role property value. The role the position entailed.
func (m *PositionDetail) GetRole()(*string) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
// Gets the startMonthYear property value. The start month and year of the position.
func (m *PositionDetail) GetStartMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startMonthYear
    }
}
// Gets the summary property value. Short summary of the position.
func (m *PositionDetail) GetSummary()(*string) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
// The deserialization information for the current model
func (m *PositionDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["company"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCompanyDetail() })
        if err != nil {
            return err
        }
        m.SetCompany(val.(*CompanyDetail))
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["endMonthYear"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEndMonthYear(val)
        return nil
    }
    res["jobTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJobTitle(val)
        return nil
    }
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRole(val)
        return nil
    }
    res["startMonthYear"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStartMonthYear(val)
        return nil
    }
    res["summary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSummary(val)
        return nil
    }
    return res
}
func (m *PositionDetail) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PositionDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("company", m.GetCompany())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("endMonthYear", m.GetEndMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jobTitle", m.GetJobTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("role", m.GetRole())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("startMonthYear", m.GetStartMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("summary", m.GetSummary())
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
func (m *PositionDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the company property value. Detail about the company or employer.
// Parameters:
//  - value : Value to set for the company property.
func (m *PositionDetail) SetCompany(value *CompanyDetail)() {
    m.company = value
}
// Sets the description property value. Description of the position in question.
// Parameters:
//  - value : Value to set for the description property.
func (m *PositionDetail) SetDescription(value *string)() {
    m.description = value
}
// Sets the endMonthYear property value. When the position ended.
// Parameters:
//  - value : Value to set for the endMonthYear property.
func (m *PositionDetail) SetEndMonthYear(value *string)() {
    m.endMonthYear = value
}
// Sets the jobTitle property value. The title held when in that position.
// Parameters:
//  - value : Value to set for the jobTitle property.
func (m *PositionDetail) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// Sets the role property value. The role the position entailed.
// Parameters:
//  - value : Value to set for the role property.
func (m *PositionDetail) SetRole(value *string)() {
    m.role = value
}
// Sets the startMonthYear property value. The start month and year of the position.
// Parameters:
//  - value : Value to set for the startMonthYear property.
func (m *PositionDetail) SetStartMonthYear(value *string)() {
    m.startMonthYear = value
}
// Sets the summary property value. Short summary of the position.
// Parameters:
//  - value : Value to set for the summary property.
func (m *PositionDetail) SetSummary(value *string)() {
    m.summary = value
}
