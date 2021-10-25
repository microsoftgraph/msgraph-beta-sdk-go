package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AttributeMappingParameterSchema struct {
    additionalData map[string]interface{};
    allowMultipleOccurrences *bool;
    name *string;
    required *bool;
    type_escaped *AttributeType;
}
func NewAttributeMappingParameterSchema()(*AttributeMappingParameterSchema) {
    m := &AttributeMappingParameterSchema{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AttributeMappingParameterSchema) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AttributeMappingParameterSchema) GetAllowMultipleOccurrences()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleOccurrences
    }
}
func (m *AttributeMappingParameterSchema) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *AttributeMappingParameterSchema) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
func (m *AttributeMappingParameterSchema) GetType_escaped()(*AttributeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *AttributeMappingParameterSchema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowMultipleOccurrences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowMultipleOccurrences(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["required"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRequired(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeType)
        if err != nil {
            return err
        }
        cast := val.(AttributeType)
        m.SetType_escaped(&cast)
        return nil
    }
    return res
}
func (m *AttributeMappingParameterSchema) IsNil()(bool) {
    return m == nil
}
func (m *AttributeMappingParameterSchema) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleOccurrences", m.GetAllowMultipleOccurrences())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AttributeMappingParameterSchema) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AttributeMappingParameterSchema) SetAllowMultipleOccurrences(value *bool)() {
    m.allowMultipleOccurrences = value
}
func (m *AttributeMappingParameterSchema) SetName(value *string)() {
    m.name = value
}
func (m *AttributeMappingParameterSchema) SetRequired(value *bool)() {
    m.required = value
}
func (m *AttributeMappingParameterSchema) SetType_escaped(value *AttributeType)() {
    m.type_escaped = value
}
