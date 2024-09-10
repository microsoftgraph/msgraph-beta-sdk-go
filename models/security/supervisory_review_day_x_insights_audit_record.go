package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SupervisoryReviewDayXInsightsAuditRecord struct {
    AuditData
}
// NewSupervisoryReviewDayXInsightsAuditRecord instantiates a new SupervisoryReviewDayXInsightsAuditRecord and sets the default values.
func NewSupervisoryReviewDayXInsightsAuditRecord()(*SupervisoryReviewDayXInsightsAuditRecord) {
    m := &SupervisoryReviewDayXInsightsAuditRecord{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.supervisoryReviewDayXInsightsAuditRecord"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSupervisoryReviewDayXInsightsAuditRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSupervisoryReviewDayXInsightsAuditRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSupervisoryReviewDayXInsightsAuditRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SupervisoryReviewDayXInsightsAuditRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SupervisoryReviewDayXInsightsAuditRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type SupervisoryReviewDayXInsightsAuditRecordable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
