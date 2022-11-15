package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Picture provides operations to manage the collection of accessReviewDecision entities.
type Picture struct {
    Entity
    // The content property
    content []byte
    // The contentType property
    contentType *string
    // The height property
    height *int32
    // The width property
    width *int32
}
// NewPicture instantiates a new picture and sets the default values.
func NewPicture()(*Picture) {
    m := &Picture{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.picture";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePictureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePictureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPicture(), nil
}
// GetContent gets the content property value. The content property
func (m *Picture) GetContent()([]byte) {
    return m.content
}
// GetContentType gets the contentType property value. The contentType property
func (m *Picture) GetContentType()(*string) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Picture) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetContent)
    res["contentType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContentType)
    res["height"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetHeight)
    res["width"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWidth)
    return res
}
// GetHeight gets the height property value. The height property
func (m *Picture) GetHeight()(*int32) {
    return m.height
}
// GetWidth gets the width property value. The width property
func (m *Picture) GetWidth()(*int32) {
    return m.width
}
// Serialize serializes information the current object
func (m *Picture) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("width", m.GetWidth())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content property
func (m *Picture) SetContent(value []byte)() {
    m.content = value
}
// SetContentType sets the contentType property value. The contentType property
func (m *Picture) SetContentType(value *string)() {
    m.contentType = value
}
// SetHeight sets the height property value. The height property
func (m *Picture) SetHeight(value *int32)() {
    m.height = value
}
// SetWidth sets the width property value. The width property
func (m *Picture) SetWidth(value *int32)() {
    m.width = value
}
