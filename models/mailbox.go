package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Mailbox struct {
    DirectoryObject
}
// NewMailbox instantiates a new Mailbox and sets the default values.
func NewMailbox()(*Mailbox) {
    m := &Mailbox{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.mailbox"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMailboxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMailboxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailbox(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Mailbox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["folders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMailboxFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MailboxFolderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MailboxFolderable)
                }
            }
            m.SetFolders(res)
        }
        return nil
    }
    return res
}
// GetFolders gets the folders property value. The collection of folders in the mailbox.
// returns a []MailboxFolderable when successful
func (m *Mailbox) GetFolders()([]MailboxFolderable) {
    val, err := m.GetBackingStore().Get("folders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MailboxFolderable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Mailbox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetFolders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFolders()))
        for i, v := range m.GetFolders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("folders", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFolders sets the folders property value. The collection of folders in the mailbox.
func (m *Mailbox) SetFolders(value []MailboxFolderable)() {
    err := m.GetBackingStore().Set("folders", value)
    if err != nil {
        panic(err)
    }
}
type Mailboxable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFolders()([]MailboxFolderable)
    SetFolders(value []MailboxFolderable)()
}
