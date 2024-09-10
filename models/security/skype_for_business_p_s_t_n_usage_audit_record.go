package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SkypeForBusinessPSTNUsageAuditRecord struct {
    AuditData
}
// NewSkypeForBusinessPSTNUsageAuditRecord instantiates a new SkypeForBusinessPSTNUsageAuditRecord and sets the default values.
func NewSkypeForBusinessPSTNUsageAuditRecord()(*SkypeForBusinessPSTNUsageAuditRecord) {
    m := &SkypeForBusinessPSTNUsageAuditRecord{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.skypeForBusinessPSTNUsageAuditRecord"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSkypeForBusinessPSTNUsageAuditRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSkypeForBusinessPSTNUsageAuditRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSkypeForBusinessPSTNUsageAuditRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SkypeForBusinessPSTNUsageAuditRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SkypeForBusinessPSTNUsageAuditRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type SkypeForBusinessPSTNUsageAuditRecordable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
