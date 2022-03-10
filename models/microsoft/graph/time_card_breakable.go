package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeCardBreakable 
type TimeCardBreakable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetBreakId()(*string)
    GetEnd()(TimeCardEventable)
    GetNotes()(ItemBodyable)
    GetStart()(TimeCardEventable)
    SetBreakId(value *string)()
    SetEnd(value TimeCardEventable)()
    SetNotes(value ItemBodyable)()
    SetStart(value TimeCardEventable)()
}
