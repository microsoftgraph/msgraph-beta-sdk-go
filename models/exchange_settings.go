package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExchangeSettings struct {
    Entity
}
// NewExchangeSettings instantiates a new ExchangeSettings and sets the default values.
func NewExchangeSettings()(*ExchangeSettings) {
    m := &ExchangeSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateExchangeSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExchangeSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExchangeSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExchangeSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["inPlaceArchiveMailboxId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInPlaceArchiveMailboxId(val)
        }
        return nil
    }
    res["primaryMailboxId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryMailboxId(val)
        }
        return nil
    }
    return res
}
// GetInPlaceArchiveMailboxId gets the inPlaceArchiveMailboxId property value. The unique identifier for the user's in-place archive mailbox.
// returns a *string when successful
func (m *ExchangeSettings) GetInPlaceArchiveMailboxId()(*string) {
    val, err := m.GetBackingStore().Get("inPlaceArchiveMailboxId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrimaryMailboxId gets the primaryMailboxId property value. The unique identifier for the user's primary mailbox.
// returns a *string when successful
func (m *ExchangeSettings) GetPrimaryMailboxId()(*string) {
    val, err := m.GetBackingStore().Get("primaryMailboxId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExchangeSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("inPlaceArchiveMailboxId", m.GetInPlaceArchiveMailboxId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryMailboxId", m.GetPrimaryMailboxId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInPlaceArchiveMailboxId sets the inPlaceArchiveMailboxId property value. The unique identifier for the user's in-place archive mailbox.
func (m *ExchangeSettings) SetInPlaceArchiveMailboxId(value *string)() {
    err := m.GetBackingStore().Set("inPlaceArchiveMailboxId", value)
    if err != nil {
        panic(err)
    }
}
// SetPrimaryMailboxId sets the primaryMailboxId property value. The unique identifier for the user's primary mailbox.
func (m *ExchangeSettings) SetPrimaryMailboxId(value *string)() {
    err := m.GetBackingStore().Set("primaryMailboxId", value)
    if err != nil {
        panic(err)
    }
}
type ExchangeSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInPlaceArchiveMailboxId()(*string)
    GetPrimaryMailboxId()(*string)
    SetInPlaceArchiveMailboxId(value *string)()
    SetPrimaryMailboxId(value *string)()
}
