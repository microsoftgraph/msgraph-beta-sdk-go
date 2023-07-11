package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsManagementAppHealthState windows management app health state entity.
type WindowsManagementAppHealthState struct {
    Entity
}
// NewWindowsManagementAppHealthState instantiates a new windowsManagementAppHealthState and sets the default values.
func NewWindowsManagementAppHealthState()(*WindowsManagementAppHealthState) {
    m := &WindowsManagementAppHealthState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsManagementAppHealthStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsManagementAppHealthStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsManagementAppHealthState(), nil
}
// GetDeviceName gets the deviceName property value. Name of the device on which Windows management app is installed.
func (m *WindowsManagementAppHealthState) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceOSVersion gets the deviceOSVersion property value. Windows 10 OS version of the device on which Windows management app is installed.
func (m *WindowsManagementAppHealthState) GetDeviceOSVersion()(*string) {
    val, err := m.GetBackingStore().Get("deviceOSVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsManagementAppHealthState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["deviceOSVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOSVersion(val)
        }
        return nil
    }
    res["healthState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthState(val.(*HealthState))
        }
        return nil
    }
    res["installedVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstalledVersion(val)
        }
        return nil
    }
    res["lastCheckInDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckInDateTime(val)
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
// GetHealthState gets the healthState property value. Indicates health state of the Windows management app.
func (m *WindowsManagementAppHealthState) GetHealthState()(*HealthState) {
    val, err := m.GetBackingStore().Get("healthState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*HealthState)
    }
    return nil
}
// GetInstalledVersion gets the installedVersion property value. Windows management app installed version.
func (m *WindowsManagementAppHealthState) GetInstalledVersion()(*string) {
    val, err := m.GetBackingStore().Get("installedVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastCheckInDateTime gets the lastCheckInDateTime property value. Windows management app last check-in time.
func (m *WindowsManagementAppHealthState) GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastCheckInDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsManagementAppHealthState) GetOdataType()(*string) {
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
func (m *WindowsManagementAppHealthState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceOSVersion", m.GetDeviceOSVersion())
        if err != nil {
            return err
        }
    }
    if m.GetHealthState() != nil {
        cast := (*m.GetHealthState()).String()
        err = writer.WriteStringValue("healthState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("installedVersion", m.GetInstalledVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastCheckInDateTime", m.GetLastCheckInDateTime())
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
// SetDeviceName sets the deviceName property value. Name of the device on which Windows management app is installed.
func (m *WindowsManagementAppHealthState) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceOSVersion sets the deviceOSVersion property value. Windows 10 OS version of the device on which Windows management app is installed.
func (m *WindowsManagementAppHealthState) SetDeviceOSVersion(value *string)() {
    err := m.GetBackingStore().Set("deviceOSVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthState sets the healthState property value. Indicates health state of the Windows management app.
func (m *WindowsManagementAppHealthState) SetHealthState(value *HealthState)() {
    err := m.GetBackingStore().Set("healthState", value)
    if err != nil {
        panic(err)
    }
}
// SetInstalledVersion sets the installedVersion property value. Windows management app installed version.
func (m *WindowsManagementAppHealthState) SetInstalledVersion(value *string)() {
    err := m.GetBackingStore().Set("installedVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetLastCheckInDateTime sets the lastCheckInDateTime property value. Windows management app last check-in time.
func (m *WindowsManagementAppHealthState) SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastCheckInDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsManagementAppHealthState) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// WindowsManagementAppHealthStateable 
type WindowsManagementAppHealthStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceName()(*string)
    GetDeviceOSVersion()(*string)
    GetHealthState()(*HealthState)
    GetInstalledVersion()(*string)
    GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    SetDeviceName(value *string)()
    SetDeviceOSVersion(value *string)()
    SetHealthState(value *HealthState)()
    SetInstalledVersion(value *string)()
    SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
}
