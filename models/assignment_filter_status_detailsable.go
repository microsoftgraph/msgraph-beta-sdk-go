package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterStatusDetailsable 
type AssignmentFilterStatusDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceProperties()([]KeyValuePairable)
    GetEvalutionSummaries()([]AssignmentFilterEvaluationSummaryable)
    GetManagedDeviceId()(*string)
    GetPayloadId()(*string)
    GetUserId()(*string)
    SetDeviceProperties(value []KeyValuePairable)()
    SetEvalutionSummaries(value []AssignmentFilterEvaluationSummaryable)()
    SetManagedDeviceId(value *string)()
    SetPayloadId(value *string)()
    SetUserId(value *string)()
}
