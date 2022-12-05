package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody provides operations to call the clockIn method.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The atApprovedLocation property
    atApprovedLocation *bool
    // The notes property
    notes ItemBodyable
    // The onBehalfOfUserId property
    onBehalfOfUserId *string
}
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody instantiates a new TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody) {
    m := &TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAtApprovedLocation gets the atApprovedLocation property value. The atApprovedLocation property
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody) GetAtApprovedLocation()(*bool) {
    return m.atApprovedLocation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["atApprovedLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAtApprovedLocation(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val.(ItemBodyable))
        }
        return nil
    }
    res["onBehalfOfUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnBehalfOfUserId(val)
        }
        return nil
    }
    return res
}
// GetNotes gets the notes property value. The notes property
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody) GetNotes()(ItemBodyable) {
    return m.notes
}
// GetOnBehalfOfUserId gets the onBehalfOfUserId property value. The onBehalfOfUserId property
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody) GetOnBehalfOfUserId()(*string) {
    return m.onBehalfOfUserId
}
// Serialize serializes information the current object
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("atApprovedLocation", m.GetAtApprovedLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onBehalfOfUserId", m.GetOnBehalfOfUserId())
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAtApprovedLocation sets the atApprovedLocation property value. The atApprovedLocation property
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody) SetAtApprovedLocation(value *bool)() {
    m.atApprovedLocation = value
}
// SetNotes sets the notes property value. The notes property
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody) SetNotes(value ItemBodyable)() {
    m.notes = value
}
// SetOnBehalfOfUserId sets the onBehalfOfUserId property value. The onBehalfOfUserId property
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleTimeCardsClockInPostRequestBody) SetOnBehalfOfUserId(value *string)() {
    m.onBehalfOfUserId = value
}
