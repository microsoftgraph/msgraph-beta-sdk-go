package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerCompliancePolicyable 
type AndroidDeviceOwnerCompliancePolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
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
