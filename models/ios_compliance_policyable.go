package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosCompliancePolicyable 
type IosCompliancePolicyable interface {
    DeviceCompliancePolicyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel)
    GetDeviceThreatProtectionEnabled()(*bool)
    GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel)
    GetManagedEmailProfileRequired()(*bool)
    GetOsMaximumBuildVersion()(*string)
    GetOsMaximumVersion()(*string)
    GetOsMinimumBuildVersion()(*string)
    GetOsMinimumVersion()(*string)
    GetPasscodeBlockSimple()(*bool)
    GetPasscodeExpirationDays()(*int32)
    GetPasscodeMinimumCharacterSetCount()(*int32)
    GetPasscodeMinimumLength()(*int32)
    GetPasscodeMinutesOfInactivityBeforeLock()(*int32)
    GetPasscodeMinutesOfInactivityBeforeScreenTimeout()(*int32)
    GetPasscodePreviousPasscodeBlockCount()(*int32)
    GetPasscodeRequired()(*bool)
    GetPasscodeRequiredType()(*RequiredPasswordType)
    GetRestrictedApps()([]AppListItemable)
    GetSecurityBlockJailbrokenDevices()(*bool)
    SetAdvancedThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)()
    SetDeviceThreatProtectionEnabled(value *bool)()
    SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)()
    SetManagedEmailProfileRequired(value *bool)()
    SetOsMaximumBuildVersion(value *string)()
    SetOsMaximumVersion(value *string)()
    SetOsMinimumBuildVersion(value *string)()
    SetOsMinimumVersion(value *string)()
    SetPasscodeBlockSimple(value *bool)()
    SetPasscodeExpirationDays(value *int32)()
    SetPasscodeMinimumCharacterSetCount(value *int32)()
    SetPasscodeMinimumLength(value *int32)()
    SetPasscodeMinutesOfInactivityBeforeLock(value *int32)()
    SetPasscodeMinutesOfInactivityBeforeScreenTimeout(value *int32)()
    SetPasscodePreviousPasscodeBlockCount(value *int32)()
    SetPasscodeRequired(value *bool)()
    SetPasscodeRequiredType(value *RequiredPasswordType)()
    SetRestrictedApps(value []AppListItemable)()
    SetSecurityBlockJailbrokenDevices(value *bool)()
}
