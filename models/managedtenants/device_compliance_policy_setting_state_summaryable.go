package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeviceCompliancePolicySettingStateSummaryable 
type DeviceCompliancePolicySettingStateSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConflictDeviceCount()(*int32)
    GetErrorDeviceCount()(*int32)
    GetFailedDeviceCount()(*int32)
    GetIntuneAccountId()(*string)
    GetIntuneSettingId()(*string)
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotApplicableDeviceCount()(*int32)
    GetPendingDeviceCount()(*int32)
    GetPolicyType()(*string)
    GetSettingName()(*string)
    GetSucceededDeviceCount()(*int32)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    SetConflictDeviceCount(value *int32)()
    SetErrorDeviceCount(value *int32)()
    SetFailedDeviceCount(value *int32)()
    SetIntuneAccountId(value *string)()
    SetIntuneSettingId(value *string)()
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotApplicableDeviceCount(value *int32)()
    SetPendingDeviceCount(value *int32)()
    SetPolicyType(value *string)()
    SetSettingName(value *string)()
    SetSucceededDeviceCount(value *int32)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
}
