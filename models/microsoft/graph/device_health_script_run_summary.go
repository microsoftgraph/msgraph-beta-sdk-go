package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceHealthScriptRunSummary struct {
    Entity
    detectionScriptErrorDeviceCount *int32;
    detectionScriptNotApplicableDeviceCount *int32;
    detectionScriptPendingDeviceCount *int32;
    issueDetectedDeviceCount *int32;
    issueRemediatedCumulativeDeviceCount *int32;
    issueRemediatedDeviceCount *int32;
    issueReoccurredDeviceCount *int32;
    lastScriptRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    noIssueDetectedDeviceCount *int32;
    remediationScriptErrorDeviceCount *int32;
    remediationSkippedDeviceCount *int32;
}
func NewDeviceHealthScriptRunSummary()(*DeviceHealthScriptRunSummary) {
    m := &DeviceHealthScriptRunSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceHealthScriptRunSummary) GetDetectionScriptErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptErrorDeviceCount
    }
}
func (m *DeviceHealthScriptRunSummary) GetDetectionScriptNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptNotApplicableDeviceCount
    }
}
func (m *DeviceHealthScriptRunSummary) GetDetectionScriptPendingDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptPendingDeviceCount
    }
}
func (m *DeviceHealthScriptRunSummary) GetIssueDetectedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.issueDetectedDeviceCount
    }
}
func (m *DeviceHealthScriptRunSummary) GetIssueRemediatedCumulativeDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.issueRemediatedCumulativeDeviceCount
    }
}
func (m *DeviceHealthScriptRunSummary) GetIssueRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.issueRemediatedDeviceCount
    }
}
func (m *DeviceHealthScriptRunSummary) GetIssueReoccurredDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.issueReoccurredDeviceCount
    }
}
func (m *DeviceHealthScriptRunSummary) GetLastScriptRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastScriptRunDateTime
    }
}
func (m *DeviceHealthScriptRunSummary) GetNoIssueDetectedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.noIssueDetectedDeviceCount
    }
}
func (m *DeviceHealthScriptRunSummary) GetRemediationScriptErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediationScriptErrorDeviceCount
    }
}
func (m *DeviceHealthScriptRunSummary) GetRemediationSkippedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediationSkippedDeviceCount
    }
}
func (m *DeviceHealthScriptRunSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["detectionScriptErrorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDetectionScriptErrorDeviceCount(val)
        return nil
    }
    res["detectionScriptNotApplicableDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDetectionScriptNotApplicableDeviceCount(val)
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
    res["issueRemediatedCumulativeDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIssueRemediatedCumulativeDeviceCount(val)
        return nil
    }
    res["issueRemediatedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIssueRemediatedDeviceCount(val)
        return nil
    }
    res["issueReoccurredDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIssueReoccurredDeviceCount(val)
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
    res["remediationScriptErrorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRemediationScriptErrorDeviceCount(val)
        return nil
    }
    res["remediationSkippedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRemediationSkippedDeviceCount(val)
        return nil
    }
    return res
}
func (m *DeviceHealthScriptRunSummary) IsNil()(bool) {
    return m == nil
}
func (m *DeviceHealthScriptRunSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("detectionScriptNotApplicableDeviceCount", m.GetDetectionScriptNotApplicableDeviceCount())
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
        err = writer.WriteInt32Value("issueRemediatedCumulativeDeviceCount", m.GetIssueRemediatedCumulativeDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("issueRemediatedDeviceCount", m.GetIssueRemediatedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("issueReoccurredDeviceCount", m.GetIssueReoccurredDeviceCount())
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
    {
        err = writer.WriteInt32Value("remediationScriptErrorDeviceCount", m.GetRemediationScriptErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("remediationSkippedDeviceCount", m.GetRemediationSkippedDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceHealthScriptRunSummary) SetDetectionScriptErrorDeviceCount(value *int32)() {
    m.detectionScriptErrorDeviceCount = value
}
func (m *DeviceHealthScriptRunSummary) SetDetectionScriptNotApplicableDeviceCount(value *int32)() {
    m.detectionScriptNotApplicableDeviceCount = value
}
func (m *DeviceHealthScriptRunSummary) SetDetectionScriptPendingDeviceCount(value *int32)() {
    m.detectionScriptPendingDeviceCount = value
}
func (m *DeviceHealthScriptRunSummary) SetIssueDetectedDeviceCount(value *int32)() {
    m.issueDetectedDeviceCount = value
}
func (m *DeviceHealthScriptRunSummary) SetIssueRemediatedCumulativeDeviceCount(value *int32)() {
    m.issueRemediatedCumulativeDeviceCount = value
}
func (m *DeviceHealthScriptRunSummary) SetIssueRemediatedDeviceCount(value *int32)() {
    m.issueRemediatedDeviceCount = value
}
func (m *DeviceHealthScriptRunSummary) SetIssueReoccurredDeviceCount(value *int32)() {
    m.issueReoccurredDeviceCount = value
}
func (m *DeviceHealthScriptRunSummary) SetLastScriptRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastScriptRunDateTime = value
}
func (m *DeviceHealthScriptRunSummary) SetNoIssueDetectedDeviceCount(value *int32)() {
    m.noIssueDetectedDeviceCount = value
}
func (m *DeviceHealthScriptRunSummary) SetRemediationScriptErrorDeviceCount(value *int32)() {
    m.remediationScriptErrorDeviceCount = value
}
func (m *DeviceHealthScriptRunSummary) SetRemediationSkippedDeviceCount(value *int32)() {
    m.remediationSkippedDeviceCount = value
}
