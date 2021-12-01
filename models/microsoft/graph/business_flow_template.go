package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BusinessFlowTemplate 
type BusinessFlowTemplate struct {
    Entity
    // The name of the business flow template
    displayName *string;
}
// NewBusinessFlowTemplate instantiates a new businessFlowTemplate and sets the default values.
func NewBusinessFlowTemplate()(*BusinessFlowTemplate) {
    m := &BusinessFlowTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The name of the business flow template
func (m *BusinessFlowTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessFlowTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
func (m *BusinessFlowTemplate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BusinessFlowTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the business flow template
func (m *BusinessFlowTemplate) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
