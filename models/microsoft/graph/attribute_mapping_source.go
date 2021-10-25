package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AttributeMappingSource struct {
    additionalData map[string]interface{};
    expression *string;
    name *string;
    parameters []StringKeyAttributeMappingSourceValuePair;
    type_escaped *AttributeMappingSourceType;
}
func NewAttributeMappingSource()(*AttributeMappingSource) {
    m := &AttributeMappingSource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AttributeMappingSource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AttributeMappingSource) GetExpression()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expression
    }
}
func (m *AttributeMappingSource) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *AttributeMappingSource) GetParameters()([]StringKeyAttributeMappingSourceValuePair) {
    if m == nil {
        return nil
    } else {
        return m.parameters
    }
}
func (m *AttributeMappingSource) GetType_escaped()(*AttributeMappingSourceType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *AttributeMappingSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["expression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExpression(val)
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
    res["parameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewStringKeyAttributeMappingSourceValuePair() })
        if err != nil {
            return err
        }
        res := make([]StringKeyAttributeMappingSourceValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*StringKeyAttributeMappingSourceValuePair))
        }
        m.SetParameters(res)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeMappingSourceType)
        if err != nil {
            return err
        }
        cast := val.(AttributeMappingSourceType)
        m.SetType_escaped(&cast)
        return nil
    }
    return res
}
func (m *AttributeMappingSource) IsNil()(bool) {
    return m == nil
}
func (m *AttributeMappingSource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expression", m.GetExpression())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("parameters", cast)
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
func (m *AttributeMappingSource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AttributeMappingSource) SetExpression(value *string)() {
    m.expression = value
}
func (m *AttributeMappingSource) SetName(value *string)() {
    m.name = value
}
func (m *AttributeMappingSource) SetParameters(value []StringKeyAttributeMappingSourceValuePair)() {
    m.parameters = value
}
func (m *AttributeMappingSource) SetType_escaped(value *AttributeMappingSourceType)() {
    m.type_escaped = value
}
