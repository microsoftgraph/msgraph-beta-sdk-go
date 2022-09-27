package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAppIdentifier 
type WindowsAppIdentifier struct {
    MobileAppIdentifier
    // The identifier for an app, as specified in the app store.
    windowsAppId *string
}
// NewWindowsAppIdentifier instantiates a new WindowsAppIdentifier and sets the default values.
func NewWindowsAppIdentifier()(*WindowsAppIdentifier) {
    m := &WindowsAppIdentifier{
        MobileAppIdentifier: *NewMobileAppIdentifier(),
    }
    odataTypeValue := "#microsoft.graph.windowsAppIdentifier";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsAppIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAppIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAppIdentifier(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAppIdentifier) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppIdentifier.GetFieldDeserializers()
    res["windowsAppId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWindowsAppId)
    return res
}
// GetWindowsAppId gets the windowsAppId property value. The identifier for an app, as specified in the app store.
func (m *WindowsAppIdentifier) GetWindowsAppId()(*string) {
    return m.windowsAppId
}
// Serialize serializes information the current object
func (m *WindowsAppIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppIdentifier.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("windowsAppId", m.GetWindowsAppId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetWindowsAppId sets the windowsAppId property value. The identifier for an app, as specified in the app store.
func (m *WindowsAppIdentifier) SetWindowsAppId(value *string)() {
    m.windowsAppId = value
}
