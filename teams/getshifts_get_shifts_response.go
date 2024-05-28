package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetshiftsGetShiftsGetResponseable instead.
type GetshiftsGetShiftsResponse struct {
    GetshiftsGetShiftsGetResponse
}
// NewGetshiftsGetShiftsResponse instantiates a new GetshiftsGetShiftsResponse and sets the default values.
func NewGetshiftsGetShiftsResponse()(*GetshiftsGetShiftsResponse) {
    m := &GetshiftsGetShiftsResponse{
        GetshiftsGetShiftsGetResponse: *NewGetshiftsGetShiftsGetResponse(),
    }
    return m
}
// CreateGetshiftsGetShiftsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetshiftsGetShiftsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetshiftsGetShiftsResponse(), nil
}
// Deprecated: This class is obsolete. Use GetshiftsGetShiftsGetResponseable instead.
type GetshiftsGetShiftsResponseable interface {
    GetshiftsGetShiftsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
