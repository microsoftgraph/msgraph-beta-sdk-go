package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsAppHealthOSVersionPerformanceable 
type UserExperienceAnalyticsAppHealthOSVersionPerformanceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActiveDeviceCount()(*int32)
    GetMeanTimeToFailureInMinutes()(*int32)
    GetOsBuildNumber()(*string)
    GetOsVersion()(*string)
    GetOsVersionAppHealthScore()(*float64)
    GetOsVersionAppHealthStatus()(*string)
    SetActiveDeviceCount(value *int32)()
    SetMeanTimeToFailureInMinutes(value *int32)()
    SetOsBuildNumber(value *string)()
    SetOsVersion(value *string)()
    SetOsVersionAppHealthScore(value *float64)()
    SetOsVersionAppHealthStatus(value *string)()
}
