package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceRoleAssignmentRequest 
type GovernanceRoleAssignmentRequest struct {
    Entity
    // Required. Representing the type of the operation on the role assignment. The possible values are: AdminAdd , UserAdd , AdminUpdate , AdminRemove , UserRemove , UserExtend , AdminExtend , UserRenew , AdminRenew.
    TypeEscaped *string
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
    val, err := m.GetBackingStore().Get("assignmentState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    return res
}
// GetLinkedEligibleRoleAssignmentId gets the linkedEligibleRoleAssignmentId property value. If this is a request for role activation, it represents the id of the eligible assignment being referred; Otherwise, the value is null.
func (m *GovernanceRoleAssignmentRequest) GetLinkedEligibleRoleAssignmentId()(*string) {
    val, err := m.GetBackingStore().Get("linkedEligibleRoleAssignmentId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReason gets the reason property value. A message provided by users and administrators when create the request about why it is needed.
func (m *GovernanceRoleAssignmentRequest) GetReason()(*string) {
    val, err := m.GetBackingStore().Get("reason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestedDateTime gets the requestedDateTime property value. Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleAssignmentRequest) GetRequestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetResource gets the resource property value. Read-only. The resource that the request aims to.
func (m *GovernanceRoleAssignmentRequest) GetResource()(GovernanceResourceable) {
    val, err := m.GetBackingStore().Get("resource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GovernanceResourceable)
    }
    return nil
}
// GetResourceId gets the resourceId property value. Required. The unique identifier of the Azure resource that is associated with the role assignment request. Azure resources can include subscriptions, resource groups, virtual machines, and SQL databases.
func (m *GovernanceRoleAssignmentRequest) GetResourceId()(*string) {
    val, err := m.GetBackingStore().Get("resourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleDefinition gets the roleDefinition property value. Read-only. The role definition that the request aims to.
func (m *GovernanceRoleAssignmentRequest) GetRoleDefinition()(GovernanceRoleDefinitionable) {
    val, err := m.GetBackingStore().Get("roleDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GovernanceRoleDefinitionable)
    }
    return nil
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Required. The identifier of the Azure role definition that the role assignment request is associated with.
func (m *GovernanceRoleAssignmentRequest) GetRoleDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("roleDefinitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSchedule gets the schedule property value. The schedule object of the role assignment request.
func (m *GovernanceRoleAssignmentRequest) GetSchedule()(GovernanceScheduleable) {
    val, err := m.GetBackingStore().Get("schedule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GovernanceScheduleable)
    }
    return nil
}
// GetStatus gets the status property value. The status of the role assignment request.
func (m *GovernanceRoleAssignmentRequest) GetStatus()(GovernanceRoleAssignmentRequestStatusable) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GovernanceRoleAssignmentRequestStatusable)
    }
    return nil
}
// GetSubject gets the subject property value. Read-only. The user/group principal.
func (m *GovernanceRoleAssignmentRequest) GetSubject()(GovernanceSubjectable) {
    val, err := m.GetBackingStore().Get("subject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GovernanceSubjectable)
    }
    return nil
}
// GetSubjectId gets the subjectId property value. Required. The unique identifier of the principal or subject that the role assignment request is associated with. Principals can be users, groups, or service principals.
func (m *GovernanceRoleAssignmentRequest) GetSubjectId()(*string) {
    val, err := m.GetBackingStore().Get("subjectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetType gets the type property value. Required. Representing the type of the operation on the role assignment. The possible values are: AdminAdd , UserAdd , AdminUpdate , AdminRemove , UserRemove , UserExtend , AdminExtend , UserRenew , AdminRenew.
func (m *GovernanceRoleAssignmentRequest) GetType()(*string) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignmentState sets the assignmentState property value. Required. The state of the assignment. The possible values are: Eligible (for eligible assignment),  Active (if it is directly assigned), Active (by administrators, or activated on an eligible assignment by the users).
func (m *GovernanceRoleAssignmentRequest) SetAssignmentState(value *string)() {
    err := m.GetBackingStore().Set("assignmentState", value)
    if err != nil {
        panic(err)
    }
}
// SetLinkedEligibleRoleAssignmentId sets the linkedEligibleRoleAssignmentId property value. If this is a request for role activation, it represents the id of the eligible assignment being referred; Otherwise, the value is null.
func (m *GovernanceRoleAssignmentRequest) SetLinkedEligibleRoleAssignmentId(value *string)() {
    err := m.GetBackingStore().Set("linkedEligibleRoleAssignmentId", value)
    if err != nil {
        panic(err)
    }
}
// SetReason sets the reason property value. A message provided by users and administrators when create the request about why it is needed.
func (m *GovernanceRoleAssignmentRequest) SetReason(value *string)() {
    err := m.GetBackingStore().Set("reason", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedDateTime sets the requestedDateTime property value. Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleAssignmentRequest) SetRequestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetResource sets the resource property value. Read-only. The resource that the request aims to.
func (m *GovernanceRoleAssignmentRequest) SetResource(value GovernanceResourceable)() {
    err := m.GetBackingStore().Set("resource", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceId sets the resourceId property value. Required. The unique identifier of the Azure resource that is associated with the role assignment request. Azure resources can include subscriptions, resource groups, virtual machines, and SQL databases.
func (m *GovernanceRoleAssignmentRequest) SetResourceId(value *string)() {
    err := m.GetBackingStore().Set("resourceId", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinition sets the roleDefinition property value. Read-only. The role definition that the request aims to.
func (m *GovernanceRoleAssignmentRequest) SetRoleDefinition(value GovernanceRoleDefinitionable)() {
    err := m.GetBackingStore().Set("roleDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Required. The identifier of the Azure role definition that the role assignment request is associated with.
func (m *GovernanceRoleAssignmentRequest) SetRoleDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("roleDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetSchedule sets the schedule property value. The schedule object of the role assignment request.
func (m *GovernanceRoleAssignmentRequest) SetSchedule(value GovernanceScheduleable)() {
    err := m.GetBackingStore().Set("schedule", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status of the role assignment request.
func (m *GovernanceRoleAssignmentRequest) SetStatus(value GovernanceRoleAssignmentRequestStatusable)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetSubject sets the subject property value. Read-only. The user/group principal.
func (m *GovernanceRoleAssignmentRequest) SetSubject(value GovernanceSubjectable)() {
    err := m.GetBackingStore().Set("subject", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectId sets the subjectId property value. Required. The unique identifier of the principal or subject that the role assignment request is associated with. Principals can be users, groups, or service principals.
func (m *GovernanceRoleAssignmentRequest) SetSubjectId(value *string)() {
    err := m.GetBackingStore().Set("subjectId", value)
    if err != nil {
        panic(err)
    }
}
// SetType sets the type property value. Required. Representing the type of the operation on the role assignment. The possible values are: AdminAdd , UserAdd , AdminUpdate , AdminRemove , UserRemove , UserExtend , AdminExtend , UserRenew , AdminRenew.
func (m *GovernanceRoleAssignmentRequest) SetType(value *string)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// GovernanceRoleAssignmentRequestable 
type GovernanceRoleAssignmentRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
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
