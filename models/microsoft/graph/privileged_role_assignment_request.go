package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivilegedRoleAssignmentRequest 
type PrivilegedRoleAssignmentRequest struct {
    Entity
    // The state of the assignment. The value can be Eligible for eligible assignment Active - if it is directly assigned Active by administrators, or activated on an eligible assignment by the users.
    assignmentState *string;
    // The duration of a role assignment.
    duration *string;
    // The reason for the role assignment.
    reason *string;
    // Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    requestedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The id of the role.
    roleId *string;
    // The roleInfo object of the role assignment request.
    roleInfo *PrivilegedRole;
    // The schedule object of the role assignment request.
    schedule *GovernanceSchedule;
    // Read-only.The status of the role assignment request. The value can be NotStarted,Completed,RequestedApproval,Scheduled,Approved,ApprovalDenied,ApprovalAborted,Cancelling,Cancelled,Revoked,RequestExpired.
    status *string;
    // The ticketNumber for the role assignment.
    ticketNumber *string;
    // The ticketSystem for the role assignment.
    ticketSystem *string;
    // Representing the type of the operation on the role assignment. The value can be AdminAdd: Administrators add users to roles;UserAdd: Users add role assignments.
    type_escaped *string;
    // The id of the user.
    userId *string;
}
// NewPrivilegedRoleAssignmentRequest instantiates a new privilegedRoleAssignmentRequest and sets the default values.
func NewPrivilegedRoleAssignmentRequest()(*PrivilegedRoleAssignmentRequest) {
    m := &PrivilegedRoleAssignmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignmentState gets the assignmentState property value. The state of the assignment. The value can be Eligible for eligible assignment Active - if it is directly assigned Active by administrators, or activated on an eligible assignment by the users.
func (m *PrivilegedRoleAssignmentRequest) GetAssignmentState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentState
    }
}
// GetDuration gets the duration property value. The duration of a role assignment.
func (m *PrivilegedRoleAssignmentRequest) GetDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// GetReason gets the reason property value. The reason for the role assignment.
func (m *PrivilegedRoleAssignmentRequest) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// GetRequestedDateTime gets the requestedDateTime property value. Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *PrivilegedRoleAssignmentRequest) GetRequestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.requestedDateTime
    }
}
// GetRoleId gets the roleId property value. The id of the role.
func (m *PrivilegedRoleAssignmentRequest) GetRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleId
    }
}
// GetRoleInfo gets the roleInfo property value. The roleInfo object of the role assignment request.
func (m *PrivilegedRoleAssignmentRequest) GetRoleInfo()(*PrivilegedRole) {
    if m == nil {
        return nil
    } else {
        return m.roleInfo
    }
}
// GetSchedule gets the schedule property value. The schedule object of the role assignment request.
func (m *PrivilegedRoleAssignmentRequest) GetSchedule()(*GovernanceSchedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// GetStatus gets the status property value. Read-only.The status of the role assignment request. The value can be NotStarted,Completed,RequestedApproval,Scheduled,Approved,ApprovalDenied,ApprovalAborted,Cancelling,Cancelled,Revoked,RequestExpired.
func (m *PrivilegedRoleAssignmentRequest) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTicketNumber gets the ticketNumber property value. The ticketNumber for the role assignment.
func (m *PrivilegedRoleAssignmentRequest) GetTicketNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ticketNumber
    }
}
// GetTicketSystem gets the ticketSystem property value. The ticketSystem for the role assignment.
func (m *PrivilegedRoleAssignmentRequest) GetTicketSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ticketSystem
    }
}
// GetType gets the type property value. Representing the type of the operation on the role assignment. The value can be AdminAdd: Administrators add users to roles;UserAdd: Users add role assignments.
func (m *PrivilegedRoleAssignmentRequest) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetUserId gets the userId property value. The id of the user.
func (m *PrivilegedRoleAssignmentRequest) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRoleAssignmentRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentState(val)
        }
        return nil
    }
    res["duration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["requestedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedDateTime(val)
        }
        return nil
    }
    res["roleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRole() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleInfo(val.(*PrivilegedRole))
        }
        return nil
    }
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceSchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(*GovernanceSchedule))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["ticketNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketNumber(val)
        }
        return nil
    }
    res["ticketSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketSystem(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *PrivilegedRoleAssignmentRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrivilegedRoleAssignmentRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignmentState", m.GetAssignmentState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestedDateTime", m.GetRequestedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleId", m.GetRoleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleInfo", m.GetRoleInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedule", m.GetSchedule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ticketNumber", m.GetTicketNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ticketSystem", m.GetTicketSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignmentState sets the assignmentState property value. The state of the assignment. The value can be Eligible for eligible assignment Active - if it is directly assigned Active by administrators, or activated on an eligible assignment by the users.
func (m *PrivilegedRoleAssignmentRequest) SetAssignmentState(value *string)() {
    if m != nil {
        m.assignmentState = value
    }
}
// SetDuration sets the duration property value. The duration of a role assignment.
func (m *PrivilegedRoleAssignmentRequest) SetDuration(value *string)() {
    if m != nil {
        m.duration = value
    }
}
// SetReason sets the reason property value. The reason for the role assignment.
func (m *PrivilegedRoleAssignmentRequest) SetReason(value *string)() {
    if m != nil {
        m.reason = value
    }
}
// SetRequestedDateTime sets the requestedDateTime property value. Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *PrivilegedRoleAssignmentRequest) SetRequestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.requestedDateTime = value
    }
}
// SetRoleId sets the roleId property value. The id of the role.
func (m *PrivilegedRoleAssignmentRequest) SetRoleId(value *string)() {
    if m != nil {
        m.roleId = value
    }
}
// SetRoleInfo sets the roleInfo property value. The roleInfo object of the role assignment request.
func (m *PrivilegedRoleAssignmentRequest) SetRoleInfo(value *PrivilegedRole)() {
    if m != nil {
        m.roleInfo = value
    }
}
// SetSchedule sets the schedule property value. The schedule object of the role assignment request.
func (m *PrivilegedRoleAssignmentRequest) SetSchedule(value *GovernanceSchedule)() {
    if m != nil {
        m.schedule = value
    }
}
// SetStatus sets the status property value. Read-only.The status of the role assignment request. The value can be NotStarted,Completed,RequestedApproval,Scheduled,Approved,ApprovalDenied,ApprovalAborted,Cancelling,Cancelled,Revoked,RequestExpired.
func (m *PrivilegedRoleAssignmentRequest) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetTicketNumber sets the ticketNumber property value. The ticketNumber for the role assignment.
func (m *PrivilegedRoleAssignmentRequest) SetTicketNumber(value *string)() {
    if m != nil {
        m.ticketNumber = value
    }
}
// SetTicketSystem sets the ticketSystem property value. The ticketSystem for the role assignment.
func (m *PrivilegedRoleAssignmentRequest) SetTicketSystem(value *string)() {
    if m != nil {
        m.ticketSystem = value
    }
}
// SetType sets the type property value. Representing the type of the operation on the role assignment. The value can be AdminAdd: Administrators add users to roles;UserAdd: Users add role assignments.
func (m *PrivilegedRoleAssignmentRequest) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetUserId sets the userId property value. The id of the user.
func (m *PrivilegedRoleAssignmentRequest) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
