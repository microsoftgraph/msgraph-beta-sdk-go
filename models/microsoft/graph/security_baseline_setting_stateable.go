package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SecurityBaselineSettingStateable 
type SecurityBaselineSettingStateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContributingPolicies()([]SecurityBaselineContributingPolicyable)
    GetErrorCode()(*string)
    GetSettingCategoryId()(*string)
    GetSettingCategoryName()(*string)
    GetSettingId()(*string)
    GetSettingName()(*string)
    GetSourcePolicies()([]SettingSourceable)
    GetState()(*SecurityBaselineComplianceState)
    SetContributingPolicies(value []SecurityBaselineContributingPolicyable)()
    SetErrorCode(value *string)()
    SetSettingCategoryId(value *string)()
    SetSettingCategoryName(value *string)()
    SetSettingId(value *string)()
    SetSettingName(value *string)()
    SetSourcePolicies(value []SettingSourceable)()
    SetState(value *SecurityBaselineComplianceState)()
}
