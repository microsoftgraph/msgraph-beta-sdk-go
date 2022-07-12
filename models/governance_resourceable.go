package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceResourceable 
type GovernanceResourceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
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
}
