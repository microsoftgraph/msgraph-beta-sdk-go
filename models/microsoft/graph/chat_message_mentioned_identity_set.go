package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ChatMessageMentionedIdentitySet struct {
    IdentitySet
    // If present, represents a conversation (for example, team or channel) @mentioned in a message.
    conversation *TeamworkConversationIdentity;
    // If present, represents a tag @mentioned in a team message.
    tag *TeamworkTagIdentity;
}
// Instantiates a new chatMessageMentionedIdentitySet and sets the default values.
func NewChatMessageMentionedIdentitySet()(*ChatMessageMentionedIdentitySet) {
    m := &ChatMessageMentionedIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    return m
}
// Gets the conversation property value. If present, represents a conversation (for example, team or channel) @mentioned in a message.
func (m *ChatMessageMentionedIdentitySet) GetConversation()(*TeamworkConversationIdentity) {
    if m == nil {
        return nil
    } else {
        return m.conversation
    }
}
// Gets the tag property value. If present, represents a tag @mentioned in a team message.
func (m *ChatMessageMentionedIdentitySet) GetTag()(*TeamworkTagIdentity) {
    if m == nil {
        return nil
    } else {
        return m.tag
    }
}
// The deserialization information for the current model
func (m *ChatMessageMentionedIdentitySet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["conversation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkConversationIdentity() })
        if err != nil {
            return err
        }
        m.SetConversation(val.(*TeamworkConversationIdentity))
        return nil
    }
    res["tag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkTagIdentity() })
        if err != nil {
            return err
        }
        m.SetTag(val.(*TeamworkTagIdentity))
        return nil
    }
    return res
}
func (m *ChatMessageMentionedIdentitySet) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the conversation property value. If present, represents a conversation (for example, team or channel) @mentioned in a message.
// Parameters:
//  - value : Value to set for the conversation property.
func (m *ChatMessageMentionedIdentitySet) SetConversation(value *TeamworkConversationIdentity)() {
    m.conversation = value
}
// Sets the tag property value. If present, represents a tag @mentioned in a team message.
// Parameters:
//  - value : Value to set for the tag property.
func (m *ChatMessageMentionedIdentitySet) SetTag(value *TeamworkTagIdentity)() {
    m.tag = value
}
