package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use VirtualEndpointCloudPCsItemRetrieveSnapshotsGetResponseable instead.
type VirtualEndpointCloudPCsItemRetrieveSnapshotsResponse struct {
    VirtualEndpointCloudPCsItemRetrieveSnapshotsGetResponse
}
// NewVirtualEndpointCloudPCsItemRetrieveSnapshotsResponse instantiates a new VirtualEndpointCloudPCsItemRetrieveSnapshotsResponse and sets the default values.
func NewVirtualEndpointCloudPCsItemRetrieveSnapshotsResponse()(*VirtualEndpointCloudPCsItemRetrieveSnapshotsResponse) {
    m := &VirtualEndpointCloudPCsItemRetrieveSnapshotsResponse{
        VirtualEndpointCloudPCsItemRetrieveSnapshotsGetResponse: *NewVirtualEndpointCloudPCsItemRetrieveSnapshotsGetResponse(),
    }
    return m
}
// CreateVirtualEndpointCloudPCsItemRetrieveSnapshotsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEndpointCloudPCsItemRetrieveSnapshotsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpointCloudPCsItemRetrieveSnapshotsResponse(), nil
}
// Deprecated: This class is obsolete. Use VirtualEndpointCloudPCsItemRetrieveSnapshotsGetResponseable instead.
type VirtualEndpointCloudPCsItemRetrieveSnapshotsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEndpointCloudPCsItemRetrieveSnapshotsGetResponseable
}
