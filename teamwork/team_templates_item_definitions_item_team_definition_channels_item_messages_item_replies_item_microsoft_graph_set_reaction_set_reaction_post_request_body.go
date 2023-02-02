package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody 
type TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The reactionType property
    reactionType *string
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody()(*TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) GetReactionType()(*string) {
    return m.reactionType
}
// Serialize serializes information the current object
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReactionType sets the reactionType property value. The reactionType property
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBody) SetReactionType(value *string)() {
    m.reactionType = value
}
