package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignment 
type AccessPackageAssignment struct {
    Entity
    // Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
    accessPackage *AccessPackage;
    // Read-only. Nullable. Supports $filter (eq) on the id property
    accessPackageAssignmentPolicy *AccessPackageAssignmentPolicy;
    // 
    accessPackageAssignmentRequests []AccessPackageAssignmentRequest;
    // The resource roles delivered to the target user for this assignment. Read-only. Nullable.
    accessPackageAssignmentResourceRoles []AccessPackageAssignmentResourceRole;
    // The identifier of the access package. Read-only.
    accessPackageId *string;
    // The identifier of the access package assignment policy. Read-only.
    assignmentPolicyId *string;
    // The state of the access package assignment. Possible values are Delivering, Delivered, or Expired. Read-only. Supports $filter (eq).
    assignmentState *string;
    // More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
    assignmentStatus *string;
    // The identifier of the catalog containing the access package. Read-only.
    catalogId *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    expiredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Indicates whether the access package assignment is extended. Read-only.
    isExtended *bool;
    // When the access assignment is to be in place. Read-only.
    schedule *RequestSchedule;
    // The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
    target *AccessPackageSubject;
    // The ID of the subject with the assignment. Read-only.
    targetId *string;
}
// NewAccessPackageAssignment instantiates a new accessPackageAssignment and sets the default values.
func NewAccessPackageAssignment()(*AccessPackageAssignment) {
    m := &AccessPackageAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessPackage gets the accessPackage property value. Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageAssignment) GetAccessPackage()(*AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackage
    }
}
// GetAccessPackageAssignmentPolicy gets the accessPackageAssignmentPolicy property value. Read-only. Nullable. Supports $filter (eq) on the id property
func (m *AccessPackageAssignment) GetAccessPackageAssignmentPolicy()(*AccessPackageAssignmentPolicy) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentPolicy
    }
}
// GetAccessPackageAssignmentRequests gets the accessPackageAssignmentRequests property value. 
func (m *AccessPackageAssignment) GetAccessPackageAssignmentRequests()([]AccessPackageAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentRequests
    }
}
// GetAccessPackageAssignmentResourceRoles gets the accessPackageAssignmentResourceRoles property value. The resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignment) GetAccessPackageAssignmentResourceRoles()([]AccessPackageAssignmentResourceRole) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentResourceRoles
    }
}
// GetAccessPackageId gets the accessPackageId property value. The identifier of the access package. Read-only.
func (m *AccessPackageAssignment) GetAccessPackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageId
    }
}
// GetAssignmentPolicyId gets the assignmentPolicyId property value. The identifier of the access package assignment policy. Read-only.
func (m *AccessPackageAssignment) GetAssignmentPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentPolicyId
    }
}
// GetAssignmentState gets the assignmentState property value. The state of the access package assignment. Possible values are Delivering, Delivered, or Expired. Read-only. Supports $filter (eq).
func (m *AccessPackageAssignment) GetAssignmentState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentState
    }
}
// GetAssignmentStatus gets the assignmentStatus property value. More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
func (m *AccessPackageAssignment) GetAssignmentStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentStatus
    }
}
// GetCatalogId gets the catalogId property value. The identifier of the catalog containing the access package. Read-only.
func (m *AccessPackageAssignment) GetCatalogId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogId
    }
}
// GetExpiredDateTime gets the expiredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageAssignment) GetExpiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expiredDateTime
    }
}
// GetIsExtended gets the isExtended property value. Indicates whether the access package assignment is extended. Read-only.
func (m *AccessPackageAssignment) GetIsExtended()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isExtended
    }
}
// GetSchedule gets the schedule property value. When the access assignment is to be in place. Read-only.
func (m *AccessPackageAssignment) GetSchedule()(*RequestSchedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// GetTarget gets the target property value. The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
func (m *AccessPackageAssignment) GetTarget()(*AccessPackageSubject) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetTargetId gets the targetId property value. The ID of the subject with the assignment. Read-only.
func (m *AccessPackageAssignment) GetTargetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackage(val.(*AccessPackage))
        }
        return nil
    }
    res["accessPackageAssignmentPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignmentPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageAssignmentPolicy(val.(*AccessPackageAssignmentPolicy))
        }
        return nil
    }
    res["accessPackageAssignmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignmentRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageAssignmentRequest))
            }
            m.SetAccessPackageAssignmentRequests(res)
        }
        return nil
    }
    res["accessPackageAssignmentResourceRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignmentResourceRole() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentResourceRole, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageAssignmentResourceRole))
            }
            m.SetAccessPackageAssignmentResourceRoles(res)
        }
        return nil
    }
    res["accessPackageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageId(val)
        }
        return nil
    }
    res["assignmentPolicyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentPolicyId(val)
        }
        return nil
    }
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
    res["assignmentStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentStatus(val)
        }
        return nil
    }
    res["catalogId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogId(val)
        }
        return nil
    }
    res["expiredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiredDateTime(val)
        }
        return nil
    }
    res["isExtended"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExtended(val)
        }
        return nil
    }
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRequestSchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(*RequestSchedule))
        }
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageSubject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(*AccessPackageSubject))
        }
        return nil
    }
    res["targetId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetId(val)
        }
        return nil
    }
    return res
}
func (m *AccessPackageAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetAccessPackageAssignmentRequests() != nil {
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
    if m.GetAccessPackageAssignmentResourceRoles() != nil {
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
// SetAccessPackage sets the accessPackage property value. Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageAssignment) SetAccessPackage(value *AccessPackage)() {
    if m != nil {
        m.accessPackage = value
    }
}
// SetAccessPackageAssignmentPolicy sets the accessPackageAssignmentPolicy property value. Read-only. Nullable. Supports $filter (eq) on the id property
func (m *AccessPackageAssignment) SetAccessPackageAssignmentPolicy(value *AccessPackageAssignmentPolicy)() {
    if m != nil {
        m.accessPackageAssignmentPolicy = value
    }
}
// SetAccessPackageAssignmentRequests sets the accessPackageAssignmentRequests property value. 
func (m *AccessPackageAssignment) SetAccessPackageAssignmentRequests(value []AccessPackageAssignmentRequest)() {
    if m != nil {
        m.accessPackageAssignmentRequests = value
    }
}
// SetAccessPackageAssignmentResourceRoles sets the accessPackageAssignmentResourceRoles property value. The resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignment) SetAccessPackageAssignmentResourceRoles(value []AccessPackageAssignmentResourceRole)() {
    if m != nil {
        m.accessPackageAssignmentResourceRoles = value
    }
}
// SetAccessPackageId sets the accessPackageId property value. The identifier of the access package. Read-only.
func (m *AccessPackageAssignment) SetAccessPackageId(value *string)() {
    if m != nil {
        m.accessPackageId = value
    }
}
// SetAssignmentPolicyId sets the assignmentPolicyId property value. The identifier of the access package assignment policy. Read-only.
func (m *AccessPackageAssignment) SetAssignmentPolicyId(value *string)() {
    if m != nil {
        m.assignmentPolicyId = value
    }
}
// SetAssignmentState sets the assignmentState property value. The state of the access package assignment. Possible values are Delivering, Delivered, or Expired. Read-only. Supports $filter (eq).
func (m *AccessPackageAssignment) SetAssignmentState(value *string)() {
    if m != nil {
        m.assignmentState = value
    }
}
// SetAssignmentStatus sets the assignmentStatus property value. More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
func (m *AccessPackageAssignment) SetAssignmentStatus(value *string)() {
    if m != nil {
        m.assignmentStatus = value
    }
}
// SetCatalogId sets the catalogId property value. The identifier of the catalog containing the access package. Read-only.
func (m *AccessPackageAssignment) SetCatalogId(value *string)() {
    if m != nil {
        m.catalogId = value
    }
}
// SetExpiredDateTime sets the expiredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageAssignment) SetExpiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expiredDateTime = value
    }
}
// SetIsExtended sets the isExtended property value. Indicates whether the access package assignment is extended. Read-only.
func (m *AccessPackageAssignment) SetIsExtended(value *bool)() {
    if m != nil {
        m.isExtended = value
    }
}
// SetSchedule sets the schedule property value. When the access assignment is to be in place. Read-only.
func (m *AccessPackageAssignment) SetSchedule(value *RequestSchedule)() {
    if m != nil {
        m.schedule = value
    }
}
// SetTarget sets the target property value. The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
func (m *AccessPackageAssignment) SetTarget(value *AccessPackageSubject)() {
    if m != nil {
        m.target = value
    }
}
// SetTargetId sets the targetId property value. The ID of the subject with the assignment. Read-only.
func (m *AccessPackageAssignment) SetTargetId(value *string)() {
    if m != nil {
        m.targetId = value
    }
}
