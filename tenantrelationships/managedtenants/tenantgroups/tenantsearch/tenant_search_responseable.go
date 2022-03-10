package tenantsearch

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// TenantSearchResponseable 
type TenantSearchResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetValue()([]i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantGroupable)
    SetValue(value []i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantGroupable)()
}
