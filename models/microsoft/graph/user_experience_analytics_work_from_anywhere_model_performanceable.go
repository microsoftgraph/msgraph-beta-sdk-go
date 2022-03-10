package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable 
type UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCloudIdentityScore()(*float64)
    GetCloudManagementScore()(*float64)
    GetCloudProvisioningScore()(*float64)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetModelDeviceCount()(*int32)
    GetWindowsScore()(*float64)
    GetWorkFromAnywhereScore()(*float64)
    SetCloudIdentityScore(value *float64)()
    SetCloudManagementScore(value *float64)()
    SetCloudProvisioningScore(value *float64)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetModelDeviceCount(value *int32)()
    SetWindowsScore(value *float64)()
    SetWorkFromAnywhereScore(value *float64)()
}
