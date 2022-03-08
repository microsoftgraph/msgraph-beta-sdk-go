package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MicrosoftTunnelSiteable 
type MicrosoftTunnelSiteable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetInternalNetworkProbeUrl()(*string)
    GetMicrosoftTunnelConfiguration()(MicrosoftTunnelConfigurationable)
    GetMicrosoftTunnelServers()([]MicrosoftTunnelServerable)
    GetPublicAddress()(*string)
    GetRoleScopeTagIds()([]string)
    GetUpgradeAutomatically()(*bool)
    GetUpgradeAvailable()(*bool)
    GetUpgradeWindowEndTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)
    GetUpgradeWindowStartTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)
    GetUpgradeWindowUtcOffsetInMinutes()(*int32)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetInternalNetworkProbeUrl(value *string)()
    SetMicrosoftTunnelConfiguration(value MicrosoftTunnelConfigurationable)()
    SetMicrosoftTunnelServers(value []MicrosoftTunnelServerable)()
    SetPublicAddress(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetUpgradeAutomatically(value *bool)()
    SetUpgradeAvailable(value *bool)()
    SetUpgradeWindowEndTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)()
    SetUpgradeWindowStartTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)()
    SetUpgradeWindowUtcOffsetInMinutes(value *int32)()
}
