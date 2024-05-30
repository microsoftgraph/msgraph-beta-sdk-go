package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ZebrafotaconnectorDisconnectPostResponseable instead.
type ZebrafotaconnectorDisconnectResponse struct {
    ZebrafotaconnectorDisconnectPostResponse
}
// NewZebrafotaconnectorDisconnectResponse instantiates a new ZebrafotaconnectorDisconnectResponse and sets the default values.
func NewZebrafotaconnectorDisconnectResponse()(*ZebrafotaconnectorDisconnectResponse) {
    m := &ZebrafotaconnectorDisconnectResponse{
        ZebrafotaconnectorDisconnectPostResponse: *NewZebrafotaconnectorDisconnectPostResponse(),
    }
    return m
}
// CreateZebrafotaconnectorDisconnectResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateZebrafotaconnectorDisconnectResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebrafotaconnectorDisconnectResponse(), nil
}
// Deprecated: This class is obsolete. Use ZebrafotaconnectorDisconnectPostResponseable instead.
type ZebrafotaconnectorDisconnectResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ZebrafotaconnectorDisconnectPostResponseable
}
