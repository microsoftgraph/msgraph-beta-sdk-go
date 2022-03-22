package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationalActivity 
type EducationalActivity struct {
    ItemFacet
    // The month and year the user graduated or completed the activity.
    completionMonthYear *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The month and year the user completed the educational activity referenced.
    endMonthYear *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // 
    institution InstitutionDataable;
    // 
    program EducationalActivityDetailable;
    // The month and year the user commenced the activity referenced.
    startMonthYear *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
}
// NewEducationalActivity instantiates a new educationalActivity and sets the default values.
func NewEducationalActivity()(*EducationalActivity) {
    m := &EducationalActivity{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// CreateEducationalActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationalActivityFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEducationalActivity(), nil
}
// GetCompletionMonthYear gets the completionMonthYear property value. The month and year the user graduated or completed the activity.
func (m *EducationalActivity) GetCompletionMonthYear()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.completionMonthYear
    }
}
// GetEndMonthYear gets the endMonthYear property value. The month and year the user completed the educational activity referenced.
func (m *EducationalActivity) GetEndMonthYear()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.endMonthYear
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationalActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["completionMonthYear"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionMonthYear(val)
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
    res["institution"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateInstitutionDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstitution(val.(InstitutionDataable))
        }
        return nil
    }
    res["program"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationalActivityDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgram(val.(EducationalActivityDetailable))
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
    return res
}
// GetInstitution gets the institution property value. 
func (m *EducationalActivity) GetInstitution()(InstitutionDataable) {
    if m == nil {
        return nil
    } else {
        return m.institution
    }
}
// GetProgram gets the program property value. 
func (m *EducationalActivity) GetProgram()(EducationalActivityDetailable) {
    if m == nil {
        return nil
    } else {
        return m.program
    }
}
// GetStartMonthYear gets the startMonthYear property value. The month and year the user commenced the activity referenced.
func (m *EducationalActivity) GetStartMonthYear()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.startMonthYear
    }
}
// Serialize serializes information the current object
func (m *EducationalActivity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteDateOnlyValue("completionMonthYear", m.GetCompletionMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("endMonthYear", m.GetEndMonthYear())
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
        err = writer.WriteDateOnlyValue("startMonthYear", m.GetStartMonthYear())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletionMonthYear sets the completionMonthYear property value. The month and year the user graduated or completed the activity.
func (m *EducationalActivity) SetCompletionMonthYear(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.completionMonthYear = value
    }
}
// SetEndMonthYear sets the endMonthYear property value. The month and year the user completed the educational activity referenced.
func (m *EducationalActivity) SetEndMonthYear(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.endMonthYear = value
    }
}
// SetInstitution sets the institution property value. 
func (m *EducationalActivity) SetInstitution(value InstitutionDataable)() {
    if m != nil {
        m.institution = value
    }
}
// SetProgram sets the program property value. 
func (m *EducationalActivity) SetProgram(value EducationalActivityDetailable)() {
    if m != nil {
        m.program = value
    }
}
// SetStartMonthYear sets the startMonthYear property value. The month and year the user commenced the activity referenced.
func (m *EducationalActivity) SetStartMonthYear(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.startMonthYear = value
    }
}
