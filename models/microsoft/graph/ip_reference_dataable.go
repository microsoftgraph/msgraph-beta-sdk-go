package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IpReferenceDataable 
type IpReferenceDataable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAsn()(*int32)
    GetCity()(*string)
    GetCountryOrRegionCode()(*string)
    GetOrganization()(*string)
    GetState()(*string)
    GetVendor()(*string)
    SetAsn(value *int32)()
    SetCity(value *string)()
    SetCountryOrRegionCode(value *string)()
    SetOrganization(value *string)()
    SetState(value *string)()
    SetVendor(value *string)()
}
