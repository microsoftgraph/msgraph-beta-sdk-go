package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivilegedRoleAssignmentRequestable 
type PrivilegedRoleAssignmentRequestable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignmentState()(*string)
    GetDuration()(*string)
    GetReason()(*string)
    GetRequestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleId()(*string)
    GetRoleInfo()(PrivilegedRoleable)
    GetSchedule()(GovernanceScheduleable)
    GetStatus()(*string)
    GetTicketNumber()(*string)
    GetTicketSystem()(*string)
    GetType()(*string)
    GetUserId()(*string)
    SetAssignmentState(value *string)()
    SetDuration(value *string)()
    SetReason(value *string)()
    SetRequestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleId(value *string)()
    SetRoleInfo(value PrivilegedRoleable)()
    SetSchedule(value GovernanceScheduleable)()
    SetStatus(value *string)()
    SetTicketNumber(value *string)()
    SetTicketSystem(value *string)()
    SetType(value *string)()
    SetUserId(value *string)()
}
