package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MicrosoftTunnelConfigurationable 
type MicrosoftTunnelConfigurationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdvancedSettings()([]KeyValuePairable)
    GetDefaultDomainSuffix()(*string)
    GetDescription()(*string)
    GetDisableUdpConnections()(*bool)
    GetDisplayName()(*string)
    GetDnsServers()([]string)
    GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetListenPort()(*int32)
    GetNetwork()(*string)
    GetRoleScopeTagIds()([]string)
    GetRoutesExclude()([]string)
    GetRoutesInclude()([]string)
    GetSplitDNS()([]string)
    SetAdvancedSettings(value []KeyValuePairable)()
    SetDefaultDomainSuffix(value *string)()
    SetDescription(value *string)()
    SetDisableUdpConnections(value *bool)()
    SetDisplayName(value *string)()
    SetDnsServers(value []string)()
    SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetListenPort(value *int32)()
    SetNetwork(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetRoutesExclude(value []string)()
    SetRoutesInclude(value []string)()
    SetSplitDNS(value []string)()
}
