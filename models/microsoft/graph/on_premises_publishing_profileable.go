package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnPremisesPublishingProfileable 
type OnPremisesPublishingProfileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAgentGroups()([]OnPremisesAgentGroupable)
    GetAgents()([]OnPremisesAgentable)
    GetConnectorGroups()([]ConnectorGroupable)
    GetConnectors()([]Connectorable)
    GetHybridAgentUpdaterConfiguration()(HybridAgentUpdaterConfigurationable)
    GetIsEnabled()(*bool)
    GetPublishedResources()([]PublishedResourceable)
    SetAgentGroups(value []OnPremisesAgentGroupable)()
    SetAgents(value []OnPremisesAgentable)()
    SetConnectorGroups(value []ConnectorGroupable)()
    SetConnectors(value []Connectorable)()
    SetHybridAgentUpdaterConfiguration(value HybridAgentUpdaterConfigurationable)()
    SetIsEnabled(value *bool)()
    SetPublishedResources(value []PublishedResourceable)()
}
