package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MatchingDlpRule 
type MatchingDlpRule struct {
    // 
    actions []DlpActionInfo;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    isMostRestrictive *bool;
    // 
    policyId *string;
    // 
    policyName *string;
    // 
    priority *int32;
    // 
    ruleId *string;
    // 
    ruleMode *RuleMode;
    // 
    ruleName *string;
}
// NewMatchingDlpRule instantiates a new matchingDlpRule and sets the default values.
func NewMatchingDlpRule()(*MatchingDlpRule) {
    m := &MatchingDlpRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetActions gets the actions property value. 
func (m *MatchingDlpRule) GetActions()([]DlpActionInfo) {
    if m == nil {
        return nil
    } else {
        return m.actions
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MatchingDlpRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsMostRestrictive gets the isMostRestrictive property value. 
func (m *MatchingDlpRule) GetIsMostRestrictive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMostRestrictive
    }
}
// GetPolicyId gets the policyId property value. 
func (m *MatchingDlpRule) GetPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyId
    }
}
// GetPolicyName gets the policyName property value. 
func (m *MatchingDlpRule) GetPolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyName
    }
}
// GetPriority gets the priority property value. 
func (m *MatchingDlpRule) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// GetRuleId gets the ruleId property value. 
func (m *MatchingDlpRule) GetRuleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleId
    }
}
// GetRuleMode gets the ruleMode property value. 
func (m *MatchingDlpRule) GetRuleMode()(*RuleMode) {
    if m == nil {
        return nil
    } else {
        return m.ruleMode
    }
}
// GetRuleName gets the ruleName property value. 
func (m *MatchingDlpRule) GetRuleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MatchingDlpRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDlpActionInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DlpActionInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*DlpActionInfo))
            }
            m.SetActions(res)
        }
        return nil
    }
    res["isMostRestrictive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMostRestrictive(val)
        }
        return nil
    }
    res["policyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
        }
        return nil
    }
    res["policyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyName(val)
        }
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["ruleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleId(val)
        }
        return nil
    }
    res["ruleMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRuleMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleMode(val.(*RuleMode))
        }
        return nil
    }
    res["ruleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *MatchingDlpRule) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MatchingDlpRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetActions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetActions()))
        for i, v := range m.GetActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// SetActions sets the actions property value. 
func (m *MatchingDlpRule) SetActions(value []DlpActionInfo)() {
    if m != nil {
        m.actions = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MatchingDlpRule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsMostRestrictive sets the isMostRestrictive property value. 
func (m *MatchingDlpRule) SetIsMostRestrictive(value *bool)() {
    if m != nil {
        m.isMostRestrictive = value
    }
}
// SetPolicyId sets the policyId property value. 
func (m *MatchingDlpRule) SetPolicyId(value *string)() {
    if m != nil {
        m.policyId = value
    }
}
// SetPolicyName sets the policyName property value. 
func (m *MatchingDlpRule) SetPolicyName(value *string)() {
    if m != nil {
        m.policyName = value
    }
}
// SetPriority sets the priority property value. 
func (m *MatchingDlpRule) SetPriority(value *int32)() {
    if m != nil {
        m.priority = value
    }
}
// SetRuleId sets the ruleId property value. 
func (m *MatchingDlpRule) SetRuleId(value *string)() {
    if m != nil {
        m.ruleId = value
    }
}
// SetRuleMode sets the ruleMode property value. 
func (m *MatchingDlpRule) SetRuleMode(value *RuleMode)() {
    if m != nil {
        m.ruleMode = value
    }
}
// SetRuleName sets the ruleName property value. 
func (m *MatchingDlpRule) SetRuleName(value *string)() {
    if m != nil {
        m.ruleName = value
    }
}
