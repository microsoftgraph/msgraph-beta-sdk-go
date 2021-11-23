package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// educationalActivity 
type EducationalActivity struct {
    ItemFacet
    // The month and year the user graduated or completed the activity.
    completionMonthYear *string;
    // The month and year the user completed the educational activity referenced.
    endMonthYear *string;
    // 
    institution *InstitutionData;
    // 
    program *EducationalActivityDetail;
    // The month and year the user commenced the activity referenced.
    startMonthYear *string;
}
// NewEducationalActivity instantiates a new educationalActivity and sets the default values.
func NewEducationalActivity()(*EducationalActivity) {
    m := &EducationalActivity{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// GetCompletionMonthYear gets the completionMonthYear property value. The month and year the user graduated or completed the activity.
func (m *EducationalActivity) GetCompletionMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.completionMonthYear
    }
}
// GetEndMonthYear gets the endMonthYear property value. The month and year the user completed the educational activity referenced.
func (m *EducationalActivity) GetEndMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endMonthYear
    }
}
// GetInstitution gets the institution property value. 
func (m *EducationalActivity) GetInstitution()(*InstitutionData) {
    if m == nil {
        return nil
    } else {
        return m.institution
    }
}
// GetProgram gets the program property value. 
func (m *EducationalActivity) GetProgram()(*EducationalActivityDetail) {
    if m == nil {
        return nil
    } else {
        return m.program
    }
}
// GetStartMonthYear gets the startMonthYear property value. The month and year the user commenced the activity referenced.
func (m *EducationalActivity) GetStartMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startMonthYear
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationalActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["completionMonthYear"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionMonthYear(val)
        }
        return nil
    }
    res["endMonthYear"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndMonthYear(val)
        }
        return nil
    }
    res["institution"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInstitutionData() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstitution(val.(*InstitutionData))
        }
        return nil
    }
    res["program"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationalActivityDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgram(val.(*EducationalActivityDetail))
        }
        return nil
    }
    res["startMonthYear"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMonthYear(val)
        }
        return nil
    }
    return res
}
func (m *EducationalActivity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCompletionMonthYear sets the completionMonthYear property value. The month and year the user graduated or completed the activity.
func (m *EducationalActivity) SetCompletionMonthYear(value *string)() {
    m.completionMonthYear = value
}
// SetEndMonthYear sets the endMonthYear property value. The month and year the user completed the educational activity referenced.
func (m *EducationalActivity) SetEndMonthYear(value *string)() {
    m.endMonthYear = value
}
// SetInstitution sets the institution property value. 
func (m *EducationalActivity) SetInstitution(value *InstitutionData)() {
    m.institution = value
}
// SetProgram sets the program property value. 
func (m *EducationalActivity) SetProgram(value *EducationalActivityDetail)() {
    m.program = value
}
// SetStartMonthYear sets the startMonthYear property value. The month and year the user commenced the activity referenced.
func (m *EducationalActivity) SetStartMonthYear(value *string)() {
    m.startMonthYear = value
}
