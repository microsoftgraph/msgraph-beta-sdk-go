package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
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
    // The OdataType property
    odataType *string
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
    return m.additionalData
}
// GetAllowNewMessageFromBots gets the allowNewMessageFromBots property value. Indicates whether bots are allowed to post messages.
func (m *ChannelModerationSettings) GetAllowNewMessageFromBots()(*bool) {
    return m.allowNewMessageFromBots
}
// GetAllowNewMessageFromConnectors gets the allowNewMessageFromConnectors property value. Indicates whether connectors are allowed to post messages.
func (m *ChannelModerationSettings) GetAllowNewMessageFromConnectors()(*bool) {
    return m.allowNewMessageFromConnectors
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChannelModerationSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowNewMessageFromBots"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAllowNewMessageFromBots)
    res["allowNewMessageFromConnectors"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAllowNewMessageFromConnectors)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["replyRestriction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseReplyRestriction , m.SetReplyRestriction)
    res["userNewMessageRestriction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseUserNewMessageRestriction , m.SetUserNewMessageRestriction)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ChannelModerationSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetReplyRestriction gets the replyRestriction property value. Indicates who is allowed to reply to the teams channel. Possible values are: everyone, authorAndModerators, unknownFutureValue.
func (m *ChannelModerationSettings) GetReplyRestriction()(*ReplyRestriction) {
    return m.replyRestriction
}
// GetUserNewMessageRestriction gets the userNewMessageRestriction property value. Indicates who is allowed to post messages to teams channel. Possible values are: everyone, everyoneExceptGuests, moderators, unknownFutureValue.
func (m *ChannelModerationSettings) GetUserNewMessageRestriction()(*UserNewMessageRestriction) {
    return m.userNewMessageRestriction
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
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    m.additionalData = value
}
// SetAllowNewMessageFromBots sets the allowNewMessageFromBots property value. Indicates whether bots are allowed to post messages.
func (m *ChannelModerationSettings) SetAllowNewMessageFromBots(value *bool)() {
    m.allowNewMessageFromBots = value
}
// SetAllowNewMessageFromConnectors sets the allowNewMessageFromConnectors property value. Indicates whether connectors are allowed to post messages.
func (m *ChannelModerationSettings) SetAllowNewMessageFromConnectors(value *bool)() {
    m.allowNewMessageFromConnectors = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChannelModerationSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReplyRestriction sets the replyRestriction property value. Indicates who is allowed to reply to the teams channel. Possible values are: everyone, authorAndModerators, unknownFutureValue.
func (m *ChannelModerationSettings) SetReplyRestriction(value *ReplyRestriction)() {
    m.replyRestriction = value
}
// SetUserNewMessageRestriction sets the userNewMessageRestriction property value. Indicates who is allowed to post messages to teams channel. Possible values are: everyone, everyoneExceptGuests, moderators, unknownFutureValue.
func (m *ChannelModerationSettings) SetUserNewMessageRestriction(value *UserNewMessageRestriction)() {
    m.userNewMessageRestriction = value
}
