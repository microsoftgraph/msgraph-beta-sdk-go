package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PublishedResourceable 
type PublishedResourceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAgentGroups()([]OnPremisesAgentGroupable)
    GetDisplayName()(*string)
    GetPublishingType()(*OnPremisesPublishingType)
    GetResourceName()(*string)
    SetAgentGroups(value []OnPremisesAgentGroupable)()
    SetDisplayName(value *string)()
    SetPublishingType(value *OnPremisesPublishingType)()
    SetResourceName(value *string)()
}
