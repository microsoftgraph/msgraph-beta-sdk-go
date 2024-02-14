package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskAzureADUser the class used to identify an AzureAD user account for the kiosk configuration
type WindowsKioskAzureADUser struct {
    WindowsKioskUser
}
// NewWindowsKioskAzureADUser instantiates a new WindowsKioskAzureADUser and sets the default values.
func NewWindowsKioskAzureADUser()(*WindowsKioskAzureADUser) {
    m := &WindowsKioskAzureADUser{
        WindowsKioskUser: *NewWindowsKioskUser(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskAzureADUser"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsKioskAzureADUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsKioskAzureADUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskAzureADUser(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsKioskAzureADUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskUser.GetFieldDeserializers()
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetUserId gets the userId property value. The ID of the AzureAD user that will be locked to this kiosk configuration
// returns a *string when successful
func (m *WindowsKioskAzureADUser) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The user accounts that will be locked to this kiosk configuration
// returns a *string when successful
func (m *WindowsKioskAzureADUser) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsKioskAzureADUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskUser.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserId sets the userId property value. The ID of the AzureAD user that will be locked to this kiosk configuration
func (m *WindowsKioskAzureADUser) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user accounts that will be locked to this kiosk configuration
func (m *WindowsKioskAzureADUser) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
type WindowsKioskAzureADUserable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsKioskUserable
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}
