package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewRecurrenceSettingsable 
type AccessReviewRecurrenceSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDurationInDays()(*int32)
    GetRecurrenceCount()(*int32)
    GetRecurrenceEndType()(*string)
    GetRecurrenceType()(*string)
    SetDurationInDays(value *int32)()
    SetRecurrenceCount(value *int32)()
    SetRecurrenceEndType(value *string)()
    SetRecurrenceType(value *string)()
}
