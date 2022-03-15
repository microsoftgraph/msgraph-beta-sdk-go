package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MessageRecipient provides operations to manage the collection of messageTrace entities.
type MessageRecipient struct {
    Entity
    // 
    deliveryStatus *MessageStatus;
    // 
    events []MessageEventable;
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
// CreateMessageRecipientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageRecipientFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMessageRecipient(), nil
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
func (m *MessageRecipient) GetEvents()([]MessageEventable) {
    if m == nil {
        return nil
    } else {
        return m.events
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
        val, err := n.GetCollectionOfObjectValues(CreateMessageEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MessageEventable, len(val))
            for i, v := range val {
                res[i] = v.(MessageEventable)
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
// GetRecipientEmail gets the recipientEmail property value. 
func (m *MessageRecipient) GetRecipientEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientEmail
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *MessageRecipient) SetEvents(value []MessageEventable)() {
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
