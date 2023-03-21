package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignment 
type AccessPackageAssignment struct {
    Entity
}
// NewAccessPackageAssignment instantiates a new accessPackageAssignment and sets the default values.
func NewAccessPackageAssignment()(*AccessPackageAssignment) {
    m := &AccessPackageAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignment(), nil
}
// GetAccessPackage gets the accessPackage property value. Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageAssignment) GetAccessPackage()(AccessPackageable) {
    val, err := m.GetBackingStore().Get("accessPackage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageable)
    }
    return nil
}
// GetAccessPackageAssignmentPolicy gets the accessPackageAssignmentPolicy property value. Read-only. Nullable. Supports $filter (eq) on the id property
func (m *AccessPackageAssignment) GetAccessPackageAssignmentPolicy()(AccessPackageAssignmentPolicyable) {
    val, err := m.GetBackingStore().Get("accessPackageAssignmentPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageAssignmentPolicyable)
    }
    return nil
}
// GetAccessPackageAssignmentRequests gets the accessPackageAssignmentRequests property value. The accessPackageAssignmentRequests property
func (m *AccessPackageAssignment) GetAccessPackageAssignmentRequests()([]AccessPackageAssignmentRequestable) {
    val, err := m.GetBackingStore().Get("accessPackageAssignmentRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageAssignmentRequestable)
    }
    return nil
}
// GetAccessPackageAssignmentResourceRoles gets the accessPackageAssignmentResourceRoles property value. The resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignment) GetAccessPackageAssignmentResourceRoles()([]AccessPackageAssignmentResourceRoleable) {
    val, err := m.GetBackingStore().Get("accessPackageAssignmentResourceRoles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageAssignmentResourceRoleable)
    }
    return nil
}
// GetAccessPackageId gets the accessPackageId property value. The identifier of the access package. Read-only.
func (m *AccessPackageAssignment) GetAccessPackageId()(*string) {
    val, err := m.GetBackingStore().Get("accessPackageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignmentPolicyId gets the assignmentPolicyId property value. The identifier of the access package assignment policy. Read-only.
func (m *AccessPackageAssignment) GetAssignmentPolicyId()(*string) {
    val, err := m.GetBackingStore().Get("assignmentPolicyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignmentState gets the assignmentState property value. The state of the access package assignment. Possible values are Delivering, Delivered, or Expired. Read-only. Supports $filter (eq).
func (m *AccessPackageAssignment) GetAssignmentState()(*string) {
    val, err := m.GetBackingStore().Get("assignmentState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignmentStatus gets the assignmentStatus property value. More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
func (m *AccessPackageAssignment) GetAssignmentStatus()(*string) {
    val, err := m.GetBackingStore().Get("assignmentStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCatalogId gets the catalogId property value. The identifier of the catalog containing the access package. Read-only.
func (m *AccessPackageAssignment) GetCatalogId()(*string) {
    val, err := m.GetBackingStore().Get("catalogId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomExtensionCalloutInstances gets the customExtensionCalloutInstances property value. Information about all the custom extension calls that were made during the access package assignment workflow.
func (m *AccessPackageAssignment) GetCustomExtensionCalloutInstances()([]CustomExtensionCalloutInstanceable) {
    val, err := m.GetBackingStore().Get("customExtensionCalloutInstances")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomExtensionCalloutInstanceable)
    }
    return nil
}
// GetExpiredDateTime gets the expiredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignment) GetExpiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expiredDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackage(val.(AccessPackageable))
        }
        return nil
    }
    res["accessPackageAssignmentPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAssignmentPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageAssignmentPolicy(val.(AccessPackageAssignmentPolicyable))
        }
        return nil
    }
    res["accessPackageAssignmentRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAssignmentRequestable)
            }
            m.SetAccessPackageAssignmentRequests(res)
        }
        return nil
    }
    res["accessPackageAssignmentResourceRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentResourceRoleable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAssignmentResourceRoleable)
            }
            m.SetAccessPackageAssignmentResourceRoles(res)
        }
        return nil
    }
    res["accessPackageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageId(val)
        }
        return nil
    }
    res["assignmentPolicyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentPolicyId(val)
        }
        return nil
    }
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
    res["assignmentStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentStatus(val)
        }
        return nil
    }
    res["catalogId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogId(val)
        }
        return nil
    }
    res["customExtensionCalloutInstances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomExtensionCalloutInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomExtensionCalloutInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(CustomExtensionCalloutInstanceable)
            }
            m.SetCustomExtensionCalloutInstances(res)
        }
        return nil
    }
    res["expiredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiredDateTime(val)
        }
        return nil
    }
    res["isExtended"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExtended(val)
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(RequestScheduleable))
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(AccessPackageSubjectable))
        }
        return nil
    }
    res["targetId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetIsExtended gets the isExtended property value. Indicates whether the access package assignment is extended. Read-only.
func (m *AccessPackageAssignment) GetIsExtended()(*bool) {
    val, err := m.GetBackingStore().Get("isExtended")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSchedule gets the schedule property value. When the access assignment is to be in place. Read-only.
func (m *AccessPackageAssignment) GetSchedule()(RequestScheduleable) {
    val, err := m.GetBackingStore().Get("schedule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RequestScheduleable)
    }
    return nil
}
// GetTarget gets the target property value. The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
func (m *AccessPackageAssignment) GetTarget()(AccessPackageSubjectable) {
    val, err := m.GetBackingStore().Get("target")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageSubjectable)
    }
    return nil
}
// GetTargetId gets the targetId property value. The ID of the subject with the assignment. Read-only.
func (m *AccessPackageAssignment) GetTargetId()(*string) {
    val, err := m.GetBackingStore().Get("targetId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessPackageAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageAssignmentRequests()))
        for i, v := range m.GetAccessPackageAssignmentRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageAssignmentResourceRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageAssignmentResourceRoles()))
        for i, v := range m.GetAccessPackageAssignmentResourceRoles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetCustomExtensionCalloutInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomExtensionCalloutInstances()))
        for i, v := range m.GetCustomExtensionCalloutInstances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customExtensionCalloutInstances", cast)
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
func (m *AccessPackageAssignment) SetAccessPackage(value AccessPackageable)() {
    err := m.GetBackingStore().Set("accessPackage", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageAssignmentPolicy sets the accessPackageAssignmentPolicy property value. Read-only. Nullable. Supports $filter (eq) on the id property
func (m *AccessPackageAssignment) SetAccessPackageAssignmentPolicy(value AccessPackageAssignmentPolicyable)() {
    err := m.GetBackingStore().Set("accessPackageAssignmentPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageAssignmentRequests sets the accessPackageAssignmentRequests property value. The accessPackageAssignmentRequests property
func (m *AccessPackageAssignment) SetAccessPackageAssignmentRequests(value []AccessPackageAssignmentRequestable)() {
    err := m.GetBackingStore().Set("accessPackageAssignmentRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageAssignmentResourceRoles sets the accessPackageAssignmentResourceRoles property value. The resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignment) SetAccessPackageAssignmentResourceRoles(value []AccessPackageAssignmentResourceRoleable)() {
    err := m.GetBackingStore().Set("accessPackageAssignmentResourceRoles", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageId sets the accessPackageId property value. The identifier of the access package. Read-only.
func (m *AccessPackageAssignment) SetAccessPackageId(value *string)() {
    err := m.GetBackingStore().Set("accessPackageId", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentPolicyId sets the assignmentPolicyId property value. The identifier of the access package assignment policy. Read-only.
func (m *AccessPackageAssignment) SetAssignmentPolicyId(value *string)() {
    err := m.GetBackingStore().Set("assignmentPolicyId", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentState sets the assignmentState property value. The state of the access package assignment. Possible values are Delivering, Delivered, or Expired. Read-only. Supports $filter (eq).
func (m *AccessPackageAssignment) SetAssignmentState(value *string)() {
    err := m.GetBackingStore().Set("assignmentState", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentStatus sets the assignmentStatus property value. More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
func (m *AccessPackageAssignment) SetAssignmentStatus(value *string)() {
    err := m.GetBackingStore().Set("assignmentStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetCatalogId sets the catalogId property value. The identifier of the catalog containing the access package. Read-only.
func (m *AccessPackageAssignment) SetCatalogId(value *string)() {
    err := m.GetBackingStore().Set("catalogId", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomExtensionCalloutInstances sets the customExtensionCalloutInstances property value. Information about all the custom extension calls that were made during the access package assignment workflow.
func (m *AccessPackageAssignment) SetCustomExtensionCalloutInstances(value []CustomExtensionCalloutInstanceable)() {
    err := m.GetBackingStore().Set("customExtensionCalloutInstances", value)
    if err != nil {
        panic(err)
    }
}
// SetExpiredDateTime sets the expiredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignment) SetExpiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expiredDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsExtended sets the isExtended property value. Indicates whether the access package assignment is extended. Read-only.
func (m *AccessPackageAssignment) SetIsExtended(value *bool)() {
    err := m.GetBackingStore().Set("isExtended", value)
    if err != nil {
        panic(err)
    }
}
// SetSchedule sets the schedule property value. When the access assignment is to be in place. Read-only.
func (m *AccessPackageAssignment) SetSchedule(value RequestScheduleable)() {
    err := m.GetBackingStore().Set("schedule", value)
    if err != nil {
        panic(err)
    }
}
// SetTarget sets the target property value. The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
func (m *AccessPackageAssignment) SetTarget(value AccessPackageSubjectable)() {
    err := m.GetBackingStore().Set("target", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetId sets the targetId property value. The ID of the subject with the assignment. Read-only.
func (m *AccessPackageAssignment) SetTargetId(value *string)() {
    err := m.GetBackingStore().Set("targetId", value)
    if err != nil {
        panic(err)
    }
}
// AccessPackageAssignmentable 
type AccessPackageAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackage()(AccessPackageable)
    GetAccessPackageAssignmentPolicy()(AccessPackageAssignmentPolicyable)
    GetAccessPackageAssignmentRequests()([]AccessPackageAssignmentRequestable)
    GetAccessPackageAssignmentResourceRoles()([]AccessPackageAssignmentResourceRoleable)
    GetAccessPackageId()(*string)
    GetAssignmentPolicyId()(*string)
    GetAssignmentState()(*string)
    GetAssignmentStatus()(*string)
    GetCatalogId()(*string)
    GetCustomExtensionCalloutInstances()([]CustomExtensionCalloutInstanceable)
    GetExpiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsExtended()(*bool)
    GetSchedule()(RequestScheduleable)
    GetTarget()(AccessPackageSubjectable)
    GetTargetId()(*string)
    SetAccessPackage(value AccessPackageable)()
    SetAccessPackageAssignmentPolicy(value AccessPackageAssignmentPolicyable)()
    SetAccessPackageAssignmentRequests(value []AccessPackageAssignmentRequestable)()
    SetAccessPackageAssignmentResourceRoles(value []AccessPackageAssignmentResourceRoleable)()
    SetAccessPackageId(value *string)()
    SetAssignmentPolicyId(value *string)()
    SetAssignmentState(value *string)()
    SetAssignmentStatus(value *string)()
    SetCatalogId(value *string)()
    SetCustomExtensionCalloutInstances(value []CustomExtensionCalloutInstanceable)()
    SetExpiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsExtended(value *bool)()
    SetSchedule(value RequestScheduleable)()
    SetTarget(value AccessPackageSubjectable)()
    SetTargetId(value *string)()
}
