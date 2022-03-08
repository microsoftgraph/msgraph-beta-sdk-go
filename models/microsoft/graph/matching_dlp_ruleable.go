package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MatchingDlpRuleable 
type MatchingDlpRuleable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActions()([]DlpActionInfoable)
    GetIsMostRestrictive()(*bool)
    GetPolicyId()(*string)
    GetPolicyName()(*string)
    GetPriority()(*int32)
    GetRuleId()(*string)
    GetRuleMode()(*RuleMode)
    GetRuleName()(*string)
    SetActions(value []DlpActionInfoable)()
    SetIsMostRestrictive(value *bool)()
    SetPolicyId(value *string)()
    SetPolicyName(value *string)()
    SetPriority(value *int32)()
    SetRuleId(value *string)()
    SetRuleMode(value *RuleMode)()
    SetRuleName(value *string)()
}
