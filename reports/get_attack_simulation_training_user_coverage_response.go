package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetAttackSimulationTrainingUserCoverageGetResponseable instead.
type GetAttackSimulationTrainingUserCoverageResponse struct {
    GetAttackSimulationTrainingUserCoverageGetResponse
}
// NewGetAttackSimulationTrainingUserCoverageResponse instantiates a new GetAttackSimulationTrainingUserCoverageResponse and sets the default values.
func NewGetAttackSimulationTrainingUserCoverageResponse()(*GetAttackSimulationTrainingUserCoverageResponse) {
    m := &GetAttackSimulationTrainingUserCoverageResponse{
        GetAttackSimulationTrainingUserCoverageGetResponse: *NewGetAttackSimulationTrainingUserCoverageGetResponse(),
    }
    return m
}
// CreateGetAttackSimulationTrainingUserCoverageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetAttackSimulationTrainingUserCoverageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetAttackSimulationTrainingUserCoverageResponse(), nil
}
// Deprecated: This class is obsolete. Use GetAttackSimulationTrainingUserCoverageGetResponseable instead.
type GetAttackSimulationTrainingUserCoverageResponseable interface {
    GetAttackSimulationTrainingUserCoverageGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
