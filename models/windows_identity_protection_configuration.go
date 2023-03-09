package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsIdentityProtectionConfiguration 
type WindowsIdentityProtectionConfiguration struct {
    DeviceConfiguration
}
// NewWindowsIdentityProtectionConfiguration instantiates a new WindowsIdentityProtectionConfiguration and sets the default values.
func NewWindowsIdentityProtectionConfiguration()(*WindowsIdentityProtectionConfiguration) {
    m := &WindowsIdentityProtectionConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsIdentityProtectionConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsIdentityProtectionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsIdentityProtectionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsIdentityProtectionConfiguration(), nil
}
// GetEnhancedAntiSpoofingForFacialFeaturesEnabled gets the enhancedAntiSpoofingForFacialFeaturesEnabled property value. Boolean value used to enable enhanced anti-spoofing for facial feature recognition on Windows Hello face authentication.
func (m *WindowsIdentityProtectionConfiguration) GetEnhancedAntiSpoofingForFacialFeaturesEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("enhancedAntiSpoofingForFacialFeaturesEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsIdentityProtectionConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["enhancedAntiSpoofingForFacialFeaturesEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnhancedAntiSpoofingForFacialFeaturesEnabled(val)
        }
        return nil
    }
    res["pinExpirationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinExpirationInDays(val)
        }
        return nil
    }
    res["pinLowercaseCharactersUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinLowercaseCharactersUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["pinMaximumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinMaximumLength(val)
        }
        return nil
    }
    res["pinMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinMinimumLength(val)
        }
        return nil
    }
    res["pinPreviousBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinPreviousBlockCount(val)
        }
        return nil
    }
    res["pinRecoveryEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinRecoveryEnabled(val)
        }
        return nil
    }
    res["pinSpecialCharactersUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinSpecialCharactersUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["pinUppercaseCharactersUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinUppercaseCharactersUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["securityDeviceRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityDeviceRequired(val)
        }
        return nil
    }
    res["unlockWithBiometricsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnlockWithBiometricsEnabled(val)
        }
        return nil
    }
    res["useCertificatesForOnPremisesAuthEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseCertificatesForOnPremisesAuthEnabled(val)
        }
        return nil
    }
    res["useSecurityKeyForSignin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseSecurityKeyForSignin(val)
        }
        return nil
    }
    res["windowsHelloForBusinessBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsHelloForBusinessBlocked(val)
        }
        return nil
    }
    return res
}
// GetPinExpirationInDays gets the pinExpirationInDays property value. Integer value specifies the period (in days) that a PIN can be used before the system requires the user to change it. Valid values are 0 to 730 inclusive. Valid values 0 to 730
func (m *WindowsIdentityProtectionConfiguration) GetPinExpirationInDays()(*int32) {
    val, err := m.GetBackingStore().Get("pinExpirationInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPinLowercaseCharactersUsage gets the pinLowercaseCharactersUsage property value. Possible values of the ConfigurationUsage list.
func (m *WindowsIdentityProtectionConfiguration) GetPinLowercaseCharactersUsage()(*ConfigurationUsage) {
    val, err := m.GetBackingStore().Get("pinLowercaseCharactersUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConfigurationUsage)
    }
    return nil
}
// GetPinMaximumLength gets the pinMaximumLength property value. Integer value that sets the maximum number of characters allowed for the work PIN. Valid values are 4 to 127 inclusive and greater than or equal to the value set for the minimum PIN. Valid values 4 to 127
func (m *WindowsIdentityProtectionConfiguration) GetPinMaximumLength()(*int32) {
    val, err := m.GetBackingStore().Get("pinMaximumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPinMinimumLength gets the pinMinimumLength property value. Integer value that sets the minimum number of characters required for the Windows Hello for Business PIN. Valid values are 4 to 127 inclusive and less than or equal to the value set for the maximum PIN. Valid values 4 to 127
func (m *WindowsIdentityProtectionConfiguration) GetPinMinimumLength()(*int32) {
    val, err := m.GetBackingStore().Get("pinMinimumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPinPreviousBlockCount gets the pinPreviousBlockCount property value. Controls the ability to prevent users from using past PINs. This must be set between 0 and 50, inclusive, and the current PIN of the user is included in that count. If set to 0, previous PINs are not stored. PIN history is not preserved through a PIN reset. Valid values 0 to 50
func (m *WindowsIdentityProtectionConfiguration) GetPinPreviousBlockCount()(*int32) {
    val, err := m.GetBackingStore().Get("pinPreviousBlockCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPinRecoveryEnabled gets the pinRecoveryEnabled property value. Boolean value that enables a user to change their PIN by using the Windows Hello for Business PIN recovery service.
func (m *WindowsIdentityProtectionConfiguration) GetPinRecoveryEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("pinRecoveryEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPinSpecialCharactersUsage gets the pinSpecialCharactersUsage property value. Possible values of the ConfigurationUsage list.
func (m *WindowsIdentityProtectionConfiguration) GetPinSpecialCharactersUsage()(*ConfigurationUsage) {
    val, err := m.GetBackingStore().Get("pinSpecialCharactersUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConfigurationUsage)
    }
    return nil
}
// GetPinUppercaseCharactersUsage gets the pinUppercaseCharactersUsage property value. Possible values of the ConfigurationUsage list.
func (m *WindowsIdentityProtectionConfiguration) GetPinUppercaseCharactersUsage()(*ConfigurationUsage) {
    val, err := m.GetBackingStore().Get("pinUppercaseCharactersUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConfigurationUsage)
    }
    return nil
}
// GetSecurityDeviceRequired gets the securityDeviceRequired property value. Controls whether to require a Trusted Platform Module (TPM) for provisioning Windows Hello for Business. A TPM provides an additional security benefit in that data stored on it cannot be used on other devices. If set to False, all devices can provision Windows Hello for Business even if there is not a usable TPM.
func (m *WindowsIdentityProtectionConfiguration) GetSecurityDeviceRequired()(*bool) {
    val, err := m.GetBackingStore().Get("securityDeviceRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUnlockWithBiometricsEnabled gets the unlockWithBiometricsEnabled property value. Controls the use of biometric gestures, such as face and fingerprint, as an alternative to the Windows Hello for Business PIN.  If set to False, biometric gestures are not allowed. Users must still configure a PIN as a backup in case of failures.
func (m *WindowsIdentityProtectionConfiguration) GetUnlockWithBiometricsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("unlockWithBiometricsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUseCertificatesForOnPremisesAuthEnabled gets the useCertificatesForOnPremisesAuthEnabled property value. Boolean value that enables Windows Hello for Business to use certificates to authenticate on-premise resources.
func (m *WindowsIdentityProtectionConfiguration) GetUseCertificatesForOnPremisesAuthEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("useCertificatesForOnPremisesAuthEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUseSecurityKeyForSignin gets the useSecurityKeyForSignin property value. Boolean value used to enable the Windows Hello security key as a logon credential.
func (m *WindowsIdentityProtectionConfiguration) GetUseSecurityKeyForSignin()(*bool) {
    val, err := m.GetBackingStore().Get("useSecurityKeyForSignin")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWindowsHelloForBusinessBlocked gets the windowsHelloForBusinessBlocked property value. Boolean value that blocks Windows Hello for Business as a method for signing into Windows.
func (m *WindowsIdentityProtectionConfiguration) GetWindowsHelloForBusinessBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("windowsHelloForBusinessBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsIdentityProtectionConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("enhancedAntiSpoofingForFacialFeaturesEnabled", m.GetEnhancedAntiSpoofingForFacialFeaturesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinExpirationInDays", m.GetPinExpirationInDays())
        if err != nil {
            return err
        }
    }
    if m.GetPinLowercaseCharactersUsage() != nil {
        cast := (*m.GetPinLowercaseCharactersUsage()).String()
        err = writer.WriteStringValue("pinLowercaseCharactersUsage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinMaximumLength", m.GetPinMaximumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinMinimumLength", m.GetPinMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinPreviousBlockCount", m.GetPinPreviousBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("pinRecoveryEnabled", m.GetPinRecoveryEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetPinSpecialCharactersUsage() != nil {
        cast := (*m.GetPinSpecialCharactersUsage()).String()
        err = writer.WriteStringValue("pinSpecialCharactersUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPinUppercaseCharactersUsage() != nil {
        cast := (*m.GetPinUppercaseCharactersUsage()).String()
        err = writer.WriteStringValue("pinUppercaseCharactersUsage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityDeviceRequired", m.GetSecurityDeviceRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("unlockWithBiometricsEnabled", m.GetUnlockWithBiometricsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useCertificatesForOnPremisesAuthEnabled", m.GetUseCertificatesForOnPremisesAuthEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useSecurityKeyForSignin", m.GetUseSecurityKeyForSignin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsHelloForBusinessBlocked", m.GetWindowsHelloForBusinessBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnhancedAntiSpoofingForFacialFeaturesEnabled sets the enhancedAntiSpoofingForFacialFeaturesEnabled property value. Boolean value used to enable enhanced anti-spoofing for facial feature recognition on Windows Hello face authentication.
func (m *WindowsIdentityProtectionConfiguration) SetEnhancedAntiSpoofingForFacialFeaturesEnabled(value *bool)() {
    err := m.GetBackingStore().Set("enhancedAntiSpoofingForFacialFeaturesEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetPinExpirationInDays sets the pinExpirationInDays property value. Integer value specifies the period (in days) that a PIN can be used before the system requires the user to change it. Valid values are 0 to 730 inclusive. Valid values 0 to 730
func (m *WindowsIdentityProtectionConfiguration) SetPinExpirationInDays(value *int32)() {
    err := m.GetBackingStore().Set("pinExpirationInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetPinLowercaseCharactersUsage sets the pinLowercaseCharactersUsage property value. Possible values of the ConfigurationUsage list.
func (m *WindowsIdentityProtectionConfiguration) SetPinLowercaseCharactersUsage(value *ConfigurationUsage)() {
    err := m.GetBackingStore().Set("pinLowercaseCharactersUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetPinMaximumLength sets the pinMaximumLength property value. Integer value that sets the maximum number of characters allowed for the work PIN. Valid values are 4 to 127 inclusive and greater than or equal to the value set for the minimum PIN. Valid values 4 to 127
func (m *WindowsIdentityProtectionConfiguration) SetPinMaximumLength(value *int32)() {
    err := m.GetBackingStore().Set("pinMaximumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetPinMinimumLength sets the pinMinimumLength property value. Integer value that sets the minimum number of characters required for the Windows Hello for Business PIN. Valid values are 4 to 127 inclusive and less than or equal to the value set for the maximum PIN. Valid values 4 to 127
func (m *WindowsIdentityProtectionConfiguration) SetPinMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("pinMinimumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetPinPreviousBlockCount sets the pinPreviousBlockCount property value. Controls the ability to prevent users from using past PINs. This must be set between 0 and 50, inclusive, and the current PIN of the user is included in that count. If set to 0, previous PINs are not stored. PIN history is not preserved through a PIN reset. Valid values 0 to 50
func (m *WindowsIdentityProtectionConfiguration) SetPinPreviousBlockCount(value *int32)() {
    err := m.GetBackingStore().Set("pinPreviousBlockCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPinRecoveryEnabled sets the pinRecoveryEnabled property value. Boolean value that enables a user to change their PIN by using the Windows Hello for Business PIN recovery service.
func (m *WindowsIdentityProtectionConfiguration) SetPinRecoveryEnabled(value *bool)() {
    err := m.GetBackingStore().Set("pinRecoveryEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetPinSpecialCharactersUsage sets the pinSpecialCharactersUsage property value. Possible values of the ConfigurationUsage list.
func (m *WindowsIdentityProtectionConfiguration) SetPinSpecialCharactersUsage(value *ConfigurationUsage)() {
    err := m.GetBackingStore().Set("pinSpecialCharactersUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetPinUppercaseCharactersUsage sets the pinUppercaseCharactersUsage property value. Possible values of the ConfigurationUsage list.
func (m *WindowsIdentityProtectionConfiguration) SetPinUppercaseCharactersUsage(value *ConfigurationUsage)() {
    err := m.GetBackingStore().Set("pinUppercaseCharactersUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityDeviceRequired sets the securityDeviceRequired property value. Controls whether to require a Trusted Platform Module (TPM) for provisioning Windows Hello for Business. A TPM provides an additional security benefit in that data stored on it cannot be used on other devices. If set to False, all devices can provision Windows Hello for Business even if there is not a usable TPM.
func (m *WindowsIdentityProtectionConfiguration) SetSecurityDeviceRequired(value *bool)() {
    err := m.GetBackingStore().Set("securityDeviceRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetUnlockWithBiometricsEnabled sets the unlockWithBiometricsEnabled property value. Controls the use of biometric gestures, such as face and fingerprint, as an alternative to the Windows Hello for Business PIN.  If set to False, biometric gestures are not allowed. Users must still configure a PIN as a backup in case of failures.
func (m *WindowsIdentityProtectionConfiguration) SetUnlockWithBiometricsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("unlockWithBiometricsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetUseCertificatesForOnPremisesAuthEnabled sets the useCertificatesForOnPremisesAuthEnabled property value. Boolean value that enables Windows Hello for Business to use certificates to authenticate on-premise resources.
func (m *WindowsIdentityProtectionConfiguration) SetUseCertificatesForOnPremisesAuthEnabled(value *bool)() {
    err := m.GetBackingStore().Set("useCertificatesForOnPremisesAuthEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetUseSecurityKeyForSignin sets the useSecurityKeyForSignin property value. Boolean value used to enable the Windows Hello security key as a logon credential.
func (m *WindowsIdentityProtectionConfiguration) SetUseSecurityKeyForSignin(value *bool)() {
    err := m.GetBackingStore().Set("useSecurityKeyForSignin", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsHelloForBusinessBlocked sets the windowsHelloForBusinessBlocked property value. Boolean value that blocks Windows Hello for Business as a method for signing into Windows.
func (m *WindowsIdentityProtectionConfiguration) SetWindowsHelloForBusinessBlocked(value *bool)() {
    err := m.GetBackingStore().Set("windowsHelloForBusinessBlocked", value)
    if err != nil {
        panic(err)
    }
}
// WindowsIdentityProtectionConfigurationable 
type WindowsIdentityProtectionConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnhancedAntiSpoofingForFacialFeaturesEnabled()(*bool)
    GetPinExpirationInDays()(*int32)
    GetPinLowercaseCharactersUsage()(*ConfigurationUsage)
    GetPinMaximumLength()(*int32)
    GetPinMinimumLength()(*int32)
    GetPinPreviousBlockCount()(*int32)
    GetPinRecoveryEnabled()(*bool)
    GetPinSpecialCharactersUsage()(*ConfigurationUsage)
    GetPinUppercaseCharactersUsage()(*ConfigurationUsage)
    GetSecurityDeviceRequired()(*bool)
    GetUnlockWithBiometricsEnabled()(*bool)
    GetUseCertificatesForOnPremisesAuthEnabled()(*bool)
    GetUseSecurityKeyForSignin()(*bool)
    GetWindowsHelloForBusinessBlocked()(*bool)
    SetEnhancedAntiSpoofingForFacialFeaturesEnabled(value *bool)()
    SetPinExpirationInDays(value *int32)()
    SetPinLowercaseCharactersUsage(value *ConfigurationUsage)()
    SetPinMaximumLength(value *int32)()
    SetPinMinimumLength(value *int32)()
    SetPinPreviousBlockCount(value *int32)()
    SetPinRecoveryEnabled(value *bool)()
    SetPinSpecialCharactersUsage(value *ConfigurationUsage)()
    SetPinUppercaseCharactersUsage(value *ConfigurationUsage)()
    SetSecurityDeviceRequired(value *bool)()
    SetUnlockWithBiometricsEnabled(value *bool)()
    SetUseCertificatesForOnPremisesAuthEnabled(value *bool)()
    SetUseSecurityKeyForSignin(value *bool)()
    SetWindowsHelloForBusinessBlocked(value *bool)()
}
