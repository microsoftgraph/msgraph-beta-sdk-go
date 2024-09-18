package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCloudPCsItemRetrieveSnapshotsGetResponseable instead.
type ItemCloudPCsItemRetrieveSnapshotsResponse struct {
    ItemCloudPCsItemRetrieveSnapshotsGetResponse
}
// NewItemCloudPCsItemRetrieveSnapshotsResponse instantiates a new ItemCloudPCsItemRetrieveSnapshotsResponse and sets the default values.
func NewItemCloudPCsItemRetrieveSnapshotsResponse()(*ItemCloudPCsItemRetrieveSnapshotsResponse) {
    m := &ItemCloudPCsItemRetrieveSnapshotsResponse{
        ItemCloudPCsItemRetrieveSnapshotsGetResponse: *NewItemCloudPCsItemRetrieveSnapshotsGetResponse(),
    }
    return m
}
// CreateItemCloudPCsItemRetrieveSnapshotsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCloudPCsItemRetrieveSnapshotsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCloudPCsItemRetrieveSnapshotsResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCloudPCsItemRetrieveSnapshotsGetResponseable instead.
type ItemCloudPCsItemRetrieveSnapshotsResponseable interface {
    ItemCloudPCsItemRetrieveSnapshotsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
