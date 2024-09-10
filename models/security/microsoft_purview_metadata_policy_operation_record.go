package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MicrosoftPurviewMetadataPolicyOperationRecord struct {
    AuditData
}
// NewMicrosoftPurviewMetadataPolicyOperationRecord instantiates a new MicrosoftPurviewMetadataPolicyOperationRecord and sets the default values.
func NewMicrosoftPurviewMetadataPolicyOperationRecord()(*MicrosoftPurviewMetadataPolicyOperationRecord) {
    m := &MicrosoftPurviewMetadataPolicyOperationRecord{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.microsoftPurviewMetadataPolicyOperationRecord"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMicrosoftPurviewMetadataPolicyOperationRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMicrosoftPurviewMetadataPolicyOperationRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftPurviewMetadataPolicyOperationRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MicrosoftPurviewMetadataPolicyOperationRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MicrosoftPurviewMetadataPolicyOperationRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type MicrosoftPurviewMetadataPolicyOperationRecordable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
