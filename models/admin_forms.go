package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminForms 
type AdminForms struct {
    Entity
}
// NewAdminForms instantiates a new adminForms and sets the default values.
func NewAdminForms()(*AdminForms) {
    m := &AdminForms{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAdminFormsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminFormsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminForms(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminForms) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFormsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(FormsSettingsable))
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. The settings property
func (m *AdminForms) GetSettings()(FormsSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(FormsSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AdminForms) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettings sets the settings property value. The settings property
func (m *AdminForms) SetSettings(value FormsSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// AdminFormsable 
type AdminFormsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSettings()(FormsSettingsable)
    SetSettings(value FormsSettingsable)()
}
