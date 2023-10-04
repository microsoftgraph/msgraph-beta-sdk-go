package planner

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TasksDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type TasksDeltaResponse struct {
    TasksDeltaGetResponse
}
// NewTasksDeltaResponse instantiates a new TasksDeltaResponse and sets the default values.
func NewTasksDeltaResponse()(*TasksDeltaResponse) {
    m := &TasksDeltaResponse{
        TasksDeltaGetResponse: *NewTasksDeltaGetResponse(),
    }
    return m
}
// CreateTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTasksDeltaResponse(), nil
}
// TasksDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type TasksDeltaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TasksDeltaGetResponseable
}
