package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChatMessageMentionedIdentitySet 
type ChatMessageMentionedIdentitySet struct {
    IdentitySet
    // If present, represents a conversation (for example, team or channel) @mentioned in a message.
    conversation *TeamworkConversationIdentity;
    // If present, represents a tag @mentioned in a team message.
    tag *TeamworkTagIdentity;
}
// NewChatMessageMentionedIdentitySet instantiates a new chatMessageMentionedIdentitySet and sets the default values.
func NewChatMessageMentionedIdentitySet()(*ChatMessageMentionedIdentitySet) {
    m := &ChatMessageMentionedIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    return m
}
// GetConversation gets the conversation property value. If present, represents a conversation (for example, team or channel) @mentioned in a message.
func (m *ChatMessageMentionedIdentitySet) GetConversation()(*TeamworkConversationIdentity) {
    if m == nil {
        return nil
    } else {
        return m.conversation
    }
}
// GetTag gets the tag property value. If present, represents a tag @mentioned in a team message.
func (m *ChatMessageMentionedIdentitySet) GetTag()(*TeamworkTagIdentity) {
    if m == nil {
        return nil
    } else {
        return m.tag
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageMentionedIdentitySet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["conversation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkConversationIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversation(val.(*TeamworkConversationIdentity))
        }
        return nil
    }
    res["tag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkTagIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTag(val.(*TeamworkTagIdentity))
        }
        return nil
    }
    return res
}
func (m *ChatMessageMentionedIdentitySet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ChatMessageMentionedIdentitySet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *ChatMessageMentionedIdentitySet) SetConversation(value *TeamworkConversationIdentity)() {
    if m != nil {
        m.conversation = value
    }
}
// SetTag sets the tag property value. If present, represents a tag @mentioned in a team message.
func (m *ChatMessageMentionedIdentitySet) SetTag(value *TeamworkTagIdentity)() {
    if m != nil {
        m.tag = value
    }
}
