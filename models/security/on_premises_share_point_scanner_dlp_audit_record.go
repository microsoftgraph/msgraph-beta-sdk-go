package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OnPremisesSharePointScannerDlpAuditRecord struct {
    AuditData
}
// NewOnPremisesSharePointScannerDlpAuditRecord instantiates a new OnPremisesSharePointScannerDlpAuditRecord and sets the default values.
func NewOnPremisesSharePointScannerDlpAuditRecord()(*OnPremisesSharePointScannerDlpAuditRecord) {
    m := &OnPremisesSharePointScannerDlpAuditRecord{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.onPremisesSharePointScannerDlpAuditRecord"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnPremisesSharePointScannerDlpAuditRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnPremisesSharePointScannerDlpAuditRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesSharePointScannerDlpAuditRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OnPremisesSharePointScannerDlpAuditRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OnPremisesSharePointScannerDlpAuditRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type OnPremisesSharePointScannerDlpAuditRecordable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
