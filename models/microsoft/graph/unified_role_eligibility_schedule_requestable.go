package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleEligibilityScheduleRequestable 
type UnifiedRoleEligibilityScheduleRequestable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Requestable
    GetAction()(*string)
    GetAppScope()(AppScopeable)
    GetAppScopeId()(*string)
    GetDirectoryScope()(DirectoryObjectable)
    GetDirectoryScopeId()(*string)
    GetIsValidationOnly()(*bool)
    GetJustification()(*string)
    GetPrincipal()(DirectoryObjectable)
    GetPrincipalId()(*string)
    GetRoleDefinition()(UnifiedRoleDefinitionable)
    GetRoleDefinitionId()(*string)
    GetScheduleInfo()(RequestScheduleable)
    GetTargetSchedule()(UnifiedRoleEligibilityScheduleable)
    GetTargetScheduleId()(*string)
    GetTicketInfo()(TicketInfoable)
    SetAction(value *string)()
    SetAppScope(value AppScopeable)()
    SetAppScopeId(value *string)()
    SetDirectoryScope(value DirectoryObjectable)()
    SetDirectoryScopeId(value *string)()
    SetIsValidationOnly(value *bool)()
    SetJustification(value *string)()
    SetPrincipal(value DirectoryObjectable)()
    SetPrincipalId(value *string)()
    SetRoleDefinition(value UnifiedRoleDefinitionable)()
    SetRoleDefinitionId(value *string)()
    SetScheduleInfo(value RequestScheduleable)()
    SetTargetSchedule(value UnifiedRoleEligibilityScheduleable)()
    SetTargetScheduleId(value *string)()
    SetTicketInfo(value TicketInfoable)()
}
