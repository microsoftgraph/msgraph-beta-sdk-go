package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemFindroomsFindRoomsGetResponseable instead.
type ItemFindroomsFindRoomsResponse struct {
    ItemFindroomsFindRoomsGetResponse
}
// NewItemFindroomsFindRoomsResponse instantiates a new ItemFindroomsFindRoomsResponse and sets the default values.
func NewItemFindroomsFindRoomsResponse()(*ItemFindroomsFindRoomsResponse) {
    m := &ItemFindroomsFindRoomsResponse{
        ItemFindroomsFindRoomsGetResponse: *NewItemFindroomsFindRoomsGetResponse(),
    }
    return m
}
// CreateItemFindroomsFindRoomsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemFindroomsFindRoomsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemFindroomsFindRoomsResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemFindroomsFindRoomsGetResponseable instead.
type ItemFindroomsFindRoomsResponseable interface {
    ItemFindroomsFindRoomsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
