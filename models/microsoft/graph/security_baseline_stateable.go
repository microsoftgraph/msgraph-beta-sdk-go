package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SecurityBaselineStateable 
type SecurityBaselineStateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetSecurityBaselineTemplateId()(*string)
    GetSettingStates()([]SecurityBaselineSettingStateable)
    GetState()(*SecurityBaselineComplianceState)
    GetUserPrincipalName()(*string)
    SetDisplayName(value *string)()
    SetSecurityBaselineTemplateId(value *string)()
    SetSettingStates(value []SecurityBaselineSettingStateable)()
    SetState(value *SecurityBaselineComplianceState)()
    SetUserPrincipalName(value *string)()
}
