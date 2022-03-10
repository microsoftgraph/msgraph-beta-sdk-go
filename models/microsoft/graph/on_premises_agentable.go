package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnPremisesAgentable 
type OnPremisesAgentable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAgentGroups()([]OnPremisesAgentGroupable)
    GetExternalIp()(*string)
    GetMachineName()(*string)
    GetStatus()(*AgentStatus)
    GetSupportedPublishingTypes()([]OnPremisesPublishingType)
    SetAgentGroups(value []OnPremisesAgentGroupable)()
    SetExternalIp(value *string)()
    SetMachineName(value *string)()
    SetStatus(value *AgentStatus)()
    SetSupportedPublishingTypes(value []OnPremisesPublishingType)()
}
