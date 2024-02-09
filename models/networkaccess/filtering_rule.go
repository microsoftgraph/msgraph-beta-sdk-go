package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FilteringRule struct {
    PolicyRule
}
// NewFilteringRule instantiates a new FilteringRule and sets the default values.
func NewFilteringRule()(*FilteringRule) {
    m := &FilteringRule{
        PolicyRule: *NewPolicyRule(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.filteringRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFilteringRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilteringRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.networkaccess.fqdnFilteringRule":
                        return NewFqdnFilteringRule(), nil
                    case "#microsoft.graph.networkaccess.webCategoryFilteringRule":
                        return NewWebCategoryFilteringRule(), nil
                }
            }
        }
    }
    return NewFilteringRule(), nil
}
// GetDestinations gets the destinations property value. Possible destinations and types of destinations accessed by the user in accordance with the network filtering policy, such as IP addresses and FQDNs/URLs.
// returns a []RuleDestinationable when successful
func (m *FilteringRule) GetDestinations()([]RuleDestinationable) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FilteringRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyRule.GetFieldDeserializers()
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
// returns a *NetworkDestinationType when successful
func (m *FilteringRule) GetRuleType()(*NetworkDestinationType) {
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
func (m *FilteringRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyRule.Serialize(writer)
    if err != nil {
        return err
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
// SetDestinations sets the destinations property value. Possible destinations and types of destinations accessed by the user in accordance with the network filtering policy, such as IP addresses and FQDNs/URLs.
func (m *FilteringRule) SetDestinations(value []RuleDestinationable)() {
    err := m.GetBackingStore().Set("destinations", value)
    if err != nil {
        panic(err)
    }
}
// SetRuleType sets the ruleType property value. The ruleType property
func (m *FilteringRule) SetRuleType(value *NetworkDestinationType)() {
    err := m.GetBackingStore().Set("ruleType", value)
    if err != nil {
        panic(err)
    }
}
type FilteringRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyRuleable
    GetDestinations()([]RuleDestinationable)
    GetRuleType()(*NetworkDestinationType)
    SetDestinations(value []RuleDestinationable)()
    SetRuleType(value *NetworkDestinationType)()
}
