package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetAttackSimulationSimulationUserCoverageResponse 
// Deprecated: This class is obsolete. Use getAttackSimulationSimulationUserCoverageGetResponse instead.
type GetAttackSimulationSimulationUserCoverageResponse struct {
    GetAttackSimulationSimulationUserCoverageGetResponse
}
// NewGetAttackSimulationSimulationUserCoverageResponse instantiates a new GetAttackSimulationSimulationUserCoverageResponse and sets the default values.
func NewGetAttackSimulationSimulationUserCoverageResponse()(*GetAttackSimulationSimulationUserCoverageResponse) {
    m := &GetAttackSimulationSimulationUserCoverageResponse{
        GetAttackSimulationSimulationUserCoverageGetResponse: *NewGetAttackSimulationSimulationUserCoverageGetResponse(),
    }
    return m
}
// CreateGetAttackSimulationSimulationUserCoverageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetAttackSimulationSimulationUserCoverageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetAttackSimulationSimulationUserCoverageResponse(), nil
}
// GetAttackSimulationSimulationUserCoverageResponseable 
// Deprecated: This class is obsolete. Use getAttackSimulationSimulationUserCoverageGetResponse instead.
type GetAttackSimulationSimulationUserCoverageResponseable interface {
    GetAttackSimulationSimulationUserCoverageGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
