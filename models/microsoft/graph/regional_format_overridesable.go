package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RegionalFormatOverridesable 
type RegionalFormatOverridesable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCalendar()(*string)
    GetFirstDayOfWeek()(*string)
    GetLongDateFormat()(*string)
    GetLongTimeFormat()(*string)
    GetShortDateFormat()(*string)
    GetShortTimeFormat()(*string)
    GetTimeZone()(*string)
    SetCalendar(value *string)()
    SetFirstDayOfWeek(value *string)()
    SetLongDateFormat(value *string)()
    SetLongTimeFormat(value *string)()
    SetShortDateFormat(value *string)()
    SetShortTimeFormat(value *string)()
    SetTimeZone(value *string)()
}
