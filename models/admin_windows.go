package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminWindows 
type AdminWindows struct {
    Entity
    // The updates property
    updates AdminWindowsUpdatesable
}
// NewAdminWindows instantiates a new AdminWindows and sets the default values.
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
// GetUpdates gets the updates property value. The updates property
func (m *AdminWindows) GetUpdates()(AdminWindowsUpdatesable) {
    return m.updates
}
// Serialize serializes information the current object
func (m *AdminWindows) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("updates", m.GetUpdates())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUpdates sets the updates property value. The updates property
func (m *AdminWindows) SetUpdates(value AdminWindowsUpdatesable)() {
    m.updates = value
}
