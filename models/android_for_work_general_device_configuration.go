package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkGeneralDeviceConfiguration 
type AndroidForWorkGeneralDeviceConfiguration struct {
    DeviceConfiguration
    // Indicates whether or not to block face unlock.
    passwordBlockFaceUnlock *bool
    // Indicates whether or not to block fingerprint unlock.
    passwordBlockFingerprintUnlock *bool
    // Indicates whether or not to block iris unlock.
    passwordBlockIrisUnlock *bool
    // Indicates whether or not to block Smart Lock and other trust agents.
    passwordBlockTrustAgents *bool
    // Number of days before the password expires. Valid values 1 to 365
    passwordExpirationDays *int32
    // Minimum length of passwords. Valid values 4 to 16
    passwordMinimumLength *int32
    // Minutes of inactivity before the screen times out.
    passwordMinutesOfInactivityBeforeScreenTimeout *int32
    // Number of previous passwords to block. Valid values 0 to 24
    passwordPreviousPasswordBlockCount *int32
    // Type of password that is required. Possible values are: deviceDefault, lowSecurityBiometric, required, atLeastNumeric, numericComplex, atLeastAlphabetic, atLeastAlphanumeric, alphanumericWithSymbols.
    passwordRequiredType *AndroidForWorkRequiredPasswordType
    // Number of sign in failures allowed before factory reset. Valid values 1 to 16
    passwordSignInFailureCountBeforeFactoryReset *int32
    // Require the Android Verify apps feature is turned on.
    securityRequireVerifyApps *bool
    // Enable lockdown mode for always-on VPN.
    vpnAlwaysOnPackageIdentifier *string
    // Enable lockdown mode for always-on VPN.
    vpnEnableAlwaysOnLockdownMode *bool
    // Allow widgets from work profile apps.
    workProfileAllowWidgets *bool
    // Block users from adding/removing accounts in work profile.
    workProfileBlockAddingAccounts *bool
    // Block work profile camera.
    workProfileBlockCamera *bool
    // Block display work profile caller ID in personal profile.
    workProfileBlockCrossProfileCallerId *bool
    // Block work profile contacts availability in personal profile.
    workProfileBlockCrossProfileContactsSearch *bool
    // Boolean that indicates if the setting disallow cross profile copy/paste is enabled.
    workProfileBlockCrossProfileCopyPaste *bool
    // Indicates whether or not to block notifications while device locked.
    workProfileBlockNotificationsWhileDeviceLocked *bool
    // Prevent app installations from unknown sources in the personal profile.
    workProfileBlockPersonalAppInstallsFromUnknownSources *bool
    // Block screen capture in work profile.
    workProfileBlockScreenCapture *bool
    // Allow bluetooth devices to access enterprise contacts.
    workProfileBluetoothEnableContactSharing *bool
    // Type of data sharing that is allowed. Possible values are: deviceDefault, preventAny, allowPersonalToWork, noRestrictions.
    workProfileDataSharingType *AndroidForWorkCrossProfileDataSharingType
    // Type of password that is required. Possible values are: deviceDefault, prompt, autoGrant, autoDeny.
    workProfileDefaultAppPermissionPolicy *AndroidForWorkDefaultAppPermissionPolicyType
    // Indicates whether or not to block face unlock for work profile.
    workProfilePasswordBlockFaceUnlock *bool
    // Indicates whether or not to block fingerprint unlock for work profile.
    workProfilePasswordBlockFingerprintUnlock *bool
    // Indicates whether or not to block iris unlock for work profile.
    workProfilePasswordBlockIrisUnlock *bool
    // Indicates whether or not to block Smart Lock and other trust agents for work profile.
    workProfilePasswordBlockTrustAgents *bool
    // Number of days before the work profile password expires. Valid values 1 to 365
    workProfilePasswordExpirationDays *int32
    // Minimum length of work profile password. Valid values 4 to 16
    workProfilePasswordMinimumLength *int32
    // Minimum # of letter characters required in work profile password. Valid values 1 to 10
    workProfilePasswordMinLetterCharacters *int32
    // Minimum # of lower-case characters required in work profile password. Valid values 1 to 10
    workProfilePasswordMinLowerCaseCharacters *int32
    // Minimum # of non-letter characters required in work profile password. Valid values 1 to 10
    workProfilePasswordMinNonLetterCharacters *int32
    // Minimum # of numeric characters required in work profile password. Valid values 1 to 10
    workProfilePasswordMinNumericCharacters *int32
    // Minimum # of symbols required in work profile password. Valid values 1 to 10
    workProfilePasswordMinSymbolCharacters *int32
    // Minimum # of upper-case characters required in work profile password. Valid values 1 to 10
    workProfilePasswordMinUpperCaseCharacters *int32
    // Minutes of inactivity before the screen times out.
    workProfilePasswordMinutesOfInactivityBeforeScreenTimeout *int32
    // Number of previous work profile passwords to block. Valid values 0 to 24
    workProfilePasswordPreviousPasswordBlockCount *int32
    // Type of work profile password that is required. Possible values are: deviceDefault, lowSecurityBiometric, required, atLeastNumeric, numericComplex, atLeastAlphabetic, atLeastAlphanumeric, alphanumericWithSymbols.
    workProfilePasswordRequiredType *AndroidForWorkRequiredPasswordType
    // Number of sign in failures allowed before work profile is removed and all corporate data deleted. Valid values 1 to 16
    workProfilePasswordSignInFailureCountBeforeFactoryReset *int32
    // Password is required or not for work profile
    workProfileRequirePassword *bool
}
// NewAndroidForWorkGeneralDeviceConfiguration instantiates a new AndroidForWorkGeneralDeviceConfiguration and sets the default values.
func NewAndroidForWorkGeneralDeviceConfiguration()(*AndroidForWorkGeneralDeviceConfiguration) {
    m := &AndroidForWorkGeneralDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    return m
}
// CreateAndroidForWorkGeneralDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkGeneralDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkGeneralDeviceConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkGeneralDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["passwordBlockFaceUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockFaceUnlock(val)
        }
        return nil
    }
    res["passwordBlockFingerprintUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockFingerprintUnlock(val)
        }
        return nil
    }
    res["passwordBlockIrisUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockIrisUnlock(val)
        }
        return nil
    }
    res["passwordBlockTrustAgents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockTrustAgents(val)
        }
        return nil
    }
    res["passwordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordExpirationDays(val)
        }
        return nil
    }
    res["passwordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLength(val)
        }
        return nil
    }
    res["passwordMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeScreenTimeout(val)
        }
        return nil
    }
    res["passwordPreviousPasswordBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPreviousPasswordBlockCount(val)
        }
        return nil
    }
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*AndroidForWorkRequiredPasswordType))
        }
        return nil
    }
    res["passwordSignInFailureCountBeforeFactoryReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordSignInFailureCountBeforeFactoryReset(val)
        }
        return nil
    }
    res["securityRequireVerifyApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireVerifyApps(val)
        }
        return nil
    }
    res["vpnAlwaysOnPackageIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnAlwaysOnPackageIdentifier(val)
        }
        return nil
    }
    res["vpnEnableAlwaysOnLockdownMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnEnableAlwaysOnLockdownMode(val)
        }
        return nil
    }
    res["workProfileAllowWidgets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileAllowWidgets(val)
        }
        return nil
    }
    res["workProfileBlockAddingAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockAddingAccounts(val)
        }
        return nil
    }
    res["workProfileBlockCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockCamera(val)
        }
        return nil
    }
    res["workProfileBlockCrossProfileCallerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockCrossProfileCallerId(val)
        }
        return nil
    }
    res["workProfileBlockCrossProfileContactsSearch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockCrossProfileContactsSearch(val)
        }
        return nil
    }
    res["workProfileBlockCrossProfileCopyPaste"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockCrossProfileCopyPaste(val)
        }
        return nil
    }
    res["workProfileBlockNotificationsWhileDeviceLocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockNotificationsWhileDeviceLocked(val)
        }
        return nil
    }
    res["workProfileBlockPersonalAppInstallsFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockPersonalAppInstallsFromUnknownSources(val)
        }
        return nil
    }
    res["workProfileBlockScreenCapture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockScreenCapture(val)
        }
        return nil
    }
    res["workProfileBluetoothEnableContactSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBluetoothEnableContactSharing(val)
        }
        return nil
    }
    res["workProfileDataSharingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkCrossProfileDataSharingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileDataSharingType(val.(*AndroidForWorkCrossProfileDataSharingType))
        }
        return nil
    }
    res["workProfileDefaultAppPermissionPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkDefaultAppPermissionPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileDefaultAppPermissionPolicy(val.(*AndroidForWorkDefaultAppPermissionPolicyType))
        }
        return nil
    }
    res["workProfilePasswordBlockFaceUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordBlockFaceUnlock(val)
        }
        return nil
    }
    res["workProfilePasswordBlockFingerprintUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordBlockFingerprintUnlock(val)
        }
        return nil
    }
    res["workProfilePasswordBlockIrisUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordBlockIrisUnlock(val)
        }
        return nil
    }
    res["workProfilePasswordBlockTrustAgents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordBlockTrustAgents(val)
        }
        return nil
    }
    res["workProfilePasswordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordExpirationDays(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumLength(val)
        }
        return nil
    }
    res["workProfilePasswordMinLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinLetterCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinLowerCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinLowerCaseCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinNonLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinNonLetterCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinNumericCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinNumericCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinSymbolCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinSymbolCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinUpperCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinUpperCaseCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout(val)
        }
        return nil
    }
    res["workProfilePasswordPreviousPasswordBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordPreviousPasswordBlockCount(val)
        }
        return nil
    }
    res["workProfilePasswordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordRequiredType(val.(*AndroidForWorkRequiredPasswordType))
        }
        return nil
    }
    res["workProfilePasswordSignInFailureCountBeforeFactoryReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(val)
        }
        return nil
    }
    res["workProfileRequirePassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileRequirePassword(val)
        }
        return nil
    }
    return res
}
// GetPasswordBlockFaceUnlock gets the passwordBlockFaceUnlock property value. Indicates whether or not to block face unlock.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetPasswordBlockFaceUnlock()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordBlockFaceUnlock
    }
}
// GetPasswordBlockFingerprintUnlock gets the passwordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlock()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordBlockFingerprintUnlock
    }
}
// GetPasswordBlockIrisUnlock gets the passwordBlockIrisUnlock property value. Indicates whether or not to block iris unlock.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetPasswordBlockIrisUnlock()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordBlockIrisUnlock
    }
}
// GetPasswordBlockTrustAgents gets the passwordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetPasswordBlockTrustAgents()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordBlockTrustAgents
    }
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidForWorkGeneralDeviceConfiguration) GetPasswordExpirationDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordExpirationDays
    }
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum length of passwords. Valid values 4 to 16
func (m *AndroidForWorkGeneralDeviceConfiguration) GetPasswordMinimumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMinimumLength
    }
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMinutesOfInactivityBeforeScreenTimeout
    }
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 0 to 24
func (m *AndroidForWorkGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordPreviousPasswordBlockCount
    }
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Type of password that is required. Possible values are: deviceDefault, lowSecurityBiometric, required, atLeastNumeric, numericComplex, atLeastAlphabetic, atLeastAlphanumeric, alphanumericWithSymbols.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetPasswordRequiredType()(*AndroidForWorkRequiredPasswordType) {
    if m == nil {
        return nil
    } else {
        return m.passwordRequiredType
    }
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before factory reset. Valid values 1 to 16
func (m *AndroidForWorkGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordSignInFailureCountBeforeFactoryReset
    }
}
// GetSecurityRequireVerifyApps gets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetSecurityRequireVerifyApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityRequireVerifyApps
    }
}
// GetVpnAlwaysOnPackageIdentifier gets the vpnAlwaysOnPackageIdentifier property value. Enable lockdown mode for always-on VPN.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetVpnAlwaysOnPackageIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vpnAlwaysOnPackageIdentifier
    }
}
// GetVpnEnableAlwaysOnLockdownMode gets the vpnEnableAlwaysOnLockdownMode property value. Enable lockdown mode for always-on VPN.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetVpnEnableAlwaysOnLockdownMode()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.vpnEnableAlwaysOnLockdownMode
    }
}
// GetWorkProfileAllowWidgets gets the workProfileAllowWidgets property value. Allow widgets from work profile apps.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileAllowWidgets()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfileAllowWidgets
    }
}
// GetWorkProfileBlockAddingAccounts gets the workProfileBlockAddingAccounts property value. Block users from adding/removing accounts in work profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileBlockAddingAccounts()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfileBlockAddingAccounts
    }
}
// GetWorkProfileBlockCamera gets the workProfileBlockCamera property value. Block work profile camera.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileBlockCamera()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfileBlockCamera
    }
}
// GetWorkProfileBlockCrossProfileCallerId gets the workProfileBlockCrossProfileCallerId property value. Block display work profile caller ID in personal profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCallerId()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfileBlockCrossProfileCallerId
    }
}
// GetWorkProfileBlockCrossProfileContactsSearch gets the workProfileBlockCrossProfileContactsSearch property value. Block work profile contacts availability in personal profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileContactsSearch()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfileBlockCrossProfileContactsSearch
    }
}
// GetWorkProfileBlockCrossProfileCopyPaste gets the workProfileBlockCrossProfileCopyPaste property value. Boolean that indicates if the setting disallow cross profile copy/paste is enabled.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCopyPaste()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfileBlockCrossProfileCopyPaste
    }
}
// GetWorkProfileBlockNotificationsWhileDeviceLocked gets the workProfileBlockNotificationsWhileDeviceLocked property value. Indicates whether or not to block notifications while device locked.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileBlockNotificationsWhileDeviceLocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfileBlockNotificationsWhileDeviceLocked
    }
}
// GetWorkProfileBlockPersonalAppInstallsFromUnknownSources gets the workProfileBlockPersonalAppInstallsFromUnknownSources property value. Prevent app installations from unknown sources in the personal profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileBlockPersonalAppInstallsFromUnknownSources()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfileBlockPersonalAppInstallsFromUnknownSources
    }
}
// GetWorkProfileBlockScreenCapture gets the workProfileBlockScreenCapture property value. Block screen capture in work profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileBlockScreenCapture()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfileBlockScreenCapture
    }
}
// GetWorkProfileBluetoothEnableContactSharing gets the workProfileBluetoothEnableContactSharing property value. Allow bluetooth devices to access enterprise contacts.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileBluetoothEnableContactSharing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfileBluetoothEnableContactSharing
    }
}
// GetWorkProfileDataSharingType gets the workProfileDataSharingType property value. Type of data sharing that is allowed. Possible values are: deviceDefault, preventAny, allowPersonalToWork, noRestrictions.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileDataSharingType()(*AndroidForWorkCrossProfileDataSharingType) {
    if m == nil {
        return nil
    } else {
        return m.workProfileDataSharingType
    }
}
// GetWorkProfileDefaultAppPermissionPolicy gets the workProfileDefaultAppPermissionPolicy property value. Type of password that is required. Possible values are: deviceDefault, prompt, autoGrant, autoDeny.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileDefaultAppPermissionPolicy()(*AndroidForWorkDefaultAppPermissionPolicyType) {
    if m == nil {
        return nil
    } else {
        return m.workProfileDefaultAppPermissionPolicy
    }
}
// GetWorkProfilePasswordBlockFaceUnlock gets the workProfilePasswordBlockFaceUnlock property value. Indicates whether or not to block face unlock for work profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordBlockFaceUnlock()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordBlockFaceUnlock
    }
}
// GetWorkProfilePasswordBlockFingerprintUnlock gets the workProfilePasswordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock for work profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordBlockFingerprintUnlock()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordBlockFingerprintUnlock
    }
}
// GetWorkProfilePasswordBlockIrisUnlock gets the workProfilePasswordBlockIrisUnlock property value. Indicates whether or not to block iris unlock for work profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordBlockIrisUnlock()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordBlockIrisUnlock
    }
}
// GetWorkProfilePasswordBlockTrustAgents gets the workProfilePasswordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents for work profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordBlockTrustAgents()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordBlockTrustAgents
    }
}
// GetWorkProfilePasswordExpirationDays gets the workProfilePasswordExpirationDays property value. Number of days before the work profile password expires. Valid values 1 to 365
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordExpirationDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordExpirationDays
    }
}
// GetWorkProfilePasswordMinimumLength gets the workProfilePasswordMinimumLength property value. Minimum length of work profile password. Valid values 4 to 16
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordMinimumLength
    }
}
// GetWorkProfilePasswordMinLetterCharacters gets the workProfilePasswordMinLetterCharacters property value. Minimum # of letter characters required in work profile password. Valid values 1 to 10
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordMinLetterCharacters()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordMinLetterCharacters
    }
}
// GetWorkProfilePasswordMinLowerCaseCharacters gets the workProfilePasswordMinLowerCaseCharacters property value. Minimum # of lower-case characters required in work profile password. Valid values 1 to 10
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordMinLowerCaseCharacters()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordMinLowerCaseCharacters
    }
}
// GetWorkProfilePasswordMinNonLetterCharacters gets the workProfilePasswordMinNonLetterCharacters property value. Minimum # of non-letter characters required in work profile password. Valid values 1 to 10
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordMinNonLetterCharacters()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordMinNonLetterCharacters
    }
}
// GetWorkProfilePasswordMinNumericCharacters gets the workProfilePasswordMinNumericCharacters property value. Minimum # of numeric characters required in work profile password. Valid values 1 to 10
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordMinNumericCharacters()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordMinNumericCharacters
    }
}
// GetWorkProfilePasswordMinSymbolCharacters gets the workProfilePasswordMinSymbolCharacters property value. Minimum # of symbols required in work profile password. Valid values 1 to 10
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordMinSymbolCharacters()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordMinSymbolCharacters
    }
}
// GetWorkProfilePasswordMinUpperCaseCharacters gets the workProfilePasswordMinUpperCaseCharacters property value. Minimum # of upper-case characters required in work profile password. Valid values 1 to 10
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordMinUpperCaseCharacters()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordMinUpperCaseCharacters
    }
}
// GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout gets the workProfilePasswordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordMinutesOfInactivityBeforeScreenTimeout
    }
}
// GetWorkProfilePasswordPreviousPasswordBlockCount gets the workProfilePasswordPreviousPasswordBlockCount property value. Number of previous work profile passwords to block. Valid values 0 to 24
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordPreviousPasswordBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordPreviousPasswordBlockCount
    }
}
// GetWorkProfilePasswordRequiredType gets the workProfilePasswordRequiredType property value. Type of work profile password that is required. Possible values are: deviceDefault, lowSecurityBiometric, required, atLeastNumeric, numericComplex, atLeastAlphabetic, atLeastAlphanumeric, alphanumericWithSymbols.
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordRequiredType()(*AndroidForWorkRequiredPasswordType) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordRequiredType
    }
}
// GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset gets the workProfilePasswordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before work profile is removed and all corporate data deleted. Valid values 1 to 16
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.workProfilePasswordSignInFailureCountBeforeFactoryReset
    }
}
// GetWorkProfileRequirePassword gets the workProfileRequirePassword property value. Password is required or not for work profile
func (m *AndroidForWorkGeneralDeviceConfiguration) GetWorkProfileRequirePassword()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.workProfileRequirePassword
    }
}
// Serialize serializes information the current object
func (m *AndroidForWorkGeneralDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("passwordBlockFaceUnlock", m.GetPasswordBlockFaceUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockFingerprintUnlock", m.GetPasswordBlockFingerprintUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockIrisUnlock", m.GetPasswordBlockIrisUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockTrustAgents", m.GetPasswordBlockTrustAgents())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationDays", m.GetPasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeScreenTimeout", m.GetPasswordMinutesOfInactivityBeforeScreenTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordPreviousPasswordBlockCount", m.GetPasswordPreviousPasswordBlockCount())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRequiredType() != nil {
        cast := (*m.GetPasswordRequiredType()).String()
        err = writer.WriteStringValue("passwordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordSignInFailureCountBeforeFactoryReset", m.GetPasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireVerifyApps", m.GetSecurityRequireVerifyApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vpnAlwaysOnPackageIdentifier", m.GetVpnAlwaysOnPackageIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("vpnEnableAlwaysOnLockdownMode", m.GetVpnEnableAlwaysOnLockdownMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileAllowWidgets", m.GetWorkProfileAllowWidgets())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockAddingAccounts", m.GetWorkProfileBlockAddingAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockCamera", m.GetWorkProfileBlockCamera())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockCrossProfileCallerId", m.GetWorkProfileBlockCrossProfileCallerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockCrossProfileContactsSearch", m.GetWorkProfileBlockCrossProfileContactsSearch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockCrossProfileCopyPaste", m.GetWorkProfileBlockCrossProfileCopyPaste())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockNotificationsWhileDeviceLocked", m.GetWorkProfileBlockNotificationsWhileDeviceLocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockPersonalAppInstallsFromUnknownSources", m.GetWorkProfileBlockPersonalAppInstallsFromUnknownSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockScreenCapture", m.GetWorkProfileBlockScreenCapture())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBluetoothEnableContactSharing", m.GetWorkProfileBluetoothEnableContactSharing())
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfileDataSharingType() != nil {
        cast := (*m.GetWorkProfileDataSharingType()).String()
        err = writer.WriteStringValue("workProfileDataSharingType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfileDefaultAppPermissionPolicy() != nil {
        cast := (*m.GetWorkProfileDefaultAppPermissionPolicy()).String()
        err = writer.WriteStringValue("workProfileDefaultAppPermissionPolicy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfilePasswordBlockFaceUnlock", m.GetWorkProfilePasswordBlockFaceUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfilePasswordBlockFingerprintUnlock", m.GetWorkProfilePasswordBlockFingerprintUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfilePasswordBlockIrisUnlock", m.GetWorkProfilePasswordBlockIrisUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfilePasswordBlockTrustAgents", m.GetWorkProfilePasswordBlockTrustAgents())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordExpirationDays", m.GetWorkProfilePasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumLength", m.GetWorkProfilePasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinLetterCharacters", m.GetWorkProfilePasswordMinLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinLowerCaseCharacters", m.GetWorkProfilePasswordMinLowerCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinNonLetterCharacters", m.GetWorkProfilePasswordMinNonLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinNumericCharacters", m.GetWorkProfilePasswordMinNumericCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinSymbolCharacters", m.GetWorkProfilePasswordMinSymbolCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinUpperCaseCharacters", m.GetWorkProfilePasswordMinUpperCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinutesOfInactivityBeforeScreenTimeout", m.GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordPreviousPasswordBlockCount", m.GetWorkProfilePasswordPreviousPasswordBlockCount())
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfilePasswordRequiredType() != nil {
        cast := (*m.GetWorkProfilePasswordRequiredType()).String()
        err = writer.WriteStringValue("workProfilePasswordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordSignInFailureCountBeforeFactoryReset", m.GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileRequirePassword", m.GetWorkProfileRequirePassword())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPasswordBlockFaceUnlock sets the passwordBlockFaceUnlock property value. Indicates whether or not to block face unlock.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetPasswordBlockFaceUnlock(value *bool)() {
    if m != nil {
        m.passwordBlockFaceUnlock = value
    }
}
// SetPasswordBlockFingerprintUnlock sets the passwordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetPasswordBlockFingerprintUnlock(value *bool)() {
    if m != nil {
        m.passwordBlockFingerprintUnlock = value
    }
}
// SetPasswordBlockIrisUnlock sets the passwordBlockIrisUnlock property value. Indicates whether or not to block iris unlock.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetPasswordBlockIrisUnlock(value *bool)() {
    if m != nil {
        m.passwordBlockIrisUnlock = value
    }
}
// SetPasswordBlockTrustAgents sets the passwordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetPasswordBlockTrustAgents(value *bool)() {
    if m != nil {
        m.passwordBlockTrustAgents = value
    }
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidForWorkGeneralDeviceConfiguration) SetPasswordExpirationDays(value *int32)() {
    if m != nil {
        m.passwordExpirationDays = value
    }
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum length of passwords. Valid values 4 to 16
func (m *AndroidForWorkGeneralDeviceConfiguration) SetPasswordMinimumLength(value *int32)() {
    if m != nil {
        m.passwordMinimumLength = value
    }
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    if m != nil {
        m.passwordMinutesOfInactivityBeforeScreenTimeout = value
    }
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 0 to 24
func (m *AndroidForWorkGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    if m != nil {
        m.passwordPreviousPasswordBlockCount = value
    }
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Type of password that is required. Possible values are: deviceDefault, lowSecurityBiometric, required, atLeastNumeric, numericComplex, atLeastAlphabetic, atLeastAlphanumeric, alphanumericWithSymbols.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetPasswordRequiredType(value *AndroidForWorkRequiredPasswordType)() {
    if m != nil {
        m.passwordRequiredType = value
    }
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before factory reset. Valid values 1 to 16
func (m *AndroidForWorkGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    if m != nil {
        m.passwordSignInFailureCountBeforeFactoryReset = value
    }
}
// SetSecurityRequireVerifyApps sets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetSecurityRequireVerifyApps(value *bool)() {
    if m != nil {
        m.securityRequireVerifyApps = value
    }
}
// SetVpnAlwaysOnPackageIdentifier sets the vpnAlwaysOnPackageIdentifier property value. Enable lockdown mode for always-on VPN.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetVpnAlwaysOnPackageIdentifier(value *string)() {
    if m != nil {
        m.vpnAlwaysOnPackageIdentifier = value
    }
}
// SetVpnEnableAlwaysOnLockdownMode sets the vpnEnableAlwaysOnLockdownMode property value. Enable lockdown mode for always-on VPN.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetVpnEnableAlwaysOnLockdownMode(value *bool)() {
    if m != nil {
        m.vpnEnableAlwaysOnLockdownMode = value
    }
}
// SetWorkProfileAllowWidgets sets the workProfileAllowWidgets property value. Allow widgets from work profile apps.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileAllowWidgets(value *bool)() {
    if m != nil {
        m.workProfileAllowWidgets = value
    }
}
// SetWorkProfileBlockAddingAccounts sets the workProfileBlockAddingAccounts property value. Block users from adding/removing accounts in work profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileBlockAddingAccounts(value *bool)() {
    if m != nil {
        m.workProfileBlockAddingAccounts = value
    }
}
// SetWorkProfileBlockCamera sets the workProfileBlockCamera property value. Block work profile camera.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileBlockCamera(value *bool)() {
    if m != nil {
        m.workProfileBlockCamera = value
    }
}
// SetWorkProfileBlockCrossProfileCallerId sets the workProfileBlockCrossProfileCallerId property value. Block display work profile caller ID in personal profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileCallerId(value *bool)() {
    if m != nil {
        m.workProfileBlockCrossProfileCallerId = value
    }
}
// SetWorkProfileBlockCrossProfileContactsSearch sets the workProfileBlockCrossProfileContactsSearch property value. Block work profile contacts availability in personal profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileContactsSearch(value *bool)() {
    if m != nil {
        m.workProfileBlockCrossProfileContactsSearch = value
    }
}
// SetWorkProfileBlockCrossProfileCopyPaste sets the workProfileBlockCrossProfileCopyPaste property value. Boolean that indicates if the setting disallow cross profile copy/paste is enabled.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileCopyPaste(value *bool)() {
    if m != nil {
        m.workProfileBlockCrossProfileCopyPaste = value
    }
}
// SetWorkProfileBlockNotificationsWhileDeviceLocked sets the workProfileBlockNotificationsWhileDeviceLocked property value. Indicates whether or not to block notifications while device locked.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileBlockNotificationsWhileDeviceLocked(value *bool)() {
    if m != nil {
        m.workProfileBlockNotificationsWhileDeviceLocked = value
    }
}
// SetWorkProfileBlockPersonalAppInstallsFromUnknownSources sets the workProfileBlockPersonalAppInstallsFromUnknownSources property value. Prevent app installations from unknown sources in the personal profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileBlockPersonalAppInstallsFromUnknownSources(value *bool)() {
    if m != nil {
        m.workProfileBlockPersonalAppInstallsFromUnknownSources = value
    }
}
// SetWorkProfileBlockScreenCapture sets the workProfileBlockScreenCapture property value. Block screen capture in work profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileBlockScreenCapture(value *bool)() {
    if m != nil {
        m.workProfileBlockScreenCapture = value
    }
}
// SetWorkProfileBluetoothEnableContactSharing sets the workProfileBluetoothEnableContactSharing property value. Allow bluetooth devices to access enterprise contacts.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileBluetoothEnableContactSharing(value *bool)() {
    if m != nil {
        m.workProfileBluetoothEnableContactSharing = value
    }
}
// SetWorkProfileDataSharingType sets the workProfileDataSharingType property value. Type of data sharing that is allowed. Possible values are: deviceDefault, preventAny, allowPersonalToWork, noRestrictions.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileDataSharingType(value *AndroidForWorkCrossProfileDataSharingType)() {
    if m != nil {
        m.workProfileDataSharingType = value
    }
}
// SetWorkProfileDefaultAppPermissionPolicy sets the workProfileDefaultAppPermissionPolicy property value. Type of password that is required. Possible values are: deviceDefault, prompt, autoGrant, autoDeny.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileDefaultAppPermissionPolicy(value *AndroidForWorkDefaultAppPermissionPolicyType)() {
    if m != nil {
        m.workProfileDefaultAppPermissionPolicy = value
    }
}
// SetWorkProfilePasswordBlockFaceUnlock sets the workProfilePasswordBlockFaceUnlock property value. Indicates whether or not to block face unlock for work profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordBlockFaceUnlock(value *bool)() {
    if m != nil {
        m.workProfilePasswordBlockFaceUnlock = value
    }
}
// SetWorkProfilePasswordBlockFingerprintUnlock sets the workProfilePasswordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock for work profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordBlockFingerprintUnlock(value *bool)() {
    if m != nil {
        m.workProfilePasswordBlockFingerprintUnlock = value
    }
}
// SetWorkProfilePasswordBlockIrisUnlock sets the workProfilePasswordBlockIrisUnlock property value. Indicates whether or not to block iris unlock for work profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordBlockIrisUnlock(value *bool)() {
    if m != nil {
        m.workProfilePasswordBlockIrisUnlock = value
    }
}
// SetWorkProfilePasswordBlockTrustAgents sets the workProfilePasswordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents for work profile.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordBlockTrustAgents(value *bool)() {
    if m != nil {
        m.workProfilePasswordBlockTrustAgents = value
    }
}
// SetWorkProfilePasswordExpirationDays sets the workProfilePasswordExpirationDays property value. Number of days before the work profile password expires. Valid values 1 to 365
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordExpirationDays(value *int32)() {
    if m != nil {
        m.workProfilePasswordExpirationDays = value
    }
}
// SetWorkProfilePasswordMinimumLength sets the workProfilePasswordMinimumLength property value. Minimum length of work profile password. Valid values 4 to 16
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLength(value *int32)() {
    if m != nil {
        m.workProfilePasswordMinimumLength = value
    }
}
// SetWorkProfilePasswordMinLetterCharacters sets the workProfilePasswordMinLetterCharacters property value. Minimum # of letter characters required in work profile password. Valid values 1 to 10
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordMinLetterCharacters(value *int32)() {
    if m != nil {
        m.workProfilePasswordMinLetterCharacters = value
    }
}
// SetWorkProfilePasswordMinLowerCaseCharacters sets the workProfilePasswordMinLowerCaseCharacters property value. Minimum # of lower-case characters required in work profile password. Valid values 1 to 10
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordMinLowerCaseCharacters(value *int32)() {
    if m != nil {
        m.workProfilePasswordMinLowerCaseCharacters = value
    }
}
// SetWorkProfilePasswordMinNonLetterCharacters sets the workProfilePasswordMinNonLetterCharacters property value. Minimum # of non-letter characters required in work profile password. Valid values 1 to 10
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordMinNonLetterCharacters(value *int32)() {
    if m != nil {
        m.workProfilePasswordMinNonLetterCharacters = value
    }
}
// SetWorkProfilePasswordMinNumericCharacters sets the workProfilePasswordMinNumericCharacters property value. Minimum # of numeric characters required in work profile password. Valid values 1 to 10
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordMinNumericCharacters(value *int32)() {
    if m != nil {
        m.workProfilePasswordMinNumericCharacters = value
    }
}
// SetWorkProfilePasswordMinSymbolCharacters sets the workProfilePasswordMinSymbolCharacters property value. Minimum # of symbols required in work profile password. Valid values 1 to 10
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordMinSymbolCharacters(value *int32)() {
    if m != nil {
        m.workProfilePasswordMinSymbolCharacters = value
    }
}
// SetWorkProfilePasswordMinUpperCaseCharacters sets the workProfilePasswordMinUpperCaseCharacters property value. Minimum # of upper-case characters required in work profile password. Valid values 1 to 10
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordMinUpperCaseCharacters(value *int32)() {
    if m != nil {
        m.workProfilePasswordMinUpperCaseCharacters = value
    }
}
// SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout sets the workProfilePasswordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    if m != nil {
        m.workProfilePasswordMinutesOfInactivityBeforeScreenTimeout = value
    }
}
// SetWorkProfilePasswordPreviousPasswordBlockCount sets the workProfilePasswordPreviousPasswordBlockCount property value. Number of previous work profile passwords to block. Valid values 0 to 24
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordPreviousPasswordBlockCount(value *int32)() {
    if m != nil {
        m.workProfilePasswordPreviousPasswordBlockCount = value
    }
}
// SetWorkProfilePasswordRequiredType sets the workProfilePasswordRequiredType property value. Type of work profile password that is required. Possible values are: deviceDefault, lowSecurityBiometric, required, atLeastNumeric, numericComplex, atLeastAlphabetic, atLeastAlphanumeric, alphanumericWithSymbols.
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordRequiredType(value *AndroidForWorkRequiredPasswordType)() {
    if m != nil {
        m.workProfilePasswordRequiredType = value
    }
}
// SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset sets the workProfilePasswordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before work profile is removed and all corporate data deleted. Valid values 1 to 16
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    if m != nil {
        m.workProfilePasswordSignInFailureCountBeforeFactoryReset = value
    }
}
// SetWorkProfileRequirePassword sets the workProfileRequirePassword property value. Password is required or not for work profile
func (m *AndroidForWorkGeneralDeviceConfiguration) SetWorkProfileRequirePassword(value *bool)() {
    if m != nil {
        m.workProfileRequirePassword = value
    }
}
