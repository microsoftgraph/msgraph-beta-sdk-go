package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerCompliancePolicy this topic provides descriptions of the declared methods, properties and relationships exposed by the AndroidDeviceOwnerCompliancePolicy resource.
type AndroidDeviceOwnerCompliancePolicy struct {
    DeviceCompliancePolicy
    // The OdataType property
    OdataType *string
}
// NewAndroidDeviceOwnerCompliancePolicy instantiates a new androidDeviceOwnerCompliancePolicy and sets the default values.
func NewAndroidDeviceOwnerCompliancePolicy()(*AndroidDeviceOwnerCompliancePolicy) {
    m := &AndroidDeviceOwnerCompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerCompliancePolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidDeviceOwnerCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerCompliancePolicy(), nil
}
// GetAdvancedThreatProtectionRequiredSecurityLevel gets the advancedThreatProtectionRequiredSecurityLevel property value. MDATP Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *AndroidDeviceOwnerCompliancePolicy) GetAdvancedThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionRequiredSecurityLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceThreatProtectionLevel)
    }
    return nil
}
// GetDeviceThreatProtectionEnabled gets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *AndroidDeviceOwnerCompliancePolicy) GetDeviceThreatProtectionEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("deviceThreatProtectionEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeviceThreatProtectionRequiredSecurityLevel gets the deviceThreatProtectionRequiredSecurityLevel property value. Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *AndroidDeviceOwnerCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
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
func (m *AndroidDeviceOwnerCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceCompliancePolicy.GetFieldDeserializers()
    res["advancedThreatProtectionRequiredSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceThreatProtectionLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionRequiredSecurityLevel(val.(*DeviceThreatProtectionLevel))
        }
        return nil
    }
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
    res["passwordMinimumLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLetterCharacters(val)
        }
        return nil
    }
    res["passwordMinimumLowerCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLowerCaseCharacters(val)
        }
        return nil
    }
    res["passwordMinimumNonLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumNonLetterCharacters(val)
        }
        return nil
    }
    res["passwordMinimumNumericCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumNumericCharacters(val)
        }
        return nil
    }
    res["passwordMinimumSymbolCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumSymbolCharacters(val)
        }
        return nil
    }
    res["passwordMinimumUpperCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumUpperCaseCharacters(val)
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
    res["passwordPreviousPasswordCountToBlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPreviousPasswordCountToBlock(val)
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
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*AndroidDeviceOwnerRequiredPasswordType))
        }
        return nil
    }
    res["securityRequireIntuneAppIntegrity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireIntuneAppIntegrity(val)
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
func (m *AndroidDeviceOwnerCompliancePolicy) GetMinAndroidSecurityPatchLevel()(*string) {
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
func (m *AndroidDeviceOwnerCompliancePolicy) GetOsMaximumVersion()(*string) {
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
func (m *AndroidDeviceOwnerCompliancePolicy) GetOsMinimumVersion()(*string) {
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
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordExpirationDays()(*int32) {
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
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumLength()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumLetterCharacters gets the passwordMinimumLetterCharacters property value. Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumLetterCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumLetterCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumLowerCaseCharacters gets the passwordMinimumLowerCaseCharacters property value. Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumLowerCaseCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumLowerCaseCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumNonLetterCharacters gets the passwordMinimumNonLetterCharacters property value. Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumNonLetterCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumNonLetterCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumNumericCharacters gets the passwordMinimumNumericCharacters property value. Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumNumericCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumNumericCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumSymbolCharacters gets the passwordMinimumSymbolCharacters property value. Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumSymbolCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumSymbolCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumUpperCaseCharacters gets the passwordMinimumUpperCaseCharacters property value. Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumUpperCaseCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumUpperCaseCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinutesOfInactivityBeforeLock gets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinutesOfInactivityBeforeLock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordPreviousPasswordCountToBlock gets the passwordPreviousPasswordCountToBlock property value. Number of previous passwords to block. Valid values 1 to 24
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordPreviousPasswordCountToBlock()(*int32) {
    val, err := m.GetBackingStore().Get("passwordPreviousPasswordCountToBlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordRequired gets the passwordRequired property value. Require a password to unlock device.
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordRequired()(*bool) {
    val, err := m.GetBackingStore().Get("passwordRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Type of characters in password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType) {
    val, err := m.GetBackingStore().Get("passwordRequiredType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerRequiredPasswordType)
    }
    return nil
}
// GetSecurityRequireIntuneAppIntegrity gets the securityRequireIntuneAppIntegrity property value. If setting is set to true, checks that the Intune app installed on fully managed, dedicated, or corporate-owned work profile Android Enterprise enrolled devices, is the one provided by Microsoft from the Managed Google Playstore. If the check fails, the device will be reported as non-compliant.
func (m *AndroidDeviceOwnerCompliancePolicy) GetSecurityRequireIntuneAppIntegrity()(*bool) {
    val, err := m.GetBackingStore().Get("securityRequireIntuneAppIntegrity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityRequireSafetyNetAttestationBasicIntegrity gets the securityRequireSafetyNetAttestationBasicIntegrity property value. Require the device to pass the SafetyNet basic integrity check.
func (m *AndroidDeviceOwnerCompliancePolicy) GetSecurityRequireSafetyNetAttestationBasicIntegrity()(*bool) {
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
func (m *AndroidDeviceOwnerCompliancePolicy) GetSecurityRequireSafetyNetAttestationCertifiedDevice()(*bool) {
    val, err := m.GetBackingStore().Get("securityRequireSafetyNetAttestationCertifiedDevice")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStorageRequireEncryption gets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AndroidDeviceOwnerCompliancePolicy) GetStorageRequireEncryption()(*bool) {
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
func (m *AndroidDeviceOwnerCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceCompliancePolicy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdvancedThreatProtectionRequiredSecurityLevel() != nil {
        cast := (*m.GetAdvancedThreatProtectionRequiredSecurityLevel()).String()
        err = writer.WriteStringValue("advancedThreatProtectionRequiredSecurityLevel", &cast)
        if err != nil {
            return err
        }
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
        err = writer.WriteInt32Value("passwordMinimumLetterCharacters", m.GetPasswordMinimumLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLowerCaseCharacters", m.GetPasswordMinimumLowerCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumNonLetterCharacters", m.GetPasswordMinimumNonLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumNumericCharacters", m.GetPasswordMinimumNumericCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumSymbolCharacters", m.GetPasswordMinimumSymbolCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumUpperCaseCharacters", m.GetPasswordMinimumUpperCaseCharacters())
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
        err = writer.WriteInt32Value("passwordPreviousPasswordCountToBlock", m.GetPasswordPreviousPasswordCountToBlock())
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
        err = writer.WriteBoolValue("securityRequireIntuneAppIntegrity", m.GetSecurityRequireIntuneAppIntegrity())
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
        err = writer.WriteBoolValue("storageRequireEncryption", m.GetStorageRequireEncryption())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdvancedThreatProtectionRequiredSecurityLevel sets the advancedThreatProtectionRequiredSecurityLevel property value. MDATP Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *AndroidDeviceOwnerCompliancePolicy) SetAdvancedThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionRequiredSecurityLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceThreatProtectionEnabled sets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *AndroidDeviceOwnerCompliancePolicy) SetDeviceThreatProtectionEnabled(value *bool)() {
    err := m.GetBackingStore().Set("deviceThreatProtectionEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceThreatProtectionRequiredSecurityLevel sets the deviceThreatProtectionRequiredSecurityLevel property value. Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *AndroidDeviceOwnerCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    err := m.GetBackingStore().Set("deviceThreatProtectionRequiredSecurityLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetMinAndroidSecurityPatchLevel sets the minAndroidSecurityPatchLevel property value. Minimum Android security patch level.
func (m *AndroidDeviceOwnerCompliancePolicy) SetMinAndroidSecurityPatchLevel(value *string)() {
    err := m.GetBackingStore().Set("minAndroidSecurityPatchLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum Android version.
func (m *AndroidDeviceOwnerCompliancePolicy) SetOsMaximumVersion(value *string)() {
    err := m.GetBackingStore().Set("osMaximumVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum Android version.
func (m *AndroidDeviceOwnerCompliancePolicy) SetOsMinimumVersion(value *string)() {
    err := m.GetBackingStore().Set("osMinimumVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordExpirationDays(value *int32)() {
    err := m.GetBackingStore().Set("passwordExpirationDays", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumLetterCharacters sets the passwordMinimumLetterCharacters property value. Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumLetterCharacters(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumLetterCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumLowerCaseCharacters sets the passwordMinimumLowerCaseCharacters property value. Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumLowerCaseCharacters(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumLowerCaseCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumNonLetterCharacters sets the passwordMinimumNonLetterCharacters property value. Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumNonLetterCharacters(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumNonLetterCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumNumericCharacters sets the passwordMinimumNumericCharacters property value. Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumNumericCharacters(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumNumericCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumSymbolCharacters sets the passwordMinimumSymbolCharacters property value. Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumSymbolCharacters(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumSymbolCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumUpperCaseCharacters sets the passwordMinimumUpperCaseCharacters property value. Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumUpperCaseCharacters(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumUpperCaseCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinutesOfInactivityBeforeLock", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordPreviousPasswordCountToBlock sets the passwordPreviousPasswordCountToBlock property value. Number of previous passwords to block. Valid values 1 to 24
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordPreviousPasswordCountToBlock(value *int32)() {
    err := m.GetBackingStore().Set("passwordPreviousPasswordCountToBlock", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequired sets the passwordRequired property value. Require a password to unlock device.
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordRequired(value *bool)() {
    err := m.GetBackingStore().Set("passwordRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Type of characters in password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)() {
    err := m.GetBackingStore().Set("passwordRequiredType", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityRequireIntuneAppIntegrity sets the securityRequireIntuneAppIntegrity property value. If setting is set to true, checks that the Intune app installed on fully managed, dedicated, or corporate-owned work profile Android Enterprise enrolled devices, is the one provided by Microsoft from the Managed Google Playstore. If the check fails, the device will be reported as non-compliant.
func (m *AndroidDeviceOwnerCompliancePolicy) SetSecurityRequireIntuneAppIntegrity(value *bool)() {
    err := m.GetBackingStore().Set("securityRequireIntuneAppIntegrity", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityRequireSafetyNetAttestationBasicIntegrity sets the securityRequireSafetyNetAttestationBasicIntegrity property value. Require the device to pass the SafetyNet basic integrity check.
func (m *AndroidDeviceOwnerCompliancePolicy) SetSecurityRequireSafetyNetAttestationBasicIntegrity(value *bool)() {
    err := m.GetBackingStore().Set("securityRequireSafetyNetAttestationBasicIntegrity", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityRequireSafetyNetAttestationCertifiedDevice sets the securityRequireSafetyNetAttestationCertifiedDevice property value. Require the device to pass the SafetyNet certified device check.
func (m *AndroidDeviceOwnerCompliancePolicy) SetSecurityRequireSafetyNetAttestationCertifiedDevice(value *bool)() {
    err := m.GetBackingStore().Set("securityRequireSafetyNetAttestationCertifiedDevice", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageRequireEncryption sets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AndroidDeviceOwnerCompliancePolicy) SetStorageRequireEncryption(value *bool)() {
    err := m.GetBackingStore().Set("storageRequireEncryption", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerCompliancePolicyable 
type AndroidDeviceOwnerCompliancePolicyable interface {
    DeviceCompliancePolicyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel)
    GetDeviceThreatProtectionEnabled()(*bool)
    GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel)
    GetMinAndroidSecurityPatchLevel()(*string)
    GetOsMaximumVersion()(*string)
    GetOsMinimumVersion()(*string)
    GetPasswordExpirationDays()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinimumLetterCharacters()(*int32)
    GetPasswordMinimumLowerCaseCharacters()(*int32)
    GetPasswordMinimumNonLetterCharacters()(*int32)
    GetPasswordMinimumNumericCharacters()(*int32)
    GetPasswordMinimumSymbolCharacters()(*int32)
    GetPasswordMinimumUpperCaseCharacters()(*int32)
    GetPasswordMinutesOfInactivityBeforeLock()(*int32)
    GetPasswordPreviousPasswordCountToBlock()(*int32)
    GetPasswordRequired()(*bool)
    GetPasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType)
    GetSecurityRequireIntuneAppIntegrity()(*bool)
    GetSecurityRequireSafetyNetAttestationBasicIntegrity()(*bool)
    GetSecurityRequireSafetyNetAttestationCertifiedDevice()(*bool)
    GetStorageRequireEncryption()(*bool)
    SetAdvancedThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)()
    SetDeviceThreatProtectionEnabled(value *bool)()
    SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)()
    SetMinAndroidSecurityPatchLevel(value *string)()
    SetOsMaximumVersion(value *string)()
    SetOsMinimumVersion(value *string)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinimumLetterCharacters(value *int32)()
    SetPasswordMinimumLowerCaseCharacters(value *int32)()
    SetPasswordMinimumNonLetterCharacters(value *int32)()
    SetPasswordMinimumNumericCharacters(value *int32)()
    SetPasswordMinimumSymbolCharacters(value *int32)()
    SetPasswordMinimumUpperCaseCharacters(value *int32)()
    SetPasswordMinutesOfInactivityBeforeLock(value *int32)()
    SetPasswordPreviousPasswordCountToBlock(value *int32)()
    SetPasswordRequired(value *bool)()
    SetPasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)()
    SetSecurityRequireIntuneAppIntegrity(value *bool)()
    SetSecurityRequireSafetyNetAttestationBasicIntegrity(value *bool)()
    SetSecurityRequireSafetyNetAttestationCertifiedDevice(value *bool)()
    SetStorageRequireEncryption(value *bool)()
}
