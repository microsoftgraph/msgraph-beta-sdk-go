package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftStoreForBusinessContainedApp a class that represents a contained app of a MicrosoftStoreForBusinessApp.
type MicrosoftStoreForBusinessContainedApp struct {
    MobileContainedApp
}
// NewMicrosoftStoreForBusinessContainedApp instantiates a new microsoftStoreForBusinessContainedApp and sets the default values.
func NewMicrosoftStoreForBusinessContainedApp()(*MicrosoftStoreForBusinessContainedApp) {
    m := &MicrosoftStoreForBusinessContainedApp{
        MobileContainedApp: *NewMobileContainedApp(),
    }
    odataTypeValue := "#microsoft.graph.microsoftStoreForBusinessContainedApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMicrosoftStoreForBusinessContainedAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftStoreForBusinessContainedAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftStoreForBusinessContainedApp(), nil
}
// GetAppUserModelId gets the appUserModelId property value. The app user model ID of the contained app of a MicrosoftStoreForBusinessApp.
func (m *MicrosoftStoreForBusinessContainedApp) GetAppUserModelId()(*string) {
    val, err := m.GetBackingStore().Get("appUserModelId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftStoreForBusinessContainedApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *MicrosoftStoreForBusinessContainedApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetAppUserModelId sets the appUserModelId property value. The app user model ID of the contained app of a MicrosoftStoreForBusinessApp.
func (m *MicrosoftStoreForBusinessContainedApp) SetAppUserModelId(value *string)() {
    err := m.GetBackingStore().Set("appUserModelId", value)
    if err != nil {
        panic(err)
    }
}
// MicrosoftStoreForBusinessContainedAppable 
type MicrosoftStoreForBusinessContainedAppable interface {
    MobileContainedAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppUserModelId()(*string)
    SetAppUserModelId(value *string)()
}
