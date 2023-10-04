package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEndpointGetEffectivePermissionsResponse 
// Deprecated: This class is obsolete. Use getEffectivePermissionsGetResponse instead.
type VirtualEndpointGetEffectivePermissionsResponse struct {
    VirtualEndpointGetEffectivePermissionsGetResponse
}
// NewVirtualEndpointGetEffectivePermissionsResponse instantiates a new VirtualEndpointGetEffectivePermissionsResponse and sets the default values.
func NewVirtualEndpointGetEffectivePermissionsResponse()(*VirtualEndpointGetEffectivePermissionsResponse) {
    m := &VirtualEndpointGetEffectivePermissionsResponse{
        VirtualEndpointGetEffectivePermissionsGetResponse: *NewVirtualEndpointGetEffectivePermissionsGetResponse(),
    }
    return m
}
// CreateVirtualEndpointGetEffectivePermissionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEndpointGetEffectivePermissionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpointGetEffectivePermissionsResponse(), nil
}
// VirtualEndpointGetEffectivePermissionsResponseable 
// Deprecated: This class is obsolete. Use getEffectivePermissionsGetResponse instead.
type VirtualEndpointGetEffectivePermissionsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEndpointGetEffectivePermissionsGetResponseable
}
