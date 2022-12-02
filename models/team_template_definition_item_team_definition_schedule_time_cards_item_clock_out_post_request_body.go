package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody provides operations to call the clockOut method.
type TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The atApprovedLocation property
    atApprovedLocation *bool
    // The notes property
    notes ItemBodyable
}
// NewTeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody instantiates a new TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody()(*TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody) {
    m := &TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAtApprovedLocation gets the atApprovedLocation property value. The atApprovedLocation property
func (m *TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody) GetAtApprovedLocation()(*bool) {
    return m.atApprovedLocation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetNotes gets the notes property value. The notes property
func (m *TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody) GetNotes()(ItemBodyable) {
    return m.notes
}
// Serialize serializes information the current object
func (m *TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAtApprovedLocation sets the atApprovedLocation property value. The atApprovedLocation property
func (m *TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody) SetAtApprovedLocation(value *bool)() {
    m.atApprovedLocation = value
}
// SetNotes sets the notes property value. The notes property
func (m *TeamTemplateDefinitionItemTeamDefinitionScheduleTimeCardsItemClockOutPostRequestBody) SetNotes(value ItemBodyable)() {
    m.notes = value
}
