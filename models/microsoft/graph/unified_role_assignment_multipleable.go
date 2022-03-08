package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleAssignmentMultipleable 
type UnifiedRoleAssignmentMultipleable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppScopeIds()([]string)
    GetAppScopes()([]AppScopeable)
    GetCondition()(*string)
    GetDescription()(*string)
    GetDirectoryScopeIds()([]string)
    GetDirectoryScopes()([]DirectoryObjectable)
    GetDisplayName()(*string)
    GetPrincipalIds()([]string)
    GetPrincipals()([]DirectoryObjectable)
    GetRoleDefinition()(UnifiedRoleDefinitionable)
    GetRoleDefinitionId()(*string)
    SetAppScopeIds(value []string)()
    SetAppScopes(value []AppScopeable)()
    SetCondition(value *string)()
    SetDescription(value *string)()
    SetDirectoryScopeIds(value []string)()
    SetDirectoryScopes(value []DirectoryObjectable)()
    SetDisplayName(value *string)()
    SetPrincipalIds(value []string)()
    SetPrincipals(value []DirectoryObjectable)()
    SetRoleDefinition(value UnifiedRoleDefinitionable)()
    SetRoleDefinitionId(value *string)()
}
