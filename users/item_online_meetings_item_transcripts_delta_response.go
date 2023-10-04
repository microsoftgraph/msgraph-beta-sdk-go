package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOnlineMeetingsItemTranscriptsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemOnlineMeetingsItemTranscriptsDeltaResponse struct {
    ItemOnlineMeetingsItemTranscriptsDeltaGetResponse
}
// NewItemOnlineMeetingsItemTranscriptsDeltaResponse instantiates a new ItemOnlineMeetingsItemTranscriptsDeltaResponse and sets the default values.
func NewItemOnlineMeetingsItemTranscriptsDeltaResponse()(*ItemOnlineMeetingsItemTranscriptsDeltaResponse) {
    m := &ItemOnlineMeetingsItemTranscriptsDeltaResponse{
        ItemOnlineMeetingsItemTranscriptsDeltaGetResponse: *NewItemOnlineMeetingsItemTranscriptsDeltaGetResponse(),
    }
    return m
}
// CreateItemOnlineMeetingsItemTranscriptsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnlineMeetingsItemTranscriptsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnlineMeetingsItemTranscriptsDeltaResponse(), nil
}
// ItemOnlineMeetingsItemTranscriptsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemOnlineMeetingsItemTranscriptsDeltaResponseable interface {
    ItemOnlineMeetingsItemTranscriptsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
