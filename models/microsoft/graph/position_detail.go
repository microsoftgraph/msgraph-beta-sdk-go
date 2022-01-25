package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PositionDetail 
type PositionDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Detail about the company or employer.
    company *CompanyDetail;
    // Description of the position in question.
    description *string;
    // When the position ended.
    endMonthYear *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The title held when in that position.
    jobTitle *string;
    // The role the position entailed.
    role *string;
    // The start month and year of the position.
    startMonthYear *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // Short summary of the position.
    summary *string;
}
// NewPositionDetail instantiates a new positionDetail and sets the default values.
func NewPositionDetail()(*PositionDetail) {
    m := &PositionDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PositionDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCompany gets the company property value. Detail about the company or employer.
func (m *PositionDetail) GetCompany()(*CompanyDetail) {
    if m == nil {
        return nil
    } else {
        return m.company
    }
}
// GetDescription gets the description property value. Description of the position in question.
func (m *PositionDetail) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetEndMonthYear gets the endMonthYear property value. When the position ended.
func (m *PositionDetail) GetEndMonthYear()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.endMonthYear
    }
}
// GetJobTitle gets the jobTitle property value. The title held when in that position.
func (m *PositionDetail) GetJobTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobTitle
    }
}
// GetRole gets the role property value. The role the position entailed.
func (m *PositionDetail) GetRole()(*string) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
// GetStartMonthYear gets the startMonthYear property value. The start month and year of the position.
func (m *PositionDetail) GetStartMonthYear()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.startMonthYear
    }
}
// GetSummary gets the summary property value. Short summary of the position.
func (m *PositionDetail) GetSummary()(*string) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PositionDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["company"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCompanyDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompany(val.(*CompanyDetail))
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["endMonthYear"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndMonthYear(val)
        }
        return nil
    }
    res["jobTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
        }
        return nil
    }
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val)
        }
        return nil
    }
    res["startMonthYear"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMonthYear(val)
        }
        return nil
    }
    res["summary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
        }
        return nil
    }
    return res
}
func (m *PositionDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err := writer.WriteDateOnlyValue("endMonthYear", m.GetEndMonthYear())
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
        err := writer.WriteDateOnlyValue("startMonthYear", m.GetStartMonthYear())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PositionDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCompany sets the company property value. Detail about the company or employer.
func (m *PositionDetail) SetCompany(value *CompanyDetail)() {
    if m != nil {
        m.company = value
    }
}
// SetDescription sets the description property value. Description of the position in question.
func (m *PositionDetail) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetEndMonthYear sets the endMonthYear property value. When the position ended.
func (m *PositionDetail) SetEndMonthYear(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.endMonthYear = value
    }
}
// SetJobTitle sets the jobTitle property value. The title held when in that position.
func (m *PositionDetail) SetJobTitle(value *string)() {
    if m != nil {
        m.jobTitle = value
    }
}
// SetRole sets the role property value. The role the position entailed.
func (m *PositionDetail) SetRole(value *string)() {
    if m != nil {
        m.role = value
    }
}
// SetStartMonthYear sets the startMonthYear property value. The start month and year of the position.
func (m *PositionDetail) SetStartMonthYear(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.startMonthYear = value
    }
}
// SetSummary sets the summary property value. Short summary of the position.
func (m *PositionDetail) SetSummary(value *string)() {
    if m != nil {
        m.summary = value
    }
}
