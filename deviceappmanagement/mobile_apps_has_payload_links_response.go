package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppsHasPayloadLinksResponse 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type MobileAppsHasPayloadLinksResponse struct {
    MobileAppsHasPayloadLinksPostResponse
}
// NewMobileAppsHasPayloadLinksResponse instantiates a new MobileAppsHasPayloadLinksResponse and sets the default values.
func NewMobileAppsHasPayloadLinksResponse()(*MobileAppsHasPayloadLinksResponse) {
    m := &MobileAppsHasPayloadLinksResponse{
        MobileAppsHasPayloadLinksPostResponse: *NewMobileAppsHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateMobileAppsHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppsHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppsHasPayloadLinksResponse(), nil
}
// MobileAppsHasPayloadLinksResponseable 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type MobileAppsHasPayloadLinksResponseable interface {
    MobileAppsHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
