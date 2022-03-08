package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantContractable 
type TenantContractable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContractType()(*int32)
    GetDefaultDomainName()(*string)
    GetDisplayName()(*string)
    SetContractType(value *int32)()
    SetDefaultDomainName(value *string)()
    SetDisplayName(value *string)()
}
