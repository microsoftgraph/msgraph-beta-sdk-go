package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessPackageAssignment struct {
    Entity
    accessPackage *AccessPackage;
    accessPackageAssignmentPolicy *AccessPackageAssignmentPolicy;
    accessPackageAssignmentRequests []AccessPackageAssignmentRequest;
    accessPackageAssignmentResourceRoles []AccessPackageAssignmentResourceRole;
    accessPackageId *string;
    assignmentPolicyId *string;
    assignmentState *string;
    assignmentStatus *string;
    catalogId *string;
    expiredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    isExtended *bool;
    schedule *RequestSchedule;
    target *AccessPackageSubject;
    targetId *string;
}
func NewAccessPackageAssignment()(*AccessPackageAssignment) {
    m := &AccessPackageAssignment{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AccessPackageAssignment) GetAccessPackage()(*AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackage
    }
}
func (m *AccessPackageAssignment) GetAccessPackageAssignmentPolicy()(*AccessPackageAssignmentPolicy) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentPolicy
    }
}
func (m *AccessPackageAssignment) GetAccessPackageAssignmentRequests()([]AccessPackageAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentRequests
    }
}
func (m *AccessPackageAssignment) GetAccessPackageAssignmentResourceRoles()([]AccessPackageAssignmentResourceRole) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentResourceRoles
    }
}
func (m *AccessPackageAssignment) GetAccessPackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageId
    }
}
func (m *AccessPackageAssignment) GetAssignmentPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentPolicyId
    }
}
func (m *AccessPackageAssignment) GetAssignmentState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentState
    }
}
func (m *AccessPackageAssignment) GetAssignmentStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentStatus
    }
}
func (m *AccessPackageAssignment) GetCatalogId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogId
    }
}
func (m *AccessPackageAssignment) GetExpiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expiredDateTime
    }
}
func (m *AccessPackageAssignment) GetIsExtended()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isExtended
    }
}
func (m *AccessPackageAssignment) GetSchedule()(*RequestSchedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
func (m *AccessPackageAssignment) GetTarget()(*AccessPackageSubject) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *AccessPackageAssignment) GetTargetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetId
    }
}
func (m *AccessPackageAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        m.SetAccessPackage(val.(*AccessPackage))
        return nil
    }
    res["accessPackageAssignmentPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignmentPolicy() })
        if err != nil {
            return err
        }
        m.SetAccessPackageAssignmentPolicy(val.(*AccessPackageAssignmentPolicy))
        return nil
    }
    res["accessPackageAssignmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignmentRequest() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageAssignmentRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageAssignmentRequest))
        }
        m.SetAccessPackageAssignmentRequests(res)
        return nil
    }
    res["accessPackageAssignmentResourceRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignmentResourceRole() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageAssignmentResourceRole, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageAssignmentResourceRole))
        }
        m.SetAccessPackageAssignmentResourceRoles(res)
        return nil
    }
    res["accessPackageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAccessPackageId(val)
        return nil
    }
    res["assignmentPolicyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignmentPolicyId(val)
        return nil
    }
    res["assignmentState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignmentState(val)
        return nil
    }
    res["assignmentStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignmentStatus(val)
        return nil
    }
    res["catalogId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCatalogId(val)
        return nil
    }
    res["expiredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpiredDateTime(val)
        return nil
    }
    res["isExtended"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsExtended(val)
        return nil
    }
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRequestSchedule() })
        if err != nil {
            return err
        }
        m.SetSchedule(val.(*RequestSchedule))
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageSubject() })
        if err != nil {
            return err
        }
        m.SetTarget(val.(*AccessPackageSubject))
        return nil
    }
    res["targetId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetId(val)
        return nil
    }
    return res
}
func (m *AccessPackageAssignment) IsNil()(bool) {
    return m == nil
}
func (m *AccessPackageAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessPackage", m.GetAccessPackage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageAssignmentPolicy", m.GetAccessPackageAssignmentPolicy())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageAssignmentRequests()))
        for i, v := range m.GetAccessPackageAssignmentRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageAssignmentResourceRoles()))
        for i, v := range m.GetAccessPackageAssignmentResourceRoles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentResourceRoles", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("accessPackageId", m.GetAccessPackageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assignmentPolicyId", m.GetAssignmentPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assignmentState", m.GetAssignmentState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assignmentStatus", m.GetAssignmentStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("catalogId", m.GetCatalogId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expiredDateTime", m.GetExpiredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isExtended", m.GetIsExtended())
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
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetId", m.GetTargetId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AccessPackageAssignment) SetAccessPackage(value *AccessPackage)() {
    m.accessPackage = value
}
func (m *AccessPackageAssignment) SetAccessPackageAssignmentPolicy(value *AccessPackageAssignmentPolicy)() {
    m.accessPackageAssignmentPolicy = value
}
func (m *AccessPackageAssignment) SetAccessPackageAssignmentRequests(value []AccessPackageAssignmentRequest)() {
    m.accessPackageAssignmentRequests = value
}
func (m *AccessPackageAssignment) SetAccessPackageAssignmentResourceRoles(value []AccessPackageAssignmentResourceRole)() {
    m.accessPackageAssignmentResourceRoles = value
}
func (m *AccessPackageAssignment) SetAccessPackageId(value *string)() {
    m.accessPackageId = value
}
func (m *AccessPackageAssignment) SetAssignmentPolicyId(value *string)() {
    m.assignmentPolicyId = value
}
func (m *AccessPackageAssignment) SetAssignmentState(value *string)() {
    m.assignmentState = value
}
func (m *AccessPackageAssignment) SetAssignmentStatus(value *string)() {
    m.assignmentStatus = value
}
func (m *AccessPackageAssignment) SetCatalogId(value *string)() {
    m.catalogId = value
}
func (m *AccessPackageAssignment) SetExpiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiredDateTime = value
}
func (m *AccessPackageAssignment) SetIsExtended(value *bool)() {
    m.isExtended = value
}
func (m *AccessPackageAssignment) SetSchedule(value *RequestSchedule)() {
    m.schedule = value
}
func (m *AccessPackageAssignment) SetTarget(value *AccessPackageSubject)() {
    m.target = value
}
func (m *AccessPackageAssignment) SetTargetId(value *string)() {
    m.targetId = value
}
