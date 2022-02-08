package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceHealthScriptDeviceState 
type DeviceHealthScriptDeviceState struct {
    Entity
    // A list of the assignment filter ids used for health script applicability evaluation
    assignmentFilterIds []string;
    // Detection state from the lastest device health script execution. Possible values are: unknown, success, fail, scriptError, pending, notApplicable.
    detectionState *RunState;
    // The next timestamp of when the device health script is expected to execute
    expectedStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The last timestamp of when the device health script executed
    lastStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The last time that Intune Managment Extension synced with Intune
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The managed device on which the device health script executed
    managedDevice *ManagedDevice;
    // Error from the detection script after remediation
    postRemediationDetectionScriptError *string;
    // Detection script output after remediation
    postRemediationDetectionScriptOutput *string;
    // Error from the detection script before remediation
    preRemediationDetectionScriptError *string;
    // Output of the detection script before remediation
    preRemediationDetectionScriptOutput *string;
    // Error output of the remediation script
    remediationScriptError *string;
    // Remediation state from the lastest device health script execution. Possible values are: unknown, skipped, success, remediationFailed, scriptError.
    remediationState *RemediationState;
}
// NewDeviceHealthScriptDeviceState instantiates a new deviceHealthScriptDeviceState and sets the default values.
func NewDeviceHealthScriptDeviceState()(*DeviceHealthScriptDeviceState) {
    m := &DeviceHealthScriptDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignmentFilterIds gets the assignmentFilterIds property value. A list of the assignment filter ids used for health script applicability evaluation
func (m *DeviceHealthScriptDeviceState) GetAssignmentFilterIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterIds
    }
}
// GetDetectionState gets the detectionState property value. Detection state from the lastest device health script execution. Possible values are: unknown, success, fail, scriptError, pending, notApplicable.
func (m *DeviceHealthScriptDeviceState) GetDetectionState()(*RunState) {
    if m == nil {
        return nil
    } else {
        return m.detectionState
    }
}
// GetExpectedStateUpdateDateTime gets the expectedStateUpdateDateTime property value. The next timestamp of when the device health script is expected to execute
func (m *DeviceHealthScriptDeviceState) GetExpectedStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expectedStateUpdateDateTime
    }
}
// GetLastStateUpdateDateTime gets the lastStateUpdateDateTime property value. The last timestamp of when the device health script executed
func (m *DeviceHealthScriptDeviceState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastStateUpdateDateTime
    }
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The last time that Intune Managment Extension synced with Intune
func (m *DeviceHealthScriptDeviceState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetManagedDevice gets the managedDevice property value. The managed device on which the device health script executed
func (m *DeviceHealthScriptDeviceState) GetManagedDevice()(*ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.managedDevice
    }
}
// GetPostRemediationDetectionScriptError gets the postRemediationDetectionScriptError property value. Error from the detection script after remediation
func (m *DeviceHealthScriptDeviceState) GetPostRemediationDetectionScriptError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postRemediationDetectionScriptError
    }
}
// GetPostRemediationDetectionScriptOutput gets the postRemediationDetectionScriptOutput property value. Detection script output after remediation
func (m *DeviceHealthScriptDeviceState) GetPostRemediationDetectionScriptOutput()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postRemediationDetectionScriptOutput
    }
}
// GetPreRemediationDetectionScriptError gets the preRemediationDetectionScriptError property value. Error from the detection script before remediation
func (m *DeviceHealthScriptDeviceState) GetPreRemediationDetectionScriptError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preRemediationDetectionScriptError
    }
}
// GetPreRemediationDetectionScriptOutput gets the preRemediationDetectionScriptOutput property value. Output of the detection script before remediation
func (m *DeviceHealthScriptDeviceState) GetPreRemediationDetectionScriptOutput()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preRemediationDetectionScriptOutput
    }
}
// GetRemediationScriptError gets the remediationScriptError property value. Error output of the remediation script
func (m *DeviceHealthScriptDeviceState) GetRemediationScriptError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediationScriptError
    }
}
// GetRemediationState gets the remediationState property value. Remediation state from the lastest device health script execution. Possible values are: unknown, skipped, success, remediationFailed, scriptError.
func (m *DeviceHealthScriptDeviceState) GetRemediationState()(*RemediationState) {
    if m == nil {
        return nil
    } else {
        return m.remediationState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptDeviceState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentFilterIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["detectionState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionState(val.(*RunState))
        }
        return nil
    }
    res["expectedStateUpdateDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpectedStateUpdateDateTime(val)
        }
        return nil
    }
    res["lastStateUpdateDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastStateUpdateDateTime(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["managedDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDevice() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDevice(val.(*ManagedDevice))
        }
        return nil
    }
    res["postRemediationDetectionScriptError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostRemediationDetectionScriptError(val)
        }
        return nil
    }
    res["postRemediationDetectionScriptOutput"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostRemediationDetectionScriptOutput(val)
        }
        return nil
    }
    res["preRemediationDetectionScriptError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreRemediationDetectionScriptError(val)
        }
        return nil
    }
    res["preRemediationDetectionScriptOutput"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreRemediationDetectionScriptOutput(val)
        }
        return nil
    }
    res["remediationScriptError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationScriptError(val)
        }
        return nil
    }
    res["remediationState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *DeviceHealthScriptDeviceState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptDeviceState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.assignmentFilterIds = value
    }
}
// SetDetectionState sets the detectionState property value. Detection state from the lastest device health script execution. Possible values are: unknown, success, fail, scriptError, pending, notApplicable.
func (m *DeviceHealthScriptDeviceState) SetDetectionState(value *RunState)() {
    if m != nil {
        m.detectionState = value
    }
}
// SetExpectedStateUpdateDateTime sets the expectedStateUpdateDateTime property value. The next timestamp of when the device health script is expected to execute
func (m *DeviceHealthScriptDeviceState) SetExpectedStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expectedStateUpdateDateTime = value
    }
}
// SetLastStateUpdateDateTime sets the lastStateUpdateDateTime property value. The last timestamp of when the device health script executed
func (m *DeviceHealthScriptDeviceState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastStateUpdateDateTime = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The last time that Intune Managment Extension synced with Intune
func (m *DeviceHealthScriptDeviceState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetManagedDevice sets the managedDevice property value. The managed device on which the device health script executed
func (m *DeviceHealthScriptDeviceState) SetManagedDevice(value *ManagedDevice)() {
    if m != nil {
        m.managedDevice = value
    }
}
// SetPostRemediationDetectionScriptError sets the postRemediationDetectionScriptError property value. Error from the detection script after remediation
func (m *DeviceHealthScriptDeviceState) SetPostRemediationDetectionScriptError(value *string)() {
    if m != nil {
        m.postRemediationDetectionScriptError = value
    }
}
// SetPostRemediationDetectionScriptOutput sets the postRemediationDetectionScriptOutput property value. Detection script output after remediation
func (m *DeviceHealthScriptDeviceState) SetPostRemediationDetectionScriptOutput(value *string)() {
    if m != nil {
        m.postRemediationDetectionScriptOutput = value
    }
}
// SetPreRemediationDetectionScriptError sets the preRemediationDetectionScriptError property value. Error from the detection script before remediation
func (m *DeviceHealthScriptDeviceState) SetPreRemediationDetectionScriptError(value *string)() {
    if m != nil {
        m.preRemediationDetectionScriptError = value
    }
}
// SetPreRemediationDetectionScriptOutput sets the preRemediationDetectionScriptOutput property value. Output of the detection script before remediation
func (m *DeviceHealthScriptDeviceState) SetPreRemediationDetectionScriptOutput(value *string)() {
    if m != nil {
        m.preRemediationDetectionScriptOutput = value
    }
}
// SetRemediationScriptError sets the remediationScriptError property value. Error output of the remediation script
func (m *DeviceHealthScriptDeviceState) SetRemediationScriptError(value *string)() {
    if m != nil {
        m.remediationScriptError = value
    }
}
// SetRemediationState sets the remediationState property value. Remediation state from the lastest device health script execution. Possible values are: unknown, skipped, success, remediationFailed, scriptError.
func (m *DeviceHealthScriptDeviceState) SetRemediationState(value *RemediationState)() {
    if m != nil {
        m.remediationState = value
    }
}
