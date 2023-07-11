package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceScriptRunSummary contains properties for the run summary of a device management script.
type DeviceComplianceScriptRunSummary struct {
    Entity
}
// NewDeviceComplianceScriptRunSummary instantiates a new deviceComplianceScriptRunSummary and sets the default values.
func NewDeviceComplianceScriptRunSummary()(*DeviceComplianceScriptRunSummary) {
    m := &DeviceComplianceScriptRunSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceComplianceScriptRunSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceScriptRunSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceComplianceScriptRunSummary(), nil
}
// GetDetectionScriptErrorDeviceCount gets the detectionScriptErrorDeviceCount property value. Number of devices on which the detection script execution encountered an error and did not complete. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) GetDetectionScriptErrorDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("detectionScriptErrorDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDetectionScriptPendingDeviceCount gets the detectionScriptPendingDeviceCount property value. Number of devices which have not yet run the latest version of the device compliance script. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) GetDetectionScriptPendingDeviceCount()(*int32) {
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
func (m *DeviceComplianceScriptRunSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetIssueDetectedDeviceCount gets the issueDetectedDeviceCount property value. Number of devices for which the detection script found an issue. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) GetIssueDetectedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("issueDetectedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetLastScriptRunDateTime gets the lastScriptRunDateTime property value. Last run time for the script across all devices
func (m *DeviceComplianceScriptRunSummary) GetLastScriptRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastScriptRunDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetNoIssueDetectedDeviceCount gets the noIssueDetectedDeviceCount property value. Number of devices for which the detection script did not find an issue and the device is healthy. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) GetNoIssueDetectedDeviceCount()(*int32) {
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
func (m *DeviceComplianceScriptRunSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceComplianceScriptRunSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetectionScriptErrorDeviceCount sets the detectionScriptErrorDeviceCount property value. Number of devices on which the detection script execution encountered an error and did not complete. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) SetDetectionScriptErrorDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("detectionScriptErrorDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionScriptPendingDeviceCount sets the detectionScriptPendingDeviceCount property value. Number of devices which have not yet run the latest version of the device compliance script. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) SetDetectionScriptPendingDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("detectionScriptPendingDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIssueDetectedDeviceCount sets the issueDetectedDeviceCount property value. Number of devices for which the detection script found an issue. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) SetIssueDetectedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("issueDetectedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetLastScriptRunDateTime sets the lastScriptRunDateTime property value. Last run time for the script across all devices
func (m *DeviceComplianceScriptRunSummary) SetLastScriptRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastScriptRunDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetNoIssueDetectedDeviceCount sets the noIssueDetectedDeviceCount property value. Number of devices for which the detection script did not find an issue and the device is healthy. Valid values -2147483648 to 2147483647
func (m *DeviceComplianceScriptRunSummary) SetNoIssueDetectedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("noIssueDetectedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceComplianceScriptRunSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DeviceComplianceScriptRunSummaryable 
type DeviceComplianceScriptRunSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetectionScriptErrorDeviceCount()(*int32)
    GetDetectionScriptPendingDeviceCount()(*int32)
    GetIssueDetectedDeviceCount()(*int32)
    GetLastScriptRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNoIssueDetectedDeviceCount()(*int32)
    GetOdataType()(*string)
    SetDetectionScriptErrorDeviceCount(value *int32)()
    SetDetectionScriptPendingDeviceCount(value *int32)()
    SetIssueDetectedDeviceCount(value *int32)()
    SetLastScriptRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNoIssueDetectedDeviceCount(value *int32)()
    SetOdataType(value *string)()
}
