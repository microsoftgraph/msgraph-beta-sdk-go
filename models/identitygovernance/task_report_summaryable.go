package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TaskReportSummaryable 
type TaskReportSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFailedTasks()(*int32)
    GetOdataType()(*string)
    GetSuccessfulTasks()(*int32)
    GetTotalTasks()(*int32)
    GetUnprocessedTasks()(*int32)
    SetFailedTasks(value *int32)()
    SetOdataType(value *string)()
    SetSuccessfulTasks(value *int32)()
    SetTotalTasks(value *int32)()
    SetUnprocessedTasks(value *int32)()
}
