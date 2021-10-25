package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceComplianceScriptRunSummary struct {
    Entity
    detectionScriptErrorDeviceCount *int32;
    detectionScriptPendingDeviceCount *int32;
    issueDetectedDeviceCount *int32;
    lastScriptRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    noIssueDetectedDeviceCount *int32;
}
func NewDeviceComplianceScriptRunSummary()(*DeviceComplianceScriptRunSummary) {
    m := &DeviceComplianceScriptRunSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceComplianceScriptRunSummary) GetDetectionScriptErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptErrorDeviceCount
    }
}
func (m *DeviceComplianceScriptRunSummary) GetDetectionScriptPendingDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptPendingDeviceCount
    }
}
func (m *DeviceComplianceScriptRunSummary) GetIssueDetectedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.issueDetectedDeviceCount
    }
}
func (m *DeviceComplianceScriptRunSummary) GetLastScriptRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastScriptRunDateTime
    }
}
func (m *DeviceComplianceScriptRunSummary) GetNoIssueDetectedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.noIssueDetectedDeviceCount
    }
}
func (m *DeviceComplianceScriptRunSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["detectionScriptErrorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDetectionScriptErrorDeviceCount(val)
        return nil
    }
    res["detectionScriptPendingDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDetectionScriptPendingDeviceCount(val)
        return nil
    }
    res["issueDetectedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIssueDetectedDeviceCount(val)
        return nil
    }
    res["lastScriptRunDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastScriptRunDateTime(val)
        return nil
    }
    res["noIssueDetectedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNoIssueDetectedDeviceCount(val)
        return nil
    }
    return res
}
func (m *DeviceComplianceScriptRunSummary) IsNil()(bool) {
    return m == nil
}
func (m *DeviceComplianceScriptRunSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("detectionScriptErrorDeviceCount", m.GetDetectionScriptErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("detectionScriptPendingDeviceCount", m.GetDetectionScriptPendingDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("issueDetectedDeviceCount", m.GetIssueDetectedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastScriptRunDateTime", m.GetLastScriptRunDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("noIssueDetectedDeviceCount", m.GetNoIssueDetectedDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceComplianceScriptRunSummary) SetDetectionScriptErrorDeviceCount(value *int32)() {
    m.detectionScriptErrorDeviceCount = value
}
func (m *DeviceComplianceScriptRunSummary) SetDetectionScriptPendingDeviceCount(value *int32)() {
    m.detectionScriptPendingDeviceCount = value
}
func (m *DeviceComplianceScriptRunSummary) SetIssueDetectedDeviceCount(value *int32)() {
    m.issueDetectedDeviceCount = value
}
func (m *DeviceComplianceScriptRunSummary) SetLastScriptRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastScriptRunDateTime = value
}
func (m *DeviceComplianceScriptRunSummary) SetNoIssueDetectedDeviceCount(value *int32)() {
    m.noIssueDetectedDeviceCount = value
}
