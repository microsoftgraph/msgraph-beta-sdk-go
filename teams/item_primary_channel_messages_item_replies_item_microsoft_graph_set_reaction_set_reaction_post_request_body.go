package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody 
type ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The reactionType property
    reactionType *string
}
// NewItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody instantiates a new ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody and sets the default values.
func NewItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody()(*ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) {
    m := &ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reactionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReactionType(val)
        }
        return nil
    }
    return res
}
// GetReactionType gets the reactionType property value. The reactionType property
func (m *ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) GetReactionType()(*string) {
    return m.reactionType
}
// Serialize serializes information the current object
func (m *ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("reactionType", m.GetReactionType())
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
func (m *ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReactionType sets the reactionType property value. The reactionType property
func (m *ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) SetReactionType(value *string)() {
    m.reactionType = value
}
