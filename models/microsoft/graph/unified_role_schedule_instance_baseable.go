package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleScheduleInstanceBaseable 
type UnifiedRoleScheduleInstanceBaseable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppScope()(AppScopeable)
    GetAppScopeId()(*string)
    GetDirectoryScope()(DirectoryObjectable)
    GetDirectoryScopeId()(*string)
    GetPrincipal()(DirectoryObjectable)
    GetPrincipalId()(*string)
    GetRoleDefinition()(UnifiedRoleDefinitionable)
    GetRoleDefinitionId()(*string)
    SetAppScope(value AppScopeable)()
    SetAppScopeId(value *string)()
    SetDirectoryScope(value DirectoryObjectable)()
    SetDirectoryScopeId(value *string)()
    SetPrincipal(value DirectoryObjectable)()
    SetPrincipalId(value *string)()
    SetRoleDefinition(value UnifiedRoleDefinitionable)()
    SetRoleDefinitionId(value *string)()
}
