package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationPolicyTemplateReferenceable 
type DeviceManagementConfigurationPolicyTemplateReferenceable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetTemplateDisplayName()(*string)
    GetTemplateDisplayVersion()(*string)
    GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily)
    GetTemplateId()(*string)
    SetTemplateDisplayName(value *string)()
    SetTemplateDisplayVersion(value *string)()
    SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)()
    SetTemplateId(value *string)()
}
