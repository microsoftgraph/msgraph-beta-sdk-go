package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUniversalAppXAppAssignmentSettings 
type WindowsUniversalAppXAppAssignmentSettings struct {
    MobileAppAssignmentSettings
    // Whether or not to use device execution context for Windows Universal AppX mobile app.
    useDeviceContext *bool
}
// NewWindowsUniversalAppXAppAssignmentSettings instantiates a new WindowsUniversalAppXAppAssignmentSettings and sets the default values.
func NewWindowsUniversalAppXAppAssignmentSettings()(*WindowsUniversalAppXAppAssignmentSettings) {
    m := &WindowsUniversalAppXAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    odataTypeValue := "#microsoft.graph.windowsUniversalAppXAppAssignmentSettings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsUniversalAppXAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUniversalAppXAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUniversalAppXAppAssignmentSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUniversalAppXAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppAssignmentSettings.GetFieldDeserializers()
    res["useDeviceContext"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetUseDeviceContext)
    return res
}
// GetUseDeviceContext gets the useDeviceContext property value. Whether or not to use device execution context for Windows Universal AppX mobile app.
func (m *WindowsUniversalAppXAppAssignmentSettings) GetUseDeviceContext()(*bool) {
    return m.useDeviceContext
}
// Serialize serializes information the current object
func (m *WindowsUniversalAppXAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("useDeviceContext", m.GetUseDeviceContext())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUseDeviceContext sets the useDeviceContext property value. Whether or not to use device execution context for Windows Universal AppX mobile app.
func (m *WindowsUniversalAppXAppAssignmentSettings) SetUseDeviceContext(value *bool)() {
    m.useDeviceContext = value
}
