package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Security 
type Security struct {
    Entity
    // 
    informationProtection *InformationProtection;
}
// NewSecurity instantiates a new security and sets the default values.
func NewSecurity()(*Security) {
    m := &Security{
        Entity: *NewEntity(),
    }
    return m
}
// GetInformationProtection gets the informationProtection property value. 
func (m *Security) GetInformationProtection()(*InformationProtection) {
    if m == nil {
        return nil
    } else {
        return m.informationProtection
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Security) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["informationProtection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInformationProtection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInformationProtection(val.(*InformationProtection))
        }
        return nil
    }
    return res
}
func (m *Security) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Security) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("informationProtection", m.GetInformationProtection())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInformationProtection sets the informationProtection property value. 
func (m *Security) SetInformationProtection(value *InformationProtection)() {
    if m != nil {
        m.informationProtection = value
    }
}
