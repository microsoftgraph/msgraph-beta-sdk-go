package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesPublishingProfile 
type OnPremisesPublishingProfile struct {
    Entity
    // List of existing onPremisesAgentGroup objects. Read-only. Nullable.
    agentGroups []OnPremisesAgentGroupable;
    // List of existing onPremisesAgent objects. Read-only. Nullable.
    agents []OnPremisesAgentable;
    // List of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
    connectorGroups []ConnectorGroupable;
    // List of existing connector objects for applications published through Application Proxy. Read-only. Nullable.
    connectors []Connectorable;
    // Represents a hybridAgentUpdaterConfiguration object.
    hybridAgentUpdaterConfiguration HybridAgentUpdaterConfigurationable;
    // Represents if Azure AD Application Proxy is enabled for the tenant.
    isEnabled *bool;
    // List of existing publishedResource objects. Read-only. Nullable.
    publishedResources []PublishedResourceable;
}
// NewOnPremisesPublishingProfile instantiates a new onPremisesPublishingProfile and sets the default values.
func NewOnPremisesPublishingProfile()(*OnPremisesPublishingProfile) {
    m := &OnPremisesPublishingProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOnPremisesPublishingProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesPublishingProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesPublishingProfile(), nil
}
// GetAgentGroups gets the agentGroups property value. List of existing onPremisesAgentGroup objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetAgentGroups()([]OnPremisesAgentGroupable) {
    if m == nil {
        return nil
    } else {
        return m.agentGroups
    }
}
// GetAgents gets the agents property value. List of existing onPremisesAgent objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetAgents()([]OnPremisesAgentable) {
    if m == nil {
        return nil
    } else {
        return m.agents
    }
}
// GetConnectorGroups gets the connectorGroups property value. List of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetConnectorGroups()([]ConnectorGroupable) {
    if m == nil {
        return nil
    } else {
        return m.connectorGroups
    }
}
// GetConnectors gets the connectors property value. List of existing connector objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetConnectors()([]Connectorable) {
    if m == nil {
        return nil
    } else {
        return m.connectors
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesPublishingProfile) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentGroups"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesAgentGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesAgentGroupable, len(val))
            for i, v := range val {
                res[i] = v.(OnPremisesAgentGroupable)
            }
            m.SetAgentGroups(res)
        }
        return nil
    }
    res["agents"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesAgentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesAgentable, len(val))
            for i, v := range val {
                res[i] = v.(OnPremisesAgentable)
            }
            m.SetAgents(res)
        }
        return nil
    }
    res["connectorGroups"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectorGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectorGroupable, len(val))
            for i, v := range val {
                res[i] = v.(ConnectorGroupable)
            }
            m.SetConnectorGroups(res)
        }
        return nil
    }
    res["connectors"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Connectorable, len(val))
            for i, v := range val {
                res[i] = v.(Connectorable)
            }
            m.SetConnectors(res)
        }
        return nil
    }
    res["hybridAgentUpdaterConfiguration"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHybridAgentUpdaterConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHybridAgentUpdaterConfiguration(val.(HybridAgentUpdaterConfigurationable))
        }
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["publishedResources"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePublishedResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PublishedResourceable, len(val))
            for i, v := range val {
                res[i] = v.(PublishedResourceable)
            }
            m.SetPublishedResources(res)
        }
        return nil
    }
    return res
}
// GetHybridAgentUpdaterConfiguration gets the hybridAgentUpdaterConfiguration property value. Represents a hybridAgentUpdaterConfiguration object.
func (m *OnPremisesPublishingProfile) GetHybridAgentUpdaterConfiguration()(HybridAgentUpdaterConfigurationable) {
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
func (m *OnPremisesPublishingProfile) GetPublishedResources()([]PublishedResourceable) {
    if m == nil {
        return nil
    } else {
        return m.publishedResources
    }
}
// Serialize serializes information the current object
func (m *OnPremisesPublishingProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAgentGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAgentGroups()))
        for i, v := range m.GetAgentGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("agentGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAgents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAgents()))
        for i, v := range m.GetAgents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("agents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConnectorGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConnectorGroups()))
        for i, v := range m.GetConnectorGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("connectorGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConnectors()))
        for i, v := range m.GetConnectors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPublishedResources()))
        for i, v := range m.GetPublishedResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("publishedResources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAgentGroups sets the agentGroups property value. List of existing onPremisesAgentGroup objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetAgentGroups(value []OnPremisesAgentGroupable)() {
    if m != nil {
        m.agentGroups = value
    }
}
// SetAgents sets the agents property value. List of existing onPremisesAgent objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetAgents(value []OnPremisesAgentable)() {
    if m != nil {
        m.agents = value
    }
}
// SetConnectorGroups sets the connectorGroups property value. List of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetConnectorGroups(value []ConnectorGroupable)() {
    if m != nil {
        m.connectorGroups = value
    }
}
// SetConnectors sets the connectors property value. List of existing connector objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetConnectors(value []Connectorable)() {
    if m != nil {
        m.connectors = value
    }
}
// SetHybridAgentUpdaterConfiguration sets the hybridAgentUpdaterConfiguration property value. Represents a hybridAgentUpdaterConfiguration object.
func (m *OnPremisesPublishingProfile) SetHybridAgentUpdaterConfiguration(value HybridAgentUpdaterConfigurationable)() {
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
func (m *OnPremisesPublishingProfile) SetPublishedResources(value []PublishedResourceable)() {
    if m != nil {
        m.publishedResources = value
    }
}
