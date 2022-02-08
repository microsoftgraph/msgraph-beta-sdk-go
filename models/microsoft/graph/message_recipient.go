package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MessageRecipient 
type MessageRecipient struct {
    Entity
    // 
    deliveryStatus *MessageStatus;
    // 
    events []MessageEvent;
    // 
    recipientEmail *string;
}
// NewMessageRecipient instantiates a new messageRecipient and sets the default values.
func NewMessageRecipient()(*MessageRecipient) {
    m := &MessageRecipient{
        Entity: *NewEntity(),
    }
    return m
}
// GetDeliveryStatus gets the deliveryStatus property value. 
func (m *MessageRecipient) GetDeliveryStatus()(*MessageStatus) {
    if m == nil {
        return nil
    } else {
        return m.deliveryStatus
    }
}
// GetEvents gets the events property value. 
func (m *MessageRecipient) GetEvents()([]MessageEvent) {
    if m == nil {
        return nil
    } else {
        return m.events
    }
}
// GetRecipientEmail gets the recipientEmail property value. 
func (m *MessageRecipient) GetRecipientEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientEmail
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageRecipient) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deliveryStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMessageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryStatus(val.(*MessageStatus))
        }
        return nil
    }
    res["events"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessageEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MessageEvent, len(val))
            for i, v := range val {
                res[i] = *(v.(*MessageEvent))
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["recipientEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *MessageRecipient) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MessageRecipient) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// SetDeliveryStatus sets the deliveryStatus property value. 
func (m *MessageRecipient) SetDeliveryStatus(value *MessageStatus)() {
    if m != nil {
        m.deliveryStatus = value
    }
}
// SetEvents sets the events property value. 
func (m *MessageRecipient) SetEvents(value []MessageEvent)() {
    if m != nil {
        m.events = value
    }
}
// SetRecipientEmail sets the recipientEmail property value. 
func (m *MessageRecipient) SetRecipientEmail(value *string)() {
    if m != nil {
        m.recipientEmail = value
    }
}
