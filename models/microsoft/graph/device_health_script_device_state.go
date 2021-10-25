package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceHealthScriptDeviceState struct {
    Entity
    assignmentFilterIds []string;
    detectionState *RunState;
    expectedStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    managedDevice *ManagedDevice;
    postRemediationDetectionScriptError *string;
    postRemediationDetectionScriptOutput *string;
    preRemediationDetectionScriptError *string;
    preRemediationDetectionScriptOutput *string;
    remediationScriptError *string;
    remediationState *RemediationState;
}
func NewDeviceHealthScriptDeviceState()(*DeviceHealthScriptDeviceState) {
    m := &DeviceHealthScriptDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceHealthScriptDeviceState) GetAssignmentFilterIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterIds
    }
}
func (m *DeviceHealthScriptDeviceState) GetDetectionState()(*RunState) {
    if m == nil {
        return nil
    } else {
        return m.detectionState
    }
}
func (m *DeviceHealthScriptDeviceState) GetExpectedStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expectedStateUpdateDateTime
    }
}
func (m *DeviceHealthScriptDeviceState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastStateUpdateDateTime
    }
}
func (m *DeviceHealthScriptDeviceState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
func (m *DeviceHealthScriptDeviceState) GetManagedDevice()(*ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.managedDevice
    }
}
func (m *DeviceHealthScriptDeviceState) GetPostRemediationDetectionScriptError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postRemediationDetectionScriptError
    }
}
func (m *DeviceHealthScriptDeviceState) GetPostRemediationDetectionScriptOutput()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postRemediationDetectionScriptOutput
    }
}
func (m *DeviceHealthScriptDeviceState) GetPreRemediationDetectionScriptError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preRemediationDetectionScriptError
    }
}
func (m *DeviceHealthScriptDeviceState) GetPreRemediationDetectionScriptOutput()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preRemediationDetectionScriptOutput
    }
}
func (m *DeviceHealthScriptDeviceState) GetRemediationScriptError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediationScriptError
    }
}
func (m *DeviceHealthScriptDeviceState) GetRemediationState()(*RemediationState) {
    if m == nil {
        return nil
    } else {
        return m.remediationState
    }
}
func (m *DeviceHealthScriptDeviceState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentFilterIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAssignmentFilterIds(res)
        return nil
    }
    res["detectionState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunState)
        if err != nil {
            return err
        }
        cast := val.(RunState)
        m.SetDetectionState(&cast)
        return nil
    }
    res["expectedStateUpdateDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpectedStateUpdateDateTime(val)
        return nil
    }
    res["lastStateUpdateDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastStateUpdateDateTime(val)
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSyncDateTime(val)
        return nil
    }
    res["managedDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDevice() })
        if err != nil {
            return err
        }
        m.SetManagedDevice(val.(*ManagedDevice))
        return nil
    }
    res["postRemediationDetectionScriptError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostRemediationDetectionScriptError(val)
        return nil
    }
    res["postRemediationDetectionScriptOutput"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostRemediationDetectionScriptOutput(val)
        return nil
    }
    res["preRemediationDetectionScriptError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreRemediationDetectionScriptError(val)
        return nil
    }
    res["preRemediationDetectionScriptOutput"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreRemediationDetectionScriptOutput(val)
        return nil
    }
    res["remediationScriptError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRemediationScriptError(val)
        return nil
    }
    res["remediationState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRemediationState)
        if err != nil {
            return err
        }
        cast := val.(RemediationState)
        m.SetRemediationState(&cast)
        return nil
    }
    return res
}
func (m *DeviceHealthScriptDeviceState) IsNil()(bool) {
    return m == nil
}
func (m *DeviceHealthScriptDeviceState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("assignmentFilterIds", m.GetAssignmentFilterIds())
        if err != nil {
            return err
        }
    }
    if m.GetDetectionState() != nil {
        cast := m.GetDetectionState().String()
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
        cast := m.GetRemediationState().String()
        err = writer.WriteStringValue("remediationState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceHealthScriptDeviceState) SetAssignmentFilterIds(value []string)() {
    m.assignmentFilterIds = value
}
func (m *DeviceHealthScriptDeviceState) SetDetectionState(value *RunState)() {
    m.detectionState = value
}
func (m *DeviceHealthScriptDeviceState) SetExpectedStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expectedStateUpdateDateTime = value
}
func (m *DeviceHealthScriptDeviceState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastStateUpdateDateTime = value
}
func (m *DeviceHealthScriptDeviceState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
func (m *DeviceHealthScriptDeviceState) SetManagedDevice(value *ManagedDevice)() {
    m.managedDevice = value
}
func (m *DeviceHealthScriptDeviceState) SetPostRemediationDetectionScriptError(value *string)() {
    m.postRemediationDetectionScriptError = value
}
func (m *DeviceHealthScriptDeviceState) SetPostRemediationDetectionScriptOutput(value *string)() {
    m.postRemediationDetectionScriptOutput = value
}
func (m *DeviceHealthScriptDeviceState) SetPreRemediationDetectionScriptError(value *string)() {
    m.preRemediationDetectionScriptError = value
}
func (m *DeviceHealthScriptDeviceState) SetPreRemediationDetectionScriptOutput(value *string)() {
    m.preRemediationDetectionScriptOutput = value
}
func (m *DeviceHealthScriptDeviceState) SetRemediationScriptError(value *string)() {
    m.remediationScriptError = value
}
func (m *DeviceHealthScriptDeviceState) SetRemediationState(value *RemediationState)() {
    m.remediationState = value
}
