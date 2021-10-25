package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AppliedConditionalAccessPolicy struct {
    additionalData map[string]interface{};
    conditionsNotSatisfied *ConditionalAccessConditions;
    conditionsSatisfied *ConditionalAccessConditions;
    displayName *string;
    enforcedGrantControls []string;
    enforcedSessionControls []string;
    excludeRulesSatisfied []ConditionalAccessRuleSatisfied;
    id *string;
    includeRulesSatisfied []ConditionalAccessRuleSatisfied;
    result *AppliedConditionalAccessPolicyResult;
}
func NewAppliedConditionalAccessPolicy()(*AppliedConditionalAccessPolicy) {
    m := &AppliedConditionalAccessPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AppliedConditionalAccessPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AppliedConditionalAccessPolicy) GetConditionsNotSatisfied()(*ConditionalAccessConditions) {
    if m == nil {
        return nil
    } else {
        return m.conditionsNotSatisfied
    }
}
func (m *AppliedConditionalAccessPolicy) GetConditionsSatisfied()(*ConditionalAccessConditions) {
    if m == nil {
        return nil
    } else {
        return m.conditionsSatisfied
    }
}
func (m *AppliedConditionalAccessPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *AppliedConditionalAccessPolicy) GetEnforcedGrantControls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedGrantControls
    }
}
func (m *AppliedConditionalAccessPolicy) GetEnforcedSessionControls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedSessionControls
    }
}
func (m *AppliedConditionalAccessPolicy) GetExcludeRulesSatisfied()([]ConditionalAccessRuleSatisfied) {
    if m == nil {
        return nil
    } else {
        return m.excludeRulesSatisfied
    }
}
func (m *AppliedConditionalAccessPolicy) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *AppliedConditionalAccessPolicy) GetIncludeRulesSatisfied()([]ConditionalAccessRuleSatisfied) {
    if m == nil {
        return nil
    } else {
        return m.includeRulesSatisfied
    }
}
func (m *AppliedConditionalAccessPolicy) GetResult()(*AppliedConditionalAccessPolicyResult) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
func (m *AppliedConditionalAccessPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["conditionsNotSatisfied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessConditions)
        if err != nil {
            return err
        }
        cast := val.(ConditionalAccessConditions)
        m.SetConditionsNotSatisfied(&cast)
        return nil
    }
    res["conditionsSatisfied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessConditions)
        if err != nil {
            return err
        }
        cast := val.(ConditionalAccessConditions)
        m.SetConditionsSatisfied(&cast)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["enforcedGrantControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetEnforcedGrantControls(res)
        return nil
    }
    res["enforcedSessionControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetEnforcedSessionControls(res)
        return nil
    }
    res["excludeRulesSatisfied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessRuleSatisfied() })
        if err != nil {
            return err
        }
        res := make([]ConditionalAccessRuleSatisfied, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConditionalAccessRuleSatisfied))
        }
        m.SetExcludeRulesSatisfied(res)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["includeRulesSatisfied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessRuleSatisfied() })
        if err != nil {
            return err
        }
        res := make([]ConditionalAccessRuleSatisfied, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConditionalAccessRuleSatisfied))
        }
        m.SetIncludeRulesSatisfied(res)
        return nil
    }
    res["result"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppliedConditionalAccessPolicyResult)
        if err != nil {
            return err
        }
        cast := val.(AppliedConditionalAccessPolicyResult)
        m.SetResult(&cast)
        return nil
    }
    return res
}
func (m *AppliedConditionalAccessPolicy) IsNil()(bool) {
    return m == nil
}
func (m *AppliedConditionalAccessPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetConditionsNotSatisfied() != nil {
        cast := m.GetConditionsNotSatisfied().String()
        err := writer.WriteStringValue("conditionsNotSatisfied", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionsSatisfied() != nil {
        cast := m.GetConditionsSatisfied().String()
        err := writer.WriteStringValue("conditionsSatisfied", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("enforcedGrantControls", m.GetEnforcedGrantControls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("enforcedSessionControls", m.GetEnforcedSessionControls())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExcludeRulesSatisfied()))
        for i, v := range m.GetExcludeRulesSatisfied() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("excludeRulesSatisfied", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIncludeRulesSatisfied()))
        for i, v := range m.GetIncludeRulesSatisfied() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("includeRulesSatisfied", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResult() != nil {
        cast := m.GetResult().String()
        err := writer.WriteStringValue("result", &cast)
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
func (m *AppliedConditionalAccessPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AppliedConditionalAccessPolicy) SetConditionsNotSatisfied(value *ConditionalAccessConditions)() {
    m.conditionsNotSatisfied = value
}
func (m *AppliedConditionalAccessPolicy) SetConditionsSatisfied(value *ConditionalAccessConditions)() {
    m.conditionsSatisfied = value
}
func (m *AppliedConditionalAccessPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *AppliedConditionalAccessPolicy) SetEnforcedGrantControls(value []string)() {
    m.enforcedGrantControls = value
}
func (m *AppliedConditionalAccessPolicy) SetEnforcedSessionControls(value []string)() {
    m.enforcedSessionControls = value
}
func (m *AppliedConditionalAccessPolicy) SetExcludeRulesSatisfied(value []ConditionalAccessRuleSatisfied)() {
    m.excludeRulesSatisfied = value
}
func (m *AppliedConditionalAccessPolicy) SetId(value *string)() {
    m.id = value
}
func (m *AppliedConditionalAccessPolicy) SetIncludeRulesSatisfied(value []ConditionalAccessRuleSatisfied)() {
    m.includeRulesSatisfied = value
}
func (m *AppliedConditionalAccessPolicy) SetResult(value *AppliedConditionalAccessPolicyResult)() {
    m.result = value
}
