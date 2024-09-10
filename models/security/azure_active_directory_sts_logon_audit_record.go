package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AzureActiveDirectoryStsLogonAuditRecord struct {
    AuditData
}
// NewAzureActiveDirectoryStsLogonAuditRecord instantiates a new AzureActiveDirectoryStsLogonAuditRecord and sets the default values.
func NewAzureActiveDirectoryStsLogonAuditRecord()(*AzureActiveDirectoryStsLogonAuditRecord) {
    m := &AzureActiveDirectoryStsLogonAuditRecord{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.azureActiveDirectoryStsLogonAuditRecord"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureActiveDirectoryStsLogonAuditRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAzureActiveDirectoryStsLogonAuditRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureActiveDirectoryStsLogonAuditRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AzureActiveDirectoryStsLogonAuditRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AzureActiveDirectoryStsLogonAuditRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type AzureActiveDirectoryStsLogonAuditRecordable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
