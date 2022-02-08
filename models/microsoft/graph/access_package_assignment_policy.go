package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentPolicy 
type AccessPackageAssignmentPolicy struct {
    Entity
    // The access package with this policy. Read-only. Nullable. Supports $expand.
    accessPackage *AccessPackage;
    // 
    accessPackageCatalog *AccessPackageCatalog;
    // Identifier of the access package.
    accessPackageId *string;
    // Who must review, and how often, the assignments to the access package from this policy. This property is null if reviews are not required.
    accessReviewSettings *AssignmentReviewSettings;
    // Indicates whether a user can extend the access package assignment duration after approval.
    canExtend *bool;
    // Read-only.
    createdBy *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    customExtensionHandlers []CustomExtensionHandler;
    // The description of the policy.
    description *string;
    // The display name of the policy. Supports $filter (eq).
    displayName *string;
    // The number of days in which assignments from this policy last until they are expired.
    durationInDays *int32;
    // The expiration date for assignments created in this policy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only.
    modifiedBy *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Questions that are posed to the  requestor.
    questions []AccessPackageQuestion;
    // Who must approve requests for access package in this policy.
    requestApprovalSettings *ApprovalSettings;
    // Who can request this access package from this policy.
    requestorSettings *RequestorSettings;
}
// NewAccessPackageAssignmentPolicy instantiates a new accessPackageAssignmentPolicy and sets the default values.
func NewAccessPackageAssignmentPolicy()(*AccessPackageAssignmentPolicy) {
    m := &AccessPackageAssignmentPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessPackage gets the accessPackage property value. The access package with this policy. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentPolicy) GetAccessPackage()(*AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackage
    }
}
// GetAccessPackageCatalog gets the accessPackageCatalog property value. 
func (m *AccessPackageAssignmentPolicy) GetAccessPackageCatalog()(*AccessPackageCatalog) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageCatalog
    }
}
// GetAccessPackageId gets the accessPackageId property value. Identifier of the access package.
func (m *AccessPackageAssignmentPolicy) GetAccessPackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageId
    }
}
// GetAccessReviewSettings gets the accessReviewSettings property value. Who must review, and how often, the assignments to the access package from this policy. This property is null if reviews are not required.
func (m *AccessPackageAssignmentPolicy) GetAccessReviewSettings()(*AssignmentReviewSettings) {
    if m == nil {
        return nil
    } else {
        return m.accessReviewSettings
    }
}
// GetCanExtend gets the canExtend property value. Indicates whether a user can extend the access package assignment duration after approval.
func (m *AccessPackageAssignmentPolicy) GetCanExtend()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.canExtend
    }
}
// GetCreatedBy gets the createdBy property value. Read-only.
func (m *AccessPackageAssignmentPolicy) GetCreatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCustomExtensionHandlers gets the customExtensionHandlers property value. 
func (m *AccessPackageAssignmentPolicy) GetCustomExtensionHandlers()([]CustomExtensionHandler) {
    if m == nil {
        return nil
    } else {
        return m.customExtensionHandlers
    }
}
// GetDescription gets the description property value. The description of the policy.
func (m *AccessPackageAssignmentPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name of the policy. Supports $filter (eq).
func (m *AccessPackageAssignmentPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDurationInDays gets the durationInDays property value. The number of days in which assignments from this policy last until they are expired.
func (m *AccessPackageAssignmentPolicy) GetDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInDays
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. The expiration date for assignments created in this policy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetModifiedBy gets the modifiedBy property value. Read-only.
func (m *AccessPackageAssignmentPolicy) GetModifiedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.modifiedBy
    }
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// GetQuestions gets the questions property value. Questions that are posed to the  requestor.
func (m *AccessPackageAssignmentPolicy) GetQuestions()([]AccessPackageQuestion) {
    if m == nil {
        return nil
    } else {
        return m.questions
    }
}
// GetRequestApprovalSettings gets the requestApprovalSettings property value. Who must approve requests for access package in this policy.
func (m *AccessPackageAssignmentPolicy) GetRequestApprovalSettings()(*ApprovalSettings) {
    if m == nil {
        return nil
    } else {
        return m.requestApprovalSettings
    }
}
// GetRequestorSettings gets the requestorSettings property value. Who can request this access package from this policy.
func (m *AccessPackageAssignmentPolicy) GetRequestorSettings()(*RequestorSettings) {
    if m == nil {
        return nil
    } else {
        return m.requestorSettings
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["accessPackageCatalog"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageCatalog() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageCatalog(val.(*AccessPackageCatalog))
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
    res["accessReviewSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignmentReviewSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessReviewSettings(val.(*AssignmentReviewSettings))
        }
        return nil
    }
    res["canExtend"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanExtend(val)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["customExtensionHandlers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomExtensionHandler() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomExtensionHandler, len(val))
            for i, v := range val {
                res[i] = *(v.(*CustomExtensionHandler))
            }
            m.SetCustomExtensionHandlers(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["durationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInDays(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["modifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedBy(val)
        }
        return nil
    }
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["questions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageQuestion() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageQuestion, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageQuestion))
            }
            m.SetQuestions(res)
        }
        return nil
    }
    res["requestApprovalSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApprovalSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestApprovalSettings(val.(*ApprovalSettings))
        }
        return nil
    }
    res["requestorSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRequestorSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestorSettings(val.(*RequestorSettings))
        }
        return nil
    }
    return res
}
func (m *AccessPackageAssignmentPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomExtensionHandlers()))
        for i, v := range m.GetCustomExtensionHandlers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetQuestions()))
        for i, v := range m.GetQuestions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    return nil
}
// SetAccessPackage sets the accessPackage property value. The access package with this policy. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentPolicy) SetAccessPackage(value *AccessPackage)() {
    if m != nil {
        m.accessPackage = value
    }
}
// SetAccessPackageCatalog sets the accessPackageCatalog property value. 
func (m *AccessPackageAssignmentPolicy) SetAccessPackageCatalog(value *AccessPackageCatalog)() {
    if m != nil {
        m.accessPackageCatalog = value
    }
}
// SetAccessPackageId sets the accessPackageId property value. Identifier of the access package.
func (m *AccessPackageAssignmentPolicy) SetAccessPackageId(value *string)() {
    if m != nil {
        m.accessPackageId = value
    }
}
// SetAccessReviewSettings sets the accessReviewSettings property value. Who must review, and how often, the assignments to the access package from this policy. This property is null if reviews are not required.
func (m *AccessPackageAssignmentPolicy) SetAccessReviewSettings(value *AssignmentReviewSettings)() {
    if m != nil {
        m.accessReviewSettings = value
    }
}
// SetCanExtend sets the canExtend property value. Indicates whether a user can extend the access package assignment duration after approval.
func (m *AccessPackageAssignmentPolicy) SetCanExtend(value *bool)() {
    if m != nil {
        m.canExtend = value
    }
}
// SetCreatedBy sets the createdBy property value. Read-only.
func (m *AccessPackageAssignmentPolicy) SetCreatedBy(value *string)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCustomExtensionHandlers sets the customExtensionHandlers property value. 
func (m *AccessPackageAssignmentPolicy) SetCustomExtensionHandlers(value []CustomExtensionHandler)() {
    if m != nil {
        m.customExtensionHandlers = value
    }
}
// SetDescription sets the description property value. The description of the policy.
func (m *AccessPackageAssignmentPolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the policy. Supports $filter (eq).
func (m *AccessPackageAssignmentPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDurationInDays sets the durationInDays property value. The number of days in which assignments from this policy last until they are expired.
func (m *AccessPackageAssignmentPolicy) SetDurationInDays(value *int32)() {
    if m != nil {
        m.durationInDays = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expiration date for assignments created in this policy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetModifiedBy sets the modifiedBy property value. Read-only.
func (m *AccessPackageAssignmentPolicy) SetModifiedBy(value *string)() {
    if m != nil {
        m.modifiedBy = value
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.modifiedDateTime = value
    }
}
// SetQuestions sets the questions property value. Questions that are posed to the  requestor.
func (m *AccessPackageAssignmentPolicy) SetQuestions(value []AccessPackageQuestion)() {
    if m != nil {
        m.questions = value
    }
}
// SetRequestApprovalSettings sets the requestApprovalSettings property value. Who must approve requests for access package in this policy.
func (m *AccessPackageAssignmentPolicy) SetRequestApprovalSettings(value *ApprovalSettings)() {
    if m != nil {
        m.requestApprovalSettings = value
    }
}
// SetRequestorSettings sets the requestorSettings property value. Who can request this access package from this policy.
func (m *AccessPackageAssignmentPolicy) SetRequestorSettings(value *RequestorSettings)() {
    if m != nil {
        m.requestorSettings = value
    }
}
