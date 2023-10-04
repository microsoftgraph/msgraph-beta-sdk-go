package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetAttackSimulationRepeatOffendersResponse 
// Deprecated: This class is obsolete. Use getAttackSimulationRepeatOffendersGetResponse instead.
type GetAttackSimulationRepeatOffendersResponse struct {
    GetAttackSimulationRepeatOffendersGetResponse
}
// NewGetAttackSimulationRepeatOffendersResponse instantiates a new GetAttackSimulationRepeatOffendersResponse and sets the default values.
func NewGetAttackSimulationRepeatOffendersResponse()(*GetAttackSimulationRepeatOffendersResponse) {
    m := &GetAttackSimulationRepeatOffendersResponse{
        GetAttackSimulationRepeatOffendersGetResponse: *NewGetAttackSimulationRepeatOffendersGetResponse(),
    }
    return m
}
// CreateGetAttackSimulationRepeatOffendersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetAttackSimulationRepeatOffendersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetAttackSimulationRepeatOffendersResponse(), nil
}
// GetAttackSimulationRepeatOffendersResponseable 
// Deprecated: This class is obsolete. Use getAttackSimulationRepeatOffendersGetResponse instead.
type GetAttackSimulationRepeatOffendersResponseable interface {
    GetAttackSimulationRepeatOffendersGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
