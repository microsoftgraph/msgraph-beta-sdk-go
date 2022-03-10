package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDisplayConfigurationable 
type TeamworkDisplayConfigurationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConfiguredDisplays()([]TeamworkConfiguredPeripheralable)
    GetDisplayCount()(*int32)
    GetInBuiltDisplayScreenConfiguration()(TeamworkDisplayScreenConfigurationable)
    GetIsContentDuplicationAllowed()(*bool)
    GetIsDualDisplayModeEnabled()(*bool)
    SetConfiguredDisplays(value []TeamworkConfiguredPeripheralable)()
    SetDisplayCount(value *int32)()
    SetInBuiltDisplayScreenConfiguration(value TeamworkDisplayScreenConfigurationable)()
    SetIsContentDuplicationAllowed(value *bool)()
    SetIsDualDisplayModeEnabled(value *bool)()
}
