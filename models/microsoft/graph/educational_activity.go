package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new educationalActivity and sets the default values.
func NewEducationalActivity()(*EducationalActivity) {
    m := &EducationalActivity{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the completionMonthYear property value. The month and year the user graduated or completed the activity.
func (m *EducationalActivity) GetCompletionMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.completionMonthYear
    }
}
// Gets the endMonthYear property value. The month and year the user completed the educational activity referenced.
func (m *EducationalActivity) GetEndMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endMonthYear
    }
}
// Gets the institution property value. 
func (m *EducationalActivity) GetInstitution()(*InstitutionData) {
    if m == nil {
        return nil
    } else {
        return m.institution
    }
}
// Gets the program property value. 
func (m *EducationalActivity) GetProgram()(*EducationalActivityDetail) {
    if m == nil {
        return nil
    } else {
        return m.program
    }
}
// Gets the startMonthYear property value. The month and year the user commenced the activity referenced.
func (m *EducationalActivity) GetStartMonthYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startMonthYear
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the completionMonthYear property value. The month and year the user graduated or completed the activity.
// Parameters:
//  - value : Value to set for the completionMonthYear property.
func (m *EducationalActivity) SetCompletionMonthYear(value *string)() {
    m.completionMonthYear = value
}
// Sets the endMonthYear property value. The month and year the user completed the educational activity referenced.
// Parameters:
//  - value : Value to set for the endMonthYear property.
func (m *EducationalActivity) SetEndMonthYear(value *string)() {
    m.endMonthYear = value
}
// Sets the institution property value. 
// Parameters:
//  - value : Value to set for the institution property.
func (m *EducationalActivity) SetInstitution(value *InstitutionData)() {
    m.institution = value
}
// Sets the program property value. 
// Parameters:
//  - value : Value to set for the program property.
func (m *EducationalActivity) SetProgram(value *EducationalActivityDetail)() {
    m.program = value
}
// Sets the startMonthYear property value. The month and year the user commenced the activity referenced.
// Parameters:
//  - value : Value to set for the startMonthYear property.
func (m *EducationalActivity) SetStartMonthYear(value *string)() {
    m.startMonthYear = value
}
