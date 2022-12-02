package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponse provides operations to call the delta method.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponse struct {
    BaseDeltaFunctionResponse
    // The value property
    value []ChatMessageable
}
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponse instantiates a new TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponse and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponse()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponse) {
    m := &TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponse{
        BaseDeltaFunctionResponse: *NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChatMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageable, len(val))
            for i, v := range val {
                res[i] = v.(ChatMessageable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponse) GetValue()([]ChatMessageable) {
    return m.value
}
// Serialize serializes information the current object
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseDeltaFunctionResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesDeltaResponse) SetValue(value []ChatMessageable)() {
    m.value = value
}
