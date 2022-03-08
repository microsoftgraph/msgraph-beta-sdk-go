package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeCardable 
type TimeCardable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    ChangeTrackedEntityable
    GetBreaks()([]TimeCardBreakable)
    GetClockInEvent()(TimeCardEventable)
    GetClockOutEvent()(TimeCardEventable)
    GetConfirmedBy()(*ConfirmedBy)
    GetNotes()(ItemBodyable)
    GetOriginalEntry()(TimeCardEntryable)
    GetState()(*TimeCardState)
    GetUserId()(*string)
    SetBreaks(value []TimeCardBreakable)()
    SetClockInEvent(value TimeCardEventable)()
    SetClockOutEvent(value TimeCardEventable)()
    SetConfirmedBy(value *ConfirmedBy)()
    SetNotes(value ItemBodyable)()
    SetOriginalEntry(value TimeCardEntryable)()
    SetState(value *TimeCardState)()
    SetUserId(value *string)()
}
