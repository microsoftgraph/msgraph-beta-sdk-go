package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use MobileappsHaspayloadlinksHasPayloadLinksPostResponseable instead.
type MobileappsHaspayloadlinksHasPayloadLinksResponse struct {
    MobileappsHaspayloadlinksHasPayloadLinksPostResponse
}
// NewMobileappsHaspayloadlinksHasPayloadLinksResponse instantiates a new MobileappsHaspayloadlinksHasPayloadLinksResponse and sets the default values.
func NewMobileappsHaspayloadlinksHasPayloadLinksResponse()(*MobileappsHaspayloadlinksHasPayloadLinksResponse) {
    m := &MobileappsHaspayloadlinksHasPayloadLinksResponse{
        MobileappsHaspayloadlinksHasPayloadLinksPostResponse: *NewMobileappsHaspayloadlinksHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateMobileappsHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMobileappsHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileappsHaspayloadlinksHasPayloadLinksResponse(), nil
}
// Deprecated: This class is obsolete. Use MobileappsHaspayloadlinksHasPayloadLinksPostResponseable instead.
type MobileappsHaspayloadlinksHasPayloadLinksResponseable interface {
    MobileappsHaspayloadlinksHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
