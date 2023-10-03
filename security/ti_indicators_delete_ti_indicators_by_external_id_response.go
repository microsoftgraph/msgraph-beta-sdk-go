package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TiIndicatorsDeleteTiIndicatorsByExternalIdResponse 
// Deprecated: This class is obsolete. Use deleteTiIndicatorsByExternalIdPostResponse instead.
type TiIndicatorsDeleteTiIndicatorsByExternalIdResponse struct {
    TiIndicatorsDeleteTiIndicatorsByExternalIdPostResponse
}
// NewTiIndicatorsDeleteTiIndicatorsByExternalIdResponse instantiates a new TiIndicatorsDeleteTiIndicatorsByExternalIdResponse and sets the default values.
func NewTiIndicatorsDeleteTiIndicatorsByExternalIdResponse()(*TiIndicatorsDeleteTiIndicatorsByExternalIdResponse) {
    m := &TiIndicatorsDeleteTiIndicatorsByExternalIdResponse{
        TiIndicatorsDeleteTiIndicatorsByExternalIdPostResponse: *NewTiIndicatorsDeleteTiIndicatorsByExternalIdPostResponse(),
    }
    return m
}
// CreateTiIndicatorsDeleteTiIndicatorsByExternalIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTiIndicatorsDeleteTiIndicatorsByExternalIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTiIndicatorsDeleteTiIndicatorsByExternalIdResponse(), nil
}
// TiIndicatorsDeleteTiIndicatorsByExternalIdResponseable 
// Deprecated: This class is obsolete. Use deleteTiIndicatorsByExternalIdPostResponse instead.
type TiIndicatorsDeleteTiIndicatorsByExternalIdResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TiIndicatorsDeleteTiIndicatorsByExternalIdPostResponseable
}
