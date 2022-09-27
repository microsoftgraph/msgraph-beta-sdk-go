package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsStoreApp 
type WindowsStoreApp struct {
    MobileApp
    // The Windows app store URL.
    appStoreUrl *string
}
// NewWindowsStoreApp instantiates a new WindowsStoreApp and sets the default values.
func NewWindowsStoreApp()(*WindowsStoreApp) {
    m := &WindowsStoreApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsStoreApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsStoreAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsStoreAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsStoreApp(), nil
}
// GetAppStoreUrl gets the appStoreUrl property value. The Windows app store URL.
func (m *WindowsStoreApp) GetAppStoreUrl()(*string) {
    return m.appStoreUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsStoreApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appStoreUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppStoreUrl)
    return res
}
// Serialize serializes information the current object
func (m *WindowsStoreApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appStoreUrl", m.GetAppStoreUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppStoreUrl sets the appStoreUrl property value. The Windows app store URL.
func (m *WindowsStoreApp) SetAppStoreUrl(value *string)() {
    m.appStoreUrl = value
}
