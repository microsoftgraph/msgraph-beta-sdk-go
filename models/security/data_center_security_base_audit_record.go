package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DataCenterSecurityBaseAuditRecord struct {
    AuditData
}
// NewDataCenterSecurityBaseAuditRecord instantiates a new DataCenterSecurityBaseAuditRecord and sets the default values.
func NewDataCenterSecurityBaseAuditRecord()(*DataCenterSecurityBaseAuditRecord) {
    m := &DataCenterSecurityBaseAuditRecord{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.dataCenterSecurityBaseAuditRecord"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDataCenterSecurityBaseAuditRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDataCenterSecurityBaseAuditRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataCenterSecurityBaseAuditRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DataCenterSecurityBaseAuditRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DataCenterSecurityBaseAuditRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type DataCenterSecurityBaseAuditRecordable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
