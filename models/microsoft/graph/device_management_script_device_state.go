package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementScriptDeviceState 
type DeviceManagementScriptDeviceState struct {
    Entity
    // Error code corresponding to erroneous execution of the device management script.
    errorCode *int32;
    // Error description corresponding to erroneous execution of the device management script.
    errorDescription *string;
    // Latest time the device management script executes.
    lastStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The managed devices that executes the device management script.
    managedDevice *ManagedDevice;
    // Details of execution output.
    resultMessage *string;
    // State of latest run of the device management script. Possible values are: unknown, success, fail, scriptError, pending, notApplicable.
    runState *RunState;
}
// NewDeviceManagementScriptDeviceState instantiates a new deviceManagementScriptDeviceState and sets the default values.
func NewDeviceManagementScriptDeviceState()(*DeviceManagementScriptDeviceState) {
    m := &DeviceManagementScriptDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// GetErrorCode gets the errorCode property value. Error code corresponding to erroneous execution of the device management script.
func (m *DeviceManagementScriptDeviceState) GetErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// GetErrorDescription gets the errorDescription property value. Error description corresponding to erroneous execution of the device management script.
func (m *DeviceManagementScriptDeviceState) GetErrorDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorDescription
    }
}
// GetLastStateUpdateDateTime gets the lastStateUpdateDateTime property value. Latest time the device management script executes.
func (m *DeviceManagementScriptDeviceState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastStateUpdateDateTime
    }
}
// GetManagedDevice gets the managedDevice property value. The managed devices that executes the device management script.
func (m *DeviceManagementScriptDeviceState) GetManagedDevice()(*ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.managedDevice
    }
}
// GetResultMessage gets the resultMessage property value. Details of execution output.
func (m *DeviceManagementScriptDeviceState) GetResultMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resultMessage
    }
}
// GetRunState gets the runState property value. State of latest run of the device management script. Possible values are: unknown, success, fail, scriptError, pending, notApplicable.
func (m *DeviceManagementScriptDeviceState) GetRunState()(*RunState) {
    if m == nil {
        return nil
    } else {
        return m.runState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementScriptDeviceState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["errorDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDescription(val)
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
    res["resultMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultMessage(val)
        }
        return nil
    }
    res["runState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunState(val.(*RunState))
        }
        return nil
    }
    return res
}
func (m *DeviceManagementScriptDeviceState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptDeviceState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorDescription", m.GetErrorDescription())
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
        err = writer.WriteObjectValue("managedDevice", m.GetManagedDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resultMessage", m.GetResultMessage())
        if err != nil {
            return err
        }
    }
    if m.GetRunState() != nil {
        cast := (*m.GetRunState()).String()
        err = writer.WriteStringValue("runState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetErrorCode sets the errorCode property value. Error code corresponding to erroneous execution of the device management script.
func (m *DeviceManagementScriptDeviceState) SetErrorCode(value *int32)() {
    if m != nil {
        m.errorCode = value
    }
}
// SetErrorDescription sets the errorDescription property value. Error description corresponding to erroneous execution of the device management script.
func (m *DeviceManagementScriptDeviceState) SetErrorDescription(value *string)() {
    if m != nil {
        m.errorDescription = value
    }
}
// SetLastStateUpdateDateTime sets the lastStateUpdateDateTime property value. Latest time the device management script executes.
func (m *DeviceManagementScriptDeviceState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastStateUpdateDateTime = value
    }
}
// SetManagedDevice sets the managedDevice property value. The managed devices that executes the device management script.
func (m *DeviceManagementScriptDeviceState) SetManagedDevice(value *ManagedDevice)() {
    if m != nil {
        m.managedDevice = value
    }
}
// SetResultMessage sets the resultMessage property value. Details of execution output.
func (m *DeviceManagementScriptDeviceState) SetResultMessage(value *string)() {
    if m != nil {
        m.resultMessage = value
    }
}
// SetRunState sets the runState property value. State of latest run of the device management script. Possible values are: unknown, success, fail, scriptError, pending, notApplicable.
func (m *DeviceManagementScriptDeviceState) SetRunState(value *RunState)() {
    if m != nil {
        m.runState = value
    }
}
