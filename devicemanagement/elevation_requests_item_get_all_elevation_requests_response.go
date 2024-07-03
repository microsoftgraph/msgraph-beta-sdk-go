package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ElevationRequestsItemGetAllElevationRequestsPostResponseable instead.
type ElevationRequestsItemGetAllElevationRequestsResponse struct {
    ElevationRequestsItemGetAllElevationRequestsPostResponse
}
// NewElevationRequestsItemGetAllElevationRequestsResponse instantiates a new ElevationRequestsItemGetAllElevationRequestsResponse and sets the default values.
func NewElevationRequestsItemGetAllElevationRequestsResponse()(*ElevationRequestsItemGetAllElevationRequestsResponse) {
    m := &ElevationRequestsItemGetAllElevationRequestsResponse{
        ElevationRequestsItemGetAllElevationRequestsPostResponse: *NewElevationRequestsItemGetAllElevationRequestsPostResponse(),
    }
    return m
}
// CreateElevationRequestsItemGetAllElevationRequestsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateElevationRequestsItemGetAllElevationRequestsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewElevationRequestsItemGetAllElevationRequestsResponse(), nil
}
// Deprecated: This class is obsolete. Use ElevationRequestsItemGetAllElevationRequestsPostResponseable instead.
type ElevationRequestsItemGetAllElevationRequestsResponseable interface {
    ElevationRequestsItemGetAllElevationRequestsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
