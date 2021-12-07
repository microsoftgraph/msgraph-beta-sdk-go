package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamDiscoverySettings 
type TeamDiscoverySettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If set to true, the team is visible via search and suggestions from the Teams client.
    showInTeamsSearchAndSuggestions *bool;
}
// NewTeamDiscoverySettings instantiates a new teamDiscoverySettings and sets the default values.
func NewTeamDiscoverySettings()(*TeamDiscoverySettings) {
    m := &TeamDiscoverySettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamDiscoverySettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetShowInTeamsSearchAndSuggestions gets the showInTeamsSearchAndSuggestions property value. If set to true, the team is visible via search and suggestions from the Teams client.
func (m *TeamDiscoverySettings) GetShowInTeamsSearchAndSuggestions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showInTeamsSearchAndSuggestions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamDiscoverySettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["showInTeamsSearchAndSuggestions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowInTeamsSearchAndSuggestions(val)
        }
        return nil
    }
    return res
}
func (m *TeamDiscoverySettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamDiscoverySettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.additionalData = value
    }
}
// SetShowInTeamsSearchAndSuggestions sets the showInTeamsSearchAndSuggestions property value. If set to true, the team is visible via search and suggestions from the Teams client.
func (m *TeamDiscoverySettings) SetShowInTeamsSearchAndSuggestions(value *bool)() {
    if m != nil {
        m.showInTeamsSearchAndSuggestions = value
    }
}
