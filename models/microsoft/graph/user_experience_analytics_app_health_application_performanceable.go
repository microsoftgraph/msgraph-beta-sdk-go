package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsAppHealthApplicationPerformanceable 
type UserExperienceAnalyticsAppHealthApplicationPerformanceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActiveDeviceCount()(*int32)
    GetAppCrashCount()(*int32)
    GetAppDisplayName()(*string)
    GetAppHangCount()(*int32)
    GetAppHealthScore()(*float64)
    GetAppHealthStatus()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetAppUsageDuration()(*int32)
    GetMeanTimeToFailureInMinutes()(*int32)
    SetActiveDeviceCount(value *int32)()
    SetAppCrashCount(value *int32)()
    SetAppDisplayName(value *string)()
    SetAppHangCount(value *int32)()
    SetAppHealthScore(value *float64)()
    SetAppHealthStatus(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetAppUsageDuration(value *int32)()
    SetMeanTimeToFailureInMinutes(value *int32)()
}
