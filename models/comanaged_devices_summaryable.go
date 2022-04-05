package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagedDevicesSummaryable 
type ComanagedDevicesSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompliancePolicyCount()(*int32)
    GetConfigurationSettingsCount()(*int32)
    GetEndpointProtectionCount()(*int32)
    GetInventoryCount()(*int32)
    GetModernAppsCount()(*int32)
    GetOfficeAppsCount()(*int32)
    GetResourceAccessCount()(*int32)
    GetTotalComanagedCount()(*int32)
    GetWindowsUpdateForBusinessCount()(*int32)
    SetCompliancePolicyCount(value *int32)()
    SetConfigurationSettingsCount(value *int32)()
    SetEndpointProtectionCount(value *int32)()
    SetInventoryCount(value *int32)()
    SetModernAppsCount(value *int32)()
    SetOfficeAppsCount(value *int32)()
    SetResourceAccessCount(value *int32)()
    SetTotalComanagedCount(value *int32)()
    SetWindowsUpdateForBusinessCount(value *int32)()
}
