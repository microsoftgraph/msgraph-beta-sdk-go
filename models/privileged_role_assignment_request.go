package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRoleAssignmentRequest 
type PrivilegedRoleAssignmentRequest struct {
    Entity
    // Representing the type of the operation on the role assignment. The value can be AdminAdd: Administrators add users to roles;UserAdd: Users add role assignments.
    TypeEscaped *string
}
// NewPrivilegedRoleAssignmentRequest instantiates a new privilegedRoleAssignmentRequest and sets the default values.
func NewPrivilegedRoleAssignmentRequest()(*PrivilegedRoleAssignmentRequest) {
    m := &PrivilegedRoleAssignmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedRoleAssignmentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedRoleAssignmentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedRoleAssignmentRequest(), nil
}
// GetAssignmentState gets the assignmentState property value. The state of the assignment. The value can be Eligible for eligible assignment Active - if it is directly assigned Active by administrators, or activated on an eligible assignment by the users.
func (m *PrivilegedRoleAssignmentRequest) GetAssignmentState()(*string) {
    val, err := m.GetBackingStore().Get("assignmentState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDuration gets the duration property value. The duration of a role assignment.
func (m *PrivilegedRoleAssignmentRequest) GetDuration()(*string) {
    val, err := m.GetBackingStore().Get("duration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRoleAssignmentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentState(val)
        }
        return nil
    }
    res["duration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["requestedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedDateTime(val)
        }
        return nil
    }
    res["roleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleInfo(val.(PrivilegedRoleable))
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(GovernanceScheduleable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["ticketNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketNumber(val)
        }
        return nil
    }
    res["ticketSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketSystem(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetReason gets the reason property value. The reason for the role assignment.
func (m *PrivilegedRoleAssignmentRequest) GetReason()(*string) {
    val, err := m.GetBackingStore().Get("reason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestedDateTime gets the requestedDateTime property value. Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *PrivilegedRoleAssignmentRequest) GetRequestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRoleId gets the roleId property value. The id of the role.
func (m *PrivilegedRoleAssignmentRequest) GetRoleId()(*string) {
    val, err := m.GetBackingStore().Get("roleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleInfo gets the roleInfo property value. The roleInfo object of the role assignment request.
func (m *PrivilegedRoleAssignmentRequest) GetRoleInfo()(PrivilegedRoleable) {
    val, err := m.GetBackingStore().Get("roleInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PrivilegedRoleable)
    }
    return nil
}
// GetSchedule gets the schedule property value. The schedule object of the role assignment request.
func (m *PrivilegedRoleAssignmentRequest) GetSchedule()(GovernanceScheduleable) {
    val, err := m.GetBackingStore().Get("schedule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GovernanceScheduleable)
    }
    return nil
}
// GetStatus gets the status property value. Read-only.The status of the role assignment request. The value can be NotStarted,Completed,RequestedApproval,Scheduled,Approved,ApprovalDenied,ApprovalAborted,Cancelling,Cancelled,Revoked,RequestExpired.
func (m *PrivilegedRoleAssignmentRequest) GetStatus()(*string) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTicketNumber gets the ticketNumber property value. The ticketNumber for the role assignment.
func (m *PrivilegedRoleAssignmentRequest) GetTicketNumber()(*string) {
    val, err := m.GetBackingStore().Get("ticketNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTicketSystem gets the ticketSystem property value. The ticketSystem for the role assignment.
func (m *PrivilegedRoleAssignmentRequest) GetTicketSystem()(*string) {
    val, err := m.GetBackingStore().Get("ticketSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetType gets the type property value. Representing the type of the operation on the role assignment. The value can be AdminAdd: Administrators add users to roles;UserAdd: Users add role assignments.
func (m *PrivilegedRoleAssignmentRequest) GetType()(*string) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. The id of the user.
func (m *PrivilegedRoleAssignmentRequest) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrivilegedRoleAssignmentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    err := m.GetBackingStore().Set("assignmentState", value)
    if err != nil {
        panic(err)
    }
}
// SetDuration sets the duration property value. The duration of a role assignment.
func (m *PrivilegedRoleAssignmentRequest) SetDuration(value *string)() {
    err := m.GetBackingStore().Set("duration", value)
    if err != nil {
        panic(err)
    }
}
// SetReason sets the reason property value. The reason for the role assignment.
func (m *PrivilegedRoleAssignmentRequest) SetReason(value *string)() {
    err := m.GetBackingStore().Set("reason", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedDateTime sets the requestedDateTime property value. Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *PrivilegedRoleAssignmentRequest) SetRequestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleId sets the roleId property value. The id of the role.
func (m *PrivilegedRoleAssignmentRequest) SetRoleId(value *string)() {
    err := m.GetBackingStore().Set("roleId", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleInfo sets the roleInfo property value. The roleInfo object of the role assignment request.
func (m *PrivilegedRoleAssignmentRequest) SetRoleInfo(value PrivilegedRoleable)() {
    err := m.GetBackingStore().Set("roleInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetSchedule sets the schedule property value. The schedule object of the role assignment request.
func (m *PrivilegedRoleAssignmentRequest) SetSchedule(value GovernanceScheduleable)() {
    err := m.GetBackingStore().Set("schedule", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Read-only.The status of the role assignment request. The value can be NotStarted,Completed,RequestedApproval,Scheduled,Approved,ApprovalDenied,ApprovalAborted,Cancelling,Cancelled,Revoked,RequestExpired.
func (m *PrivilegedRoleAssignmentRequest) SetStatus(value *string)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTicketNumber sets the ticketNumber property value. The ticketNumber for the role assignment.
func (m *PrivilegedRoleAssignmentRequest) SetTicketNumber(value *string)() {
    err := m.GetBackingStore().Set("ticketNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetTicketSystem sets the ticketSystem property value. The ticketSystem for the role assignment.
func (m *PrivilegedRoleAssignmentRequest) SetTicketSystem(value *string)() {
    err := m.GetBackingStore().Set("ticketSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetType sets the type property value. Representing the type of the operation on the role assignment. The value can be AdminAdd: Administrators add users to roles;UserAdd: Users add role assignments.
func (m *PrivilegedRoleAssignmentRequest) SetType(value *string)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The id of the user.
func (m *PrivilegedRoleAssignmentRequest) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// PrivilegedRoleAssignmentRequestable 
type PrivilegedRoleAssignmentRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
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
