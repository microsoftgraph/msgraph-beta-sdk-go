package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantDetailedInformationable 
type TenantDetailedInformationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCity()(*string)
    GetCountryCode()(*string)
    GetCountryName()(*string)
    GetDefaultDomainName()(*string)
    GetDisplayName()(*string)
    GetIndustryName()(*string)
    GetRegion()(*string)
    GetSegmentName()(*string)
    GetTenantId()(*string)
    GetVerticalName()(*string)
    SetCity(value *string)()
    SetCountryCode(value *string)()
    SetCountryName(value *string)()
    SetDefaultDomainName(value *string)()
    SetDisplayName(value *string)()
    SetIndustryName(value *string)()
    SetRegion(value *string)()
    SetSegmentName(value *string)()
    SetTenantId(value *string)()
    SetVerticalName(value *string)()
}
