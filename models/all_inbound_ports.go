package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AllInboundPorts struct {
    InboundPorts
}
// NewAllInboundPorts instantiates a new AllInboundPorts and sets the default values.
func NewAllInboundPorts()(*AllInboundPorts) {
    m := &AllInboundPorts{
        InboundPorts: *NewInboundPorts(),
    }
    odataTypeValue := "#microsoft.graph.allInboundPorts"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAllInboundPortsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAllInboundPortsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllInboundPorts(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AllInboundPorts) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InboundPorts.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AllInboundPorts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InboundPorts.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type AllInboundPortsable interface {
    InboundPortsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
