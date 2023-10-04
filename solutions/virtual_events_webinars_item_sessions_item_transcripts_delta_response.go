package solutions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEventsWebinarsItemSessionsItemTranscriptsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type VirtualEventsWebinarsItemSessionsItemTranscriptsDeltaResponse struct {
    VirtualEventsWebinarsItemSessionsItemTranscriptsDeltaGetResponse
}
// NewVirtualEventsWebinarsItemSessionsItemTranscriptsDeltaResponse instantiates a new VirtualEventsWebinarsItemSessionsItemTranscriptsDeltaResponse and sets the default values.
func NewVirtualEventsWebinarsItemSessionsItemTranscriptsDeltaResponse()(*VirtualEventsWebinarsItemSessionsItemTranscriptsDeltaResponse) {
    m := &VirtualEventsWebinarsItemSessionsItemTranscriptsDeltaResponse{
        VirtualEventsWebinarsItemSessionsItemTranscriptsDeltaGetResponse: *NewVirtualEventsWebinarsItemSessionsItemTranscriptsDeltaGetResponse(),
    }
    return m
}
// CreateVirtualEventsWebinarsItemSessionsItemTranscriptsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEventsWebinarsItemSessionsItemTranscriptsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventsWebinarsItemSessionsItemTranscriptsDeltaResponse(), nil
}
// VirtualEventsWebinarsItemSessionsItemTranscriptsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type VirtualEventsWebinarsItemSessionsItemTranscriptsDeltaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEventsWebinarsItemSessionsItemTranscriptsDeltaGetResponseable
}
