package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskLocalGroup the class used to identify a local group for the kiosk configuration
type WindowsKioskLocalGroup struct {
    WindowsKioskUser
}
// NewWindowsKioskLocalGroup instantiates a new windowsKioskLocalGroup and sets the default values.
func NewWindowsKioskLocalGroup()(*WindowsKioskLocalGroup) {
    m := &WindowsKioskLocalGroup{
        WindowsKioskUser: *NewWindowsKioskUser(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskLocalGroup"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsKioskLocalGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskLocalGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskLocalGroup(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskLocalGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskUser.GetFieldDeserializers()
    res["groupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupName(val)
        }
        return nil
    }
    return res
}
// GetGroupName gets the groupName property value. The name of the local group that will be locked to this kiosk configuration
func (m *WindowsKioskLocalGroup) GetGroupName()(*string) {
    val, err := m.GetBackingStore().Get("groupName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsKioskLocalGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskUser.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupName", m.GetGroupName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupName sets the groupName property value. The name of the local group that will be locked to this kiosk configuration
func (m *WindowsKioskLocalGroup) SetGroupName(value *string)() {
    err := m.GetBackingStore().Set("groupName", value)
    if err != nil {
        panic(err)
    }
}
// WindowsKioskLocalGroupable 
type WindowsKioskLocalGroupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsKioskUserable
    GetGroupName()(*string)
    SetGroupName(value *string)()
}
