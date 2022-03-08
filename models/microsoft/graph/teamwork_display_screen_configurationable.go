package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDisplayScreenConfigurationable 
type TeamworkDisplayScreenConfigurationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetBacklightBrightness()(*int32)
    GetBacklightTimeout()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetIsHighContrastEnabled()(*bool)
    GetIsScreensaverEnabled()(*bool)
    GetScreensaverTimeout()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    SetBacklightBrightness(value *int32)()
    SetBacklightTimeout(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetIsHighContrastEnabled(value *bool)()
    SetIsScreensaverEnabled(value *bool)()
    SetScreensaverTimeout(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
}
