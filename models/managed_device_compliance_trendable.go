package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceComplianceTrendable 
type ManagedDeviceComplianceTrendable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompliantDeviceCount()(*int32)
    GetConfigManagerDeviceCount()(*int32)
    GetCountDateTime()(*string)
    GetErrorDeviceCount()(*int32)
    GetInGracePeriodDeviceCount()(*int32)
    GetNoncompliantDeviceCount()(*int32)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    GetUnknownDeviceCount()(*int32)
    SetCompliantDeviceCount(value *int32)()
    SetConfigManagerDeviceCount(value *int32)()
    SetCountDateTime(value *string)()
    SetErrorDeviceCount(value *int32)()
    SetInGracePeriodDeviceCount(value *int32)()
    SetNoncompliantDeviceCount(value *int32)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetUnknownDeviceCount(value *int32)()
}
