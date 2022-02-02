package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnPremisesPublishingProfile 
type OnPremisesPublishingProfile struct {
    Entity
    // List of existing onPremisesAgentGroup objects. Read-only. Nullable.
    agentGroups []OnPremisesAgentGroup;
    // List of existing onPremisesAgent objects. Read-only. Nullable.
    agents []OnPremisesAgent;
    // List of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
    connectorGroups []ConnectorGroup;
    // List of existing connector objects for applications published through Application Proxy. Read-only. Nullable.
    connectors []Connector;
    // Represents a hybridAgentUpdaterConfiguration object.
    hybridAgentUpdaterConfiguration *HybridAgentUpdaterConfiguration;
    // Represents if Azure AD Application Proxy is enabled for the tenant.
    isEnabled *bool;
    // List of existing publishedResource objects. Read-only. Nullable.
    publishedResources []PublishedResource;
}
// NewOnPremisesPublishingProfile instantiates a new onPremisesPublishingProfile and sets the default values.
func NewOnPremisesPublishingProfile()(*OnPremisesPublishingProfile) {
    m := &OnPremisesPublishingProfile{
        Entity: *NewEntity(),
    }
    return m
}
// GetAgentGroups gets the agentGroups property value. List of existing onPremisesAgentGroup objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetAgentGroups()([]OnPremisesAgentGroup) {
    if m == nil {
        return nil
    } else {
        return m.agentGroups
    }
}
// GetAgents gets the agents property value. List of existing onPremisesAgent objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetAgents()([]OnPremisesAgent) {
    if m == nil {
        return nil
    } else {
        return m.agents
    }
}
// GetConnectorGroups gets the connectorGroups property value. List of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetConnectorGroups()([]ConnectorGroup) {
    if m == nil {
        return nil
    } else {
        return m.connectorGroups
    }
}
// GetConnectors gets the connectors property value. List of existing connector objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetConnectors()([]Connector) {
    if m == nil {
        return nil
    } else {
        return m.connectors
    }
}
// GetHybridAgentUpdaterConfiguration gets the hybridAgentUpdaterConfiguration property value. Represents a hybridAgentUpdaterConfiguration object.
func (m *OnPremisesPublishingProfile) GetHybridAgentUpdaterConfiguration()(*HybridAgentUpdaterConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.hybridAgentUpdaterConfiguration
    }
}
// GetIsEnabled gets the isEnabled property value. Represents if Azure AD Application Proxy is enabled for the tenant.
func (m *OnPremisesPublishingProfile) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetPublishedResources gets the publishedResources property value. List of existing publishedResource objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetPublishedResources()([]PublishedResource) {
    if m == nil {
        return nil
    } else {
        return m.publishedResources
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesPublishingProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesAgentGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesAgentGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnPremisesAgentGroup))
            }
            m.SetAgentGroups(res)
        }
        return nil
    }
    res["agents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesAgent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesAgent, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnPremisesAgent))
            }
            m.SetAgents(res)
        }
        return nil
    }
    res["connectorGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnectorGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectorGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConnectorGroup))
            }
            m.SetConnectorGroups(res)
        }
        return nil
    }
    res["connectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnector() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Connector, len(val))
            for i, v := range val {
                res[i] = *(v.(*Connector))
            }
            m.SetConnectors(res)
        }
        return nil
    }
    res["hybridAgentUpdaterConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHybridAgentUpdaterConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHybridAgentUpdaterConfiguration(val.(*HybridAgentUpdaterConfiguration))
        }
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["publishedResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublishedResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PublishedResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*PublishedResource))
            }
            m.SetPublishedResources(res)
        }
        return nil
    }
    return res
}
func (m *OnPremisesPublishingProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OnPremisesPublishingProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAgentGroups() != nil {
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
    if m.GetAgents() != nil {
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
    if m.GetConnectorGroups() != nil {
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
    if m.GetConnectors() != nil {
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
    if m.GetPublishedResources() != nil {
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
// SetAgentGroups sets the agentGroups property value. List of existing onPremisesAgentGroup objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetAgentGroups(value []OnPremisesAgentGroup)() {
    if m != nil {
        m.agentGroups = value
    }
}
// SetAgents sets the agents property value. List of existing onPremisesAgent objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetAgents(value []OnPremisesAgent)() {
    if m != nil {
        m.agents = value
    }
}
// SetConnectorGroups sets the connectorGroups property value. List of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetConnectorGroups(value []ConnectorGroup)() {
    if m != nil {
        m.connectorGroups = value
    }
}
// SetConnectors sets the connectors property value. List of existing connector objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetConnectors(value []Connector)() {
    if m != nil {
        m.connectors = value
    }
}
// SetHybridAgentUpdaterConfiguration sets the hybridAgentUpdaterConfiguration property value. Represents a hybridAgentUpdaterConfiguration object.
func (m *OnPremisesPublishingProfile) SetHybridAgentUpdaterConfiguration(value *HybridAgentUpdaterConfiguration)() {
    if m != nil {
        m.hybridAgentUpdaterConfiguration = value
    }
}
// SetIsEnabled sets the isEnabled property value. Represents if Azure AD Application Proxy is enabled for the tenant.
func (m *OnPremisesPublishingProfile) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetPublishedResources sets the publishedResources property value. List of existing publishedResource objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetPublishedResources(value []PublishedResource)() {
    if m != nil {
        m.publishedResources = value
    }
}
