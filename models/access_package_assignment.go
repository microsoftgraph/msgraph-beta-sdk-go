package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignment 
type AccessPackageAssignment struct {
    Entity
    // Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
    accessPackage AccessPackageable
    // Read-only. Nullable. Supports $filter (eq) on the id property
    accessPackageAssignmentPolicy AccessPackageAssignmentPolicyable
    // The accessPackageAssignmentRequests property
    accessPackageAssignmentRequests []AccessPackageAssignmentRequestable
    // The resource roles delivered to the target user for this assignment. Read-only. Nullable.
    accessPackageAssignmentResourceRoles []AccessPackageAssignmentResourceRoleable
    // The identifier of the access package. Read-only.
    accessPackageId *string
    // The identifier of the access package assignment policy. Read-only.
    assignmentPolicyId *string
    // The state of the access package assignment. Possible values are Delivering, Delivered, or Expired. Read-only. Supports $filter (eq).
    assignmentState *string
    // More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
    assignmentStatus *string
    // The identifier of the catalog containing the access package. Read-only.
    catalogId *string
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    expiredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates whether the access package assignment is extended. Read-only.
    isExtended *bool
    // When the access assignment is to be in place. Read-only.
    schedule RequestScheduleable
    // The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
    target AccessPackageSubjectable
    // The ID of the subject with the assignment. Read-only.
    targetId *string
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
    return m.accessPackage
}
// GetAccessPackageAssignmentPolicy gets the accessPackageAssignmentPolicy property value. Read-only. Nullable. Supports $filter (eq) on the id property
func (m *AccessPackageAssignment) GetAccessPackageAssignmentPolicy()(AccessPackageAssignmentPolicyable) {
    return m.accessPackageAssignmentPolicy
}
// GetAccessPackageAssignmentRequests gets the accessPackageAssignmentRequests property value. The accessPackageAssignmentRequests property
func (m *AccessPackageAssignment) GetAccessPackageAssignmentRequests()([]AccessPackageAssignmentRequestable) {
    return m.accessPackageAssignmentRequests
}
// GetAccessPackageAssignmentResourceRoles gets the accessPackageAssignmentResourceRoles property value. The resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignment) GetAccessPackageAssignmentResourceRoles()([]AccessPackageAssignmentResourceRoleable) {
    return m.accessPackageAssignmentResourceRoles
}
// GetAccessPackageId gets the accessPackageId property value. The identifier of the access package. Read-only.
func (m *AccessPackageAssignment) GetAccessPackageId()(*string) {
    return m.accessPackageId
}
// GetAssignmentPolicyId gets the assignmentPolicyId property value. The identifier of the access package assignment policy. Read-only.
func (m *AccessPackageAssignment) GetAssignmentPolicyId()(*string) {
    return m.assignmentPolicyId
}
// GetAssignmentState gets the assignmentState property value. The state of the access package assignment. Possible values are Delivering, Delivered, or Expired. Read-only. Supports $filter (eq).
func (m *AccessPackageAssignment) GetAssignmentState()(*string) {
    return m.assignmentState
}
// GetAssignmentStatus gets the assignmentStatus property value. More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
func (m *AccessPackageAssignment) GetAssignmentStatus()(*string) {
    return m.assignmentStatus
}
// GetCatalogId gets the catalogId property value. The identifier of the catalog containing the access package. Read-only.
func (m *AccessPackageAssignment) GetCatalogId()(*string) {
    return m.catalogId
}
// GetExpiredDateTime gets the expiredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignment) GetExpiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expiredDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessPackageFromDiscriminatorValue , m.SetAccessPackage)
    res["accessPackageAssignmentPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessPackageAssignmentPolicyFromDiscriminatorValue , m.SetAccessPackageAssignmentPolicy)
    res["accessPackageAssignmentRequests"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageAssignmentRequestFromDiscriminatorValue , m.SetAccessPackageAssignmentRequests)
    res["accessPackageAssignmentResourceRoles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue , m.SetAccessPackageAssignmentResourceRoles)
    res["accessPackageId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAccessPackageId)
    res["assignmentPolicyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAssignmentPolicyId)
    res["assignmentState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAssignmentState)
    res["assignmentStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAssignmentStatus)
    res["catalogId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCatalogId)
    res["expiredDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetExpiredDateTime)
    res["isExtended"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsExtended)
    res["schedule"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRequestScheduleFromDiscriminatorValue , m.SetSchedule)
    res["target"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessPackageSubjectFromDiscriminatorValue , m.SetTarget)
    res["targetId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTargetId)
    return res
}
// GetIsExtended gets the isExtended property value. Indicates whether the access package assignment is extended. Read-only.
func (m *AccessPackageAssignment) GetIsExtended()(*bool) {
    return m.isExtended
}
// GetSchedule gets the schedule property value. When the access assignment is to be in place. Read-only.
func (m *AccessPackageAssignment) GetSchedule()(RequestScheduleable) {
    return m.schedule
}
// GetTarget gets the target property value. The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
func (m *AccessPackageAssignment) GetTarget()(AccessPackageSubjectable) {
    return m.target
}
// GetTargetId gets the targetId property value. The ID of the subject with the assignment. Read-only.
func (m *AccessPackageAssignment) GetTargetId()(*string) {
    return m.targetId
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAccessPackageAssignmentRequests())
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageAssignmentResourceRoles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAccessPackageAssignmentResourceRoles())
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
func (m *AccessPackageAssignment) SetAccessPackage(value AccessPackageable)() {
    m.accessPackage = value
}
// SetAccessPackageAssignmentPolicy sets the accessPackageAssignmentPolicy property value. Read-only. Nullable. Supports $filter (eq) on the id property
func (m *AccessPackageAssignment) SetAccessPackageAssignmentPolicy(value AccessPackageAssignmentPolicyable)() {
    m.accessPackageAssignmentPolicy = value
}
// SetAccessPackageAssignmentRequests sets the accessPackageAssignmentRequests property value. The accessPackageAssignmentRequests property
func (m *AccessPackageAssignment) SetAccessPackageAssignmentRequests(value []AccessPackageAssignmentRequestable)() {
    m.accessPackageAssignmentRequests = value
}
// SetAccessPackageAssignmentResourceRoles sets the accessPackageAssignmentResourceRoles property value. The resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignment) SetAccessPackageAssignmentResourceRoles(value []AccessPackageAssignmentResourceRoleable)() {
    m.accessPackageAssignmentResourceRoles = value
}
// SetAccessPackageId sets the accessPackageId property value. The identifier of the access package. Read-only.
func (m *AccessPackageAssignment) SetAccessPackageId(value *string)() {
    m.accessPackageId = value
}
// SetAssignmentPolicyId sets the assignmentPolicyId property value. The identifier of the access package assignment policy. Read-only.
func (m *AccessPackageAssignment) SetAssignmentPolicyId(value *string)() {
    m.assignmentPolicyId = value
}
// SetAssignmentState sets the assignmentState property value. The state of the access package assignment. Possible values are Delivering, Delivered, or Expired. Read-only. Supports $filter (eq).
func (m *AccessPackageAssignment) SetAssignmentState(value *string)() {
    m.assignmentState = value
}
// SetAssignmentStatus sets the assignmentStatus property value. More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
func (m *AccessPackageAssignment) SetAssignmentStatus(value *string)() {
    m.assignmentStatus = value
}
// SetCatalogId sets the catalogId property value. The identifier of the catalog containing the access package. Read-only.
func (m *AccessPackageAssignment) SetCatalogId(value *string)() {
    m.catalogId = value
}
// SetExpiredDateTime sets the expiredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignment) SetExpiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiredDateTime = value
}
// SetIsExtended sets the isExtended property value. Indicates whether the access package assignment is extended. Read-only.
func (m *AccessPackageAssignment) SetIsExtended(value *bool)() {
    m.isExtended = value
}
// SetSchedule sets the schedule property value. When the access assignment is to be in place. Read-only.
func (m *AccessPackageAssignment) SetSchedule(value RequestScheduleable)() {
    m.schedule = value
}
// SetTarget sets the target property value. The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
func (m *AccessPackageAssignment) SetTarget(value AccessPackageSubjectable)() {
    m.target = value
}
// SetTargetId sets the targetId property value. The ID of the subject with the assignment. Read-only.
func (m *AccessPackageAssignment) SetTargetId(value *string)() {
    m.targetId = value
}
