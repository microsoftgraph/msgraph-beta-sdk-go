package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type StringKeyLongValuePair struct {
    additionalData map[string]interface{};
    key *string;
    value *int64;
}
func NewStringKeyLongValuePair()(*StringKeyLongValuePair) {
    m := &StringKeyLongValuePair{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *StringKeyLongValuePair) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *StringKeyLongValuePair) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
func (m *StringKeyLongValuePair) GetValue()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *StringKeyLongValuePair) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *StringKeyLongValuePair) IsNil()(bool) {
    return m == nil
}
func (m *StringKeyLongValuePair) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("value", m.GetValue())
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
func (m *StringKeyLongValuePair) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *StringKeyLongValuePair) SetKey(value *string)() {
    m.key = value
}
func (m *StringKeyLongValuePair) SetValue(value *int64)() {
    m.value = value
}
