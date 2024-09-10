package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExchangeMailboxAuditBaseRecord struct {
    AuditData
}
// NewExchangeMailboxAuditBaseRecord instantiates a new ExchangeMailboxAuditBaseRecord and sets the default values.
func NewExchangeMailboxAuditBaseRecord()(*ExchangeMailboxAuditBaseRecord) {
    m := &ExchangeMailboxAuditBaseRecord{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.exchangeMailboxAuditBaseRecord"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateExchangeMailboxAuditBaseRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExchangeMailboxAuditBaseRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExchangeMailboxAuditBaseRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExchangeMailboxAuditBaseRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ExchangeMailboxAuditBaseRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type ExchangeMailboxAuditBaseRecordable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
