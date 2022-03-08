package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable 
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetOsCheckFailedPercentage()(*float64)
    GetProcessor64BitCheckFailedPercentage()(*float64)
    GetProcessorCoreCountCheckFailedPercentage()(*float64)
    GetProcessorFamilyCheckFailedPercentage()(*float64)
    GetProcessorSpeedCheckFailedPercentage()(*float64)
    GetRamCheckFailedPercentage()(*float64)
    GetSecureBootCheckFailedPercentage()(*float64)
    GetStorageCheckFailedPercentage()(*float64)
    GetTotalDeviceCount()(*int32)
    GetTpmCheckFailedPercentage()(*float64)
    GetUpgradeEligibleDeviceCount()(*int32)
    SetOsCheckFailedPercentage(value *float64)()
    SetProcessor64BitCheckFailedPercentage(value *float64)()
    SetProcessorCoreCountCheckFailedPercentage(value *float64)()
    SetProcessorFamilyCheckFailedPercentage(value *float64)()
    SetProcessorSpeedCheckFailedPercentage(value *float64)()
    SetRamCheckFailedPercentage(value *float64)()
    SetSecureBootCheckFailedPercentage(value *float64)()
    SetStorageCheckFailedPercentage(value *float64)()
    SetTotalDeviceCount(value *int32)()
    SetTpmCheckFailedPercentage(value *float64)()
    SetUpgradeEligibleDeviceCount(value *int32)()
}
