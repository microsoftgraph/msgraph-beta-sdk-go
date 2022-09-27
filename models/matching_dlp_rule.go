package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MatchingDlpRule 
type MatchingDlpRule struct {
    // The actions property
    actions []DlpActionInfoable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The isMostRestrictive property
    isMostRestrictive *bool
    // The OdataType property
    odataType *string
    // The policyId property
    policyId *string
    // The policyName property
    policyName *string
    // The priority property
    priority *int32
    // The ruleId property
    ruleId *string
    // The ruleMode property
    ruleMode *RuleMode
    // The ruleName property
    ruleName *string
}
// NewMatchingDlpRule instantiates a new matchingDlpRule and sets the default values.
func NewMatchingDlpRule()(*MatchingDlpRule) {
    m := &MatchingDlpRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.matchingDlpRule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMatchingDlpRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMatchingDlpRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMatchingDlpRule(), nil
}
// GetActions gets the actions property value. The actions property
func (m *MatchingDlpRule) GetActions()([]DlpActionInfoable) {
    return m.actions
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MatchingDlpRule) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MatchingDlpRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDlpActionInfoFromDiscriminatorValue , m.SetActions)
    res["isMostRestrictive"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsMostRestrictive)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["policyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPolicyId)
    res["policyName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPolicyName)
    res["priority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPriority)
    res["ruleId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRuleId)
    res["ruleMode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRuleMode , m.SetRuleMode)
    res["ruleName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRuleName)
    return res
}
// GetIsMostRestrictive gets the isMostRestrictive property value. The isMostRestrictive property
func (m *MatchingDlpRule) GetIsMostRestrictive()(*bool) {
    return m.isMostRestrictive
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MatchingDlpRule) GetOdataType()(*string) {
    return m.odataType
}
// GetPolicyId gets the policyId property value. The policyId property
func (m *MatchingDlpRule) GetPolicyId()(*string) {
    return m.policyId
}
// GetPolicyName gets the policyName property value. The policyName property
func (m *MatchingDlpRule) GetPolicyName()(*string) {
    return m.policyName
}
// GetPriority gets the priority property value. The priority property
func (m *MatchingDlpRule) GetPriority()(*int32) {
    return m.priority
}
// GetRuleId gets the ruleId property value. The ruleId property
func (m *MatchingDlpRule) GetRuleId()(*string) {
    return m.ruleId
}
// GetRuleMode gets the ruleMode property value. The ruleMode property
func (m *MatchingDlpRule) GetRuleMode()(*RuleMode) {
    return m.ruleMode
}
// GetRuleName gets the ruleName property value. The ruleName property
func (m *MatchingDlpRule) GetRuleName()(*string) {
    return m.ruleName
}
// Serialize serializes information the current object
func (m *MatchingDlpRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetActions())
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
    m.actions = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MatchingDlpRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsMostRestrictive sets the isMostRestrictive property value. The isMostRestrictive property
func (m *MatchingDlpRule) SetIsMostRestrictive(value *bool)() {
    m.isMostRestrictive = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MatchingDlpRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPolicyId sets the policyId property value. The policyId property
func (m *MatchingDlpRule) SetPolicyId(value *string)() {
    m.policyId = value
}
// SetPolicyName sets the policyName property value. The policyName property
func (m *MatchingDlpRule) SetPolicyName(value *string)() {
    m.policyName = value
}
// SetPriority sets the priority property value. The priority property
func (m *MatchingDlpRule) SetPriority(value *int32)() {
    m.priority = value
}
// SetRuleId sets the ruleId property value. The ruleId property
func (m *MatchingDlpRule) SetRuleId(value *string)() {
    m.ruleId = value
}
// SetRuleMode sets the ruleMode property value. The ruleMode property
func (m *MatchingDlpRule) SetRuleMode(value *RuleMode)() {
    m.ruleMode = value
}
// SetRuleName sets the ruleName property value. The ruleName property
func (m *MatchingDlpRule) SetRuleName(value *string)() {
    m.ruleName = value
}
