package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsBatteryHealthDeviceAppImpactable 
type UserExperienceAnalyticsBatteryHealthDeviceAppImpactable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppDisplayName()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetBatteryUsagePercentage()(*float64)
    GetDeviceId()(*string)
    GetIsForegroundApp()(*bool)
    SetAppDisplayName(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetBatteryUsagePercentage(value *float64)()
    SetDeviceId(value *string)()
    SetIsForegroundApp(value *bool)()
}
