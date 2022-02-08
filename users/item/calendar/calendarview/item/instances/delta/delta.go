package delta

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// Delta 
type Delta struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookItem
    // true if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Optional. Default is true.
    allowNewTimeProposals *bool;
    // The collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
    attachments []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attachment;
    // The collection of attendees for the event.
    attendees []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attendee;
    // The body of the message associated with the event. It can be in HTML or text format.
    body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody;
    // The preview of the message associated with the event. It is in text format.
    bodyPreview *string;
    // The calendar that contains the event. Navigation property. Read-only.
    calendar *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar;
    // Contains occurrenceId property values of cancelled instances in a recurring series, if the event is the series master. Instances in a recurring series that are cancelled are called cancelledOccurences.Returned only on $select in a Get operation which specifies the id of a series master event (that is, the seriesMasterId property value).
    cancelledOccurrences []string;
    // The date, time, and time zone that the event ends. By default, the end time is in UTC.
    end *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone;
    // 
    exceptionOccurrences []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event;
    // The collection of open extensions defined for the event. Nullable.
    extensions []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension;
    // Set to true if the event has attachments.
    hasAttachments *bool;
    // When set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false.
    hideAttendees *bool;
    // 
    importance *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Importance;
    // The occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
    instances []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event;
    // 
    isAllDay *bool;
    // 
    isCancelled *bool;
    // 
    isDraft *bool;
    // 
    isOnlineMeeting *bool;
    // 
    isOrganizer *bool;
    // 
    isReminderOn *bool;
    // 
    location *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location;
    // 
    locations []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location;
    // The collection of multi-value extended properties defined for the event. Read-only. Nullable.
    multiValueExtendedProperties []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty;
    // 
    occurrenceId *string;
    // 
    onlineMeeting *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeetingInfo;
    // 
    onlineMeetingProvider *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeetingProviderType;
    // 
    onlineMeetingUrl *string;
    // 
    organizer *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient;
    // 
    originalEndTimeZone *string;
    // 
    originalStart *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    originalStartTimeZone *string;
    // 
    recurrence *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PatternedRecurrence;
    // 
    reminderMinutesBeforeStart *int32;
    // 
    responseRequested *bool;
    // 
    responseStatus *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResponseStatus;
    // 
    sensitivity *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Sensitivity;
    // 
    seriesMasterId *string;
    // 
    showAs *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FreeBusyStatus;
    // The collection of single-value extended properties defined for the event. Read-only. Nullable.
    singleValueExtendedProperties []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty;
    // 
    start *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone;
    // 
    subject *string;
    // 
    transactionId *string;
    // 
    type_escaped *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EventType;
    // 
    uid *string;
    // 
    webLink *string;
}
// NewDelta instantiates a new delta and sets the default values.
func NewDelta()(*Delta) {
    m := &Delta{
        OutlookItem: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookItem(),
    }
    return m
}
// GetAllowNewTimeProposals gets the allowNewTimeProposals property value. true if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Optional. Default is true.
func (m *Delta) GetAllowNewTimeProposals()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowNewTimeProposals
    }
}
// GetAttachments gets the attachments property value. The collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
func (m *Delta) GetAttachments()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attachment) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
// GetAttendees gets the attendees property value. The collection of attendees for the event.
func (m *Delta) GetAttendees()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attendee) {
    if m == nil {
        return nil
    } else {
        return m.attendees
    }
}
// GetBody gets the body property value. The body of the message associated with the event. It can be in HTML or text format.
func (m *Delta) GetBody()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// GetBodyPreview gets the bodyPreview property value. The preview of the message associated with the event. It is in text format.
func (m *Delta) GetBodyPreview()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bodyPreview
    }
}
// GetCalendar gets the calendar property value. The calendar that contains the event. Navigation property. Read-only.
func (m *Delta) GetCalendar()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar) {
    if m == nil {
        return nil
    } else {
        return m.calendar
    }
}
// GetCancelledOccurrences gets the cancelledOccurrences property value. Contains occurrenceId property values of cancelled instances in a recurring series, if the event is the series master. Instances in a recurring series that are cancelled are called cancelledOccurences.Returned only on $select in a Get operation which specifies the id of a series master event (that is, the seriesMasterId property value).
func (m *Delta) GetCancelledOccurrences()([]string) {
    if m == nil {
        return nil
    } else {
        return m.cancelledOccurrences
    }
}
// GetEnd gets the end property value. The date, time, and time zone that the event ends. By default, the end time is in UTC.
func (m *Delta) GetEnd()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
// GetExceptionOccurrences gets the exceptionOccurrences property value. 
func (m *Delta) GetExceptionOccurrences()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event) {
    if m == nil {
        return nil
    } else {
        return m.exceptionOccurrences
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the event. Nullable.
func (m *Delta) GetExtensions()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetHasAttachments gets the hasAttachments property value. Set to true if the event has attachments.
func (m *Delta) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// GetHideAttendees gets the hideAttendees property value. When set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false.
func (m *Delta) GetHideAttendees()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideAttendees
    }
}
// GetImportance gets the importance property value. 
func (m *Delta) GetImportance()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// GetInstances gets the instances property value. The occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *Delta) GetInstances()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event) {
    if m == nil {
        return nil
    } else {
        return m.instances
    }
}
// GetIsAllDay gets the isAllDay property value. 
func (m *Delta) GetIsAllDay()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAllDay
    }
}
// GetIsCancelled gets the isCancelled property value. 
func (m *Delta) GetIsCancelled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCancelled
    }
}
// GetIsDraft gets the isDraft property value. 
func (m *Delta) GetIsDraft()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDraft
    }
}
// GetIsOnlineMeeting gets the isOnlineMeeting property value. 
func (m *Delta) GetIsOnlineMeeting()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOnlineMeeting
    }
}
// GetIsOrganizer gets the isOrganizer property value. 
func (m *Delta) GetIsOrganizer()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOrganizer
    }
}
// GetIsReminderOn gets the isReminderOn property value. 
func (m *Delta) GetIsReminderOn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReminderOn
    }
}
// GetLocation gets the location property value. 
func (m *Delta) GetLocation()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// GetLocations gets the locations property value. 
func (m *Delta) GetLocations()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location) {
    if m == nil {
        return nil
    } else {
        return m.locations
    }
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the event. Read-only. Nullable.
func (m *Delta) GetMultiValueExtendedProperties()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// GetOccurrenceId gets the occurrenceId property value. 
func (m *Delta) GetOccurrenceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.occurrenceId
    }
}
// GetOnlineMeeting gets the onlineMeeting property value. 
func (m *Delta) GetOnlineMeeting()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeetingInfo) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeeting
    }
}
// GetOnlineMeetingProvider gets the onlineMeetingProvider property value. 
func (m *Delta) GetOnlineMeetingProvider()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeetingProviderType) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetingProvider
    }
}
// GetOnlineMeetingUrl gets the onlineMeetingUrl property value. 
func (m *Delta) GetOnlineMeetingUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetingUrl
    }
}
// GetOrganizer gets the organizer property value. 
func (m *Delta) GetOrganizer()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient) {
    if m == nil {
        return nil
    } else {
        return m.organizer
    }
}
// GetOriginalEndTimeZone gets the originalEndTimeZone property value. 
func (m *Delta) GetOriginalEndTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originalEndTimeZone
    }
}
// GetOriginalStart gets the originalStart property value. 
func (m *Delta) GetOriginalStart()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.originalStart
    }
}
// GetOriginalStartTimeZone gets the originalStartTimeZone property value. 
func (m *Delta) GetOriginalStartTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originalStartTimeZone
    }
}
// GetRecurrence gets the recurrence property value. 
func (m *Delta) GetRecurrence()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PatternedRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// GetReminderMinutesBeforeStart gets the reminderMinutesBeforeStart property value. 
func (m *Delta) GetReminderMinutesBeforeStart()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.reminderMinutesBeforeStart
    }
}
// GetResponseRequested gets the responseRequested property value. 
func (m *Delta) GetResponseRequested()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.responseRequested
    }
}
// GetResponseStatus gets the responseStatus property value. 
func (m *Delta) GetResponseStatus()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResponseStatus) {
    if m == nil {
        return nil
    } else {
        return m.responseStatus
    }
}
// GetSensitivity gets the sensitivity property value. 
func (m *Delta) GetSensitivity()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Sensitivity) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
// GetSeriesMasterId gets the seriesMasterId property value. 
func (m *Delta) GetSeriesMasterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.seriesMasterId
    }
}
// GetShowAs gets the showAs property value. 
func (m *Delta) GetShowAs()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FreeBusyStatus) {
    if m == nil {
        return nil
    } else {
        return m.showAs
    }
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the event. Read-only. Nullable.
func (m *Delta) GetSingleValueExtendedProperties()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// GetStart gets the start property value. 
func (m *Delta) GetStart()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
// GetSubject gets the subject property value. 
func (m *Delta) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// GetTransactionId gets the transactionId property value. 
func (m *Delta) GetTransactionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.transactionId
    }
}
// GetType gets the type property value. 
func (m *Delta) GetType()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EventType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetUid gets the uid property value. 
func (m *Delta) GetUid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uid
    }
}
// GetWebLink gets the webLink property value. 
func (m *Delta) GetWebLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webLink
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Delta) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["allowNewTimeProposals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowNewTimeProposals(val)
        }
        return nil
    }
    res["attachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAttachment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attachment, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attachment))
            }
            m.SetAttachments(res)
        }
        return nil
    }
    res["attendees"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAttendee() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attendee, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attendee))
            }
            m.SetAttendees(res)
        }
        return nil
    }
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody))
        }
        return nil
    }
    res["bodyPreview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyPreview(val)
        }
        return nil
    }
    res["calendar"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCalendar() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendar(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar))
        }
        return nil
    }
    res["cancelledOccurrences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCancelledOccurrences(res)
        }
        return nil
    }
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone))
        }
        return nil
    }
    res["exceptionOccurrences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event))
            }
            m.SetExceptionOccurrences(res)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension))
            }
            m.SetExtensions(res)
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
    res["hideAttendees"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideAttendees(val)
        }
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportance(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Importance))
        }
        return nil
    }
    res["instances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event))
            }
            m.SetInstances(res)
        }
        return nil
    }
    res["isAllDay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAllDay(val)
        }
        return nil
    }
    res["isCancelled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCancelled(val)
        }
        return nil
    }
    res["isDraft"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDraft(val)
        }
        return nil
    }
    res["isOnlineMeeting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOnlineMeeting(val)
        }
        return nil
    }
    res["isOrganizer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOrganizer(val)
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
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location))
        }
        return nil
    }
    res["locations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location))
            }
            m.SetLocations(res)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty))
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["occurrenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOccurrenceId(val)
        }
        return nil
    }
    res["onlineMeeting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnlineMeetingInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineMeeting(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeetingInfo))
        }
        return nil
    }
    res["onlineMeetingProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseOnlineMeetingProviderType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineMeetingProvider(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeetingProviderType))
        }
        return nil
    }
    res["onlineMeetingUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineMeetingUrl(val)
        }
        return nil
    }
    res["organizer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizer(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient))
        }
        return nil
    }
    res["originalEndTimeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalEndTimeZone(val)
        }
        return nil
    }
    res["originalStart"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalStart(val)
        }
        return nil
    }
    res["originalStartTimeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalStartTimeZone(val)
        }
        return nil
    }
    res["recurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPatternedRecurrence() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PatternedRecurrence))
        }
        return nil
    }
    res["reminderMinutesBeforeStart"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderMinutesBeforeStart(val)
        }
        return nil
    }
    res["responseRequested"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseRequested(val)
        }
        return nil
    }
    res["responseStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewResponseStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseStatus(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResponseStatus))
        }
        return nil
    }
    res["sensitivity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseSensitivity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivity(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Sensitivity))
        }
        return nil
    }
    res["seriesMasterId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeriesMasterId(val)
        }
        return nil
    }
    res["showAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseFreeBusyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowAs(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FreeBusyStatus))
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty))
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone))
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
    res["transactionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransactionId(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseEventType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EventType))
        }
        return nil
    }
    res["uid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUid(val)
        }
        return nil
    }
    res["webLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebLink(val)
        }
        return nil
    }
    return res
}
func (m *Delta) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Delta) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowNewTimeProposals", m.GetAllowNewTimeProposals())
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
    if m.GetAttendees() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("attendees", cast)
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
        err = writer.WriteStringValue("bodyPreview", m.GetBodyPreview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("calendar", m.GetCalendar())
        if err != nil {
            return err
        }
    }
    if m.GetCancelledOccurrences() != nil {
        err = writer.WriteCollectionOfStringValues("cancelledOccurrences", m.GetCancelledOccurrences())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    if m.GetExceptionOccurrences() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExceptionOccurrences()))
        for i, v := range m.GetExceptionOccurrences() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("exceptionOccurrences", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
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
    {
        err = writer.WriteBoolValue("hideAttendees", m.GetHideAttendees())
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
    if m.GetInstances() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInstances()))
        for i, v := range m.GetInstances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("instances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAllDay", m.GetIsAllDay())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCancelled", m.GetIsCancelled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDraft", m.GetIsDraft())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOnlineMeeting", m.GetIsOnlineMeeting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOrganizer", m.GetIsOrganizer())
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
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    if m.GetLocations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocations()))
        for i, v := range m.GetLocations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("locations", cast)
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
        err = writer.WriteStringValue("occurrenceId", m.GetOccurrenceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onlineMeeting", m.GetOnlineMeeting())
        if err != nil {
            return err
        }
    }
    if m.GetOnlineMeetingProvider() != nil {
        cast := (*m.GetOnlineMeetingProvider()).String()
        err = writer.WriteStringValue("onlineMeetingProvider", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onlineMeetingUrl", m.GetOnlineMeetingUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("organizer", m.GetOrganizer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originalEndTimeZone", m.GetOriginalEndTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("originalStart", m.GetOriginalStart())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originalStartTimeZone", m.GetOriginalStartTimeZone())
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
        err = writer.WriteInt32Value("reminderMinutesBeforeStart", m.GetReminderMinutesBeforeStart())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("responseRequested", m.GetResponseRequested())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("responseStatus", m.GetResponseStatus())
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
    {
        err = writer.WriteStringValue("seriesMasterId", m.GetSeriesMasterId())
        if err != nil {
            return err
        }
    }
    if m.GetShowAs() != nil {
        cast := (*m.GetShowAs()).String()
        err = writer.WriteStringValue("showAs", &cast)
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
        err = writer.WriteObjectValue("start", m.GetStart())
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
    {
        err = writer.WriteStringValue("transactionId", m.GetTransactionId())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uid", m.GetUid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webLink", m.GetWebLink())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowNewTimeProposals sets the allowNewTimeProposals property value. true if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Optional. Default is true.
func (m *Delta) SetAllowNewTimeProposals(value *bool)() {
    if m != nil {
        m.allowNewTimeProposals = value
    }
}
// SetAttachments sets the attachments property value. The collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
func (m *Delta) SetAttachments(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attachment)() {
    if m != nil {
        m.attachments = value
    }
}
// SetAttendees sets the attendees property value. The collection of attendees for the event.
func (m *Delta) SetAttendees(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Attendee)() {
    if m != nil {
        m.attendees = value
    }
}
// SetBody sets the body property value. The body of the message associated with the event. It can be in HTML or text format.
func (m *Delta) SetBody(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody)() {
    if m != nil {
        m.body = value
    }
}
// SetBodyPreview sets the bodyPreview property value. The preview of the message associated with the event. It is in text format.
func (m *Delta) SetBodyPreview(value *string)() {
    if m != nil {
        m.bodyPreview = value
    }
}
// SetCalendar sets the calendar property value. The calendar that contains the event. Navigation property. Read-only.
func (m *Delta) SetCalendar(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar)() {
    if m != nil {
        m.calendar = value
    }
}
// SetCancelledOccurrences sets the cancelledOccurrences property value. Contains occurrenceId property values of cancelled instances in a recurring series, if the event is the series master. Instances in a recurring series that are cancelled are called cancelledOccurences.Returned only on $select in a Get operation which specifies the id of a series master event (that is, the seriesMasterId property value).
func (m *Delta) SetCancelledOccurrences(value []string)() {
    if m != nil {
        m.cancelledOccurrences = value
    }
}
// SetEnd sets the end property value. The date, time, and time zone that the event ends. By default, the end time is in UTC.
func (m *Delta) SetEnd(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone)() {
    if m != nil {
        m.end = value
    }
}
// SetExceptionOccurrences sets the exceptionOccurrences property value. 
func (m *Delta) SetExceptionOccurrences(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event)() {
    if m != nil {
        m.exceptionOccurrences = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the event. Nullable.
func (m *Delta) SetExtensions(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension)() {
    if m != nil {
        m.extensions = value
    }
}
// SetHasAttachments sets the hasAttachments property value. Set to true if the event has attachments.
func (m *Delta) SetHasAttachments(value *bool)() {
    if m != nil {
        m.hasAttachments = value
    }
}
// SetHideAttendees sets the hideAttendees property value. When set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false.
func (m *Delta) SetHideAttendees(value *bool)() {
    if m != nil {
        m.hideAttendees = value
    }
}
// SetImportance sets the importance property value. 
func (m *Delta) SetImportance(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Importance)() {
    if m != nil {
        m.importance = value
    }
}
// SetInstances sets the instances property value. The occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *Delta) SetInstances(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event)() {
    if m != nil {
        m.instances = value
    }
}
// SetIsAllDay sets the isAllDay property value. 
func (m *Delta) SetIsAllDay(value *bool)() {
    if m != nil {
        m.isAllDay = value
    }
}
// SetIsCancelled sets the isCancelled property value. 
func (m *Delta) SetIsCancelled(value *bool)() {
    if m != nil {
        m.isCancelled = value
    }
}
// SetIsDraft sets the isDraft property value. 
func (m *Delta) SetIsDraft(value *bool)() {
    if m != nil {
        m.isDraft = value
    }
}
// SetIsOnlineMeeting sets the isOnlineMeeting property value. 
func (m *Delta) SetIsOnlineMeeting(value *bool)() {
    if m != nil {
        m.isOnlineMeeting = value
    }
}
// SetIsOrganizer sets the isOrganizer property value. 
func (m *Delta) SetIsOrganizer(value *bool)() {
    if m != nil {
        m.isOrganizer = value
    }
}
// SetIsReminderOn sets the isReminderOn property value. 
func (m *Delta) SetIsReminderOn(value *bool)() {
    if m != nil {
        m.isReminderOn = value
    }
}
// SetLocation sets the location property value. 
func (m *Delta) SetLocation(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location)() {
    if m != nil {
        m.location = value
    }
}
// SetLocations sets the locations property value. 
func (m *Delta) SetLocations(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location)() {
    if m != nil {
        m.locations = value
    }
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the event. Read-only. Nullable.
func (m *Delta) SetMultiValueExtendedProperties(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty)() {
    if m != nil {
        m.multiValueExtendedProperties = value
    }
}
// SetOccurrenceId sets the occurrenceId property value. 
func (m *Delta) SetOccurrenceId(value *string)() {
    if m != nil {
        m.occurrenceId = value
    }
}
// SetOnlineMeeting sets the onlineMeeting property value. 
func (m *Delta) SetOnlineMeeting(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeetingInfo)() {
    if m != nil {
        m.onlineMeeting = value
    }
}
// SetOnlineMeetingProvider sets the onlineMeetingProvider property value. 
func (m *Delta) SetOnlineMeetingProvider(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeetingProviderType)() {
    if m != nil {
        m.onlineMeetingProvider = value
    }
}
// SetOnlineMeetingUrl sets the onlineMeetingUrl property value. 
func (m *Delta) SetOnlineMeetingUrl(value *string)() {
    if m != nil {
        m.onlineMeetingUrl = value
    }
}
// SetOrganizer sets the organizer property value. 
func (m *Delta) SetOrganizer(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient)() {
    if m != nil {
        m.organizer = value
    }
}
// SetOriginalEndTimeZone sets the originalEndTimeZone property value. 
func (m *Delta) SetOriginalEndTimeZone(value *string)() {
    if m != nil {
        m.originalEndTimeZone = value
    }
}
// SetOriginalStart sets the originalStart property value. 
func (m *Delta) SetOriginalStart(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.originalStart = value
    }
}
// SetOriginalStartTimeZone sets the originalStartTimeZone property value. 
func (m *Delta) SetOriginalStartTimeZone(value *string)() {
    if m != nil {
        m.originalStartTimeZone = value
    }
}
// SetRecurrence sets the recurrence property value. 
func (m *Delta) SetRecurrence(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PatternedRecurrence)() {
    if m != nil {
        m.recurrence = value
    }
}
// SetReminderMinutesBeforeStart sets the reminderMinutesBeforeStart property value. 
func (m *Delta) SetReminderMinutesBeforeStart(value *int32)() {
    if m != nil {
        m.reminderMinutesBeforeStart = value
    }
}
// SetResponseRequested sets the responseRequested property value. 
func (m *Delta) SetResponseRequested(value *bool)() {
    if m != nil {
        m.responseRequested = value
    }
}
// SetResponseStatus sets the responseStatus property value. 
func (m *Delta) SetResponseStatus(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResponseStatus)() {
    if m != nil {
        m.responseStatus = value
    }
}
// SetSensitivity sets the sensitivity property value. 
func (m *Delta) SetSensitivity(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Sensitivity)() {
    if m != nil {
        m.sensitivity = value
    }
}
// SetSeriesMasterId sets the seriesMasterId property value. 
func (m *Delta) SetSeriesMasterId(value *string)() {
    if m != nil {
        m.seriesMasterId = value
    }
}
// SetShowAs sets the showAs property value. 
func (m *Delta) SetShowAs(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FreeBusyStatus)() {
    if m != nil {
        m.showAs = value
    }
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the event. Read-only. Nullable.
func (m *Delta) SetSingleValueExtendedProperties(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty)() {
    if m != nil {
        m.singleValueExtendedProperties = value
    }
}
// SetStart sets the start property value. 
func (m *Delta) SetStart(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone)() {
    if m != nil {
        m.start = value
    }
}
// SetSubject sets the subject property value. 
func (m *Delta) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
// SetTransactionId sets the transactionId property value. 
func (m *Delta) SetTransactionId(value *string)() {
    if m != nil {
        m.transactionId = value
    }
}
// SetType sets the type property value. 
func (m *Delta) SetType(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EventType)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetUid sets the uid property value. 
func (m *Delta) SetUid(value *string)() {
    if m != nil {
        m.uid = value
    }
}
// SetWebLink sets the webLink property value. 
func (m *Delta) SetWebLink(value *string)() {
    if m != nil {
        m.webLink = value
    }
}
