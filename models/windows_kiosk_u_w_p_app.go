package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskUWPApp 
type WindowsKioskUWPApp struct {
    WindowsKioskAppBase
}
// NewWindowsKioskUWPApp instantiates a new WindowsKioskUWPApp and sets the default values.
func NewWindowsKioskUWPApp()(*WindowsKioskUWPApp) {
    m := &WindowsKioskUWPApp{
        WindowsKioskAppBase: *NewWindowsKioskAppBase(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskUWPApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsKioskUWPAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskUWPAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskUWPApp(), nil
}
// GetAppId gets the appId property value. This references an Intune App that will be target to the same assignments as Kiosk configuration
func (m *WindowsKioskUWPApp) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppUserModelId gets the appUserModelId property value. This is the only Application User Model ID (AUMID) that will be available to launch use while in Kiosk Mode
func (m *WindowsKioskUWPApp) GetAppUserModelId()(*string) {
    val, err := m.GetBackingStore().Get("appUserModelId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContainedAppId gets the containedAppId property value. This references an contained App from an Intune App
func (m *WindowsKioskUWPApp) GetContainedAppId()(*string) {
    val, err := m.GetBackingStore().Get("containedAppId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskUWPApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskAppBase.GetFieldDeserializers()
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["appUserModelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppUserModelId(val)
        }
        return nil
    }
    res["containedAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainedAppId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsKioskUWPApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskAppBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appUserModelId", m.GetAppUserModelId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("containedAppId", m.GetContainedAppId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. This references an Intune App that will be target to the same assignments as Kiosk configuration
func (m *WindowsKioskUWPApp) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetAppUserModelId sets the appUserModelId property value. This is the only Application User Model ID (AUMID) that will be available to launch use while in Kiosk Mode
func (m *WindowsKioskUWPApp) SetAppUserModelId(value *string)() {
    err := m.GetBackingStore().Set("appUserModelId", value)
    if err != nil {
        panic(err)
    }
}
// SetContainedAppId sets the containedAppId property value. This references an contained App from an Intune App
func (m *WindowsKioskUWPApp) SetContainedAppId(value *string)() {
    err := m.GetBackingStore().Set("containedAppId", value)
    if err != nil {
        panic(err)
    }
}
// WindowsKioskUWPAppable 
type WindowsKioskUWPAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsKioskAppBaseable
    GetAppId()(*string)
    GetAppUserModelId()(*string)
    GetContainedAppId()(*string)
    SetAppId(value *string)()
    SetAppUserModelId(value *string)()
    SetContainedAppId(value *string)()
}
