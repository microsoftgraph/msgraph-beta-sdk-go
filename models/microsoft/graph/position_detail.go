package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PositionDetail struct {
    additionalData map[string]interface{};
    company *CompanyDetail;
    description *string;
    endMonthYear *string;
    jobTitle *string;
    role *string;
    startMonthYear *string;
    summary *string;
}
func NewPositionDetail()(*PositionDetail) {
    m := &PositionDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PositionDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PositionDetail) GetCompany()(*CompanyDetail) {
    if m == nil {
        return nil
    } else {
        return m.company
    }
}
func (m *PositionDetail) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *PositionDetail) GetEndMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endMonthYear
    }
}
func (m *PositionDetail) GetJobTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobTitle
    }
}
func (m *PositionDetail) GetRole()(*string) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
func (m *PositionDetail) GetStartMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startMonthYear
    }
}
func (m *PositionDetail) GetSummary()(*string) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
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
func (m *PositionDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PositionDetail) SetCompany(value *CompanyDetail)() {
    m.company = value
}
func (m *PositionDetail) SetDescription(value *string)() {
    m.description = value
}
func (m *PositionDetail) SetEndMonthYear(value *string)() {
    m.endMonthYear = value
}
func (m *PositionDetail) SetJobTitle(value *string)() {
    m.jobTitle = value
}
func (m *PositionDetail) SetRole(value *string)() {
    m.role = value
}
func (m *PositionDetail) SetStartMonthYear(value *string)() {
    m.startMonthYear = value
}
func (m *PositionDetail) SetSummary(value *string)() {
    m.summary = value
}
