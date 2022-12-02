package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody provides operations to call the unsetReaction method.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The reactionType property
    reactionType *string
}
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody instantiates a new TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody) {
    m := &TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody) GetReactionType()(*string) {
    return m.reactionType
}
// Serialize serializes information the current object
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetReactionType sets the reactionType property value. The reactionType property
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesItemUnsetReactionPostRequestBody) SetReactionType(value *string)() {
    m.reactionType = value
}
