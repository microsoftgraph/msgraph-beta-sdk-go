package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcOverviewable 
type CloudPcOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumberOfCloudPcConnectionStatusFailed()(*int32)
    GetNumberOfCloudPcConnectionStatusPassed()(*int32)
    GetNumberOfCloudPcConnectionStatusPending()(*int32)
    GetNumberOfCloudPcConnectionStatusRunning()(*int32)
    GetNumberOfCloudPcConnectionStatusUnkownFutureValue()(*int32)
    GetNumberOfCloudPcStatusDeprovisioning()(*int32)
    GetNumberOfCloudPcStatusFailed()(*int32)
    GetNumberOfCloudPcStatusInGracePeriod()(*int32)
    GetNumberOfCloudPcStatusNotProvisioned()(*int32)
    GetNumberOfCloudPcStatusProvisioned()(*int32)
    GetNumberOfCloudPcStatusProvisioning()(*int32)
    GetNumberOfCloudPcStatusUnknown()(*int32)
    GetNumberOfCloudPcStatusUpgrading()(*int32)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    GetTotalBusinessLicenses()(*int32)
    GetTotalCloudPcConnectionStatus()(*int32)
    GetTotalCloudPcStatus()(*int32)
    GetTotalEnterpriseLicenses()(*int32)
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumberOfCloudPcConnectionStatusFailed(value *int32)()
    SetNumberOfCloudPcConnectionStatusPassed(value *int32)()
    SetNumberOfCloudPcConnectionStatusPending(value *int32)()
    SetNumberOfCloudPcConnectionStatusRunning(value *int32)()
    SetNumberOfCloudPcConnectionStatusUnkownFutureValue(value *int32)()
    SetNumberOfCloudPcStatusDeprovisioning(value *int32)()
    SetNumberOfCloudPcStatusFailed(value *int32)()
    SetNumberOfCloudPcStatusInGracePeriod(value *int32)()
    SetNumberOfCloudPcStatusNotProvisioned(value *int32)()
    SetNumberOfCloudPcStatusProvisioned(value *int32)()
    SetNumberOfCloudPcStatusProvisioning(value *int32)()
    SetNumberOfCloudPcStatusUnknown(value *int32)()
    SetNumberOfCloudPcStatusUpgrading(value *int32)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetTotalBusinessLicenses(value *int32)()
    SetTotalCloudPcConnectionStatus(value *int32)()
    SetTotalCloudPcStatus(value *int32)()
    SetTotalEnterpriseLicenses(value *int32)()
}
