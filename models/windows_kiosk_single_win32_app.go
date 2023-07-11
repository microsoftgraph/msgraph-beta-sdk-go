package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskSingleWin32App the app base class used to identify the application info for the kiosk configuration
type WindowsKioskSingleWin32App struct {
    WindowsKioskAppConfiguration
}
// NewWindowsKioskSingleWin32App instantiates a new windowsKioskSingleWin32App and sets the default values.
func NewWindowsKioskSingleWin32App()(*WindowsKioskSingleWin32App) {
    m := &WindowsKioskSingleWin32App{
        WindowsKioskAppConfiguration: *NewWindowsKioskAppConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskSingleWin32App"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsKioskSingleWin32AppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskSingleWin32AppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskSingleWin32App(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskSingleWin32App) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskAppConfiguration.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["win32App"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsKioskWin32AppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWin32App(val.(WindowsKioskWin32Appable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsKioskSingleWin32App) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWin32App gets the win32App property value. The win32App property
func (m *WindowsKioskSingleWin32App) GetWin32App()(WindowsKioskWin32Appable) {
    val, err := m.GetBackingStore().Get("win32App")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsKioskWin32Appable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsKioskSingleWin32App) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskAppConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("win32App", m.GetWin32App())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsKioskSingleWin32App) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetWin32App sets the win32App property value. The win32App property
func (m *WindowsKioskSingleWin32App) SetWin32App(value WindowsKioskWin32Appable)() {
    err := m.GetBackingStore().Set("win32App", value)
    if err != nil {
        panic(err)
    }
}
// WindowsKioskSingleWin32Appable 
type WindowsKioskSingleWin32Appable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsKioskAppConfigurationable
    GetOdataType()(*string)
    GetWin32App()(WindowsKioskWin32Appable)
    SetOdataType(value *string)()
    SetWin32App(value WindowsKioskWin32Appable)()
}
