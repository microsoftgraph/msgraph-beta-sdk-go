package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AllowedValue struct {
    Entity
    // 
    isActive *bool;
}
// Instantiates a new allowedValue and sets the default values.
func NewAllowedValue()(*AllowedValue) {
    m := &AllowedValue{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the isActive property value. 
func (m *AllowedValue) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// The deserialization information for the current model
func (m *AllowedValue) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsActive(val)
        return nil
    }
    return res
}
func (m *AllowedValue) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AllowedValue) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the isActive property value. 
// Parameters:
//  - value : Value to set for the isActive property.
func (m *AllowedValue) SetIsActive(value *bool)() {
    m.isActive = value
}
