package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ZebraFotaConnectorApproveFotaAppsResponse struct {
    ZebraFotaConnectorApproveFotaAppsPostResponse
}
// NewZebraFotaConnectorApproveFotaAppsResponse instantiates a new ZebraFotaConnectorApproveFotaAppsResponse and sets the default values.
func NewZebraFotaConnectorApproveFotaAppsResponse()(*ZebraFotaConnectorApproveFotaAppsResponse) {
    m := &ZebraFotaConnectorApproveFotaAppsResponse{
        ZebraFotaConnectorApproveFotaAppsPostResponse: *NewZebraFotaConnectorApproveFotaAppsPostResponse(),
    }
    return m
}
// CreateZebraFotaConnectorApproveFotaAppsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateZebraFotaConnectorApproveFotaAppsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaConnectorApproveFotaAppsResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ZebraFotaConnectorApproveFotaAppsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ZebraFotaConnectorApproveFotaAppsPostResponseable
}
