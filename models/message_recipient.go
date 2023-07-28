package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageRecipient 
type MessageRecipient struct {
    Entity
}
// NewMessageRecipient instantiates a new messageRecipient and sets the default values.
func NewMessageRecipient()(*MessageRecipient) {
    m := &MessageRecipient{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMessageRecipientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageRecipientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageRecipient(), nil
}
// GetDeliveryStatus gets the deliveryStatus property value. The deliveryStatus property
func (m *MessageRecipient) GetDeliveryStatus()(*MessageStatus) {
    val, err := m.GetBackingStore().Get("deliveryStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MessageStatus)
    }
    return nil
}
// GetEvents gets the events property value. The events property
func (m *MessageRecipient) GetEvents()([]MessageEventable) {
    val, err := m.GetBackingStore().Get("events")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MessageEventable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageRecipient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deliveryStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMessageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryStatus(val.(*MessageStatus))
        }
        return nil
    }
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMessageEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MessageEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MessageEventable)
                }
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["recipientEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientEmail(val)
        }
        return nil
    }
    return res
}
// GetRecipientEmail gets the recipientEmail property value. The recipientEmail property
func (m *MessageRecipient) GetRecipientEmail()(*string) {
    val, err := m.GetBackingStore().Get("recipientEmail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MessageRecipient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeliveryStatus() != nil {
        cast := (*m.GetDeliveryStatus()).String()
        err = writer.WriteStringValue("deliveryStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("events", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recipientEmail", m.GetRecipientEmail())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeliveryStatus sets the deliveryStatus property value. The deliveryStatus property
func (m *MessageRecipient) SetDeliveryStatus(value *MessageStatus)() {
    err := m.GetBackingStore().Set("deliveryStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetEvents sets the events property value. The events property
func (m *MessageRecipient) SetEvents(value []MessageEventable)() {
    err := m.GetBackingStore().Set("events", value)
    if err != nil {
        panic(err)
    }
}
// SetRecipientEmail sets the recipientEmail property value. The recipientEmail property
func (m *MessageRecipient) SetRecipientEmail(value *string)() {
    err := m.GetBackingStore().Set("recipientEmail", value)
    if err != nil {
        panic(err)
    }
}
// MessageRecipientable 
type MessageRecipientable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeliveryStatus()(*MessageStatus)
    GetEvents()([]MessageEventable)
    GetRecipientEmail()(*string)
    SetDeliveryStatus(value *MessageStatus)()
    SetEvents(value []MessageEventable)()
    SetRecipientEmail(value *string)()
}
