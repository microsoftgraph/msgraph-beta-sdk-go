package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TunnelConfigurationIKEv2Default struct {
    TunnelConfiguration
}
// NewTunnelConfigurationIKEv2Default instantiates a new TunnelConfigurationIKEv2Default and sets the default values.
func NewTunnelConfigurationIKEv2Default()(*TunnelConfigurationIKEv2Default) {
    m := &TunnelConfigurationIKEv2Default{
        TunnelConfiguration: *NewTunnelConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.tunnelConfigurationIKEv2Default"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTunnelConfigurationIKEv2DefaultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTunnelConfigurationIKEv2DefaultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTunnelConfigurationIKEv2Default(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TunnelConfigurationIKEv2Default) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TunnelConfiguration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TunnelConfigurationIKEv2Default) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TunnelConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type TunnelConfigurationIKEv2Defaultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TunnelConfigurationable
}
