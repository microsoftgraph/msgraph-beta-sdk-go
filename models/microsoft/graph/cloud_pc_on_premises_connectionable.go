package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcOnPremisesConnectionable 
type CloudPcOnPremisesConnectionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdDomainName()(*string)
    GetAdDomainPassword()(*string)
    GetAdDomainUsername()(*string)
    GetDisplayName()(*string)
    GetHealthCheckStatus()(*CloudPcOnPremisesConnectionStatus)
    GetHealthCheckStatusDetails()(CloudPcOnPremisesConnectionStatusDetailsable)
    GetInUse()(*bool)
    GetOrganizationalUnit()(*string)
    GetResourceGroupId()(*string)
    GetSubnetId()(*string)
    GetSubscriptionId()(*string)
    GetSubscriptionName()(*string)
    GetType()(*CloudPcOnPremisesConnectionType)
    GetVirtualNetworkId()(*string)
    SetAdDomainName(value *string)()
    SetAdDomainPassword(value *string)()
    SetAdDomainUsername(value *string)()
    SetDisplayName(value *string)()
    SetHealthCheckStatus(value *CloudPcOnPremisesConnectionStatus)()
    SetHealthCheckStatusDetails(value CloudPcOnPremisesConnectionStatusDetailsable)()
    SetInUse(value *bool)()
    SetOrganizationalUnit(value *string)()
    SetResourceGroupId(value *string)()
    SetSubnetId(value *string)()
    SetSubscriptionId(value *string)()
    SetSubscriptionName(value *string)()
    SetType(value *CloudPcOnPremisesConnectionType)()
    SetVirtualNetworkId(value *string)()
}
