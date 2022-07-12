package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedAppProtection 
type AndroidManagedAppProtection struct {
    TargetedManagedAppProtection
    // Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work.
    allowedAndroidDeviceManufacturers *string
    // List of device models allowed, as a string, for the managed app to work.
    allowedAndroidDeviceModels []string
    // An admin initiated action to be applied on a managed app.
    appActionIfAndroidDeviceManufacturerNotAllowed *ManagedAppRemediationAction
    // An admin initiated action to be applied on a managed app.
    appActionIfAndroidDeviceModelNotAllowed *ManagedAppRemediationAction
    // An admin initiated action to be applied on a managed app.
    appActionIfAndroidSafetyNetAppsVerificationFailed *ManagedAppRemediationAction
    // An admin initiated action to be applied on a managed app.
    appActionIfAndroidSafetyNetDeviceAttestationFailed *ManagedAppRemediationAction
    // An admin initiated action to be applied on a managed app.
    appActionIfDeviceLockNotSet *ManagedAppRemediationAction
    // If the device does not have a passcode of high complexity or higher, trigger the stored action.
    appActionIfDevicePasscodeComplexityLessThanHigh *ManagedAppRemediationAction
    // If the device does not have a passcode of low complexity or higher, trigger the stored action.
    appActionIfDevicePasscodeComplexityLessThanLow *ManagedAppRemediationAction
    // If the device does not have a passcode of medium complexity or higher, trigger the stored action.
    appActionIfDevicePasscodeComplexityLessThanMedium *ManagedAppRemediationAction
    // If Keyboard Restriction is enabled, only keyboards in this approved list will be allowed. A key should be Android package id for a keyboard and value should be a friendly name
    approvedKeyboards []KeyValuePairable
    // List of apps to which the policy is deployed.
    apps []ManagedMobileAppable
    // Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True.
    biometricAuthenticationBlocked *bool
    // Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
    blockAfterCompanyPortalUpdateDeferralInDays *int32
    // Whether the app should connect to the configured VPN on launch.
    connectToVpnOnLaunch *bool
    // Friendly name of the preferred custom browser to open weblink on Android.
    customBrowserDisplayName *string
    // Unique identifier of a custom browser to open weblink on Android.
    customBrowserPackageId *string
    // Friendly name of a custom dialer app to click-to-open a phone number on Android.
    customDialerAppDisplayName *string
    // PackageId of a custom dialer app to click-to-open a phone number on Android.
    customDialerAppPackageId *string
    // Count of apps to which the current policy is deployed.
    deployedAppCount *int32
    // Navigation property to deployment summary of the configuration.
    deploymentSummary ManagedAppPolicyDeploymentSummaryable
    // Defines if any kind of lock must be required on android device
    deviceLockRequired *bool
    // When this setting is enabled, app level encryption is disabled if device level encryption is enabled
    disableAppEncryptionIfDeviceEncryptionIsEnabled *bool
    // Indicates whether application data for managed apps should be encrypted
    encryptAppData *bool
    // App packages in this list will be exempt from the policy and will be able to receive data from managed apps.
    exemptedAppPackages []KeyValuePairable
    // If null, this setting will be ignored. If false both fingerprints and biometrics will not be enabled. If true, both fingerprints and biometrics will be enabled.
    fingerprintAndBiometricEnabled *bool
    // Indicates if keyboard restriction is enabled. If enabled list of approved keyboards must be provided as well.
    keyboardsRestricted *bool
    // Minimum version of the Company portal that must be installed on the device or app access will be blocked
    minimumRequiredCompanyPortalVersion *string
    // Define the oldest required Android security patch level a user can have to gain secure access to the app.
    minimumRequiredPatchVersion *string
    // Minimum version of the Company portal that must be installed on the device or the user will receive a warning
    minimumWarningCompanyPortalVersion *string
    // Define the oldest recommended Android security patch level a user can have for secure access to the app.
    minimumWarningPatchVersion *string
    // Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
    minimumWipeCompanyPortalVersion *string
    // Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data.
    minimumWipePatchVersion *string
    // Require user to apply Class 3 Biometrics on their Android device.
    requireClass3Biometrics *bool
    // An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
    requiredAndroidSafetyNetAppsVerificationType *AndroidManagedAppSafetyNetAppsVerificationType
    // An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
    requiredAndroidSafetyNetDeviceAttestationType *AndroidManagedAppSafetyNetDeviceAttestationType
    // An admin enforced Android SafetyNet evaluation type requirement on a managed app.
    requiredAndroidSafetyNetEvaluationType *AndroidManagedAppSafetyNetEvaluationType
    // A PIN prompt will override biometric prompts if class 3 biometrics are updated on the device.
    requirePinAfterBiometricChange *bool
    // Indicates whether a managed user can take screen captures of managed apps
    screenCaptureBlocked *bool
    // Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
    warnAfterCompanyPortalUpdateDeferralInDays *int32
    // Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
    wipeAfterCompanyPortalUpdateDeferralInDays *int32
}
// NewAndroidManagedAppProtection instantiates a new AndroidManagedAppProtection and sets the default values.
func NewAndroidManagedAppProtection()(*AndroidManagedAppProtection) {
    m := &AndroidManagedAppProtection{
        TargetedManagedAppProtection: *NewTargetedManagedAppProtection(),
    }
    return m
}
// CreateAndroidManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedAppProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedAppProtection(), nil
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
// GetAppActionIfAndroidDeviceManufacturerNotAllowed gets the appActionIfAndroidDeviceManufacturerNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidDeviceManufacturerNotAllowed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfAndroidDeviceManufacturerNotAllowed
    }
}
// GetAppActionIfAndroidDeviceModelNotAllowed gets the appActionIfAndroidDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidDeviceModelNotAllowed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfAndroidDeviceModelNotAllowed
    }
}
// GetAppActionIfAndroidSafetyNetAppsVerificationFailed gets the appActionIfAndroidSafetyNetAppsVerificationFailed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidSafetyNetAppsVerificationFailed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfAndroidSafetyNetAppsVerificationFailed
    }
}
// GetAppActionIfAndroidSafetyNetDeviceAttestationFailed gets the appActionIfAndroidSafetyNetDeviceAttestationFailed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidSafetyNetDeviceAttestationFailed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfAndroidSafetyNetDeviceAttestationFailed
    }
}
// GetAppActionIfDeviceLockNotSet gets the appActionIfDeviceLockNotSet property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) GetAppActionIfDeviceLockNotSet()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfDeviceLockNotSet
    }
}
// GetAppActionIfDevicePasscodeComplexityLessThanHigh gets the appActionIfDevicePasscodeComplexityLessThanHigh property value. If the device does not have a passcode of high complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) GetAppActionIfDevicePasscodeComplexityLessThanHigh()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfDevicePasscodeComplexityLessThanHigh
    }
}
// GetAppActionIfDevicePasscodeComplexityLessThanLow gets the appActionIfDevicePasscodeComplexityLessThanLow property value. If the device does not have a passcode of low complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) GetAppActionIfDevicePasscodeComplexityLessThanLow()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfDevicePasscodeComplexityLessThanLow
    }
}
// GetAppActionIfDevicePasscodeComplexityLessThanMedium gets the appActionIfDevicePasscodeComplexityLessThanMedium property value. If the device does not have a passcode of medium complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) GetAppActionIfDevicePasscodeComplexityLessThanMedium()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfDevicePasscodeComplexityLessThanMedium
    }
}
// GetApprovedKeyboards gets the approvedKeyboards property value. If Keyboard Restriction is enabled, only keyboards in this approved list will be allowed. A key should be Android package id for a keyboard and value should be a friendly name
func (m *AndroidManagedAppProtection) GetApprovedKeyboards()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.approvedKeyboards
    }
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
func (m *AndroidManagedAppProtection) GetApps()([]ManagedMobileAppable) {
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
// GetCustomBrowserDisplayName gets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android.
func (m *AndroidManagedAppProtection) GetCustomBrowserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customBrowserDisplayName
    }
}
// GetCustomBrowserPackageId gets the customBrowserPackageId property value. Unique identifier of a custom browser to open weblink on Android.
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
func (m *AndroidManagedAppProtection) GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable) {
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
func (m *AndroidManagedAppProtection) GetExemptedAppPackages()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.exemptedAppPackages
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedAppProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TargetedManagedAppProtection.GetFieldDeserializers()
    res["allowedAndroidDeviceManufacturers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedAndroidDeviceManufacturers(val)
        }
        return nil
    }
    res["allowedAndroidDeviceModels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["appActionIfAndroidDeviceManufacturerNotAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAndroidDeviceManufacturerNotAllowed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfAndroidDeviceModelNotAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAndroidDeviceModelNotAllowed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfAndroidSafetyNetAppsVerificationFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAndroidSafetyNetAppsVerificationFailed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfAndroidSafetyNetDeviceAttestationFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfDeviceLockNotSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfDeviceLockNotSet(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfDevicePasscodeComplexityLessThanHigh"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfDevicePasscodeComplexityLessThanHigh(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfDevicePasscodeComplexityLessThanLow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfDevicePasscodeComplexityLessThanLow(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfDevicePasscodeComplexityLessThanMedium"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfDevicePasscodeComplexityLessThanMedium(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["approvedKeyboards"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetApprovedKeyboards(res)
        }
        return nil
    }
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedMobileAppable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedMobileAppable)
            }
            m.SetApps(res)
        }
        return nil
    }
    res["biometricAuthenticationBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBiometricAuthenticationBlocked(val)
        }
        return nil
    }
    res["blockAfterCompanyPortalUpdateDeferralInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockAfterCompanyPortalUpdateDeferralInDays(val)
        }
        return nil
    }
    res["connectToVpnOnLaunch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectToVpnOnLaunch(val)
        }
        return nil
    }
    res["customBrowserDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBrowserDisplayName(val)
        }
        return nil
    }
    res["customBrowserPackageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBrowserPackageId(val)
        }
        return nil
    }
    res["customDialerAppDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomDialerAppDisplayName(val)
        }
        return nil
    }
    res["customDialerAppPackageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomDialerAppPackageId(val)
        }
        return nil
    }
    res["deployedAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedAppCount(val)
        }
        return nil
    }
    res["deploymentSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentSummary(val.(ManagedAppPolicyDeploymentSummaryable))
        }
        return nil
    }
    res["deviceLockRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLockRequired(val)
        }
        return nil
    }
    res["disableAppEncryptionIfDeviceEncryptionIsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(val)
        }
        return nil
    }
    res["encryptAppData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptAppData(val)
        }
        return nil
    }
    res["exemptedAppPackages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetExemptedAppPackages(res)
        }
        return nil
    }
    res["fingerprintAndBiometricEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprintAndBiometricEnabled(val)
        }
        return nil
    }
    res["keyboardsRestricted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardsRestricted(val)
        }
        return nil
    }
    res["minimumRequiredCompanyPortalVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredCompanyPortalVersion(val)
        }
        return nil
    }
    res["minimumRequiredPatchVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredPatchVersion(val)
        }
        return nil
    }
    res["minimumWarningCompanyPortalVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningCompanyPortalVersion(val)
        }
        return nil
    }
    res["minimumWarningPatchVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningPatchVersion(val)
        }
        return nil
    }
    res["minimumWipeCompanyPortalVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipeCompanyPortalVersion(val)
        }
        return nil
    }
    res["minimumWipePatchVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipePatchVersion(val)
        }
        return nil
    }
    res["requireClass3Biometrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireClass3Biometrics(val)
        }
        return nil
    }
    res["requiredAndroidSafetyNetAppsVerificationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedAppSafetyNetAppsVerificationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredAndroidSafetyNetAppsVerificationType(val.(*AndroidManagedAppSafetyNetAppsVerificationType))
        }
        return nil
    }
    res["requiredAndroidSafetyNetDeviceAttestationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedAppSafetyNetDeviceAttestationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredAndroidSafetyNetDeviceAttestationType(val.(*AndroidManagedAppSafetyNetDeviceAttestationType))
        }
        return nil
    }
    res["requiredAndroidSafetyNetEvaluationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedAppSafetyNetEvaluationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredAndroidSafetyNetEvaluationType(val.(*AndroidManagedAppSafetyNetEvaluationType))
        }
        return nil
    }
    res["requirePinAfterBiometricChange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequirePinAfterBiometricChange(val)
        }
        return nil
    }
    res["screenCaptureBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenCaptureBlocked(val)
        }
        return nil
    }
    res["warnAfterCompanyPortalUpdateDeferralInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWarnAfterCompanyPortalUpdateDeferralInDays(val)
        }
        return nil
    }
    res["wipeAfterCompanyPortalUpdateDeferralInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetFingerprintAndBiometricEnabled gets the fingerprintAndBiometricEnabled property value. If null, this setting will be ignored. If false both fingerprints and biometrics will not be enabled. If true, both fingerprints and biometrics will be enabled.
func (m *AndroidManagedAppProtection) GetFingerprintAndBiometricEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fingerprintAndBiometricEnabled
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
// GetRequireClass3Biometrics gets the requireClass3Biometrics property value. Require user to apply Class 3 Biometrics on their Android device.
func (m *AndroidManagedAppProtection) GetRequireClass3Biometrics()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requireClass3Biometrics
    }
}
// GetRequiredAndroidSafetyNetAppsVerificationType gets the requiredAndroidSafetyNetAppsVerificationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *AndroidManagedAppProtection) GetRequiredAndroidSafetyNetAppsVerificationType()(*AndroidManagedAppSafetyNetAppsVerificationType) {
    if m == nil {
        return nil
    } else {
        return m.requiredAndroidSafetyNetAppsVerificationType
    }
}
// GetRequiredAndroidSafetyNetDeviceAttestationType gets the requiredAndroidSafetyNetDeviceAttestationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *AndroidManagedAppProtection) GetRequiredAndroidSafetyNetDeviceAttestationType()(*AndroidManagedAppSafetyNetDeviceAttestationType) {
    if m == nil {
        return nil
    } else {
        return m.requiredAndroidSafetyNetDeviceAttestationType
    }
}
// GetRequiredAndroidSafetyNetEvaluationType gets the requiredAndroidSafetyNetEvaluationType property value. An admin enforced Android SafetyNet evaluation type requirement on a managed app.
func (m *AndroidManagedAppProtection) GetRequiredAndroidSafetyNetEvaluationType()(*AndroidManagedAppSafetyNetEvaluationType) {
    if m == nil {
        return nil
    } else {
        return m.requiredAndroidSafetyNetEvaluationType
    }
}
// GetRequirePinAfterBiometricChange gets the requirePinAfterBiometricChange property value. A PIN prompt will override biometric prompts if class 3 biometrics are updated on the device.
func (m *AndroidManagedAppProtection) GetRequirePinAfterBiometricChange()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requirePinAfterBiometricChange
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
// Serialize serializes information the current object
func (m *AndroidManagedAppProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetAppActionIfDevicePasscodeComplexityLessThanHigh() != nil {
        cast := (*m.GetAppActionIfDevicePasscodeComplexityLessThanHigh()).String()
        err = writer.WriteStringValue("appActionIfDevicePasscodeComplexityLessThanHigh", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfDevicePasscodeComplexityLessThanLow() != nil {
        cast := (*m.GetAppActionIfDevicePasscodeComplexityLessThanLow()).String()
        err = writer.WriteStringValue("appActionIfDevicePasscodeComplexityLessThanLow", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfDevicePasscodeComplexityLessThanMedium() != nil {
        cast := (*m.GetAppActionIfDevicePasscodeComplexityLessThanMedium()).String()
        err = writer.WriteStringValue("appActionIfDevicePasscodeComplexityLessThanMedium", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApprovedKeyboards() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApprovedKeyboards()))
        for i, v := range m.GetApprovedKeyboards() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("approvedKeyboards", cast)
        if err != nil {
            return err
        }
    }
    if m.GetApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExemptedAppPackages()))
        for i, v := range m.GetExemptedAppPackages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("exemptedAppPackages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fingerprintAndBiometricEnabled", m.GetFingerprintAndBiometricEnabled())
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
    {
        err = writer.WriteBoolValue("requireClass3Biometrics", m.GetRequireClass3Biometrics())
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
        err = writer.WriteBoolValue("requirePinAfterBiometricChange", m.GetRequirePinAfterBiometricChange())
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
// SetAppActionIfAndroidDeviceManufacturerNotAllowed sets the appActionIfAndroidDeviceManufacturerNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidDeviceManufacturerNotAllowed(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfAndroidDeviceManufacturerNotAllowed = value
    }
}
// SetAppActionIfAndroidDeviceModelNotAllowed sets the appActionIfAndroidDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidDeviceModelNotAllowed(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfAndroidDeviceModelNotAllowed = value
    }
}
// SetAppActionIfAndroidSafetyNetAppsVerificationFailed sets the appActionIfAndroidSafetyNetAppsVerificationFailed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidSafetyNetAppsVerificationFailed(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfAndroidSafetyNetAppsVerificationFailed = value
    }
}
// SetAppActionIfAndroidSafetyNetDeviceAttestationFailed sets the appActionIfAndroidSafetyNetDeviceAttestationFailed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfAndroidSafetyNetDeviceAttestationFailed = value
    }
}
// SetAppActionIfDeviceLockNotSet sets the appActionIfDeviceLockNotSet property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfDeviceLockNotSet(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfDeviceLockNotSet = value
    }
}
// SetAppActionIfDevicePasscodeComplexityLessThanHigh sets the appActionIfDevicePasscodeComplexityLessThanHigh property value. If the device does not have a passcode of high complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) SetAppActionIfDevicePasscodeComplexityLessThanHigh(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfDevicePasscodeComplexityLessThanHigh = value
    }
}
// SetAppActionIfDevicePasscodeComplexityLessThanLow sets the appActionIfDevicePasscodeComplexityLessThanLow property value. If the device does not have a passcode of low complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) SetAppActionIfDevicePasscodeComplexityLessThanLow(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfDevicePasscodeComplexityLessThanLow = value
    }
}
// SetAppActionIfDevicePasscodeComplexityLessThanMedium sets the appActionIfDevicePasscodeComplexityLessThanMedium property value. If the device does not have a passcode of medium complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) SetAppActionIfDevicePasscodeComplexityLessThanMedium(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfDevicePasscodeComplexityLessThanMedium = value
    }
}
// SetApprovedKeyboards sets the approvedKeyboards property value. If Keyboard Restriction is enabled, only keyboards in this approved list will be allowed. A key should be Android package id for a keyboard and value should be a friendly name
func (m *AndroidManagedAppProtection) SetApprovedKeyboards(value []KeyValuePairable)() {
    if m != nil {
        m.approvedKeyboards = value
    }
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *AndroidManagedAppProtection) SetApps(value []ManagedMobileAppable)() {
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
// SetCustomBrowserDisplayName sets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android.
func (m *AndroidManagedAppProtection) SetCustomBrowserDisplayName(value *string)() {
    if m != nil {
        m.customBrowserDisplayName = value
    }
}
// SetCustomBrowserPackageId sets the customBrowserPackageId property value. Unique identifier of a custom browser to open weblink on Android.
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
func (m *AndroidManagedAppProtection) SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)() {
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
func (m *AndroidManagedAppProtection) SetExemptedAppPackages(value []KeyValuePairable)() {
    if m != nil {
        m.exemptedAppPackages = value
    }
}
// SetFingerprintAndBiometricEnabled sets the fingerprintAndBiometricEnabled property value. If null, this setting will be ignored. If false both fingerprints and biometrics will not be enabled. If true, both fingerprints and biometrics will be enabled.
func (m *AndroidManagedAppProtection) SetFingerprintAndBiometricEnabled(value *bool)() {
    if m != nil {
        m.fingerprintAndBiometricEnabled = value
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
// SetRequireClass3Biometrics sets the requireClass3Biometrics property value. Require user to apply Class 3 Biometrics on their Android device.
func (m *AndroidManagedAppProtection) SetRequireClass3Biometrics(value *bool)() {
    if m != nil {
        m.requireClass3Biometrics = value
    }
}
// SetRequiredAndroidSafetyNetAppsVerificationType sets the requiredAndroidSafetyNetAppsVerificationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *AndroidManagedAppProtection) SetRequiredAndroidSafetyNetAppsVerificationType(value *AndroidManagedAppSafetyNetAppsVerificationType)() {
    if m != nil {
        m.requiredAndroidSafetyNetAppsVerificationType = value
    }
}
// SetRequiredAndroidSafetyNetDeviceAttestationType sets the requiredAndroidSafetyNetDeviceAttestationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *AndroidManagedAppProtection) SetRequiredAndroidSafetyNetDeviceAttestationType(value *AndroidManagedAppSafetyNetDeviceAttestationType)() {
    if m != nil {
        m.requiredAndroidSafetyNetDeviceAttestationType = value
    }
}
// SetRequiredAndroidSafetyNetEvaluationType sets the requiredAndroidSafetyNetEvaluationType property value. An admin enforced Android SafetyNet evaluation type requirement on a managed app.
func (m *AndroidManagedAppProtection) SetRequiredAndroidSafetyNetEvaluationType(value *AndroidManagedAppSafetyNetEvaluationType)() {
    if m != nil {
        m.requiredAndroidSafetyNetEvaluationType = value
    }
}
// SetRequirePinAfterBiometricChange sets the requirePinAfterBiometricChange property value. A PIN prompt will override biometric prompts if class 3 biometrics are updated on the device.
func (m *AndroidManagedAppProtection) SetRequirePinAfterBiometricChange(value *bool)() {
    if m != nil {
        m.requirePinAfterBiometricChange = value
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
