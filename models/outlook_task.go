package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutlookTask 
type OutlookTask struct {
    OutlookItem
    // The name of the person who has been assigned the task in Outlook. Read-only.
    assignedTo *string
    // The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
    attachments []Attachmentable
    // The task body that typically contains information about the task. Note that only HTML type is supported.
    body ItemBodyable
    // The date in the specified time zone that the task was finished.
    completedDateTime DateTimeTimeZoneable
    // The date in the specified time zone that the task is to be finished.
    dueDateTime DateTimeTimeZoneable
    // Set to true if the task has attachments.
    hasAttachments *bool
    // The importance property
    importance *Importance
    // The isReminderOn property
    isReminderOn *bool
    // The collection of multi-value extended properties defined for the task. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The owner property
    owner *string
    // The parentFolderId property
    parentFolderId *string
    // The recurrence property
    recurrence PatternedRecurrenceable
    // The reminderDateTime property
    reminderDateTime DateTimeTimeZoneable
    // The sensitivity property
    sensitivity *Sensitivity
    // The collection of single-value extended properties defined for the task. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
    // The startDateTime property
    startDateTime DateTimeTimeZoneable
    // The status property
    status *TaskStatus
    // The subject property
    subject *string
}
// NewOutlookTask instantiates a new OutlookTask and sets the default values.
func NewOutlookTask()(*OutlookTask) {
    m := &OutlookTask{
        OutlookItem: *NewOutlookItem(),
    }
    odataTypeValue := "#microsoft.graph.outlookTask";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOutlookTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutlookTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutlookTask(), nil
}
// GetAssignedTo gets the assignedTo property value. The name of the person who has been assigned the task in Outlook. Read-only.
func (m *OutlookTask) GetAssignedTo()(*string) {
    return m.assignedTo
}
// GetAttachments gets the attachments property value. The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
func (m *OutlookTask) GetAttachments()([]Attachmentable) {
    return m.attachments
}
// GetBody gets the body property value. The task body that typically contains information about the task. Note that only HTML type is supported.
func (m *OutlookTask) GetBody()(ItemBodyable) {
    return m.body
}
// GetCompletedDateTime gets the completedDateTime property value. The date in the specified time zone that the task was finished.
func (m *OutlookTask) GetCompletedDateTime()(DateTimeTimeZoneable) {
    return m.completedDateTime
}
// GetDueDateTime gets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *OutlookTask) GetDueDateTime()(DateTimeTimeZoneable) {
    return m.dueDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["assignedTo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAssignedTo)
    res["attachments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAttachmentFromDiscriminatorValue , m.SetAttachments)
    res["body"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemBodyFromDiscriminatorValue , m.SetBody)
    res["completedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue , m.SetCompletedDateTime)
    res["dueDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue , m.SetDueDateTime)
    res["hasAttachments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetHasAttachments)
    res["importance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseImportance , m.SetImportance)
    res["isReminderOn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsReminderOn)
    res["multiValueExtendedProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue , m.SetMultiValueExtendedProperties)
    res["owner"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOwner)
    res["parentFolderId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetParentFolderId)
    res["recurrence"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePatternedRecurrenceFromDiscriminatorValue , m.SetRecurrence)
    res["reminderDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue , m.SetReminderDateTime)
    res["sensitivity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSensitivity , m.SetSensitivity)
    res["singleValueExtendedProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue , m.SetSingleValueExtendedProperties)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue , m.SetStartDateTime)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTaskStatus , m.SetStatus)
    res["subject"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubject)
    return res
}
// GetHasAttachments gets the hasAttachments property value. Set to true if the task has attachments.
func (m *OutlookTask) GetHasAttachments()(*bool) {
    return m.hasAttachments
}
// GetImportance gets the importance property value. The importance property
func (m *OutlookTask) GetImportance()(*Importance) {
    return m.importance
}
// GetIsReminderOn gets the isReminderOn property value. The isReminderOn property
func (m *OutlookTask) GetIsReminderOn()(*bool) {
    return m.isReminderOn
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    return m.multiValueExtendedProperties
}
// GetOwner gets the owner property value. The owner property
func (m *OutlookTask) GetOwner()(*string) {
    return m.owner
}
// GetParentFolderId gets the parentFolderId property value. The parentFolderId property
func (m *OutlookTask) GetParentFolderId()(*string) {
    return m.parentFolderId
}
// GetRecurrence gets the recurrence property value. The recurrence property
func (m *OutlookTask) GetRecurrence()(PatternedRecurrenceable) {
    return m.recurrence
}
// GetReminderDateTime gets the reminderDateTime property value. The reminderDateTime property
func (m *OutlookTask) GetReminderDateTime()(DateTimeTimeZoneable) {
    return m.reminderDateTime
}
// GetSensitivity gets the sensitivity property value. The sensitivity property
func (m *OutlookTask) GetSensitivity()(*Sensitivity) {
    return m.sensitivity
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    return m.singleValueExtendedProperties
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *OutlookTask) GetStartDateTime()(DateTimeTimeZoneable) {
    return m.startDateTime
}
// GetStatus gets the status property value. The status property
func (m *OutlookTask) GetStatus()(*TaskStatus) {
    return m.status
}
// GetSubject gets the subject property value. The subject property
func (m *OutlookTask) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *OutlookTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    if m.GetAttachments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAttachments())
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := (*m.GetImportance()).String()
        err = writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReminderOn", m.GetIsReminderOn())
        if err != nil {
            return err
        }
    }
    if m.GetMultiValueExtendedProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMultiValueExtendedProperties())
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reminderDateTime", m.GetReminderDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSensitivity() != nil {
        cast := (*m.GetSensitivity()).String()
        err = writer.WriteStringValue("sensitivity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSingleValueExtendedProperties())
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTo sets the assignedTo property value. The name of the person who has been assigned the task in Outlook. Read-only.
func (m *OutlookTask) SetAssignedTo(value *string)() {
    m.assignedTo = value
}
// SetAttachments sets the attachments property value. The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
func (m *OutlookTask) SetAttachments(value []Attachmentable)() {
    m.attachments = value
}
// SetBody sets the body property value. The task body that typically contains information about the task. Note that only HTML type is supported.
func (m *OutlookTask) SetBody(value ItemBodyable)() {
    m.body = value
}
// SetCompletedDateTime sets the completedDateTime property value. The date in the specified time zone that the task was finished.
func (m *OutlookTask) SetCompletedDateTime(value DateTimeTimeZoneable)() {
    m.completedDateTime = value
}
// SetDueDateTime sets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *OutlookTask) SetDueDateTime(value DateTimeTimeZoneable)() {
    m.dueDateTime = value
}
// SetHasAttachments sets the hasAttachments property value. Set to true if the task has attachments.
func (m *OutlookTask) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// SetImportance sets the importance property value. The importance property
func (m *OutlookTask) SetImportance(value *Importance)() {
    m.importance = value
}
// SetIsReminderOn sets the isReminderOn property value. The isReminderOn property
func (m *OutlookTask) SetIsReminderOn(value *bool)() {
    m.isReminderOn = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    m.multiValueExtendedProperties = value
}
// SetOwner sets the owner property value. The owner property
func (m *OutlookTask) SetOwner(value *string)() {
    m.owner = value
}
// SetParentFolderId sets the parentFolderId property value. The parentFolderId property
func (m *OutlookTask) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// SetRecurrence sets the recurrence property value. The recurrence property
func (m *OutlookTask) SetRecurrence(value PatternedRecurrenceable)() {
    m.recurrence = value
}
// SetReminderDateTime sets the reminderDateTime property value. The reminderDateTime property
func (m *OutlookTask) SetReminderDateTime(value DateTimeTimeZoneable)() {
    m.reminderDateTime = value
}
// SetSensitivity sets the sensitivity property value. The sensitivity property
func (m *OutlookTask) SetSensitivity(value *Sensitivity)() {
    m.sensitivity = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    m.singleValueExtendedProperties = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *OutlookTask) SetStartDateTime(value DateTimeTimeZoneable)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *OutlookTask) SetStatus(value *TaskStatus)() {
    m.status = value
}
// SetSubject sets the subject property value. The subject property
func (m *OutlookTask) SetSubject(value *string)() {
    m.subject = value
}
