package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleManagementPolicyable 
type UnifiedRoleManagementPolicyable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEffectiveRules()([]UnifiedRoleManagementPolicyRuleable)
    GetIsOrganizationDefault()(*bool)
    GetLastModifiedBy()(Identityable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRules()([]UnifiedRoleManagementPolicyRuleable)
    GetScopeId()(*string)
    GetScopeType()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEffectiveRules(value []UnifiedRoleManagementPolicyRuleable)()
    SetIsOrganizationDefault(value *bool)()
    SetLastModifiedBy(value Identityable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRules(value []UnifiedRoleManagementPolicyRuleable)()
    SetScopeId(value *string)()
    SetScopeType(value *string)()
}
