package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkCompliancePolicy this class contains compliance settings for Android for Work.
type AndroidForWorkCompliancePolicy struct {
    DeviceCompliancePolicy
    // The OdataType property
    OdataType *string
}
// NewAndroidForWorkCompliancePolicy instantiates a new androidForWorkCompliancePolicy and sets the default values.
func NewAndroidForWorkCompliancePolicy()(*AndroidForWorkCompliancePolicy) {
    m := &AndroidForWorkCompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    odataTypeValue := "#microsoft.graph.androidForWorkCompliancePolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidForWorkCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkCompliancePolicy(), nil
}
// GetDeviceThreatProtectionEnabled gets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *AndroidForWorkCompliancePolicy) GetDeviceThreatProtectionEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("deviceThreatProtectionEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeviceThreatProtectionRequiredSecurityLevel gets the deviceThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *AndroidForWorkCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    val, err := m.GetBackingStore().Get("deviceThreatProtectionRequiredSecurityLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceThreatProtectionLevel)
    }
    return nil
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
    res["requiredPasswordComplexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidRequiredPasswordComplexity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredPasswordComplexity(val.(*AndroidRequiredPasswordComplexity))
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
    val, err := m.GetBackingStore().Get("minAndroidSecurityPatchLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Maximum Android version.
func (m *AndroidForWorkCompliancePolicy) GetOsMaximumVersion()(*string) {
    val, err := m.GetBackingStore().Get("osMaximumVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Minimum Android version.
func (m *AndroidForWorkCompliancePolicy) GetOsMinimumVersion()(*string) {
    val, err := m.GetBackingStore().Get("osMinimumVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidForWorkCompliancePolicy) GetPasswordExpirationDays()(*int32) {
    val, err := m.GetBackingStore().Get("passwordExpirationDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AndroidForWorkCompliancePolicy) GetPasswordMinimumLength()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinutesOfInactivityBeforeLock gets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *AndroidForWorkCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinutesOfInactivityBeforeLock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 1 to 24
func (m *AndroidForWorkCompliancePolicy) GetPasswordPreviousPasswordBlockCount()(*int32) {
    val, err := m.GetBackingStore().Get("passwordPreviousPasswordBlockCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordRequired gets the passwordRequired property value. Require a password to unlock device.
func (m *AndroidForWorkCompliancePolicy) GetPasswordRequired()(*bool) {
    val, err := m.GetBackingStore().Get("passwordRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Android required password type.
func (m *AndroidForWorkCompliancePolicy) GetPasswordRequiredType()(*AndroidRequiredPasswordType) {
    val, err := m.GetBackingStore().Get("passwordRequiredType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidRequiredPasswordType)
    }
    return nil
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign-in failures allowed before factory reset. Valid values 1 to 16
func (m *AndroidForWorkCompliancePolicy) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    val, err := m.GetBackingStore().Get("passwordSignInFailureCountBeforeFactoryReset")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRequiredPasswordComplexity gets the requiredPasswordComplexity property value. The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
func (m *AndroidForWorkCompliancePolicy) GetRequiredPasswordComplexity()(*AndroidRequiredPasswordComplexity) {
    val, err := m.GetBackingStore().Get("requiredPasswordComplexity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidRequiredPasswordComplexity)
    }
    return nil
}
// GetSecurityBlockJailbrokenDevices gets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *AndroidForWorkCompliancePolicy) GetSecurityBlockJailbrokenDevices()(*bool) {
    val, err := m.GetBackingStore().Get("securityBlockJailbrokenDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityDisableUsbDebugging gets the securityDisableUsbDebugging property value. Disable USB debugging on Android devices.
func (m *AndroidForWorkCompliancePolicy) GetSecurityDisableUsbDebugging()(*bool) {
    val, err := m.GetBackingStore().Get("securityDisableUsbDebugging")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityPreventInstallAppsFromUnknownSources gets the securityPreventInstallAppsFromUnknownSources property value. Require that devices disallow installation of apps from unknown sources.
func (m *AndroidForWorkCompliancePolicy) GetSecurityPreventInstallAppsFromUnknownSources()(*bool) {
    val, err := m.GetBackingStore().Get("securityPreventInstallAppsFromUnknownSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityRequireCompanyPortalAppIntegrity gets the securityRequireCompanyPortalAppIntegrity property value. Require the device to pass the Company Portal client app runtime integrity check.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequireCompanyPortalAppIntegrity()(*bool) {
    val, err := m.GetBackingStore().Get("securityRequireCompanyPortalAppIntegrity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityRequiredAndroidSafetyNetEvaluationType gets the securityRequiredAndroidSafetyNetEvaluationType property value. An enum representing the Android SafetyNet attestation evaluation types.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequiredAndroidSafetyNetEvaluationType()(*AndroidSafetyNetEvaluationType) {
    val, err := m.GetBackingStore().Get("securityRequiredAndroidSafetyNetEvaluationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidSafetyNetEvaluationType)
    }
    return nil
}
// GetSecurityRequireGooglePlayServices gets the securityRequireGooglePlayServices property value. Require Google Play Services to be installed and enabled on the device.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequireGooglePlayServices()(*bool) {
    val, err := m.GetBackingStore().Get("securityRequireGooglePlayServices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityRequireSafetyNetAttestationBasicIntegrity gets the securityRequireSafetyNetAttestationBasicIntegrity property value. Require the device to pass the SafetyNet basic integrity check.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequireSafetyNetAttestationBasicIntegrity()(*bool) {
    val, err := m.GetBackingStore().Get("securityRequireSafetyNetAttestationBasicIntegrity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityRequireSafetyNetAttestationCertifiedDevice gets the securityRequireSafetyNetAttestationCertifiedDevice property value. Require the device to pass the SafetyNet certified device check.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequireSafetyNetAttestationCertifiedDevice()(*bool) {
    val, err := m.GetBackingStore().Get("securityRequireSafetyNetAttestationCertifiedDevice")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityRequireUpToDateSecurityProviders gets the securityRequireUpToDateSecurityProviders property value. Require the device to have up to date security providers. The device will require Google Play Services to be enabled and up to date.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequireUpToDateSecurityProviders()(*bool) {
    val, err := m.GetBackingStore().Get("securityRequireUpToDateSecurityProviders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityRequireVerifyApps gets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidForWorkCompliancePolicy) GetSecurityRequireVerifyApps()(*bool) {
    val, err := m.GetBackingStore().Get("securityRequireVerifyApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStorageRequireEncryption gets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AndroidForWorkCompliancePolicy) GetStorageRequireEncryption()(*bool) {
    val, err := m.GetBackingStore().Get("storageRequireEncryption")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
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
    if m.GetRequiredPasswordComplexity() != nil {
        cast := (*m.GetRequiredPasswordComplexity()).String()
        err = writer.WriteStringValue("requiredPasswordComplexity", &cast)
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
    err := m.GetBackingStore().Set("deviceThreatProtectionEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceThreatProtectionRequiredSecurityLevel sets the deviceThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *AndroidForWorkCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    err := m.GetBackingStore().Set("deviceThreatProtectionRequiredSecurityLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetMinAndroidSecurityPatchLevel sets the minAndroidSecurityPatchLevel property value. Minimum Android security patch level.
func (m *AndroidForWorkCompliancePolicy) SetMinAndroidSecurityPatchLevel(value *string)() {
    err := m.GetBackingStore().Set("minAndroidSecurityPatchLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum Android version.
func (m *AndroidForWorkCompliancePolicy) SetOsMaximumVersion(value *string)() {
    err := m.GetBackingStore().Set("osMaximumVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum Android version.
func (m *AndroidForWorkCompliancePolicy) SetOsMinimumVersion(value *string)() {
    err := m.GetBackingStore().Set("osMinimumVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidForWorkCompliancePolicy) SetPasswordExpirationDays(value *int32)() {
    err := m.GetBackingStore().Set("passwordExpirationDays", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AndroidForWorkCompliancePolicy) SetPasswordMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *AndroidForWorkCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinutesOfInactivityBeforeLock", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 1 to 24
func (m *AndroidForWorkCompliancePolicy) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    err := m.GetBackingStore().Set("passwordPreviousPasswordBlockCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequired sets the passwordRequired property value. Require a password to unlock device.
func (m *AndroidForWorkCompliancePolicy) SetPasswordRequired(value *bool)() {
    err := m.GetBackingStore().Set("passwordRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Android required password type.
func (m *AndroidForWorkCompliancePolicy) SetPasswordRequiredType(value *AndroidRequiredPasswordType)() {
    err := m.GetBackingStore().Set("passwordRequiredType", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign-in failures allowed before factory reset. Valid values 1 to 16
func (m *AndroidForWorkCompliancePolicy) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    err := m.GetBackingStore().Set("passwordSignInFailureCountBeforeFactoryReset", value)
    if err != nil {
        panic(err)
    }
}
// SetRequiredPasswordComplexity sets the requiredPasswordComplexity property value. The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
func (m *AndroidForWorkCompliancePolicy) SetRequiredPasswordComplexity(value *AndroidRequiredPasswordComplexity)() {
    err := m.GetBackingStore().Set("requiredPasswordComplexity", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityBlockJailbrokenDevices sets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *AndroidForWorkCompliancePolicy) SetSecurityBlockJailbrokenDevices(value *bool)() {
    err := m.GetBackingStore().Set("securityBlockJailbrokenDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityDisableUsbDebugging sets the securityDisableUsbDebugging property value. Disable USB debugging on Android devices.
func (m *AndroidForWorkCompliancePolicy) SetSecurityDisableUsbDebugging(value *bool)() {
    err := m.GetBackingStore().Set("securityDisableUsbDebugging", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityPreventInstallAppsFromUnknownSources sets the securityPreventInstallAppsFromUnknownSources property value. Require that devices disallow installation of apps from unknown sources.
func (m *AndroidForWorkCompliancePolicy) SetSecurityPreventInstallAppsFromUnknownSources(value *bool)() {
    err := m.GetBackingStore().Set("securityPreventInstallAppsFromUnknownSources", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityRequireCompanyPortalAppIntegrity sets the securityRequireCompanyPortalAppIntegrity property value. Require the device to pass the Company Portal client app runtime integrity check.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequireCompanyPortalAppIntegrity(value *bool)() {
    err := m.GetBackingStore().Set("securityRequireCompanyPortalAppIntegrity", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityRequiredAndroidSafetyNetEvaluationType sets the securityRequiredAndroidSafetyNetEvaluationType property value. An enum representing the Android SafetyNet attestation evaluation types.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequiredAndroidSafetyNetEvaluationType(value *AndroidSafetyNetEvaluationType)() {
    err := m.GetBackingStore().Set("securityRequiredAndroidSafetyNetEvaluationType", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityRequireGooglePlayServices sets the securityRequireGooglePlayServices property value. Require Google Play Services to be installed and enabled on the device.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequireGooglePlayServices(value *bool)() {
    err := m.GetBackingStore().Set("securityRequireGooglePlayServices", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityRequireSafetyNetAttestationBasicIntegrity sets the securityRequireSafetyNetAttestationBasicIntegrity property value. Require the device to pass the SafetyNet basic integrity check.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequireSafetyNetAttestationBasicIntegrity(value *bool)() {
    err := m.GetBackingStore().Set("securityRequireSafetyNetAttestationBasicIntegrity", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityRequireSafetyNetAttestationCertifiedDevice sets the securityRequireSafetyNetAttestationCertifiedDevice property value. Require the device to pass the SafetyNet certified device check.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequireSafetyNetAttestationCertifiedDevice(value *bool)() {
    err := m.GetBackingStore().Set("securityRequireSafetyNetAttestationCertifiedDevice", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityRequireUpToDateSecurityProviders sets the securityRequireUpToDateSecurityProviders property value. Require the device to have up to date security providers. The device will require Google Play Services to be enabled and up to date.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequireUpToDateSecurityProviders(value *bool)() {
    err := m.GetBackingStore().Set("securityRequireUpToDateSecurityProviders", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityRequireVerifyApps sets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidForWorkCompliancePolicy) SetSecurityRequireVerifyApps(value *bool)() {
    err := m.GetBackingStore().Set("securityRequireVerifyApps", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageRequireEncryption sets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AndroidForWorkCompliancePolicy) SetStorageRequireEncryption(value *bool)() {
    err := m.GetBackingStore().Set("storageRequireEncryption", value)
    if err != nil {
        panic(err)
    }
}
// AndroidForWorkCompliancePolicyable 
type AndroidForWorkCompliancePolicyable interface {
    DeviceCompliancePolicyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceThreatProtectionEnabled()(*bool)
    GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel)
    GetMinAndroidSecurityPatchLevel()(*string)
    GetOsMaximumVersion()(*string)
    GetOsMinimumVersion()(*string)
    GetPasswordExpirationDays()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinutesOfInactivityBeforeLock()(*int32)
    GetPasswordPreviousPasswordBlockCount()(*int32)
    GetPasswordRequired()(*bool)
    GetPasswordRequiredType()(*AndroidRequiredPasswordType)
    GetPasswordSignInFailureCountBeforeFactoryReset()(*int32)
    GetRequiredPasswordComplexity()(*AndroidRequiredPasswordComplexity)
    GetSecurityBlockJailbrokenDevices()(*bool)
    GetSecurityDisableUsbDebugging()(*bool)
    GetSecurityPreventInstallAppsFromUnknownSources()(*bool)
    GetSecurityRequireCompanyPortalAppIntegrity()(*bool)
    GetSecurityRequiredAndroidSafetyNetEvaluationType()(*AndroidSafetyNetEvaluationType)
    GetSecurityRequireGooglePlayServices()(*bool)
    GetSecurityRequireSafetyNetAttestationBasicIntegrity()(*bool)
    GetSecurityRequireSafetyNetAttestationCertifiedDevice()(*bool)
    GetSecurityRequireUpToDateSecurityProviders()(*bool)
    GetSecurityRequireVerifyApps()(*bool)
    GetStorageRequireEncryption()(*bool)
    SetDeviceThreatProtectionEnabled(value *bool)()
    SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)()
    SetMinAndroidSecurityPatchLevel(value *string)()
    SetOsMaximumVersion(value *string)()
    SetOsMinimumVersion(value *string)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinutesOfInactivityBeforeLock(value *int32)()
    SetPasswordPreviousPasswordBlockCount(value *int32)()
    SetPasswordRequired(value *bool)()
    SetPasswordRequiredType(value *AndroidRequiredPasswordType)()
    SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)()
    SetRequiredPasswordComplexity(value *AndroidRequiredPasswordComplexity)()
    SetSecurityBlockJailbrokenDevices(value *bool)()
    SetSecurityDisableUsbDebugging(value *bool)()
    SetSecurityPreventInstallAppsFromUnknownSources(value *bool)()
    SetSecurityRequireCompanyPortalAppIntegrity(value *bool)()
    SetSecurityRequiredAndroidSafetyNetEvaluationType(value *AndroidSafetyNetEvaluationType)()
    SetSecurityRequireGooglePlayServices(value *bool)()
    SetSecurityRequireSafetyNetAttestationBasicIntegrity(value *bool)()
    SetSecurityRequireSafetyNetAttestationCertifiedDevice(value *bool)()
    SetSecurityRequireUpToDateSecurityProviders(value *bool)()
    SetSecurityRequireVerifyApps(value *bool)()
    SetStorageRequireEncryption(value *bool)()
}
