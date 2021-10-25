package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MatchingDlpRule struct {
    actions []DlpActionInfo;
    additionalData map[string]interface{};
    isMostRestrictive *bool;
    policyId *string;
    policyName *string;
    priority *int32;
    ruleId *string;
    ruleMode *RuleMode;
    ruleName *string;
}
func NewMatchingDlpRule()(*MatchingDlpRule) {
    m := &MatchingDlpRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MatchingDlpRule) GetActions()([]DlpActionInfo) {
    if m == nil {
        return nil
    } else {
        return m.actions
    }
}
func (m *MatchingDlpRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MatchingDlpRule) GetIsMostRestrictive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMostRestrictive
    }
}
func (m *MatchingDlpRule) GetPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyId
    }
}
func (m *MatchingDlpRule) GetPolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyName
    }
}
func (m *MatchingDlpRule) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
func (m *MatchingDlpRule) GetRuleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleId
    }
}
func (m *MatchingDlpRule) GetRuleMode()(*RuleMode) {
    if m == nil {
        return nil
    } else {
        return m.ruleMode
    }
}
func (m *MatchingDlpRule) GetRuleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleName
    }
}
func (m *MatchingDlpRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDlpActionInfo() })
        if err != nil {
            return err
        }
        res := make([]DlpActionInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*DlpActionInfo))
        }
        m.SetActions(res)
        return nil
    }
    res["isMostRestrictive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMostRestrictive(val)
        return nil
    }
    res["policyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyId(val)
        return nil
    }
    res["policyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyName(val)
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPriority(val)
        return nil
    }
    res["ruleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRuleId(val)
        return nil
    }
    res["ruleMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRuleMode)
        if err != nil {
            return err
        }
        cast := val.(RuleMode)
        m.SetRuleMode(&cast)
        return nil
    }
    res["ruleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRuleName(val)
        return nil
    }
    return res
}
func (m *MatchingDlpRule) IsNil()(bool) {
    return m == nil
}
func (m *MatchingDlpRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
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
        cast := m.GetRuleMode().String()
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
func (m *MatchingDlpRule) SetActions(value []DlpActionInfo)() {
    m.actions = value
}
func (m *MatchingDlpRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MatchingDlpRule) SetIsMostRestrictive(value *bool)() {
    m.isMostRestrictive = value
}
func (m *MatchingDlpRule) SetPolicyId(value *string)() {
    m.policyId = value
}
func (m *MatchingDlpRule) SetPolicyName(value *string)() {
    m.policyName = value
}
func (m *MatchingDlpRule) SetPriority(value *int32)() {
    m.priority = value
}
func (m *MatchingDlpRule) SetRuleId(value *string)() {
    m.ruleId = value
}
func (m *MatchingDlpRule) SetRuleMode(value *RuleMode)() {
    m.ruleMode = value
}
func (m *MatchingDlpRule) SetRuleName(value *string)() {
    m.ruleName = value
}
