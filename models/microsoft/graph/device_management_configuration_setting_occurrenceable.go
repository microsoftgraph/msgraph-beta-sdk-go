package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationSettingOccurrenceable 
type DeviceManagementConfigurationSettingOccurrenceable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetMaxDeviceOccurrence()(*int32)
    GetMinDeviceOccurrence()(*int32)
    SetMaxDeviceOccurrence(value *int32)()
    SetMinDeviceOccurrence(value *int32)()
}
