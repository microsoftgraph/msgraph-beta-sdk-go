package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type StringKeyAttributeMappingSourceValuePair struct {
    additionalData map[string]interface{};
    key *string;
    value *AttributeMappingSource;
}
func NewStringKeyAttributeMappingSourceValuePair()(*StringKeyAttributeMappingSourceValuePair) {
    m := &StringKeyAttributeMappingSourceValuePair{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *StringKeyAttributeMappingSourceValuePair) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *StringKeyAttributeMappingSourceValuePair) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
func (m *StringKeyAttributeMappingSourceValuePair) GetValue()(*AttributeMappingSource) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *StringKeyAttributeMappingSourceValuePair) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttributeMappingSource() })
        if err != nil {
            return err
        }
        m.SetValue(val.(*AttributeMappingSource))
        return nil
    }
    return res
}
func (m *StringKeyAttributeMappingSourceValuePair) IsNil()(bool) {
    return m == nil
}
func (m *StringKeyAttributeMappingSourceValuePair) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("value", m.GetValue())
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
func (m *StringKeyAttributeMappingSourceValuePair) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *StringKeyAttributeMappingSourceValuePair) SetKey(value *string)() {
    m.key = value
}
func (m *StringKeyAttributeMappingSourceValuePair) SetValue(value *AttributeMappingSource)() {
    m.value = value
}
