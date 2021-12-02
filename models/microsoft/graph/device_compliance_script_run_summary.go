package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceComplianceScriptRunSummary 
type DeviceComplianceScriptRunSummary struct {
    Entity
    // Number of devices on which the detection script execution encountered an error and did not complete. Valid values -2147483648 to 2147483647
    detectionScriptErrorDeviceCount *int32;
    // Number of devices which have not yet run the latest version of the device compliance script. Valid values -2147483648 to 2147483647
    detectionScriptPendingDeviceCount *int32;
    // Number of devices for which the detection script found an issue. Valid values -2147483648 to 2147483647
    issueDetectedDeviceCount *int32;
    // Last run time for the script across all devices
    lastScriptRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Number of devices for which the detection script did not find an issue and the device is healthy. Valid values -2147483648 to 2147483647
    noIssueDetectedDeviceCount *int32;
}
// NewDeviceComplianceScriptRunSummary instantiates a new deviceComplianceScriptRunSummary and sets the default values.
func NewDeviceComplianceScriptRunSummary()(*DeviceComplianceScriptRunSummary) {
    m := &DeviceComplianceScriptRunSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetDetectionScriptErrorDeviceCount gets the detectionScriptErrorDeviceCount property value. Number of devices on which the detection script execution encountered an error and did not complete. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) GetDetectionScriptErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptErrorDeviceCount
    }
}
// GetDetectionScriptPendingDeviceCount gets the detectionScriptPendingDeviceCount property value. Number of devices which have not yet run the latest version of the device compliance script. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) GetDetectionScriptPendingDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptPendingDeviceCount
    }
}
// GetIssueDetectedDeviceCount gets the issueDetectedDeviceCount property value. Number of devices for which the detection script found an issue. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) GetIssueDetectedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.issueDetectedDeviceCount
    }
}
// GetLastScriptRunDateTime gets the lastScriptRunDateTime property value. Last run time for the script across all devices
func (m *DeviceComplianceScriptRunSummary) GetLastScriptRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastScriptRunDateTime
    }
}
// GetNoIssueDetectedDeviceCount gets the noIssueDetectedDeviceCount property value. Number of devices for which the detection script did not find an issue and the device is healthy. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) GetNoIssueDetectedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.noIssueDetectedDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceScriptRunSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    return res
}
func (m *DeviceComplianceScriptRunSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetDetectionScriptErrorDeviceCount sets the detectionScriptErrorDeviceCount property value. Number of devices on which the detection script execution encountered an error and did not complete. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) SetDetectionScriptErrorDeviceCount(value *int32)() {
    if m != nil {
        m.detectionScriptErrorDeviceCount = value
    }
}
// SetDetectionScriptPendingDeviceCount sets the detectionScriptPendingDeviceCount property value. Number of devices which have not yet run the latest version of the device compliance script. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) SetDetectionScriptPendingDeviceCount(value *int32)() {
    if m != nil {
        m.detectionScriptPendingDeviceCount = value
    }
}
// SetIssueDetectedDeviceCount sets the issueDetectedDeviceCount property value. Number of devices for which the detection script found an issue. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) SetIssueDetectedDeviceCount(value *int32)() {
    if m != nil {
        m.issueDetectedDeviceCount = value
    }
}
// SetLastScriptRunDateTime sets the lastScriptRunDateTime property value. Last run time for the script across all devices
func (m *DeviceComplianceScriptRunSummary) SetLastScriptRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastScriptRunDateTime = value
    }
}
// SetNoIssueDetectedDeviceCount sets the noIssueDetectedDeviceCount property value. Number of devices for which the detection script did not find an issue and the device is healthy. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) SetNoIssueDetectedDeviceCount(value *int32)() {
    if m != nil {
        m.noIssueDetectedDeviceCount = value
    }
}
