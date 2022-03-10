package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppIntentAndStateDetailable 
type MobileAppIntentAndStateDetailable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationId()(*string)
    GetDisplayName()(*string)
    GetDisplayVersion()(*string)
    GetInstallState()(*ResultantAppState)
    GetMobileAppIntent()(*MobileAppIntent)
    GetSupportedDeviceTypes()([]MobileAppSupportedDeviceTypeable)
    SetApplicationId(value *string)()
    SetDisplayName(value *string)()
    SetDisplayVersion(value *string)()
    SetInstallState(value *ResultantAppState)()
    SetMobileAppIntent(value *MobileAppIntent)()
    SetSupportedDeviceTypes(value []MobileAppSupportedDeviceTypeable)()
}
