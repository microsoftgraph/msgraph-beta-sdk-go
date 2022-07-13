package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceRoleAssignmentRequest provides operations to manage the collection of accessReviewDecision entities.
type GovernanceRoleAssignmentRequest struct {
    Entity
    // Required. The state of the assignment. The possible values are: Eligible (for eligible assignment),  Active (if it is directly assigned), Active (by administrators, or activated on an eligible assignment by the users).
    assignmentState *string
    // If this is a request for role activation, it represents the id of the eligible assignment being referred; Otherwise, the value is null.
    linkedEligibleRoleAssignmentId *string
    // A message provided by users and administrators when create the request about why it is needed.
    reason *string
    // Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    requestedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Read-only. The resource that the request aims to.
    resource GovernanceResourceable
    // Required. The unique identifier of the Azure resource that is associated with the role assignment request. Azure resources can include subscriptions, resource groups, virtual machines, and SQL databases.
    resourceId *string
    // Read-only. The role definition that the request aims to.
    roleDefinition GovernanceRoleDefinitionable
    // Required. The identifier of the Azure role definition that the role assignment request is associated with.
    roleDefinitionId *string
    // The schedule object of the role assignment request.
    schedule GovernanceScheduleable
    // The status of the role assignment request.
    status GovernanceRoleAssignmentRequestStatusable
    // Read-only. The user/group principal.
    subject GovernanceSubjectable
    // Required. The unique identifier of the principal or subject that the role assignment request is associated with. Principals can be users, groups, or service principals.
    subjectId *string
}
// NewGovernanceRoleAssignmentRequest instantiates a new governanceRoleAssignmentRequest and sets the default values.
func NewGovernanceRoleAssignmentRequest()(*GovernanceRoleAssignmentRequest) {
    m := &GovernanceRoleAssignmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceRoleAssignmentRequest(), nil
}
// GetAssignmentState gets the assignmentState property value. Required. The state of the assignment. The possible values are: Eligible (for eligible assignment),  Active (if it is directly assigned), Active (by administrators, or activated on an eligible assignment by the users).
func (m *GovernanceRoleAssignmentRequest) GetAssignmentState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceRoleAssignmentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["linkedEligibleRoleAssignmentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkedEligibleRoleAssignmentId(val)
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
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(GovernanceResourceable))
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["roleDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinition(val.(GovernanceRoleDefinitionable))
        }
        return nil
    }
    res["roleDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionId(val)
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
        val, err := n.GetObjectValue(CreateGovernanceRoleAssignmentRequestStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(GovernanceRoleAssignmentRequestStatusable))
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val.(GovernanceSubjectable))
        }
        return nil
    }
    res["subjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectId(val)
        }
        return nil
    }
    return res
}
// GetLinkedEligibleRoleAssignmentId gets the linkedEligibleRoleAssignmentId property value. If this is a request for role activation, it represents the id of the eligible assignment being referred; Otherwise, the value is null.
func (m *GovernanceRoleAssignmentRequest) GetLinkedEligibleRoleAssignmentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.linkedEligibleRoleAssignmentId
    }
}
// GetReason gets the reason property value. A message provided by users and administrators when create the request about why it is needed.
func (m *GovernanceRoleAssignmentRequest) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// GetRequestedDateTime gets the requestedDateTime property value. Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleAssignmentRequest) GetRequestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.requestedDateTime
    }
}
// GetResource gets the resource property value. Read-only. The resource that the request aims to.
func (m *GovernanceRoleAssignmentRequest) GetResource()(GovernanceResourceable) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetResourceId gets the resourceId property value. Required. The unique identifier of the Azure resource that is associated with the role assignment request. Azure resources can include subscriptions, resource groups, virtual machines, and SQL databases.
func (m *GovernanceRoleAssignmentRequest) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// GetRoleDefinition gets the roleDefinition property value. Read-only. The role definition that the request aims to.
func (m *GovernanceRoleAssignmentRequest) GetRoleDefinition()(GovernanceRoleDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Required. The identifier of the Azure role definition that the role assignment request is associated with.
func (m *GovernanceRoleAssignmentRequest) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
// GetSchedule gets the schedule property value. The schedule object of the role assignment request.
func (m *GovernanceRoleAssignmentRequest) GetSchedule()(GovernanceScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// GetStatus gets the status property value. The status of the role assignment request.
func (m *GovernanceRoleAssignmentRequest) GetStatus()(GovernanceRoleAssignmentRequestStatusable) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetSubject gets the subject property value. Read-only. The user/group principal.
func (m *GovernanceRoleAssignmentRequest) GetSubject()(GovernanceSubjectable) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// GetSubjectId gets the subjectId property value. Required. The unique identifier of the principal or subject that the role assignment request is associated with. Principals can be users, groups, or service principals.
func (m *GovernanceRoleAssignmentRequest) GetSubjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectId
    }
}
// Serialize serializes information the current object
func (m *GovernanceRoleAssignmentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("linkedEligibleRoleAssignmentId", m.GetLinkedEligibleRoleAssignmentId())
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
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleDefinition", m.GetRoleDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleDefinitionId", m.GetRoleDefinitionId())
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
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectId", m.GetSubjectId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignmentState sets the assignmentState property value. Required. The state of the assignment. The possible values are: Eligible (for eligible assignment),  Active (if it is directly assigned), Active (by administrators, or activated on an eligible assignment by the users).
func (m *GovernanceRoleAssignmentRequest) SetAssignmentState(value *string)() {
    if m != nil {
        m.assignmentState = value
    }
}
// SetLinkedEligibleRoleAssignmentId sets the linkedEligibleRoleAssignmentId property value. If this is a request for role activation, it represents the id of the eligible assignment being referred; Otherwise, the value is null.
func (m *GovernanceRoleAssignmentRequest) SetLinkedEligibleRoleAssignmentId(value *string)() {
    if m != nil {
        m.linkedEligibleRoleAssignmentId = value
    }
}
// SetReason sets the reason property value. A message provided by users and administrators when create the request about why it is needed.
func (m *GovernanceRoleAssignmentRequest) SetReason(value *string)() {
    if m != nil {
        m.reason = value
    }
}
// SetRequestedDateTime sets the requestedDateTime property value. Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleAssignmentRequest) SetRequestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.requestedDateTime = value
    }
}
// SetResource sets the resource property value. Read-only. The resource that the request aims to.
func (m *GovernanceRoleAssignmentRequest) SetResource(value GovernanceResourceable)() {
    if m != nil {
        m.resource = value
    }
}
// SetResourceId sets the resourceId property value. Required. The unique identifier of the Azure resource that is associated with the role assignment request. Azure resources can include subscriptions, resource groups, virtual machines, and SQL databases.
func (m *GovernanceRoleAssignmentRequest) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
// SetRoleDefinition sets the roleDefinition property value. Read-only. The role definition that the request aims to.
func (m *GovernanceRoleAssignmentRequest) SetRoleDefinition(value GovernanceRoleDefinitionable)() {
    if m != nil {
        m.roleDefinition = value
    }
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Required. The identifier of the Azure role definition that the role assignment request is associated with.
func (m *GovernanceRoleAssignmentRequest) SetRoleDefinitionId(value *string)() {
    if m != nil {
        m.roleDefinitionId = value
    }
}
// SetSchedule sets the schedule property value. The schedule object of the role assignment request.
func (m *GovernanceRoleAssignmentRequest) SetSchedule(value GovernanceScheduleable)() {
    if m != nil {
        m.schedule = value
    }
}
// SetStatus sets the status property value. The status of the role assignment request.
func (m *GovernanceRoleAssignmentRequest) SetStatus(value GovernanceRoleAssignmentRequestStatusable)() {
    if m != nil {
        m.status = value
    }
}
// SetSubject sets the subject property value. Read-only. The user/group principal.
func (m *GovernanceRoleAssignmentRequest) SetSubject(value GovernanceSubjectable)() {
    if m != nil {
        m.subject = value
    }
}
// SetSubjectId sets the subjectId property value. Required. The unique identifier of the principal or subject that the role assignment request is associated with. Principals can be users, groups, or service principals.
func (m *GovernanceRoleAssignmentRequest) SetSubjectId(value *string)() {
    if m != nil {
        m.subjectId = value
    }
}
