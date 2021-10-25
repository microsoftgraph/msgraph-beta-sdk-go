package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceHealthScriptParameter struct {
    additionalData map[string]interface{};
    applyDefaultValueWhenNotAssigned *bool;
    description *string;
    isRequired *bool;
    name *string;
}
func NewDeviceHealthScriptParameter()(*DeviceHealthScriptParameter) {
    m := &DeviceHealthScriptParameter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceHealthScriptParameter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceHealthScriptParameter) GetApplyDefaultValueWhenNotAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applyDefaultValueWhenNotAssigned
    }
}
func (m *DeviceHealthScriptParameter) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *DeviceHealthScriptParameter) GetIsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequired
    }
}
func (m *DeviceHealthScriptParameter) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *DeviceHealthScriptParameter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applyDefaultValueWhenNotAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetApplyDefaultValueWhenNotAssigned(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["isRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRequired(val)
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
    return res
}
func (m *DeviceHealthScriptParameter) IsNil()(bool) {
    return m == nil
}
func (m *DeviceHealthScriptParameter) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("applyDefaultValueWhenNotAssigned", m.GetApplyDefaultValueWhenNotAssigned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceHealthScriptParameter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceHealthScriptParameter) SetApplyDefaultValueWhenNotAssigned(value *bool)() {
    m.applyDefaultValueWhenNotAssigned = value
}
func (m *DeviceHealthScriptParameter) SetDescription(value *string)() {
    m.description = value
}
func (m *DeviceHealthScriptParameter) SetIsRequired(value *bool)() {
    m.isRequired = value
}
func (m *DeviceHealthScriptParameter) SetName(value *string)() {
    m.name = value
}
