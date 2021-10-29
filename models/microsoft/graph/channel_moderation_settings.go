package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ChannelModerationSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether bots are allowed to post messages.
    allowNewMessageFromBots *bool;
    // Indicates whether connectors are allowed to post messages.
    allowNewMessageFromConnectors *bool;
    // Indicates who is allowed to reply to the teams channel. Possible values are: everyone, authorAndModerators, unknownFutureValue.
    replyRestriction *ReplyRestriction;
    // Indicates who is allowed to post messages to teams channel. Possible values are: everyone, everyoneExceptGuests, moderators, unknownFutureValue.
    userNewMessageRestriction *UserNewMessageRestriction;
}
// Instantiates a new channelModerationSettings and sets the default values.
func NewChannelModerationSettings()(*ChannelModerationSettings) {
    m := &ChannelModerationSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChannelModerationSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowNewMessageFromBots property value. Indicates whether bots are allowed to post messages.
func (m *ChannelModerationSettings) GetAllowNewMessageFromBots()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowNewMessageFromBots
    }
}
// Gets the allowNewMessageFromConnectors property value. Indicates whether connectors are allowed to post messages.
func (m *ChannelModerationSettings) GetAllowNewMessageFromConnectors()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowNewMessageFromConnectors
    }
}
// Gets the replyRestriction property value. Indicates who is allowed to reply to the teams channel. Possible values are: everyone, authorAndModerators, unknownFutureValue.
func (m *ChannelModerationSettings) GetReplyRestriction()(*ReplyRestriction) {
    if m == nil {
        return nil
    } else {
        return m.replyRestriction
    }
}
// Gets the userNewMessageRestriction property value. Indicates who is allowed to post messages to teams channel. Possible values are: everyone, everyoneExceptGuests, moderators, unknownFutureValue.
func (m *ChannelModerationSettings) GetUserNewMessageRestriction()(*UserNewMessageRestriction) {
    if m == nil {
        return nil
    } else {
        return m.userNewMessageRestriction
    }
}
// The deserialization information for the current model
func (m *ChannelModerationSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowNewMessageFromBots"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowNewMessageFromBots(val)
        return nil
    }
    res["allowNewMessageFromConnectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowNewMessageFromConnectors(val)
        return nil
    }
    res["replyRestriction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseReplyRestriction)
        if err != nil {
            return err
        }
        cast := val.(ReplyRestriction)
        m.SetReplyRestriction(&cast)
        return nil
    }
    res["userNewMessageRestriction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserNewMessageRestriction)
        if err != nil {
            return err
        }
        cast := val.(UserNewMessageRestriction)
        m.SetUserNewMessageRestriction(&cast)
        return nil
    }
    return res
}
func (m *ChannelModerationSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ChannelModerationSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowNewMessageFromBots", m.GetAllowNewMessageFromBots())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowNewMessageFromConnectors", m.GetAllowNewMessageFromConnectors())
        if err != nil {
            return err
        }
    }
    if m.GetReplyRestriction() != nil {
        cast := m.GetReplyRestriction().String()
        err := writer.WriteStringValue("replyRestriction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserNewMessageRestriction() != nil {
        cast := m.GetUserNewMessageRestriction().String()
        err := writer.WriteStringValue("userNewMessageRestriction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ChannelModerationSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowNewMessageFromBots property value. Indicates whether bots are allowed to post messages.
// Parameters:
//  - value : Value to set for the allowNewMessageFromBots property.
func (m *ChannelModerationSettings) SetAllowNewMessageFromBots(value *bool)() {
    m.allowNewMessageFromBots = value
}
// Sets the allowNewMessageFromConnectors property value. Indicates whether connectors are allowed to post messages.
// Parameters:
//  - value : Value to set for the allowNewMessageFromConnectors property.
func (m *ChannelModerationSettings) SetAllowNewMessageFromConnectors(value *bool)() {
    m.allowNewMessageFromConnectors = value
}
// Sets the replyRestriction property value. Indicates who is allowed to reply to the teams channel. Possible values are: everyone, authorAndModerators, unknownFutureValue.
// Parameters:
//  - value : Value to set for the replyRestriction property.
func (m *ChannelModerationSettings) SetReplyRestriction(value *ReplyRestriction)() {
    m.replyRestriction = value
}
// Sets the userNewMessageRestriction property value. Indicates who is allowed to post messages to teams channel. Possible values are: everyone, everyoneExceptGuests, moderators, unknownFutureValue.
// Parameters:
//  - value : Value to set for the userNewMessageRestriction property.
func (m *ChannelModerationSettings) SetUserNewMessageRestriction(value *UserNewMessageRestriction)() {
    m.userNewMessageRestriction = value
}
