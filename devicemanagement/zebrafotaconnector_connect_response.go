package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ZebrafotaconnectorConnectPostResponseable instead.
type ZebrafotaconnectorConnectResponse struct {
    ZebrafotaconnectorConnectPostResponse
}
// NewZebrafotaconnectorConnectResponse instantiates a new ZebrafotaconnectorConnectResponse and sets the default values.
func NewZebrafotaconnectorConnectResponse()(*ZebrafotaconnectorConnectResponse) {
    m := &ZebrafotaconnectorConnectResponse{
        ZebrafotaconnectorConnectPostResponse: *NewZebrafotaconnectorConnectPostResponse(),
    }
    return m
}
// CreateZebrafotaconnectorConnectResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateZebrafotaconnectorConnectResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebrafotaconnectorConnectResponse(), nil
}
// Deprecated: This class is obsolete. Use ZebrafotaconnectorConnectPostResponseable instead.
type ZebrafotaconnectorConnectResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ZebrafotaconnectorConnectPostResponseable
}
