package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GovernanceRoleAssignment struct {
    Entity
    // The state of the assignment. The value can be Eligible for eligible assignment or Active if it is directly assigned Active by administrators, or activated on an eligible assignment by the users.
    assignmentState *string;
    // For a non-permanent role assignment, this is the time when the role assignment will be expired. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The external ID the resource that is used to identify the role assignment in the provider.
    externalId *string;
    // Read-only. If this is an active assignment and created due to activation on an eligible assignment, it represents the object of that eligible assignment; Otherwise, the value is null.
    linkedEligibleRoleAssignment *GovernanceRoleAssignment;
    // If this is an active assignment and created due to activation on an eligible assignment, it represents the ID of that eligible assignment; Otherwise, the value is null.
    linkedEligibleRoleAssignmentId *string;
    // The type of member. The value can be: Inherited (if the role assignment is inherited from a parent resource scope), Group (if the role assignment is not inherited, but comes from the membership of a group assignment), or User (if the role assignment is neither inherited nor from a group assignment).
    memberType *string;
    // Read-only. The resource associated with the role assignment.
    resource *GovernanceResource;
    // Required. The ID of the resource which the role assignment is associated with.
    resourceId *string;
    // Read-only. The role definition associated with the role assignment.
    roleDefinition *GovernanceRoleDefinition;
    // Required. The ID of the role definition which the role assignment is associated with.
    roleDefinitionId *string;
    // The start time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    status *string;
    // Read-only. The subject associated with the role assignment.
    subject *GovernanceSubject;
    // Required. The ID of the subject which the role assignment is associated with.
    subjectId *string;
}
// Instantiates a new governanceRoleAssignment and sets the default values.
func NewGovernanceRoleAssignment()(*GovernanceRoleAssignment) {
    m := &GovernanceRoleAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignmentState property value. The state of the assignment. The value can be Eligible for eligible assignment or Active if it is directly assigned Active by administrators, or activated on an eligible assignment by the users.
func (m *GovernanceRoleAssignment) GetAssignmentState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentState
    }
}
// Gets the endDateTime property value. For a non-permanent role assignment, this is the time when the role assignment will be expired. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleAssignment) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// Gets the externalId property value. The external ID the resource that is used to identify the role assignment in the provider.
func (m *GovernanceRoleAssignment) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// Gets the linkedEligibleRoleAssignment property value. Read-only. If this is an active assignment and created due to activation on an eligible assignment, it represents the object of that eligible assignment; Otherwise, the value is null.
func (m *GovernanceRoleAssignment) GetLinkedEligibleRoleAssignment()(*GovernanceRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.linkedEligibleRoleAssignment
    }
}
// Gets the linkedEligibleRoleAssignmentId property value. If this is an active assignment and created due to activation on an eligible assignment, it represents the ID of that eligible assignment; Otherwise, the value is null.
func (m *GovernanceRoleAssignment) GetLinkedEligibleRoleAssignmentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.linkedEligibleRoleAssignmentId
    }
}
// Gets the memberType property value. The type of member. The value can be: Inherited (if the role assignment is inherited from a parent resource scope), Group (if the role assignment is not inherited, but comes from the membership of a group assignment), or User (if the role assignment is neither inherited nor from a group assignment).
func (m *GovernanceRoleAssignment) GetMemberType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberType
    }
}
// Gets the resource property value. Read-only. The resource associated with the role assignment.
func (m *GovernanceRoleAssignment) GetResource()(*GovernanceResource) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// Gets the resourceId property value. Required. The ID of the resource which the role assignment is associated with.
func (m *GovernanceRoleAssignment) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// Gets the roleDefinition property value. Read-only. The role definition associated with the role assignment.
func (m *GovernanceRoleAssignment) GetRoleDefinition()(*GovernanceRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
// Gets the roleDefinitionId property value. Required. The ID of the role definition which the role assignment is associated with.
func (m *GovernanceRoleAssignment) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
// Gets the startDateTime property value. The start time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleAssignment) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the status property value. 
func (m *GovernanceRoleAssignment) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the subject property value. Read-only. The subject associated with the role assignment.
func (m *GovernanceRoleAssignment) GetSubject()(*GovernanceSubject) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// Gets the subjectId property value. Required. The ID of the subject which the role assignment is associated with.
func (m *GovernanceRoleAssignment) GetSubjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectId
    }
}
// The deserialization information for the current model
func (m *GovernanceRoleAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignmentState(val)
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEndDateTime(val)
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    res["linkedEligibleRoleAssignment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleAssignment() })
        if err != nil {
            return err
        }
        m.SetLinkedEligibleRoleAssignment(val.(*GovernanceRoleAssignment))
        return nil
    }
    res["linkedEligibleRoleAssignmentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLinkedEligibleRoleAssignmentId(val)
        return nil
    }
    res["memberType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMemberType(val)
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceResource() })
        if err != nil {
            return err
        }
        m.SetResource(val.(*GovernanceResource))
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceId(val)
        return nil
    }
    res["roleDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleDefinition() })
        if err != nil {
            return err
        }
        m.SetRoleDefinition(val.(*GovernanceRoleDefinition))
        return nil
    }
    res["roleDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRoleDefinitionId(val)
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceSubject() })
        if err != nil {
            return err
        }
        m.SetSubject(val.(*GovernanceSubject))
        return nil
    }
    res["subjectId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubjectId(val)
        return nil
    }
    return res
}
func (m *GovernanceRoleAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GovernanceRoleAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("linkedEligibleRoleAssignment", m.GetLinkedEligibleRoleAssignment())
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
        err = writer.WriteStringValue("memberType", m.GetMemberType())
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
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
// Sets the assignmentState property value. The state of the assignment. The value can be Eligible for eligible assignment or Active if it is directly assigned Active by administrators, or activated on an eligible assignment by the users.
// Parameters:
//  - value : Value to set for the assignmentState property.
func (m *GovernanceRoleAssignment) SetAssignmentState(value *string)() {
    m.assignmentState = value
}
// Sets the endDateTime property value. For a non-permanent role assignment, this is the time when the role assignment will be expired. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the endDateTime property.
func (m *GovernanceRoleAssignment) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// Sets the externalId property value. The external ID the resource that is used to identify the role assignment in the provider.
// Parameters:
//  - value : Value to set for the externalId property.
func (m *GovernanceRoleAssignment) SetExternalId(value *string)() {
    m.externalId = value
}
// Sets the linkedEligibleRoleAssignment property value. Read-only. If this is an active assignment and created due to activation on an eligible assignment, it represents the object of that eligible assignment; Otherwise, the value is null.
// Parameters:
//  - value : Value to set for the linkedEligibleRoleAssignment property.
func (m *GovernanceRoleAssignment) SetLinkedEligibleRoleAssignment(value *GovernanceRoleAssignment)() {
    m.linkedEligibleRoleAssignment = value
}
// Sets the linkedEligibleRoleAssignmentId property value. If this is an active assignment and created due to activation on an eligible assignment, it represents the ID of that eligible assignment; Otherwise, the value is null.
// Parameters:
//  - value : Value to set for the linkedEligibleRoleAssignmentId property.
func (m *GovernanceRoleAssignment) SetLinkedEligibleRoleAssignmentId(value *string)() {
    m.linkedEligibleRoleAssignmentId = value
}
// Sets the memberType property value. The type of member. The value can be: Inherited (if the role assignment is inherited from a parent resource scope), Group (if the role assignment is not inherited, but comes from the membership of a group assignment), or User (if the role assignment is neither inherited nor from a group assignment).
// Parameters:
//  - value : Value to set for the memberType property.
func (m *GovernanceRoleAssignment) SetMemberType(value *string)() {
    m.memberType = value
}
// Sets the resource property value. Read-only. The resource associated with the role assignment.
// Parameters:
//  - value : Value to set for the resource property.
func (m *GovernanceRoleAssignment) SetResource(value *GovernanceResource)() {
    m.resource = value
}
// Sets the resourceId property value. Required. The ID of the resource which the role assignment is associated with.
// Parameters:
//  - value : Value to set for the resourceId property.
func (m *GovernanceRoleAssignment) SetResourceId(value *string)() {
    m.resourceId = value
}
// Sets the roleDefinition property value. Read-only. The role definition associated with the role assignment.
// Parameters:
//  - value : Value to set for the roleDefinition property.
func (m *GovernanceRoleAssignment) SetRoleDefinition(value *GovernanceRoleDefinition)() {
    m.roleDefinition = value
}
// Sets the roleDefinitionId property value. Required. The ID of the role definition which the role assignment is associated with.
// Parameters:
//  - value : Value to set for the roleDefinitionId property.
func (m *GovernanceRoleAssignment) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
// Sets the startDateTime property value. The start time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *GovernanceRoleAssignment) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *GovernanceRoleAssignment) SetStatus(value *string)() {
    m.status = value
}
// Sets the subject property value. Read-only. The subject associated with the role assignment.
// Parameters:
//  - value : Value to set for the subject property.
func (m *GovernanceRoleAssignment) SetSubject(value *GovernanceSubject)() {
    m.subject = value
}
// Sets the subjectId property value. Required. The ID of the subject which the role assignment is associated with.
// Parameters:
//  - value : Value to set for the subjectId property.
func (m *GovernanceRoleAssignment) SetSubjectId(value *string)() {
    m.subjectId = value
}
