package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DefaultManagedAppProtection 
type DefaultManagedAppProtection struct {
    ManagedAppProtection
    // Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work. (Android only)
    allowedAndroidDeviceManufacturers *string
    // List of device models allowed, as a string, for the managed app to work. (Android Only)
    allowedAndroidDeviceModels []string
    // Semicolon seperated list of device models allowed, as a string, for the managed app to work. (iOS Only)
    allowedIosDeviceModels *string
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
    // If the device does not have a passcode of high complexity or higher, trigger the stored action. Possible values are: block, wipe, warn.
    appActionIfDevicePasscodeComplexityLessThanHigh *ManagedAppRemediationAction
    // If the device does not have a passcode of low complexity or higher, trigger the stored action. Possible values are: block, wipe, warn.
    appActionIfDevicePasscodeComplexityLessThanLow *ManagedAppRemediationAction
    // If the device does not have a passcode of medium complexity or higher, trigger the stored action. Possible values are: block, wipe, warn.
    appActionIfDevicePasscodeComplexityLessThanMedium *ManagedAppRemediationAction
    // An admin initiated action to be applied on a managed app.
    appActionIfIosDeviceModelNotAllowed *ManagedAppRemediationAction
    // Represents the level to which app data is encrypted for managed apps
    appDataEncryptionType *ManagedAppDataEncryptionType
    // List of apps to which the policy is deployed.
    apps []ManagedMobileAppable
    // Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True. (Android Only)
    biometricAuthenticationBlocked *bool
    // Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
    blockAfterCompanyPortalUpdateDeferralInDays *int32
    // Whether the app should connect to the configured VPN on launch (Android only).
    connectToVpnOnLaunch *bool
    // Friendly name of the preferred custom browser to open weblink on Android. (Android only)
    customBrowserDisplayName *string
    // Unique identifier of a custom browser to open weblink on Android. (Android only)
    customBrowserPackageId *string
    // A custom browser protocol to open weblink on iOS. (iOS only)
    customBrowserProtocol *string
    // Friendly name of a custom dialer app to click-to-open a phone number on Android.
    customDialerAppDisplayName *string
    // PackageId of a custom dialer app to click-to-open a phone number on Android.
    customDialerAppPackageId *string
    // Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
    customDialerAppProtocol *string
    // A set of string key and string value pairs to be sent to the affected users, unalterned by this service
    customSettings []KeyValuePairable
    // Count of apps to which the current policy is deployed.
    deployedAppCount *int32
    // Navigation property to deployment summary of the configuration.
    deploymentSummary ManagedAppPolicyDeploymentSummaryable
    // Defines if any kind of lock must be required on device. (android only)
    deviceLockRequired *bool
    // When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only)
    disableAppEncryptionIfDeviceEncryptionIsEnabled *bool
    // Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps. (iOS Only)
    disableProtectionOfManagedOutboundOpenInData *bool
    // Indicates whether managed-app data should be encrypted. (Android only)
    encryptAppData *bool
    // Android App packages in this list will be exempt from the policy and will be able to receive data from managed apps. (Android only)
    exemptedAppPackages []KeyValuePairable
    // iOS Apps in this list will be exempt from the policy and will be able to receive data from managed apps. (iOS Only)
    exemptedAppProtocols []KeyValuePairable
    // Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only)
    faceIdBlocked *bool
    // Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False. (iOS Only)
    filterOpenInToOnlyManagedApps *bool
    // Indicate to the client to enable both biometrics and fingerprints for the app.
    fingerprintAndBiometricEnabled *bool
    // Minimum version of the Company portal that must be installed on the device or app access will be blocked
    minimumRequiredCompanyPortalVersion *string
    // Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only)
    minimumRequiredPatchVersion *string
    // Versions less than the specified version will block the managed app from accessing company data. (iOS Only)
    minimumRequiredSdkVersion *string
    // Minimum version of the Company portal that must be installed on the device or the user will receive a warning
    minimumWarningCompanyPortalVersion *string
    // Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only)
    minimumWarningPatchVersion *string
    // Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
    minimumWipeCompanyPortalVersion *string
    // Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data. (Android only)
    minimumWipePatchVersion *string
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumWipeSdkVersion *string
    // Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps. (iOS Only)
    protectInboundDataFromUnknownSources *bool
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
    // Indicates whether screen capture is blocked. (Android only)
    screenCaptureBlocked *bool
    // Defines if third party keyboards are allowed while accessing a managed app. (iOS Only)
    thirdPartyKeyboardsBlocked *bool
    // Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
    warnAfterCompanyPortalUpdateDeferralInDays *int32
    // Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
    wipeAfterCompanyPortalUpdateDeferralInDays *int32
}
// NewDefaultManagedAppProtection instantiates a new DefaultManagedAppProtection and sets the default values.
func NewDefaultManagedAppProtection()(*DefaultManagedAppProtection) {
    m := &DefaultManagedAppProtection{
        ManagedAppProtection: *NewManagedAppProtection(),
    }
    odataTypeValue := "#microsoft.graph.defaultManagedAppProtection";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDefaultManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDefaultManagedAppProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDefaultManagedAppProtection(), nil
}
// GetAllowedAndroidDeviceManufacturers gets the allowedAndroidDeviceManufacturers property value. Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work. (Android only)
func (m *DefaultManagedAppProtection) GetAllowedAndroidDeviceManufacturers()(*string) {
    return m.allowedAndroidDeviceManufacturers
}
// GetAllowedAndroidDeviceModels gets the allowedAndroidDeviceModels property value. List of device models allowed, as a string, for the managed app to work. (Android Only)
func (m *DefaultManagedAppProtection) GetAllowedAndroidDeviceModels()([]string) {
    return m.allowedAndroidDeviceModels
}
// GetAllowedIosDeviceModels gets the allowedIosDeviceModels property value. Semicolon seperated list of device models allowed, as a string, for the managed app to work. (iOS Only)
func (m *DefaultManagedAppProtection) GetAllowedIosDeviceModels()(*string) {
    return m.allowedIosDeviceModels
}
// GetAppActionIfAndroidDeviceManufacturerNotAllowed gets the appActionIfAndroidDeviceManufacturerNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *DefaultManagedAppProtection) GetAppActionIfAndroidDeviceManufacturerNotAllowed()(*ManagedAppRemediationAction) {
    return m.appActionIfAndroidDeviceManufacturerNotAllowed
}
// GetAppActionIfAndroidDeviceModelNotAllowed gets the appActionIfAndroidDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *DefaultManagedAppProtection) GetAppActionIfAndroidDeviceModelNotAllowed()(*ManagedAppRemediationAction) {
    return m.appActionIfAndroidDeviceModelNotAllowed
}
// GetAppActionIfAndroidSafetyNetAppsVerificationFailed gets the appActionIfAndroidSafetyNetAppsVerificationFailed property value. An admin initiated action to be applied on a managed app.
func (m *DefaultManagedAppProtection) GetAppActionIfAndroidSafetyNetAppsVerificationFailed()(*ManagedAppRemediationAction) {
    return m.appActionIfAndroidSafetyNetAppsVerificationFailed
}
// GetAppActionIfAndroidSafetyNetDeviceAttestationFailed gets the appActionIfAndroidSafetyNetDeviceAttestationFailed property value. An admin initiated action to be applied on a managed app.
func (m *DefaultManagedAppProtection) GetAppActionIfAndroidSafetyNetDeviceAttestationFailed()(*ManagedAppRemediationAction) {
    return m.appActionIfAndroidSafetyNetDeviceAttestationFailed
}
// GetAppActionIfDeviceLockNotSet gets the appActionIfDeviceLockNotSet property value. An admin initiated action to be applied on a managed app.
func (m *DefaultManagedAppProtection) GetAppActionIfDeviceLockNotSet()(*ManagedAppRemediationAction) {
    return m.appActionIfDeviceLockNotSet
}
// GetAppActionIfDevicePasscodeComplexityLessThanHigh gets the appActionIfDevicePasscodeComplexityLessThanHigh property value. If the device does not have a passcode of high complexity or higher, trigger the stored action. Possible values are: block, wipe, warn.
func (m *DefaultManagedAppProtection) GetAppActionIfDevicePasscodeComplexityLessThanHigh()(*ManagedAppRemediationAction) {
    return m.appActionIfDevicePasscodeComplexityLessThanHigh
}
// GetAppActionIfDevicePasscodeComplexityLessThanLow gets the appActionIfDevicePasscodeComplexityLessThanLow property value. If the device does not have a passcode of low complexity or higher, trigger the stored action. Possible values are: block, wipe, warn.
func (m *DefaultManagedAppProtection) GetAppActionIfDevicePasscodeComplexityLessThanLow()(*ManagedAppRemediationAction) {
    return m.appActionIfDevicePasscodeComplexityLessThanLow
}
// GetAppActionIfDevicePasscodeComplexityLessThanMedium gets the appActionIfDevicePasscodeComplexityLessThanMedium property value. If the device does not have a passcode of medium complexity or higher, trigger the stored action. Possible values are: block, wipe, warn.
func (m *DefaultManagedAppProtection) GetAppActionIfDevicePasscodeComplexityLessThanMedium()(*ManagedAppRemediationAction) {
    return m.appActionIfDevicePasscodeComplexityLessThanMedium
}
// GetAppActionIfIosDeviceModelNotAllowed gets the appActionIfIosDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *DefaultManagedAppProtection) GetAppActionIfIosDeviceModelNotAllowed()(*ManagedAppRemediationAction) {
    return m.appActionIfIosDeviceModelNotAllowed
}
// GetAppDataEncryptionType gets the appDataEncryptionType property value. Represents the level to which app data is encrypted for managed apps
func (m *DefaultManagedAppProtection) GetAppDataEncryptionType()(*ManagedAppDataEncryptionType) {
    return m.appDataEncryptionType
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
func (m *DefaultManagedAppProtection) GetApps()([]ManagedMobileAppable) {
    return m.apps
}
// GetBiometricAuthenticationBlocked gets the biometricAuthenticationBlocked property value. Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True. (Android Only)
func (m *DefaultManagedAppProtection) GetBiometricAuthenticationBlocked()(*bool) {
    return m.biometricAuthenticationBlocked
}
// GetBlockAfterCompanyPortalUpdateDeferralInDays gets the blockAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
func (m *DefaultManagedAppProtection) GetBlockAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    return m.blockAfterCompanyPortalUpdateDeferralInDays
}
// GetConnectToVpnOnLaunch gets the connectToVpnOnLaunch property value. Whether the app should connect to the configured VPN on launch (Android only).
func (m *DefaultManagedAppProtection) GetConnectToVpnOnLaunch()(*bool) {
    return m.connectToVpnOnLaunch
}
// GetCustomBrowserDisplayName gets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android. (Android only)
func (m *DefaultManagedAppProtection) GetCustomBrowserDisplayName()(*string) {
    return m.customBrowserDisplayName
}
// GetCustomBrowserPackageId gets the customBrowserPackageId property value. Unique identifier of a custom browser to open weblink on Android. (Android only)
func (m *DefaultManagedAppProtection) GetCustomBrowserPackageId()(*string) {
    return m.customBrowserPackageId
}
// GetCustomBrowserProtocol gets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS. (iOS only)
func (m *DefaultManagedAppProtection) GetCustomBrowserProtocol()(*string) {
    return m.customBrowserProtocol
}
// GetCustomDialerAppDisplayName gets the customDialerAppDisplayName property value. Friendly name of a custom dialer app to click-to-open a phone number on Android.
func (m *DefaultManagedAppProtection) GetCustomDialerAppDisplayName()(*string) {
    return m.customDialerAppDisplayName
}
// GetCustomDialerAppPackageId gets the customDialerAppPackageId property value. PackageId of a custom dialer app to click-to-open a phone number on Android.
func (m *DefaultManagedAppProtection) GetCustomDialerAppPackageId()(*string) {
    return m.customDialerAppPackageId
}
// GetCustomDialerAppProtocol gets the customDialerAppProtocol property value. Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
func (m *DefaultManagedAppProtection) GetCustomDialerAppProtocol()(*string) {
    return m.customDialerAppProtocol
}
// GetCustomSettings gets the customSettings property value. A set of string key and string value pairs to be sent to the affected users, unalterned by this service
func (m *DefaultManagedAppProtection) GetCustomSettings()([]KeyValuePairable) {
    return m.customSettings
}
// GetDeployedAppCount gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *DefaultManagedAppProtection) GetDeployedAppCount()(*int32) {
    return m.deployedAppCount
}
// GetDeploymentSummary gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *DefaultManagedAppProtection) GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable) {
    return m.deploymentSummary
}
// GetDeviceLockRequired gets the deviceLockRequired property value. Defines if any kind of lock must be required on device. (android only)
func (m *DefaultManagedAppProtection) GetDeviceLockRequired()(*bool) {
    return m.deviceLockRequired
}
// GetDisableAppEncryptionIfDeviceEncryptionIsEnabled gets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only)
func (m *DefaultManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool) {
    return m.disableAppEncryptionIfDeviceEncryptionIsEnabled
}
// GetDisableProtectionOfManagedOutboundOpenInData gets the disableProtectionOfManagedOutboundOpenInData property value. Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps. (iOS Only)
func (m *DefaultManagedAppProtection) GetDisableProtectionOfManagedOutboundOpenInData()(*bool) {
    return m.disableProtectionOfManagedOutboundOpenInData
}
// GetEncryptAppData gets the encryptAppData property value. Indicates whether managed-app data should be encrypted. (Android only)
func (m *DefaultManagedAppProtection) GetEncryptAppData()(*bool) {
    return m.encryptAppData
}
// GetExemptedAppPackages gets the exemptedAppPackages property value. Android App packages in this list will be exempt from the policy and will be able to receive data from managed apps. (Android only)
func (m *DefaultManagedAppProtection) GetExemptedAppPackages()([]KeyValuePairable) {
    return m.exemptedAppPackages
}
// GetExemptedAppProtocols gets the exemptedAppProtocols property value. iOS Apps in this list will be exempt from the policy and will be able to receive data from managed apps. (iOS Only)
func (m *DefaultManagedAppProtection) GetExemptedAppProtocols()([]KeyValuePairable) {
    return m.exemptedAppProtocols
}
// GetFaceIdBlocked gets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only)
func (m *DefaultManagedAppProtection) GetFaceIdBlocked()(*bool) {
    return m.faceIdBlocked
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultManagedAppProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedAppProtection.GetFieldDeserializers()
    res["allowedAndroidDeviceManufacturers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAllowedAndroidDeviceManufacturers)
    res["allowedAndroidDeviceModels"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAllowedAndroidDeviceModels)
    res["allowedIosDeviceModels"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAllowedIosDeviceModels)
    res["appActionIfAndroidDeviceManufacturerNotAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfAndroidDeviceManufacturerNotAllowed)
    res["appActionIfAndroidDeviceModelNotAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfAndroidDeviceModelNotAllowed)
    res["appActionIfAndroidSafetyNetAppsVerificationFailed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfAndroidSafetyNetAppsVerificationFailed)
    res["appActionIfAndroidSafetyNetDeviceAttestationFailed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfAndroidSafetyNetDeviceAttestationFailed)
    res["appActionIfDeviceLockNotSet"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfDeviceLockNotSet)
    res["appActionIfDevicePasscodeComplexityLessThanHigh"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfDevicePasscodeComplexityLessThanHigh)
    res["appActionIfDevicePasscodeComplexityLessThanLow"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfDevicePasscodeComplexityLessThanLow)
    res["appActionIfDevicePasscodeComplexityLessThanMedium"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfDevicePasscodeComplexityLessThanMedium)
    res["appActionIfIosDeviceModelNotAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppRemediationAction , m.SetAppActionIfIosDeviceModelNotAllowed)
    res["appDataEncryptionType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedAppDataEncryptionType , m.SetAppDataEncryptionType)
    res["apps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedMobileAppFromDiscriminatorValue , m.SetApps)
    res["biometricAuthenticationBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBiometricAuthenticationBlocked)
    res["blockAfterCompanyPortalUpdateDeferralInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetBlockAfterCompanyPortalUpdateDeferralInDays)
    res["connectToVpnOnLaunch"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetConnectToVpnOnLaunch)
    res["customBrowserDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomBrowserDisplayName)
    res["customBrowserPackageId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomBrowserPackageId)
    res["customBrowserProtocol"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomBrowserProtocol)
    res["customDialerAppDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomDialerAppDisplayName)
    res["customDialerAppPackageId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomDialerAppPackageId)
    res["customDialerAppProtocol"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomDialerAppProtocol)
    res["customSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetCustomSettings)
    res["deployedAppCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDeployedAppCount)
    res["deploymentSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue , m.SetDeploymentSummary)
    res["deviceLockRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDeviceLockRequired)
    res["disableAppEncryptionIfDeviceEncryptionIsEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDisableAppEncryptionIfDeviceEncryptionIsEnabled)
    res["disableProtectionOfManagedOutboundOpenInData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDisableProtectionOfManagedOutboundOpenInData)
    res["encryptAppData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEncryptAppData)
    res["exemptedAppPackages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetExemptedAppPackages)
    res["exemptedAppProtocols"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetExemptedAppProtocols)
    res["faceIdBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetFaceIdBlocked)
    res["filterOpenInToOnlyManagedApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetFilterOpenInToOnlyManagedApps)
    res["fingerprintAndBiometricEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetFingerprintAndBiometricEnabled)
    res["minimumRequiredCompanyPortalVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumRequiredCompanyPortalVersion)
    res["minimumRequiredPatchVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumRequiredPatchVersion)
    res["minimumRequiredSdkVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumRequiredSdkVersion)
    res["minimumWarningCompanyPortalVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumWarningCompanyPortalVersion)
    res["minimumWarningPatchVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumWarningPatchVersion)
    res["minimumWipeCompanyPortalVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumWipeCompanyPortalVersion)
    res["minimumWipePatchVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumWipePatchVersion)
    res["minimumWipeSdkVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinimumWipeSdkVersion)
    res["protectInboundDataFromUnknownSources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetProtectInboundDataFromUnknownSources)
    res["requireClass3Biometrics"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequireClass3Biometrics)
    res["requiredAndroidSafetyNetAppsVerificationType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidManagedAppSafetyNetAppsVerificationType , m.SetRequiredAndroidSafetyNetAppsVerificationType)
    res["requiredAndroidSafetyNetDeviceAttestationType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidManagedAppSafetyNetDeviceAttestationType , m.SetRequiredAndroidSafetyNetDeviceAttestationType)
    res["requiredAndroidSafetyNetEvaluationType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidManagedAppSafetyNetEvaluationType , m.SetRequiredAndroidSafetyNetEvaluationType)
    res["requirePinAfterBiometricChange"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequirePinAfterBiometricChange)
    res["screenCaptureBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetScreenCaptureBlocked)
    res["thirdPartyKeyboardsBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetThirdPartyKeyboardsBlocked)
    res["warnAfterCompanyPortalUpdateDeferralInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWarnAfterCompanyPortalUpdateDeferralInDays)
    res["wipeAfterCompanyPortalUpdateDeferralInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWipeAfterCompanyPortalUpdateDeferralInDays)
    return res
}
// GetFilterOpenInToOnlyManagedApps gets the filterOpenInToOnlyManagedApps property value. Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False. (iOS Only)
func (m *DefaultManagedAppProtection) GetFilterOpenInToOnlyManagedApps()(*bool) {
    return m.filterOpenInToOnlyManagedApps
}
// GetFingerprintAndBiometricEnabled gets the fingerprintAndBiometricEnabled property value. Indicate to the client to enable both biometrics and fingerprints for the app.
func (m *DefaultManagedAppProtection) GetFingerprintAndBiometricEnabled()(*bool) {
    return m.fingerprintAndBiometricEnabled
}
// GetMinimumRequiredCompanyPortalVersion gets the minimumRequiredCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or app access will be blocked
func (m *DefaultManagedAppProtection) GetMinimumRequiredCompanyPortalVersion()(*string) {
    return m.minimumRequiredCompanyPortalVersion
}
// GetMinimumRequiredPatchVersion gets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only)
func (m *DefaultManagedAppProtection) GetMinimumRequiredPatchVersion()(*string) {
    return m.minimumRequiredPatchVersion
}
// GetMinimumRequiredSdkVersion gets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data. (iOS Only)
func (m *DefaultManagedAppProtection) GetMinimumRequiredSdkVersion()(*string) {
    return m.minimumRequiredSdkVersion
}
// GetMinimumWarningCompanyPortalVersion gets the minimumWarningCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the user will receive a warning
func (m *DefaultManagedAppProtection) GetMinimumWarningCompanyPortalVersion()(*string) {
    return m.minimumWarningCompanyPortalVersion
}
// GetMinimumWarningPatchVersion gets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only)
func (m *DefaultManagedAppProtection) GetMinimumWarningPatchVersion()(*string) {
    return m.minimumWarningPatchVersion
}
// GetMinimumWipeCompanyPortalVersion gets the minimumWipeCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
func (m *DefaultManagedAppProtection) GetMinimumWipeCompanyPortalVersion()(*string) {
    return m.minimumWipeCompanyPortalVersion
}
// GetMinimumWipePatchVersion gets the minimumWipePatchVersion property value. Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data. (Android only)
func (m *DefaultManagedAppProtection) GetMinimumWipePatchVersion()(*string) {
    return m.minimumWipePatchVersion
}
// GetMinimumWipeSdkVersion gets the minimumWipeSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *DefaultManagedAppProtection) GetMinimumWipeSdkVersion()(*string) {
    return m.minimumWipeSdkVersion
}
// GetProtectInboundDataFromUnknownSources gets the protectInboundDataFromUnknownSources property value. Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps. (iOS Only)
func (m *DefaultManagedAppProtection) GetProtectInboundDataFromUnknownSources()(*bool) {
    return m.protectInboundDataFromUnknownSources
}
// GetRequireClass3Biometrics gets the requireClass3Biometrics property value. Require user to apply Class 3 Biometrics on their Android device.
func (m *DefaultManagedAppProtection) GetRequireClass3Biometrics()(*bool) {
    return m.requireClass3Biometrics
}
// GetRequiredAndroidSafetyNetAppsVerificationType gets the requiredAndroidSafetyNetAppsVerificationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *DefaultManagedAppProtection) GetRequiredAndroidSafetyNetAppsVerificationType()(*AndroidManagedAppSafetyNetAppsVerificationType) {
    return m.requiredAndroidSafetyNetAppsVerificationType
}
// GetRequiredAndroidSafetyNetDeviceAttestationType gets the requiredAndroidSafetyNetDeviceAttestationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *DefaultManagedAppProtection) GetRequiredAndroidSafetyNetDeviceAttestationType()(*AndroidManagedAppSafetyNetDeviceAttestationType) {
    return m.requiredAndroidSafetyNetDeviceAttestationType
}
// GetRequiredAndroidSafetyNetEvaluationType gets the requiredAndroidSafetyNetEvaluationType property value. An admin enforced Android SafetyNet evaluation type requirement on a managed app.
func (m *DefaultManagedAppProtection) GetRequiredAndroidSafetyNetEvaluationType()(*AndroidManagedAppSafetyNetEvaluationType) {
    return m.requiredAndroidSafetyNetEvaluationType
}
// GetRequirePinAfterBiometricChange gets the requirePinAfterBiometricChange property value. A PIN prompt will override biometric prompts if class 3 biometrics are updated on the device.
func (m *DefaultManagedAppProtection) GetRequirePinAfterBiometricChange()(*bool) {
    return m.requirePinAfterBiometricChange
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether screen capture is blocked. (Android only)
func (m *DefaultManagedAppProtection) GetScreenCaptureBlocked()(*bool) {
    return m.screenCaptureBlocked
}
// GetThirdPartyKeyboardsBlocked gets the thirdPartyKeyboardsBlocked property value. Defines if third party keyboards are allowed while accessing a managed app. (iOS Only)
func (m *DefaultManagedAppProtection) GetThirdPartyKeyboardsBlocked()(*bool) {
    return m.thirdPartyKeyboardsBlocked
}
// GetWarnAfterCompanyPortalUpdateDeferralInDays gets the warnAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
func (m *DefaultManagedAppProtection) GetWarnAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    return m.warnAfterCompanyPortalUpdateDeferralInDays
}
// GetWipeAfterCompanyPortalUpdateDeferralInDays gets the wipeAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
func (m *DefaultManagedAppProtection) GetWipeAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    return m.wipeAfterCompanyPortalUpdateDeferralInDays
}
// Serialize serializes information the current object
func (m *DefaultManagedAppProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetAllowedAndroidDeviceModels() != nil {
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
    if m.GetAppActionIfIosDeviceModelNotAllowed() != nil {
        cast := (*m.GetAppActionIfIosDeviceModelNotAllowed()).String()
        err = writer.WriteStringValue("appActionIfIosDeviceModelNotAllowed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppDataEncryptionType() != nil {
        cast := (*m.GetAppDataEncryptionType()).String()
        err = writer.WriteStringValue("appDataEncryptionType", &cast)
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
    if m.GetCustomSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomSettings())
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
    if m.GetExemptedAppPackages() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExemptedAppPackages())
        err = writer.WriteCollectionOfObjectValues("exemptedAppPackages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExemptedAppProtocols() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExemptedAppProtocols())
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
        err = writer.WriteBoolValue("fingerprintAndBiometricEnabled", m.GetFingerprintAndBiometricEnabled())
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
// SetAllowedAndroidDeviceManufacturers sets the allowedAndroidDeviceManufacturers property value. Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work. (Android only)
func (m *DefaultManagedAppProtection) SetAllowedAndroidDeviceManufacturers(value *string)() {
    m.allowedAndroidDeviceManufacturers = value
}
// SetAllowedAndroidDeviceModels sets the allowedAndroidDeviceModels property value. List of device models allowed, as a string, for the managed app to work. (Android Only)
func (m *DefaultManagedAppProtection) SetAllowedAndroidDeviceModels(value []string)() {
    m.allowedAndroidDeviceModels = value
}
// SetAllowedIosDeviceModels sets the allowedIosDeviceModels property value. Semicolon seperated list of device models allowed, as a string, for the managed app to work. (iOS Only)
func (m *DefaultManagedAppProtection) SetAllowedIosDeviceModels(value *string)() {
    m.allowedIosDeviceModels = value
}
// SetAppActionIfAndroidDeviceManufacturerNotAllowed sets the appActionIfAndroidDeviceManufacturerNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *DefaultManagedAppProtection) SetAppActionIfAndroidDeviceManufacturerNotAllowed(value *ManagedAppRemediationAction)() {
    m.appActionIfAndroidDeviceManufacturerNotAllowed = value
}
// SetAppActionIfAndroidDeviceModelNotAllowed sets the appActionIfAndroidDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *DefaultManagedAppProtection) SetAppActionIfAndroidDeviceModelNotAllowed(value *ManagedAppRemediationAction)() {
    m.appActionIfAndroidDeviceModelNotAllowed = value
}
// SetAppActionIfAndroidSafetyNetAppsVerificationFailed sets the appActionIfAndroidSafetyNetAppsVerificationFailed property value. An admin initiated action to be applied on a managed app.
func (m *DefaultManagedAppProtection) SetAppActionIfAndroidSafetyNetAppsVerificationFailed(value *ManagedAppRemediationAction)() {
    m.appActionIfAndroidSafetyNetAppsVerificationFailed = value
}
// SetAppActionIfAndroidSafetyNetDeviceAttestationFailed sets the appActionIfAndroidSafetyNetDeviceAttestationFailed property value. An admin initiated action to be applied on a managed app.
func (m *DefaultManagedAppProtection) SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(value *ManagedAppRemediationAction)() {
    m.appActionIfAndroidSafetyNetDeviceAttestationFailed = value
}
// SetAppActionIfDeviceLockNotSet sets the appActionIfDeviceLockNotSet property value. An admin initiated action to be applied on a managed app.
func (m *DefaultManagedAppProtection) SetAppActionIfDeviceLockNotSet(value *ManagedAppRemediationAction)() {
    m.appActionIfDeviceLockNotSet = value
}
// SetAppActionIfDevicePasscodeComplexityLessThanHigh sets the appActionIfDevicePasscodeComplexityLessThanHigh property value. If the device does not have a passcode of high complexity or higher, trigger the stored action. Possible values are: block, wipe, warn.
func (m *DefaultManagedAppProtection) SetAppActionIfDevicePasscodeComplexityLessThanHigh(value *ManagedAppRemediationAction)() {
    m.appActionIfDevicePasscodeComplexityLessThanHigh = value
}
// SetAppActionIfDevicePasscodeComplexityLessThanLow sets the appActionIfDevicePasscodeComplexityLessThanLow property value. If the device does not have a passcode of low complexity or higher, trigger the stored action. Possible values are: block, wipe, warn.
func (m *DefaultManagedAppProtection) SetAppActionIfDevicePasscodeComplexityLessThanLow(value *ManagedAppRemediationAction)() {
    m.appActionIfDevicePasscodeComplexityLessThanLow = value
}
// SetAppActionIfDevicePasscodeComplexityLessThanMedium sets the appActionIfDevicePasscodeComplexityLessThanMedium property value. If the device does not have a passcode of medium complexity or higher, trigger the stored action. Possible values are: block, wipe, warn.
func (m *DefaultManagedAppProtection) SetAppActionIfDevicePasscodeComplexityLessThanMedium(value *ManagedAppRemediationAction)() {
    m.appActionIfDevicePasscodeComplexityLessThanMedium = value
}
// SetAppActionIfIosDeviceModelNotAllowed sets the appActionIfIosDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *DefaultManagedAppProtection) SetAppActionIfIosDeviceModelNotAllowed(value *ManagedAppRemediationAction)() {
    m.appActionIfIosDeviceModelNotAllowed = value
}
// SetAppDataEncryptionType sets the appDataEncryptionType property value. Represents the level to which app data is encrypted for managed apps
func (m *DefaultManagedAppProtection) SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)() {
    m.appDataEncryptionType = value
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *DefaultManagedAppProtection) SetApps(value []ManagedMobileAppable)() {
    m.apps = value
}
// SetBiometricAuthenticationBlocked sets the biometricAuthenticationBlocked property value. Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True. (Android Only)
func (m *DefaultManagedAppProtection) SetBiometricAuthenticationBlocked(value *bool)() {
    m.biometricAuthenticationBlocked = value
}
// SetBlockAfterCompanyPortalUpdateDeferralInDays sets the blockAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
func (m *DefaultManagedAppProtection) SetBlockAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    m.blockAfterCompanyPortalUpdateDeferralInDays = value
}
// SetConnectToVpnOnLaunch sets the connectToVpnOnLaunch property value. Whether the app should connect to the configured VPN on launch (Android only).
func (m *DefaultManagedAppProtection) SetConnectToVpnOnLaunch(value *bool)() {
    m.connectToVpnOnLaunch = value
}
// SetCustomBrowserDisplayName sets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android. (Android only)
func (m *DefaultManagedAppProtection) SetCustomBrowserDisplayName(value *string)() {
    m.customBrowserDisplayName = value
}
// SetCustomBrowserPackageId sets the customBrowserPackageId property value. Unique identifier of a custom browser to open weblink on Android. (Android only)
func (m *DefaultManagedAppProtection) SetCustomBrowserPackageId(value *string)() {
    m.customBrowserPackageId = value
}
// SetCustomBrowserProtocol sets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS. (iOS only)
func (m *DefaultManagedAppProtection) SetCustomBrowserProtocol(value *string)() {
    m.customBrowserProtocol = value
}
// SetCustomDialerAppDisplayName sets the customDialerAppDisplayName property value. Friendly name of a custom dialer app to click-to-open a phone number on Android.
func (m *DefaultManagedAppProtection) SetCustomDialerAppDisplayName(value *string)() {
    m.customDialerAppDisplayName = value
}
// SetCustomDialerAppPackageId sets the customDialerAppPackageId property value. PackageId of a custom dialer app to click-to-open a phone number on Android.
func (m *DefaultManagedAppProtection) SetCustomDialerAppPackageId(value *string)() {
    m.customDialerAppPackageId = value
}
// SetCustomDialerAppProtocol sets the customDialerAppProtocol property value. Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
func (m *DefaultManagedAppProtection) SetCustomDialerAppProtocol(value *string)() {
    m.customDialerAppProtocol = value
}
// SetCustomSettings sets the customSettings property value. A set of string key and string value pairs to be sent to the affected users, unalterned by this service
func (m *DefaultManagedAppProtection) SetCustomSettings(value []KeyValuePairable)() {
    m.customSettings = value
}
// SetDeployedAppCount sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *DefaultManagedAppProtection) SetDeployedAppCount(value *int32)() {
    m.deployedAppCount = value
}
// SetDeploymentSummary sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *DefaultManagedAppProtection) SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)() {
    m.deploymentSummary = value
}
// SetDeviceLockRequired sets the deviceLockRequired property value. Defines if any kind of lock must be required on device. (android only)
func (m *DefaultManagedAppProtection) SetDeviceLockRequired(value *bool)() {
    m.deviceLockRequired = value
}
// SetDisableAppEncryptionIfDeviceEncryptionIsEnabled sets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only)
func (m *DefaultManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)() {
    m.disableAppEncryptionIfDeviceEncryptionIsEnabled = value
}
// SetDisableProtectionOfManagedOutboundOpenInData sets the disableProtectionOfManagedOutboundOpenInData property value. Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps. (iOS Only)
func (m *DefaultManagedAppProtection) SetDisableProtectionOfManagedOutboundOpenInData(value *bool)() {
    m.disableProtectionOfManagedOutboundOpenInData = value
}
// SetEncryptAppData sets the encryptAppData property value. Indicates whether managed-app data should be encrypted. (Android only)
func (m *DefaultManagedAppProtection) SetEncryptAppData(value *bool)() {
    m.encryptAppData = value
}
// SetExemptedAppPackages sets the exemptedAppPackages property value. Android App packages in this list will be exempt from the policy and will be able to receive data from managed apps. (Android only)
func (m *DefaultManagedAppProtection) SetExemptedAppPackages(value []KeyValuePairable)() {
    m.exemptedAppPackages = value
}
// SetExemptedAppProtocols sets the exemptedAppProtocols property value. iOS Apps in this list will be exempt from the policy and will be able to receive data from managed apps. (iOS Only)
func (m *DefaultManagedAppProtection) SetExemptedAppProtocols(value []KeyValuePairable)() {
    m.exemptedAppProtocols = value
}
// SetFaceIdBlocked sets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only)
func (m *DefaultManagedAppProtection) SetFaceIdBlocked(value *bool)() {
    m.faceIdBlocked = value
}
// SetFilterOpenInToOnlyManagedApps sets the filterOpenInToOnlyManagedApps property value. Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False. (iOS Only)
func (m *DefaultManagedAppProtection) SetFilterOpenInToOnlyManagedApps(value *bool)() {
    m.filterOpenInToOnlyManagedApps = value
}
// SetFingerprintAndBiometricEnabled sets the fingerprintAndBiometricEnabled property value. Indicate to the client to enable both biometrics and fingerprints for the app.
func (m *DefaultManagedAppProtection) SetFingerprintAndBiometricEnabled(value *bool)() {
    m.fingerprintAndBiometricEnabled = value
}
// SetMinimumRequiredCompanyPortalVersion sets the minimumRequiredCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or app access will be blocked
func (m *DefaultManagedAppProtection) SetMinimumRequiredCompanyPortalVersion(value *string)() {
    m.minimumRequiredCompanyPortalVersion = value
}
// SetMinimumRequiredPatchVersion sets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only)
func (m *DefaultManagedAppProtection) SetMinimumRequiredPatchVersion(value *string)() {
    m.minimumRequiredPatchVersion = value
}
// SetMinimumRequiredSdkVersion sets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data. (iOS Only)
func (m *DefaultManagedAppProtection) SetMinimumRequiredSdkVersion(value *string)() {
    m.minimumRequiredSdkVersion = value
}
// SetMinimumWarningCompanyPortalVersion sets the minimumWarningCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the user will receive a warning
func (m *DefaultManagedAppProtection) SetMinimumWarningCompanyPortalVersion(value *string)() {
    m.minimumWarningCompanyPortalVersion = value
}
// SetMinimumWarningPatchVersion sets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only)
func (m *DefaultManagedAppProtection) SetMinimumWarningPatchVersion(value *string)() {
    m.minimumWarningPatchVersion = value
}
// SetMinimumWipeCompanyPortalVersion sets the minimumWipeCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
func (m *DefaultManagedAppProtection) SetMinimumWipeCompanyPortalVersion(value *string)() {
    m.minimumWipeCompanyPortalVersion = value
}
// SetMinimumWipePatchVersion sets the minimumWipePatchVersion property value. Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data. (Android only)
func (m *DefaultManagedAppProtection) SetMinimumWipePatchVersion(value *string)() {
    m.minimumWipePatchVersion = value
}
// SetMinimumWipeSdkVersion sets the minimumWipeSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *DefaultManagedAppProtection) SetMinimumWipeSdkVersion(value *string)() {
    m.minimumWipeSdkVersion = value
}
// SetProtectInboundDataFromUnknownSources sets the protectInboundDataFromUnknownSources property value. Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps. (iOS Only)
func (m *DefaultManagedAppProtection) SetProtectInboundDataFromUnknownSources(value *bool)() {
    m.protectInboundDataFromUnknownSources = value
}
// SetRequireClass3Biometrics sets the requireClass3Biometrics property value. Require user to apply Class 3 Biometrics on their Android device.
func (m *DefaultManagedAppProtection) SetRequireClass3Biometrics(value *bool)() {
    m.requireClass3Biometrics = value
}
// SetRequiredAndroidSafetyNetAppsVerificationType sets the requiredAndroidSafetyNetAppsVerificationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *DefaultManagedAppProtection) SetRequiredAndroidSafetyNetAppsVerificationType(value *AndroidManagedAppSafetyNetAppsVerificationType)() {
    m.requiredAndroidSafetyNetAppsVerificationType = value
}
// SetRequiredAndroidSafetyNetDeviceAttestationType sets the requiredAndroidSafetyNetDeviceAttestationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *DefaultManagedAppProtection) SetRequiredAndroidSafetyNetDeviceAttestationType(value *AndroidManagedAppSafetyNetDeviceAttestationType)() {
    m.requiredAndroidSafetyNetDeviceAttestationType = value
}
// SetRequiredAndroidSafetyNetEvaluationType sets the requiredAndroidSafetyNetEvaluationType property value. An admin enforced Android SafetyNet evaluation type requirement on a managed app.
func (m *DefaultManagedAppProtection) SetRequiredAndroidSafetyNetEvaluationType(value *AndroidManagedAppSafetyNetEvaluationType)() {
    m.requiredAndroidSafetyNetEvaluationType = value
}
// SetRequirePinAfterBiometricChange sets the requirePinAfterBiometricChange property value. A PIN prompt will override biometric prompts if class 3 biometrics are updated on the device.
func (m *DefaultManagedAppProtection) SetRequirePinAfterBiometricChange(value *bool)() {
    m.requirePinAfterBiometricChange = value
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether screen capture is blocked. (Android only)
func (m *DefaultManagedAppProtection) SetScreenCaptureBlocked(value *bool)() {
    m.screenCaptureBlocked = value
}
// SetThirdPartyKeyboardsBlocked sets the thirdPartyKeyboardsBlocked property value. Defines if third party keyboards are allowed while accessing a managed app. (iOS Only)
func (m *DefaultManagedAppProtection) SetThirdPartyKeyboardsBlocked(value *bool)() {
    m.thirdPartyKeyboardsBlocked = value
}
// SetWarnAfterCompanyPortalUpdateDeferralInDays sets the warnAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
func (m *DefaultManagedAppProtection) SetWarnAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    m.warnAfterCompanyPortalUpdateDeferralInDays = value
}
// SetWipeAfterCompanyPortalUpdateDeferralInDays sets the wipeAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
func (m *DefaultManagedAppProtection) SetWipeAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    m.wipeAfterCompanyPortalUpdateDeferralInDays = value
}
