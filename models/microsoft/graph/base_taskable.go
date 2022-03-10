package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BaseTaskable 
type BaseTaskable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetBodyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetChecklistItems()([]ChecklistItemable)
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetDueDateTime()(DateTimeTimeZoneable)
    GetExtensions()([]Extensionable)
    GetImportance()(*Importance)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLinkedResources()([]LinkedResource_v2able)
    GetParentList()(BaseTaskListable)
    GetRecurrence()(PatternedRecurrenceable)
    GetStartDateTime()(DateTimeTimeZoneable)
    GetStatus()(*TaskStatus_v2)
    GetTextBody()(*string)
    GetViewpoint()(TaskViewpointable)
    SetBodyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetChecklistItems(value []ChecklistItemable)()
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetDueDateTime(value DateTimeTimeZoneable)()
    SetExtensions(value []Extensionable)()
    SetImportance(value *Importance)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLinkedResources(value []LinkedResource_v2able)()
    SetParentList(value BaseTaskListable)()
    SetRecurrence(value PatternedRecurrenceable)()
    SetStartDateTime(value DateTimeTimeZoneable)()
    SetStatus(value *TaskStatus_v2)()
    SetTextBody(value *string)()
    SetViewpoint(value TaskViewpointable)()
}
