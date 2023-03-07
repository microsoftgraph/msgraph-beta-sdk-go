package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageTrace 
type MessageTrace struct {
    Entity
}
// NewMessageTrace instantiates a new messageTrace and sets the default values.
func NewMessageTrace()(*MessageTrace) {
    m := &MessageTrace{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMessageTraceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageTraceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageTrace(), nil
}
// GetDestinationIPAddress gets the destinationIPAddress property value. The destinationIPAddress property
func (m *MessageTrace) GetDestinationIPAddress()(*string) {
    val, err := m.GetBackingStore().Get("destinationIPAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageTrace) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["destinationIPAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationIPAddress(val)
        }
        return nil
    }
    res["messageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageId(val)
        }
        return nil
    }
    res["receivedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedDateTime(val)
        }
        return nil
    }
    res["recipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMessageRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MessageRecipientable, len(val))
            for i, v := range val {
                res[i] = v.(MessageRecipientable)
            }
            m.SetRecipients(res)
        }
        return nil
    }
    res["senderEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderEmail(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["sourceIPAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceIPAddress(val)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetMessageId gets the messageId property value. The messageId property
func (m *MessageTrace) GetMessageId()(*string) {
    val, err := m.GetBackingStore().Get("messageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReceivedDateTime gets the receivedDateTime property value. The receivedDateTime property
func (m *MessageTrace) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("receivedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRecipients gets the recipients property value. The recipients property
func (m *MessageTrace) GetRecipients()([]MessageRecipientable) {
    val, err := m.GetBackingStore().Get("recipients")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MessageRecipientable)
    }
    return nil
}
// GetSenderEmail gets the senderEmail property value. The senderEmail property
func (m *MessageTrace) GetSenderEmail()(*string) {
    val, err := m.GetBackingStore().Get("senderEmail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSize gets the size property value. The size property
func (m *MessageTrace) GetSize()(*int32) {
    val, err := m.GetBackingStore().Get("size")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSourceIPAddress gets the sourceIPAddress property value. The sourceIPAddress property
func (m *MessageTrace) GetSourceIPAddress()(*string) {
    val, err := m.GetBackingStore().Get("sourceIPAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubject gets the subject property value. The subject property
func (m *MessageTrace) GetSubject()(*string) {
    val, err := m.GetBackingStore().Get("subject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MessageTrace) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecipients()))
        for i, v := range m.GetRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
// SetDestinationIPAddress sets the destinationIPAddress property value. The destinationIPAddress property
func (m *MessageTrace) SetDestinationIPAddress(value *string)() {
    err := m.GetBackingStore().Set("destinationIPAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetMessageId sets the messageId property value. The messageId property
func (m *MessageTrace) SetMessageId(value *string)() {
    err := m.GetBackingStore().Set("messageId", value)
    if err != nil {
        panic(err)
    }
}
// SetReceivedDateTime sets the receivedDateTime property value. The receivedDateTime property
func (m *MessageTrace) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("receivedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRecipients sets the recipients property value. The recipients property
func (m *MessageTrace) SetRecipients(value []MessageRecipientable)() {
    err := m.GetBackingStore().Set("recipients", value)
    if err != nil {
        panic(err)
    }
}
// SetSenderEmail sets the senderEmail property value. The senderEmail property
func (m *MessageTrace) SetSenderEmail(value *string)() {
    err := m.GetBackingStore().Set("senderEmail", value)
    if err != nil {
        panic(err)
    }
}
// SetSize sets the size property value. The size property
func (m *MessageTrace) SetSize(value *int32)() {
    err := m.GetBackingStore().Set("size", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceIPAddress sets the sourceIPAddress property value. The sourceIPAddress property
func (m *MessageTrace) SetSourceIPAddress(value *string)() {
    err := m.GetBackingStore().Set("sourceIPAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetSubject sets the subject property value. The subject property
func (m *MessageTrace) SetSubject(value *string)() {
    err := m.GetBackingStore().Set("subject", value)
    if err != nil {
        panic(err)
    }
}
// MessageTraceable 
type MessageTraceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDestinationIPAddress()(*string)
    GetMessageId()(*string)
    GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecipients()([]MessageRecipientable)
    GetSenderEmail()(*string)
    GetSize()(*int32)
    GetSourceIPAddress()(*string)
    GetSubject()(*string)
    SetDestinationIPAddress(value *string)()
    SetMessageId(value *string)()
    SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecipients(value []MessageRecipientable)()
    SetSenderEmail(value *string)()
    SetSize(value *int32)()
    SetSourceIPAddress(value *string)()
    SetSubject(value *string)()
}
