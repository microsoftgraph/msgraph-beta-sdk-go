package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable 
type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppCrashCount()(*int32)
    GetAppDisplayName()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetAppUsageDuration()(*int32)
    GetAppVersion()(*string)
    GetMeanTimeToFailureInMinutes()(*int32)
    SetAppCrashCount(value *int32)()
    SetAppDisplayName(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetAppUsageDuration(value *int32)()
    SetAppVersion(value *string)()
    SetMeanTimeToFailureInMinutes(value *int32)()
}
