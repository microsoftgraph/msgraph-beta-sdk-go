package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TunnelConfigurationIKEv2Default 
type TunnelConfigurationIKEv2Default struct {
    TunnelConfiguration
}
// NewTunnelConfigurationIKEv2Default instantiates a new tunnelConfigurationIKEv2Default and sets the default values.
func NewTunnelConfigurationIKEv2Default()(*TunnelConfigurationIKEv2Default) {
    m := &TunnelConfigurationIKEv2Default{
        TunnelConfiguration: *NewTunnelConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.tunnelConfigurationIKEv2Default"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTunnelConfigurationIKEv2DefaultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTunnelConfigurationIKEv2DefaultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTunnelConfigurationIKEv2Default(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TunnelConfigurationIKEv2Default) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TunnelConfiguration.GetFieldDeserializers()
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TunnelConfigurationIKEv2Default) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TunnelConfigurationIKEv2Default) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TunnelConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TunnelConfigurationIKEv2Default) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// TunnelConfigurationIKEv2Defaultable 
type TunnelConfigurationIKEv2Defaultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TunnelConfigurationable
    GetOdataType()(*string)
    SetOdataType(value *string)()
}
