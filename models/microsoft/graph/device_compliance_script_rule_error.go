package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceComplianceScriptRuleError struct {
    DeviceComplianceScriptError
    settingName *string;
}
func NewDeviceComplianceScriptRuleError()(*DeviceComplianceScriptRuleError) {
    m := &DeviceComplianceScriptRuleError{
        DeviceComplianceScriptError: *NewDeviceComplianceScriptError(),
    }
    return m
}
func (m *DeviceComplianceScriptRuleError) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
func (m *DeviceComplianceScriptRuleError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DeviceComplianceScriptError.GetFieldDeserializers()
    res["settingName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingName(val)
        return nil
    }
    return res
}
func (m *DeviceComplianceScriptRuleError) IsNil()(bool) {
    return m == nil
}
func (m *DeviceComplianceScriptRuleError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DeviceComplianceScriptError.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceComplianceScriptRuleError) SetSettingName(value *string)() {
    m.settingName = value
}
