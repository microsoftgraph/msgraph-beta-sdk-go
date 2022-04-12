package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelModerationSettings 
type ChannelModerationSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether bots are allowed to post messages.
    allowNewMessageFromBots *bool
    // Indicates whether connectors are allowed to post messages.
    allowNewMessageFromConnectors *bool
    // Indicates who is allowed to reply to the teams channel. Possible values are: everyone, authorAndModerators, unknownFutureValue.
    replyRestriction *ReplyRestriction
    // Indicates who is allowed to post messages to teams channel. Possible values are: everyone, everyoneExceptGuests, moderators, unknownFutureValue.
    userNewMessageRestriction *UserNewMessageRestriction
}
// NewChannelModerationSettings instantiates a new channelModerationSettings and sets the default values.
func NewChannelModerationSettings()(*ChannelModerationSettings) {
    m := &ChannelModerationSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChannelModerationSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelModerationSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannelModerationSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChannelModerationSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowNewMessageFromBots gets the allowNewMessageFromBots property value. Indicates whether bots are allowed to post messages.
func (m *ChannelModerationSettings) GetAllowNewMessageFromBots()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowNewMessageFromBots
    }
}
// GetAllowNewMessageFromConnectors gets the allowNewMessageFromConnectors property value. Indicates whether connectors are allowed to post messages.
func (m *ChannelModerationSettings) GetAllowNewMessageFromConnectors()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowNewMessageFromConnectors
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChannelModerationSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowNewMessageFromBots"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowNewMessageFromBots(val)
        }
        return nil
    }
    res["allowNewMessageFromConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowNewMessageFromConnectors(val)
        }
        return nil
    }
    res["replyRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseReplyRestriction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplyRestriction(val.(*ReplyRestriction))
        }
        return nil
    }
    res["userNewMessageRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserNewMessageRestriction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserNewMessageRestriction(val.(*UserNewMessageRestriction))
        }
        return nil
    }
    return res
}
// GetReplyRestriction gets the replyRestriction property value. Indicates who is allowed to reply to the teams channel. Possible values are: everyone, authorAndModerators, unknownFutureValue.
func (m *ChannelModerationSettings) GetReplyRestriction()(*ReplyRestriction) {
    if m == nil {
        return nil
    } else {
        return m.replyRestriction
    }
}
// GetUserNewMessageRestriction gets the userNewMessageRestriction property value. Indicates who is allowed to post messages to teams channel. Possible values are: everyone, everyoneExceptGuests, moderators, unknownFutureValue.
func (m *ChannelModerationSettings) GetUserNewMessageRestriction()(*UserNewMessageRestriction) {
    if m == nil {
        return nil
    } else {
        return m.userNewMessageRestriction
    }
}
// Serialize serializes information the current object
func (m *ChannelModerationSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetReplyRestriction()).String()
        err := writer.WriteStringValue("replyRestriction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserNewMessageRestriction() != nil {
        cast := (*m.GetUserNewMessageRestriction()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChannelModerationSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowNewMessageFromBots sets the allowNewMessageFromBots property value. Indicates whether bots are allowed to post messages.
func (m *ChannelModerationSettings) SetAllowNewMessageFromBots(value *bool)() {
    if m != nil {
        m.allowNewMessageFromBots = value
    }
}
// SetAllowNewMessageFromConnectors sets the allowNewMessageFromConnectors property value. Indicates whether connectors are allowed to post messages.
func (m *ChannelModerationSettings) SetAllowNewMessageFromConnectors(value *bool)() {
    if m != nil {
        m.allowNewMessageFromConnectors = value
    }
}
// SetReplyRestriction sets the replyRestriction property value. Indicates who is allowed to reply to the teams channel. Possible values are: everyone, authorAndModerators, unknownFutureValue.
func (m *ChannelModerationSettings) SetReplyRestriction(value *ReplyRestriction)() {
    if m != nil {
        m.replyRestriction = value
    }
}
// SetUserNewMessageRestriction sets the userNewMessageRestriction property value. Indicates who is allowed to post messages to teams channel. Possible values are: everyone, everyoneExceptGuests, moderators, unknownFutureValue.
func (m *ChannelModerationSettings) SetUserNewMessageRestriction(value *UserNewMessageRestriction)() {
    if m != nil {
        m.userNewMessageRestriction = value
    }
}
