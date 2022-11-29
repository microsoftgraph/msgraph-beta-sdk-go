package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DocumentCommentReply provides operations to manage the collection of accessReview entities.
type DocumentCommentReply struct {
    Entity
    // The content property
    content *string
    // The location property
    location *string
}
// NewDocumentCommentReply instantiates a new documentCommentReply and sets the default values.
func NewDocumentCommentReply()(*DocumentCommentReply) {
    m := &DocumentCommentReply{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDocumentCommentReplyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDocumentCommentReplyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDocumentCommentReply(), nil
}
// GetContent gets the content property value. The content property
func (m *DocumentCommentReply) GetContent()(*string) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DocumentCommentReply) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContent)
    res["location"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLocation)
    return res
}
// GetLocation gets the location property value. The location property
func (m *DocumentCommentReply) GetLocation()(*string) {
    return m.location
}
// Serialize serializes information the current object
func (m *DocumentCommentReply) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content property
func (m *DocumentCommentReply) SetContent(value *string)() {
    m.content = value
}
// SetLocation sets the location property value. The location property
func (m *DocumentCommentReply) SetLocation(value *string)() {
    m.location = value
}
