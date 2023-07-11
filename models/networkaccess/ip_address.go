package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IpAddress 
type IpAddress struct {
    RuleDestination
    // The OdataType property
    OdataType *string
}
// NewIpAddress instantiates a new ipAddress and sets the default values.
func NewIpAddress()(*IpAddress) {
    m := &IpAddress{
        RuleDestination: *NewRuleDestination(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.ipAddress"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIpAddressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIpAddressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpAddress(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IpAddress) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RuleDestination.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. Defines the IP address used in a destination for a rule.
func (m *IpAddress) GetValue()(*string) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IpAddress) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RuleDestination.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. Defines the IP address used in a destination for a rule.
func (m *IpAddress) SetValue(value *string)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// IpAddressable 
type IpAddressable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RuleDestinationable
    GetValue()(*string)
    SetValue(value *string)()
}
