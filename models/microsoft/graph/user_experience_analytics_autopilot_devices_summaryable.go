package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsAutopilotDevicesSummaryable 
type UserExperienceAnalyticsAutopilotDevicesSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDevicesNotAutopilotRegistered()(*int32)
    GetDevicesWithoutAutopilotProfileAssigned()(*int32)
    GetTotalWindows10DevicesWithoutTenantAttached()(*int32)
    SetDevicesNotAutopilotRegistered(value *int32)()
    SetDevicesWithoutAutopilotProfileAssigned(value *int32)()
    SetTotalWindows10DevicesWithoutTenantAttached(value *int32)()
}
