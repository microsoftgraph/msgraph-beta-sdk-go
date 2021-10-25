package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DlpPoliciesJobResult struct {
    additionalData map[string]interface{};
    matchingRules []MatchingDlpRule;
}
func NewDlpPoliciesJobResult()(*DlpPoliciesJobResult) {
    m := &DlpPoliciesJobResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DlpPoliciesJobResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DlpPoliciesJobResult) GetMatchingRules()([]MatchingDlpRule) {
    if m == nil {
        return nil
    } else {
        return m.matchingRules
    }
}
func (m *DlpPoliciesJobResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["matchingRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMatchingDlpRule() })
        if err != nil {
            return err
        }
        res := make([]MatchingDlpRule, len(val))
        for i, v := range val {
            res[i] = *(v.(*MatchingDlpRule))
        }
        m.SetMatchingRules(res)
        return nil
    }
    return res
}
func (m *DlpPoliciesJobResult) IsNil()(bool) {
    return m == nil
}
func (m *DlpPoliciesJobResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMatchingRules()))
        for i, v := range m.GetMatchingRules() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("matchingRules", cast)
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
func (m *DlpPoliciesJobResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DlpPoliciesJobResult) SetMatchingRules(value []MatchingDlpRule)() {
    m.matchingRules = value
}
