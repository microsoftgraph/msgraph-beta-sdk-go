package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody provides operations to call the unsetReaction method.
type TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The reactionType property
    reactionType *string
}
// NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody instantiates a new TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody) {
    m := &TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody) GetReactionType()(*string) {
    return m.reactionType
}
// Serialize serializes information the current object
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetReactionType sets the reactionType property value. The reactionType property
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionPostRequestBody) SetReactionType(value *string)() {
    m.reactionType = value
}
