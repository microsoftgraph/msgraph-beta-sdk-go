package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MessageRecipient struct {
    Entity
    // 
    deliveryStatus *MessageStatus;
    // 
    events []MessageEvent;
    // 
    recipientEmail *string;
}
// Instantiates a new messageRecipient and sets the default values.
func NewMessageRecipient()(*MessageRecipient) {
    m := &MessageRecipient{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deliveryStatus property value. 
func (m *MessageRecipient) GetDeliveryStatus()(*MessageStatus) {
    if m == nil {
        return nil
    } else {
        return m.deliveryStatus
    }
}
// Gets the events property value. 
func (m *MessageRecipient) GetEvents()([]MessageEvent) {
    if m == nil {
        return nil
    } else {
        return m.events
    }
}
// Gets the recipientEmail property value. 
func (m *MessageRecipient) GetRecipientEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientEmail
    }
}
// The deserialization information for the current model
func (m *MessageRecipient) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deliveryStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMessageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MessageStatus)
            m.SetDeliveryStatus(&cast)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MessageRecipient) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeliveryStatus() != nil {
        cast := m.GetDeliveryStatus().String()
        err = writer.WriteStringValue("deliveryStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
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
// Sets the deliveryStatus property value. 
// Parameters:
//  - value : Value to set for the deliveryStatus property.
func (m *MessageRecipient) SetDeliveryStatus(value *MessageStatus)() {
    m.deliveryStatus = value
}
// Sets the events property value. 
// Parameters:
//  - value : Value to set for the events property.
func (m *MessageRecipient) SetEvents(value []MessageEvent)() {
    m.events = value
}
// Sets the recipientEmail property value. 
// Parameters:
//  - value : Value to set for the recipientEmail property.
func (m *MessageRecipient) SetRecipientEmail(value *string)() {
    m.recipientEmail = value
}
