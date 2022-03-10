package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsImpactingProcessable 
type UserExperienceAnalyticsImpactingProcessable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCategory()(*string)
    GetDescription()(*string)
    GetDeviceId()(*string)
    GetImpactValue()(*float64)
    GetProcessName()(*string)
    GetPublisher()(*string)
    SetCategory(value *string)()
    SetDescription(value *string)()
    SetDeviceId(value *string)()
    SetImpactValue(value *float64)()
    SetProcessName(value *string)()
    SetPublisher(value *string)()
}
