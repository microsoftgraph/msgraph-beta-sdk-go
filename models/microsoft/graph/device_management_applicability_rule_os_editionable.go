package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementApplicabilityRuleOsEditionable 
type DeviceManagementApplicabilityRuleOsEditionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetName()(*string)
    GetOsEditionTypes()([]Windows10EditionType)
    GetRuleType()(*DeviceManagementApplicabilityRuleType)
    SetName(value *string)()
    SetOsEditionTypes(value []Windows10EditionType)()
    SetRuleType(value *DeviceManagementApplicabilityRuleType)()
}
