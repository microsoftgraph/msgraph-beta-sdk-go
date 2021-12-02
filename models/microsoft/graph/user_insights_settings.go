package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserInsightsSettings 
type UserInsightsSettings struct {
    Entity
    // true if user's itemInsights and meeting hours insights are enabled; false if user's itemInsights and meeting hours insights are disabled. Default is true. Optional.
    isEnabled *bool;
}
// NewUserInsightsSettings instantiates a new userInsightsSettings and sets the default values.
func NewUserInsightsSettings()(*UserInsightsSettings) {
    m := &UserInsightsSettings{
        Entity: *NewEntity(),
    }
    return m
}
// GetIsEnabled gets the isEnabled property value. true if user's itemInsights and meeting hours insights are enabled; false if user's itemInsights and meeting hours insights are disabled. Default is true. Optional.
func (m *UserInsightsSettings) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserInsightsSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    return res
}
func (m *UserInsightsSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserInsightsSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsEnabled sets the isEnabled property value. true if user's itemInsights and meeting hours insights are enabled; false if user's itemInsights and meeting hours insights are disabled. Default is true. Optional.
func (m *UserInsightsSettings) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
