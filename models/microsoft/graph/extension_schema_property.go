package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ExtensionSchemaProperty struct {
    additionalData map[string]interface{};
    name *string;
    type_escaped *string;
}
func NewExtensionSchemaProperty()(*ExtensionSchemaProperty) {
    m := &ExtensionSchemaProperty{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ExtensionSchemaProperty) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ExtensionSchemaProperty) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *ExtensionSchemaProperty) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *ExtensionSchemaProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    return res
}
func (m *ExtensionSchemaProperty) IsNil()(bool) {
    return m == nil
}
func (m *ExtensionSchemaProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
func (m *ExtensionSchemaProperty) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ExtensionSchemaProperty) SetName(value *string)() {
    m.name = value
}
func (m *ExtensionSchemaProperty) SetType_escaped(value *string)() {
    m.type_escaped = value
}
