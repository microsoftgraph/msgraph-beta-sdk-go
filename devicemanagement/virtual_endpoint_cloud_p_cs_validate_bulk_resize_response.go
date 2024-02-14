package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type VirtualEndpointCloudPCsValidateBulkResizeResponse struct {
    VirtualEndpointCloudPCsValidateBulkResizePostResponse
}
// NewVirtualEndpointCloudPCsValidateBulkResizeResponse instantiates a new VirtualEndpointCloudPCsValidateBulkResizeResponse and sets the default values.
func NewVirtualEndpointCloudPCsValidateBulkResizeResponse()(*VirtualEndpointCloudPCsValidateBulkResizeResponse) {
    m := &VirtualEndpointCloudPCsValidateBulkResizeResponse{
        VirtualEndpointCloudPCsValidateBulkResizePostResponse: *NewVirtualEndpointCloudPCsValidateBulkResizePostResponse(),
    }
    return m
}
// CreateVirtualEndpointCloudPCsValidateBulkResizeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEndpointCloudPCsValidateBulkResizeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpointCloudPCsValidateBulkResizeResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type VirtualEndpointCloudPCsValidateBulkResizeResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEndpointCloudPCsValidateBulkResizePostResponseable
}
