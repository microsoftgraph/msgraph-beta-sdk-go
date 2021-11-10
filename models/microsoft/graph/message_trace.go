package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MessageTrace struct {
    Entity
    // 
    destinationIPAddress *string;
    // 
    messageId *string;
    // 
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    recipients []MessageRecipient;
    // 
    senderEmail *string;
    // 
    size *int32;
    // 
    sourceIPAddress *string;
    // 
    subject *string;
}
// Instantiates a new messageTrace and sets the default values.
func NewMessageTrace()(*MessageTrace) {
    m := &MessageTrace{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the destinationIPAddress property value. 
func (m *MessageTrace) GetDestinationIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationIPAddress
    }
}
// Gets the messageId property value. 
func (m *MessageTrace) GetMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageId
    }
}
// Gets the receivedDateTime property value. 
func (m *MessageTrace) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.receivedDateTime
    }
}
// Gets the recipients property value. 
func (m *MessageTrace) GetRecipients()([]MessageRecipient) {
    if m == nil {
        return nil
    } else {
        return m.recipients
    }
}
// Gets the senderEmail property value. 
func (m *MessageTrace) GetSenderEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.senderEmail
    }
}
// Gets the size property value. 
func (m *MessageTrace) GetSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// Gets the sourceIPAddress property value. 
func (m *MessageTrace) GetSourceIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceIPAddress
    }
}
// Gets the subject property value. 
func (m *MessageTrace) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// The deserialization information for the current model
func (m *MessageTrace) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["destinationIPAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationIPAddress(val)
        }
        return nil
    }
    res["messageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageId(val)
        }
        return nil
    }
    res["receivedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedDateTime(val)
        }
        return nil
    }
    res["recipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessageRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MessageRecipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*MessageRecipient))
            }
            m.SetRecipients(res)
        }
        return nil
    }
    res["senderEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderEmail(val)
        }
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["sourceIPAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceIPAddress(val)
        }
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
func (m *MessageTrace) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MessageTrace) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("destinationIPAddress", m.GetDestinationIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("messageId", m.GetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTime", m.GetReceivedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecipients()))
        for i, v := range m.GetRecipients() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("recipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderEmail", m.GetSenderEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceIPAddress", m.GetSourceIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the destinationIPAddress property value. 
// Parameters:
//  - value : Value to set for the destinationIPAddress property.
func (m *MessageTrace) SetDestinationIPAddress(value *string)() {
    m.destinationIPAddress = value
}
// Sets the messageId property value. 
// Parameters:
//  - value : Value to set for the messageId property.
func (m *MessageTrace) SetMessageId(value *string)() {
    m.messageId = value
}
// Sets the receivedDateTime property value. 
// Parameters:
//  - value : Value to set for the receivedDateTime property.
func (m *MessageTrace) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTime = value
}
// Sets the recipients property value. 
// Parameters:
//  - value : Value to set for the recipients property.
func (m *MessageTrace) SetRecipients(value []MessageRecipient)() {
    m.recipients = value
}
// Sets the senderEmail property value. 
// Parameters:
//  - value : Value to set for the senderEmail property.
func (m *MessageTrace) SetSenderEmail(value *string)() {
    m.senderEmail = value
}
// Sets the size property value. 
// Parameters:
//  - value : Value to set for the size property.
func (m *MessageTrace) SetSize(value *int32)() {
    m.size = value
}
// Sets the sourceIPAddress property value. 
// Parameters:
//  - value : Value to set for the sourceIPAddress property.
func (m *MessageTrace) SetSourceIPAddress(value *string)() {
    m.sourceIPAddress = value
}
// Sets the subject property value. 
// Parameters:
//  - value : Value to set for the subject property.
func (m *MessageTrace) SetSubject(value *string)() {
    m.subject = value
}
