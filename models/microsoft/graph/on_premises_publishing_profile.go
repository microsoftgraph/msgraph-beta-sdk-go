package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OnPremisesPublishingProfile struct {
    Entity
    agentGroups []OnPremisesAgentGroup;
    agents []OnPremisesAgent;
    connectorGroups []ConnectorGroup;
    connectors []Connector;
    hybridAgentUpdaterConfiguration *HybridAgentUpdaterConfiguration;
    isEnabled *bool;
    publishedResources []PublishedResource;
}
func NewOnPremisesPublishingProfile()(*OnPremisesPublishingProfile) {
    m := &OnPremisesPublishingProfile{
        Entity: *NewEntity(),
    }
    return m
}
func (m *OnPremisesPublishingProfile) GetAgentGroups()([]OnPremisesAgentGroup) {
    if m == nil {
        return nil
    } else {
        return m.agentGroups
    }
}
func (m *OnPremisesPublishingProfile) GetAgents()([]OnPremisesAgent) {
    if m == nil {
        return nil
    } else {
        return m.agents
    }
}
func (m *OnPremisesPublishingProfile) GetConnectorGroups()([]ConnectorGroup) {
    if m == nil {
        return nil
    } else {
        return m.connectorGroups
    }
}
func (m *OnPremisesPublishingProfile) GetConnectors()([]Connector) {
    if m == nil {
        return nil
    } else {
        return m.connectors
    }
}
func (m *OnPremisesPublishingProfile) GetHybridAgentUpdaterConfiguration()(*HybridAgentUpdaterConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.hybridAgentUpdaterConfiguration
    }
}
func (m *OnPremisesPublishingProfile) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
func (m *OnPremisesPublishingProfile) GetPublishedResources()([]PublishedResource) {
    if m == nil {
        return nil
    } else {
        return m.publishedResources
    }
}
func (m *OnPremisesPublishingProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesAgentGroup() })
        if err != nil {
            return err
        }
        res := make([]OnPremisesAgentGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnPremisesAgentGroup))
        }
        m.SetAgentGroups(res)
        return nil
    }
    res["agents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesAgent() })
        if err != nil {
            return err
        }
        res := make([]OnPremisesAgent, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnPremisesAgent))
        }
        m.SetAgents(res)
        return nil
    }
    res["connectorGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnectorGroup() })
        if err != nil {
            return err
        }
        res := make([]ConnectorGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConnectorGroup))
        }
        m.SetConnectorGroups(res)
        return nil
    }
    res["connectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnector() })
        if err != nil {
            return err
        }
        res := make([]Connector, len(val))
        for i, v := range val {
            res[i] = *(v.(*Connector))
        }
        m.SetConnectors(res)
        return nil
    }
    res["hybridAgentUpdaterConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHybridAgentUpdaterConfiguration() })
        if err != nil {
            return err
        }
        m.SetHybridAgentUpdaterConfiguration(val.(*HybridAgentUpdaterConfiguration))
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabled(val)
        return nil
    }
    res["publishedResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublishedResource() })
        if err != nil {
            return err
        }
        res := make([]PublishedResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*PublishedResource))
        }
        m.SetPublishedResources(res)
        return nil
    }
    return res
}
func (m *OnPremisesPublishingProfile) IsNil()(bool) {
    return m == nil
}
func (m *OnPremisesPublishingProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgentGroups()))
        for i, v := range m.GetAgentGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("agentGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgents()))
        for i, v := range m.GetAgents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("agents", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConnectorGroups()))
        for i, v := range m.GetConnectorGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("connectorGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConnectors()))
        for i, v := range m.GetConnectors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("connectors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("hybridAgentUpdaterConfiguration", m.GetHybridAgentUpdaterConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPublishedResources()))
        for i, v := range m.GetPublishedResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("publishedResources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *OnPremisesPublishingProfile) SetAgentGroups(value []OnPremisesAgentGroup)() {
    m.agentGroups = value
}
func (m *OnPremisesPublishingProfile) SetAgents(value []OnPremisesAgent)() {
    m.agents = value
}
func (m *OnPremisesPublishingProfile) SetConnectorGroups(value []ConnectorGroup)() {
    m.connectorGroups = value
}
func (m *OnPremisesPublishingProfile) SetConnectors(value []Connector)() {
    m.connectors = value
}
func (m *OnPremisesPublishingProfile) SetHybridAgentUpdaterConfiguration(value *HybridAgentUpdaterConfiguration)() {
    m.hybridAgentUpdaterConfiguration = value
}
func (m *OnPremisesPublishingProfile) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
func (m *OnPremisesPublishingProfile) SetPublishedResources(value []PublishedResource)() {
    m.publishedResources = value
}
