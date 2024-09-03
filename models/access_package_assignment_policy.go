package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageAssignmentPolicy struct {
    Entity
}
// NewAccessPackageAssignmentPolicy instantiates a new AccessPackageAssignmentPolicy and sets the default values.
func NewAccessPackageAssignmentPolicy()(*AccessPackageAssignmentPolicy) {
    m := &AccessPackageAssignmentPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageAssignmentPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageAssignmentPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentPolicy(), nil
}
// GetAccessPackage gets the accessPackage property value. The access package with this policy. Read-only. Nullable. Supports $expand.
// returns a AccessPackageable when successful
func (m *AccessPackageAssignmentPolicy) GetAccessPackage()(AccessPackageable) {
    val, err := m.GetBackingStore().Get("accessPackage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageable)
    }
    return nil
}
// GetAccessPackageCatalog gets the accessPackageCatalog property value. The accessPackageCatalog property
// returns a AccessPackageCatalogable when successful
func (m *AccessPackageAssignmentPolicy) GetAccessPackageCatalog()(AccessPackageCatalogable) {
    val, err := m.GetBackingStore().Get("accessPackageCatalog")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageCatalogable)
    }
    return nil
}
// GetAccessPackageId gets the accessPackageId property value. Identifier of the access package.
// returns a *string when successful
func (m *AccessPackageAssignmentPolicy) GetAccessPackageId()(*string) {
    val, err := m.GetBackingStore().Get("accessPackageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAccessPackageNotificationSettings gets the accessPackageNotificationSettings property value. The accessPackageNotificationSettings property
// returns a AccessPackageNotificationSettingsable when successful
func (m *AccessPackageAssignmentPolicy) GetAccessPackageNotificationSettings()(AccessPackageNotificationSettingsable) {
    val, err := m.GetBackingStore().Get("accessPackageNotificationSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageNotificationSettingsable)
    }
    return nil
}
// GetAccessReviewSettings gets the accessReviewSettings property value. Who must review, and how often, the assignments to the access package from this policy. This property is null if reviews aren't required.
// returns a AssignmentReviewSettingsable when successful
func (m *AccessPackageAssignmentPolicy) GetAccessReviewSettings()(AssignmentReviewSettingsable) {
    val, err := m.GetBackingStore().Get("accessReviewSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AssignmentReviewSettingsable)
    }
    return nil
}
// GetCanExtend gets the canExtend property value. Indicates whether a user can extend the access package assignment duration after approval.
// returns a *bool when successful
func (m *AccessPackageAssignmentPolicy) GetCanExtend()(*bool) {
    val, err := m.GetBackingStore().Get("canExtend")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. The createdBy property
// returns a *string when successful
func (m *AccessPackageAssignmentPolicy) GetCreatedBy()(*string) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *AccessPackageAssignmentPolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCustomExtensionHandlers gets the customExtensionHandlers property value. The collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
// returns a []CustomExtensionHandlerable when successful
func (m *AccessPackageAssignmentPolicy) GetCustomExtensionHandlers()([]CustomExtensionHandlerable) {
    val, err := m.GetBackingStore().Get("customExtensionHandlers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomExtensionHandlerable)
    }
    return nil
}
// GetCustomExtensionStageSettings gets the customExtensionStageSettings property value. The collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
// returns a []CustomExtensionStageSettingable when successful
func (m *AccessPackageAssignmentPolicy) GetCustomExtensionStageSettings()([]CustomExtensionStageSettingable) {
    val, err := m.GetBackingStore().Get("customExtensionStageSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomExtensionStageSettingable)
    }
    return nil
}
// GetDescription gets the description property value. The description of the policy.
// returns a *string when successful
func (m *AccessPackageAssignmentPolicy) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the policy. Supports $filter (eq).
// returns a *string when successful
func (m *AccessPackageAssignmentPolicy) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDurationInDays gets the durationInDays property value. The number of days in which assignments from this policy last until they're expired.
// returns a *int32 when successful
func (m *AccessPackageAssignmentPolicy) GetDurationInDays()(*int32) {
    val, err := m.GetBackingStore().Get("durationInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The expiration date for assignments created in this policy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *AccessPackageAssignmentPolicy) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageAssignmentPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["accessPackageCatalog"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageCatalog(val.(AccessPackageCatalogable))
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
    res["accessPackageNotificationSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageNotificationSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageNotificationSettings(val.(AccessPackageNotificationSettingsable))
        }
        return nil
    }
    res["accessReviewSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAssignmentReviewSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessReviewSettings(val.(AssignmentReviewSettingsable))
        }
        return nil
    }
    res["canExtend"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanExtend(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["customExtensionHandlers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomExtensionHandlerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomExtensionHandlerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomExtensionHandlerable)
                }
            }
            m.SetCustomExtensionHandlers(res)
        }
        return nil
    }
    res["customExtensionStageSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomExtensionStageSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomExtensionStageSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomExtensionStageSettingable)
                }
            }
            m.SetCustomExtensionStageSettings(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["durationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInDays(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["modifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedBy(val)
        }
        return nil
    }
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["questions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageQuestionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageQuestionable)
                }
            }
            m.SetQuestions(res)
        }
        return nil
    }
    res["requestApprovalSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApprovalSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestApprovalSettings(val.(ApprovalSettingsable))
        }
        return nil
    }
    res["requestorSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestorSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestorSettings(val.(RequestorSettingsable))
        }
        return nil
    }
    res["verifiableCredentialSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVerifiableCredentialSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiableCredentialSettings(val.(VerifiableCredentialSettingsable))
        }
        return nil
    }
    return res
}
// GetModifiedBy gets the modifiedBy property value. The modifiedBy property
// returns a *string when successful
func (m *AccessPackageAssignmentPolicy) GetModifiedBy()(*string) {
    val, err := m.GetBackingStore().Get("modifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *AccessPackageAssignmentPolicy) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("modifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetQuestions gets the questions property value. Questions that are posed to the  requestor.
// returns a []AccessPackageQuestionable when successful
func (m *AccessPackageAssignmentPolicy) GetQuestions()([]AccessPackageQuestionable) {
    val, err := m.GetBackingStore().Get("questions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageQuestionable)
    }
    return nil
}
// GetRequestApprovalSettings gets the requestApprovalSettings property value. Who must approve requests for access package in this policy.
// returns a ApprovalSettingsable when successful
func (m *AccessPackageAssignmentPolicy) GetRequestApprovalSettings()(ApprovalSettingsable) {
    val, err := m.GetBackingStore().Get("requestApprovalSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ApprovalSettingsable)
    }
    return nil
}
// GetRequestorSettings gets the requestorSettings property value. Who can request this access package from this policy.
// returns a RequestorSettingsable when successful
func (m *AccessPackageAssignmentPolicy) GetRequestorSettings()(RequestorSettingsable) {
    val, err := m.GetBackingStore().Get("requestorSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RequestorSettingsable)
    }
    return nil
}
// GetVerifiableCredentialSettings gets the verifiableCredentialSettings property value. Settings for verifiable credentials set up through the Microsoft Entra Verified I D service. These settings represent the verifiable credentials that a requestor of an access package in this policy can present to be assigned the access package.
// returns a VerifiableCredentialSettingsable when successful
func (m *AccessPackageAssignmentPolicy) GetVerifiableCredentialSettings()(VerifiableCredentialSettingsable) {
    val, err := m.GetBackingStore().Get("verifiableCredentialSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VerifiableCredentialSettingsable)
    }
    return nil
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
        err = writer.WriteObjectValue("accessPackageNotificationSettings", m.GetAccessPackageNotificationSettings())
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomExtensionHandlers()))
        for i, v := range m.GetCustomExtensionHandlers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customExtensionHandlers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomExtensionStageSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomExtensionStageSettings()))
        for i, v := range m.GetCustomExtensionStageSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customExtensionStageSettings", cast)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetQuestions()))
        for i, v := range m.GetQuestions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    {
        err = writer.WriteObjectValue("verifiableCredentialSettings", m.GetVerifiableCredentialSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackage sets the accessPackage property value. The access package with this policy. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentPolicy) SetAccessPackage(value AccessPackageable)() {
    err := m.GetBackingStore().Set("accessPackage", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageCatalog sets the accessPackageCatalog property value. The accessPackageCatalog property
func (m *AccessPackageAssignmentPolicy) SetAccessPackageCatalog(value AccessPackageCatalogable)() {
    err := m.GetBackingStore().Set("accessPackageCatalog", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageId sets the accessPackageId property value. Identifier of the access package.
func (m *AccessPackageAssignmentPolicy) SetAccessPackageId(value *string)() {
    err := m.GetBackingStore().Set("accessPackageId", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageNotificationSettings sets the accessPackageNotificationSettings property value. The accessPackageNotificationSettings property
func (m *AccessPackageAssignmentPolicy) SetAccessPackageNotificationSettings(value AccessPackageNotificationSettingsable)() {
    err := m.GetBackingStore().Set("accessPackageNotificationSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessReviewSettings sets the accessReviewSettings property value. Who must review, and how often, the assignments to the access package from this policy. This property is null if reviews aren't required.
func (m *AccessPackageAssignmentPolicy) SetAccessReviewSettings(value AssignmentReviewSettingsable)() {
    err := m.GetBackingStore().Set("accessReviewSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetCanExtend sets the canExtend property value. Indicates whether a user can extend the access package assignment duration after approval.
func (m *AccessPackageAssignmentPolicy) SetCanExtend(value *bool)() {
    err := m.GetBackingStore().Set("canExtend", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *AccessPackageAssignmentPolicy) SetCreatedBy(value *string)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomExtensionHandlers sets the customExtensionHandlers property value. The collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
func (m *AccessPackageAssignmentPolicy) SetCustomExtensionHandlers(value []CustomExtensionHandlerable)() {
    err := m.GetBackingStore().Set("customExtensionHandlers", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomExtensionStageSettings sets the customExtensionStageSettings property value. The collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
func (m *AccessPackageAssignmentPolicy) SetCustomExtensionStageSettings(value []CustomExtensionStageSettingable)() {
    err := m.GetBackingStore().Set("customExtensionStageSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the policy.
func (m *AccessPackageAssignmentPolicy) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the policy. Supports $filter (eq).
func (m *AccessPackageAssignmentPolicy) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDurationInDays sets the durationInDays property value. The number of days in which assignments from this policy last until they're expired.
func (m *AccessPackageAssignmentPolicy) SetDurationInDays(value *int32)() {
    err := m.GetBackingStore().Set("durationInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expiration date for assignments created in this policy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetModifiedBy sets the modifiedBy property value. The modifiedBy property
func (m *AccessPackageAssignmentPolicy) SetModifiedBy(value *string)() {
    err := m.GetBackingStore().Set("modifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("modifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetQuestions sets the questions property value. Questions that are posed to the  requestor.
func (m *AccessPackageAssignmentPolicy) SetQuestions(value []AccessPackageQuestionable)() {
    err := m.GetBackingStore().Set("questions", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestApprovalSettings sets the requestApprovalSettings property value. Who must approve requests for access package in this policy.
func (m *AccessPackageAssignmentPolicy) SetRequestApprovalSettings(value ApprovalSettingsable)() {
    err := m.GetBackingStore().Set("requestApprovalSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestorSettings sets the requestorSettings property value. Who can request this access package from this policy.
func (m *AccessPackageAssignmentPolicy) SetRequestorSettings(value RequestorSettingsable)() {
    err := m.GetBackingStore().Set("requestorSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetVerifiableCredentialSettings sets the verifiableCredentialSettings property value. Settings for verifiable credentials set up through the Microsoft Entra Verified I D service. These settings represent the verifiable credentials that a requestor of an access package in this policy can present to be assigned the access package.
func (m *AccessPackageAssignmentPolicy) SetVerifiableCredentialSettings(value VerifiableCredentialSettingsable)() {
    err := m.GetBackingStore().Set("verifiableCredentialSettings", value)
    if err != nil {
        panic(err)
    }
}
type AccessPackageAssignmentPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackage()(AccessPackageable)
    GetAccessPackageCatalog()(AccessPackageCatalogable)
    GetAccessPackageId()(*string)
    GetAccessPackageNotificationSettings()(AccessPackageNotificationSettingsable)
    GetAccessReviewSettings()(AssignmentReviewSettingsable)
    GetCanExtend()(*bool)
    GetCreatedBy()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomExtensionHandlers()([]CustomExtensionHandlerable)
    GetCustomExtensionStageSettings()([]CustomExtensionStageSettingable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDurationInDays()(*int32)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetModifiedBy()(*string)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetQuestions()([]AccessPackageQuestionable)
    GetRequestApprovalSettings()(ApprovalSettingsable)
    GetRequestorSettings()(RequestorSettingsable)
    GetVerifiableCredentialSettings()(VerifiableCredentialSettingsable)
    SetAccessPackage(value AccessPackageable)()
    SetAccessPackageCatalog(value AccessPackageCatalogable)()
    SetAccessPackageId(value *string)()
    SetAccessPackageNotificationSettings(value AccessPackageNotificationSettingsable)()
    SetAccessReviewSettings(value AssignmentReviewSettingsable)()
    SetCanExtend(value *bool)()
    SetCreatedBy(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomExtensionHandlers(value []CustomExtensionHandlerable)()
    SetCustomExtensionStageSettings(value []CustomExtensionStageSettingable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDurationInDays(value *int32)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetModifiedBy(value *string)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetQuestions(value []AccessPackageQuestionable)()
    SetRequestApprovalSettings(value ApprovalSettingsable)()
    SetRequestorSettings(value RequestorSettingsable)()
    SetVerifiableCredentialSettings(value VerifiableCredentialSettingsable)()
}
