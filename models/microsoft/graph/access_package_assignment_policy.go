package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackageAssignmentPolicy struct {
    Entity
    // The access package with this policy. Read-only. Nullable.
    accessPackage *AccessPackage;
    // 
    accessPackageCatalog *AccessPackageCatalog;
    // ID of the access package.
    accessPackageId *string;
    // Who must review, and how often, the assignments to the access package from this policy. This property is null if reviews are not required.
    accessReviewSettings *AssignmentReviewSettings;
    // Indicates whether a user can extend the access package assignment duration after approval.
    canExtend *bool;
    // Read-only.
    createdBy *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description of the policy.
    description *string;
    // The display name of the policy.
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
// Instantiates a new accessPackageAssignmentPolicy and sets the default values.
func NewAccessPackageAssignmentPolicy()(*AccessPackageAssignmentPolicy) {
    m := &AccessPackageAssignmentPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessPackage property value. The access package with this policy. Read-only. Nullable.
func (m *AccessPackageAssignmentPolicy) GetAccessPackage()(*AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackage
    }
}
// Gets the accessPackageCatalog property value. 
func (m *AccessPackageAssignmentPolicy) GetAccessPackageCatalog()(*AccessPackageCatalog) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageCatalog
    }
}
// Gets the accessPackageId property value. ID of the access package.
func (m *AccessPackageAssignmentPolicy) GetAccessPackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageId
    }
}
// Gets the accessReviewSettings property value. Who must review, and how often, the assignments to the access package from this policy. This property is null if reviews are not required.
func (m *AccessPackageAssignmentPolicy) GetAccessReviewSettings()(*AssignmentReviewSettings) {
    if m == nil {
        return nil
    } else {
        return m.accessReviewSettings
    }
}
// Gets the canExtend property value. Indicates whether a user can extend the access package assignment duration after approval.
func (m *AccessPackageAssignmentPolicy) GetCanExtend()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.canExtend
    }
}
// Gets the createdBy property value. Read-only.
func (m *AccessPackageAssignmentPolicy) GetCreatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. The description of the policy.
func (m *AccessPackageAssignmentPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name of the policy.
func (m *AccessPackageAssignmentPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the durationInDays property value. The number of days in which assignments from this policy last until they are expired.
func (m *AccessPackageAssignmentPolicy) GetDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInDays
    }
}
// Gets the expirationDateTime property value. The expiration date for assignments created in this policy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the modifiedBy property value. Read-only.
func (m *AccessPackageAssignmentPolicy) GetModifiedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.modifiedBy
    }
}
// Gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// Gets the questions property value. Questions that are posed to the  requestor.
func (m *AccessPackageAssignmentPolicy) GetQuestions()([]AccessPackageQuestion) {
    if m == nil {
        return nil
    } else {
        return m.questions
    }
}
// Gets the requestApprovalSettings property value. Who must approve requests for access package in this policy.
func (m *AccessPackageAssignmentPolicy) GetRequestApprovalSettings()(*ApprovalSettings) {
    if m == nil {
        return nil
    } else {
        return m.requestApprovalSettings
    }
}
// Gets the requestorSettings property value. Who can request this access package from this policy.
func (m *AccessPackageAssignmentPolicy) GetRequestorSettings()(*RequestorSettings) {
    if m == nil {
        return nil
    } else {
        return m.requestorSettings
    }
}
// The deserialization information for the current model
func (m *AccessPackageAssignmentPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        m.SetAccessPackage(val.(*AccessPackage))
        return nil
    }
    res["accessPackageCatalog"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageCatalog() })
        if err != nil {
            return err
        }
        m.SetAccessPackageCatalog(val.(*AccessPackageCatalog))
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
    res["accessReviewSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignmentReviewSettings() })
        if err != nil {
            return err
        }
        m.SetAccessReviewSettings(val.(*AssignmentReviewSettings))
        return nil
    }
    res["canExtend"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCanExtend(val)
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCreatedBy(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["durationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDurationInDays(val)
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["modifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModifiedBy(val)
        return nil
    }
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetModifiedDateTime(val)
        return nil
    }
    res["questions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageQuestion() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageQuestion, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageQuestion))
        }
        m.SetQuestions(res)
        return nil
    }
    res["requestApprovalSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApprovalSettings() })
        if err != nil {
            return err
        }
        m.SetRequestApprovalSettings(val.(*ApprovalSettings))
        return nil
    }
    res["requestorSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRequestorSettings() })
        if err != nil {
            return err
        }
        m.SetRequestorSettings(val.(*RequestorSettings))
        return nil
    }
    return res
}
func (m *AccessPackageAssignmentPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
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
// Sets the accessPackage property value. The access package with this policy. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the accessPackage property.
func (m *AccessPackageAssignmentPolicy) SetAccessPackage(value *AccessPackage)() {
    m.accessPackage = value
}
// Sets the accessPackageCatalog property value. 
// Parameters:
//  - value : Value to set for the accessPackageCatalog property.
func (m *AccessPackageAssignmentPolicy) SetAccessPackageCatalog(value *AccessPackageCatalog)() {
    m.accessPackageCatalog = value
}
// Sets the accessPackageId property value. ID of the access package.
// Parameters:
//  - value : Value to set for the accessPackageId property.
func (m *AccessPackageAssignmentPolicy) SetAccessPackageId(value *string)() {
    m.accessPackageId = value
}
// Sets the accessReviewSettings property value. Who must review, and how often, the assignments to the access package from this policy. This property is null if reviews are not required.
// Parameters:
//  - value : Value to set for the accessReviewSettings property.
func (m *AccessPackageAssignmentPolicy) SetAccessReviewSettings(value *AssignmentReviewSettings)() {
    m.accessReviewSettings = value
}
// Sets the canExtend property value. Indicates whether a user can extend the access package assignment duration after approval.
// Parameters:
//  - value : Value to set for the canExtend property.
func (m *AccessPackageAssignmentPolicy) SetCanExtend(value *bool)() {
    m.canExtend = value
}
// Sets the createdBy property value. Read-only.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *AccessPackageAssignmentPolicy) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *AccessPackageAssignmentPolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. The description of the policy.
// Parameters:
//  - value : Value to set for the description property.
func (m *AccessPackageAssignmentPolicy) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name of the policy.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AccessPackageAssignmentPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the durationInDays property value. The number of days in which assignments from this policy last until they are expired.
// Parameters:
//  - value : Value to set for the durationInDays property.
func (m *AccessPackageAssignmentPolicy) SetDurationInDays(value *int32)() {
    m.durationInDays = value
}
// Sets the expirationDateTime property value. The expiration date for assignments created in this policy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *AccessPackageAssignmentPolicy) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the modifiedBy property value. Read-only.
// Parameters:
//  - value : Value to set for the modifiedBy property.
func (m *AccessPackageAssignmentPolicy) SetModifiedBy(value *string)() {
    m.modifiedBy = value
}
// Sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the modifiedDateTime property.
func (m *AccessPackageAssignmentPolicy) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// Sets the questions property value. Questions that are posed to the  requestor.
// Parameters:
//  - value : Value to set for the questions property.
func (m *AccessPackageAssignmentPolicy) SetQuestions(value []AccessPackageQuestion)() {
    m.questions = value
}
// Sets the requestApprovalSettings property value. Who must approve requests for access package in this policy.
// Parameters:
//  - value : Value to set for the requestApprovalSettings property.
func (m *AccessPackageAssignmentPolicy) SetRequestApprovalSettings(value *ApprovalSettings)() {
    m.requestApprovalSettings = value
}
// Sets the requestorSettings property value. Who can request this access package from this policy.
// Parameters:
//  - value : Value to set for the requestorSettings property.
func (m *AccessPackageAssignmentPolicy) SetRequestorSettings(value *RequestorSettings)() {
    m.requestorSettings = value
}
