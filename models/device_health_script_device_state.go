package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptDeviceState contains properties for device run state of the device health script.
type DeviceHealthScriptDeviceState struct {
    Entity
    // A list of the assignment filter ids used for health script applicability evaluation
    assignmentFilterIds []string
    // Indicates the type of execution status of the device management script.
    detectionState *RunState
    // The next timestamp of when the device health script is expected to execute
    expectedStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last timestamp of when the device health script executed
    lastStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last time that Intune Managment Extension synced with Intune
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The managed device on which the device health script executed
    managedDevice ManagedDeviceable
    // Error from the detection script after remediation
    postRemediationDetectionScriptError *string
    // Detection script output after remediation
    postRemediationDetectionScriptOutput *string
    // Error from the detection script before remediation
    preRemediationDetectionScriptError *string
    // Output of the detection script before remediation
    preRemediationDetectionScriptOutput *string
    // Error output of the remediation script
    remediationScriptError *string
    // Indicates the type of execution status of the device management script.
    remediationState *RemediationState
}
// NewDeviceHealthScriptDeviceState instantiates a new deviceHealthScriptDeviceState and sets the default values.
func NewDeviceHealthScriptDeviceState()(*DeviceHealthScriptDeviceState) {
    m := &DeviceHealthScriptDeviceState{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceHealthScriptDeviceState";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceHealthScriptDeviceStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptDeviceStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptDeviceState(), nil
}
// GetAssignmentFilterIds gets the assignmentFilterIds property value. A list of the assignment filter ids used for health script applicability evaluation
func (m *DeviceHealthScriptDeviceState) GetAssignmentFilterIds()([]string) {
    return m.assignmentFilterIds
}
// GetDetectionState gets the detectionState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptDeviceState) GetDetectionState()(*RunState) {
    return m.detectionState
}
// GetExpectedStateUpdateDateTime gets the expectedStateUpdateDateTime property value. The next timestamp of when the device health script is expected to execute
func (m *DeviceHealthScriptDeviceState) GetExpectedStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expectedStateUpdateDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptDeviceState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentFilterIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAssignmentFilterIds)
    res["detectionState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRunState , m.SetDetectionState)
    res["expectedStateUpdateDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetExpectedStateUpdateDateTime)
    res["lastStateUpdateDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastStateUpdateDateTime)
    res["lastSyncDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastSyncDateTime)
    res["managedDevice"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateManagedDeviceFromDiscriminatorValue , m.SetManagedDevice)
    res["postRemediationDetectionScriptError"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPostRemediationDetectionScriptError)
    res["postRemediationDetectionScriptOutput"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPostRemediationDetectionScriptOutput)
    res["preRemediationDetectionScriptError"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPreRemediationDetectionScriptError)
    res["preRemediationDetectionScriptOutput"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPreRemediationDetectionScriptOutput)
    res["remediationScriptError"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRemediationScriptError)
    res["remediationState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRemediationState , m.SetRemediationState)
    return res
}
// GetLastStateUpdateDateTime gets the lastStateUpdateDateTime property value. The last timestamp of when the device health script executed
func (m *DeviceHealthScriptDeviceState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastStateUpdateDateTime
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The last time that Intune Managment Extension synced with Intune
func (m *DeviceHealthScriptDeviceState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetManagedDevice gets the managedDevice property value. The managed device on which the device health script executed
func (m *DeviceHealthScriptDeviceState) GetManagedDevice()(ManagedDeviceable) {
    return m.managedDevice
}
// GetPostRemediationDetectionScriptError gets the postRemediationDetectionScriptError property value. Error from the detection script after remediation
func (m *DeviceHealthScriptDeviceState) GetPostRemediationDetectionScriptError()(*string) {
    return m.postRemediationDetectionScriptError
}
// GetPostRemediationDetectionScriptOutput gets the postRemediationDetectionScriptOutput property value. Detection script output after remediation
func (m *DeviceHealthScriptDeviceState) GetPostRemediationDetectionScriptOutput()(*string) {
    return m.postRemediationDetectionScriptOutput
}
// GetPreRemediationDetectionScriptError gets the preRemediationDetectionScriptError property value. Error from the detection script before remediation
func (m *DeviceHealthScriptDeviceState) GetPreRemediationDetectionScriptError()(*string) {
    return m.preRemediationDetectionScriptError
}
// GetPreRemediationDetectionScriptOutput gets the preRemediationDetectionScriptOutput property value. Output of the detection script before remediation
func (m *DeviceHealthScriptDeviceState) GetPreRemediationDetectionScriptOutput()(*string) {
    return m.preRemediationDetectionScriptOutput
}
// GetRemediationScriptError gets the remediationScriptError property value. Error output of the remediation script
func (m *DeviceHealthScriptDeviceState) GetRemediationScriptError()(*string) {
    return m.remediationScriptError
}
// GetRemediationState gets the remediationState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptDeviceState) GetRemediationState()(*RemediationState) {
    return m.remediationState
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
    m.assignmentFilterIds = value
}
// SetDetectionState sets the detectionState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptDeviceState) SetDetectionState(value *RunState)() {
    m.detectionState = value
}
// SetExpectedStateUpdateDateTime sets the expectedStateUpdateDateTime property value. The next timestamp of when the device health script is expected to execute
func (m *DeviceHealthScriptDeviceState) SetExpectedStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expectedStateUpdateDateTime = value
}
// SetLastStateUpdateDateTime sets the lastStateUpdateDateTime property value. The last timestamp of when the device health script executed
func (m *DeviceHealthScriptDeviceState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastStateUpdateDateTime = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The last time that Intune Managment Extension synced with Intune
func (m *DeviceHealthScriptDeviceState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetManagedDevice sets the managedDevice property value. The managed device on which the device health script executed
func (m *DeviceHealthScriptDeviceState) SetManagedDevice(value ManagedDeviceable)() {
    m.managedDevice = value
}
// SetPostRemediationDetectionScriptError sets the postRemediationDetectionScriptError property value. Error from the detection script after remediation
func (m *DeviceHealthScriptDeviceState) SetPostRemediationDetectionScriptError(value *string)() {
    m.postRemediationDetectionScriptError = value
}
// SetPostRemediationDetectionScriptOutput sets the postRemediationDetectionScriptOutput property value. Detection script output after remediation
func (m *DeviceHealthScriptDeviceState) SetPostRemediationDetectionScriptOutput(value *string)() {
    m.postRemediationDetectionScriptOutput = value
}
// SetPreRemediationDetectionScriptError sets the preRemediationDetectionScriptError property value. Error from the detection script before remediation
func (m *DeviceHealthScriptDeviceState) SetPreRemediationDetectionScriptError(value *string)() {
    m.preRemediationDetectionScriptError = value
}
// SetPreRemediationDetectionScriptOutput sets the preRemediationDetectionScriptOutput property value. Output of the detection script before remediation
func (m *DeviceHealthScriptDeviceState) SetPreRemediationDetectionScriptOutput(value *string)() {
    m.preRemediationDetectionScriptOutput = value
}
// SetRemediationScriptError sets the remediationScriptError property value. Error output of the remediation script
func (m *DeviceHealthScriptDeviceState) SetRemediationScriptError(value *string)() {
    m.remediationScriptError = value
}
// SetRemediationState sets the remediationState property value. Indicates the type of execution status of the device management script.
func (m *DeviceHealthScriptDeviceState) SetRemediationState(value *RemediationState)() {
    m.remediationState = value
}
