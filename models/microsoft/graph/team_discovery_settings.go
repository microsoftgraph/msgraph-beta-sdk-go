package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeamDiscoverySettings struct {
    additionalData map[string]interface{};
    showInTeamsSearchAndSuggestions *bool;
}
func NewTeamDiscoverySettings()(*TeamDiscoverySettings) {
    m := &TeamDiscoverySettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TeamDiscoverySettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TeamDiscoverySettings) GetShowInTeamsSearchAndSuggestions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showInTeamsSearchAndSuggestions
    }
}
func (m *TeamDiscoverySettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["showInTeamsSearchAndSuggestions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowInTeamsSearchAndSuggestions(val)
        return nil
    }
    return res
}
func (m *TeamDiscoverySettings) IsNil()(bool) {
    return m == nil
}
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
func (m *TeamDiscoverySettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TeamDiscoverySettings) SetShowInTeamsSearchAndSuggestions(value *bool)() {
    m.showInTeamsSearchAndSuggestions = value
}
