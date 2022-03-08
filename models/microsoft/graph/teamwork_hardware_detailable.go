package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkHardwareDetailable 
type TeamworkHardwareDetailable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetMacAddresses()([]string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetSerialNumber()(*string)
    GetUniqueId()(*string)
    SetMacAddresses(value []string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetSerialNumber(value *string)()
    SetUniqueId(value *string)()
}
