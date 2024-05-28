package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use VirtualendpointCloudpcsBulkresizeBulkResizePostResponseable instead.
type VirtualendpointCloudpcsBulkresizeBulkResizeResponse struct {
    VirtualendpointCloudpcsBulkresizeBulkResizePostResponse
}
// NewVirtualendpointCloudpcsBulkresizeBulkResizeResponse instantiates a new VirtualendpointCloudpcsBulkresizeBulkResizeResponse and sets the default values.
func NewVirtualendpointCloudpcsBulkresizeBulkResizeResponse()(*VirtualendpointCloudpcsBulkresizeBulkResizeResponse) {
    m := &VirtualendpointCloudpcsBulkresizeBulkResizeResponse{
        VirtualendpointCloudpcsBulkresizeBulkResizePostResponse: *NewVirtualendpointCloudpcsBulkresizeBulkResizePostResponse(),
    }
    return m
}
// CreateVirtualendpointCloudpcsBulkresizeBulkResizeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualendpointCloudpcsBulkresizeBulkResizeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualendpointCloudpcsBulkresizeBulkResizeResponse(), nil
}
// Deprecated: This class is obsolete. Use VirtualendpointCloudpcsBulkresizeBulkResizePostResponseable instead.
type VirtualendpointCloudpcsBulkresizeBulkResizeResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualendpointCloudpcsBulkresizeBulkResizePostResponseable
}
