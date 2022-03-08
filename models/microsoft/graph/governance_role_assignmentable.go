package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceRoleAssignmentable 
type GovernanceRoleAssignmentable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignmentState()(*string)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExternalId()(*string)
    GetLinkedEligibleRoleAssignment()(GovernanceRoleAssignmentable)
    GetLinkedEligibleRoleAssignmentId()(*string)
    GetMemberType()(*string)
    GetResource()(GovernanceResourceable)
    GetResourceId()(*string)
    GetRoleDefinition()(GovernanceRoleDefinitionable)
    GetRoleDefinitionId()(*string)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*string)
    GetSubject()(GovernanceSubjectable)
    GetSubjectId()(*string)
    SetAssignmentState(value *string)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExternalId(value *string)()
    SetLinkedEligibleRoleAssignment(value GovernanceRoleAssignmentable)()
    SetLinkedEligibleRoleAssignmentId(value *string)()
    SetMemberType(value *string)()
    SetResource(value GovernanceResourceable)()
    SetResourceId(value *string)()
    SetRoleDefinition(value GovernanceRoleDefinitionable)()
    SetRoleDefinitionId(value *string)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *string)()
    SetSubject(value GovernanceSubjectable)()
    SetSubjectId(value *string)()
}
