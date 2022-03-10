package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleDefinitionable 
type RoleDefinitionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsBuiltIn()(*bool)
    GetIsBuiltInRoleDefinition()(*bool)
    GetPermissions()([]RolePermissionable)
    GetRoleAssignments()([]RoleAssignmentable)
    GetRolePermissions()([]RolePermissionable)
    GetRoleScopeTagIds()([]string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsBuiltIn(value *bool)()
    SetIsBuiltInRoleDefinition(value *bool)()
    SetPermissions(value []RolePermissionable)()
    SetRoleAssignments(value []RoleAssignmentable)()
    SetRolePermissions(value []RolePermissionable)()
    SetRoleScopeTagIds(value []string)()
}
