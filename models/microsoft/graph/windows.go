package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Windows provides operations to manage the admin singleton.
type Windows struct {
    Entity
    // Entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
    updates Updatesable;
}
// NewWindows instantiates a new windows and sets the default values.
func NewWindows()(*Windows) {
    m := &Windows{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWindows(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["updates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateUpdatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdates(val.(Updatesable))
        }
        return nil
    }
    return res
}
// GetUpdates gets the updates property value. Entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *Windows) GetUpdates()(Updatesable) {
    if m == nil {
        return nil
    } else {
        return m.updates
    }
}
func (m *Windows) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetUpdates sets the updates property value. Entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *Windows) SetUpdates(value Updatesable)() {
    if m != nil {
        m.updates = value
    }
}
