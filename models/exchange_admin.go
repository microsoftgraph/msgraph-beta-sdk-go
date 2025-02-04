package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExchangeAdmin struct {
    Entity
}
// NewExchangeAdmin instantiates a new ExchangeAdmin and sets the default values.
func NewExchangeAdmin()(*ExchangeAdmin) {
    m := &ExchangeAdmin{
        Entity: *NewEntity(),
    }
    return m
}
// CreateExchangeAdminFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExchangeAdminFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExchangeAdmin(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExchangeAdmin) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["mailboxes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMailboxFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Mailboxable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Mailboxable)
                }
            }
            m.SetMailboxes(res)
        }
        return nil
    }
    res["messageTraces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMessageTraceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MessageTraceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MessageTraceable)
                }
            }
            m.SetMessageTraces(res)
        }
        return nil
    }
    return res
}
// GetMailboxes gets the mailboxes property value. Represents a user's mailboxes.
// returns a []Mailboxable when successful
func (m *ExchangeAdmin) GetMailboxes()([]Mailboxable) {
    val, err := m.GetBackingStore().Get("mailboxes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Mailboxable)
    }
    return nil
}
// GetMessageTraces gets the messageTraces property value. The messageTraces property
// returns a []MessageTraceable when successful
func (m *ExchangeAdmin) GetMessageTraces()([]MessageTraceable) {
    val, err := m.GetBackingStore().Get("messageTraces")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MessageTraceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExchangeAdmin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMailboxes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMailboxes()))
        for i, v := range m.GetMailboxes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mailboxes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMessageTraces() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMessageTraces()))
        for i, v := range m.GetMessageTraces() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("messageTraces", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMailboxes sets the mailboxes property value. Represents a user's mailboxes.
func (m *ExchangeAdmin) SetMailboxes(value []Mailboxable)() {
    err := m.GetBackingStore().Set("mailboxes", value)
    if err != nil {
        panic(err)
    }
}
// SetMessageTraces sets the messageTraces property value. The messageTraces property
func (m *ExchangeAdmin) SetMessageTraces(value []MessageTraceable)() {
    err := m.GetBackingStore().Set("messageTraces", value)
    if err != nil {
        panic(err)
    }
}
type ExchangeAdminable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMailboxes()([]Mailboxable)
    GetMessageTraces()([]MessageTraceable)
    SetMailboxes(value []Mailboxable)()
    SetMessageTraces(value []MessageTraceable)()
}
