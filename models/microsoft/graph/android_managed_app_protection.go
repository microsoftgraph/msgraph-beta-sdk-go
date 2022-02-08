package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AndroidManagedAppProtection 
type AndroidManagedAppProtection struct {
    TargetedManagedAppProtection
    // Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work.
    allowedAndroidDeviceManufacturers *string;
    // List of device models allowed, as a string, for the managed app to work.
    allowedAndroidDeviceModels []string;
    // Defines a managed app behavior, either block or wipe, if the specified device manufacturer is not allowed. Possible values are: block, wipe, warn.
    appActionIfAndroidDeviceManufacturerNotAllowed *ManagedAppRemediationAction;
    // Defines a managed app behavior, either block or wipe, if the specified device model is not allowed.
    appActionIfAndroidDeviceModelNotAllowed *ManagedAppRemediationAction;
    // Defines a managed app behavior, either warn or block, if the specified Android App Verification requirment fails. Possible values are: block, wipe, warn.
    appActionIfAndroidSafetyNetAppsVerificationFailed *ManagedAppRemediationAction;
    // Defines a managed app behavior, either warn or block, if the specified Android SafetyNet Attestation requirment fails. Possible values are: block, wipe, warn.
    appActionIfAndroidSafetyNetDeviceAttestationFailed *ManagedAppRemediationAction;
    // Defines a managed app behavior, either warn, block or wipe, if the screen lock is required on android device but is not set.
    appActionIfDeviceLockNotSet *ManagedAppRemediationAction;
    // If Keyboard Restriction is enabled, only keyboards in this approved list will be allowed. A key should be Android package id for a keyboard and value should be a friendly name
    approvedKeyboards []KeyValuePair;
    // List of apps to which the policy is deployed.
    apps []ManagedMobileApp;
    // Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True.
    biometricAuthenticationBlocked *bool;
    // Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
    blockAfterCompanyPortalUpdateDeferralInDays *int32;
    // Whether the app should connect to the configured VPN on launch.
    connectToVpnOnLaunch *bool;
    // Friendly name of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
    customBrowserDisplayName *string;
    // Unique identifier of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
    customBrowserPackageId *string;
    // Friendly name of a custom dialer app to click-to-open a phone number on Android.
    customDialerAppDisplayName *string;
    // PackageId of a custom dialer app to click-to-open a phone number on Android.
    customDialerAppPackageId *string;
    // Count of apps to which the current policy is deployed.
    deployedAppCount *int32;
    // Navigation property to deployment summary of the configuration.
    deploymentSummary *ManagedAppPolicyDeploymentSummary;
    // Defines if any kind of lock must be required on android device
    deviceLockRequired *bool;
    // When this setting is enabled, app level encryption is disabled if device level encryption is enabled
    disableAppEncryptionIfDeviceEncryptionIsEnabled *bool;
    // Indicates whether application data for managed apps should be encrypted
    encryptAppData *bool;
    // App packages in this list will be exempt from the policy and will be able to receive data from managed apps.
    exemptedAppPackages []KeyValuePair;
    // Indicates if keyboard restriction is enabled. If enabled list of approved keyboards must be provided as well.
    keyboardsRestricted *bool;
    // Minimum version of the Company portal that must be installed on the device or app access will be blocked
    minimumRequiredCompanyPortalVersion *string;
    // Define the oldest required Android security patch level a user can have to gain secure access to the app.
    minimumRequiredPatchVersion *string;
    // Minimum version of the Company portal that must be installed on the device or the user will receive a warning
    minimumWarningCompanyPortalVersion *string;
    // Define the oldest recommended Android security patch level a user can have for secure access to the app.
    minimumWarningPatchVersion *string;
    // Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
    minimumWipeCompanyPortalVersion *string;
    // Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data.
    minimumWipePatchVersion *string;
    // Defines the Android SafetyNet Apps Verification requirement for a managed app to work. Possible values are: none, enabled.
    requiredAndroidSafetyNetAppsVerificationType *AndroidManagedAppSafetyNetAppsVerificationType;
    // Defines the Android SafetyNet Device Attestation requirement for a managed app to work. Possible values are: none, basicIntegrity, basicIntegrityAndDeviceCertification.
    requiredAndroidSafetyNetDeviceAttestationType *AndroidManagedAppSafetyNetDeviceAttestationType;
    // Defines the Android SafetyNet evaluation type requirement for a managed app to work.
    requiredAndroidSafetyNetEvaluationType *AndroidManagedAppSafetyNetEvaluationType;
    // Indicates whether a managed user can take screen captures of managed apps
    screenCaptureBlocked *bool;
    // Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
    warnAfterCompanyPortalUpdateDeferralInDays *int32;
    // Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
    wipeAfterCompanyPortalUpdateDeferralInDays *int32;
}
// NewAndroidManagedAppProtection instantiates a new androidManagedAppProtection and sets the default values.
func NewAndroidManagedAppProtection()(*AndroidManagedAppProtection) {
    m := &AndroidManagedAppProtection{
        TargetedManagedAppProtection: *NewTargetedManagedAppProtection(),
    }
    return m
}
// GetAllowedAndroidDeviceManufacturers gets the allowedAndroidDeviceManufacturers property value. Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work.
func (m *AndroidManagedAppProtection) GetAllowedAndroidDeviceManufacturers()(*string) {
    if m == nil {
        return nil
    } else {
        return m.allowedAndroidDeviceManufacturers
    }
}
// GetAllowedAndroidDeviceModels gets the allowedAndroidDeviceModels property value. List of device models allowed, as a string, for the managed app to work.
func (m *AndroidManagedAppProtection) GetAllowedAndroidDeviceModels()([]string) {
    if m == nil {
        return nil
    } else {
        return m.allowedAndroidDeviceModels
    }
}
// GetAppActionIfAndroidDeviceManufacturerNotAllowed gets the appActionIfAndroidDeviceManufacturerNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device manufacturer is not allowed. Possible values are: block, wipe, warn.
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidDeviceManufacturerNotAllowed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfAndroidDeviceManufacturerNotAllowed
    }
}
// GetAppActionIfAndroidDeviceModelNotAllowed gets the appActionIfAndroidDeviceModelNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device model is not allowed.
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidDeviceModelNotAllowed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfAndroidDeviceModelNotAllowed
    }
}
// GetAppActionIfAndroidSafetyNetAppsVerificationFailed gets the appActionIfAndroidSafetyNetAppsVerificationFailed property value. Defines a managed app behavior, either warn or block, if the specified Android App Verification requirment fails. Possible values are: block, wipe, warn.
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidSafetyNetAppsVerificationFailed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfAndroidSafetyNetAppsVerificationFailed
    }
}
// GetAppActionIfAndroidSafetyNetDeviceAttestationFailed gets the appActionIfAndroidSafetyNetDeviceAttestationFailed property value. Defines a managed app behavior, either warn or block, if the specified Android SafetyNet Attestation requirment fails. Possible values are: block, wipe, warn.
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidSafetyNetDeviceAttestationFailed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfAndroidSafetyNetDeviceAttestationFailed
    }
}
// GetAppActionIfDeviceLockNotSet gets the appActionIfDeviceLockNotSet property value. Defines a managed app behavior, either warn, block or wipe, if the screen lock is required on android device but is not set.
func (m *AndroidManagedAppProtection) GetAppActionIfDeviceLockNotSet()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfDeviceLockNotSet
    }
}
// GetApprovedKeyboards gets the approvedKeyboards property value. If Keyboard Restriction is enabled, only keyboards in this approved list will be allowed. A key should be Android package id for a keyboard and value should be a friendly name
func (m *AndroidManagedAppProtection) GetApprovedKeyboards()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.approvedKeyboards
    }
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
func (m *AndroidManagedAppProtection) GetApps()([]ManagedMobileApp) {
    if m == nil {
        return nil
    } else {
        return m.apps
    }
}
// GetBiometricAuthenticationBlocked gets the biometricAuthenticationBlocked property value. Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True.
func (m *AndroidManagedAppProtection) GetBiometricAuthenticationBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.biometricAuthenticationBlocked
    }
}
// GetBlockAfterCompanyPortalUpdateDeferralInDays gets the blockAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
func (m *AndroidManagedAppProtection) GetBlockAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.blockAfterCompanyPortalUpdateDeferralInDays
    }
}
// GetConnectToVpnOnLaunch gets the connectToVpnOnLaunch property value. Whether the app should connect to the configured VPN on launch.
func (m *AndroidManagedAppProtection) GetConnectToVpnOnLaunch()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.connectToVpnOnLaunch
    }
}
// GetCustomBrowserDisplayName gets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *AndroidManagedAppProtection) GetCustomBrowserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customBrowserDisplayName
    }
}
// GetCustomBrowserPackageId gets the customBrowserPackageId property value. Unique identifier of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *AndroidManagedAppProtection) GetCustomBrowserPackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customBrowserPackageId
    }
}
// GetCustomDialerAppDisplayName gets the customDialerAppDisplayName property value. Friendly name of a custom dialer app to click-to-open a phone number on Android.
func (m *AndroidManagedAppProtection) GetCustomDialerAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customDialerAppDisplayName
    }
}
// GetCustomDialerAppPackageId gets the customDialerAppPackageId property value. PackageId of a custom dialer app to click-to-open a phone number on Android.
func (m *AndroidManagedAppProtection) GetCustomDialerAppPackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customDialerAppPackageId
    }
}
// GetDeployedAppCount gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *AndroidManagedAppProtection) GetDeployedAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deployedAppCount
    }
}
// GetDeploymentSummary gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *AndroidManagedAppProtection) GetDeploymentSummary()(*ManagedAppPolicyDeploymentSummary) {
    if m == nil {
        return nil
    } else {
        return m.deploymentSummary
    }
}
// GetDeviceLockRequired gets the deviceLockRequired property value. Defines if any kind of lock must be required on android device
func (m *AndroidManagedAppProtection) GetDeviceLockRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceLockRequired
    }
}
// GetDisableAppEncryptionIfDeviceEncryptionIsEnabled gets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled
func (m *AndroidManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableAppEncryptionIfDeviceEncryptionIsEnabled
    }
}
// GetEncryptAppData gets the encryptAppData property value. Indicates whether application data for managed apps should be encrypted
func (m *AndroidManagedAppProtection) GetEncryptAppData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.encryptAppData
    }
}
// GetExemptedAppPackages gets the exemptedAppPackages property value. App packages in this list will be exempt from the policy and will be able to receive data from managed apps.
func (m *AndroidManagedAppProtection) GetExemptedAppPackages()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.exemptedAppPackages
    }
}
// GetKeyboardsRestricted gets the keyboardsRestricted property value. Indicates if keyboard restriction is enabled. If enabled list of approved keyboards must be provided as well.
func (m *AndroidManagedAppProtection) GetKeyboardsRestricted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keyboardsRestricted
    }
}
// GetMinimumRequiredCompanyPortalVersion gets the minimumRequiredCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or app access will be blocked
func (m *AndroidManagedAppProtection) GetMinimumRequiredCompanyPortalVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredCompanyPortalVersion
    }
}
// GetMinimumRequiredPatchVersion gets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app.
func (m *AndroidManagedAppProtection) GetMinimumRequiredPatchVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredPatchVersion
    }
}
// GetMinimumWarningCompanyPortalVersion gets the minimumWarningCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the user will receive a warning
func (m *AndroidManagedAppProtection) GetMinimumWarningCompanyPortalVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningCompanyPortalVersion
    }
}
// GetMinimumWarningPatchVersion gets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app.
func (m *AndroidManagedAppProtection) GetMinimumWarningPatchVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningPatchVersion
    }
}
// GetMinimumWipeCompanyPortalVersion gets the minimumWipeCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
func (m *AndroidManagedAppProtection) GetMinimumWipeCompanyPortalVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWipeCompanyPortalVersion
    }
}
// GetMinimumWipePatchVersion gets the minimumWipePatchVersion property value. Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data.
func (m *AndroidManagedAppProtection) GetMinimumWipePatchVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWipePatchVersion
    }
}
// GetRequiredAndroidSafetyNetAppsVerificationType gets the requiredAndroidSafetyNetAppsVerificationType property value. Defines the Android SafetyNet Apps Verification requirement for a managed app to work. Possible values are: none, enabled.
func (m *AndroidManagedAppProtection) GetRequiredAndroidSafetyNetAppsVerificationType()(*AndroidManagedAppSafetyNetAppsVerificationType) {
    if m == nil {
        return nil
    } else {
        return m.requiredAndroidSafetyNetAppsVerificationType
    }
}
// GetRequiredAndroidSafetyNetDeviceAttestationType gets the requiredAndroidSafetyNetDeviceAttestationType property value. Defines the Android SafetyNet Device Attestation requirement for a managed app to work. Possible values are: none, basicIntegrity, basicIntegrityAndDeviceCertification.
func (m *AndroidManagedAppProtection) GetRequiredAndroidSafetyNetDeviceAttestationType()(*AndroidManagedAppSafetyNetDeviceAttestationType) {
    if m == nil {
        return nil
    } else {
        return m.requiredAndroidSafetyNetDeviceAttestationType
    }
}
// GetRequiredAndroidSafetyNetEvaluationType gets the requiredAndroidSafetyNetEvaluationType property value. Defines the Android SafetyNet evaluation type requirement for a managed app to work.
func (m *AndroidManagedAppProtection) GetRequiredAndroidSafetyNetEvaluationType()(*AndroidManagedAppSafetyNetEvaluationType) {
    if m == nil {
        return nil
    } else {
        return m.requiredAndroidSafetyNetEvaluationType
    }
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether a managed user can take screen captures of managed apps
func (m *AndroidManagedAppProtection) GetScreenCaptureBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.screenCaptureBlocked
    }
}
// GetWarnAfterCompanyPortalUpdateDeferralInDays gets the warnAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
func (m *AndroidManagedAppProtection) GetWarnAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.warnAfterCompanyPortalUpdateDeferralInDays
    }
}
// GetWipeAfterCompanyPortalUpdateDeferralInDays gets the wipeAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
func (m *AndroidManagedAppProtection) GetWipeAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.wipeAfterCompanyPortalUpdateDeferralInDays
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedAppProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.TargetedManagedAppProtection.GetFieldDeserializers()
    res["allowedAndroidDeviceManufacturers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedAndroidDeviceManufacturers(val)
        }
        return nil
    }
    res["allowedAndroidDeviceModels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedAndroidDeviceModels(res)
        }
        return nil
    }
    res["appActionIfAndroidDeviceManufacturerNotAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAndroidDeviceManufacturerNotAllowed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfAndroidDeviceModelNotAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAndroidDeviceModelNotAllowed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfAndroidSafetyNetAppsVerificationFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAndroidSafetyNetAppsVerificationFailed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfAndroidSafetyNetDeviceAttestationFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfDeviceLockNotSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfDeviceLockNotSet(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["approvedKeyboards"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValuePair))
            }
            m.SetApprovedKeyboards(res)
        }
        return nil
    }
    res["apps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedMobileApp() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedMobileApp, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedMobileApp))
            }
            m.SetApps(res)
        }
        return nil
    }
    res["biometricAuthenticationBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBiometricAuthenticationBlocked(val)
        }
        return nil
    }
    res["blockAfterCompanyPortalUpdateDeferralInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockAfterCompanyPortalUpdateDeferralInDays(val)
        }
        return nil
    }
    res["connectToVpnOnLaunch"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectToVpnOnLaunch(val)
        }
        return nil
    }
    res["customBrowserDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBrowserDisplayName(val)
        }
        return nil
    }
    res["customBrowserPackageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBrowserPackageId(val)
        }
        return nil
    }
    res["customDialerAppDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomDialerAppDisplayName(val)
        }
        return nil
    }
    res["customDialerAppPackageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomDialerAppPackageId(val)
        }
        return nil
    }
    res["deployedAppCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedAppCount(val)
        }
        return nil
    }
    res["deploymentSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicyDeploymentSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentSummary(val.(*ManagedAppPolicyDeploymentSummary))
        }
        return nil
    }
    res["deviceLockRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLockRequired(val)
        }
        return nil
    }
    res["disableAppEncryptionIfDeviceEncryptionIsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(val)
        }
        return nil
    }
    res["encryptAppData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptAppData(val)
        }
        return nil
    }
    res["exemptedAppPackages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValuePair))
            }
            m.SetExemptedAppPackages(res)
        }
        return nil
    }
    res["keyboardsRestricted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardsRestricted(val)
        }
        return nil
    }
    res["minimumRequiredCompanyPortalVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredCompanyPortalVersion(val)
        }
        return nil
    }
    res["minimumRequiredPatchVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredPatchVersion(val)
        }
        return nil
    }
    res["minimumWarningCompanyPortalVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningCompanyPortalVersion(val)
        }
        return nil
    }
    res["minimumWarningPatchVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningPatchVersion(val)
        }
        return nil
    }
    res["minimumWipeCompanyPortalVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipeCompanyPortalVersion(val)
        }
        return nil
    }
    res["minimumWipePatchVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipePatchVersion(val)
        }
        return nil
    }
    res["requiredAndroidSafetyNetAppsVerificationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedAppSafetyNetAppsVerificationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredAndroidSafetyNetAppsVerificationType(val.(*AndroidManagedAppSafetyNetAppsVerificationType))
        }
        return nil
    }
    res["requiredAndroidSafetyNetDeviceAttestationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedAppSafetyNetDeviceAttestationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredAndroidSafetyNetDeviceAttestationType(val.(*AndroidManagedAppSafetyNetDeviceAttestationType))
        }
        return nil
    }
    res["requiredAndroidSafetyNetEvaluationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedAppSafetyNetEvaluationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredAndroidSafetyNetEvaluationType(val.(*AndroidManagedAppSafetyNetEvaluationType))
        }
        return nil
    }
    res["screenCaptureBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenCaptureBlocked(val)
        }
        return nil
    }
    res["warnAfterCompanyPortalUpdateDeferralInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWarnAfterCompanyPortalUpdateDeferralInDays(val)
        }
        return nil
    }
    res["wipeAfterCompanyPortalUpdateDeferralInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWipeAfterCompanyPortalUpdateDeferralInDays(val)
        }
        return nil
    }
    return res
}
func (m *AndroidManagedAppProtection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AndroidManagedAppProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.TargetedManagedAppProtection.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("allowedAndroidDeviceManufacturers", m.GetAllowedAndroidDeviceManufacturers())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedAndroidDeviceModels() != nil {
        err = writer.WriteCollectionOfStringValues("allowedAndroidDeviceModels", m.GetAllowedAndroidDeviceModels())
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAndroidDeviceManufacturerNotAllowed() != nil {
        cast := (*m.GetAppActionIfAndroidDeviceManufacturerNotAllowed()).String()
        err = writer.WriteStringValue("appActionIfAndroidDeviceManufacturerNotAllowed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAndroidDeviceModelNotAllowed() != nil {
        cast := (*m.GetAppActionIfAndroidDeviceModelNotAllowed()).String()
        err = writer.WriteStringValue("appActionIfAndroidDeviceModelNotAllowed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAndroidSafetyNetAppsVerificationFailed() != nil {
        cast := (*m.GetAppActionIfAndroidSafetyNetAppsVerificationFailed()).String()
        err = writer.WriteStringValue("appActionIfAndroidSafetyNetAppsVerificationFailed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAndroidSafetyNetDeviceAttestationFailed() != nil {
        cast := (*m.GetAppActionIfAndroidSafetyNetDeviceAttestationFailed()).String()
        err = writer.WriteStringValue("appActionIfAndroidSafetyNetDeviceAttestationFailed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfDeviceLockNotSet() != nil {
        cast := (*m.GetAppActionIfDeviceLockNotSet()).String()
        err = writer.WriteStringValue("appActionIfDeviceLockNotSet", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApprovedKeyboards() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApprovedKeyboards()))
        for i, v := range m.GetApprovedKeyboards() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("approvedKeyboards", cast)
        if err != nil {
            return err
        }
    }
    if m.GetApps() != nil {
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
        err = writer.WriteBoolValue("encryptAppData", m.GetEncryptAppData())
        if err != nil {
            return err
        }
    }
    if m.GetExemptedAppPackages() != nil {
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
        err = writer.WriteBoolValue("keyboardsRestricted", m.GetKeyboardsRestricted())
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
    if m.GetRequiredAndroidSafetyNetAppsVerificationType() != nil {
        cast := (*m.GetRequiredAndroidSafetyNetAppsVerificationType()).String()
        err = writer.WriteStringValue("requiredAndroidSafetyNetAppsVerificationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequiredAndroidSafetyNetDeviceAttestationType() != nil {
        cast := (*m.GetRequiredAndroidSafetyNetDeviceAttestationType()).String()
        err = writer.WriteStringValue("requiredAndroidSafetyNetDeviceAttestationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequiredAndroidSafetyNetEvaluationType() != nil {
        cast := (*m.GetRequiredAndroidSafetyNetEvaluationType()).String()
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
// SetAllowedAndroidDeviceManufacturers sets the allowedAndroidDeviceManufacturers property value. Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work.
func (m *AndroidManagedAppProtection) SetAllowedAndroidDeviceManufacturers(value *string)() {
    if m != nil {
        m.allowedAndroidDeviceManufacturers = value
    }
}
// SetAllowedAndroidDeviceModels sets the allowedAndroidDeviceModels property value. List of device models allowed, as a string, for the managed app to work.
func (m *AndroidManagedAppProtection) SetAllowedAndroidDeviceModels(value []string)() {
    if m != nil {
        m.allowedAndroidDeviceModels = value
    }
}
// SetAppActionIfAndroidDeviceManufacturerNotAllowed sets the appActionIfAndroidDeviceManufacturerNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device manufacturer is not allowed. Possible values are: block, wipe, warn.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidDeviceManufacturerNotAllowed(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfAndroidDeviceManufacturerNotAllowed = value
    }
}
// SetAppActionIfAndroidDeviceModelNotAllowed sets the appActionIfAndroidDeviceModelNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device model is not allowed.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidDeviceModelNotAllowed(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfAndroidDeviceModelNotAllowed = value
    }
}
// SetAppActionIfAndroidSafetyNetAppsVerificationFailed sets the appActionIfAndroidSafetyNetAppsVerificationFailed property value. Defines a managed app behavior, either warn or block, if the specified Android App Verification requirment fails. Possible values are: block, wipe, warn.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidSafetyNetAppsVerificationFailed(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfAndroidSafetyNetAppsVerificationFailed = value
    }
}
// SetAppActionIfAndroidSafetyNetDeviceAttestationFailed sets the appActionIfAndroidSafetyNetDeviceAttestationFailed property value. Defines a managed app behavior, either warn or block, if the specified Android SafetyNet Attestation requirment fails. Possible values are: block, wipe, warn.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfAndroidSafetyNetDeviceAttestationFailed = value
    }
}
// SetAppActionIfDeviceLockNotSet sets the appActionIfDeviceLockNotSet property value. Defines a managed app behavior, either warn, block or wipe, if the screen lock is required on android device but is not set.
func (m *AndroidManagedAppProtection) SetAppActionIfDeviceLockNotSet(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfDeviceLockNotSet = value
    }
}
// SetApprovedKeyboards sets the approvedKeyboards property value. If Keyboard Restriction is enabled, only keyboards in this approved list will be allowed. A key should be Android package id for a keyboard and value should be a friendly name
func (m *AndroidManagedAppProtection) SetApprovedKeyboards(value []KeyValuePair)() {
    if m != nil {
        m.approvedKeyboards = value
    }
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *AndroidManagedAppProtection) SetApps(value []ManagedMobileApp)() {
    if m != nil {
        m.apps = value
    }
}
// SetBiometricAuthenticationBlocked sets the biometricAuthenticationBlocked property value. Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True.
func (m *AndroidManagedAppProtection) SetBiometricAuthenticationBlocked(value *bool)() {
    if m != nil {
        m.biometricAuthenticationBlocked = value
    }
}
// SetBlockAfterCompanyPortalUpdateDeferralInDays sets the blockAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
func (m *AndroidManagedAppProtection) SetBlockAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    if m != nil {
        m.blockAfterCompanyPortalUpdateDeferralInDays = value
    }
}
// SetConnectToVpnOnLaunch sets the connectToVpnOnLaunch property value. Whether the app should connect to the configured VPN on launch.
func (m *AndroidManagedAppProtection) SetConnectToVpnOnLaunch(value *bool)() {
    if m != nil {
        m.connectToVpnOnLaunch = value
    }
}
// SetCustomBrowserDisplayName sets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *AndroidManagedAppProtection) SetCustomBrowserDisplayName(value *string)() {
    if m != nil {
        m.customBrowserDisplayName = value
    }
}
// SetCustomBrowserPackageId sets the customBrowserPackageId property value. Unique identifier of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *AndroidManagedAppProtection) SetCustomBrowserPackageId(value *string)() {
    if m != nil {
        m.customBrowserPackageId = value
    }
}
// SetCustomDialerAppDisplayName sets the customDialerAppDisplayName property value. Friendly name of a custom dialer app to click-to-open a phone number on Android.
func (m *AndroidManagedAppProtection) SetCustomDialerAppDisplayName(value *string)() {
    if m != nil {
        m.customDialerAppDisplayName = value
    }
}
// SetCustomDialerAppPackageId sets the customDialerAppPackageId property value. PackageId of a custom dialer app to click-to-open a phone number on Android.
func (m *AndroidManagedAppProtection) SetCustomDialerAppPackageId(value *string)() {
    if m != nil {
        m.customDialerAppPackageId = value
    }
}
// SetDeployedAppCount sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *AndroidManagedAppProtection) SetDeployedAppCount(value *int32)() {
    if m != nil {
        m.deployedAppCount = value
    }
}
// SetDeploymentSummary sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *AndroidManagedAppProtection) SetDeploymentSummary(value *ManagedAppPolicyDeploymentSummary)() {
    if m != nil {
        m.deploymentSummary = value
    }
}
// SetDeviceLockRequired sets the deviceLockRequired property value. Defines if any kind of lock must be required on android device
func (m *AndroidManagedAppProtection) SetDeviceLockRequired(value *bool)() {
    if m != nil {
        m.deviceLockRequired = value
    }
}
// SetDisableAppEncryptionIfDeviceEncryptionIsEnabled sets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled
func (m *AndroidManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)() {
    if m != nil {
        m.disableAppEncryptionIfDeviceEncryptionIsEnabled = value
    }
}
// SetEncryptAppData sets the encryptAppData property value. Indicates whether application data for managed apps should be encrypted
func (m *AndroidManagedAppProtection) SetEncryptAppData(value *bool)() {
    if m != nil {
        m.encryptAppData = value
    }
}
// SetExemptedAppPackages sets the exemptedAppPackages property value. App packages in this list will be exempt from the policy and will be able to receive data from managed apps.
func (m *AndroidManagedAppProtection) SetExemptedAppPackages(value []KeyValuePair)() {
    if m != nil {
        m.exemptedAppPackages = value
    }
}
// SetKeyboardsRestricted sets the keyboardsRestricted property value. Indicates if keyboard restriction is enabled. If enabled list of approved keyboards must be provided as well.
func (m *AndroidManagedAppProtection) SetKeyboardsRestricted(value *bool)() {
    if m != nil {
        m.keyboardsRestricted = value
    }
}
// SetMinimumRequiredCompanyPortalVersion sets the minimumRequiredCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or app access will be blocked
func (m *AndroidManagedAppProtection) SetMinimumRequiredCompanyPortalVersion(value *string)() {
    if m != nil {
        m.minimumRequiredCompanyPortalVersion = value
    }
}
// SetMinimumRequiredPatchVersion sets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app.
func (m *AndroidManagedAppProtection) SetMinimumRequiredPatchVersion(value *string)() {
    if m != nil {
        m.minimumRequiredPatchVersion = value
    }
}
// SetMinimumWarningCompanyPortalVersion sets the minimumWarningCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the user will receive a warning
func (m *AndroidManagedAppProtection) SetMinimumWarningCompanyPortalVersion(value *string)() {
    if m != nil {
        m.minimumWarningCompanyPortalVersion = value
    }
}
// SetMinimumWarningPatchVersion sets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app.
func (m *AndroidManagedAppProtection) SetMinimumWarningPatchVersion(value *string)() {
    if m != nil {
        m.minimumWarningPatchVersion = value
    }
}
// SetMinimumWipeCompanyPortalVersion sets the minimumWipeCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
func (m *AndroidManagedAppProtection) SetMinimumWipeCompanyPortalVersion(value *string)() {
    if m != nil {
        m.minimumWipeCompanyPortalVersion = value
    }
}
// SetMinimumWipePatchVersion sets the minimumWipePatchVersion property value. Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data.
func (m *AndroidManagedAppProtection) SetMinimumWipePatchVersion(value *string)() {
    if m != nil {
        m.minimumWipePatchVersion = value
    }
}
// SetRequiredAndroidSafetyNetAppsVerificationType sets the requiredAndroidSafetyNetAppsVerificationType property value. Defines the Android SafetyNet Apps Verification requirement for a managed app to work. Possible values are: none, enabled.
func (m *AndroidManagedAppProtection) SetRequiredAndroidSafetyNetAppsVerificationType(value *AndroidManagedAppSafetyNetAppsVerificationType)() {
    if m != nil {
        m.requiredAndroidSafetyNetAppsVerificationType = value
    }
}
// SetRequiredAndroidSafetyNetDeviceAttestationType sets the requiredAndroidSafetyNetDeviceAttestationType property value. Defines the Android SafetyNet Device Attestation requirement for a managed app to work. Possible values are: none, basicIntegrity, basicIntegrityAndDeviceCertification.
func (m *AndroidManagedAppProtection) SetRequiredAndroidSafetyNetDeviceAttestationType(value *AndroidManagedAppSafetyNetDeviceAttestationType)() {
    if m != nil {
        m.requiredAndroidSafetyNetDeviceAttestationType = value
    }
}
// SetRequiredAndroidSafetyNetEvaluationType sets the requiredAndroidSafetyNetEvaluationType property value. Defines the Android SafetyNet evaluation type requirement for a managed app to work.
func (m *AndroidManagedAppProtection) SetRequiredAndroidSafetyNetEvaluationType(value *AndroidManagedAppSafetyNetEvaluationType)() {
    if m != nil {
        m.requiredAndroidSafetyNetEvaluationType = value
    }
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether a managed user can take screen captures of managed apps
func (m *AndroidManagedAppProtection) SetScreenCaptureBlocked(value *bool)() {
    if m != nil {
        m.screenCaptureBlocked = value
    }
}
// SetWarnAfterCompanyPortalUpdateDeferralInDays sets the warnAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
func (m *AndroidManagedAppProtection) SetWarnAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    if m != nil {
        m.warnAfterCompanyPortalUpdateDeferralInDays = value
    }
}
// SetWipeAfterCompanyPortalUpdateDeferralInDays sets the wipeAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
func (m *AndroidManagedAppProtection) SetWipeAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    if m != nil {
        m.wipeAfterCompanyPortalUpdateDeferralInDays = value
    }
}
