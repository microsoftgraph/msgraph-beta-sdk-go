package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcOverviewable 
type CloudPcOverviewable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
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
    GetTotalCloudPcConnectionStatus()(*int32)
    GetTotalCloudPcStatus()(*int32)
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
    SetTotalCloudPcConnectionStatus(value *int32)()
    SetTotalCloudPcStatus(value *int32)()
}
