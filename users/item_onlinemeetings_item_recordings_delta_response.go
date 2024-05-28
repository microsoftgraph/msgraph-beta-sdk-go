package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemOnlinemeetingsItemRecordingsDeltaGetResponseable instead.
type ItemOnlinemeetingsItemRecordingsDeltaResponse struct {
    ItemOnlinemeetingsItemRecordingsDeltaGetResponse
}
// NewItemOnlinemeetingsItemRecordingsDeltaResponse instantiates a new ItemOnlinemeetingsItemRecordingsDeltaResponse and sets the default values.
func NewItemOnlinemeetingsItemRecordingsDeltaResponse()(*ItemOnlinemeetingsItemRecordingsDeltaResponse) {
    m := &ItemOnlinemeetingsItemRecordingsDeltaResponse{
        ItemOnlinemeetingsItemRecordingsDeltaGetResponse: *NewItemOnlinemeetingsItemRecordingsDeltaGetResponse(),
    }
    return m
}
// CreateItemOnlinemeetingsItemRecordingsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOnlinemeetingsItemRecordingsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnlinemeetingsItemRecordingsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemOnlinemeetingsItemRecordingsDeltaGetResponseable instead.
type ItemOnlinemeetingsItemRecordingsDeltaResponseable interface {
    ItemOnlinemeetingsItemRecordingsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
