package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ZebrafotadeploymentsItemCancelPostResponseable instead.
type ZebrafotadeploymentsItemCancelResponse struct {
    ZebrafotadeploymentsItemCancelPostResponse
}
// NewZebrafotadeploymentsItemCancelResponse instantiates a new ZebrafotadeploymentsItemCancelResponse and sets the default values.
func NewZebrafotadeploymentsItemCancelResponse()(*ZebrafotadeploymentsItemCancelResponse) {
    m := &ZebrafotadeploymentsItemCancelResponse{
        ZebrafotadeploymentsItemCancelPostResponse: *NewZebrafotadeploymentsItemCancelPostResponse(),
    }
    return m
}
// CreateZebrafotadeploymentsItemCancelResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateZebrafotadeploymentsItemCancelResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebrafotadeploymentsItemCancelResponse(), nil
}
// Deprecated: This class is obsolete. Use ZebrafotadeploymentsItemCancelPostResponseable instead.
type ZebrafotadeploymentsItemCancelResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ZebrafotadeploymentsItemCancelPostResponseable
}
