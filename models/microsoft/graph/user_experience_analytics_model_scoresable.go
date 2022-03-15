package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsModelScoresable 
type UserExperienceAnalyticsModelScoresable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppReliabilityScore()(*float64)
    GetEndpointAnalyticsScore()(*float64)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetModelDeviceCount()(*int32)
    GetStartupPerformanceScore()(*float64)
    GetWorkFromAnywhereScore()(*float64)
    SetAppReliabilityScore(value *float64)()
    SetEndpointAnalyticsScore(value *float64)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetModelDeviceCount(value *int32)()
    SetStartupPerformanceScore(value *float64)()
    SetWorkFromAnywhereScore(value *float64)()
}
