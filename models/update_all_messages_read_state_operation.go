package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UpdateAllMessagesReadStateOperation struct {
    MailFolderOperation
}
// NewUpdateAllMessagesReadStateOperation instantiates a new UpdateAllMessagesReadStateOperation and sets the default values.
func NewUpdateAllMessagesReadStateOperation()(*UpdateAllMessagesReadStateOperation) {
    m := &UpdateAllMessagesReadStateOperation{
        MailFolderOperation: *NewMailFolderOperation(),
    }
    return m
}
// CreateUpdateAllMessagesReadStateOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUpdateAllMessagesReadStateOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateAllMessagesReadStateOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UpdateAllMessagesReadStateOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MailFolderOperation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *UpdateAllMessagesReadStateOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MailFolderOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type UpdateAllMessagesReadStateOperationable interface {
    MailFolderOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
