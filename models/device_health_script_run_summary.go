package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptRunSummary contains properties for the run summary of a device management script.
type DeviceHealthScriptRunSummary struct {
    Entity
}
// NewDeviceHealthScriptRunSummary instantiates a new deviceHealthScriptRunSummary and sets the default values.
func NewDeviceHealthScriptRunSummary()(*DeviceHealthScriptRunSummary) {
    m := &DeviceHealthScriptRunSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceHealthScriptRunSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptRunSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptRunSummary(), nil
}
// GetDetectionScriptErrorDeviceCount gets the detectionScriptErrorDeviceCount property value. Number of devices on which the detection script execution encountered an error and did not complete
func (m *DeviceHealthScriptRunSummary) GetDetectionScriptErrorDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("detectionScriptErrorDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDetectionScriptNotApplicableDeviceCount gets the detectionScriptNotApplicableDeviceCount property value. Number of devices for which the detection script was not applicable
func (m *DeviceHealthScriptRunSummary) GetDetectionScriptNotApplicableDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("detectionScriptNotApplicableDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDetectionScriptPendingDeviceCount gets the detectionScriptPendingDeviceCount property value. Number of devices which have not yet run the latest version of the device health script
func (m *DeviceHealthScriptRunSummary) GetDetectionScriptPendingDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("detectionScriptPendingDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptRunSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["detectionScriptErrorDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptErrorDeviceCount(val)
        }
        return nil
    }
    res["detectionScriptNotApplicableDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptNotApplicableDeviceCount(val)
        }
        return nil
    }
    res["detectionScriptPendingDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptPendingDeviceCount(val)
        }
        return nil
    }
    res["issueDetectedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueDetectedDeviceCount(val)
        }
        return nil
    }
    res["issueRemediatedCumulativeDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueRemediatedCumulativeDeviceCount(val)
        }
        return nil
    }
    res["issueRemediatedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueRemediatedDeviceCount(val)
        }
        return nil
    }
    res["issueReoccurredDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueReoccurredDeviceCount(val)
        }
        return nil
    }
    res["lastScriptRunDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastScriptRunDateTime(val)
        }
        return nil
    }
    res["noIssueDetectedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoIssueDetectedDeviceCount(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["remediationScriptErrorDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationScriptErrorDeviceCount(val)
        }
        return nil
    }
    res["remediationSkippedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetIssueDetectedDeviceCount gets the issueDetectedDeviceCount property value. Number of devices for which the detection script found an issue
func (m *DeviceHealthScriptRunSummary) GetIssueDetectedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("issueDetectedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetIssueRemediatedCumulativeDeviceCount gets the issueRemediatedCumulativeDeviceCount property value. Number of devices that were remediated over the last 30 days
func (m *DeviceHealthScriptRunSummary) GetIssueRemediatedCumulativeDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("issueRemediatedCumulativeDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetIssueRemediatedDeviceCount gets the issueRemediatedDeviceCount property value. Number of devices for which the remediation script was able to resolve the detected issue
func (m *DeviceHealthScriptRunSummary) GetIssueRemediatedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("issueRemediatedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetIssueReoccurredDeviceCount gets the issueReoccurredDeviceCount property value. Number of devices for which the remediation script executed successfully but failed to resolve the detected issue
func (m *DeviceHealthScriptRunSummary) GetIssueReoccurredDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("issueReoccurredDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetLastScriptRunDateTime gets the lastScriptRunDateTime property value. Last run time for the script across all devices
func (m *DeviceHealthScriptRunSummary) GetLastScriptRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastScriptRunDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetNoIssueDetectedDeviceCount gets the noIssueDetectedDeviceCount property value. Number of devices for which the detection script did not find an issue and the device is healthy
func (m *DeviceHealthScriptRunSummary) GetNoIssueDetectedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("noIssueDetectedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptRunSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediationScriptErrorDeviceCount gets the remediationScriptErrorDeviceCount property value. Number of devices for which the remediation script execution encountered an error and did not complete
func (m *DeviceHealthScriptRunSummary) GetRemediationScriptErrorDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("remediationScriptErrorDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRemediationSkippedDeviceCount gets the remediationSkippedDeviceCount property value. Number of devices for which remediation was skipped
func (m *DeviceHealthScriptRunSummary) GetRemediationSkippedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("remediationSkippedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptRunSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetDetectionScriptErrorDeviceCount sets the detectionScriptErrorDeviceCount property value. Number of devices on which the detection script execution encountered an error and did not complete
func (m *DeviceHealthScriptRunSummary) SetDetectionScriptErrorDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("detectionScriptErrorDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionScriptNotApplicableDeviceCount sets the detectionScriptNotApplicableDeviceCount property value. Number of devices for which the detection script was not applicable
func (m *DeviceHealthScriptRunSummary) SetDetectionScriptNotApplicableDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("detectionScriptNotApplicableDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionScriptPendingDeviceCount sets the detectionScriptPendingDeviceCount property value. Number of devices which have not yet run the latest version of the device health script
func (m *DeviceHealthScriptRunSummary) SetDetectionScriptPendingDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("detectionScriptPendingDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIssueDetectedDeviceCount sets the issueDetectedDeviceCount property value. Number of devices for which the detection script found an issue
func (m *DeviceHealthScriptRunSummary) SetIssueDetectedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("issueDetectedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIssueRemediatedCumulativeDeviceCount sets the issueRemediatedCumulativeDeviceCount property value. Number of devices that were remediated over the last 30 days
func (m *DeviceHealthScriptRunSummary) SetIssueRemediatedCumulativeDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("issueRemediatedCumulativeDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIssueRemediatedDeviceCount sets the issueRemediatedDeviceCount property value. Number of devices for which the remediation script was able to resolve the detected issue
func (m *DeviceHealthScriptRunSummary) SetIssueRemediatedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("issueRemediatedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIssueReoccurredDeviceCount sets the issueReoccurredDeviceCount property value. Number of devices for which the remediation script executed successfully but failed to resolve the detected issue
func (m *DeviceHealthScriptRunSummary) SetIssueReoccurredDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("issueReoccurredDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetLastScriptRunDateTime sets the lastScriptRunDateTime property value. Last run time for the script across all devices
func (m *DeviceHealthScriptRunSummary) SetLastScriptRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastScriptRunDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetNoIssueDetectedDeviceCount sets the noIssueDetectedDeviceCount property value. Number of devices for which the detection script did not find an issue and the device is healthy
func (m *DeviceHealthScriptRunSummary) SetNoIssueDetectedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("noIssueDetectedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptRunSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationScriptErrorDeviceCount sets the remediationScriptErrorDeviceCount property value. Number of devices for which the remediation script execution encountered an error and did not complete
func (m *DeviceHealthScriptRunSummary) SetRemediationScriptErrorDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("remediationScriptErrorDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationSkippedDeviceCount sets the remediationSkippedDeviceCount property value. Number of devices for which remediation was skipped
func (m *DeviceHealthScriptRunSummary) SetRemediationSkippedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("remediationSkippedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// DeviceHealthScriptRunSummaryable 
type DeviceHealthScriptRunSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetectionScriptErrorDeviceCount()(*int32)
    GetDetectionScriptNotApplicableDeviceCount()(*int32)
    GetDetectionScriptPendingDeviceCount()(*int32)
    GetIssueDetectedDeviceCount()(*int32)
    GetIssueRemediatedCumulativeDeviceCount()(*int32)
    GetIssueRemediatedDeviceCount()(*int32)
    GetIssueReoccurredDeviceCount()(*int32)
    GetLastScriptRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNoIssueDetectedDeviceCount()(*int32)
    GetOdataType()(*string)
    GetRemediationScriptErrorDeviceCount()(*int32)
    GetRemediationSkippedDeviceCount()(*int32)
    SetDetectionScriptErrorDeviceCount(value *int32)()
    SetDetectionScriptNotApplicableDeviceCount(value *int32)()
    SetDetectionScriptPendingDeviceCount(value *int32)()
    SetIssueDetectedDeviceCount(value *int32)()
    SetIssueRemediatedCumulativeDeviceCount(value *int32)()
    SetIssueRemediatedDeviceCount(value *int32)()
    SetIssueReoccurredDeviceCount(value *int32)()
    SetLastScriptRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNoIssueDetectedDeviceCount(value *int32)()
    SetOdataType(value *string)()
    SetRemediationScriptErrorDeviceCount(value *int32)()
    SetRemediationSkippedDeviceCount(value *int32)()
}
