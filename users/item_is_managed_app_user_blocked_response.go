package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemIsManagedAppUserBlockedGetResponseable instead.
type ItemIsManagedAppUserBlockedResponse struct {
    ItemIsManagedAppUserBlockedGetResponse
}
// NewItemIsManagedAppUserBlockedResponse instantiates a new ItemIsManagedAppUserBlockedResponse and sets the default values.
func NewItemIsManagedAppUserBlockedResponse()(*ItemIsManagedAppUserBlockedResponse) {
    m := &ItemIsManagedAppUserBlockedResponse{
        ItemIsManagedAppUserBlockedGetResponse: *NewItemIsManagedAppUserBlockedGetResponse(),
    }
    return m
}
// CreateItemIsManagedAppUserBlockedResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemIsManagedAppUserBlockedResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemIsManagedAppUserBlockedResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemIsManagedAppUserBlockedGetResponseable instead.
type ItemIsManagedAppUserBlockedResponseable interface {
    ItemIsManagedAppUserBlockedGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
