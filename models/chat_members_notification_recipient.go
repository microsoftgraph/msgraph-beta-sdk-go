package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMembersNotificationRecipient 
type ChatMembersNotificationRecipient struct {
    TeamworkNotificationRecipient
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The chat's identifier.
    chatId *string
}
// NewChatMembersNotificationRecipient instantiates a new ChatMembersNotificationRecipient and sets the default values.
func NewChatMembersNotificationRecipient()(*ChatMembersNotificationRecipient) {
    m := &ChatMembersNotificationRecipient{
        TeamworkNotificationRecipient: *NewTeamworkNotificationRecipient(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChatMembersNotificationRecipientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMembersNotificationRecipientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMembersNotificationRecipient(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMembersNotificationRecipient) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetChatId gets the chatId property value. The chat's identifier.
func (m *ChatMembersNotificationRecipient) GetChatId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.chatId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMembersNotificationRecipient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeamworkNotificationRecipient.GetFieldDeserializers()
    res["chatId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ChatMembersNotificationRecipient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeamworkNotificationRecipient.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("chatId", m.GetChatId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMembersNotificationRecipient) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetChatId sets the chatId property value. The chat's identifier.
func (m *ChatMembersNotificationRecipient) SetChatId(value *string)() {
    if m != nil {
        m.chatId = value
    }
}
