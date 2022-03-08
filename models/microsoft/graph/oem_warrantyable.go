package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OemWarrantyable 
type OemWarrantyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdditionalWarranties()([]WarrantyOfferable)
    GetBaseWarranties()([]WarrantyOfferable)
    GetDeviceConfigurationUrl()(*string)
    GetDeviceWarrantyUrl()(*string)
    SetAdditionalWarranties(value []WarrantyOfferable)()
    SetBaseWarranties(value []WarrantyOfferable)()
    SetDeviceConfigurationUrl(value *string)()
    SetDeviceWarrantyUrl(value *string)()
}
