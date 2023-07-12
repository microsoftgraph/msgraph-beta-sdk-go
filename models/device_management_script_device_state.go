package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementScriptDeviceState contains properties for device run state of the device management script.
type DeviceManagementScriptDeviceState struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewDeviceManagementScriptDeviceState instantiates a new deviceManagementScriptDeviceState and sets the default values.
func NewDeviceManagementScriptDeviceState()(*DeviceManagementScriptDeviceState) {
    m := &DeviceManagementScriptDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementScriptDeviceState(), nil
}
// GetErrorCode gets the errorCode property value. Error code corresponding to erroneous execution of the device management script.
func (m *DeviceManagementScriptDeviceState) GetErrorCode()(*int32) {
    val, err := m.GetBackingStore().Get("errorCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetErrorDescription gets the errorDescription property value. Error description corresponding to erroneous execution of the device management script.
func (m *DeviceManagementScriptDeviceState) GetErrorDescription()(*string) {
    val, err := m.GetBackingStore().Get("errorDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementScriptDeviceState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["errorDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDescription(val)
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
    res["resultMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultMessage(val)
        }
        return nil
    }
    res["runState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetLastStateUpdateDateTime gets the lastStateUpdateDateTime property value. Latest time the device management script executes.
func (m *DeviceManagementScriptDeviceState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastStateUpdateDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetManagedDevice gets the managedDevice property value. The managed devices that executes the device management script.
func (m *DeviceManagementScriptDeviceState) GetManagedDevice()(ManagedDeviceable) {
    val, err := m.GetBackingStore().Get("managedDevice")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagedDeviceable)
    }
    return nil
}
// GetResultMessage gets the resultMessage property value. Details of execution output.
func (m *DeviceManagementScriptDeviceState) GetResultMessage()(*string) {
    val, err := m.GetBackingStore().Get("resultMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRunState gets the runState property value. Indicates the type of execution status of the device management script.
func (m *DeviceManagementScriptDeviceState) GetRunState()(*RunState) {
    val, err := m.GetBackingStore().Get("runState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RunState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptDeviceState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    err := m.GetBackingStore().Set("errorCode", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorDescription sets the errorDescription property value. Error description corresponding to erroneous execution of the device management script.
func (m *DeviceManagementScriptDeviceState) SetErrorDescription(value *string)() {
    err := m.GetBackingStore().Set("errorDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetLastStateUpdateDateTime sets the lastStateUpdateDateTime property value. Latest time the device management script executes.
func (m *DeviceManagementScriptDeviceState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastStateUpdateDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDevice sets the managedDevice property value. The managed devices that executes the device management script.
func (m *DeviceManagementScriptDeviceState) SetManagedDevice(value ManagedDeviceable)() {
    err := m.GetBackingStore().Set("managedDevice", value)
    if err != nil {
        panic(err)
    }
}
// SetResultMessage sets the resultMessage property value. Details of execution output.
func (m *DeviceManagementScriptDeviceState) SetResultMessage(value *string)() {
    err := m.GetBackingStore().Set("resultMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetRunState sets the runState property value. Indicates the type of execution status of the device management script.
func (m *DeviceManagementScriptDeviceState) SetRunState(value *RunState)() {
    err := m.GetBackingStore().Set("runState", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementScriptDeviceStateable 
type DeviceManagementScriptDeviceStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrorCode()(*int32)
    GetErrorDescription()(*string)
    GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDevice()(ManagedDeviceable)
    GetResultMessage()(*string)
    GetRunState()(*RunState)
    SetErrorCode(value *int32)()
    SetErrorDescription(value *string)()
    SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDevice(value ManagedDeviceable)()
    SetResultMessage(value *string)()
    SetRunState(value *RunState)()
}
