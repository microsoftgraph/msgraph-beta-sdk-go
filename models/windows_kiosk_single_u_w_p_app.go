package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskSingleUWPApp the class used to identify the UWP app info for the kiosk configuration
type WindowsKioskSingleUWPApp struct {
    WindowsKioskAppConfiguration
}
// NewWindowsKioskSingleUWPApp instantiates a new windowsKioskSingleUWPApp and sets the default values.
func NewWindowsKioskSingleUWPApp()(*WindowsKioskSingleUWPApp) {
    m := &WindowsKioskSingleUWPApp{
        WindowsKioskAppConfiguration: *NewWindowsKioskAppConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskSingleUWPApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsKioskSingleUWPAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskSingleUWPAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskSingleUWPApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskSingleUWPApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskAppConfiguration.GetFieldDeserializers()
    res["uwpApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsKioskUWPAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUwpApp(val.(WindowsKioskUWPAppable))
        }
        return nil
    }
    return res
}
// GetUwpApp gets the uwpApp property value. The uwpApp property
func (m *WindowsKioskSingleUWPApp) GetUwpApp()(WindowsKioskUWPAppable) {
    val, err := m.GetBackingStore().Get("uwpApp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsKioskUWPAppable)
    }
    return nil
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
    err := m.GetBackingStore().Set("uwpApp", value)
    if err != nil {
        panic(err)
    }
}
// WindowsKioskSingleUWPAppable 
type WindowsKioskSingleUWPAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsKioskAppConfigurationable
    GetUwpApp()(WindowsKioskUWPAppable)
    SetUwpApp(value WindowsKioskUWPAppable)()
}
