package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkCompliancePolicy 
type AndroidForWorkCompliancePolicy struct {
    DeviceCompliancePolicy
    // Require that devices have enabled device threat protection.
    deviceThreatProtectionEnabled *bool
    // Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
    deviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel
    // Minimum Android security patch level.
    minAndroidSecurityPatchLevel *string
    // Maximum Android version.
    osMaximumVersion *string
    // Minimum Android version.
    osMinimumVersion *string
    // Number of days before the password expires. Valid values 1 to 365
    passwordExpirationDays *int32
    // Minimum password length. Valid values 4 to 16
    passwordMinimumLength *int32
    // Minutes of inactivity before a password is required.
    passwordMinutesOfInactivityBeforeLock *int32
    // Number of previous passwords to block. Valid values 1 to 24
    passwordPreviousPasswordBlockCount *int32
    // Require a password to unlock device.
    passwordRequired *bool
    // Type of characters in password. Possible values are: deviceDefault, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, numeric, numericComplex, any.
    passwordRequiredType *AndroidRequiredPasswordType
    // Number of sign-in failures allowed before factory reset. Valid values 1 to 16
    passwordSignInFailureCountBeforeFactoryReset *int32
    // Devices must not be jailbroken or rooted.
    securityBlockJailbrokenDevices *bool
    // Disable USB debugging on Android devices.
    securityDisableUsbDebugging *bool
    // Require that devices disallow installation of apps from unknown sources.
    securityPreventInstallAppsFromUnknownSources *bool
    // Require the device to pass the Company Portal client app runtime integrity check.
    securityRequireCompanyPortalAppIntegrity *bool
    // Require a specific SafetyNet evaluation type for compliance. Possible values are: basic, hardwareBacked.
    securityRequiredAndroidSafetyNetEvaluationType *AndroidSafetyNetEvaluationType
    // Require Google Play Services to be installed and enabled on the device.
    securityRequireGooglePlayServices *bool
    // Require the device to pass the SafetyNet basic integrity check.
    securityRequireSafetyNetAttestationBasicIntegrity *bool
    // Require the device to pass the SafetyNet certified device check.
    securityRequireSafetyNetAttestationCertifiedDevice *bool
    // Require the device to have up to date security providers. The device will require Google Play Services to be enabled and up to date.
    securityRequireUpToDateSecurityProviders *bool
    // Require the Android Verify apps feature is turned on.
    securityRequireVerifyApps *bool
    // Require encryption on Android devices.
    storageRequireEncryption *bool
}
// NewAndroidForWorkCompliancePolicy instantiates a new AndroidForWorkCompliancePolicy and sets the default values.
func NewAndroidForWorkCompliancePolicy()(*AndroidForWorkCompliancePolicy) {
    m := &AndroidForWorkCompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    return m
}
// CreateAndroidForWorkCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkCompliancePolicy(), nil
}
// GetDeviceThreatProtectionEnabled gets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *AndroidForWorkCompliancePolicy) GetDeviceThreatProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceThreatProtectionEnabled
    }
}
// GetDeviceThreatProtectionRequiredSecurityLevel gets the deviceThreatProtectionRequiredSecurityLevel property value. Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *AndroidForWorkCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    if m == nil {
        return nil
    } else {
        return m.deviceThreatProtectionRequiredSecurityLevel
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceCompliancePolicy.GetFieldDeserializers()
    res["deviceThreatProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceThreatProtectionEnabled(val)
        }
        return nil
    }
    res["deviceThreatProtectionRequiredSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceThreatProtectionLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceThreatProtectionRequiredSecurityLevel(val.(*DeviceThreatProtectionLevel))
        }
        return nil
    }
    res["minAndroidSecurityPatchLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinAndroidSecurityPatchLevel(val)
        }
        return nil
    }
    res["osMaximumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMaximumVersion(val)
        }
        return nil
    }
    res["osMinimumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMinimumVersion(val)
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
    res["passwordMinutesOfInactivityBeforeLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeLock(val)
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
    res["passwordRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequired(val)
        }
        return nil
    }
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*AndroidRequiredPasswordType))
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
    res["securityBlockJailbrokenDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityBlockJailbrokenDevices(val)
        }
        return nil
    }
    res["securityDisableUsbDebugging"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityDisableUsbDebugging(val)
        }
        return nil
    }
    res["securityPreventInstallAppsFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityPreventInstallAppsFromUnknownSources(val)
        }
        return nil
    }
    res["securityRequireCompanyPortalAppIntegrity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireCompanyPortalAppIntegrity(val)
        }
        return nil
    }
    res["securityRequiredAndroidSafetyNetEvaluationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidSafetyNetEvaluationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequiredAndroidSafetyNetEvaluationType(val.(*AndroidSafetyNetEvaluationType))
        }
        return nil
    }
    res["securityRequireGooglePlayServices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireGooglePlayServices(val)
        }
        return nil
    }
    res["securityRequireSafetyNetAttestationBasicIntegrity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireSafetyNetAttestationBasicIntegrity(val)
        }
        return nil
    }
    res["securityRequireSafetyNetAttestationCertifiedDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireSafetyNetAttestationCertifiedDevice(val)
        }
        return nil
    }
    res["securityRequireUpToDateSecurityProviders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireUpToDateSecurityProviders(val)
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
    res["storageRequireEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageRequireEncryption(val)
        }
        return nil
    }
    return res
}
// GetMinAndroidSecurityPatchLevel gets the minAndroidSecurityPatchLevel property value. Minimum Android security patch level.
func (m *AndroidForWorkCompliancePolicy) GetMinAndroidSecurityPatchLevel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minAndroidSecurityPatchLevel
    }
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Maximum Android version.
func (m *AndroidForWorkCompliancePolicy) GetOsMaximumVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osMaximumVersion
    }
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Minimum Android version.
func (m *AndroidForWorkCompliancePolicy) GetOsMinimumVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osMinimumVersion
    }
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidForWorkCompliancePolicy) GetPasswordExpirationDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordExpirationDays
    }
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AndroidForWorkCompliancePolicy) GetPasswordMinimumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMinimumLength
    }
}
// GetPasswordMinutesOfInactivityBeforeLock gets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *AndroidForWorkCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMinutesOfInactivityBeforeLock
    }
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 1 to 24
func (m *AndroidForWorkCompliancePolicy) GetPasswordPreviousPasswordBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordPreviousPasswordBlockCount
    }
}
// GetPasswordRequired gets the passwordRequired property value. Require a password to unlock device.
func (m *AndroidForWorkCompliancePolicy) GetPasswordRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordRequired
    }
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Type of characters in password. Possible values are: deviceDefault, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, numeric, numericComplex, any.
func (m *AndroidForWorkCompliancePolicy) GetPasswordRequiredType()(*AndroidRequiredPasswordType) {
    if m == nil {
        return nil
    } else {
        return m.passwordRequiredType
    }
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign-in failures allowed before factory reset. Valid values 1 to 16
func (m *AndroidForWorkCompliancePolicy) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordSignInFailureCountBeforeFactoryReset
    }
}
// GetSecurityBlockJailbrokenDevices gets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *AndroidForWorkCompliancePolicy) GetSecurityBlockJailbrokenDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityBlockJailbrokenDevices
    }
}
// GetSecurityDisableUsbDebugging gets the securityDisableUsbDebugging property value. Disable USB debugging on Android devices.
func (m *AndroidForWorkCompliancePolicy) GetSecurityDisableUsbDebugging()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityDisableUsbDebugging
    }
}
// GetSecurityPreventInstallAppsFromUnknownSources gets the securityPreventInstallAppsFromUnknownSources property value. Require that devices disallow installation of apps from unknown sources.
func (m *AndroidForWorkCompliancePolicy) GetSecurityPreventInstallAppsFromUnknownSources()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityPreventInstallAppsFromUnknownSources
    }
}
// GetSecurityRequireCompanyPortalAppIntegrity gets the securityRequireCompanyPortalAppIntegrity property value. Require the device to pass the Company Portal client app runtime integrity check.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequireCompanyPortalAppIntegrity()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityRequireCompanyPortalAppIntegrity
    }
}
// GetSecurityRequiredAndroidSafetyNetEvaluationType gets the securityRequiredAndroidSafetyNetEvaluationType property value. Require a specific SafetyNet evaluation type for compliance. Possible values are: basic, hardwareBacked.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequiredAndroidSafetyNetEvaluationType()(*AndroidSafetyNetEvaluationType) {
    if m == nil {
        return nil
    } else {
        return m.securityRequiredAndroidSafetyNetEvaluationType
    }
}
// GetSecurityRequireGooglePlayServices gets the securityRequireGooglePlayServices property value. Require Google Play Services to be installed and enabled on the device.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequireGooglePlayServices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityRequireGooglePlayServices
    }
}
// GetSecurityRequireSafetyNetAttestationBasicIntegrity gets the securityRequireSafetyNetAttestationBasicIntegrity property value. Require the device to pass the SafetyNet basic integrity check.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequireSafetyNetAttestationBasicIntegrity()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityRequireSafetyNetAttestationBasicIntegrity
    }
}
// GetSecurityRequireSafetyNetAttestationCertifiedDevice gets the securityRequireSafetyNetAttestationCertifiedDevice property value. Require the device to pass the SafetyNet certified device check.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequireSafetyNetAttestationCertifiedDevice()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityRequireSafetyNetAttestationCertifiedDevice
    }
}
// GetSecurityRequireUpToDateSecurityProviders gets the securityRequireUpToDateSecurityProviders property value. Require the device to have up to date security providers. The device will require Google Play Services to be enabled and up to date.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequireUpToDateSecurityProviders()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityRequireUpToDateSecurityProviders
    }
}
// GetSecurityRequireVerifyApps gets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequireVerifyApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityRequireVerifyApps
    }
}
// GetStorageRequireEncryption gets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AndroidForWorkCompliancePolicy) GetStorageRequireEncryption()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.storageRequireEncryption
    }
}
// Serialize serializes information the current object
func (m *AndroidForWorkCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceCompliancePolicy.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("deviceThreatProtectionEnabled", m.GetDeviceThreatProtectionEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceThreatProtectionRequiredSecurityLevel() != nil {
        cast := (*m.GetDeviceThreatProtectionRequiredSecurityLevel()).String()
        err = writer.WriteStringValue("deviceThreatProtectionRequiredSecurityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minAndroidSecurityPatchLevel", m.GetMinAndroidSecurityPatchLevel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMaximumVersion", m.GetOsMaximumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMinimumVersion", m.GetOsMinimumVersion())
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
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeLock", m.GetPasswordMinutesOfInactivityBeforeLock())
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
    {
        err = writer.WriteBoolValue("passwordRequired", m.GetPasswordRequired())
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
        err = writer.WriteBoolValue("securityBlockJailbrokenDevices", m.GetSecurityBlockJailbrokenDevices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityDisableUsbDebugging", m.GetSecurityDisableUsbDebugging())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityPreventInstallAppsFromUnknownSources", m.GetSecurityPreventInstallAppsFromUnknownSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireCompanyPortalAppIntegrity", m.GetSecurityRequireCompanyPortalAppIntegrity())
        if err != nil {
            return err
        }
    }
    if m.GetSecurityRequiredAndroidSafetyNetEvaluationType() != nil {
        cast := (*m.GetSecurityRequiredAndroidSafetyNetEvaluationType()).String()
        err = writer.WriteStringValue("securityRequiredAndroidSafetyNetEvaluationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireGooglePlayServices", m.GetSecurityRequireGooglePlayServices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireSafetyNetAttestationBasicIntegrity", m.GetSecurityRequireSafetyNetAttestationBasicIntegrity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireSafetyNetAttestationCertifiedDevice", m.GetSecurityRequireSafetyNetAttestationCertifiedDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireUpToDateSecurityProviders", m.GetSecurityRequireUpToDateSecurityProviders())
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
        err = writer.WriteBoolValue("storageRequireEncryption", m.GetStorageRequireEncryption())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceThreatProtectionEnabled sets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *AndroidForWorkCompliancePolicy) SetDeviceThreatProtectionEnabled(value *bool)() {
    if m != nil {
        m.deviceThreatProtectionEnabled = value
    }
}
// SetDeviceThreatProtectionRequiredSecurityLevel sets the deviceThreatProtectionRequiredSecurityLevel property value. Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *AndroidForWorkCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    if m != nil {
        m.deviceThreatProtectionRequiredSecurityLevel = value
    }
}
// SetMinAndroidSecurityPatchLevel sets the minAndroidSecurityPatchLevel property value. Minimum Android security patch level.
func (m *AndroidForWorkCompliancePolicy) SetMinAndroidSecurityPatchLevel(value *string)() {
    if m != nil {
        m.minAndroidSecurityPatchLevel = value
    }
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum Android version.
func (m *AndroidForWorkCompliancePolicy) SetOsMaximumVersion(value *string)() {
    if m != nil {
        m.osMaximumVersion = value
    }
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum Android version.
func (m *AndroidForWorkCompliancePolicy) SetOsMinimumVersion(value *string)() {
    if m != nil {
        m.osMinimumVersion = value
    }
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidForWorkCompliancePolicy) SetPasswordExpirationDays(value *int32)() {
    if m != nil {
        m.passwordExpirationDays = value
    }
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AndroidForWorkCompliancePolicy) SetPasswordMinimumLength(value *int32)() {
    if m != nil {
        m.passwordMinimumLength = value
    }
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *AndroidForWorkCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    if m != nil {
        m.passwordMinutesOfInactivityBeforeLock = value
    }
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 1 to 24
func (m *AndroidForWorkCompliancePolicy) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    if m != nil {
        m.passwordPreviousPasswordBlockCount = value
    }
}
// SetPasswordRequired sets the passwordRequired property value. Require a password to unlock device.
func (m *AndroidForWorkCompliancePolicy) SetPasswordRequired(value *bool)() {
    if m != nil {
        m.passwordRequired = value
    }
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Type of characters in password. Possible values are: deviceDefault, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, numeric, numericComplex, any.
func (m *AndroidForWorkCompliancePolicy) SetPasswordRequiredType(value *AndroidRequiredPasswordType)() {
    if m != nil {
        m.passwordRequiredType = value
    }
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign-in failures allowed before factory reset. Valid values 1 to 16
func (m *AndroidForWorkCompliancePolicy) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    if m != nil {
        m.passwordSignInFailureCountBeforeFactoryReset = value
    }
}
// SetSecurityBlockJailbrokenDevices sets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *AndroidForWorkCompliancePolicy) SetSecurityBlockJailbrokenDevices(value *bool)() {
    if m != nil {
        m.securityBlockJailbrokenDevices = value
    }
}
// SetSecurityDisableUsbDebugging sets the securityDisableUsbDebugging property value. Disable USB debugging on Android devices.
func (m *AndroidForWorkCompliancePolicy) SetSecurityDisableUsbDebugging(value *bool)() {
    if m != nil {
        m.securityDisableUsbDebugging = value
    }
}
// SetSecurityPreventInstallAppsFromUnknownSources sets the securityPreventInstallAppsFromUnknownSources property value. Require that devices disallow installation of apps from unknown sources.
func (m *AndroidForWorkCompliancePolicy) SetSecurityPreventInstallAppsFromUnknownSources(value *bool)() {
    if m != nil {
        m.securityPreventInstallAppsFromUnknownSources = value
    }
}
// SetSecurityRequireCompanyPortalAppIntegrity sets the securityRequireCompanyPortalAppIntegrity property value. Require the device to pass the Company Portal client app runtime integrity check.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequireCompanyPortalAppIntegrity(value *bool)() {
    if m != nil {
        m.securityRequireCompanyPortalAppIntegrity = value
    }
}
// SetSecurityRequiredAndroidSafetyNetEvaluationType sets the securityRequiredAndroidSafetyNetEvaluationType property value. Require a specific SafetyNet evaluation type for compliance. Possible values are: basic, hardwareBacked.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequiredAndroidSafetyNetEvaluationType(value *AndroidSafetyNetEvaluationType)() {
    if m != nil {
        m.securityRequiredAndroidSafetyNetEvaluationType = value
    }
}
// SetSecurityRequireGooglePlayServices sets the securityRequireGooglePlayServices property value. Require Google Play Services to be installed and enabled on the device.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequireGooglePlayServices(value *bool)() {
    if m != nil {
        m.securityRequireGooglePlayServices = value
    }
}
// SetSecurityRequireSafetyNetAttestationBasicIntegrity sets the securityRequireSafetyNetAttestationBasicIntegrity property value. Require the device to pass the SafetyNet basic integrity check.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequireSafetyNetAttestationBasicIntegrity(value *bool)() {
    if m != nil {
        m.securityRequireSafetyNetAttestationBasicIntegrity = value
    }
}
// SetSecurityRequireSafetyNetAttestationCertifiedDevice sets the securityRequireSafetyNetAttestationCertifiedDevice property value. Require the device to pass the SafetyNet certified device check.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequireSafetyNetAttestationCertifiedDevice(value *bool)() {
    if m != nil {
        m.securityRequireSafetyNetAttestationCertifiedDevice = value
    }
}
// SetSecurityRequireUpToDateSecurityProviders sets the securityRequireUpToDateSecurityProviders property value. Require the device to have up to date security providers. The device will require Google Play Services to be enabled and up to date.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequireUpToDateSecurityProviders(value *bool)() {
    if m != nil {
        m.securityRequireUpToDateSecurityProviders = value
    }
}
// SetSecurityRequireVerifyApps sets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequireVerifyApps(value *bool)() {
    if m != nil {
        m.securityRequireVerifyApps = value
    }
}
// SetStorageRequireEncryption sets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AndroidForWorkCompliancePolicy) SetStorageRequireEncryption(value *bool)() {
    if m != nil {
        m.storageRequireEncryption = value
    }
}
