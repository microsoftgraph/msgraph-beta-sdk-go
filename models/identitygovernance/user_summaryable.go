package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSummaryable 
type UserSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFailedTasks()(*int32)
    GetFailedUsers()(*int32)
    GetOdataType()(*string)
    GetSuccessfulUsers()(*int32)
    GetTotalTasks()(*int32)
    GetTotalUsers()(*int32)
    SetFailedTasks(value *int32)()
    SetFailedUsers(value *int32)()
    SetOdataType(value *string)()
    SetSuccessfulUsers(value *int32)()
    SetTotalTasks(value *int32)()
    SetTotalUsers(value *int32)()
}
