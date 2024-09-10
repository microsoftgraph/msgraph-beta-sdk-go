package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PowerPlatformAdminEnvironmentAuditRecord struct {
    AuditData
}
// NewPowerPlatformAdminEnvironmentAuditRecord instantiates a new PowerPlatformAdminEnvironmentAuditRecord and sets the default values.
func NewPowerPlatformAdminEnvironmentAuditRecord()(*PowerPlatformAdminEnvironmentAuditRecord) {
    m := &PowerPlatformAdminEnvironmentAuditRecord{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.powerPlatformAdminEnvironmentAuditRecord"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePowerPlatformAdminEnvironmentAuditRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePowerPlatformAdminEnvironmentAuditRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPowerPlatformAdminEnvironmentAuditRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PowerPlatformAdminEnvironmentAuditRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *PowerPlatformAdminEnvironmentAuditRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type PowerPlatformAdminEnvironmentAuditRecordable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
