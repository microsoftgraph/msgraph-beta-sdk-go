package solutions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEventsEventsItemSessionsItemTranscriptsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type VirtualEventsEventsItemSessionsItemTranscriptsDeltaResponse struct {
    VirtualEventsEventsItemSessionsItemTranscriptsDeltaGetResponse
}
// NewVirtualEventsEventsItemSessionsItemTranscriptsDeltaResponse instantiates a new VirtualEventsEventsItemSessionsItemTranscriptsDeltaResponse and sets the default values.
func NewVirtualEventsEventsItemSessionsItemTranscriptsDeltaResponse()(*VirtualEventsEventsItemSessionsItemTranscriptsDeltaResponse) {
    m := &VirtualEventsEventsItemSessionsItemTranscriptsDeltaResponse{
        VirtualEventsEventsItemSessionsItemTranscriptsDeltaGetResponse: *NewVirtualEventsEventsItemSessionsItemTranscriptsDeltaGetResponse(),
    }
    return m
}
// CreateVirtualEventsEventsItemSessionsItemTranscriptsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEventsEventsItemSessionsItemTranscriptsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventsEventsItemSessionsItemTranscriptsDeltaResponse(), nil
}
// VirtualEventsEventsItemSessionsItemTranscriptsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type VirtualEventsEventsItemSessionsItemTranscriptsDeltaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEventsEventsItemSessionsItemTranscriptsDeltaGetResponseable
}
