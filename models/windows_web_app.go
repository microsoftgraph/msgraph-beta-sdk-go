package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsWebApp 
type WindowsWebApp struct {
    MobileApp
    // The web app URL.
    appUrl *string
}
// NewWindowsWebApp instantiates a new WindowsWebApp and sets the default values.
func NewWindowsWebApp()(*WindowsWebApp) {
    m := &WindowsWebApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsWebApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsWebAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsWebAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsWebApp(), nil
}
// GetAppUrl gets the appUrl property value. The web app URL.
func (m *WindowsWebApp) GetAppUrl()(*string) {
    return m.appUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsWebApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppUrl)
    return res
}
// Serialize serializes information the current object
func (m *WindowsWebApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appUrl", m.GetAppUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppUrl sets the appUrl property value. The web app URL.
func (m *WindowsWebApp) SetAppUrl(value *string)() {
    m.appUrl = value
}
