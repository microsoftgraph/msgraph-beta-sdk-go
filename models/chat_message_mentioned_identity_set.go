package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ChatMessageMentionedIdentitySet struct {
    IdentitySet
}
// NewChatMessageMentionedIdentitySet instantiates a new ChatMessageMentionedIdentitySet and sets the default values.
func NewChatMessageMentionedIdentitySet()(*ChatMessageMentionedIdentitySet) {
    m := &ChatMessageMentionedIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    odataTypeValue := "#microsoft.graph.chatMessageMentionedIdentitySet"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateChatMessageMentionedIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateChatMessageMentionedIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessageMentionedIdentitySet(), nil
}
// GetConversation gets the conversation property value. If present, represents a conversation (for example, team or channel) @mentioned in a message.
// returns a TeamworkConversationIdentityable when successful
func (m *ChatMessageMentionedIdentitySet) GetConversation()(TeamworkConversationIdentityable) {
    val, err := m.GetBackingStore().Get("conversation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkConversationIdentityable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ChatMessageMentionedIdentitySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["conversation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkConversationIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversation(val.(TeamworkConversationIdentityable))
        }
        return nil
    }
    res["tag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkTagIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTag(val.(TeamworkTagIdentityable))
        }
        return nil
    }
    return res
}
// GetTag gets the tag property value. If present, represents a tag @mentioned in a team message.
// returns a TeamworkTagIdentityable when successful
func (m *ChatMessageMentionedIdentitySet) GetTag()(TeamworkTagIdentityable) {
    val, err := m.GetBackingStore().Get("tag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkTagIdentityable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ChatMessageMentionedIdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("conversation", m.GetConversation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tag", m.GetTag())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConversation sets the conversation property value. If present, represents a conversation (for example, team or channel) @mentioned in a message.
func (m *ChatMessageMentionedIdentitySet) SetConversation(value TeamworkConversationIdentityable)() {
    err := m.GetBackingStore().Set("conversation", value)
    if err != nil {
        panic(err)
    }
}
// SetTag sets the tag property value. If present, represents a tag @mentioned in a team message.
func (m *ChatMessageMentionedIdentitySet) SetTag(value TeamworkTagIdentityable)() {
    err := m.GetBackingStore().Set("tag", value)
    if err != nil {
        panic(err)
    }
}
type ChatMessageMentionedIdentitySetable interface {
    IdentitySetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConversation()(TeamworkConversationIdentityable)
    GetTag()(TeamworkTagIdentityable)
    SetConversation(value TeamworkConversationIdentityable)()
    SetTag(value TeamworkTagIdentityable)()
}
