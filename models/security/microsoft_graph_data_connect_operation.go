package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MicrosoftGraphDataConnectOperation struct {
    AuditData
}
// NewMicrosoftGraphDataConnectOperation instantiates a new MicrosoftGraphDataConnectOperation and sets the default values.
func NewMicrosoftGraphDataConnectOperation()(*MicrosoftGraphDataConnectOperation) {
    m := &MicrosoftGraphDataConnectOperation{
        AuditData: *NewAuditData(),
    }
    odataTypeValue := "#microsoft.graph.security.microsoftGraphDataConnectOperation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMicrosoftGraphDataConnectOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMicrosoftGraphDataConnectOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphDataConnectOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MicrosoftGraphDataConnectOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuditData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MicrosoftGraphDataConnectOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuditData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type MicrosoftGraphDataConnectOperationable interface {
    AuditDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
