package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DlpPoliciesJobResult 
type DlpPoliciesJobResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    matchingRules []MatchingDlpRule;
}
// NewDlpPoliciesJobResult instantiates a new dlpPoliciesJobResult and sets the default values.
func NewDlpPoliciesJobResult()(*DlpPoliciesJobResult) {
    m := &DlpPoliciesJobResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpPoliciesJobResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMatchingRules gets the matchingRules property value. 
func (m *DlpPoliciesJobResult) GetMatchingRules()([]MatchingDlpRule) {
    if m == nil {
        return nil
    } else {
        return m.matchingRules
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DlpPoliciesJobResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["matchingRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMatchingDlpRule() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MatchingDlpRule, len(val))
            for i, v := range val {
                res[i] = *(v.(*MatchingDlpRule))
            }
            m.SetMatchingRules(res)
        }
        return nil
    }
    return res
}
func (m *DlpPoliciesJobResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DlpPoliciesJobResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetMatchingRules() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpPoliciesJobResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMatchingRules sets the matchingRules property value. 
func (m *DlpPoliciesJobResult) SetMatchingRules(value []MatchingDlpRule)() {
    if m != nil {
        m.matchingRules = value
    }
}
