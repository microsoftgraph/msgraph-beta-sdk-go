package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemIsmanagedappuserblockedIsManagedAppUserBlockedGetResponseable instead.
type ItemIsmanagedappuserblockedIsManagedAppUserBlockedResponse struct {
    ItemIsmanagedappuserblockedIsManagedAppUserBlockedGetResponse
}
// NewItemIsmanagedappuserblockedIsManagedAppUserBlockedResponse instantiates a new ItemIsmanagedappuserblockedIsManagedAppUserBlockedResponse and sets the default values.
func NewItemIsmanagedappuserblockedIsManagedAppUserBlockedResponse()(*ItemIsmanagedappuserblockedIsManagedAppUserBlockedResponse) {
    m := &ItemIsmanagedappuserblockedIsManagedAppUserBlockedResponse{
        ItemIsmanagedappuserblockedIsManagedAppUserBlockedGetResponse: *NewItemIsmanagedappuserblockedIsManagedAppUserBlockedGetResponse(),
    }
    return m
}
// CreateItemIsmanagedappuserblockedIsManagedAppUserBlockedResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemIsmanagedappuserblockedIsManagedAppUserBlockedResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemIsmanagedappuserblockedIsManagedAppUserBlockedResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemIsmanagedappuserblockedIsManagedAppUserBlockedGetResponseable instead.
type ItemIsmanagedappuserblockedIsManagedAppUserBlockedResponseable interface {
    ItemIsmanagedappuserblockedIsManagedAppUserBlockedGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
