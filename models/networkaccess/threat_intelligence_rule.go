// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ThreatIntelligenceRule struct {
    PolicyRule
}
// NewThreatIntelligenceRule instantiates a new ThreatIntelligenceRule and sets the default values.
func NewThreatIntelligenceRule()(*ThreatIntelligenceRule) {
    m := &ThreatIntelligenceRule{
        PolicyRule: *NewPolicyRule(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.threatIntelligenceRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateThreatIntelligenceRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateThreatIntelligenceRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThreatIntelligenceRule(), nil
}
// GetAction gets the action property value. The action property
// returns a *ThreatIntelligenceAction when successful
func (m *ThreatIntelligenceRule) GetAction()(*ThreatIntelligenceAction) {
    val, err := m.GetBackingStore().Get("action")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ThreatIntelligenceAction)
    }
    return nil
}
// GetDescription gets the description property value. A description of the threat intelligence rule. Supports $filter (eq).
// returns a *string when successful
func (m *ThreatIntelligenceRule) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ThreatIntelligenceRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyRule.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatIntelligenceAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*ThreatIntelligenceAction))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["matchingConditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateThreatIntelligenceMatchingConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchingConditions(val.(ThreatIntelligenceMatchingConditionsable))
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateThreatIntelligenceRuleSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(ThreatIntelligenceRuleSettingsable))
        }
        return nil
    }
    return res
}
// GetMatchingConditions gets the matchingConditions property value. The matchingConditions property
// returns a ThreatIntelligenceMatchingConditionsable when successful
func (m *ThreatIntelligenceRule) GetMatchingConditions()(ThreatIntelligenceMatchingConditionsable) {
    val, err := m.GetBackingStore().Get("matchingConditions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ThreatIntelligenceMatchingConditionsable)
    }
    return nil
}
// GetPriority gets the priority property value. The priority of the rule which determines the order of rule evaluation. Lower values indicate higher priority. Supports $filter (eq).
// returns a *int64 when successful
func (m *ThreatIntelligenceRule) GetPriority()(*int64) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSettings gets the settings property value. The settings property
// returns a ThreatIntelligenceRuleSettingsable when successful
func (m *ThreatIntelligenceRule) GetSettings()(ThreatIntelligenceRuleSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ThreatIntelligenceRuleSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ThreatIntelligenceRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("matchingConditions", m.GetMatchingConditions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The action property
func (m *ThreatIntelligenceRule) SetAction(value *ThreatIntelligenceAction)() {
    err := m.GetBackingStore().Set("action", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. A description of the threat intelligence rule. Supports $filter (eq).
func (m *ThreatIntelligenceRule) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetMatchingConditions sets the matchingConditions property value. The matchingConditions property
func (m *ThreatIntelligenceRule) SetMatchingConditions(value ThreatIntelligenceMatchingConditionsable)() {
    err := m.GetBackingStore().Set("matchingConditions", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. The priority of the rule which determines the order of rule evaluation. Lower values indicate higher priority. Supports $filter (eq).
func (m *ThreatIntelligenceRule) SetPriority(value *int64)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. The settings property
func (m *ThreatIntelligenceRule) SetSettings(value ThreatIntelligenceRuleSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
type ThreatIntelligenceRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyRuleable
    GetAction()(*ThreatIntelligenceAction)
    GetDescription()(*string)
    GetMatchingConditions()(ThreatIntelligenceMatchingConditionsable)
    GetPriority()(*int64)
    GetSettings()(ThreatIntelligenceRuleSettingsable)
    SetAction(value *ThreatIntelligenceAction)()
    SetDescription(value *string)()
    SetMatchingConditions(value ThreatIntelligenceMatchingConditionsable)()
    SetPriority(value *int64)()
    SetSettings(value ThreatIntelligenceRuleSettingsable)()
}
