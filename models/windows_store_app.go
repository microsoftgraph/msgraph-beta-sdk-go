package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsStoreApp 
type WindowsStoreApp struct {
    MobileApp
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Windows app store URL.
    appStoreUrl *string
}
// NewWindowsStoreApp instantiates a new WindowsStoreApp and sets the default values.
func NewWindowsStoreApp()(*WindowsStoreApp) {
    m := &WindowsStoreApp{
        MobileApp: *NewMobileApp(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsStoreAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsStoreAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsStoreApp(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsStoreApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppStoreUrl gets the appStoreUrl property value. The Windows app store URL.
func (m *WindowsStoreApp) GetAppStoreUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appStoreUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsStoreApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appStoreUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreUrl(val)
        }
        return nil
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsStoreApp) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppStoreUrl sets the appStoreUrl property value. The Windows app store URL.
func (m *WindowsStoreApp) SetAppStoreUrl(value *string)() {
    if m != nil {
        m.appStoreUrl = value
    }
}
