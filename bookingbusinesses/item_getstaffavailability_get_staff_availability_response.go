package bookingbusinesses

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemGetstaffavailabilityGetStaffAvailabilityPostResponseable instead.
type ItemGetstaffavailabilityGetStaffAvailabilityResponse struct {
    ItemGetstaffavailabilityGetStaffAvailabilityPostResponse
}
// NewItemGetstaffavailabilityGetStaffAvailabilityResponse instantiates a new ItemGetstaffavailabilityGetStaffAvailabilityResponse and sets the default values.
func NewItemGetstaffavailabilityGetStaffAvailabilityResponse()(*ItemGetstaffavailabilityGetStaffAvailabilityResponse) {
    m := &ItemGetstaffavailabilityGetStaffAvailabilityResponse{
        ItemGetstaffavailabilityGetStaffAvailabilityPostResponse: *NewItemGetstaffavailabilityGetStaffAvailabilityPostResponse(),
    }
    return m
}
// CreateItemGetstaffavailabilityGetStaffAvailabilityResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemGetstaffavailabilityGetStaffAvailabilityResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetstaffavailabilityGetStaffAvailabilityResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemGetstaffavailabilityGetStaffAvailabilityPostResponseable instead.
type ItemGetstaffavailabilityGetStaffAvailabilityResponseable interface {
    ItemGetstaffavailabilityGetStaffAvailabilityPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
