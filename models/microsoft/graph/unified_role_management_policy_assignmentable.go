package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleManagementPolicyAssignmentable 
type UnifiedRoleManagementPolicyAssignmentable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetPolicy()(UnifiedRoleManagementPolicyable)
    GetPolicyId()(*string)
    GetRoleDefinitionId()(*string)
    GetScopeId()(*string)
    GetScopeType()(*string)
    SetPolicy(value UnifiedRoleManagementPolicyable)()
    SetPolicyId(value *string)()
    SetRoleDefinitionId(value *string)()
    SetScopeId(value *string)()
    SetScopeType(value *string)()
}
