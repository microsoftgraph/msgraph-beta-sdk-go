package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagementConditionExpressionString 
type ManagementConditionExpressionString struct {
    ManagementConditionExpression
    // The management condition statement expression string value.
    value *string;
}
// NewManagementConditionExpressionString instantiates a new managementConditionExpressionString and sets the default values.
func NewManagementConditionExpressionString()(*ManagementConditionExpressionString) {
    m := &ManagementConditionExpressionString{
        ManagementConditionExpression: *NewManagementConditionExpression(),
    }
    return m
}
// GetValue gets the value property value. The management condition statement expression string value.
func (m *ManagementConditionExpressionString) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementConditionExpressionString) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagementConditionExpression.GetFieldDeserializers()
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
func (m *ManagementConditionExpressionString) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementConditionExpressionString) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagementConditionExpression.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The management condition statement expression string value.
func (m *ManagementConditionExpressionString) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
