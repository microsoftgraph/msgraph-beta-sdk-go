package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommentAction 
type CommentAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If true, this activity was a reply to an existing comment thread.
    isReply *bool
    // The identity of the user who started the comment thread.
    parentAuthor IdentitySetable
    // The identities of the users participating in this comment thread.
    participants []IdentitySetable
}
// NewCommentAction instantiates a new commentAction and sets the default values.
func NewCommentAction()(*CommentAction) {
    m := &CommentAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCommentActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommentActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommentAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommentAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommentAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isReply"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReply(val)
        }
        return nil
    }
    res["parentAuthor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentAuthor(val.(IdentitySetable))
        }
        return nil
    }
    res["participants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentitySetable, len(val))
            for i, v := range val {
                res[i] = v.(IdentitySetable)
            }
            m.SetParticipants(res)
        }
        return nil
    }
    return res
}
// GetIsReply gets the isReply property value. If true, this activity was a reply to an existing comment thread.
func (m *CommentAction) GetIsReply()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReply
    }
}
// GetParentAuthor gets the parentAuthor property value. The identity of the user who started the comment thread.
func (m *CommentAction) GetParentAuthor()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.parentAuthor
    }
}
// GetParticipants gets the participants property value. The identities of the users participating in this comment thread.
func (m *CommentAction) GetParticipants()([]IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
// Serialize serializes information the current object
func (m *CommentAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetParticipants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParticipants()))
        for i, v := range m.GetParticipants() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommentAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsReply sets the isReply property value. If true, this activity was a reply to an existing comment thread.
func (m *CommentAction) SetIsReply(value *bool)() {
    if m != nil {
        m.isReply = value
    }
}
// SetParentAuthor sets the parentAuthor property value. The identity of the user who started the comment thread.
func (m *CommentAction) SetParentAuthor(value IdentitySetable)() {
    if m != nil {
        m.parentAuthor = value
    }
}
// SetParticipants sets the participants property value. The identities of the users participating in this comment thread.
func (m *CommentAction) SetParticipants(value []IdentitySetable)() {
    if m != nil {
        m.participants = value
    }
}
