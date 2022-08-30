package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RunSummaryable 
type RunSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFailedRuns()(*int32)
    GetFailedTasks()(*int32)
    GetOdataType()(*string)
    GetSuccessfulRuns()(*int32)
    GetTotalRuns()(*int32)
    GetTotalTasks()(*int32)
    GetTotalUsers()(*int32)
    SetFailedRuns(value *int32)()
    SetFailedTasks(value *int32)()
    SetOdataType(value *string)()
    SetSuccessfulRuns(value *int32)()
    SetTotalRuns(value *int32)()
    SetTotalTasks(value *int32)()
    SetTotalUsers(value *int32)()
}
