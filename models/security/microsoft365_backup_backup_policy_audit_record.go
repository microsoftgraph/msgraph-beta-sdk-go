package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Microsoft365BackupBackupPolicyAuditRecord struct {
    AuditData
}
// NewMicrosoft365BackupBackupPolicyAuditRecord instantiates a new Microsoft365BackupBackupPolicyAuditRecord and sets the default values.
func NewMicrosoft365BackupBackupPolicyAuditRecord()(*Microsoft365BackupBackupPolicyAuditRecord) {
    m := &Microsoft365BackupBackupPolicyAuditRecord{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.microsoft365BackupBackupPolicyAuditRecord"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMicrosoft365BackupBackupPolicyAuditRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMicrosoft365BackupBackupPolicyAuditRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoft365BackupBackupPolicyAuditRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Microsoft365BackupBackupPolicyAuditRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *Microsoft365BackupBackupPolicyAuditRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type Microsoft365BackupBackupPolicyAuditRecordable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
