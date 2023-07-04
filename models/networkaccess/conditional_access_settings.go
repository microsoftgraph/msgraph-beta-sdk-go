package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ConditionalAccessSettings 
type ConditionalAccessSettings struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewConditionalAccessSettings instantiates a new ConditionalAccessSettings and sets the default values.
func NewConditionalAccessSettings()(*ConditionalAccessSettings) {
    m := &ConditionalAccessSettings{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateConditionalAccessSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["signalingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignalingStatus(val.(*Status))
        }
        return nil
    }
    return res
}
// GetSignalingStatus gets the signalingStatus property value. The signalingStatus property
func (m *ConditionalAccessSettings) GetSignalingStatus()(*Status) {
    val, err := m.GetBackingStore().Get("signalingStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Status)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ConditionalAccessSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSignalingStatus() != nil {
        cast := (*m.GetSignalingStatus()).String()
        err = writer.WriteStringValue("signalingStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSignalingStatus sets the signalingStatus property value. The signalingStatus property
func (m *ConditionalAccessSettings) SetSignalingStatus(value *Status)() {
    err := m.GetBackingStore().Set("signalingStatus", value)
    if err != nil {
        panic(err)
    }
}
// ConditionalAccessSettingsable 
type ConditionalAccessSettingsable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSignalingStatus()(*Status)
    SetSignalingStatus(value *Status)()
}
