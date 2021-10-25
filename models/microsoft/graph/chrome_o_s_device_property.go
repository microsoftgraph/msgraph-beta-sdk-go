package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ChromeOSDeviceProperty struct {
    additionalData map[string]interface{};
    name *string;
    updatable *bool;
    value *string;
    valueType *string;
}
func NewChromeOSDeviceProperty()(*ChromeOSDeviceProperty) {
    m := &ChromeOSDeviceProperty{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ChromeOSDeviceProperty) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ChromeOSDeviceProperty) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *ChromeOSDeviceProperty) GetUpdatable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.updatable
    }
}
func (m *ChromeOSDeviceProperty) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *ChromeOSDeviceProperty) GetValueType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
func (m *ChromeOSDeviceProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["updatable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUpdatable(val)
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
    res["valueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValueType(val)
        return nil
    }
    return res
}
func (m *ChromeOSDeviceProperty) IsNil()(bool) {
    return m == nil
}
func (m *ChromeOSDeviceProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("updatable", m.GetUpdatable())
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
        err := writer.WriteStringValue("valueType", m.GetValueType())
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
func (m *ChromeOSDeviceProperty) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ChromeOSDeviceProperty) SetName(value *string)() {
    m.name = value
}
func (m *ChromeOSDeviceProperty) SetUpdatable(value *bool)() {
    m.updatable = value
}
func (m *ChromeOSDeviceProperty) SetValue(value *string)() {
    m.value = value
}
func (m *ChromeOSDeviceProperty) SetValueType(value *string)() {
    m.valueType = value
}
