package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody 
type ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The reactionType property
    reactionType *string
}
// NewItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody instantiates a new ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody and sets the default values.
func NewItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody()(*ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody) {
    m := &ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody) GetReactionType()(*string) {
    return m.reactionType
}
// Serialize serializes information the current object
func (m *ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReactionType sets the reactionType property value. The reactionType property
func (m *ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphUnsetReactionUnsetReactionPostRequestBody) SetReactionType(value *string)() {
    m.reactionType = value
}
