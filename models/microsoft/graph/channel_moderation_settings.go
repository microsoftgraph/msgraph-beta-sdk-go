package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ChannelModerationSettings struct {
    additionalData map[string]interface{};
    allowNewMessageFromBots *bool;
    allowNewMessageFromConnectors *bool;
    replyRestriction *ReplyRestriction;
    userNewMessageRestriction *UserNewMessageRestriction;
}
func NewChannelModerationSettings()(*ChannelModerationSettings) {
    m := &ChannelModerationSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ChannelModerationSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ChannelModerationSettings) GetAllowNewMessageFromBots()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowNewMessageFromBots
    }
}
func (m *ChannelModerationSettings) GetAllowNewMessageFromConnectors()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowNewMessageFromConnectors
    }
}
func (m *ChannelModerationSettings) GetReplyRestriction()(*ReplyRestriction) {
    if m == nil {
        return nil
    } else {
        return m.replyRestriction
    }
}
func (m *ChannelModerationSettings) GetUserNewMessageRestriction()(*UserNewMessageRestriction) {
    if m == nil {
        return nil
    } else {
        return m.userNewMessageRestriction
    }
}
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
func (m *ChannelModerationSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ChannelModerationSettings) SetAllowNewMessageFromBots(value *bool)() {
    m.allowNewMessageFromBots = value
}
func (m *ChannelModerationSettings) SetAllowNewMessageFromConnectors(value *bool)() {
    m.allowNewMessageFromConnectors = value
}
func (m *ChannelModerationSettings) SetReplyRestriction(value *ReplyRestriction)() {
    m.replyRestriction = value
}
func (m *ChannelModerationSettings) SetUserNewMessageRestriction(value *UserNewMessageRestriction)() {
    m.userNewMessageRestriction = value
}
