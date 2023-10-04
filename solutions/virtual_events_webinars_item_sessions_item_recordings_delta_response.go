package solutions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEventsWebinarsItemSessionsItemRecordingsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type VirtualEventsWebinarsItemSessionsItemRecordingsDeltaResponse struct {
    VirtualEventsWebinarsItemSessionsItemRecordingsDeltaGetResponse
}
// NewVirtualEventsWebinarsItemSessionsItemRecordingsDeltaResponse instantiates a new VirtualEventsWebinarsItemSessionsItemRecordingsDeltaResponse and sets the default values.
func NewVirtualEventsWebinarsItemSessionsItemRecordingsDeltaResponse()(*VirtualEventsWebinarsItemSessionsItemRecordingsDeltaResponse) {
    m := &VirtualEventsWebinarsItemSessionsItemRecordingsDeltaResponse{
        VirtualEventsWebinarsItemSessionsItemRecordingsDeltaGetResponse: *NewVirtualEventsWebinarsItemSessionsItemRecordingsDeltaGetResponse(),
    }
    return m
}
// CreateVirtualEventsWebinarsItemSessionsItemRecordingsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEventsWebinarsItemSessionsItemRecordingsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventsWebinarsItemSessionsItemRecordingsDeltaResponse(), nil
}
// VirtualEventsWebinarsItemSessionsItemRecordingsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type VirtualEventsWebinarsItemSessionsItemRecordingsDeltaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEventsWebinarsItemSessionsItemRecordingsDeltaGetResponseable
}
