package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationalActivity struct {
    ItemFacet
    completionMonthYear *string;
    endMonthYear *string;
    institution *InstitutionData;
    program *EducationalActivityDetail;
    startMonthYear *string;
}
func NewEducationalActivity()(*EducationalActivity) {
    m := &EducationalActivity{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
func (m *EducationalActivity) GetCompletionMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.completionMonthYear
    }
}
func (m *EducationalActivity) GetEndMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endMonthYear
    }
}
func (m *EducationalActivity) GetInstitution()(*InstitutionData) {
    if m == nil {
        return nil
    } else {
        return m.institution
    }
}
func (m *EducationalActivity) GetProgram()(*EducationalActivityDetail) {
    if m == nil {
        return nil
    } else {
        return m.program
    }
}
func (m *EducationalActivity) GetStartMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startMonthYear
    }
}
func (m *EducationalActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["completionMonthYear"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCompletionMonthYear(val)
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
    res["institution"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInstitutionData() })
        if err != nil {
            return err
        }
        m.SetInstitution(val.(*InstitutionData))
        return nil
    }
    res["program"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationalActivityDetail() })
        if err != nil {
            return err
        }
        m.SetProgram(val.(*EducationalActivityDetail))
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
    return res
}
func (m *EducationalActivity) IsNil()(bool) {
    return m == nil
}
func (m *EducationalActivity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("completionMonthYear", m.GetCompletionMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("endMonthYear", m.GetEndMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("institution", m.GetInstitution())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("program", m.GetProgram())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("startMonthYear", m.GetStartMonthYear())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *EducationalActivity) SetCompletionMonthYear(value *string)() {
    m.completionMonthYear = value
}
func (m *EducationalActivity) SetEndMonthYear(value *string)() {
    m.endMonthYear = value
}
func (m *EducationalActivity) SetInstitution(value *InstitutionData)() {
    m.institution = value
}
func (m *EducationalActivity) SetProgram(value *EducationalActivityDetail)() {
    m.program = value
}
func (m *EducationalActivity) SetStartMonthYear(value *string)() {
    m.startMonthYear = value
}
