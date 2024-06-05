package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExchangeRestoreSession struct {
    RestoreSessionBase
}
// NewExchangeRestoreSession instantiates a new ExchangeRestoreSession and sets the default values.
func NewExchangeRestoreSession()(*ExchangeRestoreSession) {
    m := &ExchangeRestoreSession{
        RestoreSessionBase: *NewRestoreSessionBase(),
    }
    odataTypeValue := "#microsoft.graph.exchangeRestoreSession"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateExchangeRestoreSessionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExchangeRestoreSessionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExchangeRestoreSession(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExchangeRestoreSession) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RestoreSessionBase.GetFieldDeserializers()
    res["mailboxRestoreArtifacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMailboxRestoreArtifactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MailboxRestoreArtifactable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MailboxRestoreArtifactable)
                }
            }
            m.SetMailboxRestoreArtifacts(res)
        }
        return nil
    }
    return res
}
// GetMailboxRestoreArtifacts gets the mailboxRestoreArtifacts property value. A collection of restore points and destination details that can be used to restore Exchange mailboxes.
// returns a []MailboxRestoreArtifactable when successful
func (m *ExchangeRestoreSession) GetMailboxRestoreArtifacts()([]MailboxRestoreArtifactable) {
    val, err := m.GetBackingStore().Get("mailboxRestoreArtifacts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MailboxRestoreArtifactable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExchangeRestoreSession) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RestoreSessionBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMailboxRestoreArtifacts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMailboxRestoreArtifacts()))
        for i, v := range m.GetMailboxRestoreArtifacts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mailboxRestoreArtifacts", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMailboxRestoreArtifacts sets the mailboxRestoreArtifacts property value. A collection of restore points and destination details that can be used to restore Exchange mailboxes.
func (m *ExchangeRestoreSession) SetMailboxRestoreArtifacts(value []MailboxRestoreArtifactable)() {
    err := m.GetBackingStore().Set("mailboxRestoreArtifacts", value)
    if err != nil {
        panic(err)
    }
}
type ExchangeRestoreSessionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RestoreSessionBaseable
    GetMailboxRestoreArtifacts()([]MailboxRestoreArtifactable)
    SetMailboxRestoreArtifacts(value []MailboxRestoreArtifactable)()
}
