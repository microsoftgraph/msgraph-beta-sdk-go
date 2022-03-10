package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IosManagedAppProtectionable 
type IosManagedAppProtectionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    TargetedManagedAppProtectionable
    GetAllowedIosDeviceModels()(*string)
    GetAppActionIfIosDeviceModelNotAllowed()(*ManagedAppRemediationAction)
    GetAppDataEncryptionType()(*ManagedAppDataEncryptionType)
    GetApps()([]ManagedMobileAppable)
    GetCustomBrowserProtocol()(*string)
    GetCustomDialerAppProtocol()(*string)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetDisableProtectionOfManagedOutboundOpenInData()(*bool)
    GetExemptedAppProtocols()([]KeyValuePairable)
    GetExemptedUniversalLinks()([]string)
    GetFaceIdBlocked()(*bool)
    GetFilterOpenInToOnlyManagedApps()(*bool)
    GetManagedUniversalLinks()([]string)
    GetMinimumRequiredSdkVersion()(*string)
    GetMinimumWipeSdkVersion()(*string)
    GetProtectInboundDataFromUnknownSources()(*bool)
    GetThirdPartyKeyboardsBlocked()(*bool)
    SetAllowedIosDeviceModels(value *string)()
    SetAppActionIfIosDeviceModelNotAllowed(value *ManagedAppRemediationAction)()
    SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)()
    SetApps(value []ManagedMobileAppable)()
    SetCustomBrowserProtocol(value *string)()
    SetCustomDialerAppProtocol(value *string)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetDisableProtectionOfManagedOutboundOpenInData(value *bool)()
    SetExemptedAppProtocols(value []KeyValuePairable)()
    SetExemptedUniversalLinks(value []string)()
    SetFaceIdBlocked(value *bool)()
    SetFilterOpenInToOnlyManagedApps(value *bool)()
    SetManagedUniversalLinks(value []string)()
    SetMinimumRequiredSdkVersion(value *string)()
    SetMinimumWipeSdkVersion(value *string)()
    SetProtectInboundDataFromUnknownSources(value *bool)()
    SetThirdPartyKeyboardsBlocked(value *bool)()
}
