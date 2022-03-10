package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementSettingsable 
type DeviceManagementSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAndroidDeviceAdministratorEnrollmentEnabled()(*bool)
    GetDerivedCredentialProvider()(*DerivedCredentialProviderType)
    GetDerivedCredentialUrl()(*string)
    GetDeviceComplianceCheckinThresholdDays()(*int32)
    GetDeviceInactivityBeforeRetirementInDay()(*int32)
    GetEnableAutopilotDiagnostics()(*bool)
    GetEnableDeviceGroupMembershipReport()(*bool)
    GetEnableEnhancedTroubleshootingExperience()(*bool)
    GetEnableLogCollection()(*bool)
    GetEnhancedJailBreak()(*bool)
    GetIgnoreDevicesForUnsupportedSettingsEnabled()(*bool)
    GetIsScheduledActionEnabled()(*bool)
    GetSecureByDefault()(*bool)
    SetAndroidDeviceAdministratorEnrollmentEnabled(value *bool)()
    SetDerivedCredentialProvider(value *DerivedCredentialProviderType)()
    SetDerivedCredentialUrl(value *string)()
    SetDeviceComplianceCheckinThresholdDays(value *int32)()
    SetDeviceInactivityBeforeRetirementInDay(value *int32)()
    SetEnableAutopilotDiagnostics(value *bool)()
    SetEnableDeviceGroupMembershipReport(value *bool)()
    SetEnableEnhancedTroubleshootingExperience(value *bool)()
    SetEnableLogCollection(value *bool)()
    SetEnhancedJailBreak(value *bool)()
    SetIgnoreDevicesForUnsupportedSettingsEnabled(value *bool)()
    SetIsScheduledActionEnabled(value *bool)()
    SetSecureByDefault(value *bool)()
}
