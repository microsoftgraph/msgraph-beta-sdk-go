package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommentAction 
type CommentAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If true, this activity was a reply to an existing comment thread.
    isReply *bool
    // The OdataType property
    odataType *string
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
    odataTypeValue := "#microsoft.graph.commentAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCommentActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommentActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommentAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommentAction) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommentAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isReply"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsReply)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["parentAuthor"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetParentAuthor)
    res["participants"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIdentitySetFromDiscriminatorValue , m.SetParticipants)
    return res
}
// GetIsReply gets the isReply property value. If true, this activity was a reply to an existing comment thread.
func (m *CommentAction) GetIsReply()(*bool) {
    return m.isReply
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CommentAction) GetOdataType()(*string) {
    return m.odataType
}
// GetParentAuthor gets the parentAuthor property value. The identity of the user who started the comment thread.
func (m *CommentAction) GetParentAuthor()(IdentitySetable) {
    return m.parentAuthor
}
// GetParticipants gets the participants property value. The identities of the users participating in this comment thread.
func (m *CommentAction) GetParticipants()([]IdentitySetable) {
    return m.participants
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetParticipants())
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
    m.additionalData = value
}
// SetIsReply sets the isReply property value. If true, this activity was a reply to an existing comment thread.
func (m *CommentAction) SetIsReply(value *bool)() {
    m.isReply = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CommentAction) SetOdataType(value *string)() {
    m.odataType = value
}
// SetParentAuthor sets the parentAuthor property value. The identity of the user who started the comment thread.
func (m *CommentAction) SetParentAuthor(value IdentitySetable)() {
    m.parentAuthor = value
}
// SetParticipants sets the participants property value. The identities of the users participating in this comment thread.
func (m *CommentAction) SetParticipants(value []IdentitySetable)() {
    m.participants = value
}
