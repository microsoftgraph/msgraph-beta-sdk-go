package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementScriptDeviceState struct {
    Entity
    errorCode *int32;
    errorDescription *string;
    lastStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    managedDevice *ManagedDevice;
    resultMessage *string;
    runState *RunState;
}
func NewDeviceManagementScriptDeviceState()(*DeviceManagementScriptDeviceState) {
    m := &DeviceManagementScriptDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementScriptDeviceState) GetErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
func (m *DeviceManagementScriptDeviceState) GetErrorDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorDescription
    }
}
func (m *DeviceManagementScriptDeviceState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastStateUpdateDateTime
    }
}
func (m *DeviceManagementScriptDeviceState) GetManagedDevice()(*ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.managedDevice
    }
}
func (m *DeviceManagementScriptDeviceState) GetResultMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resultMessage
    }
}
func (m *DeviceManagementScriptDeviceState) GetRunState()(*RunState) {
    if m == nil {
        return nil
    } else {
        return m.runState
    }
}
func (m *DeviceManagementScriptDeviceState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorCode(val)
        return nil
    }
    res["errorDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorDescription(val)
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
    res["managedDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDevice() })
        if err != nil {
            return err
        }
        m.SetManagedDevice(val.(*ManagedDevice))
        return nil
    }
    res["resultMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResultMessage(val)
        return nil
    }
    res["runState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunState)
        if err != nil {
            return err
        }
        cast := val.(RunState)
        m.SetRunState(&cast)
        return nil
    }
    return res
}
func (m *DeviceManagementScriptDeviceState) IsNil()(bool) {
    return m == nil
}
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
        cast := m.GetRunState().String()
        err = writer.WriteStringValue("runState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceManagementScriptDeviceState) SetErrorCode(value *int32)() {
    m.errorCode = value
}
func (m *DeviceManagementScriptDeviceState) SetErrorDescription(value *string)() {
    m.errorDescription = value
}
func (m *DeviceManagementScriptDeviceState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastStateUpdateDateTime = value
}
func (m *DeviceManagementScriptDeviceState) SetManagedDevice(value *ManagedDevice)() {
    m.managedDevice = value
}
func (m *DeviceManagementScriptDeviceState) SetResultMessage(value *string)() {
    m.resultMessage = value
}
func (m *DeviceManagementScriptDeviceState) SetRunState(value *RunState)() {
    m.runState = value
}
