package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody 
type ItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The reactionType property
    reactionType *string
}
// NewItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody instantiates a new ItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody and sets the default values.
func NewItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody()(*ItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) {
    m := &ItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) GetReactionType()(*string) {
    return m.reactionType
}
// Serialize serializes information the current object
func (m *ItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReactionType sets the reactionType property value. The reactionType property
func (m *ItemTeamChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) SetReactionType(value *string)() {
    m.reactionType = value
}
