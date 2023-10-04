package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TiIndicatorsDeleteTiIndicatorsResponse 
// Deprecated: This class is obsolete. Use deleteTiIndicatorsPostResponse instead.
type TiIndicatorsDeleteTiIndicatorsResponse struct {
    TiIndicatorsDeleteTiIndicatorsPostResponse
}
// NewTiIndicatorsDeleteTiIndicatorsResponse instantiates a new TiIndicatorsDeleteTiIndicatorsResponse and sets the default values.
func NewTiIndicatorsDeleteTiIndicatorsResponse()(*TiIndicatorsDeleteTiIndicatorsResponse) {
    m := &TiIndicatorsDeleteTiIndicatorsResponse{
        TiIndicatorsDeleteTiIndicatorsPostResponse: *NewTiIndicatorsDeleteTiIndicatorsPostResponse(),
    }
    return m
}
// CreateTiIndicatorsDeleteTiIndicatorsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTiIndicatorsDeleteTiIndicatorsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTiIndicatorsDeleteTiIndicatorsResponse(), nil
}
// TiIndicatorsDeleteTiIndicatorsResponseable 
// Deprecated: This class is obsolete. Use deleteTiIndicatorsPostResponse instead.
type TiIndicatorsDeleteTiIndicatorsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TiIndicatorsDeleteTiIndicatorsPostResponseable
}
