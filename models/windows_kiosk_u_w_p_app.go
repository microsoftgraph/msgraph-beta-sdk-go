package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskUWPApp 
type WindowsKioskUWPApp struct {
    WindowsKioskAppBase
    // This references an Intune App that will be target to the same assignments as Kiosk configuration
    appId *string
    // This is the only Application User Model ID (AUMID) that will be available to launch use while in Kiosk Mode
    appUserModelId *string
    // This references an contained App from an Intune App
    containedAppId *string
}
// NewWindowsKioskUWPApp instantiates a new WindowsKioskUWPApp and sets the default values.
func NewWindowsKioskUWPApp()(*WindowsKioskUWPApp) {
    m := &WindowsKioskUWPApp{
        WindowsKioskAppBase: *NewWindowsKioskAppBase(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskUWPApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsKioskUWPAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskUWPAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskUWPApp(), nil
}
// GetAppId gets the appId property value. This references an Intune App that will be target to the same assignments as Kiosk configuration
func (m *WindowsKioskUWPApp) GetAppId()(*string) {
    return m.appId
}
// GetAppUserModelId gets the appUserModelId property value. This is the only Application User Model ID (AUMID) that will be available to launch use while in Kiosk Mode
func (m *WindowsKioskUWPApp) GetAppUserModelId()(*string) {
    return m.appUserModelId
}
// GetContainedAppId gets the containedAppId property value. This references an contained App from an Intune App
func (m *WindowsKioskUWPApp) GetContainedAppId()(*string) {
    return m.containedAppId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskUWPApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskAppBase.GetFieldDeserializers()
    res["appId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppId)
    res["appUserModelId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppUserModelId)
    res["containedAppId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContainedAppId)
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
    m.appId = value
}
// SetAppUserModelId sets the appUserModelId property value. This is the only Application User Model ID (AUMID) that will be available to launch use while in Kiosk Mode
func (m *WindowsKioskUWPApp) SetAppUserModelId(value *string)() {
    m.appUserModelId = value
}
// SetContainedAppId sets the containedAppId property value. This references an contained App from an Intune App
func (m *WindowsKioskUWPApp) SetContainedAppId(value *string)() {
    m.containedAppId = value
}
