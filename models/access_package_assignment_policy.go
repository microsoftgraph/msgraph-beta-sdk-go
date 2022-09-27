package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentPolicy provides operations to manage the collection of accessReviewDecision entities.
type AccessPackageAssignmentPolicy struct {
    Entity
    // The access package with this policy. Read-only. Nullable. Supports $expand.
    accessPackage AccessPackageable
    // The accessPackageCatalog property
    accessPackageCatalog AccessPackageCatalogable
    // Identifier of the access package.
    accessPackageId *string
    // Who must review, and how often, the assignments to the access package from this policy. This property is null if reviews are not required.
    accessReviewSettings AssignmentReviewSettingsable
    // Indicates whether a user can extend the access package assignment duration after approval.
    canExtend *bool
    // The createdBy property
    createdBy *string
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
    customExtensionHandlers []CustomExtensionHandlerable
    // The description of the policy.
    description *string
    // The display name of the policy. Supports $filter (eq).
    displayName *string
    // The number of days in which assignments from this policy last until they are expired.
    durationInDays *int32
    // The expiration date for assignments created in this policy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The modifiedBy property
    modifiedBy *string
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Questions that are posed to the  requestor.
    questions []AccessPackageQuestionable
    // Who must approve requests for access package in this policy.
    requestApprovalSettings ApprovalSettingsable
    // Who can request this access package from this policy.
    requestorSettings RequestorSettingsable
}
// NewAccessPackageAssignmentPolicy instantiates a new accessPackageAssignmentPolicy and sets the default values.
func NewAccessPackageAssignmentPolicy()(*AccessPackageAssignmentPolicy) {
    m := &AccessPackageAssignmentPolicy{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageAssignmentPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageAssignmentPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentPolicy(), nil
}
// GetAccessPackage gets the accessPackage property value. The access package with this policy. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentPolicy) GetAccessPackage()(AccessPackageable) {
    return m.accessPackage
}
// GetAccessPackageCatalog gets the accessPackageCatalog property value. The accessPackageCatalog property
func (m *AccessPackageAssignmentPolicy) GetAccessPackageCatalog()(AccessPackageCatalogable) {
    return m.accessPackageCatalog
}
// GetAccessPackageId gets the accessPackageId property value. Identifier of the access package.
func (m *AccessPackageAssignmentPolicy) GetAccessPackageId()(*string) {
    return m.accessPackageId
}
// GetAccessReviewSettings gets the accessReviewSettings property value. Who must review, and how often, the assignments to the access package from this policy. This property is null if reviews are not required.
func (m *AccessPackageAssignmentPolicy) GetAccessReviewSettings()(AssignmentReviewSettingsable) {
    return m.accessReviewSettings
}
// GetCanExtend gets the canExtend property value. Indicates whether a user can extend the access package assignment duration after approval.
func (m *AccessPackageAssignmentPolicy) GetCanExtend()(*bool) {
    return m.canExtend
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *AccessPackageAssignmentPolicy) GetCreatedBy()(*string) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustomExtensionHandlers gets the customExtensionHandlers property value. The collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
func (m *AccessPackageAssignmentPolicy) GetCustomExtensionHandlers()([]CustomExtensionHandlerable) {
    return m.customExtensionHandlers
}
// GetDescription gets the description property value. The description of the policy.
func (m *AccessPackageAssignmentPolicy) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the policy. Supports $filter (eq).
func (m *AccessPackageAssignmentPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetDurationInDays gets the durationInDays property value. The number of days in which assignments from this policy last until they are expired.
func (m *AccessPackageAssignmentPolicy) GetDurationInDays()(*int32) {
    return m.durationInDays
}
// GetExpirationDateTime gets the expirationDateTime property value. The expiration date for assignments created in this policy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessPackageFromDiscriminatorValue , m.SetAccessPackage)
    res["accessPackageCatalog"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessPackageCatalogFromDiscriminatorValue , m.SetAccessPackageCatalog)
    res["accessPackageId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAccessPackageId)
    res["accessReviewSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAssignmentReviewSettingsFromDiscriminatorValue , m.SetAccessReviewSettings)
    res["canExtend"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCanExtend)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["customExtensionHandlers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomExtensionHandlerFromDiscriminatorValue , m.SetCustomExtensionHandlers)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["durationInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDurationInDays)
    res["expirationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetExpirationDateTime)
    res["modifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModifiedBy)
    res["modifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetModifiedDateTime)
    res["questions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageQuestionFromDiscriminatorValue , m.SetQuestions)
    res["requestApprovalSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateApprovalSettingsFromDiscriminatorValue , m.SetRequestApprovalSettings)
    res["requestorSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRequestorSettingsFromDiscriminatorValue , m.SetRequestorSettings)
    return res
}
// GetModifiedBy gets the modifiedBy property value. The modifiedBy property
func (m *AccessPackageAssignmentPolicy) GetModifiedBy()(*string) {
    return m.modifiedBy
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetQuestions gets the questions property value. Questions that are posed to the  requestor.
func (m *AccessPackageAssignmentPolicy) GetQuestions()([]AccessPackageQuestionable) {
    return m.questions
}
// GetRequestApprovalSettings gets the requestApprovalSettings property value. Who must approve requests for access package in this policy.
func (m *AccessPackageAssignmentPolicy) GetRequestApprovalSettings()(ApprovalSettingsable) {
    return m.requestApprovalSettings
}
// GetRequestorSettings gets the requestorSettings property value. Who can request this access package from this policy.
func (m *AccessPackageAssignmentPolicy) GetRequestorSettings()(RequestorSettingsable) {
    return m.requestorSettings
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("accessPackageCatalog", m.GetAccessPackageCatalog())
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
        err = writer.WriteObjectValue("accessReviewSettings", m.GetAccessReviewSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("canExtend", m.GetCanExtend())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetCustomExtensionHandlers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomExtensionHandlers())
        err = writer.WriteCollectionOfObjectValues("customExtensionHandlers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("durationInDays", m.GetDurationInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("modifiedBy", m.GetModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetQuestions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetQuestions())
        err = writer.WriteCollectionOfObjectValues("questions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestApprovalSettings", m.GetRequestApprovalSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestorSettings", m.GetRequestorSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackage sets the accessPackage property value. The access package with this policy. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentPolicy) SetAccessPackage(value AccessPackageable)() {
    m.accessPackage = value
}
// SetAccessPackageCatalog sets the accessPackageCatalog property value. The accessPackageCatalog property
func (m *AccessPackageAssignmentPolicy) SetAccessPackageCatalog(value AccessPackageCatalogable)() {
    m.accessPackageCatalog = value
}
// SetAccessPackageId sets the accessPackageId property value. Identifier of the access package.
func (m *AccessPackageAssignmentPolicy) SetAccessPackageId(value *string)() {
    m.accessPackageId = value
}
// SetAccessReviewSettings sets the accessReviewSettings property value. Who must review, and how often, the assignments to the access package from this policy. This property is null if reviews are not required.
func (m *AccessPackageAssignmentPolicy) SetAccessReviewSettings(value AssignmentReviewSettingsable)() {
    m.accessReviewSettings = value
}
// SetCanExtend sets the canExtend property value. Indicates whether a user can extend the access package assignment duration after approval.
func (m *AccessPackageAssignmentPolicy) SetCanExtend(value *bool)() {
    m.canExtend = value
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *AccessPackageAssignmentPolicy) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustomExtensionHandlers sets the customExtensionHandlers property value. The collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
func (m *AccessPackageAssignmentPolicy) SetCustomExtensionHandlers(value []CustomExtensionHandlerable)() {
    m.customExtensionHandlers = value
}
// SetDescription sets the description property value. The description of the policy.
func (m *AccessPackageAssignmentPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the policy. Supports $filter (eq).
func (m *AccessPackageAssignmentPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDurationInDays sets the durationInDays property value. The number of days in which assignments from this policy last until they are expired.
func (m *AccessPackageAssignmentPolicy) SetDurationInDays(value *int32)() {
    m.durationInDays = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The expiration date for assignments created in this policy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetModifiedBy sets the modifiedBy property value. The modifiedBy property
func (m *AccessPackageAssignmentPolicy) SetModifiedBy(value *string)() {
    m.modifiedBy = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetQuestions sets the questions property value. Questions that are posed to the  requestor.
func (m *AccessPackageAssignmentPolicy) SetQuestions(value []AccessPackageQuestionable)() {
    m.questions = value
}
// SetRequestApprovalSettings sets the requestApprovalSettings property value. Who must approve requests for access package in this policy.
func (m *AccessPackageAssignmentPolicy) SetRequestApprovalSettings(value ApprovalSettingsable)() {
    m.requestApprovalSettings = value
}
// SetRequestorSettings sets the requestorSettings property value. Who can request this access package from this policy.
func (m *AccessPackageAssignmentPolicy) SetRequestorSettings(value RequestorSettingsable)() {
    m.requestorSettings = value
}
