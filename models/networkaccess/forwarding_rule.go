package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ForwardingRule 
type ForwardingRule struct {
    PolicyRule
}
// NewForwardingRule instantiates a new ForwardingRule and sets the default values.
func NewForwardingRule()(*ForwardingRule) {
    m := &ForwardingRule{
        PolicyRule: *NewPolicyRule(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.forwardingRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateForwardingRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateForwardingRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.networkaccess.m365ForwardingRule":
                        return NewM365ForwardingRule(), nil
                    case "#microsoft.graph.networkaccess.privateAccessForwardingRule":
                        return NewPrivateAccessForwardingRule(), nil
                }
            }
        }
    }
    return NewForwardingRule(), nil
}
// GetAction gets the action property value. The action property
func (m *ForwardingRule) GetAction()(*ForwardingRuleAction) {
    val, err := m.GetBackingStore().Get("action")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ForwardingRuleAction)
    }
    return nil
}
// GetDestinations gets the destinations property value. The destinations property
func (m *ForwardingRule) GetDestinations()([]RuleDestinationable) {
    val, err := m.GetBackingStore().Get("destinations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RuleDestinationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ForwardingRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyRule.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseForwardingRuleAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*ForwardingRuleAction))
        }
        return nil
    }
    res["destinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRuleDestinationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RuleDestinationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RuleDestinationable)
                }
            }
            m.SetDestinations(res)
        }
        return nil
    }
    res["ruleType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkDestinationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleType(val.(*NetworkDestinationType))
        }
        return nil
    }
    return res
}
// GetRuleType gets the ruleType property value. The ruleType property
func (m *ForwardingRule) GetRuleType()(*NetworkDestinationType) {
    val, err := m.GetBackingStore().Get("ruleType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NetworkDestinationType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ForwardingRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyRule.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDestinations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDestinations()))
        for i, v := range m.GetDestinations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("destinations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRuleType() != nil {
        cast := (*m.GetRuleType()).String()
        err = writer.WriteStringValue("ruleType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The action property
func (m *ForwardingRule) SetAction(value *ForwardingRuleAction)() {
    err := m.GetBackingStore().Set("action", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinations sets the destinations property value. The destinations property
func (m *ForwardingRule) SetDestinations(value []RuleDestinationable)() {
    err := m.GetBackingStore().Set("destinations", value)
    if err != nil {
        panic(err)
    }
}
// SetRuleType sets the ruleType property value. The ruleType property
func (m *ForwardingRule) SetRuleType(value *NetworkDestinationType)() {
    err := m.GetBackingStore().Set("ruleType", value)
    if err != nil {
        panic(err)
    }
}
// ForwardingRuleable 
type ForwardingRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyRuleable
    GetAction()(*ForwardingRuleAction)
    GetDestinations()([]RuleDestinationable)
    GetRuleType()(*NetworkDestinationType)
    SetAction(value *ForwardingRuleAction)()
    SetDestinations(value []RuleDestinationable)()
    SetRuleType(value *NetworkDestinationType)()
}
