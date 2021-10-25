package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceCompliancePolicyScript struct {
    additionalData map[string]interface{};
    deviceComplianceScriptId *string;
    rulesContent []byte;
}
func NewDeviceCompliancePolicyScript()(*DeviceCompliancePolicyScript) {
    m := &DeviceCompliancePolicyScript{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceCompliancePolicyScript) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceCompliancePolicyScript) GetDeviceComplianceScriptId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScriptId
    }
}
func (m *DeviceCompliancePolicyScript) GetRulesContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.rulesContent
    }
}
func (m *DeviceCompliancePolicyScript) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceComplianceScriptId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceComplianceScriptId(val)
        return nil
    }
    res["rulesContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetRulesContent(val)
        return nil
    }
    return res
}
func (m *DeviceCompliancePolicyScript) IsNil()(bool) {
    return m == nil
}
func (m *DeviceCompliancePolicyScript) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceComplianceScriptId", m.GetDeviceComplianceScriptId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("rulesContent", m.GetRulesContent())
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
func (m *DeviceCompliancePolicyScript) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceCompliancePolicyScript) SetDeviceComplianceScriptId(value *string)() {
    m.deviceComplianceScriptId = value
}
func (m *DeviceCompliancePolicyScript) SetRulesContent(value []byte)() {
    m.rulesContent = value
}
