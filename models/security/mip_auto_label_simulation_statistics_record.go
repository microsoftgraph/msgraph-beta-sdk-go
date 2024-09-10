package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MipAutoLabelSimulationStatisticsRecord struct {
    AuditData
}
// NewMipAutoLabelSimulationStatisticsRecord instantiates a new MipAutoLabelSimulationStatisticsRecord and sets the default values.
func NewMipAutoLabelSimulationStatisticsRecord()(*MipAutoLabelSimulationStatisticsRecord) {
    m := &MipAutoLabelSimulationStatisticsRecord{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.mipAutoLabelSimulationStatisticsRecord"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMipAutoLabelSimulationStatisticsRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMipAutoLabelSimulationStatisticsRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMipAutoLabelSimulationStatisticsRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MipAutoLabelSimulationStatisticsRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MipAutoLabelSimulationStatisticsRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type MipAutoLabelSimulationStatisticsRecordable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
