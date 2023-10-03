package solutions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEventsEventsItemSessionsItemRecordingsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type VirtualEventsEventsItemSessionsItemRecordingsDeltaResponse struct {
    VirtualEventsEventsItemSessionsItemRecordingsDeltaGetResponse
}
// NewVirtualEventsEventsItemSessionsItemRecordingsDeltaResponse instantiates a new VirtualEventsEventsItemSessionsItemRecordingsDeltaResponse and sets the default values.
func NewVirtualEventsEventsItemSessionsItemRecordingsDeltaResponse()(*VirtualEventsEventsItemSessionsItemRecordingsDeltaResponse) {
    m := &VirtualEventsEventsItemSessionsItemRecordingsDeltaResponse{
        VirtualEventsEventsItemSessionsItemRecordingsDeltaGetResponse: *NewVirtualEventsEventsItemSessionsItemRecordingsDeltaGetResponse(),
    }
    return m
}
// CreateVirtualEventsEventsItemSessionsItemRecordingsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEventsEventsItemSessionsItemRecordingsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventsEventsItemSessionsItemRecordingsDeltaResponse(), nil
}
// VirtualEventsEventsItemSessionsItemRecordingsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type VirtualEventsEventsItemSessionsItemRecordingsDeltaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEventsEventsItemSessionsItemRecordingsDeltaGetResponseable
}
