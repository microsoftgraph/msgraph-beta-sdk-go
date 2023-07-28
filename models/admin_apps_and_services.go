package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminAppsAndServices 
type AdminAppsAndServices struct {
    Entity
}
// NewAdminAppsAndServices instantiates a new adminAppsAndServices and sets the default values.
func NewAdminAppsAndServices()(*AdminAppsAndServices) {
    m := &AdminAppsAndServices{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAdminAppsAndServicesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminAppsAndServicesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminAppsAndServices(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminAppsAndServices) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppsAndServicesSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(AppsAndServicesSettingsable))
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. The settings property
func (m *AdminAppsAndServices) GetSettings()(AppsAndServicesSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AppsAndServicesSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AdminAppsAndServices) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AdminAppsAndServices) SetSettings(value AppsAndServicesSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// AdminAppsAndServicesable 
type AdminAppsAndServicesable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSettings()(AppsAndServicesSettingsable)
    SetSettings(value AppsAndServicesSettingsable)()
}
