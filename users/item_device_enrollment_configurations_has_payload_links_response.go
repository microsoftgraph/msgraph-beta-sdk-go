package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemDeviceEnrollmentConfigurationsHasPayloadLinksResponse 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type ItemDeviceEnrollmentConfigurationsHasPayloadLinksResponse struct {
    ItemDeviceEnrollmentConfigurationsHasPayloadLinksPostResponse
}
// NewItemDeviceEnrollmentConfigurationsHasPayloadLinksResponse instantiates a new ItemDeviceEnrollmentConfigurationsHasPayloadLinksResponse and sets the default values.
func NewItemDeviceEnrollmentConfigurationsHasPayloadLinksResponse()(*ItemDeviceEnrollmentConfigurationsHasPayloadLinksResponse) {
    m := &ItemDeviceEnrollmentConfigurationsHasPayloadLinksResponse{
        ItemDeviceEnrollmentConfigurationsHasPayloadLinksPostResponse: *NewItemDeviceEnrollmentConfigurationsHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateItemDeviceEnrollmentConfigurationsHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemDeviceEnrollmentConfigurationsHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDeviceEnrollmentConfigurationsHasPayloadLinksResponse(), nil
}
// ItemDeviceEnrollmentConfigurationsHasPayloadLinksResponseable 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type ItemDeviceEnrollmentConfigurationsHasPayloadLinksResponseable interface {
    ItemDeviceEnrollmentConfigurationsHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
