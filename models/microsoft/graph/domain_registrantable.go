package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DomainRegistrantable 
type DomainRegistrantable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCountryOrRegionCode()(*string)
    GetOrganization()(*string)
    GetUrl()(*string)
    GetVendor()(*string)
    SetCountryOrRegionCode(value *string)()
    SetOrganization(value *string)()
    SetUrl(value *string)()
    SetVendor(value *string)()
}
