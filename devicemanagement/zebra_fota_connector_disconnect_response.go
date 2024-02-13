package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ZebraFotaConnectorDisconnectResponse struct {
    ZebraFotaConnectorDisconnectPostResponse
}
// NewZebraFotaConnectorDisconnectResponse instantiates a new ZebraFotaConnectorDisconnectResponse and sets the default values.
func NewZebraFotaConnectorDisconnectResponse()(*ZebraFotaConnectorDisconnectResponse) {
    m := &ZebraFotaConnectorDisconnectResponse{
        ZebraFotaConnectorDisconnectPostResponse: *NewZebraFotaConnectorDisconnectPostResponse(),
    }
    return m
}
// CreateZebraFotaConnectorDisconnectResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateZebraFotaConnectorDisconnectResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaConnectorDisconnectResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ZebraFotaConnectorDisconnectResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ZebraFotaConnectorDisconnectPostResponseable
}
