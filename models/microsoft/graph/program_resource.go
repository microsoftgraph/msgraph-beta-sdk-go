package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProgramResource 
type ProgramResource struct {
    Identity
    // Type of the resource, indicating whether it is a group or an app.
    type_escaped *string;
}
// NewProgramResource instantiates a new programResource and sets the default values.
func NewProgramResource()(*ProgramResource) {
    m := &ProgramResource{
        Identity: *NewIdentity(),
    }
    return m
}
// GetType gets the type property value. Type of the resource, indicating whether it is a group or an app.
func (m *ProgramResource) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProgramResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
func (m *ProgramResource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ProgramResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetType sets the type property value. Type of the resource, indicating whether it is a group or an app.
func (m *ProgramResource) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
