package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsDeviceStartupProcessPerformanceable 
type UserExperienceAnalyticsDeviceStartupProcessPerformanceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceCount()(*int64)
    GetMedianImpactInMs()(*int32)
    GetMedianImpactInMs2()(*int64)
    GetProcessName()(*string)
    GetProductName()(*string)
    GetPublisher()(*string)
    GetTotalImpactInMs()(*int32)
    GetTotalImpactInMs2()(*int64)
    SetDeviceCount(value *int64)()
    SetMedianImpactInMs(value *int32)()
    SetMedianImpactInMs2(value *int64)()
    SetProcessName(value *string)()
    SetProductName(value *string)()
    SetPublisher(value *string)()
    SetTotalImpactInMs(value *int32)()
    SetTotalImpactInMs2(value *int64)()
}
