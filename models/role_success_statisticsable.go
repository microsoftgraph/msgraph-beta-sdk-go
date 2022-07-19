package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleSuccessStatisticsable 
type RoleSuccessStatisticsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPermanentFail()(*int64)
    GetPermanentSuccess()(*int64)
    GetRemoveFail()(*int64)
    GetRemoveSuccess()(*int64)
    GetRoleId()(*string)
    GetRoleName()(*string)
    GetTemporaryFail()(*int64)
    GetTemporarySuccess()(*int64)
    GetUnknownFail()(*int64)
    SetOdataType(value *string)()
    SetPermanentFail(value *int64)()
    SetPermanentSuccess(value *int64)()
    SetRemoveFail(value *int64)()
    SetRemoveSuccess(value *int64)()
    SetRoleId(value *string)()
    SetRoleName(value *string)()
    SetTemporaryFail(value *int64)()
    SetTemporarySuccess(value *int64)()
    SetUnknownFail(value *int64)()
}
