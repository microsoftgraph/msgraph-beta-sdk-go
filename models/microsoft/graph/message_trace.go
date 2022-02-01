package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MessageTrace 
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
// NewMessageTrace instantiates a new messageTrace and sets the default values.
func NewMessageTrace()(*MessageTrace) {
    m := &MessageTrace{
        Entity: *NewEntity(),
    }
    return m
}
// GetDestinationIPAddress gets the destinationIPAddress property value. 
func (m *MessageTrace) GetDestinationIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationIPAddress
    }
}
// GetMessageId gets the messageId property value. 
func (m *MessageTrace) GetMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageId
    }
}
// GetReceivedDateTime gets the receivedDateTime property value. 
func (m *MessageTrace) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.receivedDateTime
    }
}
// GetRecipients gets the recipients property value. 
func (m *MessageTrace) GetRecipients()([]MessageRecipient) {
    if m == nil {
        return nil
    } else {
        return m.recipients
    }
}
// GetSenderEmail gets the senderEmail property value. 
func (m *MessageTrace) GetSenderEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.senderEmail
    }
}
// GetSize gets the size property value. 
func (m *MessageTrace) GetSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// GetSourceIPAddress gets the sourceIPAddress property value. 
func (m *MessageTrace) GetSourceIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceIPAddress
    }
}
// GetSubject gets the subject property value. 
func (m *MessageTrace) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
    if m.GetRecipients() != nil {
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
// SetDestinationIPAddress sets the destinationIPAddress property value. 
func (m *MessageTrace) SetDestinationIPAddress(value *string)() {
    if m != nil {
        m.destinationIPAddress = value
    }
}
// SetMessageId sets the messageId property value. 
func (m *MessageTrace) SetMessageId(value *string)() {
    if m != nil {
        m.messageId = value
    }
}
// SetReceivedDateTime sets the receivedDateTime property value. 
func (m *MessageTrace) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.receivedDateTime = value
    }
}
// SetRecipients sets the recipients property value. 
func (m *MessageTrace) SetRecipients(value []MessageRecipient)() {
    if m != nil {
        m.recipients = value
    }
}
// SetSenderEmail sets the senderEmail property value. 
func (m *MessageTrace) SetSenderEmail(value *string)() {
    if m != nil {
        m.senderEmail = value
    }
}
// SetSize sets the size property value. 
func (m *MessageTrace) SetSize(value *int32)() {
    if m != nil {
        m.size = value
    }
}
// SetSourceIPAddress sets the sourceIPAddress property value. 
func (m *MessageTrace) SetSourceIPAddress(value *string)() {
    if m != nil {
        m.sourceIPAddress = value
    }
}
// SetSubject sets the subject property value. 
func (m *MessageTrace) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
