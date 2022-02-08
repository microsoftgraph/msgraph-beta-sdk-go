package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutlookTask 
type OutlookTask struct {
    OutlookItem
    // The name of the person who has been assigned the task in Outlook. Read-only.
    assignedTo *string;
    // The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
    attachments []Attachment;
    // The task body that typically contains information about the task. Note that only HTML type is supported.
    body *ItemBody;
    // The date in the specified time zone that the task was finished.
    completedDateTime *DateTimeTimeZone;
    // The date in the specified time zone that the task is to be finished.
    dueDateTime *DateTimeTimeZone;
    // Set to true if the task has attachments.
    hasAttachments *bool;
    // 
    importance *Importance;
    // 
    isReminderOn *bool;
    // The collection of multi-value extended properties defined for the task. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedProperty;
    // 
    owner *string;
    // 
    parentFolderId *string;
    // 
    recurrence *PatternedRecurrence;
    // 
    reminderDateTime *DateTimeTimeZone;
    // 
    sensitivity *Sensitivity;
    // The collection of single-value extended properties defined for the task. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
    // 
    startDateTime *DateTimeTimeZone;
    // 
    status *TaskStatus;
    // 
    subject *string;
}
// NewOutlookTask instantiates a new outlookTask and sets the default values.
func NewOutlookTask()(*OutlookTask) {
    m := &OutlookTask{
        OutlookItem: *NewOutlookItem(),
    }
    return m
}
// GetAssignedTo gets the assignedTo property value. The name of the person who has been assigned the task in Outlook. Read-only.
func (m *OutlookTask) GetAssignedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
// GetAttachments gets the attachments property value. The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
func (m *OutlookTask) GetAttachments()([]Attachment) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
// GetBody gets the body property value. The task body that typically contains information about the task. Note that only HTML type is supported.
func (m *OutlookTask) GetBody()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// GetCompletedDateTime gets the completedDateTime property value. The date in the specified time zone that the task was finished.
func (m *OutlookTask) GetCompletedDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// GetDueDateTime gets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *OutlookTask) GetDueDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// GetHasAttachments gets the hasAttachments property value. Set to true if the task has attachments.
func (m *OutlookTask) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// GetImportance gets the importance property value. 
func (m *OutlookTask) GetImportance()(*Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// GetIsReminderOn gets the isReminderOn property value. 
func (m *OutlookTask) GetIsReminderOn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReminderOn
    }
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// GetOwner gets the owner property value. 
func (m *OutlookTask) GetOwner()(*string) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetParentFolderId gets the parentFolderId property value. 
func (m *OutlookTask) GetParentFolderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentFolderId
    }
}
// GetRecurrence gets the recurrence property value. 
func (m *OutlookTask) GetRecurrence()(*PatternedRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// GetReminderDateTime gets the reminderDateTime property value. 
func (m *OutlookTask) GetReminderDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.reminderDateTime
    }
}
// GetSensitivity gets the sensitivity property value. 
func (m *OutlookTask) GetSensitivity()(*Sensitivity) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// GetStartDateTime gets the startDateTime property value. 
func (m *OutlookTask) GetStartDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetStatus gets the status property value. 
func (m *OutlookTask) GetStatus()(*TaskStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetSubject gets the subject property value. 
func (m *OutlookTask) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookTask) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["assignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val)
        }
        return nil
    }
    res["attachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttachment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Attachment, len(val))
            for i, v := range val {
                res[i] = *(v.(*Attachment))
            }
            m.SetAttachments(res)
        }
        return nil
    }
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(*ItemBody))
        }
        return nil
    }
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["dueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["hasAttachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasAttachments(val)
        }
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportance(val.(*Importance))
        }
        return nil
    }
    res["isReminderOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReminderOn(val)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*MultiValueLegacyExtendedProperty))
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["parentFolderId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentFolderId(val)
        }
        return nil
    }
    res["recurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPatternedRecurrence() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(*PatternedRecurrence))
        }
        return nil
    }
    res["reminderDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["sensitivity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivity(val.(*Sensitivity))
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*SingleValueLegacyExtendedProperty))
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTaskStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*TaskStatus))
        }
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
func (m *OutlookTask) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OutlookTask) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    if m != nil {
        m.assignedTo = value
    }
}
// SetAttachments sets the attachments property value. The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
func (m *OutlookTask) SetAttachments(value []Attachment)() {
    if m != nil {
        m.attachments = value
    }
}
// SetBody sets the body property value. The task body that typically contains information about the task. Note that only HTML type is supported.
func (m *OutlookTask) SetBody(value *ItemBody)() {
    if m != nil {
        m.body = value
    }
}
// SetCompletedDateTime sets the completedDateTime property value. The date in the specified time zone that the task was finished.
func (m *OutlookTask) SetCompletedDateTime(value *DateTimeTimeZone)() {
    if m != nil {
        m.completedDateTime = value
    }
}
// SetDueDateTime sets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *OutlookTask) SetDueDateTime(value *DateTimeTimeZone)() {
    if m != nil {
        m.dueDateTime = value
    }
}
// SetHasAttachments sets the hasAttachments property value. Set to true if the task has attachments.
func (m *OutlookTask) SetHasAttachments(value *bool)() {
    if m != nil {
        m.hasAttachments = value
    }
}
// SetImportance sets the importance property value. 
func (m *OutlookTask) SetImportance(value *Importance)() {
    if m != nil {
        m.importance = value
    }
}
// SetIsReminderOn sets the isReminderOn property value. 
func (m *OutlookTask) SetIsReminderOn(value *bool)() {
    if m != nil {
        m.isReminderOn = value
    }
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    if m != nil {
        m.multiValueExtendedProperties = value
    }
}
// SetOwner sets the owner property value. 
func (m *OutlookTask) SetOwner(value *string)() {
    if m != nil {
        m.owner = value
    }
}
// SetParentFolderId sets the parentFolderId property value. 
func (m *OutlookTask) SetParentFolderId(value *string)() {
    if m != nil {
        m.parentFolderId = value
    }
}
// SetRecurrence sets the recurrence property value. 
func (m *OutlookTask) SetRecurrence(value *PatternedRecurrence)() {
    if m != nil {
        m.recurrence = value
    }
}
// SetReminderDateTime sets the reminderDateTime property value. 
func (m *OutlookTask) SetReminderDateTime(value *DateTimeTimeZone)() {
    if m != nil {
        m.reminderDateTime = value
    }
}
// SetSensitivity sets the sensitivity property value. 
func (m *OutlookTask) SetSensitivity(value *Sensitivity)() {
    if m != nil {
        m.sensitivity = value
    }
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    if m != nil {
        m.singleValueExtendedProperties = value
    }
}
// SetStartDateTime sets the startDateTime property value. 
func (m *OutlookTask) SetStartDateTime(value *DateTimeTimeZone)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetStatus sets the status property value. 
func (m *OutlookTask) SetStatus(value *TaskStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetSubject sets the subject property value. 
func (m *OutlookTask) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
