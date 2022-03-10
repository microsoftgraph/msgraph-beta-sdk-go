package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceRoleAssignmentRequestable 
type GovernanceRoleAssignmentRequestable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignmentState()(*string)
    GetLinkedEligibleRoleAssignmentId()(*string)
    GetReason()(*string)
    GetRequestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetResource()(GovernanceResourceable)
    GetResourceId()(*string)
    GetRoleDefinition()(GovernanceRoleDefinitionable)
    GetRoleDefinitionId()(*string)
    GetSchedule()(GovernanceScheduleable)
    GetStatus()(GovernanceRoleAssignmentRequestStatusable)
    GetSubject()(GovernanceSubjectable)
    GetSubjectId()(*string)
    GetType()(*string)
    SetAssignmentState(value *string)()
    SetLinkedEligibleRoleAssignmentId(value *string)()
    SetReason(value *string)()
    SetRequestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetResource(value GovernanceResourceable)()
    SetResourceId(value *string)()
    SetRoleDefinition(value GovernanceRoleDefinitionable)()
    SetRoleDefinitionId(value *string)()
    SetSchedule(value GovernanceScheduleable)()
    SetStatus(value GovernanceRoleAssignmentRequestStatusable)()
    SetSubject(value GovernanceSubjectable)()
    SetSubjectId(value *string)()
    SetType(value *string)()
}
