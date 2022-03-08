package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkNetworkConfigurationable 
type TeamworkNetworkConfigurationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDefaultGateway()(*string)
    GetDomainName()(*string)
    GetHostName()(*string)
    GetIpAddress()(*string)
    GetIsDhcpEnabled()(*bool)
    GetIsPCPortEnabled()(*bool)
    GetPrimaryDns()(*string)
    GetSecondaryDns()(*string)
    GetSubnetMask()(*string)
    SetDefaultGateway(value *string)()
    SetDomainName(value *string)()
    SetHostName(value *string)()
    SetIpAddress(value *string)()
    SetIsDhcpEnabled(value *bool)()
    SetIsPCPortEnabled(value *bool)()
    SetPrimaryDns(value *string)()
    SetSecondaryDns(value *string)()
    SetSubnetMask(value *string)()
}
