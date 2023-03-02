package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptDeviceState contains properties for device run state of the device health script.
type DeviceHealthScriptDeviceState struct {
    Entity
}
// NewDeviceHealthScriptDeviceState instantiates a new deviceHealthScriptDeviceState and sets the default values.
func NewDeviceHealthScriptDeviceState()(*DeviceHealthScriptDeviceState) {
    m := &DeviceHealthScriptDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceHealthScriptDeviceStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptDeviceStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptDeviceState(), nil
}
// GetAssignmentFilterIds gets the assignmentFilterIds property value. A list of the assignment filter ids used for health script applicability evaluation
func (m *DeviceHealthScriptDeviceState) GetAssignmentFilterIds()([]string) {
    val, err := m.GetBackingStore().Get("assignmentFilterIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDetectionState gets the detectionState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptDeviceState) GetDetectionState()(*RunState) {
    val, err := m.GetBackingStore().Get("detectionState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RunState)
    }
    return nil
}
// GetExpectedStateUpdateDateTime gets the expectedStateUpdateDateTime property value. The next timestamp of when the device health script is expected to execute
func (m *DeviceHealthScriptDeviceState) GetExpectedStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expectedStateUpdateDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptDeviceState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentFilterIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssignmentFilterIds(res)
        }
        return nil
    }
    res["detectionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionState(val.(*RunState))
        }
        return nil
    }
    res["expectedStateUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpectedStateUpdateDateTime(val)
        }
        return nil
    }
    res["lastStateUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastStateUpdateDateTime(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["managedDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDevice(val.(ManagedDeviceable))
        }
        return nil
    }
    res["postRemediationDetectionScriptError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostRemediationDetectionScriptError(val)
        }
        return nil
    }
    res["postRemediationDetectionScriptOutput"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostRemediationDetectionScriptOutput(val)
        }
        return nil
    }
    res["preRemediationDetectionScriptError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreRemediationDetectionScriptError(val)
        }
        return nil
    }
    res["preRemediationDetectionScriptOutput"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreRemediationDetectionScriptOutput(val)
        }
        return nil
    }
    res["remediationScriptError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationScriptError(val)
        }
        return nil
    }
    res["remediationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRemediationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationState(val.(*RemediationState))
        }
        return nil
    }
    return res
}
// GetLastStateUpdateDateTime gets the lastStateUpdateDateTime property value. The last timestamp of when the device health script executed
func (m *DeviceHealthScriptDeviceState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastStateUpdateDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The last time that Intune Managment Extension synced with Intune
func (m *DeviceHealthScriptDeviceState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastSyncDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetManagedDevice gets the managedDevice property value. The managed device on which the device health script executed
func (m *DeviceHealthScriptDeviceState) GetManagedDevice()(ManagedDeviceable) {
    val, err := m.GetBackingStore().Get("managedDevice")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagedDeviceable)
    }
    return nil
}
// GetPostRemediationDetectionScriptError gets the postRemediationDetectionScriptError property value. Error from the detection script after remediation
func (m *DeviceHealthScriptDeviceState) GetPostRemediationDetectionScriptError()(*string) {
    val, err := m.GetBackingStore().Get("postRemediationDetectionScriptError")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPostRemediationDetectionScriptOutput gets the postRemediationDetectionScriptOutput property value. Detection script output after remediation
func (m *DeviceHealthScriptDeviceState) GetPostRemediationDetectionScriptOutput()(*string) {
    val, err := m.GetBackingStore().Get("postRemediationDetectionScriptOutput")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreRemediationDetectionScriptError gets the preRemediationDetectionScriptError property value. Error from the detection script before remediation
func (m *DeviceHealthScriptDeviceState) GetPreRemediationDetectionScriptError()(*string) {
    val, err := m.GetBackingStore().Get("preRemediationDetectionScriptError")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreRemediationDetectionScriptOutput gets the preRemediationDetectionScriptOutput property value. Output of the detection script before remediation
func (m *DeviceHealthScriptDeviceState) GetPreRemediationDetectionScriptOutput()(*string) {
    val, err := m.GetBackingStore().Get("preRemediationDetectionScriptOutput")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediationScriptError gets the remediationScriptError property value. Error output of the remediation script
func (m *DeviceHealthScriptDeviceState) GetRemediationScriptError()(*string) {
    val, err := m.GetBackingStore().Get("remediationScriptError")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediationState gets the remediationState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptDeviceState) GetRemediationState()(*RemediationState) {
    val, err := m.GetBackingStore().Get("remediationState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RemediationState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptDeviceState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignmentFilterIds() != nil {
        err = writer.WriteCollectionOfStringValues("assignmentFilterIds", m.GetAssignmentFilterIds())
        if err != nil {
            return err
        }
    }
    if m.GetDetectionState() != nil {
        cast := (*m.GetDetectionState()).String()
        err = writer.WriteStringValue("detectionState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expectedStateUpdateDateTime", m.GetExpectedStateUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastStateUpdateDateTime", m.GetLastStateUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managedDevice", m.GetManagedDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postRemediationDetectionScriptError", m.GetPostRemediationDetectionScriptError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postRemediationDetectionScriptOutput", m.GetPostRemediationDetectionScriptOutput())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preRemediationDetectionScriptError", m.GetPreRemediationDetectionScriptError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preRemediationDetectionScriptOutput", m.GetPreRemediationDetectionScriptOutput())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remediationScriptError", m.GetRemediationScriptError())
        if err != nil {
            return err
        }
    }
    if m.GetRemediationState() != nil {
        cast := (*m.GetRemediationState()).String()
        err = writer.WriteStringValue("remediationState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignmentFilterIds sets the assignmentFilterIds property value. A list of the assignment filter ids used for health script applicability evaluation
func (m *DeviceHealthScriptDeviceState) SetAssignmentFilterIds(value []string)() {
    err := m.GetBackingStore().Set("assignmentFilterIds", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionState sets the detectionState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptDeviceState) SetDetectionState(value *RunState)() {
    err := m.GetBackingStore().Set("detectionState", value)
    if err != nil {
        panic(err)
    }
}
// SetExpectedStateUpdateDateTime sets the expectedStateUpdateDateTime property value. The next timestamp of when the device health script is expected to execute
func (m *DeviceHealthScriptDeviceState) SetExpectedStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expectedStateUpdateDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastStateUpdateDateTime sets the lastStateUpdateDateTime property value. The last timestamp of when the device health script executed
func (m *DeviceHealthScriptDeviceState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastStateUpdateDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The last time that Intune Managment Extension synced with Intune
func (m *DeviceHealthScriptDeviceState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastSyncDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDevice sets the managedDevice property value. The managed device on which the device health script executed
func (m *DeviceHealthScriptDeviceState) SetManagedDevice(value ManagedDeviceable)() {
    err := m.GetBackingStore().Set("managedDevice", value)
    if err != nil {
        panic(err)
    }
}
// SetPostRemediationDetectionScriptError sets the postRemediationDetectionScriptError property value. Error from the detection script after remediation
func (m *DeviceHealthScriptDeviceState) SetPostRemediationDetectionScriptError(value *string)() {
    err := m.GetBackingStore().Set("postRemediationDetectionScriptError", value)
    if err != nil {
        panic(err)
    }
}
// SetPostRemediationDetectionScriptOutput sets the postRemediationDetectionScriptOutput property value. Detection script output after remediation
func (m *DeviceHealthScriptDeviceState) SetPostRemediationDetectionScriptOutput(value *string)() {
    err := m.GetBackingStore().Set("postRemediationDetectionScriptOutput", value)
    if err != nil {
        panic(err)
    }
}
// SetPreRemediationDetectionScriptError sets the preRemediationDetectionScriptError property value. Error from the detection script before remediation
func (m *DeviceHealthScriptDeviceState) SetPreRemediationDetectionScriptError(value *string)() {
    err := m.GetBackingStore().Set("preRemediationDetectionScriptError", value)
    if err != nil {
        panic(err)
    }
}
// SetPreRemediationDetectionScriptOutput sets the preRemediationDetectionScriptOutput property value. Output of the detection script before remediation
func (m *DeviceHealthScriptDeviceState) SetPreRemediationDetectionScriptOutput(value *string)() {
    err := m.GetBackingStore().Set("preRemediationDetectionScriptOutput", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationScriptError sets the remediationScriptError property value. Error output of the remediation script
func (m *DeviceHealthScriptDeviceState) SetRemediationScriptError(value *string)() {
    err := m.GetBackingStore().Set("remediationScriptError", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationState sets the remediationState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptDeviceState) SetRemediationState(value *RemediationState)() {
    err := m.GetBackingStore().Set("remediationState", value)
    if err != nil {
        panic(err)
    }
}
// DeviceHealthScriptDeviceStateable 
type DeviceHealthScriptDeviceStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentFilterIds()([]string)
    GetDetectionState()(*RunState)
    GetExpectedStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDevice()(ManagedDeviceable)
    GetPostRemediationDetectionScriptError()(*string)
    GetPostRemediationDetectionScriptOutput()(*string)
    GetPreRemediationDetectionScriptError()(*string)
    GetPreRemediationDetectionScriptOutput()(*string)
    GetRemediationScriptError()(*string)
    GetRemediationState()(*RemediationState)
    SetAssignmentFilterIds(value []string)()
    SetDetectionState(value *RunState)()
    SetExpectedStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDevice(value ManagedDeviceable)()
    SetPostRemediationDetectionScriptError(value *string)()
    SetPostRemediationDetectionScriptOutput(value *string)()
    SetPreRemediationDetectionScriptError(value *string)()
    SetPreRemediationDetectionScriptOutput(value *string)()
    SetRemediationScriptError(value *string)()
    SetRemediationState(value *RemediationState)()
}
