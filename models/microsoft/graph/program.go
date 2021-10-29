package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Program struct {
    Entity
    // Controls associated with the program.
    controls []ProgramControl;
    // The description of the program.
    description *string;
    // The name of the program.  Required on create.
    displayName *string;
}
// Instantiates a new program and sets the default values.
func NewProgram()(*Program) {
    m := &Program{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the controls property value. Controls associated with the program.
func (m *Program) GetControls()([]ProgramControl) {
    if m == nil {
        return nil
    } else {
        return m.controls
    }
}
// Gets the description property value. The description of the program.
func (m *Program) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The name of the program.  Required on create.
func (m *Program) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// The deserialization information for the current model
func (m *Program) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["controls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProgramControl() })
        if err != nil {
            return err
        }
        res := make([]ProgramControl, len(val))
        for i, v := range val {
            res[i] = *(v.(*ProgramControl))
        }
        m.SetControls(res)
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    return res
}
func (m *Program) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Program) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetControls()))
        for i, v := range m.GetControls() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("controls", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    return nil
}
// Sets the controls property value. Controls associated with the program.
// Parameters:
//  - value : Value to set for the controls property.
func (m *Program) SetControls(value []ProgramControl)() {
    m.controls = value
}
// Sets the description property value. The description of the program.
// Parameters:
//  - value : Value to set for the description property.
func (m *Program) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The name of the program.  Required on create.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Program) SetDisplayName(value *string)() {
    m.displayName = value
}
