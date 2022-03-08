package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// NetworkInterfaceable 
type NetworkInterfaceable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetIpV4Address()(*string)
    GetIpV6Address()(*string)
    GetLocalIpV6Address()(*string)
    GetMacAddress()(*string)
    SetDescription(value *string)()
    SetIpV4Address(value *string)()
    SetIpV6Address(value *string)()
    SetLocalIpV6Address(value *string)()
    SetMacAddress(value *string)()
}
