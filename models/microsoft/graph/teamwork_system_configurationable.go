package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkSystemConfigurationable 
type TeamworkSystemConfigurationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDateTimeConfiguration()(TeamworkDateTimeConfigurationable)
    GetDefaultPassword()(*string)
    GetDeviceLockTimeout()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetIsDeviceLockEnabled()(*bool)
    GetIsLoggingEnabled()(*bool)
    GetIsPowerSavingEnabled()(*bool)
    GetIsScreenCaptureEnabled()(*bool)
    GetIsSilentModeEnabled()(*bool)
    GetLanguage()(*string)
    GetLockPin()(*string)
    GetLoggingLevel()(*string)
    GetNetworkConfiguration()(TeamworkNetworkConfigurationable)
    SetDateTimeConfiguration(value TeamworkDateTimeConfigurationable)()
    SetDefaultPassword(value *string)()
    SetDeviceLockTimeout(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetIsDeviceLockEnabled(value *bool)()
    SetIsLoggingEnabled(value *bool)()
    SetIsPowerSavingEnabled(value *bool)()
    SetIsScreenCaptureEnabled(value *bool)()
    SetIsSilentModeEnabled(value *bool)()
    SetLanguage(value *string)()
    SetLockPin(value *string)()
    SetLoggingLevel(value *string)()
    SetNetworkConfiguration(value TeamworkNetworkConfigurationable)()
}
