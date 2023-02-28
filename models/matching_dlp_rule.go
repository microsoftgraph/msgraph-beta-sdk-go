package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// MatchingDlpRule 
type MatchingDlpRule struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMatchingDlpRule instantiates a new matchingDlpRule and sets the default values.
func NewMatchingDlpRule()(*MatchingDlpRule) {
    m := &MatchingDlpRule{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMatchingDlpRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMatchingDlpRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMatchingDlpRule(), nil
}
// GetActions gets the actions property value. The actions property
func (m *MatchingDlpRule) GetActions()([]DlpActionInfoable) {
    val, err := m.GetBackingStore().Get("actions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DlpActionInfoable)
    }
    return nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MatchingDlpRule) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *MatchingDlpRule) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MatchingDlpRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDlpActionInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DlpActionInfoable, len(val))
            for i, v := range val {
                res[i] = v.(DlpActionInfoable)
            }
            m.SetActions(res)
        }
        return nil
    }
    res["isMostRestrictive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMostRestrictive(val)
        }
        return nil
    }
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
    res["policyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
        }
        return nil
    }
    res["policyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyName(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["ruleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleId(val)
        }
        return nil
    }
    res["ruleMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRuleMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleMode(val.(*RuleMode))
        }
        return nil
    }
    res["ruleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleName(val)
        }
        return nil
    }
    return res
}
// GetIsMostRestrictive gets the isMostRestrictive property value. The isMostRestrictive property
func (m *MatchingDlpRule) GetIsMostRestrictive()(*bool) {
    val, err := m.GetBackingStore().Get("isMostRestrictive")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MatchingDlpRule) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyId gets the policyId property value. The policyId property
func (m *MatchingDlpRule) GetPolicyId()(*string) {
    val, err := m.GetBackingStore().Get("policyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyName gets the policyName property value. The policyName property
func (m *MatchingDlpRule) GetPolicyName()(*string) {
    val, err := m.GetBackingStore().Get("policyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPriority gets the priority property value. The priority property
func (m *MatchingDlpRule) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRuleId gets the ruleId property value. The ruleId property
func (m *MatchingDlpRule) GetRuleId()(*string) {
    val, err := m.GetBackingStore().Get("ruleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRuleMode gets the ruleMode property value. The ruleMode property
func (m *MatchingDlpRule) GetRuleMode()(*RuleMode) {
    val, err := m.GetBackingStore().Get("ruleMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RuleMode)
    }
    return nil
}
// GetRuleName gets the ruleName property value. The ruleName property
func (m *MatchingDlpRule) GetRuleName()(*string) {
    val, err := m.GetBackingStore().Get("ruleName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MatchingDlpRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActions()))
        for i, v := range m.GetActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("actions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMostRestrictive", m.GetIsMostRestrictive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyId", m.GetPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyName", m.GetPolicyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleId", m.GetRuleId())
        if err != nil {
            return err
        }
    }
    if m.GetRuleMode() != nil {
        cast := (*m.GetRuleMode()).String()
        err := writer.WriteStringValue("ruleMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleName", m.GetRuleName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. The actions property
func (m *MatchingDlpRule) SetActions(value []DlpActionInfoable)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MatchingDlpRule) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *MatchingDlpRule) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsMostRestrictive sets the isMostRestrictive property value. The isMostRestrictive property
func (m *MatchingDlpRule) SetIsMostRestrictive(value *bool)() {
    err := m.GetBackingStore().Set("isMostRestrictive", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MatchingDlpRule) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyId sets the policyId property value. The policyId property
func (m *MatchingDlpRule) SetPolicyId(value *string)() {
    err := m.GetBackingStore().Set("policyId", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyName sets the policyName property value. The policyName property
func (m *MatchingDlpRule) SetPolicyName(value *string)() {
    err := m.GetBackingStore().Set("policyName", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. The priority property
func (m *MatchingDlpRule) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetRuleId sets the ruleId property value. The ruleId property
func (m *MatchingDlpRule) SetRuleId(value *string)() {
    err := m.GetBackingStore().Set("ruleId", value)
    if err != nil {
        panic(err)
    }
}
// SetRuleMode sets the ruleMode property value. The ruleMode property
func (m *MatchingDlpRule) SetRuleMode(value *RuleMode)() {
    err := m.GetBackingStore().Set("ruleMode", value)
    if err != nil {
        panic(err)
    }
}
// SetRuleName sets the ruleName property value. The ruleName property
func (m *MatchingDlpRule) SetRuleName(value *string)() {
    err := m.GetBackingStore().Set("ruleName", value)
    if err != nil {
        panic(err)
    }
}
// MatchingDlpRuleable 
type MatchingDlpRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]DlpActionInfoable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsMostRestrictive()(*bool)
    GetOdataType()(*string)
    GetPolicyId()(*string)
    GetPolicyName()(*string)
    GetPriority()(*int32)
    GetRuleId()(*string)
    GetRuleMode()(*RuleMode)
    GetRuleName()(*string)
    SetActions(value []DlpActionInfoable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsMostRestrictive(value *bool)()
    SetOdataType(value *string)()
    SetPolicyId(value *string)()
    SetPolicyName(value *string)()
    SetPriority(value *int32)()
    SetRuleId(value *string)()
    SetRuleMode(value *RuleMode)()
    SetRuleName(value *string)()
}
