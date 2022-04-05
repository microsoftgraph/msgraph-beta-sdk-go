package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsCloudManagementDevicesSummaryable 
type UserExperienceAnalyticsCloudManagementDevicesSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCoManagedDeviceCount()(*int32)
    GetIntuneDeviceCount()(*int32)
    GetTenantAttachDeviceCount()(*int32)
    SetCoManagedDeviceCount(value *int32)()
    SetIntuneDeviceCount(value *int32)()
    SetTenantAttachDeviceCount(value *int32)()
}
