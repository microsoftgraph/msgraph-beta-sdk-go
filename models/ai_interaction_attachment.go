// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AiInteractionAttachment struct {
    Entity
}
// NewAiInteractionAttachment instantiates a new AiInteractionAttachment and sets the default values.
func NewAiInteractionAttachment()(*AiInteractionAttachment) {
    m := &AiInteractionAttachment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAiInteractionAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAiInteractionAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAiInteractionAttachment(), nil
}
// GetAttachmentId gets the attachmentId property value. The identifier for the attachment. This identifier is only unique within the message scope.
// returns a *string when successful
func (m *AiInteractionAttachment) GetAttachmentId()(*string) {
    val, err := m.GetBackingStore().Get("attachmentId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContent gets the content property value. The content of the attachment.
// returns a *string when successful
func (m *AiInteractionAttachment) GetContent()(*string) {
    val, err := m.GetBackingStore().Get("content")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContentType gets the contentType property value. The type of the content. For example, reference, file, and image/imageType.
// returns a *string when successful
func (m *AiInteractionAttachment) GetContentType()(*string) {
    val, err := m.GetBackingStore().Get("contentType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContentUrl gets the contentUrl property value. The URL of the content.
// returns a *string when successful
func (m *AiInteractionAttachment) GetContentUrl()(*string) {
    val, err := m.GetBackingStore().Get("contentUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AiInteractionAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attachmentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentId(val)
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["contentUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentUrl(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the attachment.
// returns a *string when successful
func (m *AiInteractionAttachment) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AiInteractionAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("attachmentId", m.GetAttachmentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("content", m.GetContent())
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
        err = writer.WriteStringValue("contentUrl", m.GetContentUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttachmentId sets the attachmentId property value. The identifier for the attachment. This identifier is only unique within the message scope.
func (m *AiInteractionAttachment) SetAttachmentId(value *string)() {
    err := m.GetBackingStore().Set("attachmentId", value)
    if err != nil {
        panic(err)
    }
}
// SetContent sets the content property value. The content of the attachment.
func (m *AiInteractionAttachment) SetContent(value *string)() {
    err := m.GetBackingStore().Set("content", value)
    if err != nil {
        panic(err)
    }
}
// SetContentType sets the contentType property value. The type of the content. For example, reference, file, and image/imageType.
func (m *AiInteractionAttachment) SetContentType(value *string)() {
    err := m.GetBackingStore().Set("contentType", value)
    if err != nil {
        panic(err)
    }
}
// SetContentUrl sets the contentUrl property value. The URL of the content.
func (m *AiInteractionAttachment) SetContentUrl(value *string)() {
    err := m.GetBackingStore().Set("contentUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name of the attachment.
func (m *AiInteractionAttachment) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
type AiInteractionAttachmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachmentId()(*string)
    GetContent()(*string)
    GetContentType()(*string)
    GetContentUrl()(*string)
    GetName()(*string)
    SetAttachmentId(value *string)()
    SetContent(value *string)()
    SetContentType(value *string)()
    SetContentUrl(value *string)()
    SetName(value *string)()
}
