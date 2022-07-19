package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryable 
type UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutopilotDevicesSummary()(UserExperienceAnalyticsAutopilotDevicesSummaryable)
    GetCloudIdentityDevicesSummary()(UserExperienceAnalyticsCloudIdentityDevicesSummaryable)
    GetCloudManagementDevicesSummary()(UserExperienceAnalyticsCloudManagementDevicesSummaryable)
    GetCoManagedDevices()(*int32)
    GetDevicesNotAutopilotRegistered()(*int32)
    GetDevicesWithoutAutopilotProfileAssigned()(*int32)
    GetDevicesWithoutCloudIdentity()(*int32)
    GetIntuneDevices()(*int32)
    GetOdataType()(*string)
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
    SetOdataType(value *string)()
    SetTenantAttachDevices(value *int32)()
    SetTotalDevices(value *int32)()
    SetUnsupportedOSversionDevices(value *int32)()
    SetWindows10Devices(value *int32)()
    SetWindows10DevicesSummary(value UserExperienceAnalyticsWindows10DevicesSummaryable)()
    SetWindows10DevicesWithoutTenantAttach(value *int32)()
}
