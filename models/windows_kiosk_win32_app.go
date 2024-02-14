package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskWin32App kioskModeApp v4 for Win32 app support
type WindowsKioskWin32App struct {
    WindowsKioskAppBase
}
// NewWindowsKioskWin32App instantiates a new WindowsKioskWin32App and sets the default values.
func NewWindowsKioskWin32App()(*WindowsKioskWin32App) {
    m := &WindowsKioskWin32App{
        WindowsKioskAppBase: *NewWindowsKioskAppBase(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskWin32App"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsKioskWin32AppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsKioskWin32AppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskWin32App(), nil
}
// GetClassicAppPath gets the classicAppPath property value. This is the classicapppath to be used by v4 Win32 app while in Kiosk Mode
// returns a *string when successful
func (m *WindowsKioskWin32App) GetClassicAppPath()(*string) {
    val, err := m.GetBackingStore().Get("classicAppPath")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEdgeKiosk gets the edgeKiosk property value. Edge kiosk (url) for Edge kiosk mode
// returns a *string when successful
func (m *WindowsKioskWin32App) GetEdgeKiosk()(*string) {
    val, err := m.GetBackingStore().Get("edgeKiosk")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEdgeKioskIdleTimeoutMinutes gets the edgeKioskIdleTimeoutMinutes property value. Edge kiosk idle timeout in minutes for Edge kiosk mode. Valid values 0 to 1440
// returns a *int32 when successful
func (m *WindowsKioskWin32App) GetEdgeKioskIdleTimeoutMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("edgeKioskIdleTimeoutMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEdgeKioskType gets the edgeKioskType property value. Edge kiosk type
// returns a *WindowsEdgeKioskType when successful
func (m *WindowsKioskWin32App) GetEdgeKioskType()(*WindowsEdgeKioskType) {
    val, err := m.GetBackingStore().Get("edgeKioskType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsEdgeKioskType)
    }
    return nil
}
// GetEdgeNoFirstRun gets the edgeNoFirstRun property value. Edge first run flag for Edge kiosk mode
// returns a *bool when successful
func (m *WindowsKioskWin32App) GetEdgeNoFirstRun()(*bool) {
    val, err := m.GetBackingStore().Get("edgeNoFirstRun")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsKioskWin32App) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskAppBase.GetFieldDeserializers()
    res["classicAppPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassicAppPath(val)
        }
        return nil
    }
    res["edgeKiosk"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeKiosk(val)
        }
        return nil
    }
    res["edgeKioskIdleTimeoutMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeKioskIdleTimeoutMinutes(val)
        }
        return nil
    }
    res["edgeKioskType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsEdgeKioskType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeKioskType(val.(*WindowsEdgeKioskType))
        }
        return nil
    }
    res["edgeNoFirstRun"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeNoFirstRun(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsKioskWin32App) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskAppBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("classicAppPath", m.GetClassicAppPath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("edgeKiosk", m.GetEdgeKiosk())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("edgeKioskIdleTimeoutMinutes", m.GetEdgeKioskIdleTimeoutMinutes())
        if err != nil {
            return err
        }
    }
    if m.GetEdgeKioskType() != nil {
        cast := (*m.GetEdgeKioskType()).String()
        err = writer.WriteStringValue("edgeKioskType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeNoFirstRun", m.GetEdgeNoFirstRun())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassicAppPath sets the classicAppPath property value. This is the classicapppath to be used by v4 Win32 app while in Kiosk Mode
func (m *WindowsKioskWin32App) SetClassicAppPath(value *string)() {
    err := m.GetBackingStore().Set("classicAppPath", value)
    if err != nil {
        panic(err)
    }
}
// SetEdgeKiosk sets the edgeKiosk property value. Edge kiosk (url) for Edge kiosk mode
func (m *WindowsKioskWin32App) SetEdgeKiosk(value *string)() {
    err := m.GetBackingStore().Set("edgeKiosk", value)
    if err != nil {
        panic(err)
    }
}
// SetEdgeKioskIdleTimeoutMinutes sets the edgeKioskIdleTimeoutMinutes property value. Edge kiosk idle timeout in minutes for Edge kiosk mode. Valid values 0 to 1440
func (m *WindowsKioskWin32App) SetEdgeKioskIdleTimeoutMinutes(value *int32)() {
    err := m.GetBackingStore().Set("edgeKioskIdleTimeoutMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetEdgeKioskType sets the edgeKioskType property value. Edge kiosk type
func (m *WindowsKioskWin32App) SetEdgeKioskType(value *WindowsEdgeKioskType)() {
    err := m.GetBackingStore().Set("edgeKioskType", value)
    if err != nil {
        panic(err)
    }
}
// SetEdgeNoFirstRun sets the edgeNoFirstRun property value. Edge first run flag for Edge kiosk mode
func (m *WindowsKioskWin32App) SetEdgeNoFirstRun(value *bool)() {
    err := m.GetBackingStore().Set("edgeNoFirstRun", value)
    if err != nil {
        panic(err)
    }
}
type WindowsKioskWin32Appable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsKioskAppBaseable
    GetClassicAppPath()(*string)
    GetEdgeKiosk()(*string)
    GetEdgeKioskIdleTimeoutMinutes()(*int32)
    GetEdgeKioskType()(*WindowsEdgeKioskType)
    GetEdgeNoFirstRun()(*bool)
    SetClassicAppPath(value *string)()
    SetEdgeKiosk(value *string)()
    SetEdgeKioskIdleTimeoutMinutes(value *int32)()
    SetEdgeKioskType(value *WindowsEdgeKioskType)()
    SetEdgeNoFirstRun(value *bool)()
}
