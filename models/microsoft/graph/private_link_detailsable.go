package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivateLinkDetailsable 
type PrivateLinkDetailsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetPolicyId()(*string)
    GetPolicyName()(*string)
    GetPolicyTenantId()(*string)
    GetResourceId()(*string)
    SetPolicyId(value *string)()
    SetPolicyName(value *string)()
    SetPolicyTenantId(value *string)()
    SetResourceId(value *string)()
}
