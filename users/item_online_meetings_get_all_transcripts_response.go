package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemOnlineMeetingsGetAllTranscriptsGetResponseable instead.
type ItemOnlineMeetingsGetAllTranscriptsResponse struct {
    ItemOnlineMeetingsGetAllTranscriptsGetResponse
}
// NewItemOnlineMeetingsGetAllTranscriptsResponse instantiates a new ItemOnlineMeetingsGetAllTranscriptsResponse and sets the default values.
func NewItemOnlineMeetingsGetAllTranscriptsResponse()(*ItemOnlineMeetingsGetAllTranscriptsResponse) {
    m := &ItemOnlineMeetingsGetAllTranscriptsResponse{
        ItemOnlineMeetingsGetAllTranscriptsGetResponse: *NewItemOnlineMeetingsGetAllTranscriptsGetResponse(),
    }
    return m
}
// CreateItemOnlineMeetingsGetAllTranscriptsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOnlineMeetingsGetAllTranscriptsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnlineMeetingsGetAllTranscriptsResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemOnlineMeetingsGetAllTranscriptsGetResponseable instead.
type ItemOnlineMeetingsGetAllTranscriptsResponseable interface {
    ItemOnlineMeetingsGetAllTranscriptsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
