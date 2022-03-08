package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable 
type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppCrashCount()(*int32)
    GetAppDisplayName()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetAppVersion()(*string)
    GetDeviceCountWithCrashes()(*int32)
    GetIsLatestUsedVersion()(*bool)
    GetIsMostUsedVersion()(*bool)
    SetAppCrashCount(value *int32)()
    SetAppDisplayName(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetAppVersion(value *string)()
    SetDeviceCountWithCrashes(value *int32)()
    SetIsLatestUsedVersion(value *bool)()
    SetIsMostUsedVersion(value *bool)()
}
