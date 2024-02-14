package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhoneXAP contains properties and inherited properties for Windows Phone XAP Line Of Business apps.
type WindowsPhoneXAP struct {
    MobileLobApp
}
// NewWindowsPhoneXAP instantiates a new WindowsPhoneXAP and sets the default values.
func NewWindowsPhoneXAP()(*WindowsPhoneXAP) {
    m := &WindowsPhoneXAP{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhoneXAP"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsPhoneXAPFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsPhoneXAPFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPhoneXAP(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsPhoneXAP) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["identityVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityVersion(val)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(WindowsMinimumOperatingSystemable))
        }
        return nil
    }
    res["productIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductIdentifier(val)
        }
        return nil
    }
    return res
}
// GetIdentityVersion gets the identityVersion property value. The identity version.
// returns a *string when successful
func (m *WindowsPhoneXAP) GetIdentityVersion()(*string) {
    val, err := m.GetBackingStore().Get("identityVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
// returns a WindowsMinimumOperatingSystemable when successful
func (m *WindowsPhoneXAP) GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable) {
    val, err := m.GetBackingStore().Get("minimumSupportedOperatingSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsMinimumOperatingSystemable)
    }
    return nil
}
// GetProductIdentifier gets the productIdentifier property value. The Product Identifier.
// returns a *string when successful
func (m *WindowsPhoneXAP) GetProductIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("productIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsPhoneXAP) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("identityVersion", m.GetIdentityVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productIdentifier", m.GetProductIdentifier())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentityVersion sets the identityVersion property value. The identity version.
func (m *WindowsPhoneXAP) SetIdentityVersion(value *string)() {
    err := m.GetBackingStore().Set("identityVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
func (m *WindowsPhoneXAP) SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)() {
    err := m.GetBackingStore().Set("minimumSupportedOperatingSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetProductIdentifier sets the productIdentifier property value. The Product Identifier.
func (m *WindowsPhoneXAP) SetProductIdentifier(value *string)() {
    err := m.GetBackingStore().Set("productIdentifier", value)
    if err != nil {
        panic(err)
    }
}
type WindowsPhoneXAPable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentityVersion()(*string)
    GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable)
    GetProductIdentifier()(*string)
    SetIdentityVersion(value *string)()
    SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)()
    SetProductIdentifier(value *string)()
}
