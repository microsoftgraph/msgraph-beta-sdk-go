package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeCardEntryable 
type TimeCardEntryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetBreaks()([]TimeCardBreakable)
    GetClockInEvent()(TimeCardEventable)
    GetClockOutEvent()(TimeCardEventable)
    SetBreaks(value []TimeCardBreakable)()
    SetClockInEvent(value TimeCardEventable)()
    SetClockOutEvent(value TimeCardEventable)()
}
