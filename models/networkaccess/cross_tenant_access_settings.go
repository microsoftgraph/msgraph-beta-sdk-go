package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CrossTenantAccessSettings 
type CrossTenantAccessSettings struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewCrossTenantAccessSettings instantiates a new CrossTenantAccessSettings and sets the default values.
func NewCrossTenantAccessSettings()(*CrossTenantAccessSettings) {
    m := &CrossTenantAccessSettings{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateCrossTenantAccessSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantAccessSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["networkPacketTaggingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkPacketTaggingStatus(val.(*Status))
        }
        return nil
    }
    return res
}
// GetNetworkPacketTaggingStatus gets the networkPacketTaggingStatus property value. The networkPacketTaggingStatus property
func (m *CrossTenantAccessSettings) GetNetworkPacketTaggingStatus()(*Status) {
    val, err := m.GetBackingStore().Get("networkPacketTaggingStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Status)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CrossTenantAccessSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetNetworkPacketTaggingStatus() != nil {
        cast := (*m.GetNetworkPacketTaggingStatus()).String()
        err = writer.WriteStringValue("networkPacketTaggingStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetNetworkPacketTaggingStatus sets the networkPacketTaggingStatus property value. The networkPacketTaggingStatus property
func (m *CrossTenantAccessSettings) SetNetworkPacketTaggingStatus(value *Status)() {
    err := m.GetBackingStore().Set("networkPacketTaggingStatus", value)
    if err != nil {
        panic(err)
    }
}
// CrossTenantAccessSettingsable 
type CrossTenantAccessSettingsable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNetworkPacketTaggingStatus()(*Status)
    SetNetworkPacketTaggingStatus(value *Status)()
}
