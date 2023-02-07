package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody 
type MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The Message property
    message ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable
}
// NewMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody instantiates a new MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody and sets the default values.
func NewMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody()(*MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) {
    m := &MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable))
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The Message property
func (m *MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) GetMessage()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable) {
    return m.message
}
// Serialize serializes information the current object
func (m *MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
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
func (m *MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetMessage sets the message property value. The Message property
func (m *MailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) SetMessage(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable)() {
    m.message = value
}
