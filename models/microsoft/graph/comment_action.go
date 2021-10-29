package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CommentAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If true, this activity was a reply to an existing comment thread.
    isReply *bool;
    // The identity of the user who started the comment thread.
    parentAuthor *IdentitySet;
    // The identities of the users participating in this comment thread.
    participants []IdentitySet;
}
// Instantiates a new commentAction and sets the default values.
func NewCommentAction()(*CommentAction) {
    m := &CommentAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommentAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the isReply property value. If true, this activity was a reply to an existing comment thread.
func (m *CommentAction) GetIsReply()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReply
    }
}
// Gets the parentAuthor property value. The identity of the user who started the comment thread.
func (m *CommentAction) GetParentAuthor()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.parentAuthor
    }
}
// Gets the participants property value. The identities of the users participating in this comment thread.
func (m *CommentAction) GetParticipants()([]IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
// The deserialization information for the current model
func (m *CommentAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isReply"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsReply(val)
        return nil
    }
    res["parentAuthor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetParentAuthor(val.(*IdentitySet))
        return nil
    }
    res["participants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        res := make([]IdentitySet, len(val))
        for i, v := range val {
            res[i] = *(v.(*IdentitySet))
        }
        m.SetParticipants(res)
        return nil
    }
    return res
}
func (m *CommentAction) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CommentAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isReply", m.GetIsReply())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parentAuthor", m.GetParentAuthor())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetParticipants()))
        for i, v := range m.GetParticipants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("participants", cast)
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
func (m *CommentAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the isReply property value. If true, this activity was a reply to an existing comment thread.
// Parameters:
//  - value : Value to set for the isReply property.
func (m *CommentAction) SetIsReply(value *bool)() {
    m.isReply = value
}
// Sets the parentAuthor property value. The identity of the user who started the comment thread.
// Parameters:
//  - value : Value to set for the parentAuthor property.
func (m *CommentAction) SetParentAuthor(value *IdentitySet)() {
    m.parentAuthor = value
}
// Sets the participants property value. The identities of the users participating in this comment thread.
// Parameters:
//  - value : Value to set for the participants property.
func (m *CommentAction) SetParticipants(value []IdentitySet)() {
    m.participants = value
}
