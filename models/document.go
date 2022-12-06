package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Document 
type Document struct {
    Entity
    // The comments property
    comments []DocumentCommentable
}
// NewDocument instantiates a new Document and sets the default values.
func NewDocument()(*Document) {
    m := &Document{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDocumentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDocumentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDocument(), nil
}
// GetComments gets the comments property value. The comments property
func (m *Document) GetComments()([]DocumentCommentable) {
    return m.comments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Document) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["comments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDocumentCommentFromDiscriminatorValue , m.SetComments)
    return res
}
// Serialize serializes information the current object
func (m *Document) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetComments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetComments())
        err = writer.WriteCollectionOfObjectValues("comments", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComments sets the comments property value. The comments property
func (m *Document) SetComments(value []DocumentCommentable)() {
    m.comments = value
}
