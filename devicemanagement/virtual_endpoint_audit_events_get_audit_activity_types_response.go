package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEndpointAuditEventsGetAuditActivityTypesResponse 
// Deprecated: This class is obsolete. Use getAuditActivityTypesGetResponse instead.
type VirtualEndpointAuditEventsGetAuditActivityTypesResponse struct {
    VirtualEndpointAuditEventsGetAuditActivityTypesGetResponse
}
// NewVirtualEndpointAuditEventsGetAuditActivityTypesResponse instantiates a new VirtualEndpointAuditEventsGetAuditActivityTypesResponse and sets the default values.
func NewVirtualEndpointAuditEventsGetAuditActivityTypesResponse()(*VirtualEndpointAuditEventsGetAuditActivityTypesResponse) {
    m := &VirtualEndpointAuditEventsGetAuditActivityTypesResponse{
        VirtualEndpointAuditEventsGetAuditActivityTypesGetResponse: *NewVirtualEndpointAuditEventsGetAuditActivityTypesGetResponse(),
    }
    return m
}
// CreateVirtualEndpointAuditEventsGetAuditActivityTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEndpointAuditEventsGetAuditActivityTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpointAuditEventsGetAuditActivityTypesResponse(), nil
}
// VirtualEndpointAuditEventsGetAuditActivityTypesResponseable 
// Deprecated: This class is obsolete. Use getAuditActivityTypesGetResponse instead.
type VirtualEndpointAuditEventsGetAuditActivityTypesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEndpointAuditEventsGetAuditActivityTypesGetResponseable
}
