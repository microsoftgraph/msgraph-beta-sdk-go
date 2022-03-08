package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PhysicalAddressable 
type PhysicalAddressable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCity()(*string)
    GetCountryOrRegion()(*string)
    GetPostalCode()(*string)
    GetPostOfficeBox()(*string)
    GetState()(*string)
    GetStreet()(*string)
    GetType()(*PhysicalAddressType)
    SetCity(value *string)()
    SetCountryOrRegion(value *string)()
    SetPostalCode(value *string)()
    SetPostOfficeBox(value *string)()
    SetState(value *string)()
    SetStreet(value *string)()
    SetType(value *PhysicalAddressType)()
}
