package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSitesItemContentModelsItemGetAppliedDrivesGetResponseable instead.
type ItemSitesItemContentModelsItemGetAppliedDrivesResponse struct {
    ItemSitesItemContentModelsItemGetAppliedDrivesGetResponse
}
// NewItemSitesItemContentModelsItemGetAppliedDrivesResponse instantiates a new ItemSitesItemContentModelsItemGetAppliedDrivesResponse and sets the default values.
func NewItemSitesItemContentModelsItemGetAppliedDrivesResponse()(*ItemSitesItemContentModelsItemGetAppliedDrivesResponse) {
    m := &ItemSitesItemContentModelsItemGetAppliedDrivesResponse{
        ItemSitesItemContentModelsItemGetAppliedDrivesGetResponse: *NewItemSitesItemContentModelsItemGetAppliedDrivesGetResponse(),
    }
    return m
}
// CreateItemSitesItemContentModelsItemGetAppliedDrivesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSitesItemContentModelsItemGetAppliedDrivesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemContentModelsItemGetAppliedDrivesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSitesItemContentModelsItemGetAppliedDrivesGetResponseable instead.
type ItemSitesItemContentModelsItemGetAppliedDrivesResponseable interface {
    ItemSitesItemContentModelsItemGetAppliedDrivesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
