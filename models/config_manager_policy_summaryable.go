package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConfigManagerPolicySummaryable 
type ConfigManagerPolicySummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompliantDeviceCount()(*int32)
    GetEnforcedDeviceCount()(*int32)
    GetFailedDeviceCount()(*int32)
    GetNonCompliantDeviceCount()(*int32)
    GetOdataType()(*string)
    GetPendingDeviceCount()(*int32)
    GetTargetedDeviceCount()(*int32)
    SetCompliantDeviceCount(value *int32)()
    SetEnforcedDeviceCount(value *int32)()
    SetFailedDeviceCount(value *int32)()
    SetNonCompliantDeviceCount(value *int32)()
    SetOdataType(value *string)()
    SetPendingDeviceCount(value *int32)()
    SetTargetedDeviceCount(value *int32)()
}
