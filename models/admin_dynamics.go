package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminDynamics 
type AdminDynamics struct {
    Entity
}
// NewAdminDynamics instantiates a new adminDynamics and sets the default values.
func NewAdminDynamics()(*AdminDynamics) {
    m := &AdminDynamics{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAdminDynamicsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminDynamicsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminDynamics(), nil
}
// GetCustomerVoice gets the customerVoice property value. The customerVoice property
func (m *AdminDynamics) GetCustomerVoice()(CustomerVoiceSettingsable) {
    val, err := m.GetBackingStore().Get("customerVoice")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CustomerVoiceSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminDynamics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customerVoice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomerVoiceSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerVoice(val.(CustomerVoiceSettingsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AdminDynamics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("customerVoice", m.GetCustomerVoice())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomerVoice sets the customerVoice property value. The customerVoice property
func (m *AdminDynamics) SetCustomerVoice(value CustomerVoiceSettingsable)() {
    err := m.GetBackingStore().Set("customerVoice", value)
    if err != nil {
        panic(err)
    }
}
// AdminDynamicsable 
type AdminDynamicsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomerVoice()(CustomerVoiceSettingsable)
    SetCustomerVoice(value CustomerVoiceSettingsable)()
}
