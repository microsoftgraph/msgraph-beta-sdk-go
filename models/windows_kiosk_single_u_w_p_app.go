package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskSingleUWPApp 
type WindowsKioskSingleUWPApp struct {
    WindowsKioskAppConfiguration
    // The uwpApp property
    uwpApp WindowsKioskUWPAppable
}
// NewWindowsKioskSingleUWPApp instantiates a new WindowsKioskSingleUWPApp and sets the default values.
func NewWindowsKioskSingleUWPApp()(*WindowsKioskSingleUWPApp) {
    m := &WindowsKioskSingleUWPApp{
        WindowsKioskAppConfiguration: *NewWindowsKioskAppConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskSingleUWPApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsKioskSingleUWPAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskSingleUWPAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskSingleUWPApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskSingleUWPApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskAppConfiguration.GetFieldDeserializers()
    res["uwpApp"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindowsKioskUWPAppFromDiscriminatorValue , m.SetUwpApp)
    return res
}
// GetUwpApp gets the uwpApp property value. The uwpApp property
func (m *WindowsKioskSingleUWPApp) GetUwpApp()(WindowsKioskUWPAppable) {
    return m.uwpApp
}
// Serialize serializes information the current object
func (m *WindowsKioskSingleUWPApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskAppConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("uwpApp", m.GetUwpApp())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUwpApp sets the uwpApp property value. The uwpApp property
func (m *WindowsKioskSingleUWPApp) SetUwpApp(value WindowsKioskUWPAppable)() {
    m.uwpApp = value
}
