package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOnlineMeetingsItemRecordingsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemOnlineMeetingsItemRecordingsDeltaResponse struct {
    ItemOnlineMeetingsItemRecordingsDeltaGetResponse
}
// NewItemOnlineMeetingsItemRecordingsDeltaResponse instantiates a new ItemOnlineMeetingsItemRecordingsDeltaResponse and sets the default values.
func NewItemOnlineMeetingsItemRecordingsDeltaResponse()(*ItemOnlineMeetingsItemRecordingsDeltaResponse) {
    m := &ItemOnlineMeetingsItemRecordingsDeltaResponse{
        ItemOnlineMeetingsItemRecordingsDeltaGetResponse: *NewItemOnlineMeetingsItemRecordingsDeltaGetResponse(),
    }
    return m
}
// CreateItemOnlineMeetingsItemRecordingsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnlineMeetingsItemRecordingsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnlineMeetingsItemRecordingsDeltaResponse(), nil
}
// ItemOnlineMeetingsItemRecordingsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemOnlineMeetingsItemRecordingsDeltaResponseable interface {
    ItemOnlineMeetingsItemRecordingsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
