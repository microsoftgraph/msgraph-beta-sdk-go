package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedAppProtectionsHasPayloadLinksResponse 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type AndroidManagedAppProtectionsHasPayloadLinksResponse struct {
    AndroidManagedAppProtectionsHasPayloadLinksPostResponse
}
// NewAndroidManagedAppProtectionsHasPayloadLinksResponse instantiates a new AndroidManagedAppProtectionsHasPayloadLinksResponse and sets the default values.
func NewAndroidManagedAppProtectionsHasPayloadLinksResponse()(*AndroidManagedAppProtectionsHasPayloadLinksResponse) {
    m := &AndroidManagedAppProtectionsHasPayloadLinksResponse{
        AndroidManagedAppProtectionsHasPayloadLinksPostResponse: *NewAndroidManagedAppProtectionsHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateAndroidManagedAppProtectionsHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedAppProtectionsHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedAppProtectionsHasPayloadLinksResponse(), nil
}
// AndroidManagedAppProtectionsHasPayloadLinksResponseable 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type AndroidManagedAppProtectionsHasPayloadLinksResponseable interface {
    AndroidManagedAppProtectionsHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
