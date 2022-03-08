package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnPremisesAgentGroupable 
type OnPremisesAgentGroupable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAgents()([]OnPremisesAgentable)
    GetDisplayName()(*string)
    GetIsDefault()(*bool)
    GetPublishedResources()([]PublishedResourceable)
    GetPublishingType()(*OnPremisesPublishingType)
    SetAgents(value []OnPremisesAgentable)()
    SetDisplayName(value *string)()
    SetIsDefault(value *bool)()
    SetPublishedResources(value []PublishedResourceable)()
    SetPublishingType(value *OnPremisesPublishingType)()
}
