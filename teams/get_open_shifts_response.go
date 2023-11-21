package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetOpenShiftsResponse 
// Deprecated: This class is obsolete. Use getOpenShiftsGetResponse instead.
type GetOpenShiftsResponse struct {
    GetOpenShiftsGetResponse
}
// NewGetOpenShiftsResponse instantiates a new GetOpenShiftsResponse and sets the default values.
func NewGetOpenShiftsResponse()(*GetOpenShiftsResponse) {
    m := &GetOpenShiftsResponse{
        GetOpenShiftsGetResponse: *NewGetOpenShiftsGetResponse(),
    }
    return m
}
// CreateGetOpenShiftsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetOpenShiftsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetOpenShiftsResponse(), nil
}
// GetOpenShiftsResponseable 
// Deprecated: This class is obsolete. Use getOpenShiftsGetResponse instead.
type GetOpenShiftsResponseable interface {
    GetOpenShiftsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
