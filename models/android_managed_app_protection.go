package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
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
    odataTypeValue := "#microsoft.graph.androidManagedAppProtection";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedAppProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedAppProtection(), nil
}
// GetAllowedAndroidDeviceManufacturers gets the allowedAndroidDeviceManufacturers property value. Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work.
func (m *AndroidManagedAppProtection) GetAllowedAndroidDeviceManufacturers()(*string) {
    return m.allowedAndroidDeviceManufacturers
}
// GetAllowedAndroidDeviceModels gets the allowedAndroidDeviceModels property value. List of device models allowed, as a string, for the managed app to work.
func (m *AndroidManagedAppProtection) GetAllowedAndroidDeviceModels()([]string) {
    return m.allowedAndroidDeviceModels
}
// GetAppActionIfAndroidDeviceManufacturerNotAllowed gets the appActionIfAndroidDeviceManufacturerNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidDeviceManufacturerNotAllowed()(*ManagedAppRemediationAction) {
    return m.appActionIfAndroidDeviceManufacturerNotAllowed
}
// GetAppActionIfAndroidDeviceModelNotAllowed gets the appActionIfAndroidDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidDeviceModelNotAllowed()(*ManagedAppRemediationAction) {
    return m.appActionIfAndroidDeviceModelNotAllowed
}
// GetAppActionIfAndroidSafetyNetAppsVerificationFailed gets the appActionIfAndroidSafetyNetAppsVerificationFailed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidSafetyNetAppsVerificationFailed()(*ManagedAppRemediationAction) {
    return m.appActionIfAndroidSafetyNetAppsVerificationFailed
}
// GetAppActionIfAndroidSafetyNetDeviceAttestationFailed gets the appActionIfAndroidSafetyNetDeviceAttestationFailed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidSafetyNetDeviceAttestationFailed()(*ManagedAppRemediationAction) {
    return m.appActionIfAndroidSafetyNetDeviceAttestationFailed
}
// GetAppActionIfDeviceLockNotSet gets the appActionIfDeviceLockNotSet property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) GetAppActionIfDeviceLockNotSet()(*ManagedAppRemediationAction) {
    return m.appActionIfDeviceLockNotSet
}
// GetAppActionIfDevicePasscodeComplexityLessThanHigh gets the appActionIfDevicePasscodeComplexityLessThanHigh property value. If the device does not have a passcode of high complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) GetAppActionIfDevicePasscodeComplexityLessThanHigh()(*ManagedAppRemediationAction) {
    return m.appActionIfDevicePasscodeComplexityLessThanHigh
}
// GetAppActionIfDevicePasscodeComplexityLessThanLow gets the appActionIfDevicePasscodeComplexityLessThanLow property value. If the device does not have a passcode of low complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) GetAppActionIfDevicePasscodeComplexityLessThanLow()(*ManagedAppRemediationAction) {
    return m.appActionIfDevicePasscodeComplexityLessThanLow
}
// GetAppActionIfDevicePasscodeComplexityLessThanMedium gets the appActionIfDevicePasscodeComplexityLessThanMedium property value. If the device does not have a passcode of medium complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) GetAppActionIfDevicePasscodeComplexityLessThanMedium()(*ManagedAppRemediationAction) {
    return m.appActionIfDevicePasscodeComplexityLessThanMedium
}
// GetApprovedKeyboards gets the approvedKeyboards property value. If Keyboard Restriction is enabled, only keyboards in this approved list will be allowed. A key should be Android package id for a keyboard and value should be a friendly name
func (m *AndroidManagedAppProtection) GetApprovedKeyboards()([]KeyValuePairable) {
    return m.approvedKeyboards
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
func (m *AndroidManagedAppProtection) GetApps()([]ManagedMobileAppable) {
    return m.apps
}
// GetBiometricAuthenticationBlocked gets the biometricAuthenticationBlocked property value. Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True.
func (m *AndroidManagedAppProtection) GetBiometricAuthenticationBlocked()(*bool) {
    return m.biometricAuthenticationBlocked
}
// GetBlockAfterCompanyPortalUpdateDeferralInDays gets the blockAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
func (m *AndroidManagedAppProtection) GetBlockAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    return m.blockAfterCompanyPortalUpdateDeferralInDays
}
// GetConnectToVpnOnLaunch gets the connectToVpnOnLaunch property value. Whether the app should connect to the configured VPN on launch.
func (m *AndroidManagedAppProtection) GetConnectToVpnOnLaunch()(*bool) {
    return m.connectToVpnOnLaunch
}
// GetCustomBrowserDisplayName gets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android.
func (m *AndroidManagedAppProtection) GetCustomBrowserDisplayName()(*string) {
    return m.customBrowserDisplayName
}
// GetCustomBrowserPackageId gets the customBrowserPackageId property value. Unique identifier of a custom browser to open weblink on Android.
func (m *AndroidManagedAppProtection) GetCustomBrowserPackageId()(*string) {
    return m.customBrowserPackageId
}
// GetCustomDialerAppDisplayName gets the customDialerAppDisplayName property value. Friendly name of a custom dialer app to click-to-open a phone number on Android.
func (m *AndroidManagedAppProtection) GetCustomDialerAppDisplayName()(*string) {
    return m.customDialerAppDisplayName
}
// GetCustomDialerAppPackageId gets the customDialerAppPackageId property value. PackageId of a custom dialer app to click-to-open a phone number on Android.
func (m *AndroidManagedAppProtection) GetCustomDialerAppPackageId()(*string) {
    return m.customDialerAppPackageId
}
// GetDeployedAppCount gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *AndroidManagedAppProtection) GetDeployedAppCount()(*int32) {
    return m.deployedAppCount
}
// GetDeploymentSummary gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *AndroidManagedAppProtection) GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable) {
    return m.deploymentSummary
}
// GetDeviceLockRequired gets the deviceLockRequired property value. Defines if any kind of lock must be required on android device
func (m *AndroidManagedAppProtection) GetDeviceLockRequired()(*bool) {
    return m.deviceLockRequired
}
// GetDisableAppEncryptionIfDeviceEncryptionIsEnabled gets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled
func (m *AndroidManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool) {
    return m.disableAppEncryptionIfDeviceEncryptionIsEnabled
}
// GetEncryptAppData gets the encryptAppData property value. Indicates whether application data for managed apps should be encrypted
func (m *AndroidManagedAppProtection) GetEncryptAppData()(*bool) {
    return m.encryptAppData
}
// GetExemptedAppPackages gets the exemptedAppPackages property value. App packages in this list will be exempt from the policy and will be able to receive data from managed apps.
func (m *AndroidManagedAppProtection) GetExemptedAppPackages()([]KeyValuePairable) {
    return m.exemptedAppPackages
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedAppProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TargetedManagedAppProtection.GetFieldDeserializers()
    res["allowedAndroidDeviceManufacturers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAllowedAndroidDeviceManufacturers)
    res["allowedAndroidDeviceModels"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAllowedAndroidDeviceModels)
    res["appActionIfAndroidDeviceManufacturerNotAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfAndroidDeviceManufacturerNotAllowed)
    res["appActionIfAndroidDeviceModelNotAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfAndroidDeviceModelNotAllowed)
    res["appActionIfAndroidSafetyNetAppsVerificationFailed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfAndroidSafetyNetAppsVerificationFailed)
    res["appActionIfAndroidSafetyNetDeviceAttestationFailed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfAndroidSafetyNetDeviceAttestationFailed)
    res["appActionIfDeviceLockNotSet"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfDeviceLockNotSet)
    res["appActionIfDevicePasscodeComplexityLessThanHigh"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfDevicePasscodeComplexityLessThanHigh)
    res["appActionIfDevicePasscodeComplexityLessThanLow"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfDevicePasscodeComplexityLessThanLow)
    res["appActionIfDevicePasscodeComplexityLessThanMedium"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfDevicePasscodeComplexityLessThanMedium)
    res["approvedKeyboards"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetApprovedKeyboards)
    res["apps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedMobileAppFromDiscriminatorValue , m.SetApps)
    res["biometricAuthenticationBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBiometricAuthenticationBlocked)
    res["blockAfterCompanyPortalUpdateDeferralInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetBlockAfterCompanyPortalUpdateDeferralInDays)
    res["connectToVpnOnLaunch"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetConnectToVpnOnLaunch)
    res["customBrowserDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomBrowserDisplayName)
    res["customBrowserPackageId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomBrowserPackageId)
    res["customDialerAppDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomDialerAppDisplayName)
    res["customDialerAppPackageId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomDialerAppPackageId)
    res["deployedAppCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDeployedAppCount)
    res["deploymentSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue , m.SetDeploymentSummary)
    res["deviceLockRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDeviceLockRequired)
    res["disableAppEncryptionIfDeviceEncryptionIsEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDisableAppEncryptionIfDeviceEncryptionIsEnabled)
    res["encryptAppData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEncryptAppData)
    res["exemptedAppPackages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetExemptedAppPackages)
    res["fingerprintAndBiometricEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetFingerprintAndBiometricEnabled)
    res["keyboardsRestricted"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetKeyboardsRestricted)
    res["minimumRequiredCompanyPortalVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumRequiredCompanyPortalVersion)
    res["minimumRequiredPatchVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumRequiredPatchVersion)
    res["minimumWarningCompanyPortalVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumWarningCompanyPortalVersion)
    res["minimumWarningPatchVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumWarningPatchVersion)
    res["minimumWipeCompanyPortalVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumWipeCompanyPortalVersion)
    res["minimumWipePatchVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumWipePatchVersion)
    res["requireClass3Biometrics"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequireClass3Biometrics)
    res["requiredAndroidSafetyNetAppsVerificationType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidManagedAppSafetyNetAppsVerificationType , m.SetRequiredAndroidSafetyNetAppsVerificationType)
    res["requiredAndroidSafetyNetDeviceAttestationType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidManagedAppSafetyNetDeviceAttestationType , m.SetRequiredAndroidSafetyNetDeviceAttestationType)
    res["requiredAndroidSafetyNetEvaluationType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidManagedAppSafetyNetEvaluationType , m.SetRequiredAndroidSafetyNetEvaluationType)
    res["requirePinAfterBiometricChange"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequirePinAfterBiometricChange)
    res["screenCaptureBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetScreenCaptureBlocked)
    res["warnAfterCompanyPortalUpdateDeferralInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWarnAfterCompanyPortalUpdateDeferralInDays)
    res["wipeAfterCompanyPortalUpdateDeferralInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWipeAfterCompanyPortalUpdateDeferralInDays)
    return res
}
// GetFingerprintAndBiometricEnabled gets the fingerprintAndBiometricEnabled property value. If null, this setting will be ignored. If false both fingerprints and biometrics will not be enabled. If true, both fingerprints and biometrics will be enabled.
func (m *AndroidManagedAppProtection) GetFingerprintAndBiometricEnabled()(*bool) {
    return m.fingerprintAndBiometricEnabled
}
// GetKeyboardsRestricted gets the keyboardsRestricted property value. Indicates if keyboard restriction is enabled. If enabled list of approved keyboards must be provided as well.
func (m *AndroidManagedAppProtection) GetKeyboardsRestricted()(*bool) {
    return m.keyboardsRestricted
}
// GetMinimumRequiredCompanyPortalVersion gets the minimumRequiredCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or app access will be blocked
func (m *AndroidManagedAppProtection) GetMinimumRequiredCompanyPortalVersion()(*string) {
    return m.minimumRequiredCompanyPortalVersion
}
// GetMinimumRequiredPatchVersion gets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app.
func (m *AndroidManagedAppProtection) GetMinimumRequiredPatchVersion()(*string) {
    return m.minimumRequiredPatchVersion
}
// GetMinimumWarningCompanyPortalVersion gets the minimumWarningCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the user will receive a warning
func (m *AndroidManagedAppProtection) GetMinimumWarningCompanyPortalVersion()(*string) {
    return m.minimumWarningCompanyPortalVersion
}
// GetMinimumWarningPatchVersion gets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app.
func (m *AndroidManagedAppProtection) GetMinimumWarningPatchVersion()(*string) {
    return m.minimumWarningPatchVersion
}
// GetMinimumWipeCompanyPortalVersion gets the minimumWipeCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
func (m *AndroidManagedAppProtection) GetMinimumWipeCompanyPortalVersion()(*string) {
    return m.minimumWipeCompanyPortalVersion
}
// GetMinimumWipePatchVersion gets the minimumWipePatchVersion property value. Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data.
func (m *AndroidManagedAppProtection) GetMinimumWipePatchVersion()(*string) {
    return m.minimumWipePatchVersion
}
// GetRequireClass3Biometrics gets the requireClass3Biometrics property value. Require user to apply Class 3 Biometrics on their Android device.
func (m *AndroidManagedAppProtection) GetRequireClass3Biometrics()(*bool) {
    return m.requireClass3Biometrics
}
// GetRequiredAndroidSafetyNetAppsVerificationType gets the requiredAndroidSafetyNetAppsVerificationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *AndroidManagedAppProtection) GetRequiredAndroidSafetyNetAppsVerificationType()(*AndroidManagedAppSafetyNetAppsVerificationType) {
    return m.requiredAndroidSafetyNetAppsVerificationType
}
// GetRequiredAndroidSafetyNetDeviceAttestationType gets the requiredAndroidSafetyNetDeviceAttestationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *AndroidManagedAppProtection) GetRequiredAndroidSafetyNetDeviceAttestationType()(*AndroidManagedAppSafetyNetDeviceAttestationType) {
    return m.requiredAndroidSafetyNetDeviceAttestationType
}
// GetRequiredAndroidSafetyNetEvaluationType gets the requiredAndroidSafetyNetEvaluationType property value. An admin enforced Android SafetyNet evaluation type requirement on a managed app.
func (m *AndroidManagedAppProtection) GetRequiredAndroidSafetyNetEvaluationType()(*AndroidManagedAppSafetyNetEvaluationType) {
    return m.requiredAndroidSafetyNetEvaluationType
}
// GetRequirePinAfterBiometricChange gets the requirePinAfterBiometricChange property value. A PIN prompt will override biometric prompts if class 3 biometrics are updated on the device.
func (m *AndroidManagedAppProtection) GetRequirePinAfterBiometricChange()(*bool) {
    return m.requirePinAfterBiometricChange
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether a managed user can take screen captures of managed apps
func (m *AndroidManagedAppProtection) GetScreenCaptureBlocked()(*bool) {
    return m.screenCaptureBlocked
}
// GetWarnAfterCompanyPortalUpdateDeferralInDays gets the warnAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
func (m *AndroidManagedAppProtection) GetWarnAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    return m.warnAfterCompanyPortalUpdateDeferralInDays
}
// GetWipeAfterCompanyPortalUpdateDeferralInDays gets the wipeAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
func (m *AndroidManagedAppProtection) GetWipeAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    return m.wipeAfterCompanyPortalUpdateDeferralInDays
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetApprovedKeyboards())
        err = writer.WriteCollectionOfObjectValues("approvedKeyboards", cast)
        if err != nil {
            return err
        }
    }
    if m.GetApps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetApps())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExemptedAppPackages())
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
    m.allowedAndroidDeviceManufacturers = value
}
// SetAllowedAndroidDeviceModels sets the allowedAndroidDeviceModels property value. List of device models allowed, as a string, for the managed app to work.
func (m *AndroidManagedAppProtection) SetAllowedAndroidDeviceModels(value []string)() {
    m.allowedAndroidDeviceModels = value
}
// SetAppActionIfAndroidDeviceManufacturerNotAllowed sets the appActionIfAndroidDeviceManufacturerNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidDeviceManufacturerNotAllowed(value *ManagedAppRemediationAction)() {
    m.appActionIfAndroidDeviceManufacturerNotAllowed = value
}
// SetAppActionIfAndroidDeviceModelNotAllowed sets the appActionIfAndroidDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidDeviceModelNotAllowed(value *ManagedAppRemediationAction)() {
    m.appActionIfAndroidDeviceModelNotAllowed = value
}
// SetAppActionIfAndroidSafetyNetAppsVerificationFailed sets the appActionIfAndroidSafetyNetAppsVerificationFailed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidSafetyNetAppsVerificationFailed(value *ManagedAppRemediationAction)() {
    m.appActionIfAndroidSafetyNetAppsVerificationFailed = value
}
// SetAppActionIfAndroidSafetyNetDeviceAttestationFailed sets the appActionIfAndroidSafetyNetDeviceAttestationFailed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(value *ManagedAppRemediationAction)() {
    m.appActionIfAndroidSafetyNetDeviceAttestationFailed = value
}
// SetAppActionIfDeviceLockNotSet sets the appActionIfDeviceLockNotSet property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfDeviceLockNotSet(value *ManagedAppRemediationAction)() {
    m.appActionIfDeviceLockNotSet = value
}
// SetAppActionIfDevicePasscodeComplexityLessThanHigh sets the appActionIfDevicePasscodeComplexityLessThanHigh property value. If the device does not have a passcode of high complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) SetAppActionIfDevicePasscodeComplexityLessThanHigh(value *ManagedAppRemediationAction)() {
    m.appActionIfDevicePasscodeComplexityLessThanHigh = value
}
// SetAppActionIfDevicePasscodeComplexityLessThanLow sets the appActionIfDevicePasscodeComplexityLessThanLow property value. If the device does not have a passcode of low complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) SetAppActionIfDevicePasscodeComplexityLessThanLow(value *ManagedAppRemediationAction)() {
    m.appActionIfDevicePasscodeComplexityLessThanLow = value
}
// SetAppActionIfDevicePasscodeComplexityLessThanMedium sets the appActionIfDevicePasscodeComplexityLessThanMedium property value. If the device does not have a passcode of medium complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) SetAppActionIfDevicePasscodeComplexityLessThanMedium(value *ManagedAppRemediationAction)() {
    m.appActionIfDevicePasscodeComplexityLessThanMedium = value
}
// SetApprovedKeyboards sets the approvedKeyboards property value. If Keyboard Restriction is enabled, only keyboards in this approved list will be allowed. A key should be Android package id for a keyboard and value should be a friendly name
func (m *AndroidManagedAppProtection) SetApprovedKeyboards(value []KeyValuePairable)() {
    m.approvedKeyboards = value
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *AndroidManagedAppProtection) SetApps(value []ManagedMobileAppable)() {
    m.apps = value
}
// SetBiometricAuthenticationBlocked sets the biometricAuthenticationBlocked property value. Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True.
func (m *AndroidManagedAppProtection) SetBiometricAuthenticationBlocked(value *bool)() {
    m.biometricAuthenticationBlocked = value
}
// SetBlockAfterCompanyPortalUpdateDeferralInDays sets the blockAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
func (m *AndroidManagedAppProtection) SetBlockAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    m.blockAfterCompanyPortalUpdateDeferralInDays = value
}
// SetConnectToVpnOnLaunch sets the connectToVpnOnLaunch property value. Whether the app should connect to the configured VPN on launch.
func (m *AndroidManagedAppProtection) SetConnectToVpnOnLaunch(value *bool)() {
    m.connectToVpnOnLaunch = value
}
// SetCustomBrowserDisplayName sets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android.
func (m *AndroidManagedAppProtection) SetCustomBrowserDisplayName(value *string)() {
    m.customBrowserDisplayName = value
}
// SetCustomBrowserPackageId sets the customBrowserPackageId property value. Unique identifier of a custom browser to open weblink on Android.
func (m *AndroidManagedAppProtection) SetCustomBrowserPackageId(value *string)() {
    m.customBrowserPackageId = value
}
// SetCustomDialerAppDisplayName sets the customDialerAppDisplayName property value. Friendly name of a custom dialer app to click-to-open a phone number on Android.
func (m *AndroidManagedAppProtection) SetCustomDialerAppDisplayName(value *string)() {
    m.customDialerAppDisplayName = value
}
// SetCustomDialerAppPackageId sets the customDialerAppPackageId property value. PackageId of a custom dialer app to click-to-open a phone number on Android.
func (m *AndroidManagedAppProtection) SetCustomDialerAppPackageId(value *string)() {
    m.customDialerAppPackageId = value
}
// SetDeployedAppCount sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *AndroidManagedAppProtection) SetDeployedAppCount(value *int32)() {
    m.deployedAppCount = value
}
// SetDeploymentSummary sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *AndroidManagedAppProtection) SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)() {
    m.deploymentSummary = value
}
// SetDeviceLockRequired sets the deviceLockRequired property value. Defines if any kind of lock must be required on android device
func (m *AndroidManagedAppProtection) SetDeviceLockRequired(value *bool)() {
    m.deviceLockRequired = value
}
// SetDisableAppEncryptionIfDeviceEncryptionIsEnabled sets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled
func (m *AndroidManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)() {
    m.disableAppEncryptionIfDeviceEncryptionIsEnabled = value
}
// SetEncryptAppData sets the encryptAppData property value. Indicates whether application data for managed apps should be encrypted
func (m *AndroidManagedAppProtection) SetEncryptAppData(value *bool)() {
    m.encryptAppData = value
}
// SetExemptedAppPackages sets the exemptedAppPackages property value. App packages in this list will be exempt from the policy and will be able to receive data from managed apps.
func (m *AndroidManagedAppProtection) SetExemptedAppPackages(value []KeyValuePairable)() {
    m.exemptedAppPackages = value
}
// SetFingerprintAndBiometricEnabled sets the fingerprintAndBiometricEnabled property value. If null, this setting will be ignored. If false both fingerprints and biometrics will not be enabled. If true, both fingerprints and biometrics will be enabled.
func (m *AndroidManagedAppProtection) SetFingerprintAndBiometricEnabled(value *bool)() {
    m.fingerprintAndBiometricEnabled = value
}
// SetKeyboardsRestricted sets the keyboardsRestricted property value. Indicates if keyboard restriction is enabled. If enabled list of approved keyboards must be provided as well.
func (m *AndroidManagedAppProtection) SetKeyboardsRestricted(value *bool)() {
    m.keyboardsRestricted = value
}
// SetMinimumRequiredCompanyPortalVersion sets the minimumRequiredCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or app access will be blocked
func (m *AndroidManagedAppProtection) SetMinimumRequiredCompanyPortalVersion(value *string)() {
    m.minimumRequiredCompanyPortalVersion = value
}
// SetMinimumRequiredPatchVersion sets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app.
func (m *AndroidManagedAppProtection) SetMinimumRequiredPatchVersion(value *string)() {
    m.minimumRequiredPatchVersion = value
}
// SetMinimumWarningCompanyPortalVersion sets the minimumWarningCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the user will receive a warning
func (m *AndroidManagedAppProtection) SetMinimumWarningCompanyPortalVersion(value *string)() {
    m.minimumWarningCompanyPortalVersion = value
}
// SetMinimumWarningPatchVersion sets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app.
func (m *AndroidManagedAppProtection) SetMinimumWarningPatchVersion(value *string)() {
    m.minimumWarningPatchVersion = value
}
// SetMinimumWipeCompanyPortalVersion sets the minimumWipeCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
func (m *AndroidManagedAppProtection) SetMinimumWipeCompanyPortalVersion(value *string)() {
    m.minimumWipeCompanyPortalVersion = value
}
// SetMinimumWipePatchVersion sets the minimumWipePatchVersion property value. Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data.
func (m *AndroidManagedAppProtection) SetMinimumWipePatchVersion(value *string)() {
    m.minimumWipePatchVersion = value
}
// SetRequireClass3Biometrics sets the requireClass3Biometrics property value. Require user to apply Class 3 Biometrics on their Android device.
func (m *AndroidManagedAppProtection) SetRequireClass3Biometrics(value *bool)() {
    m.requireClass3Biometrics = value
}
// SetRequiredAndroidSafetyNetAppsVerificationType sets the requiredAndroidSafetyNetAppsVerificationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *AndroidManagedAppProtection) SetRequiredAndroidSafetyNetAppsVerificationType(value *AndroidManagedAppSafetyNetAppsVerificationType)() {
    m.requiredAndroidSafetyNetAppsVerificationType = value
}
// SetRequiredAndroidSafetyNetDeviceAttestationType sets the requiredAndroidSafetyNetDeviceAttestationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *AndroidManagedAppProtection) SetRequiredAndroidSafetyNetDeviceAttestationType(value *AndroidManagedAppSafetyNetDeviceAttestationType)() {
    m.requiredAndroidSafetyNetDeviceAttestationType = value
}
// SetRequiredAndroidSafetyNetEvaluationType sets the requiredAndroidSafetyNetEvaluationType property value. An admin enforced Android SafetyNet evaluation type requirement on a managed app.
func (m *AndroidManagedAppProtection) SetRequiredAndroidSafetyNetEvaluationType(value *AndroidManagedAppSafetyNetEvaluationType)() {
    m.requiredAndroidSafetyNetEvaluationType = value
}
// SetRequirePinAfterBiometricChange sets the requirePinAfterBiometricChange property value. A PIN prompt will override biometric prompts if class 3 biometrics are updated on the device.
func (m *AndroidManagedAppProtection) SetRequirePinAfterBiometricChange(value *bool)() {
    m.requirePinAfterBiometricChange = value
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether a managed user can take screen captures of managed apps
func (m *AndroidManagedAppProtection) SetScreenCaptureBlocked(value *bool)() {
    m.screenCaptureBlocked = value
}
// SetWarnAfterCompanyPortalUpdateDeferralInDays sets the warnAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
func (m *AndroidManagedAppProtection) SetWarnAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    m.warnAfterCompanyPortalUpdateDeferralInDays = value
}
// SetWipeAfterCompanyPortalUpdateDeferralInDays sets the wipeAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
func (m *AndroidManagedAppProtection) SetWipeAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    m.wipeAfterCompanyPortalUpdateDeferralInDays = value
}
