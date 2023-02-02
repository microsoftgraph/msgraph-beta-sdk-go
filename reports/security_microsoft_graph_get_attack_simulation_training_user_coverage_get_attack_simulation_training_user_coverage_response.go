package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponse 
type SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
    // The value property
    value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttackSimulationTrainingUserCoverageable
}
// NewSecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponse instantiates a new SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponse and sets the default values.
func NewSecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponse()(*SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponse) {
    m := &SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttackSimulationTrainingUserCoverageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttackSimulationTrainingUserCoverageable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttackSimulationTrainingUserCoverageable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponse) GetValue()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttackSimulationTrainingUserCoverageable) {
    return m.value
}
// Serialize serializes information the current object
func (m *SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponse) SetValue(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttackSimulationTrainingUserCoverageable)() {
    m.value = value
}
