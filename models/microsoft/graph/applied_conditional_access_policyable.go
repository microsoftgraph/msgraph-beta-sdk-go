package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppliedConditionalAccessPolicyable 
type AppliedConditionalAccessPolicyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAuthenticationStrength()(AuthenticationStrengthable)
    GetConditionsNotSatisfied()(*ConditionalAccessConditions)
    GetConditionsSatisfied()(*ConditionalAccessConditions)
    GetDisplayName()(*string)
    GetEnforcedGrantControls()([]string)
    GetEnforcedSessionControls()([]string)
    GetExcludeRulesSatisfied()([]ConditionalAccessRuleSatisfiedable)
    GetId()(*string)
    GetIncludeRulesSatisfied()([]ConditionalAccessRuleSatisfiedable)
    GetResult()(*AppliedConditionalAccessPolicyResult)
    SetAuthenticationStrength(value AuthenticationStrengthable)()
    SetConditionsNotSatisfied(value *ConditionalAccessConditions)()
    SetConditionsSatisfied(value *ConditionalAccessConditions)()
    SetDisplayName(value *string)()
    SetEnforcedGrantControls(value []string)()
    SetEnforcedSessionControls(value []string)()
    SetExcludeRulesSatisfied(value []ConditionalAccessRuleSatisfiedable)()
    SetId(value *string)()
    SetIncludeRulesSatisfied(value []ConditionalAccessRuleSatisfiedable)()
    SetResult(value *AppliedConditionalAccessPolicyResult)()
}
