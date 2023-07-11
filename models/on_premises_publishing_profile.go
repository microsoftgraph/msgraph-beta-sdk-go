package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesPublishingProfile 
type OnPremisesPublishingProfile struct {
    Entity
    // The OdataType property
    OdataType *string
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
    val, err := m.GetBackingStore().Get("agentGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OnPremisesAgentGroupable)
    }
    return nil
}
// GetAgents gets the agents property value. List of existing onPremisesAgent objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetAgents()([]OnPremisesAgentable) {
    val, err := m.GetBackingStore().Get("agents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OnPremisesAgentable)
    }
    return nil
}
// GetConnectorGroups gets the connectorGroups property value. List of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetConnectorGroups()([]ConnectorGroupable) {
    val, err := m.GetBackingStore().Get("connectorGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ConnectorGroupable)
    }
    return nil
}
// GetConnectors gets the connectors property value. List of existing connector objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetConnectors()([]Connectorable) {
    val, err := m.GetBackingStore().Get("connectors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Connectorable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesPublishingProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesAgentGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesAgentGroupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OnPremisesAgentGroupable)
                }
            }
            m.SetAgentGroups(res)
        }
        return nil
    }
    res["agents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesAgentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesAgentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OnPremisesAgentable)
                }
            }
            m.SetAgents(res)
        }
        return nil
    }
    res["connectorGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectorGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectorGroupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConnectorGroupable)
                }
            }
            m.SetConnectorGroups(res)
        }
        return nil
    }
    res["connectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Connectorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Connectorable)
                }
            }
            m.SetConnectors(res)
        }
        return nil
    }
    res["hybridAgentUpdaterConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHybridAgentUpdaterConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHybridAgentUpdaterConfiguration(val.(HybridAgentUpdaterConfigurationable))
        }
        return nil
    }
    res["isDefaultAccessEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultAccessEnabled(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["publishedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePublishedResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PublishedResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PublishedResourceable)
                }
            }
            m.SetPublishedResources(res)
        }
        return nil
    }
    return res
}
// GetHybridAgentUpdaterConfiguration gets the hybridAgentUpdaterConfiguration property value. Represents a hybridAgentUpdaterConfiguration object.
func (m *OnPremisesPublishingProfile) GetHybridAgentUpdaterConfiguration()(HybridAgentUpdaterConfigurationable) {
    val, err := m.GetBackingStore().Get("hybridAgentUpdaterConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(HybridAgentUpdaterConfigurationable)
    }
    return nil
}
// GetIsDefaultAccessEnabled gets the isDefaultAccessEnabled property value. The isDefaultAccessEnabled property
func (m *OnPremisesPublishingProfile) GetIsDefaultAccessEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isDefaultAccessEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsEnabled gets the isEnabled property value. Represents if Azure AD Application Proxy is enabled for the tenant.
func (m *OnPremisesPublishingProfile) GetIsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPublishedResources gets the publishedResources property value. List of existing publishedResource objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) GetPublishedResources()([]PublishedResourceable) {
    val, err := m.GetBackingStore().Get("publishedResources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PublishedResourceable)
    }
    return nil
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("agentGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAgents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAgents()))
        for i, v := range m.GetAgents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("agents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConnectorGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConnectorGroups()))
        for i, v := range m.GetConnectorGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("connectorGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConnectors()))
        for i, v := range m.GetConnectors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        err = writer.WriteBoolValue("isDefaultAccessEnabled", m.GetIsDefaultAccessEnabled())
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    err := m.GetBackingStore().Set("agentGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetAgents sets the agents property value. List of existing onPremisesAgent objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetAgents(value []OnPremisesAgentable)() {
    err := m.GetBackingStore().Set("agents", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectorGroups sets the connectorGroups property value. List of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetConnectorGroups(value []ConnectorGroupable)() {
    err := m.GetBackingStore().Set("connectorGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectors sets the connectors property value. List of existing connector objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetConnectors(value []Connectorable)() {
    err := m.GetBackingStore().Set("connectors", value)
    if err != nil {
        panic(err)
    }
}
// SetHybridAgentUpdaterConfiguration sets the hybridAgentUpdaterConfiguration property value. Represents a hybridAgentUpdaterConfiguration object.
func (m *OnPremisesPublishingProfile) SetHybridAgentUpdaterConfiguration(value HybridAgentUpdaterConfigurationable)() {
    err := m.GetBackingStore().Set("hybridAgentUpdaterConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDefaultAccessEnabled sets the isDefaultAccessEnabled property value. The isDefaultAccessEnabled property
func (m *OnPremisesPublishingProfile) SetIsDefaultAccessEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isDefaultAccessEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEnabled sets the isEnabled property value. Represents if Azure AD Application Proxy is enabled for the tenant.
func (m *OnPremisesPublishingProfile) SetIsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetPublishedResources sets the publishedResources property value. List of existing publishedResource objects. Read-only. Nullable.
func (m *OnPremisesPublishingProfile) SetPublishedResources(value []PublishedResourceable)() {
    err := m.GetBackingStore().Set("publishedResources", value)
    if err != nil {
        panic(err)
    }
}
// OnPremisesPublishingProfileable 
type OnPremisesPublishingProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgentGroups()([]OnPremisesAgentGroupable)
    GetAgents()([]OnPremisesAgentable)
    GetConnectorGroups()([]ConnectorGroupable)
    GetConnectors()([]Connectorable)
    GetHybridAgentUpdaterConfiguration()(HybridAgentUpdaterConfigurationable)
    GetIsDefaultAccessEnabled()(*bool)
    GetIsEnabled()(*bool)
    GetPublishedResources()([]PublishedResourceable)
    SetAgentGroups(value []OnPremisesAgentGroupable)()
    SetAgents(value []OnPremisesAgentable)()
    SetConnectorGroups(value []ConnectorGroupable)()
    SetConnectors(value []Connectorable)()
    SetHybridAgentUpdaterConfiguration(value HybridAgentUpdaterConfigurationable)()
    SetIsDefaultAccessEnabled(value *bool)()
    SetIsEnabled(value *bool)()
    SetPublishedResources(value []PublishedResourceable)()
}
