package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutOfBoxExperienceSettingsable 
type OutOfBoxExperienceSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceUsageType()(*WindowsDeviceUsageType)
    GetHideEscapeLink()(*bool)
    GetHideEULA()(*bool)
    GetHidePrivacySettings()(*bool)
    GetSkipKeyboardSelectionPage()(*bool)
    GetUserType()(*WindowsUserType)
    SetDeviceUsageType(value *WindowsDeviceUsageType)()
    SetHideEscapeLink(value *bool)()
    SetHideEULA(value *bool)()
    SetHidePrivacySettings(value *bool)()
    SetSkipKeyboardSelectionPage(value *bool)()
    SetUserType(value *WindowsUserType)()
}
