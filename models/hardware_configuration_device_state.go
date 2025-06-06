// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HardwareConfigurationDeviceState contains properties for device run state of the hardware configuration
type HardwareConfigurationDeviceState struct {
    Entity
}
// NewHardwareConfigurationDeviceState instantiates a new HardwareConfigurationDeviceState and sets the default values.
func NewHardwareConfigurationDeviceState()(*HardwareConfigurationDeviceState) {
    m := &HardwareConfigurationDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateHardwareConfigurationDeviceStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHardwareConfigurationDeviceStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHardwareConfigurationDeviceState(), nil
}
// GetAssignmentFilterIds gets the assignmentFilterIds property value. A list of identifier strings of different assignment filters applied
// returns a *string when successful
func (m *HardwareConfigurationDeviceState) GetAssignmentFilterIds()(*string) {
    val, err := m.GetBackingStore().Get("assignmentFilterIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConfigurationError gets the configurationError property value. Error from the hardware configuration execution
// returns a *string when successful
func (m *HardwareConfigurationDeviceState) GetConfigurationError()(*string) {
    val, err := m.GetBackingStore().Get("configurationError")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConfigurationOutput gets the configurationOutput property value. Output of the hardware configuration execution
// returns a *string when successful
func (m *HardwareConfigurationDeviceState) GetConfigurationOutput()(*string) {
    val, err := m.GetBackingStore().Get("configurationOutput")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConfigurationState gets the configurationState property value. Indicates the type of execution status of the device management script.
// returns a *RunState when successful
func (m *HardwareConfigurationDeviceState) GetConfigurationState()(*RunState) {
    val, err := m.GetBackingStore().Get("configurationState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RunState)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The name of the device
// returns a *string when successful
func (m *HardwareConfigurationDeviceState) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HardwareConfigurationDeviceState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentFilterIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterIds(val)
        }
        return nil
    }
    res["configurationError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationError(val)
        }
        return nil
    }
    res["configurationOutput"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationOutput(val)
        }
        return nil
    }
    res["configurationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationState(val.(*RunState))
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["internalVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalVersion(val)
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
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["upn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpn(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetInternalVersion gets the internalVersion property value. The Policy internal version
// returns a *int32 when successful
func (m *HardwareConfigurationDeviceState) GetInternalVersion()(*int32) {
    val, err := m.GetBackingStore().Get("internalVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetLastStateUpdateDateTime gets the lastStateUpdateDateTime property value. The last timestamp of when the hardware configuration executed
// returns a *Time when successful
func (m *HardwareConfigurationDeviceState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastStateUpdateDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. Operating system version of the device (E.g. 10.0.19042.1165, 10.0.19042.1288 etc.)
// returns a *string when successful
func (m *HardwareConfigurationDeviceState) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUpn gets the upn property value. User Principal Name (UPN).
// returns a *string when successful
func (m *HardwareConfigurationDeviceState) GetUpn()(*string) {
    val, err := m.GetBackingStore().Get("upn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. The unique identifier of the Entra user associated with the device for which policy is applied. Read-Only.
// returns a *string when successful
func (m *HardwareConfigurationDeviceState) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *HardwareConfigurationDeviceState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignmentFilterIds", m.GetAssignmentFilterIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("configurationError", m.GetConfigurationError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("configurationOutput", m.GetConfigurationOutput())
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationState() != nil {
        cast := (*m.GetConfigurationState()).String()
        err = writer.WriteStringValue("configurationState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("internalVersion", m.GetInternalVersion())
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
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("upn", m.GetUpn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignmentFilterIds sets the assignmentFilterIds property value. A list of identifier strings of different assignment filters applied
func (m *HardwareConfigurationDeviceState) SetAssignmentFilterIds(value *string)() {
    err := m.GetBackingStore().Set("assignmentFilterIds", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationError sets the configurationError property value. Error from the hardware configuration execution
func (m *HardwareConfigurationDeviceState) SetConfigurationError(value *string)() {
    err := m.GetBackingStore().Set("configurationError", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationOutput sets the configurationOutput property value. Output of the hardware configuration execution
func (m *HardwareConfigurationDeviceState) SetConfigurationOutput(value *string)() {
    err := m.GetBackingStore().Set("configurationOutput", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationState sets the configurationState property value. Indicates the type of execution status of the device management script.
func (m *HardwareConfigurationDeviceState) SetConfigurationState(value *RunState)() {
    err := m.GetBackingStore().Set("configurationState", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The name of the device
func (m *HardwareConfigurationDeviceState) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetInternalVersion sets the internalVersion property value. The Policy internal version
func (m *HardwareConfigurationDeviceState) SetInternalVersion(value *int32)() {
    err := m.GetBackingStore().Set("internalVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetLastStateUpdateDateTime sets the lastStateUpdateDateTime property value. The last timestamp of when the hardware configuration executed
func (m *HardwareConfigurationDeviceState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastStateUpdateDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. Operating system version of the device (E.g. 10.0.19042.1165, 10.0.19042.1288 etc.)
func (m *HardwareConfigurationDeviceState) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetUpn sets the upn property value. User Principal Name (UPN).
func (m *HardwareConfigurationDeviceState) SetUpn(value *string)() {
    err := m.GetBackingStore().Set("upn", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The unique identifier of the Entra user associated with the device for which policy is applied. Read-Only.
func (m *HardwareConfigurationDeviceState) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
type HardwareConfigurationDeviceStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentFilterIds()(*string)
    GetConfigurationError()(*string)
    GetConfigurationOutput()(*string)
    GetConfigurationState()(*RunState)
    GetDeviceName()(*string)
    GetInternalVersion()(*int32)
    GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOsVersion()(*string)
    GetUpn()(*string)
    GetUserId()(*string)
    SetAssignmentFilterIds(value *string)()
    SetConfigurationError(value *string)()
    SetConfigurationOutput(value *string)()
    SetConfigurationState(value *RunState)()
    SetDeviceName(value *string)()
    SetInternalVersion(value *int32)()
    SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOsVersion(value *string)()
    SetUpn(value *string)()
    SetUserId(value *string)()
}
