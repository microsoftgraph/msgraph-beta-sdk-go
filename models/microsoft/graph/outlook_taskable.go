package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutlookTaskable 
type OutlookTaskable interface {
    OutlookItemable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignedTo()(*string)
    GetAttachments()([]Attachmentable)
    GetBody()(ItemBodyable)
    GetCompletedDateTime()(DateTimeTimeZoneable)
    GetDueDateTime()(DateTimeTimeZoneable)
    GetHasAttachments()(*bool)
    GetImportance()(*Importance)
    GetIsReminderOn()(*bool)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetOwner()(*string)
    GetParentFolderId()(*string)
    GetRecurrence()(PatternedRecurrenceable)
    GetReminderDateTime()(DateTimeTimeZoneable)
    GetSensitivity()(*Sensitivity)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    GetStartDateTime()(DateTimeTimeZoneable)
    GetStatus()(*TaskStatus)
    GetSubject()(*string)
    SetAssignedTo(value *string)()
    SetAttachments(value []Attachmentable)()
    SetBody(value ItemBodyable)()
    SetCompletedDateTime(value DateTimeTimeZoneable)()
    SetDueDateTime(value DateTimeTimeZoneable)()
    SetHasAttachments(value *bool)()
    SetImportance(value *Importance)()
    SetIsReminderOn(value *bool)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetOwner(value *string)()
    SetParentFolderId(value *string)()
    SetRecurrence(value PatternedRecurrenceable)()
    SetReminderDateTime(value DateTimeTimeZoneable)()
    SetSensitivity(value *Sensitivity)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
    SetStartDateTime(value DateTimeTimeZoneable)()
    SetStatus(value *TaskStatus)()
    SetSubject(value *string)()
}
