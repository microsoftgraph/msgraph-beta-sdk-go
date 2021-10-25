package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceHealthScriptRemediationSummary struct {
    additionalData map[string]interface{};
    remediatedDeviceCount *int32;
    scriptCount *int32;
}
func NewDeviceHealthScriptRemediationSummary()(*DeviceHealthScriptRemediationSummary) {
    m := &DeviceHealthScriptRemediationSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceHealthScriptRemediationSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceHealthScriptRemediationSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
func (m *DeviceHealthScriptRemediationSummary) GetScriptCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.scriptCount
    }
}
func (m *DeviceHealthScriptRemediationSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["remediatedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRemediatedDeviceCount(val)
        return nil
    }
    res["scriptCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetScriptCount(val)
        return nil
    }
    return res
}
func (m *DeviceHealthScriptRemediationSummary) IsNil()(bool) {
    return m == nil
}
func (m *DeviceHealthScriptRemediationSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("remediatedDeviceCount", m.GetRemediatedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("scriptCount", m.GetScriptCount())
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
func (m *DeviceHealthScriptRemediationSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceHealthScriptRemediationSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
func (m *DeviceHealthScriptRemediationSummary) SetScriptCount(value *int32)() {
    m.scriptCount = value
}
