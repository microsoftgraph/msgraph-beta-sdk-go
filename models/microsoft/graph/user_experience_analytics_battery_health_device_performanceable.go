package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsBatteryHealthDevicePerformanceable 
type UserExperienceAnalyticsBatteryHealthDevicePerformanceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetBatteryAgeInDays()(*int32)
    GetDeviceBatteryHealthScore()(*int32)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetEstimatedRuntimeInMinutes()(*int32)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetManufacturer()(*string)
    GetMaxCapacityPercentage()(*int32)
    GetModel()(*string)
    SetBatteryAgeInDays(value *int32)()
    SetDeviceBatteryHealthScore(value *int32)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetEstimatedRuntimeInMinutes(value *int32)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetManufacturer(value *string)()
    SetMaxCapacityPercentage(value *int32)()
    SetModel(value *string)()
}
