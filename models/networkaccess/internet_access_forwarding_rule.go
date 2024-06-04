package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InternetAccessForwardingRule struct {
    ForwardingRule
}
// NewInternetAccessForwardingRule instantiates a new InternetAccessForwardingRule and sets the default values.
func NewInternetAccessForwardingRule()(*InternetAccessForwardingRule) {
    m := &InternetAccessForwardingRule{
        ForwardingRule: *NewForwardingRule(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.internetAccessForwardingRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInternetAccessForwardingRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInternetAccessForwardingRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInternetAccessForwardingRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InternetAccessForwardingRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ForwardingRule.GetFieldDeserializers()
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
    res["protocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkingProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocol(val.(*NetworkingProtocol))
        }
        return nil
    }
    return res
}
// GetPorts gets the ports property value. The ports property
// returns a []string when successful
func (m *InternetAccessForwardingRule) GetPorts()([]string) {
    val, err := m.GetBackingStore().Get("ports")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetProtocol gets the protocol property value. The protocol property
// returns a *NetworkingProtocol when successful
func (m *InternetAccessForwardingRule) GetProtocol()(*NetworkingProtocol) {
    val, err := m.GetBackingStore().Get("protocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NetworkingProtocol)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InternetAccessForwardingRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ForwardingRule.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPorts() != nil {
        err = writer.WriteCollectionOfStringValues("ports", m.GetPorts())
        if err != nil {
            return err
        }
    }
    if m.GetProtocol() != nil {
        cast := (*m.GetProtocol()).String()
        err = writer.WriteStringValue("protocol", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPorts sets the ports property value. The ports property
func (m *InternetAccessForwardingRule) SetPorts(value []string)() {
    err := m.GetBackingStore().Set("ports", value)
    if err != nil {
        panic(err)
    }
}
// SetProtocol sets the protocol property value. The protocol property
func (m *InternetAccessForwardingRule) SetProtocol(value *NetworkingProtocol)() {
    err := m.GetBackingStore().Set("protocol", value)
    if err != nil {
        panic(err)
    }
}
type InternetAccessForwardingRuleable interface {
    ForwardingRuleable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPorts()([]string)
    GetProtocol()(*NetworkingProtocol)
    SetPorts(value []string)()
    SetProtocol(value *NetworkingProtocol)()
}
