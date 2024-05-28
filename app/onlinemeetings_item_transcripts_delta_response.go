package app

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OnlinemeetingsItemTranscriptsDeltaGetResponseable instead.
type OnlinemeetingsItemTranscriptsDeltaResponse struct {
    OnlinemeetingsItemTranscriptsDeltaGetResponse
}
// NewOnlinemeetingsItemTranscriptsDeltaResponse instantiates a new OnlinemeetingsItemTranscriptsDeltaResponse and sets the default values.
func NewOnlinemeetingsItemTranscriptsDeltaResponse()(*OnlinemeetingsItemTranscriptsDeltaResponse) {
    m := &OnlinemeetingsItemTranscriptsDeltaResponse{
        OnlinemeetingsItemTranscriptsDeltaGetResponse: *NewOnlinemeetingsItemTranscriptsDeltaGetResponse(),
    }
    return m
}
// CreateOnlinemeetingsItemTranscriptsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnlinemeetingsItemTranscriptsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlinemeetingsItemTranscriptsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use OnlinemeetingsItemTranscriptsDeltaGetResponseable instead.
type OnlinemeetingsItemTranscriptsDeltaResponseable interface {
    OnlinemeetingsItemTranscriptsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
