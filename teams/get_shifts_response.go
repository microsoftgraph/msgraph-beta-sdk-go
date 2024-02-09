package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type GetShiftsResponse struct {
    GetShiftsGetResponse
}
// NewGetShiftsResponse instantiates a new GetShiftsResponse and sets the default values.
func NewGetShiftsResponse()(*GetShiftsResponse) {
    m := &GetShiftsResponse{
        GetShiftsGetResponse: *NewGetShiftsGetResponse(),
    }
    return m
}
// CreateGetShiftsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetShiftsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetShiftsResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type GetShiftsResponseable interface {
    GetShiftsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
