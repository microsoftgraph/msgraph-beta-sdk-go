package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Windows struct {
    Entity
    // Entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
    updates *Updates;
}
// Instantiates a new windows and sets the default values.
func NewWindows()(*Windows) {
    m := &Windows{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the updates property value. Entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *Windows) GetUpdates()(*Updates) {
    if m == nil {
        return nil
    } else {
        return m.updates
    }
}
// The deserialization information for the current model
func (m *Windows) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["updates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdates() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdates(val.(*Updates))
        }
        return nil
    }
    return res
}
func (m *Windows) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Windows) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("updates", m.GetUpdates())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the updates property value. Entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
// Parameters:
//  - value : Value to set for the updates property.
func (m *Windows) SetUpdates(value *Updates)() {
    m.updates = value
}
