package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CompanyDetailable 
type CompanyDetailable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddress()(PhysicalAddressable)
    GetDepartment()(*string)
    GetDisplayName()(*string)
    GetOfficeLocation()(*string)
    GetPronunciation()(*string)
    GetWebUrl()(*string)
    SetAddress(value PhysicalAddressable)()
    SetDepartment(value *string)()
    SetDisplayName(value *string)()
    SetOfficeLocation(value *string)()
    SetPronunciation(value *string)()
    SetWebUrl(value *string)()
}
