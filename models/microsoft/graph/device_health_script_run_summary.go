package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceHealthScriptRunSummary struct {
    Entity
    // Number of devices on which the detection script execution encountered an error and did not complete
    detectionScriptErrorDeviceCount *int32;
    // Number of devices for which the detection script was not applicable
    detectionScriptNotApplicableDeviceCount *int32;
    // Number of devices which have not yet run the latest version of the device health script
    detectionScriptPendingDeviceCount *int32;
    // Number of devices for which the detection script found an issue
    issueDetectedDeviceCount *int32;
    // Number of devices that were remediated over the last 30 days
    issueRemediatedCumulativeDeviceCount *int32;
    // Number of devices for which the remediation script was able to resolve the detected issue
    issueRemediatedDeviceCount *int32;
    // Number of devices for which the remediation script executed successfully but failed to resolve the detected issue
    issueReoccurredDeviceCount *int32;
    // Last run time for the script across all devices
    lastScriptRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Number of devices for which the detection script did not find an issue and the device is healthy
    noIssueDetectedDeviceCount *int32;
    // Number of devices for which the remediation script execution encountered an error and did not complete
    remediationScriptErrorDeviceCount *int32;
    // Number of devices for which remediation was skipped
    remediationSkippedDeviceCount *int32;
}
// Instantiates a new deviceHealthScriptRunSummary and sets the default values.
func NewDeviceHealthScriptRunSummary()(*DeviceHealthScriptRunSummary) {
    m := &DeviceHealthScriptRunSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the detectionScriptErrorDeviceCount property value. Number of devices on which the detection script execution encountered an error and did not complete
func (m *DeviceHealthScriptRunSummary) GetDetectionScriptErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptErrorDeviceCount
    }
}
// Gets the detectionScriptNotApplicableDeviceCount property value. Number of devices for which the detection script was not applicable
func (m *DeviceHealthScriptRunSummary) GetDetectionScriptNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptNotApplicableDeviceCount
    }
}
// Gets the detectionScriptPendingDeviceCount property value. Number of devices which have not yet run the latest version of the device health script
func (m *DeviceHealthScriptRunSummary) GetDetectionScriptPendingDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptPendingDeviceCount
    }
}
// Gets the issueDetectedDeviceCount property value. Number of devices for which the detection script found an issue
func (m *DeviceHealthScriptRunSummary) GetIssueDetectedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.issueDetectedDeviceCount
    }
}
// Gets the issueRemediatedCumulativeDeviceCount property value. Number of devices that were remediated over the last 30 days
func (m *DeviceHealthScriptRunSummary) GetIssueRemediatedCumulativeDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.issueRemediatedCumulativeDeviceCount
    }
}
// Gets the issueRemediatedDeviceCount property value. Number of devices for which the remediation script was able to resolve the detected issue
func (m *DeviceHealthScriptRunSummary) GetIssueRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.issueRemediatedDeviceCount
    }
}
// Gets the issueReoccurredDeviceCount property value. Number of devices for which the remediation script executed successfully but failed to resolve the detected issue
func (m *DeviceHealthScriptRunSummary) GetIssueReoccurredDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.issueReoccurredDeviceCount
    }
}
// Gets the lastScriptRunDateTime property value. Last run time for the script across all devices
func (m *DeviceHealthScriptRunSummary) GetLastScriptRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastScriptRunDateTime
    }
}
// Gets the noIssueDetectedDeviceCount property value. Number of devices for which the detection script did not find an issue and the device is healthy
func (m *DeviceHealthScriptRunSummary) GetNoIssueDetectedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.noIssueDetectedDeviceCount
    }
}
// Gets the remediationScriptErrorDeviceCount property value. Number of devices for which the remediation script execution encountered an error and did not complete
func (m *DeviceHealthScriptRunSummary) GetRemediationScriptErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediationScriptErrorDeviceCount
    }
}
// Gets the remediationSkippedDeviceCount property value. Number of devices for which remediation was skipped
func (m *DeviceHealthScriptRunSummary) GetRemediationSkippedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediationSkippedDeviceCount
    }
}
// The deserialization information for the current model
func (m *DeviceHealthScriptRunSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["detectionScriptErrorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptErrorDeviceCount(val)
        }
        return nil
    }
    res["detectionScriptNotApplicableDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptNotApplicableDeviceCount(val)
        }
        return nil
    }
    res["detectionScriptPendingDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptPendingDeviceCount(val)
        }
        return nil
    }
    res["issueDetectedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueDetectedDeviceCount(val)
        }
        return nil
    }
    res["issueRemediatedCumulativeDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueRemediatedCumulativeDeviceCount(val)
        }
        return nil
    }
    res["issueRemediatedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueRemediatedDeviceCount(val)
        }
        return nil
    }
    res["issueReoccurredDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueReoccurredDeviceCount(val)
        }
        return nil
    }
    res["lastScriptRunDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastScriptRunDateTime(val)
        }
        return nil
    }
    res["noIssueDetectedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoIssueDetectedDeviceCount(val)
        }
        return nil
    }
    res["remediationScriptErrorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationScriptErrorDeviceCount(val)
        }
        return nil
    }
    res["remediationSkippedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationSkippedDeviceCount(val)
        }
        return nil
    }
    return res
}
func (m *DeviceHealthScriptRunSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the detectionScriptErrorDeviceCount property value. Number of devices on which the detection script execution encountered an error and did not complete
// Parameters:
//  - value : Value to set for the detectionScriptErrorDeviceCount property.
func (m *DeviceHealthScriptRunSummary) SetDetectionScriptErrorDeviceCount(value *int32)() {
    m.detectionScriptErrorDeviceCount = value
}
// Sets the detectionScriptNotApplicableDeviceCount property value. Number of devices for which the detection script was not applicable
// Parameters:
//  - value : Value to set for the detectionScriptNotApplicableDeviceCount property.
func (m *DeviceHealthScriptRunSummary) SetDetectionScriptNotApplicableDeviceCount(value *int32)() {
    m.detectionScriptNotApplicableDeviceCount = value
}
// Sets the detectionScriptPendingDeviceCount property value. Number of devices which have not yet run the latest version of the device health script
// Parameters:
//  - value : Value to set for the detectionScriptPendingDeviceCount property.
func (m *DeviceHealthScriptRunSummary) SetDetectionScriptPendingDeviceCount(value *int32)() {
    m.detectionScriptPendingDeviceCount = value
}
// Sets the issueDetectedDeviceCount property value. Number of devices for which the detection script found an issue
// Parameters:
//  - value : Value to set for the issueDetectedDeviceCount property.
func (m *DeviceHealthScriptRunSummary) SetIssueDetectedDeviceCount(value *int32)() {
    m.issueDetectedDeviceCount = value
}
// Sets the issueRemediatedCumulativeDeviceCount property value. Number of devices that were remediated over the last 30 days
// Parameters:
//  - value : Value to set for the issueRemediatedCumulativeDeviceCount property.
func (m *DeviceHealthScriptRunSummary) SetIssueRemediatedCumulativeDeviceCount(value *int32)() {
    m.issueRemediatedCumulativeDeviceCount = value
}
// Sets the issueRemediatedDeviceCount property value. Number of devices for which the remediation script was able to resolve the detected issue
// Parameters:
//  - value : Value to set for the issueRemediatedDeviceCount property.
func (m *DeviceHealthScriptRunSummary) SetIssueRemediatedDeviceCount(value *int32)() {
    m.issueRemediatedDeviceCount = value
}
// Sets the issueReoccurredDeviceCount property value. Number of devices for which the remediation script executed successfully but failed to resolve the detected issue
// Parameters:
//  - value : Value to set for the issueReoccurredDeviceCount property.
func (m *DeviceHealthScriptRunSummary) SetIssueReoccurredDeviceCount(value *int32)() {
    m.issueReoccurredDeviceCount = value
}
// Sets the lastScriptRunDateTime property value. Last run time for the script across all devices
// Parameters:
//  - value : Value to set for the lastScriptRunDateTime property.
func (m *DeviceHealthScriptRunSummary) SetLastScriptRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastScriptRunDateTime = value
}
// Sets the noIssueDetectedDeviceCount property value. Number of devices for which the detection script did not find an issue and the device is healthy
// Parameters:
//  - value : Value to set for the noIssueDetectedDeviceCount property.
func (m *DeviceHealthScriptRunSummary) SetNoIssueDetectedDeviceCount(value *int32)() {
    m.noIssueDetectedDeviceCount = value
}
// Sets the remediationScriptErrorDeviceCount property value. Number of devices for which the remediation script execution encountered an error and did not complete
// Parameters:
//  - value : Value to set for the remediationScriptErrorDeviceCount property.
func (m *DeviceHealthScriptRunSummary) SetRemediationScriptErrorDeviceCount(value *int32)() {
    m.remediationScriptErrorDeviceCount = value
}
// Sets the remediationSkippedDeviceCount property value. Number of devices for which remediation was skipped
// Parameters:
//  - value : Value to set for the remediationSkippedDeviceCount property.
func (m *DeviceHealthScriptRunSummary) SetRemediationSkippedDeviceCount(value *int32)() {
    m.remediationSkippedDeviceCount = value
}
