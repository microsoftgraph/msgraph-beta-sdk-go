package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemContentModelsItemGetAppliedDrivesGetResponseable instead.
type ItemContentModelsItemGetAppliedDrivesResponse struct {
    ItemContentModelsItemGetAppliedDrivesGetResponse
}
// NewItemContentModelsItemGetAppliedDrivesResponse instantiates a new ItemContentModelsItemGetAppliedDrivesResponse and sets the default values.
func NewItemContentModelsItemGetAppliedDrivesResponse()(*ItemContentModelsItemGetAppliedDrivesResponse) {
    m := &ItemContentModelsItemGetAppliedDrivesResponse{
        ItemContentModelsItemGetAppliedDrivesGetResponse: *NewItemContentModelsItemGetAppliedDrivesGetResponse(),
    }
    return m
}
// CreateItemContentModelsItemGetAppliedDrivesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemContentModelsItemGetAppliedDrivesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContentModelsItemGetAppliedDrivesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemContentModelsItemGetAppliedDrivesGetResponseable instead.
type ItemContentModelsItemGetAppliedDrivesResponseable interface {
    ItemContentModelsItemGetAppliedDrivesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
