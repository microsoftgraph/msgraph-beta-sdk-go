package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GovernanceRoleAssignmentRequest struct {
    Entity
    assignmentState *string;
    linkedEligibleRoleAssignmentId *string;
    reason *string;
    requestedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    resource *GovernanceResource;
    resourceId *string;
    roleDefinition *GovernanceRoleDefinition;
    roleDefinitionId *string;
    schedule *GovernanceSchedule;
    status *GovernanceRoleAssignmentRequestStatus;
    subject *GovernanceSubject;
    subjectId *string;
    type_escpaped *string;
}
func NewGovernanceRoleAssignmentRequest()(*GovernanceRoleAssignmentRequest) {
    m := &GovernanceRoleAssignmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
func (m *GovernanceRoleAssignmentRequest) GetAssignmentState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentState
    }
}
func (m *GovernanceRoleAssignmentRequest) GetLinkedEligibleRoleAssignmentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.linkedEligibleRoleAssignmentId
    }
}
func (m *GovernanceRoleAssignmentRequest) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
func (m *GovernanceRoleAssignmentRequest) GetRequestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.requestedDateTime
    }
}
func (m *GovernanceRoleAssignmentRequest) GetResource()(*GovernanceResource) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
func (m *GovernanceRoleAssignmentRequest) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
func (m *GovernanceRoleAssignmentRequest) GetRoleDefinition()(*GovernanceRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
func (m *GovernanceRoleAssignmentRequest) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
func (m *GovernanceRoleAssignmentRequest) GetSchedule()(*GovernanceSchedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
func (m *GovernanceRoleAssignmentRequest) GetStatus()(*GovernanceRoleAssignmentRequestStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *GovernanceRoleAssignmentRequest) GetSubject()(*GovernanceSubject) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
func (m *GovernanceRoleAssignmentRequest) GetSubjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectId
    }
}
func (m *GovernanceRoleAssignmentRequest) GetType_escpaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *GovernanceRoleAssignmentRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignmentState(val)
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
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReason(val)
        return nil
    }
    res["requestedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRequestedDateTime(val)
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
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceSchedule() })
        if err != nil {
            return err
        }
        m.SetSchedule(val.(*GovernanceSchedule))
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleAssignmentRequestStatus() })
        if err != nil {
            return err
        }
        m.SetStatus(val.(*GovernanceRoleAssignmentRequestStatus))
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
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escpaped(val)
        return nil
    }
    return res
}
func (m *GovernanceRoleAssignmentRequest) IsNil()(bool) {
    return m == nil
}
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
        err = writer.WriteStringValue("type_escpaped", m.GetType_escpaped())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GovernanceRoleAssignmentRequest) SetAssignmentState(value *string)() {
    m.assignmentState = value
}
func (m *GovernanceRoleAssignmentRequest) SetLinkedEligibleRoleAssignmentId(value *string)() {
    m.linkedEligibleRoleAssignmentId = value
}
func (m *GovernanceRoleAssignmentRequest) SetReason(value *string)() {
    m.reason = value
}
func (m *GovernanceRoleAssignmentRequest) SetRequestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestedDateTime = value
}
func (m *GovernanceRoleAssignmentRequest) SetResource(value *GovernanceResource)() {
    m.resource = value
}
func (m *GovernanceRoleAssignmentRequest) SetResourceId(value *string)() {
    m.resourceId = value
}
func (m *GovernanceRoleAssignmentRequest) SetRoleDefinition(value *GovernanceRoleDefinition)() {
    m.roleDefinition = value
}
func (m *GovernanceRoleAssignmentRequest) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
func (m *GovernanceRoleAssignmentRequest) SetSchedule(value *GovernanceSchedule)() {
    m.schedule = value
}
func (m *GovernanceRoleAssignmentRequest) SetStatus(value *GovernanceRoleAssignmentRequestStatus)() {
    m.status = value
}
func (m *GovernanceRoleAssignmentRequest) SetSubject(value *GovernanceSubject)() {
    m.subject = value
}
func (m *GovernanceRoleAssignmentRequest) SetSubjectId(value *string)() {
    m.subjectId = value
}
func (m *GovernanceRoleAssignmentRequest) SetType_escpaped(value *string)() {
    m.type_escpaped = value
}
