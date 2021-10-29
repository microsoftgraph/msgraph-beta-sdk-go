package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagementConditionExpressionString struct {
    ManagementConditionExpression
    // The management condition statement expression string value.
    value *string;
}
// Instantiates a new managementConditionExpressionString and sets the default values.
func NewManagementConditionExpressionString()(*ManagementConditionExpressionString) {
    m := &ManagementConditionExpressionString{
        ManagementConditionExpression: *NewManagementConditionExpression(),
    }
    return m
}
// Gets the value property value. The management condition statement expression string value.
func (m *ManagementConditionExpressionString) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// The deserialization information for the current model
func (m *ManagementConditionExpressionString) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagementConditionExpression.GetFieldDeserializers()
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *ManagementConditionExpressionString) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the value property value. The management condition statement expression string value.
// Parameters:
//  - value : Value to set for the value property.
func (m *ManagementConditionExpressionString) SetValue(value *string)() {
    m.value = value
}
