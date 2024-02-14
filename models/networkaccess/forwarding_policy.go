package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ForwardingPolicy struct {
    Policy
}
// NewForwardingPolicy instantiates a new ForwardingPolicy and sets the default values.
func NewForwardingPolicy()(*ForwardingPolicy) {
    m := &ForwardingPolicy{
        Policy: *NewPolicy(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.forwardingPolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateForwardingPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateForwardingPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewForwardingPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ForwardingPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Policy.GetFieldDeserializers()
    res["trafficForwardingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrafficForwardingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrafficForwardingType(val.(*TrafficForwardingType))
        }
        return nil
    }
    return res
}
// GetTrafficForwardingType gets the trafficForwardingType property value. The trafficForwardingType property
// returns a *TrafficForwardingType when successful
func (m *ForwardingPolicy) GetTrafficForwardingType()(*TrafficForwardingType) {
    val, err := m.GetBackingStore().Get("trafficForwardingType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TrafficForwardingType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ForwardingPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Policy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetTrafficForwardingType() != nil {
        cast := (*m.GetTrafficForwardingType()).String()
        err = writer.WriteStringValue("trafficForwardingType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTrafficForwardingType sets the trafficForwardingType property value. The trafficForwardingType property
func (m *ForwardingPolicy) SetTrafficForwardingType(value *TrafficForwardingType)() {
    err := m.GetBackingStore().Set("trafficForwardingType", value)
    if err != nil {
        panic(err)
    }
}
type ForwardingPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Policyable
    GetTrafficForwardingType()(*TrafficForwardingType)
    SetTrafficForwardingType(value *TrafficForwardingType)()
}
