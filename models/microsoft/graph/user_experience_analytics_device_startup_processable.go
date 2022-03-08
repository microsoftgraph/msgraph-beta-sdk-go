package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsDeviceStartupProcessable 
type UserExperienceAnalyticsDeviceStartupProcessable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetManagedDeviceId()(*string)
    GetProcessName()(*string)
    GetProductName()(*string)
    GetPublisher()(*string)
    GetStartupImpactInMs()(*int32)
    SetManagedDeviceId(value *string)()
    SetProcessName(value *string)()
    SetProductName(value *string)()
    SetPublisher(value *string)()
    SetStartupImpactInMs(value *int32)()
}
