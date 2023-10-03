package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaConnectorHasActiveDeploymentsResponse 
// Deprecated: This class is obsolete. Use hasActiveDeploymentsPostResponse instead.
type ZebraFotaConnectorHasActiveDeploymentsResponse struct {
    ZebraFotaConnectorHasActiveDeploymentsPostResponse
}
// NewZebraFotaConnectorHasActiveDeploymentsResponse instantiates a new ZebraFotaConnectorHasActiveDeploymentsResponse and sets the default values.
func NewZebraFotaConnectorHasActiveDeploymentsResponse()(*ZebraFotaConnectorHasActiveDeploymentsResponse) {
    m := &ZebraFotaConnectorHasActiveDeploymentsResponse{
        ZebraFotaConnectorHasActiveDeploymentsPostResponse: *NewZebraFotaConnectorHasActiveDeploymentsPostResponse(),
    }
    return m
}
// CreateZebraFotaConnectorHasActiveDeploymentsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateZebraFotaConnectorHasActiveDeploymentsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaConnectorHasActiveDeploymentsResponse(), nil
}
// ZebraFotaConnectorHasActiveDeploymentsResponseable 
// Deprecated: This class is obsolete. Use hasActiveDeploymentsPostResponse instead.
type ZebraFotaConnectorHasActiveDeploymentsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ZebraFotaConnectorHasActiveDeploymentsPostResponseable
}
