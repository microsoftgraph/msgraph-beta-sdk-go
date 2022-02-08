package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChatMessageInfo 
type ChatMessageInfo struct {
    Entity
    // Body of the chatMessage. This will still contain markers for @mentions and attachments even though the object does not return @mentions and attachments.
    body *ItemBody;
    // Date time object representing the time at which message was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only.  If present, represents details of an event that happened in a chat, a channel, or a team, for example, members were added, and so on. For event messages, the messageType property will be set to systemEventMessage.
    eventDetail *EventMessageDetail;
    // Information about the sender of the message.
    from *ChatMessageFromIdentitySet;
    // If set to true, the original message has been deleted.
    isDeleted *bool;
    // The type of chat message. The possible values are: message, unknownFutureValue, systemEventMessage.
    messageType *ChatMessageType;
}
// NewChatMessageInfo instantiates a new chatMessageInfo and sets the default values.
func NewChatMessageInfo()(*ChatMessageInfo) {
    m := &ChatMessageInfo{
        Entity: *NewEntity(),
    }
    return m
}
// GetBody gets the body property value. Body of the chatMessage. This will still contain markers for @mentions and attachments even though the object does not return @mentions and attachments.
func (m *ChatMessageInfo) GetBody()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Date time object representing the time at which message was created.
func (m *ChatMessageInfo) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetEventDetail gets the eventDetail property value. Read-only.  If present, represents details of an event that happened in a chat, a channel, or a team, for example, members were added, and so on. For event messages, the messageType property will be set to systemEventMessage.
func (m *ChatMessageInfo) GetEventDetail()(*EventMessageDetail) {
    if m == nil {
        return nil
    } else {
        return m.eventDetail
    }
}
// GetFrom gets the from property value. Information about the sender of the message.
func (m *ChatMessageInfo) GetFrom()(*ChatMessageFromIdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
// GetIsDeleted gets the isDeleted property value. If set to true, the original message has been deleted.
func (m *ChatMessageInfo) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// GetMessageType gets the messageType property value. The type of chat message. The possible values are: message, unknownFutureValue, systemEventMessage.
func (m *ChatMessageInfo) GetMessageType()(*ChatMessageType) {
    if m == nil {
        return nil
    } else {
        return m.messageType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(*ItemBody))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["eventDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEventMessageDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDetail(val.(*EventMessageDetail))
        }
        return nil
    }
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageFromIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val.(*ChatMessageFromIdentitySet))
        }
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["messageType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageType(val.(*ChatMessageType))
        }
        return nil
    }
    return res
}
func (m *ChatMessageInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ChatMessageInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("eventDetail", m.GetEventDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("from", m.GetFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    if m.GetMessageType() != nil {
        cast := (*m.GetMessageType()).String()
        err = writer.WriteStringValue("messageType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBody sets the body property value. Body of the chatMessage. This will still contain markers for @mentions and attachments even though the object does not return @mentions and attachments.
func (m *ChatMessageInfo) SetBody(value *ItemBody)() {
    if m != nil {
        m.body = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Date time object representing the time at which message was created.
func (m *ChatMessageInfo) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetEventDetail sets the eventDetail property value. Read-only.  If present, represents details of an event that happened in a chat, a channel, or a team, for example, members were added, and so on. For event messages, the messageType property will be set to systemEventMessage.
func (m *ChatMessageInfo) SetEventDetail(value *EventMessageDetail)() {
    if m != nil {
        m.eventDetail = value
    }
}
// SetFrom sets the from property value. Information about the sender of the message.
func (m *ChatMessageInfo) SetFrom(value *ChatMessageFromIdentitySet)() {
    if m != nil {
        m.from = value
    }
}
// SetIsDeleted sets the isDeleted property value. If set to true, the original message has been deleted.
func (m *ChatMessageInfo) SetIsDeleted(value *bool)() {
    if m != nil {
        m.isDeleted = value
    }
}
// SetMessageType sets the messageType property value. The type of chat message. The possible values are: message, unknownFutureValue, systemEventMessage.
func (m *ChatMessageInfo) SetMessageType(value *ChatMessageType)() {
    if m != nil {
        m.messageType = value
    }
}
