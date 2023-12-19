package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ChannelModerationSettings 
type ChannelModerationSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewChannelModerationSettings instantiates a new channelModerationSettings and sets the default values.
func NewChannelModerationSettings()(*ChannelModerationSettings) {
    m := &ChannelModerationSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChannelModerationSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelModerationSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannelModerationSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChannelModerationSettings) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAllowNewMessageFromBots gets the allowNewMessageFromBots property value. Indicates whether bots are allowed to post messages.
func (m *ChannelModerationSettings) GetAllowNewMessageFromBots()(*bool) {
    val, err := m.GetBackingStore().Get("allowNewMessageFromBots")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowNewMessageFromConnectors gets the allowNewMessageFromConnectors property value. Indicates whether connectors are allowed to post messages.
func (m *ChannelModerationSettings) GetAllowNewMessageFromConnectors()(*bool) {
    val, err := m.GetBackingStore().Get("allowNewMessageFromConnectors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *ChannelModerationSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["replyRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChannelModerationSettings_replyRestriction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplyRestriction(val.(*ChannelModerationSettings_replyRestriction))
        }
        return nil
    }
    res["userNewMessageRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChannelModerationSettings_userNewMessageRestriction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserNewMessageRestriction(val.(*ChannelModerationSettings_userNewMessageRestriction))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ChannelModerationSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReplyRestriction gets the replyRestriction property value. Indicates who is allowed to reply to the teams channel. Possible values are: everyone, authorAndModerators, unknownFutureValue.
func (m *ChannelModerationSettings) GetReplyRestriction()(*ChannelModerationSettings_replyRestriction) {
    val, err := m.GetBackingStore().Get("replyRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ChannelModerationSettings_replyRestriction)
    }
    return nil
}
// GetUserNewMessageRestriction gets the userNewMessageRestriction property value. Indicates who is allowed to post messages to teams channel. Possible values are: everyone, everyoneExceptGuests, moderators, unknownFutureValue.
func (m *ChannelModerationSettings) GetUserNewMessageRestriction()(*ChannelModerationSettings_userNewMessageRestriction) {
    val, err := m.GetBackingStore().Get("userNewMessageRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ChannelModerationSettings_userNewMessageRestriction)
    }
    return nil
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChannelModerationSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowNewMessageFromBots sets the allowNewMessageFromBots property value. Indicates whether bots are allowed to post messages.
func (m *ChannelModerationSettings) SetAllowNewMessageFromBots(value *bool)() {
    err := m.GetBackingStore().Set("allowNewMessageFromBots", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowNewMessageFromConnectors sets the allowNewMessageFromConnectors property value. Indicates whether connectors are allowed to post messages.
func (m *ChannelModerationSettings) SetAllowNewMessageFromConnectors(value *bool)() {
    err := m.GetBackingStore().Set("allowNewMessageFromConnectors", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ChannelModerationSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChannelModerationSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetReplyRestriction sets the replyRestriction property value. Indicates who is allowed to reply to the teams channel. Possible values are: everyone, authorAndModerators, unknownFutureValue.
func (m *ChannelModerationSettings) SetReplyRestriction(value *ChannelModerationSettings_replyRestriction)() {
    err := m.GetBackingStore().Set("replyRestriction", value)
    if err != nil {
        panic(err)
    }
}
// SetUserNewMessageRestriction sets the userNewMessageRestriction property value. Indicates who is allowed to post messages to teams channel. Possible values are: everyone, everyoneExceptGuests, moderators, unknownFutureValue.
func (m *ChannelModerationSettings) SetUserNewMessageRestriction(value *ChannelModerationSettings_userNewMessageRestriction)() {
    err := m.GetBackingStore().Set("userNewMessageRestriction", value)
    if err != nil {
        panic(err)
    }
}
// ChannelModerationSettingsable 
type ChannelModerationSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowNewMessageFromBots()(*bool)
    GetAllowNewMessageFromConnectors()(*bool)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetReplyRestriction()(*ChannelModerationSettings_replyRestriction)
    GetUserNewMessageRestriction()(*ChannelModerationSettings_userNewMessageRestriction)
    SetAllowNewMessageFromBots(value *bool)()
    SetAllowNewMessageFromConnectors(value *bool)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetReplyRestriction(value *ChannelModerationSettings_replyRestriction)()
    SetUserNewMessageRestriction(value *ChannelModerationSettings_userNewMessageRestriction)()
}
