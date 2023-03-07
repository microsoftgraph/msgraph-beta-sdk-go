package tenantadmin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Sharepoint 
type Sharepoint struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewSharepoint instantiates a new Sharepoint and sets the default values.
func NewSharepoint()(*Sharepoint) {
    m := &Sharepoint{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateSharepointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharepointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharepoint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Sharepoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(Settingsable))
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. Represents the tenant-level settings for SharePoint and OneDrive.
func (m *Sharepoint) GetSettings()(Settingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Settingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Sharepoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetSettings sets the settings property value. Represents the tenant-level settings for SharePoint and OneDrive.
func (m *Sharepoint) SetSettings(value Settingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// Sharepointable 
type Sharepointable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSettings()(Settingsable)
    SetSettings(value Settingsable)()
}
