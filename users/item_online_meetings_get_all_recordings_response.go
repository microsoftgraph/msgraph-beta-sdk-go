package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOnlineMeetingsGetAllRecordingsResponse 
// Deprecated: This class is obsolete. Use getAllRecordingsGetResponse instead.
type ItemOnlineMeetingsGetAllRecordingsResponse struct {
    ItemOnlineMeetingsGetAllRecordingsGetResponse
}
// NewItemOnlineMeetingsGetAllRecordingsResponse instantiates a new ItemOnlineMeetingsGetAllRecordingsResponse and sets the default values.
func NewItemOnlineMeetingsGetAllRecordingsResponse()(*ItemOnlineMeetingsGetAllRecordingsResponse) {
    m := &ItemOnlineMeetingsGetAllRecordingsResponse{
        ItemOnlineMeetingsGetAllRecordingsGetResponse: *NewItemOnlineMeetingsGetAllRecordingsGetResponse(),
    }
    return m
}
// CreateItemOnlineMeetingsGetAllRecordingsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnlineMeetingsGetAllRecordingsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnlineMeetingsGetAllRecordingsResponse(), nil
}
// ItemOnlineMeetingsGetAllRecordingsResponseable 
// Deprecated: This class is obsolete. Use getAllRecordingsGetResponse instead.
type ItemOnlineMeetingsGetAllRecordingsResponseable interface {
    ItemOnlineMeetingsGetAllRecordingsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
