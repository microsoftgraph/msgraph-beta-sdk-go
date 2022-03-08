package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantCustomizedInformationable 
type TenantCustomizedInformationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContacts()([]TenantContactInformationable)
    GetDisplayName()(*string)
    GetTenantId()(*string)
    GetWebsite()(*string)
    SetContacts(value []TenantContactInformationable)()
    SetDisplayName(value *string)()
    SetTenantId(value *string)()
    SetWebsite(value *string)()
}
