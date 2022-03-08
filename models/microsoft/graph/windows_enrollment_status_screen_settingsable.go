package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsEnrollmentStatusScreenSettingsable 
type WindowsEnrollmentStatusScreenSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowDeviceUseBeforeProfileAndAppInstallComplete()(*bool)
    GetAllowDeviceUseOnInstallFailure()(*bool)
    GetAllowLogCollectionOnInstallFailure()(*bool)
    GetBlockDeviceSetupRetryByUser()(*bool)
    GetCustomErrorMessage()(*string)
    GetHideInstallationProgress()(*bool)
    GetInstallProgressTimeoutInMinutes()(*int32)
    SetAllowDeviceUseBeforeProfileAndAppInstallComplete(value *bool)()
    SetAllowDeviceUseOnInstallFailure(value *bool)()
    SetAllowLogCollectionOnInstallFailure(value *bool)()
    SetBlockDeviceSetupRetryByUser(value *bool)()
    SetCustomErrorMessage(value *string)()
    SetHideInstallationProgress(value *bool)()
    SetInstallProgressTimeoutInMinutes(value *int32)()
}
