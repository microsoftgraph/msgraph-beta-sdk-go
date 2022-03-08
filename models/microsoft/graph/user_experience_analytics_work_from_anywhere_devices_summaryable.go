package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryable 
type UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAutopilotDevicesSummary()(UserExperienceAnalyticsAutopilotDevicesSummaryable)
    GetCloudIdentityDevicesSummary()(UserExperienceAnalyticsCloudIdentityDevicesSummaryable)
    GetCloudManagementDevicesSummary()(UserExperienceAnalyticsCloudManagementDevicesSummaryable)
    GetCoManagedDevices()(*int32)
    GetDevicesNotAutopilotRegistered()(*int32)
    GetDevicesWithoutAutopilotProfileAssigned()(*int32)
    GetDevicesWithoutCloudIdentity()(*int32)
    GetIntuneDevices()(*int32)
    GetTenantAttachDevices()(*int32)
    GetTotalDevices()(*int32)
    GetUnsupportedOSversionDevices()(*int32)
    GetWindows10Devices()(*int32)
    GetWindows10DevicesSummary()(UserExperienceAnalyticsWindows10DevicesSummaryable)
    GetWindows10DevicesWithoutTenantAttach()(*int32)
    SetAutopilotDevicesSummary(value UserExperienceAnalyticsAutopilotDevicesSummaryable)()
    SetCloudIdentityDevicesSummary(value UserExperienceAnalyticsCloudIdentityDevicesSummaryable)()
    SetCloudManagementDevicesSummary(value UserExperienceAnalyticsCloudManagementDevicesSummaryable)()
    SetCoManagedDevices(value *int32)()
    SetDevicesNotAutopilotRegistered(value *int32)()
    SetDevicesWithoutAutopilotProfileAssigned(value *int32)()
    SetDevicesWithoutCloudIdentity(value *int32)()
    SetIntuneDevices(value *int32)()
    SetTenantAttachDevices(value *int32)()
    SetTotalDevices(value *int32)()
    SetUnsupportedOSversionDevices(value *int32)()
    SetWindows10Devices(value *int32)()
    SetWindows10DevicesSummary(value UserExperienceAnalyticsWindows10DevicesSummaryable)()
    SetWindows10DevicesWithoutTenantAttach(value *int32)()
}
