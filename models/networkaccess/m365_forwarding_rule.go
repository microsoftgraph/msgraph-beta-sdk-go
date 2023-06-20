package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// M365ForwardingRule 
type M365ForwardingRule struct {
    ForwardingRule
}
// NewM365ForwardingRule instantiates a new M365ForwardingRule and sets the default values.
func NewM365ForwardingRule()(*M365ForwardingRule) {
    m := &M365ForwardingRule{
        ForwardingRule: *NewForwardingRule(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.m365ForwardingRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateM365ForwardingRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateM365ForwardingRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewM365ForwardingRule(), nil
}
// GetCategory gets the category property value. The category property
func (m *M365ForwardingRule) GetCategory()(*ForwardingCategory) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ForwardingCategory)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *M365ForwardingRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ForwardingRule.GetFieldDeserializers()
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseForwardingCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*ForwardingCategory))
        }
        return nil
    }
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
func (m *M365ForwardingRule) GetPorts()([]string) {
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
func (m *M365ForwardingRule) GetProtocol()(*NetworkingProtocol) {
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
func (m *M365ForwardingRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ForwardingRule.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
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
// SetCategory sets the category property value. The category property
func (m *M365ForwardingRule) SetCategory(value *ForwardingCategory)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetPorts sets the ports property value. The ports property
func (m *M365ForwardingRule) SetPorts(value []string)() {
    err := m.GetBackingStore().Set("ports", value)
    if err != nil {
        panic(err)
    }
}
// SetProtocol sets the protocol property value. The protocol property
func (m *M365ForwardingRule) SetProtocol(value *NetworkingProtocol)() {
    err := m.GetBackingStore().Set("protocol", value)
    if err != nil {
        panic(err)
    }
}
// M365ForwardingRuleable 
type M365ForwardingRuleable interface {
    ForwardingRuleable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*ForwardingCategory)
    GetPorts()([]string)
    GetProtocol()(*NetworkingProtocol)
    SetCategory(value *ForwardingCategory)()
    SetPorts(value []string)()
    SetProtocol(value *NetworkingProtocol)()
}
