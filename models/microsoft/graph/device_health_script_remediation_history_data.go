package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceHealthScriptRemediationHistoryData struct {
    additionalData map[string]interface{};
    date *string;
    noIssueDeviceCount *int32;
    remediatedDeviceCount *int32;
}
func NewDeviceHealthScriptRemediationHistoryData()(*DeviceHealthScriptRemediationHistoryData) {
    m := &DeviceHealthScriptRemediationHistoryData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceHealthScriptRemediationHistoryData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceHealthScriptRemediationHistoryData) GetDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.date
    }
}
func (m *DeviceHealthScriptRemediationHistoryData) GetNoIssueDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.noIssueDeviceCount
    }
}
func (m *DeviceHealthScriptRemediationHistoryData) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
func (m *DeviceHealthScriptRemediationHistoryData) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["date"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDate(val)
        return nil
    }
    res["noIssueDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNoIssueDeviceCount(val)
        return nil
    }
    res["remediatedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRemediatedDeviceCount(val)
        return nil
    }
    return res
}
func (m *DeviceHealthScriptRemediationHistoryData) IsNil()(bool) {
    return m == nil
}
func (m *DeviceHealthScriptRemediationHistoryData) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("date", m.GetDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("noIssueDeviceCount", m.GetNoIssueDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("remediatedDeviceCount", m.GetRemediatedDeviceCount())
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
func (m *DeviceHealthScriptRemediationHistoryData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceHealthScriptRemediationHistoryData) SetDate(value *string)() {
    m.date = value
}
func (m *DeviceHealthScriptRemediationHistoryData) SetNoIssueDeviceCount(value *int32)() {
    m.noIssueDeviceCount = value
}
func (m *DeviceHealthScriptRemediationHistoryData) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
