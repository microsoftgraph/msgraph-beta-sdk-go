package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DefaultManagedAppProtection struct {
    ManagedAppProtection
    // Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work. (Android only)
    allowedAndroidDeviceManufacturers *string;
    // List of device models allowed, as a string, for the managed app to work. (Android Only)
    allowedAndroidDeviceModels []string;
    // Semicolon seperated list of device models allowed, as a string, for the managed app to work. (iOS Only)
    allowedIosDeviceModels *string;
    // Defines a managed app behavior, either block or wipe, if the specified device manufacturer is not allowed. (Android only). Possible values are: block, wipe, warn.
    appActionIfAndroidDeviceManufacturerNotAllowed *ManagedAppRemediationAction;
    // Defines a managed app behavior, either block or wipe, if the specified device model is not allowed. (Android Only). Possible values are: block, wipe, warn.
    appActionIfAndroidDeviceModelNotAllowed *ManagedAppRemediationAction;
    // Defines a managed app behavior, either warn or block, if the specified Android App Verification requirement fails. Possible values are: block, wipe, warn.
    appActionIfAndroidSafetyNetAppsVerificationFailed *ManagedAppRemediationAction;
    // Defines a managed app behavior, either warn or block, if the specified Android SafetyNet Attestation requirement fails. Possible values are: block, wipe, warn.
    appActionIfAndroidSafetyNetDeviceAttestationFailed *ManagedAppRemediationAction;
    // Defines a managed app behavior, either warn, block or wipe, if the screen lock is required on device but is not set. (android only). Possible values are: block, wipe, warn.
    appActionIfDeviceLockNotSet *ManagedAppRemediationAction;
    // Defines a managed app behavior, either block or wipe, if the specified device model is not allowed. (iOS Only). Possible values are: block, wipe, warn.
    appActionIfIosDeviceModelNotAllowed *ManagedAppRemediationAction;
    // Type of encryption which should be used for data in a managed app. (iOS Only). Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked.
    appDataEncryptionType *ManagedAppDataEncryptionType;
    // List of apps to which the policy is deployed.
    apps []ManagedMobileApp;
    // Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True. (Android Only)
    biometricAuthenticationBlocked *bool;
    // Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
    blockAfterCompanyPortalUpdateDeferralInDays *int32;
    // Whether the app should connect to the configured VPN on launch (Android only).
    connectToVpnOnLaunch *bool;
    // Friendly name of the preferred custom browser to open weblink on Android. (Android only)
    customBrowserDisplayName *string;
    // Unique identifier of a custom browser to open weblink on Android. (Android only)
    customBrowserPackageId *string;
    // A custom browser protocol to open weblink on iOS. (iOS only)
    customBrowserProtocol *string;
    // Friendly name of a custom dialer app to click-to-open a phone number on Android.
    customDialerAppDisplayName *string;
    // PackageId of a custom dialer app to click-to-open a phone number on Android.
    customDialerAppPackageId *string;
    // Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
    customDialerAppProtocol *string;
    // A set of string key and string value pairs to be sent to the affected users, unalterned by this service
    customSettings []KeyValuePair;
    // Count of apps to which the current policy is deployed.
    deployedAppCount *int32;
    // Navigation property to deployment summary of the configuration.
    deploymentSummary *ManagedAppPolicyDeploymentSummary;
    // Defines if any kind of lock must be required on device. (android only)
    deviceLockRequired *bool;
    // When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only)
    disableAppEncryptionIfDeviceEncryptionIsEnabled *bool;
    // Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps. (iOS Only)
    disableProtectionOfManagedOutboundOpenInData *bool;
    // Indicates whether managed-app data should be encrypted. (Android only)
    encryptAppData *bool;
    // Android App packages in this list will be exempt from the policy and will be able to receive data from managed apps. (Android only)
    exemptedAppPackages []KeyValuePair;
    // iOS Apps in this list will be exempt from the policy and will be able to receive data from managed apps. (iOS Only)
    exemptedAppProtocols []KeyValuePair;
    // Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only)
    faceIdBlocked *bool;
    // Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False. (iOS Only)
    filterOpenInToOnlyManagedApps *bool;
    // Minimum version of the Company portal that must be installed on the device or app access will be blocked
    minimumRequiredCompanyPortalVersion *string;
    // Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only)
    minimumRequiredPatchVersion *string;
    // Versions less than the specified version will block the managed app from accessing company data. (iOS Only)
    minimumRequiredSdkVersion *string;
    // Minimum version of the Company portal that must be installed on the device or the user will receive a warning
    minimumWarningCompanyPortalVersion *string;
    // Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only)
    minimumWarningPatchVersion *string;
    // Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
    minimumWipeCompanyPortalVersion *string;
    // Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data. (Android only)
    minimumWipePatchVersion *string;
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumWipeSdkVersion *string;
    // Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps. (iOS Only)
    protectInboundDataFromUnknownSources *bool;
    // Defines the Android SafetyNet Apps Verification requirement for a managed app to work. Possible values are: none, enabled.
    requiredAndroidSafetyNetAppsVerificationType *AndroidManagedAppSafetyNetAppsVerificationType;
    // Defines the Android SafetyNet Device Attestation requirement for a managed app to work. Possible values are: none, basicIntegrity, basicIntegrityAndDeviceCertification.
    requiredAndroidSafetyNetDeviceAttestationType *AndroidManagedAppSafetyNetDeviceAttestationType;
    // Defines the Android SafetyNet evaluation type requirement for a managed app to work. (Android Only). Possible values are: basic, hardwareBacked.
    requiredAndroidSafetyNetEvaluationType *AndroidManagedAppSafetyNetEvaluationType;
    // Indicates whether screen capture is blocked. (Android only)
    screenCaptureBlocked *bool;
    // Defines if third party keyboards are allowed while accessing a managed app. (iOS Only)
    thirdPartyKeyboardsBlocked *bool;
    // Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
    warnAfterCompanyPortalUpdateDeferralInDays *int32;
    // Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
    wipeAfterCompanyPortalUpdateDeferralInDays *int32;
}
// Instantiates a new defaultManagedAppProtection and sets the default values.
func NewDefaultManagedAppProtection()(*DefaultManagedAppProtection) {
    m := &DefaultManagedAppProtection{
        ManagedAppProtection: *NewManagedAppProtection(),
    }
    return m
}
// Gets the allowedAndroidDeviceManufacturers property value. Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work. (Android only)
func (m *DefaultManagedAppProtection) GetAllowedAndroidDeviceManufacturers()(*string) {
    if m == nil {
        return nil
    } else {
        return m.allowedAndroidDeviceManufacturers
    }
}
// Gets the allowedAndroidDeviceModels property value. List of device models allowed, as a string, for the managed app to work. (Android Only)
func (m *DefaultManagedAppProtection) GetAllowedAndroidDeviceModels()([]string) {
    if m == nil {
        return nil
    } else {
        return m.allowedAndroidDeviceModels
    }
}
// Gets the allowedIosDeviceModels property value. Semicolon seperated list of device models allowed, as a string, for the managed app to work. (iOS Only)
func (m *DefaultManagedAppProtection) GetAllowedIosDeviceModels()(*string) {
    if m == nil {
        return nil
    } else {
        return m.allowedIosDeviceModels
    }
}
// Gets the appActionIfAndroidDeviceManufacturerNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device manufacturer is not allowed. (Android only). Possible values are: block, wipe, warn.
func (m *DefaultManagedAppProtection) GetAppActionIfAndroidDeviceManufacturerNotAllowed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfAndroidDeviceManufacturerNotAllowed
    }
}
// Gets the appActionIfAndroidDeviceModelNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device model is not allowed. (Android Only). Possible values are: block, wipe, warn.
func (m *DefaultManagedAppProtection) GetAppActionIfAndroidDeviceModelNotAllowed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfAndroidDeviceModelNotAllowed
    }
}
// Gets the appActionIfAndroidSafetyNetAppsVerificationFailed property value. Defines a managed app behavior, either warn or block, if the specified Android App Verification requirement fails. Possible values are: block, wipe, warn.
func (m *DefaultManagedAppProtection) GetAppActionIfAndroidSafetyNetAppsVerificationFailed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfAndroidSafetyNetAppsVerificationFailed
    }
}
// Gets the appActionIfAndroidSafetyNetDeviceAttestationFailed property value. Defines a managed app behavior, either warn or block, if the specified Android SafetyNet Attestation requirement fails. Possible values are: block, wipe, warn.
func (m *DefaultManagedAppProtection) GetAppActionIfAndroidSafetyNetDeviceAttestationFailed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfAndroidSafetyNetDeviceAttestationFailed
    }
}
// Gets the appActionIfDeviceLockNotSet property value. Defines a managed app behavior, either warn, block or wipe, if the screen lock is required on device but is not set. (android only). Possible values are: block, wipe, warn.
func (m *DefaultManagedAppProtection) GetAppActionIfDeviceLockNotSet()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfDeviceLockNotSet
    }
}
// Gets the appActionIfIosDeviceModelNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device model is not allowed. (iOS Only). Possible values are: block, wipe, warn.
func (m *DefaultManagedAppProtection) GetAppActionIfIosDeviceModelNotAllowed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfIosDeviceModelNotAllowed
    }
}
// Gets the appDataEncryptionType property value. Type of encryption which should be used for data in a managed app. (iOS Only). Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked.
func (m *DefaultManagedAppProtection) GetAppDataEncryptionType()(*ManagedAppDataEncryptionType) {
    if m == nil {
        return nil
    } else {
        return m.appDataEncryptionType
    }
}
// Gets the apps property value. List of apps to which the policy is deployed.
func (m *DefaultManagedAppProtection) GetApps()([]ManagedMobileApp) {
    if m == nil {
        return nil
    } else {
        return m.apps
    }
}
// Gets the biometricAuthenticationBlocked property value. Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True. (Android Only)
func (m *DefaultManagedAppProtection) GetBiometricAuthenticationBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.biometricAuthenticationBlocked
    }
}
// Gets the blockAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
func (m *DefaultManagedAppProtection) GetBlockAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.blockAfterCompanyPortalUpdateDeferralInDays
    }
}
// Gets the connectToVpnOnLaunch property value. Whether the app should connect to the configured VPN on launch (Android only).
func (m *DefaultManagedAppProtection) GetConnectToVpnOnLaunch()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.connectToVpnOnLaunch
    }
}
// Gets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android. (Android only)
func (m *DefaultManagedAppProtection) GetCustomBrowserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customBrowserDisplayName
    }
}
// Gets the customBrowserPackageId property value. Unique identifier of a custom browser to open weblink on Android. (Android only)
func (m *DefaultManagedAppProtection) GetCustomBrowserPackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customBrowserPackageId
    }
}
// Gets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS. (iOS only)
func (m *DefaultManagedAppProtection) GetCustomBrowserProtocol()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customBrowserProtocol
    }
}
// Gets the customDialerAppDisplayName property value. Friendly name of a custom dialer app to click-to-open a phone number on Android.
func (m *DefaultManagedAppProtection) GetCustomDialerAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customDialerAppDisplayName
    }
}
// Gets the customDialerAppPackageId property value. PackageId of a custom dialer app to click-to-open a phone number on Android.
func (m *DefaultManagedAppProtection) GetCustomDialerAppPackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customDialerAppPackageId
    }
}
// Gets the customDialerAppProtocol property value. Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
func (m *DefaultManagedAppProtection) GetCustomDialerAppProtocol()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customDialerAppProtocol
    }
}
// Gets the customSettings property value. A set of string key and string value pairs to be sent to the affected users, unalterned by this service
func (m *DefaultManagedAppProtection) GetCustomSettings()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.customSettings
    }
}
// Gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *DefaultManagedAppProtection) GetDeployedAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deployedAppCount
    }
}
// Gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *DefaultManagedAppProtection) GetDeploymentSummary()(*ManagedAppPolicyDeploymentSummary) {
    if m == nil {
        return nil
    } else {
        return m.deploymentSummary
    }
}
// Gets the deviceLockRequired property value. Defines if any kind of lock must be required on device. (android only)
func (m *DefaultManagedAppProtection) GetDeviceLockRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceLockRequired
    }
}
// Gets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only)
func (m *DefaultManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableAppEncryptionIfDeviceEncryptionIsEnabled
    }
}
// Gets the disableProtectionOfManagedOutboundOpenInData property value. Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps. (iOS Only)
func (m *DefaultManagedAppProtection) GetDisableProtectionOfManagedOutboundOpenInData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableProtectionOfManagedOutboundOpenInData
    }
}
// Gets the encryptAppData property value. Indicates whether managed-app data should be encrypted. (Android only)
func (m *DefaultManagedAppProtection) GetEncryptAppData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.encryptAppData
    }
}
// Gets the exemptedAppPackages property value. Android App packages in this list will be exempt from the policy and will be able to receive data from managed apps. (Android only)
func (m *DefaultManagedAppProtection) GetExemptedAppPackages()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.exemptedAppPackages
    }
}
// Gets the exemptedAppProtocols property value. iOS Apps in this list will be exempt from the policy and will be able to receive data from managed apps. (iOS Only)
func (m *DefaultManagedAppProtection) GetExemptedAppProtocols()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.exemptedAppProtocols
    }
}
// Gets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only)
func (m *DefaultManagedAppProtection) GetFaceIdBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.faceIdBlocked
    }
}
// Gets the filterOpenInToOnlyManagedApps property value. Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False. (iOS Only)
func (m *DefaultManagedAppProtection) GetFilterOpenInToOnlyManagedApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.filterOpenInToOnlyManagedApps
    }
}
// Gets the minimumRequiredCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or app access will be blocked
func (m *DefaultManagedAppProtection) GetMinimumRequiredCompanyPortalVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredCompanyPortalVersion
    }
}
// Gets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only)
func (m *DefaultManagedAppProtection) GetMinimumRequiredPatchVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredPatchVersion
    }
}
// Gets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data. (iOS Only)
func (m *DefaultManagedAppProtection) GetMinimumRequiredSdkVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredSdkVersion
    }
}
// Gets the minimumWarningCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the user will receive a warning
func (m *DefaultManagedAppProtection) GetMinimumWarningCompanyPortalVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningCompanyPortalVersion
    }
}
// Gets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only)
func (m *DefaultManagedAppProtection) GetMinimumWarningPatchVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningPatchVersion
    }
}
// Gets the minimumWipeCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
func (m *DefaultManagedAppProtection) GetMinimumWipeCompanyPortalVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWipeCompanyPortalVersion
    }
}
// Gets the minimumWipePatchVersion property value. Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data. (Android only)
func (m *DefaultManagedAppProtection) GetMinimumWipePatchVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWipePatchVersion
    }
}
// Gets the minimumWipeSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *DefaultManagedAppProtection) GetMinimumWipeSdkVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWipeSdkVersion
    }
}
// Gets the protectInboundDataFromUnknownSources property value. Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps. (iOS Only)
func (m *DefaultManagedAppProtection) GetProtectInboundDataFromUnknownSources()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.protectInboundDataFromUnknownSources
    }
}
// Gets the requiredAndroidSafetyNetAppsVerificationType property value. Defines the Android SafetyNet Apps Verification requirement for a managed app to work. Possible values are: none, enabled.
func (m *DefaultManagedAppProtection) GetRequiredAndroidSafetyNetAppsVerificationType()(*AndroidManagedAppSafetyNetAppsVerificationType) {
    if m == nil {
        return nil
    } else {
        return m.requiredAndroidSafetyNetAppsVerificationType
    }
}
// Gets the requiredAndroidSafetyNetDeviceAttestationType property value. Defines the Android SafetyNet Device Attestation requirement for a managed app to work. Possible values are: none, basicIntegrity, basicIntegrityAndDeviceCertification.
func (m *DefaultManagedAppProtection) GetRequiredAndroidSafetyNetDeviceAttestationType()(*AndroidManagedAppSafetyNetDeviceAttestationType) {
    if m == nil {
        return nil
    } else {
        return m.requiredAndroidSafetyNetDeviceAttestationType
    }
}
// Gets the requiredAndroidSafetyNetEvaluationType property value. Defines the Android SafetyNet evaluation type requirement for a managed app to work. (Android Only). Possible values are: basic, hardwareBacked.
func (m *DefaultManagedAppProtection) GetRequiredAndroidSafetyNetEvaluationType()(*AndroidManagedAppSafetyNetEvaluationType) {
    if m == nil {
        return nil
    } else {
        return m.requiredAndroidSafetyNetEvaluationType
    }
}
// Gets the screenCaptureBlocked property value. Indicates whether screen capture is blocked. (Android only)
func (m *DefaultManagedAppProtection) GetScreenCaptureBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.screenCaptureBlocked
    }
}
// Gets the thirdPartyKeyboardsBlocked property value. Defines if third party keyboards are allowed while accessing a managed app. (iOS Only)
func (m *DefaultManagedAppProtection) GetThirdPartyKeyboardsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.thirdPartyKeyboardsBlocked
    }
}
// Gets the warnAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
func (m *DefaultManagedAppProtection) GetWarnAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.warnAfterCompanyPortalUpdateDeferralInDays
    }
}
// Gets the wipeAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
func (m *DefaultManagedAppProtection) GetWipeAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.wipeAfterCompanyPortalUpdateDeferralInDays
    }
}
// The deserialization information for the current model
func (m *DefaultManagedAppProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagedAppProtection.GetFieldDeserializers()
    res["allowedAndroidDeviceManufacturers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAllowedAndroidDeviceManufacturers(val)
        return nil
    }
    res["allowedAndroidDeviceModels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAllowedAndroidDeviceModels(res)
        return nil
    }
    res["allowedIosDeviceModels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAllowedIosDeviceModels(val)
        return nil
    }
    res["appActionIfAndroidDeviceManufacturerNotAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppRemediationAction)
        m.SetAppActionIfAndroidDeviceManufacturerNotAllowed(&cast)
        return nil
    }
    res["appActionIfAndroidDeviceModelNotAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppRemediationAction)
        m.SetAppActionIfAndroidDeviceModelNotAllowed(&cast)
        return nil
    }
    res["appActionIfAndroidSafetyNetAppsVerificationFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppRemediationAction)
        m.SetAppActionIfAndroidSafetyNetAppsVerificationFailed(&cast)
        return nil
    }
    res["appActionIfAndroidSafetyNetDeviceAttestationFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppRemediationAction)
        m.SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(&cast)
        return nil
    }
    res["appActionIfDeviceLockNotSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppRemediationAction)
        m.SetAppActionIfDeviceLockNotSet(&cast)
        return nil
    }
    res["appActionIfIosDeviceModelNotAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppRemediationAction)
        m.SetAppActionIfIosDeviceModelNotAllowed(&cast)
        return nil
    }
    res["appDataEncryptionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataEncryptionType)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppDataEncryptionType)
        m.SetAppDataEncryptionType(&cast)
        return nil
    }
    res["apps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedMobileApp() })
        if err != nil {
            return err
        }
        res := make([]ManagedMobileApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedMobileApp))
        }
        m.SetApps(res)
        return nil
    }
    res["biometricAuthenticationBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetBiometricAuthenticationBlocked(val)
        return nil
    }
    res["blockAfterCompanyPortalUpdateDeferralInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetBlockAfterCompanyPortalUpdateDeferralInDays(val)
        return nil
    }
    res["connectToVpnOnLaunch"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetConnectToVpnOnLaunch(val)
        return nil
    }
    res["customBrowserDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomBrowserDisplayName(val)
        return nil
    }
    res["customBrowserPackageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomBrowserPackageId(val)
        return nil
    }
    res["customBrowserProtocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomBrowserProtocol(val)
        return nil
    }
    res["customDialerAppDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomDialerAppDisplayName(val)
        return nil
    }
    res["customDialerAppPackageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomDialerAppPackageId(val)
        return nil
    }
    res["customDialerAppProtocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomDialerAppProtocol(val)
        return nil
    }
    res["customSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetCustomSettings(res)
        return nil
    }
    res["deployedAppCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeployedAppCount(val)
        return nil
    }
    res["deploymentSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicyDeploymentSummary() })
        if err != nil {
            return err
        }
        m.SetDeploymentSummary(val.(*ManagedAppPolicyDeploymentSummary))
        return nil
    }
    res["deviceLockRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDeviceLockRequired(val)
        return nil
    }
    res["disableAppEncryptionIfDeviceEncryptionIsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(val)
        return nil
    }
    res["disableProtectionOfManagedOutboundOpenInData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDisableProtectionOfManagedOutboundOpenInData(val)
        return nil
    }
    res["encryptAppData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEncryptAppData(val)
        return nil
    }
    res["exemptedAppPackages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetExemptedAppPackages(res)
        return nil
    }
    res["exemptedAppProtocols"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetExemptedAppProtocols(res)
        return nil
    }
    res["faceIdBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFaceIdBlocked(val)
        return nil
    }
    res["filterOpenInToOnlyManagedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFilterOpenInToOnlyManagedApps(val)
        return nil
    }
    res["minimumRequiredCompanyPortalVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumRequiredCompanyPortalVersion(val)
        return nil
    }
    res["minimumRequiredPatchVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumRequiredPatchVersion(val)
        return nil
    }
    res["minimumRequiredSdkVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumRequiredSdkVersion(val)
        return nil
    }
    res["minimumWarningCompanyPortalVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumWarningCompanyPortalVersion(val)
        return nil
    }
    res["minimumWarningPatchVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumWarningPatchVersion(val)
        return nil
    }
    res["minimumWipeCompanyPortalVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumWipeCompanyPortalVersion(val)
        return nil
    }
    res["minimumWipePatchVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumWipePatchVersion(val)
        return nil
    }
    res["minimumWipeSdkVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumWipeSdkVersion(val)
        return nil
    }
    res["protectInboundDataFromUnknownSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetProtectInboundDataFromUnknownSources(val)
        return nil
    }
    res["requiredAndroidSafetyNetAppsVerificationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedAppSafetyNetAppsVerificationType)
        if err != nil {
            return err
        }
        cast := val.(AndroidManagedAppSafetyNetAppsVerificationType)
        m.SetRequiredAndroidSafetyNetAppsVerificationType(&cast)
        return nil
    }
    res["requiredAndroidSafetyNetDeviceAttestationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedAppSafetyNetDeviceAttestationType)
        if err != nil {
            return err
        }
        cast := val.(AndroidManagedAppSafetyNetDeviceAttestationType)
        m.SetRequiredAndroidSafetyNetDeviceAttestationType(&cast)
        return nil
    }
    res["requiredAndroidSafetyNetEvaluationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedAppSafetyNetEvaluationType)
        if err != nil {
            return err
        }
        cast := val.(AndroidManagedAppSafetyNetEvaluationType)
        m.SetRequiredAndroidSafetyNetEvaluationType(&cast)
        return nil
    }
    res["screenCaptureBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetScreenCaptureBlocked(val)
        return nil
    }
    res["thirdPartyKeyboardsBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetThirdPartyKeyboardsBlocked(val)
        return nil
    }
    res["warnAfterCompanyPortalUpdateDeferralInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWarnAfterCompanyPortalUpdateDeferralInDays(val)
        return nil
    }
    res["wipeAfterCompanyPortalUpdateDeferralInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWipeAfterCompanyPortalUpdateDeferralInDays(val)
        return nil
    }
    return res
}
func (m *DefaultManagedAppProtection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DefaultManagedAppProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagedAppProtection.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("allowedAndroidDeviceManufacturers", m.GetAllowedAndroidDeviceManufacturers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("allowedAndroidDeviceModels", m.GetAllowedAndroidDeviceModels())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("allowedIosDeviceModels", m.GetAllowedIosDeviceModels())
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAndroidDeviceManufacturerNotAllowed() != nil {
        cast := m.GetAppActionIfAndroidDeviceManufacturerNotAllowed().String()
        err = writer.WriteStringValue("appActionIfAndroidDeviceManufacturerNotAllowed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAndroidDeviceModelNotAllowed() != nil {
        cast := m.GetAppActionIfAndroidDeviceModelNotAllowed().String()
        err = writer.WriteStringValue("appActionIfAndroidDeviceModelNotAllowed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAndroidSafetyNetAppsVerificationFailed() != nil {
        cast := m.GetAppActionIfAndroidSafetyNetAppsVerificationFailed().String()
        err = writer.WriteStringValue("appActionIfAndroidSafetyNetAppsVerificationFailed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAndroidSafetyNetDeviceAttestationFailed() != nil {
        cast := m.GetAppActionIfAndroidSafetyNetDeviceAttestationFailed().String()
        err = writer.WriteStringValue("appActionIfAndroidSafetyNetDeviceAttestationFailed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfDeviceLockNotSet() != nil {
        cast := m.GetAppActionIfDeviceLockNotSet().String()
        err = writer.WriteStringValue("appActionIfDeviceLockNotSet", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfIosDeviceModelNotAllowed() != nil {
        cast := m.GetAppActionIfIosDeviceModelNotAllowed().String()
        err = writer.WriteStringValue("appActionIfIosDeviceModelNotAllowed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppDataEncryptionType() != nil {
        cast := m.GetAppDataEncryptionType().String()
        err = writer.WriteStringValue("appDataEncryptionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("apps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("biometricAuthenticationBlocked", m.GetBiometricAuthenticationBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("blockAfterCompanyPortalUpdateDeferralInDays", m.GetBlockAfterCompanyPortalUpdateDeferralInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("connectToVpnOnLaunch", m.GetConnectToVpnOnLaunch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customBrowserDisplayName", m.GetCustomBrowserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customBrowserPackageId", m.GetCustomBrowserPackageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customBrowserProtocol", m.GetCustomBrowserProtocol())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customDialerAppDisplayName", m.GetCustomDialerAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customDialerAppPackageId", m.GetCustomDialerAppPackageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customDialerAppProtocol", m.GetCustomDialerAppProtocol())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomSettings()))
        for i, v := range m.GetCustomSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deployedAppCount", m.GetDeployedAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentSummary", m.GetDeploymentSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceLockRequired", m.GetDeviceLockRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableAppEncryptionIfDeviceEncryptionIsEnabled", m.GetDisableAppEncryptionIfDeviceEncryptionIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableProtectionOfManagedOutboundOpenInData", m.GetDisableProtectionOfManagedOutboundOpenInData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("encryptAppData", m.GetEncryptAppData())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExemptedAppPackages()))
        for i, v := range m.GetExemptedAppPackages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("exemptedAppPackages", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExemptedAppProtocols()))
        for i, v := range m.GetExemptedAppProtocols() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("exemptedAppProtocols", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("faceIdBlocked", m.GetFaceIdBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("filterOpenInToOnlyManagedApps", m.GetFilterOpenInToOnlyManagedApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredCompanyPortalVersion", m.GetMinimumRequiredCompanyPortalVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredPatchVersion", m.GetMinimumRequiredPatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredSdkVersion", m.GetMinimumRequiredSdkVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWarningCompanyPortalVersion", m.GetMinimumWarningCompanyPortalVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWarningPatchVersion", m.GetMinimumWarningPatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWipeCompanyPortalVersion", m.GetMinimumWipeCompanyPortalVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWipePatchVersion", m.GetMinimumWipePatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWipeSdkVersion", m.GetMinimumWipeSdkVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("protectInboundDataFromUnknownSources", m.GetProtectInboundDataFromUnknownSources())
        if err != nil {
            return err
        }
    }
    if m.GetRequiredAndroidSafetyNetAppsVerificationType() != nil {
        cast := m.GetRequiredAndroidSafetyNetAppsVerificationType().String()
        err = writer.WriteStringValue("requiredAndroidSafetyNetAppsVerificationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequiredAndroidSafetyNetDeviceAttestationType() != nil {
        cast := m.GetRequiredAndroidSafetyNetDeviceAttestationType().String()
        err = writer.WriteStringValue("requiredAndroidSafetyNetDeviceAttestationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequiredAndroidSafetyNetEvaluationType() != nil {
        cast := m.GetRequiredAndroidSafetyNetEvaluationType().String()
        err = writer.WriteStringValue("requiredAndroidSafetyNetEvaluationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("screenCaptureBlocked", m.GetScreenCaptureBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("thirdPartyKeyboardsBlocked", m.GetThirdPartyKeyboardsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("warnAfterCompanyPortalUpdateDeferralInDays", m.GetWarnAfterCompanyPortalUpdateDeferralInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("wipeAfterCompanyPortalUpdateDeferralInDays", m.GetWipeAfterCompanyPortalUpdateDeferralInDays())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the allowedAndroidDeviceManufacturers property value. Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work. (Android only)
// Parameters:
//  - value : Value to set for the allowedAndroidDeviceManufacturers property.
func (m *DefaultManagedAppProtection) SetAllowedAndroidDeviceManufacturers(value *string)() {
    m.allowedAndroidDeviceManufacturers = value
}
// Sets the allowedAndroidDeviceModels property value. List of device models allowed, as a string, for the managed app to work. (Android Only)
// Parameters:
//  - value : Value to set for the allowedAndroidDeviceModels property.
func (m *DefaultManagedAppProtection) SetAllowedAndroidDeviceModels(value []string)() {
    m.allowedAndroidDeviceModels = value
}
// Sets the allowedIosDeviceModels property value. Semicolon seperated list of device models allowed, as a string, for the managed app to work. (iOS Only)
// Parameters:
//  - value : Value to set for the allowedIosDeviceModels property.
func (m *DefaultManagedAppProtection) SetAllowedIosDeviceModels(value *string)() {
    m.allowedIosDeviceModels = value
}
// Sets the appActionIfAndroidDeviceManufacturerNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device manufacturer is not allowed. (Android only). Possible values are: block, wipe, warn.
// Parameters:
//  - value : Value to set for the appActionIfAndroidDeviceManufacturerNotAllowed property.
func (m *DefaultManagedAppProtection) SetAppActionIfAndroidDeviceManufacturerNotAllowed(value *ManagedAppRemediationAction)() {
    m.appActionIfAndroidDeviceManufacturerNotAllowed = value
}
// Sets the appActionIfAndroidDeviceModelNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device model is not allowed. (Android Only). Possible values are: block, wipe, warn.
// Parameters:
//  - value : Value to set for the appActionIfAndroidDeviceModelNotAllowed property.
func (m *DefaultManagedAppProtection) SetAppActionIfAndroidDeviceModelNotAllowed(value *ManagedAppRemediationAction)() {
    m.appActionIfAndroidDeviceModelNotAllowed = value
}
// Sets the appActionIfAndroidSafetyNetAppsVerificationFailed property value. Defines a managed app behavior, either warn or block, if the specified Android App Verification requirement fails. Possible values are: block, wipe, warn.
// Parameters:
//  - value : Value to set for the appActionIfAndroidSafetyNetAppsVerificationFailed property.
func (m *DefaultManagedAppProtection) SetAppActionIfAndroidSafetyNetAppsVerificationFailed(value *ManagedAppRemediationAction)() {
    m.appActionIfAndroidSafetyNetAppsVerificationFailed = value
}
// Sets the appActionIfAndroidSafetyNetDeviceAttestationFailed property value. Defines a managed app behavior, either warn or block, if the specified Android SafetyNet Attestation requirement fails. Possible values are: block, wipe, warn.
// Parameters:
//  - value : Value to set for the appActionIfAndroidSafetyNetDeviceAttestationFailed property.
func (m *DefaultManagedAppProtection) SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(value *ManagedAppRemediationAction)() {
    m.appActionIfAndroidSafetyNetDeviceAttestationFailed = value
}
// Sets the appActionIfDeviceLockNotSet property value. Defines a managed app behavior, either warn, block or wipe, if the screen lock is required on device but is not set. (android only). Possible values are: block, wipe, warn.
// Parameters:
//  - value : Value to set for the appActionIfDeviceLockNotSet property.
func (m *DefaultManagedAppProtection) SetAppActionIfDeviceLockNotSet(value *ManagedAppRemediationAction)() {
    m.appActionIfDeviceLockNotSet = value
}
// Sets the appActionIfIosDeviceModelNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device model is not allowed. (iOS Only). Possible values are: block, wipe, warn.
// Parameters:
//  - value : Value to set for the appActionIfIosDeviceModelNotAllowed property.
func (m *DefaultManagedAppProtection) SetAppActionIfIosDeviceModelNotAllowed(value *ManagedAppRemediationAction)() {
    m.appActionIfIosDeviceModelNotAllowed = value
}
// Sets the appDataEncryptionType property value. Type of encryption which should be used for data in a managed app. (iOS Only). Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked.
// Parameters:
//  - value : Value to set for the appDataEncryptionType property.
func (m *DefaultManagedAppProtection) SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)() {
    m.appDataEncryptionType = value
}
// Sets the apps property value. List of apps to which the policy is deployed.
// Parameters:
//  - value : Value to set for the apps property.
func (m *DefaultManagedAppProtection) SetApps(value []ManagedMobileApp)() {
    m.apps = value
}
// Sets the biometricAuthenticationBlocked property value. Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True. (Android Only)
// Parameters:
//  - value : Value to set for the biometricAuthenticationBlocked property.
func (m *DefaultManagedAppProtection) SetBiometricAuthenticationBlocked(value *bool)() {
    m.biometricAuthenticationBlocked = value
}
// Sets the blockAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
// Parameters:
//  - value : Value to set for the blockAfterCompanyPortalUpdateDeferralInDays property.
func (m *DefaultManagedAppProtection) SetBlockAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    m.blockAfterCompanyPortalUpdateDeferralInDays = value
}
// Sets the connectToVpnOnLaunch property value. Whether the app should connect to the configured VPN on launch (Android only).
// Parameters:
//  - value : Value to set for the connectToVpnOnLaunch property.
func (m *DefaultManagedAppProtection) SetConnectToVpnOnLaunch(value *bool)() {
    m.connectToVpnOnLaunch = value
}
// Sets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android. (Android only)
// Parameters:
//  - value : Value to set for the customBrowserDisplayName property.
func (m *DefaultManagedAppProtection) SetCustomBrowserDisplayName(value *string)() {
    m.customBrowserDisplayName = value
}
// Sets the customBrowserPackageId property value. Unique identifier of a custom browser to open weblink on Android. (Android only)
// Parameters:
//  - value : Value to set for the customBrowserPackageId property.
func (m *DefaultManagedAppProtection) SetCustomBrowserPackageId(value *string)() {
    m.customBrowserPackageId = value
}
// Sets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS. (iOS only)
// Parameters:
//  - value : Value to set for the customBrowserProtocol property.
func (m *DefaultManagedAppProtection) SetCustomBrowserProtocol(value *string)() {
    m.customBrowserProtocol = value
}
// Sets the customDialerAppDisplayName property value. Friendly name of a custom dialer app to click-to-open a phone number on Android.
// Parameters:
//  - value : Value to set for the customDialerAppDisplayName property.
func (m *DefaultManagedAppProtection) SetCustomDialerAppDisplayName(value *string)() {
    m.customDialerAppDisplayName = value
}
// Sets the customDialerAppPackageId property value. PackageId of a custom dialer app to click-to-open a phone number on Android.
// Parameters:
//  - value : Value to set for the customDialerAppPackageId property.
func (m *DefaultManagedAppProtection) SetCustomDialerAppPackageId(value *string)() {
    m.customDialerAppPackageId = value
}
// Sets the customDialerAppProtocol property value. Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
// Parameters:
//  - value : Value to set for the customDialerAppProtocol property.
func (m *DefaultManagedAppProtection) SetCustomDialerAppProtocol(value *string)() {
    m.customDialerAppProtocol = value
}
// Sets the customSettings property value. A set of string key and string value pairs to be sent to the affected users, unalterned by this service
// Parameters:
//  - value : Value to set for the customSettings property.
func (m *DefaultManagedAppProtection) SetCustomSettings(value []KeyValuePair)() {
    m.customSettings = value
}
// Sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
// Parameters:
//  - value : Value to set for the deployedAppCount property.
func (m *DefaultManagedAppProtection) SetDeployedAppCount(value *int32)() {
    m.deployedAppCount = value
}
// Sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
// Parameters:
//  - value : Value to set for the deploymentSummary property.
func (m *DefaultManagedAppProtection) SetDeploymentSummary(value *ManagedAppPolicyDeploymentSummary)() {
    m.deploymentSummary = value
}
// Sets the deviceLockRequired property value. Defines if any kind of lock must be required on device. (android only)
// Parameters:
//  - value : Value to set for the deviceLockRequired property.
func (m *DefaultManagedAppProtection) SetDeviceLockRequired(value *bool)() {
    m.deviceLockRequired = value
}
// Sets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only)
// Parameters:
//  - value : Value to set for the disableAppEncryptionIfDeviceEncryptionIsEnabled property.
func (m *DefaultManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)() {
    m.disableAppEncryptionIfDeviceEncryptionIsEnabled = value
}
// Sets the disableProtectionOfManagedOutboundOpenInData property value. Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps. (iOS Only)
// Parameters:
//  - value : Value to set for the disableProtectionOfManagedOutboundOpenInData property.
func (m *DefaultManagedAppProtection) SetDisableProtectionOfManagedOutboundOpenInData(value *bool)() {
    m.disableProtectionOfManagedOutboundOpenInData = value
}
// Sets the encryptAppData property value. Indicates whether managed-app data should be encrypted. (Android only)
// Parameters:
//  - value : Value to set for the encryptAppData property.
func (m *DefaultManagedAppProtection) SetEncryptAppData(value *bool)() {
    m.encryptAppData = value
}
// Sets the exemptedAppPackages property value. Android App packages in this list will be exempt from the policy and will be able to receive data from managed apps. (Android only)
// Parameters:
//  - value : Value to set for the exemptedAppPackages property.
func (m *DefaultManagedAppProtection) SetExemptedAppPackages(value []KeyValuePair)() {
    m.exemptedAppPackages = value
}
// Sets the exemptedAppProtocols property value. iOS Apps in this list will be exempt from the policy and will be able to receive data from managed apps. (iOS Only)
// Parameters:
//  - value : Value to set for the exemptedAppProtocols property.
func (m *DefaultManagedAppProtection) SetExemptedAppProtocols(value []KeyValuePair)() {
    m.exemptedAppProtocols = value
}
// Sets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only)
// Parameters:
//  - value : Value to set for the faceIdBlocked property.
func (m *DefaultManagedAppProtection) SetFaceIdBlocked(value *bool)() {
    m.faceIdBlocked = value
}
// Sets the filterOpenInToOnlyManagedApps property value. Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False. (iOS Only)
// Parameters:
//  - value : Value to set for the filterOpenInToOnlyManagedApps property.
func (m *DefaultManagedAppProtection) SetFilterOpenInToOnlyManagedApps(value *bool)() {
    m.filterOpenInToOnlyManagedApps = value
}
// Sets the minimumRequiredCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or app access will be blocked
// Parameters:
//  - value : Value to set for the minimumRequiredCompanyPortalVersion property.
func (m *DefaultManagedAppProtection) SetMinimumRequiredCompanyPortalVersion(value *string)() {
    m.minimumRequiredCompanyPortalVersion = value
}
// Sets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only)
// Parameters:
//  - value : Value to set for the minimumRequiredPatchVersion property.
func (m *DefaultManagedAppProtection) SetMinimumRequiredPatchVersion(value *string)() {
    m.minimumRequiredPatchVersion = value
}
// Sets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data. (iOS Only)
// Parameters:
//  - value : Value to set for the minimumRequiredSdkVersion property.
func (m *DefaultManagedAppProtection) SetMinimumRequiredSdkVersion(value *string)() {
    m.minimumRequiredSdkVersion = value
}
// Sets the minimumWarningCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the user will receive a warning
// Parameters:
//  - value : Value to set for the minimumWarningCompanyPortalVersion property.
func (m *DefaultManagedAppProtection) SetMinimumWarningCompanyPortalVersion(value *string)() {
    m.minimumWarningCompanyPortalVersion = value
}
// Sets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only)
// Parameters:
//  - value : Value to set for the minimumWarningPatchVersion property.
func (m *DefaultManagedAppProtection) SetMinimumWarningPatchVersion(value *string)() {
    m.minimumWarningPatchVersion = value
}
// Sets the minimumWipeCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
// Parameters:
//  - value : Value to set for the minimumWipeCompanyPortalVersion property.
func (m *DefaultManagedAppProtection) SetMinimumWipeCompanyPortalVersion(value *string)() {
    m.minimumWipeCompanyPortalVersion = value
}
// Sets the minimumWipePatchVersion property value. Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data. (Android only)
// Parameters:
//  - value : Value to set for the minimumWipePatchVersion property.
func (m *DefaultManagedAppProtection) SetMinimumWipePatchVersion(value *string)() {
    m.minimumWipePatchVersion = value
}
// Sets the minimumWipeSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
// Parameters:
//  - value : Value to set for the minimumWipeSdkVersion property.
func (m *DefaultManagedAppProtection) SetMinimumWipeSdkVersion(value *string)() {
    m.minimumWipeSdkVersion = value
}
// Sets the protectInboundDataFromUnknownSources property value. Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps. (iOS Only)
// Parameters:
//  - value : Value to set for the protectInboundDataFromUnknownSources property.
func (m *DefaultManagedAppProtection) SetProtectInboundDataFromUnknownSources(value *bool)() {
    m.protectInboundDataFromUnknownSources = value
}
// Sets the requiredAndroidSafetyNetAppsVerificationType property value. Defines the Android SafetyNet Apps Verification requirement for a managed app to work. Possible values are: none, enabled.
// Parameters:
//  - value : Value to set for the requiredAndroidSafetyNetAppsVerificationType property.
func (m *DefaultManagedAppProtection) SetRequiredAndroidSafetyNetAppsVerificationType(value *AndroidManagedAppSafetyNetAppsVerificationType)() {
    m.requiredAndroidSafetyNetAppsVerificationType = value
}
// Sets the requiredAndroidSafetyNetDeviceAttestationType property value. Defines the Android SafetyNet Device Attestation requirement for a managed app to work. Possible values are: none, basicIntegrity, basicIntegrityAndDeviceCertification.
// Parameters:
//  - value : Value to set for the requiredAndroidSafetyNetDeviceAttestationType property.
func (m *DefaultManagedAppProtection) SetRequiredAndroidSafetyNetDeviceAttestationType(value *AndroidManagedAppSafetyNetDeviceAttestationType)() {
    m.requiredAndroidSafetyNetDeviceAttestationType = value
}
// Sets the requiredAndroidSafetyNetEvaluationType property value. Defines the Android SafetyNet evaluation type requirement for a managed app to work. (Android Only). Possible values are: basic, hardwareBacked.
// Parameters:
//  - value : Value to set for the requiredAndroidSafetyNetEvaluationType property.
func (m *DefaultManagedAppProtection) SetRequiredAndroidSafetyNetEvaluationType(value *AndroidManagedAppSafetyNetEvaluationType)() {
    m.requiredAndroidSafetyNetEvaluationType = value
}
// Sets the screenCaptureBlocked property value. Indicates whether screen capture is blocked. (Android only)
// Parameters:
//  - value : Value to set for the screenCaptureBlocked property.
func (m *DefaultManagedAppProtection) SetScreenCaptureBlocked(value *bool)() {
    m.screenCaptureBlocked = value
}
// Sets the thirdPartyKeyboardsBlocked property value. Defines if third party keyboards are allowed while accessing a managed app. (iOS Only)
// Parameters:
//  - value : Value to set for the thirdPartyKeyboardsBlocked property.
func (m *DefaultManagedAppProtection) SetThirdPartyKeyboardsBlocked(value *bool)() {
    m.thirdPartyKeyboardsBlocked = value
}
// Sets the warnAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
// Parameters:
//  - value : Value to set for the warnAfterCompanyPortalUpdateDeferralInDays property.
func (m *DefaultManagedAppProtection) SetWarnAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    m.warnAfterCompanyPortalUpdateDeferralInDays = value
}
// Sets the wipeAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
// Parameters:
//  - value : Value to set for the wipeAfterCompanyPortalUpdateDeferralInDays property.
func (m *DefaultManagedAppProtection) SetWipeAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    m.wipeAfterCompanyPortalUpdateDeferralInDays = value
}
