package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackageAssignment struct {
    Entity
    // Read-only. Nullable.
    accessPackage *AccessPackage;
    // Read-only. Nullable.
    accessPackageAssignmentPolicy *AccessPackageAssignmentPolicy;
    // 
    accessPackageAssignmentRequests []AccessPackageAssignmentRequest;
    // The resource roles delivered to the target user for this assignment. Read-only. Nullable.
    accessPackageAssignmentResourceRoles []AccessPackageAssignmentResourceRole;
    // The identifier of the access package. Read-only.
    accessPackageId *string;
    // The identifier of the access package assignment policy. Read-only.
    assignmentPolicyId *string;
    // The state of the access package assignment. Possible values are Delivering, Delivered, or Expired. Read-only.
    assignmentState *string;
    // More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
    assignmentStatus *string;
    // The identifier of the catalog containing the access package. Read-only.
    catalogId *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    expiredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Indicates whether the access package assignment is extended. Read-only.
    isExtended *bool;
    // When the access assignment is to be in place. Read-only.
    schedule *RequestSchedule;
    // The subject of the access package assignment. Read-only. Nullable.
    target *AccessPackageSubject;
    // The ID of the subject with the assignment. Read-only.
    targetId *string;
}
// Instantiates a new accessPackageAssignment and sets the default values.
func NewAccessPackageAssignment()(*AccessPackageAssignment) {
    m := &AccessPackageAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessPackage property value. Read-only. Nullable.
func (m *AccessPackageAssignment) GetAccessPackage()(*AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackage
    }
}
// Gets the accessPackageAssignmentPolicy property value. Read-only. Nullable.
func (m *AccessPackageAssignment) GetAccessPackageAssignmentPolicy()(*AccessPackageAssignmentPolicy) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentPolicy
    }
}
// Gets the accessPackageAssignmentRequests property value. 
func (m *AccessPackageAssignment) GetAccessPackageAssignmentRequests()([]AccessPackageAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentRequests
    }
}
// Gets the accessPackageAssignmentResourceRoles property value. The resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignment) GetAccessPackageAssignmentResourceRoles()([]AccessPackageAssignmentResourceRole) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentResourceRoles
    }
}
// Gets the accessPackageId property value. The identifier of the access package. Read-only.
func (m *AccessPackageAssignment) GetAccessPackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageId
    }
}
// Gets the assignmentPolicyId property value. The identifier of the access package assignment policy. Read-only.
func (m *AccessPackageAssignment) GetAssignmentPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentPolicyId
    }
}
// Gets the assignmentState property value. The state of the access package assignment. Possible values are Delivering, Delivered, or Expired. Read-only.
func (m *AccessPackageAssignment) GetAssignmentState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentState
    }
}
// Gets the assignmentStatus property value. More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
func (m *AccessPackageAssignment) GetAssignmentStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentStatus
    }
}
// Gets the catalogId property value. The identifier of the catalog containing the access package. Read-only.
func (m *AccessPackageAssignment) GetCatalogId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogId
    }
}
// Gets the expiredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignment) GetExpiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expiredDateTime
    }
}
// Gets the isExtended property value. Indicates whether the access package assignment is extended. Read-only.
func (m *AccessPackageAssignment) GetIsExtended()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isExtended
    }
}
// Gets the schedule property value. When the access assignment is to be in place. Read-only.
func (m *AccessPackageAssignment) GetSchedule()(*RequestSchedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// Gets the target property value. The subject of the access package assignment. Read-only. Nullable.
func (m *AccessPackageAssignment) GetTarget()(*AccessPackageSubject) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// Gets the targetId property value. The ID of the subject with the assignment. Read-only.
func (m *AccessPackageAssignment) GetTargetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetId
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the accessPackage property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the accessPackage property.
func (m *AccessPackageAssignment) SetAccessPackage(value *AccessPackage)() {
    m.accessPackage = value
}
// Sets the accessPackageAssignmentPolicy property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the accessPackageAssignmentPolicy property.
func (m *AccessPackageAssignment) SetAccessPackageAssignmentPolicy(value *AccessPackageAssignmentPolicy)() {
    m.accessPackageAssignmentPolicy = value
}
// Sets the accessPackageAssignmentRequests property value. 
// Parameters:
//  - value : Value to set for the accessPackageAssignmentRequests property.
func (m *AccessPackageAssignment) SetAccessPackageAssignmentRequests(value []AccessPackageAssignmentRequest)() {
    m.accessPackageAssignmentRequests = value
}
// Sets the accessPackageAssignmentResourceRoles property value. The resource roles delivered to the target user for this assignment. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the accessPackageAssignmentResourceRoles property.
func (m *AccessPackageAssignment) SetAccessPackageAssignmentResourceRoles(value []AccessPackageAssignmentResourceRole)() {
    m.accessPackageAssignmentResourceRoles = value
}
// Sets the accessPackageId property value. The identifier of the access package. Read-only.
// Parameters:
//  - value : Value to set for the accessPackageId property.
func (m *AccessPackageAssignment) SetAccessPackageId(value *string)() {
    m.accessPackageId = value
}
// Sets the assignmentPolicyId property value. The identifier of the access package assignment policy. Read-only.
// Parameters:
//  - value : Value to set for the assignmentPolicyId property.
func (m *AccessPackageAssignment) SetAssignmentPolicyId(value *string)() {
    m.assignmentPolicyId = value
}
// Sets the assignmentState property value. The state of the access package assignment. Possible values are Delivering, Delivered, or Expired. Read-only.
// Parameters:
//  - value : Value to set for the assignmentState property.
func (m *AccessPackageAssignment) SetAssignmentState(value *string)() {
    m.assignmentState = value
}
// Sets the assignmentStatus property value. More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
// Parameters:
//  - value : Value to set for the assignmentStatus property.
func (m *AccessPackageAssignment) SetAssignmentStatus(value *string)() {
    m.assignmentStatus = value
}
// Sets the catalogId property value. The identifier of the catalog containing the access package. Read-only.
// Parameters:
//  - value : Value to set for the catalogId property.
func (m *AccessPackageAssignment) SetCatalogId(value *string)() {
    m.catalogId = value
}
// Sets the expiredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the expiredDateTime property.
func (m *AccessPackageAssignment) SetExpiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiredDateTime = value
}
// Sets the isExtended property value. Indicates whether the access package assignment is extended. Read-only.
// Parameters:
//  - value : Value to set for the isExtended property.
func (m *AccessPackageAssignment) SetIsExtended(value *bool)() {
    m.isExtended = value
}
// Sets the schedule property value. When the access assignment is to be in place. Read-only.
// Parameters:
//  - value : Value to set for the schedule property.
func (m *AccessPackageAssignment) SetSchedule(value *RequestSchedule)() {
    m.schedule = value
}
// Sets the target property value. The subject of the access package assignment. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the target property.
func (m *AccessPackageAssignment) SetTarget(value *AccessPackageSubject)() {
    m.target = value
}
// Sets the targetId property value. The ID of the subject with the assignment. Read-only.
// Parameters:
//  - value : Value to set for the targetId property.
func (m *AccessPackageAssignment) SetTargetId(value *string)() {
    m.targetId = value
}
