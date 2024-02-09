package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type IosManagedAppProtectionsHasPayloadLinksResponse struct {
    IosManagedAppProtectionsHasPayloadLinksPostResponse
}
// NewIosManagedAppProtectionsHasPayloadLinksResponse instantiates a new IosManagedAppProtectionsHasPayloadLinksResponse and sets the default values.
func NewIosManagedAppProtectionsHasPayloadLinksResponse()(*IosManagedAppProtectionsHasPayloadLinksResponse) {
    m := &IosManagedAppProtectionsHasPayloadLinksResponse{
        IosManagedAppProtectionsHasPayloadLinksPostResponse: *NewIosManagedAppProtectionsHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateIosManagedAppProtectionsHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIosManagedAppProtectionsHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosManagedAppProtectionsHasPayloadLinksResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type IosManagedAppProtectionsHasPayloadLinksResponseable interface {
    IosManagedAppProtectionsHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
