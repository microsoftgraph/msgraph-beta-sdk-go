package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GovernanceRoleAssignmentRequest struct {
    Entity
    // Required. The state of the assignment. The possible values are: Eligible (for eligible assignment),  Active (if it is directly assigned), Active (by administrators, or activated on an eligible assignment by the users).
    assignmentState *string;
    // If this is a request for role activation, it represents the id of the eligible assignment being referred; Otherwise, the value is null.
    linkedEligibleRoleAssignmentId *string;
    // A message provided by users and administrators when create the request about why it is needed.
    reason *string;
    // Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    requestedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only. The resource that the request aims to.
    resource *GovernanceResource;
    // Required. The id of the resource which the role assignment request is associated with.
    resourceId *string;
    // Read-only. The role definition that the request aims to.
    roleDefinition *GovernanceRoleDefinition;
    // Required. The id of the role definition which the role assignment request is associated with.
    roleDefinitionId *string;
    // The schedule object of the role assignment request.
    schedule *GovernanceSchedule;
    // The status of the role assignment request.
    status *GovernanceRoleAssignmentRequestStatus;
    // Read-only. The user/group principal.
    subject *GovernanceSubject;
    // Required. The id of the subject which the role assignment request is associated with.
    subjectId *string;
    // Required. Representing the type of the operation on the role assignment. The possible values are: AdminAdd , UserAdd , AdminUpdate , AdminRemove , UserRemove , UserExtend , AdminExtend , UserRenew , AdminRenew.
    type_escaped *string;
}
// Instantiates a new governanceRoleAssignmentRequest and sets the default values.
func NewGovernanceRoleAssignmentRequest()(*GovernanceRoleAssignmentRequest) {
    m := &GovernanceRoleAssignmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignmentState property value. Required. The state of the assignment. The possible values are: Eligible (for eligible assignment),  Active (if it is directly assigned), Active (by administrators, or activated on an eligible assignment by the users).
func (m *GovernanceRoleAssignmentRequest) GetAssignmentState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentState
    }
}
// Gets the linkedEligibleRoleAssignmentId property value. If this is a request for role activation, it represents the id of the eligible assignment being referred; Otherwise, the value is null.
func (m *GovernanceRoleAssignmentRequest) GetLinkedEligibleRoleAssignmentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.linkedEligibleRoleAssignmentId
    }
}
// Gets the reason property value. A message provided by users and administrators when create the request about why it is needed.
func (m *GovernanceRoleAssignmentRequest) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// Gets the requestedDateTime property value. Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleAssignmentRequest) GetRequestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.requestedDateTime
    }
}
// Gets the resource property value. Read-only. The resource that the request aims to.
func (m *GovernanceRoleAssignmentRequest) GetResource()(*GovernanceResource) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// Gets the resourceId property value. Required. The id of the resource which the role assignment request is associated with.
func (m *GovernanceRoleAssignmentRequest) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// Gets the roleDefinition property value. Read-only. The role definition that the request aims to.
func (m *GovernanceRoleAssignmentRequest) GetRoleDefinition()(*GovernanceRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
// Gets the roleDefinitionId property value. Required. The id of the role definition which the role assignment request is associated with.
func (m *GovernanceRoleAssignmentRequest) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
// Gets the schedule property value. The schedule object of the role assignment request.
func (m *GovernanceRoleAssignmentRequest) GetSchedule()(*GovernanceSchedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// Gets the status property value. The status of the role assignment request.
func (m *GovernanceRoleAssignmentRequest) GetStatus()(*GovernanceRoleAssignmentRequestStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the subject property value. Read-only. The user/group principal.
func (m *GovernanceRoleAssignmentRequest) GetSubject()(*GovernanceSubject) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// Gets the subjectId property value. Required. The id of the subject which the role assignment request is associated with.
func (m *GovernanceRoleAssignmentRequest) GetSubjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectId
    }
}
// Gets the type_escaped property value. Required. Representing the type of the operation on the role assignment. The possible values are: AdminAdd , UserAdd , AdminUpdate , AdminRemove , UserRemove , UserExtend , AdminExtend , UserRenew , AdminRenew.
func (m *GovernanceRoleAssignmentRequest) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *GovernanceRoleAssignmentRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["linkedEligibleRoleAssignmentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkedEligibleRoleAssignmentId(val)
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
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceResource() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(*GovernanceResource))
        }
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["roleDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinition(val.(*GovernanceRoleDefinition))
        }
        return nil
    }
    res["roleDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionId(val)
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleAssignmentRequestStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*GovernanceRoleAssignmentRequestStatus))
        }
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceSubject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val.(*GovernanceSubject))
        }
        return nil
    }
    res["subjectId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectId(val)
        }
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType_escaped(val)
        }
        return nil
    }
    return res
}
func (m *GovernanceRoleAssignmentRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GovernanceRoleAssignmentRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("type_escaped", m.GetType_escaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignmentState property value. Required. The state of the assignment. The possible values are: Eligible (for eligible assignment),  Active (if it is directly assigned), Active (by administrators, or activated on an eligible assignment by the users).
// Parameters:
//  - value : Value to set for the assignmentState property.
func (m *GovernanceRoleAssignmentRequest) SetAssignmentState(value *string)() {
    m.assignmentState = value
}
// Sets the linkedEligibleRoleAssignmentId property value. If this is a request for role activation, it represents the id of the eligible assignment being referred; Otherwise, the value is null.
// Parameters:
//  - value : Value to set for the linkedEligibleRoleAssignmentId property.
func (m *GovernanceRoleAssignmentRequest) SetLinkedEligibleRoleAssignmentId(value *string)() {
    m.linkedEligibleRoleAssignmentId = value
}
// Sets the reason property value. A message provided by users and administrators when create the request about why it is needed.
// Parameters:
//  - value : Value to set for the reason property.
func (m *GovernanceRoleAssignmentRequest) SetReason(value *string)() {
    m.reason = value
}
// Sets the requestedDateTime property value. Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the requestedDateTime property.
func (m *GovernanceRoleAssignmentRequest) SetRequestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestedDateTime = value
}
// Sets the resource property value. Read-only. The resource that the request aims to.
// Parameters:
//  - value : Value to set for the resource property.
func (m *GovernanceRoleAssignmentRequest) SetResource(value *GovernanceResource)() {
    m.resource = value
}
// Sets the resourceId property value. Required. The id of the resource which the role assignment request is associated with.
// Parameters:
//  - value : Value to set for the resourceId property.
func (m *GovernanceRoleAssignmentRequest) SetResourceId(value *string)() {
    m.resourceId = value
}
// Sets the roleDefinition property value. Read-only. The role definition that the request aims to.
// Parameters:
//  - value : Value to set for the roleDefinition property.
func (m *GovernanceRoleAssignmentRequest) SetRoleDefinition(value *GovernanceRoleDefinition)() {
    m.roleDefinition = value
}
// Sets the roleDefinitionId property value. Required. The id of the role definition which the role assignment request is associated with.
// Parameters:
//  - value : Value to set for the roleDefinitionId property.
func (m *GovernanceRoleAssignmentRequest) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
// Sets the schedule property value. The schedule object of the role assignment request.
// Parameters:
//  - value : Value to set for the schedule property.
func (m *GovernanceRoleAssignmentRequest) SetSchedule(value *GovernanceSchedule)() {
    m.schedule = value
}
// Sets the status property value. The status of the role assignment request.
// Parameters:
//  - value : Value to set for the status property.
func (m *GovernanceRoleAssignmentRequest) SetStatus(value *GovernanceRoleAssignmentRequestStatus)() {
    m.status = value
}
// Sets the subject property value. Read-only. The user/group principal.
// Parameters:
//  - value : Value to set for the subject property.
func (m *GovernanceRoleAssignmentRequest) SetSubject(value *GovernanceSubject)() {
    m.subject = value
}
// Sets the subjectId property value. Required. The id of the subject which the role assignment request is associated with.
// Parameters:
//  - value : Value to set for the subjectId property.
func (m *GovernanceRoleAssignmentRequest) SetSubjectId(value *string)() {
    m.subjectId = value
}
// Sets the type_escaped property value. Required. Representing the type of the operation on the role assignment. The possible values are: AdminAdd , UserAdd , AdminUpdate , AdminRemove , UserRemove , UserExtend , AdminExtend , UserRenew , AdminRenew.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *GovernanceRoleAssignmentRequest) SetType_escaped(value *string)() {
    m.type_escaped = value
}
