package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamDiscoverySettings 
type TeamDiscoverySettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // If set to true, the team is visible via search and suggestions from the Teams client.
    showInTeamsSearchAndSuggestions *bool
}
// NewTeamDiscoverySettings instantiates a new teamDiscoverySettings and sets the default values.
func NewTeamDiscoverySettings()(*TeamDiscoverySettings) {
    m := &TeamDiscoverySettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamDiscoverySettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamDiscoverySettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamDiscoverySettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamDiscoverySettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamDiscoverySettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["showInTeamsSearchAndSuggestions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetShowInTeamsSearchAndSuggestions)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamDiscoverySettings) GetOdataType()(*string) {
    return m.odataType
}
// GetShowInTeamsSearchAndSuggestions gets the showInTeamsSearchAndSuggestions property value. If set to true, the team is visible via search and suggestions from the Teams client.
func (m *TeamDiscoverySettings) GetShowInTeamsSearchAndSuggestions()(*bool) {
    return m.showInTeamsSearchAndSuggestions
}
// Serialize serializes information the current object
func (m *TeamDiscoverySettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showInTeamsSearchAndSuggestions", m.GetShowInTeamsSearchAndSuggestions())
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
func (m *TeamDiscoverySettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamDiscoverySettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetShowInTeamsSearchAndSuggestions sets the showInTeamsSearchAndSuggestions property value. If set to true, the team is visible via search and suggestions from the Teams client.
func (m *TeamDiscoverySettings) SetShowInTeamsSearchAndSuggestions(value *bool)() {
    m.showInTeamsSearchAndSuggestions = value
}
