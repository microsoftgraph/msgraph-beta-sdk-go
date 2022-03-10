package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable 
type UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActiveDeviceCount()(*int32)
    GetAppCrashCount()(*int32)
    GetAppDisplayName()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetAppUsageDuration()(*int32)
    GetMeanTimeToFailureInMinutes()(*int32)
    GetOsBuildNumber()(*string)
    GetOsVersion()(*string)
    SetActiveDeviceCount(value *int32)()
    SetAppCrashCount(value *int32)()
    SetAppDisplayName(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetAppUsageDuration(value *int32)()
    SetMeanTimeToFailureInMinutes(value *int32)()
    SetOsBuildNumber(value *string)()
    SetOsVersion(value *string)()
}
