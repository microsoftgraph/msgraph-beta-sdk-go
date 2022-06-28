package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUniversalAppXContainedApp 
type WindowsUniversalAppXContainedApp struct {
    MobileContainedApp
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The app user model ID of the contained app of a WindowsUniversalAppX app.
    appUserModelId *string
}
// NewWindowsUniversalAppXContainedApp instantiates a new WindowsUniversalAppXContainedApp and sets the default values.
func NewWindowsUniversalAppXContainedApp()(*WindowsUniversalAppXContainedApp) {
    m := &WindowsUniversalAppXContainedApp{
        MobileContainedApp: *NewMobileContainedApp(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsUniversalAppXContainedAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUniversalAppXContainedAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUniversalAppXContainedApp(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUniversalAppXContainedApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppUserModelId gets the appUserModelId property value. The app user model ID of the contained app of a WindowsUniversalAppX app.
func (m *WindowsUniversalAppXContainedApp) GetAppUserModelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appUserModelId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUniversalAppXContainedApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileContainedApp.GetFieldDeserializers()
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
    return res
}
// Serialize serializes information the current object
func (m *WindowsUniversalAppXContainedApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileContainedApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appUserModelId", m.GetAppUserModelId())
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
func (m *WindowsUniversalAppXContainedApp) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppUserModelId sets the appUserModelId property value. The app user model ID of the contained app of a WindowsUniversalAppX app.
func (m *WindowsUniversalAppXContainedApp) SetAppUserModelId(value *string)() {
    if m != nil {
        m.appUserModelId = value
    }
}
