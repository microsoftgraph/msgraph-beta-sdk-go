package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TiIndicatorsUpdateTiIndicatorsResponse 
// Deprecated: This class is obsolete. Use updateTiIndicatorsPostResponse instead.
type TiIndicatorsUpdateTiIndicatorsResponse struct {
    TiIndicatorsUpdateTiIndicatorsPostResponse
}
// NewTiIndicatorsUpdateTiIndicatorsResponse instantiates a new TiIndicatorsUpdateTiIndicatorsResponse and sets the default values.
func NewTiIndicatorsUpdateTiIndicatorsResponse()(*TiIndicatorsUpdateTiIndicatorsResponse) {
    m := &TiIndicatorsUpdateTiIndicatorsResponse{
        TiIndicatorsUpdateTiIndicatorsPostResponse: *NewTiIndicatorsUpdateTiIndicatorsPostResponse(),
    }
    return m
}
// CreateTiIndicatorsUpdateTiIndicatorsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTiIndicatorsUpdateTiIndicatorsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTiIndicatorsUpdateTiIndicatorsResponse(), nil
}
// TiIndicatorsUpdateTiIndicatorsResponseable 
// Deprecated: This class is obsolete. Use updateTiIndicatorsPostResponse instead.
type TiIndicatorsUpdateTiIndicatorsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TiIndicatorsUpdateTiIndicatorsPostResponseable
}
