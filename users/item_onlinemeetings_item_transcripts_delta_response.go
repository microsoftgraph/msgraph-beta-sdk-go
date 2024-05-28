package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemOnlinemeetingsItemTranscriptsDeltaGetResponseable instead.
type ItemOnlinemeetingsItemTranscriptsDeltaResponse struct {
    ItemOnlinemeetingsItemTranscriptsDeltaGetResponse
}
// NewItemOnlinemeetingsItemTranscriptsDeltaResponse instantiates a new ItemOnlinemeetingsItemTranscriptsDeltaResponse and sets the default values.
func NewItemOnlinemeetingsItemTranscriptsDeltaResponse()(*ItemOnlinemeetingsItemTranscriptsDeltaResponse) {
    m := &ItemOnlinemeetingsItemTranscriptsDeltaResponse{
        ItemOnlinemeetingsItemTranscriptsDeltaGetResponse: *NewItemOnlinemeetingsItemTranscriptsDeltaGetResponse(),
    }
    return m
}
// CreateItemOnlinemeetingsItemTranscriptsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOnlinemeetingsItemTranscriptsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnlinemeetingsItemTranscriptsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemOnlinemeetingsItemTranscriptsDeltaGetResponseable instead.
type ItemOnlinemeetingsItemTranscriptsDeltaResponseable interface {
    ItemOnlinemeetingsItemTranscriptsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
