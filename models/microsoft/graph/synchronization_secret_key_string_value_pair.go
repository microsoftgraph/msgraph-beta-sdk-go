package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationSecretKeyStringValuePair struct {
    additionalData map[string]interface{};
    key *SynchronizationSecret;
    value *string;
}
func NewSynchronizationSecretKeyStringValuePair()(*SynchronizationSecretKeyStringValuePair) {
    m := &SynchronizationSecretKeyStringValuePair{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SynchronizationSecretKeyStringValuePair) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SynchronizationSecretKeyStringValuePair) GetKey()(*SynchronizationSecret) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
func (m *SynchronizationSecretKeyStringValuePair) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *SynchronizationSecretKeyStringValuePair) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationSecret)
        if err != nil {
            return err
        }
        cast := val.(SynchronizationSecret)
        m.SetKey(&cast)
        return nil
    }
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
func (m *SynchronizationSecretKeyStringValuePair) IsNil()(bool) {
    return m == nil
}
func (m *SynchronizationSecretKeyStringValuePair) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetKey() != nil {
        cast := m.GetKey().String()
        err := writer.WriteStringValue("key", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *SynchronizationSecretKeyStringValuePair) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SynchronizationSecretKeyStringValuePair) SetKey(value *SynchronizationSecret)() {
    m.key = value
}
func (m *SynchronizationSecretKeyStringValuePair) SetValue(value *string)() {
    m.value = value
}
