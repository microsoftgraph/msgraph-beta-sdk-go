package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnumeratedInboundPorts struct {
    InboundPorts
}
// NewEnumeratedInboundPorts instantiates a new EnumeratedInboundPorts and sets the default values.
func NewEnumeratedInboundPorts()(*EnumeratedInboundPorts) {
    m := &EnumeratedInboundPorts{
        InboundPorts: *NewInboundPorts(),
    }
    odataTypeValue := "#microsoft.graph.enumeratedInboundPorts"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEnumeratedInboundPortsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnumeratedInboundPortsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnumeratedInboundPorts(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnumeratedInboundPorts) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InboundPorts.GetFieldDeserializers()
    res["ports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPorts(res)
        }
        return nil
    }
    return res
}
// GetPorts gets the ports property value. Collection of ports that allow inbound traffic.
// returns a []string when successful
func (m *EnumeratedInboundPorts) GetPorts()([]string) {
    val, err := m.GetBackingStore().Get("ports")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EnumeratedInboundPorts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InboundPorts.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPorts() != nil {
        err = writer.WriteCollectionOfStringValues("ports", m.GetPorts())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPorts sets the ports property value. Collection of ports that allow inbound traffic.
func (m *EnumeratedInboundPorts) SetPorts(value []string)() {
    err := m.GetBackingStore().Set("ports", value)
    if err != nil {
        panic(err)
    }
}
type EnumeratedInboundPortsable interface {
    InboundPortsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPorts()([]string)
    SetPorts(value []string)()
}
