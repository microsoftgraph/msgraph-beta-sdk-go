package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceResourceable 
type GovernanceResourceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetParent()(GovernanceResourceable)
    GetRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRegisteredRoot()(*string)
    GetRoleAssignmentRequests()([]GovernanceRoleAssignmentRequestable)
    GetRoleAssignments()([]GovernanceRoleAssignmentable)
    GetRoleDefinitions()([]GovernanceRoleDefinitionable)
    GetRoleSettings()([]GovernanceRoleSettingable)
    GetStatus()(*string)
    GetType()(*string)
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetParent(value GovernanceResourceable)()
    SetRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRegisteredRoot(value *string)()
    SetRoleAssignmentRequests(value []GovernanceRoleAssignmentRequestable)()
    SetRoleAssignments(value []GovernanceRoleAssignmentable)()
    SetRoleDefinitions(value []GovernanceRoleDefinitionable)()
    SetRoleSettings(value []GovernanceRoleSettingable)()
    SetStatus(value *string)()
    SetType(value *string)()
}
