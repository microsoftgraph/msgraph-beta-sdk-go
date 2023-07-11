package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminWindows 
type AdminWindows struct {
    Entity
}
// NewAdminWindows instantiates a new adminWindows and sets the default values.
func NewAdminWindows()(*AdminWindows) {
    m := &AdminWindows{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAdminWindowsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminWindowsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminWindows(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminWindows) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["updates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdminWindowsUpdatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdates(val.(AdminWindowsUpdatesable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AdminWindows) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUpdates gets the updates property value. Entity that acts as a container for all Windows Update for Business deployment service functionalities. Read-only.
func (m *AdminWindows) GetUpdates()(AdminWindowsUpdatesable) {
    val, err := m.GetBackingStore().Get("updates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AdminWindowsUpdatesable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AdminWindows) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
        err = writer.WriteObjectValue("updates", m.GetUpdates())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AdminWindows) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdates sets the updates property value. Entity that acts as a container for all Windows Update for Business deployment service functionalities. Read-only.
func (m *AdminWindows) SetUpdates(value AdminWindowsUpdatesable)() {
    err := m.GetBackingStore().Set("updates", value)
    if err != nil {
        panic(err)
    }
}
// AdminWindowsable 
type AdminWindowsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetUpdates()(AdminWindowsUpdatesable)
    SetOdataType(value *string)()
    SetUpdates(value AdminWindowsUpdatesable)()
}
