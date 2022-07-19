package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppRelationshipStateable 
type MobileAppRelationshipStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceId()(*string)
    GetErrorCode()(*int32)
    GetInstallState()(*ResultantAppState)
    GetInstallStateDetail()(*ResultantAppStateDetail)
    GetOdataType()(*string)
    GetSourceIds()([]string)
    GetTargetDisplayName()(*string)
    GetTargetId()(*string)
    GetTargetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetDeviceId(value *string)()
    SetErrorCode(value *int32)()
    SetInstallState(value *ResultantAppState)()
    SetInstallStateDetail(value *ResultantAppStateDetail)()
    SetOdataType(value *string)()
    SetSourceIds(value []string)()
    SetTargetDisplayName(value *string)()
    SetTargetId(value *string)()
    SetTargetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
