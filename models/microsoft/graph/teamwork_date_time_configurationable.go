package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDateTimeConfigurationable 
type TeamworkDateTimeConfigurationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDateFormat()(*string)
    GetOfficeHoursEndTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)
    GetOfficeHoursStartTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)
    GetTimeFormat()(*string)
    GetTimeZone()(*string)
    SetDateFormat(value *string)()
    SetOfficeHoursEndTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)()
    SetOfficeHoursStartTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)()
    SetTimeFormat(value *string)()
    SetTimeZone(value *string)()
}
