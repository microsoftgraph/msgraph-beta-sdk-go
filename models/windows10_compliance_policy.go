package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10CompliancePolicy this class contains compliance settings for Windows 10.
type Windows10CompliancePolicy struct {
    DeviceCompliancePolicy
}
// NewWindows10CompliancePolicy instantiates a new Windows10CompliancePolicy and sets the default values.
func NewWindows10CompliancePolicy()(*Windows10CompliancePolicy) {
    m := &Windows10CompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    odataTypeValue := "#microsoft.graph.windows10CompliancePolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10CompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindows10CompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10CompliancePolicy(), nil
}
// GetActiveFirewallRequired gets the activeFirewallRequired property value. Require active firewall on Windows devices.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetActiveFirewallRequired()(*bool) {
    val, err := m.GetBackingStore().Get("activeFirewallRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAntiSpywareRequired gets the antiSpywareRequired property value. Require any AntiSpyware solution registered with Windows Decurity Center to be on and monitoring (e.g. Symantec, Windows Defender).
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetAntiSpywareRequired()(*bool) {
    val, err := m.GetBackingStore().Get("antiSpywareRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAntivirusRequired gets the antivirusRequired property value. Require any Antivirus solution registered with Windows Decurity Center to be on and monitoring (e.g. Symantec, Windows Defender).
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetAntivirusRequired()(*bool) {
    val, err := m.GetBackingStore().Get("antivirusRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBitLockerEnabled gets the bitLockerEnabled property value. Require devices to be reported healthy by Windows Device Health Attestation - bit locker is enabled
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetBitLockerEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("bitLockerEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCodeIntegrityEnabled gets the codeIntegrityEnabled property value. Require devices to be reported as healthy by Windows Device Health Attestation.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetCodeIntegrityEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("codeIntegrityEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetConfigurationManagerComplianceRequired gets the configurationManagerComplianceRequired property value. Require to consider SCCM Compliance state into consideration for Intune Compliance State.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetConfigurationManagerComplianceRequired()(*bool) {
    val, err := m.GetBackingStore().Get("configurationManagerComplianceRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderEnabled gets the defenderEnabled property value. Require Windows Defender Antimalware on Windows devices.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetDefenderEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("defenderEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderVersion gets the defenderVersion property value. Require Windows Defender Antimalware minimum version on Windows devices.
// returns a *string when successful
func (m *Windows10CompliancePolicy) GetDefenderVersion()(*string) {
    val, err := m.GetBackingStore().Get("defenderVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceCompliancePolicyScript gets the deviceCompliancePolicyScript property value. The deviceCompliancePolicyScript property
// returns a DeviceCompliancePolicyScriptable when successful
func (m *Windows10CompliancePolicy) GetDeviceCompliancePolicyScript()(DeviceCompliancePolicyScriptable) {
    val, err := m.GetBackingStore().Get("deviceCompliancePolicyScript")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceCompliancePolicyScriptable)
    }
    return nil
}
// GetDeviceThreatProtectionEnabled gets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetDeviceThreatProtectionEnabled()(*bool) {
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
// returns a *DeviceThreatProtectionLevel when successful
func (m *Windows10CompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    val, err := m.GetBackingStore().Get("deviceThreatProtectionRequiredSecurityLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceThreatProtectionLevel)
    }
    return nil
}
// GetEarlyLaunchAntiMalwareDriverEnabled gets the earlyLaunchAntiMalwareDriverEnabled property value. Require devices to be reported as healthy by Windows Device Health Attestation - early launch antimalware driver is enabled.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetEarlyLaunchAntiMalwareDriverEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("earlyLaunchAntiMalwareDriverEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Windows10CompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceCompliancePolicy.GetFieldDeserializers()
    res["activeFirewallRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveFirewallRequired(val)
        }
        return nil
    }
    res["antiSpywareRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAntiSpywareRequired(val)
        }
        return nil
    }
    res["antivirusRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAntivirusRequired(val)
        }
        return nil
    }
    res["bitLockerEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerEnabled(val)
        }
        return nil
    }
    res["codeIntegrityEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeIntegrityEnabled(val)
        }
        return nil
    }
    res["configurationManagerComplianceRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerComplianceRequired(val)
        }
        return nil
    }
    res["defenderEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderEnabled(val)
        }
        return nil
    }
    res["defenderVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderVersion(val)
        }
        return nil
    }
    res["deviceCompliancePolicyScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCompliancePolicyScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCompliancePolicyScript(val.(DeviceCompliancePolicyScriptable))
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
    res["earlyLaunchAntiMalwareDriverEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEarlyLaunchAntiMalwareDriverEnabled(val)
        }
        return nil
    }
    res["firmwareProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareProtectionEnabled(val)
        }
        return nil
    }
    res["kernelDmaProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKernelDmaProtectionEnabled(val)
        }
        return nil
    }
    res["memoryIntegrityEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryIntegrityEnabled(val)
        }
        return nil
    }
    res["mobileOsMaximumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileOsMaximumVersion(val)
        }
        return nil
    }
    res["mobileOsMinimumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileOsMinimumVersion(val)
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
    res["passwordBlockSimple"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockSimple(val)
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
    res["passwordMinimumCharacterSetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumCharacterSetCount(val)
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
    res["passwordRequiredToUnlockFromIdle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredToUnlockFromIdle(val)
        }
        return nil
    }
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*RequiredPasswordType))
        }
        return nil
    }
    res["requireHealthyDeviceReport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireHealthyDeviceReport(val)
        }
        return nil
    }
    res["rtpEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRtpEnabled(val)
        }
        return nil
    }
    res["secureBootEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureBootEnabled(val)
        }
        return nil
    }
    res["signatureOutOfDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignatureOutOfDate(val)
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
    res["tpmRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmRequired(val)
        }
        return nil
    }
    res["validOperatingSystemBuildRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOperatingSystemVersionRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OperatingSystemVersionRangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OperatingSystemVersionRangeable)
                }
            }
            m.SetValidOperatingSystemBuildRanges(res)
        }
        return nil
    }
    res["virtualizationBasedSecurityEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualizationBasedSecurityEnabled(val)
        }
        return nil
    }
    res["wslDistributions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWslDistributionConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WslDistributionConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WslDistributionConfigurationable)
                }
            }
            m.SetWslDistributions(res)
        }
        return nil
    }
    return res
}
// GetFirmwareProtectionEnabled gets the firmwareProtectionEnabled property value. When TRUE, indicates that Firmware protection is required to be reported as healthy by Microsoft Azure Attestion. When FALSE, indicates that Firmware protection is not required to be reported as healthy. Devices that support either Dynamic Root of Trust for Measurement (DRTM) or Firmware Attack Surface Reduction (FASR) will report compliant for this setting. Default value is FALSE.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetFirmwareProtectionEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("firmwareProtectionEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKernelDmaProtectionEnabled gets the kernelDmaProtectionEnabled property value. When TRUE, indicates that Kernel Direct Memory Access (DMA) protection is required to be reported as healthy by Microsoft Azure Attestion. When FALSE, indicates that Kernel DMA Protection is not required to be reported as healthy. Default value is FALSE.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetKernelDmaProtectionEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("kernelDmaProtectionEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMemoryIntegrityEnabled gets the memoryIntegrityEnabled property value. When TRUE, indicates that Memory Integrity as known as Hypervisor-protected Code Integrity (HVCI) or Hypervisor Enforced Code Integrity protection is required to be reported as healthy by Microsoft Azure Attestion. When FALSE, indicates that Memory Integrity Protection is not required to be reported as healthy. Default value is FALSE.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetMemoryIntegrityEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("memoryIntegrityEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMobileOsMaximumVersion gets the mobileOsMaximumVersion property value. Maximum Windows Phone version.
// returns a *string when successful
func (m *Windows10CompliancePolicy) GetMobileOsMaximumVersion()(*string) {
    val, err := m.GetBackingStore().Get("mobileOsMaximumVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMobileOsMinimumVersion gets the mobileOsMinimumVersion property value. Minimum Windows Phone version.
// returns a *string when successful
func (m *Windows10CompliancePolicy) GetMobileOsMinimumVersion()(*string) {
    val, err := m.GetBackingStore().Get("mobileOsMinimumVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Maximum Windows 10 version.
// returns a *string when successful
func (m *Windows10CompliancePolicy) GetOsMaximumVersion()(*string) {
    val, err := m.GetBackingStore().Get("osMaximumVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Minimum Windows 10 version.
// returns a *string when successful
func (m *Windows10CompliancePolicy) GetOsMinimumVersion()(*string) {
    val, err := m.GetBackingStore().Get("osMinimumVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPasswordBlockSimple gets the passwordBlockSimple property value. Indicates whether or not to block simple password.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetPasswordBlockSimple()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockSimple")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. The password expiration in days.
// returns a *int32 when successful
func (m *Windows10CompliancePolicy) GetPasswordExpirationDays()(*int32) {
    val, err := m.GetBackingStore().Get("passwordExpirationDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumCharacterSetCount gets the passwordMinimumCharacterSetCount property value. The number of character sets required in the password.
// returns a *int32 when successful
func (m *Windows10CompliancePolicy) GetPasswordMinimumCharacterSetCount()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumCharacterSetCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. The minimum password length.
// returns a *int32 when successful
func (m *Windows10CompliancePolicy) GetPasswordMinimumLength()(*int32) {
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
// returns a *int32 when successful
func (m *Windows10CompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinutesOfInactivityBeforeLock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. The number of previous passwords to prevent re-use of.
// returns a *int32 when successful
func (m *Windows10CompliancePolicy) GetPasswordPreviousPasswordBlockCount()(*int32) {
    val, err := m.GetBackingStore().Get("passwordPreviousPasswordBlockCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordRequired gets the passwordRequired property value. Require a password to unlock Windows device.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetPasswordRequired()(*bool) {
    val, err := m.GetBackingStore().Get("passwordRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordRequiredToUnlockFromIdle gets the passwordRequiredToUnlockFromIdle property value. Require a password to unlock an idle device.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetPasswordRequiredToUnlockFromIdle()(*bool) {
    val, err := m.GetBackingStore().Get("passwordRequiredToUnlockFromIdle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Possible values of required passwords.
// returns a *RequiredPasswordType when successful
func (m *Windows10CompliancePolicy) GetPasswordRequiredType()(*RequiredPasswordType) {
    val, err := m.GetBackingStore().Get("passwordRequiredType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RequiredPasswordType)
    }
    return nil
}
// GetRequireHealthyDeviceReport gets the requireHealthyDeviceReport property value. Require devices to be reported as healthy by Windows Device Health Attestation.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetRequireHealthyDeviceReport()(*bool) {
    val, err := m.GetBackingStore().Get("requireHealthyDeviceReport")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRtpEnabled gets the rtpEnabled property value. Require Windows Defender Antimalware Real-Time Protection on Windows devices.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetRtpEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("rtpEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecureBootEnabled gets the secureBootEnabled property value. Require devices to be reported as healthy by Windows Device Health Attestation - secure boot is enabled.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetSecureBootEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("secureBootEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSignatureOutOfDate gets the signatureOutOfDate property value. Require Windows Defender Antimalware Signature to be up to date on Windows devices.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetSignatureOutOfDate()(*bool) {
    val, err := m.GetBackingStore().Get("signatureOutOfDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStorageRequireEncryption gets the storageRequireEncryption property value. Require encryption on windows devices.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetStorageRequireEncryption()(*bool) {
    val, err := m.GetBackingStore().Get("storageRequireEncryption")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTpmRequired gets the tpmRequired property value. Require Trusted Platform Module(TPM) to be present.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetTpmRequired()(*bool) {
    val, err := m.GetBackingStore().Get("tpmRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetValidOperatingSystemBuildRanges gets the validOperatingSystemBuildRanges property value. The valid operating system build ranges on Windows devices. This collection can contain a maximum of 10000 elements.
// returns a []OperatingSystemVersionRangeable when successful
func (m *Windows10CompliancePolicy) GetValidOperatingSystemBuildRanges()([]OperatingSystemVersionRangeable) {
    val, err := m.GetBackingStore().Get("validOperatingSystemBuildRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OperatingSystemVersionRangeable)
    }
    return nil
}
// GetVirtualizationBasedSecurityEnabled gets the virtualizationBasedSecurityEnabled property value. When TRUE, indicates that Virtualization-based Security is required to be reported as healthy by Microsoft Azure Attestion. When FALSE, indicates that Virtualization-based Security is not required to be reported as healthy. Default value is FALSE.
// returns a *bool when successful
func (m *Windows10CompliancePolicy) GetVirtualizationBasedSecurityEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("virtualizationBasedSecurityEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWslDistributions gets the wslDistributions property value. The wslDistributions property
// returns a []WslDistributionConfigurationable when successful
func (m *Windows10CompliancePolicy) GetWslDistributions()([]WslDistributionConfigurationable) {
    val, err := m.GetBackingStore().Get("wslDistributions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WslDistributionConfigurationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Windows10CompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceCompliancePolicy.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("activeFirewallRequired", m.GetActiveFirewallRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("antiSpywareRequired", m.GetAntiSpywareRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("antivirusRequired", m.GetAntivirusRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bitLockerEnabled", m.GetBitLockerEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("codeIntegrityEnabled", m.GetCodeIntegrityEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("configurationManagerComplianceRequired", m.GetConfigurationManagerComplianceRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderEnabled", m.GetDefenderEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defenderVersion", m.GetDefenderVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceCompliancePolicyScript", m.GetDeviceCompliancePolicyScript())
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
        err = writer.WriteBoolValue("earlyLaunchAntiMalwareDriverEnabled", m.GetEarlyLaunchAntiMalwareDriverEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firmwareProtectionEnabled", m.GetFirmwareProtectionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kernelDmaProtectionEnabled", m.GetKernelDmaProtectionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("memoryIntegrityEnabled", m.GetMemoryIntegrityEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobileOsMaximumVersion", m.GetMobileOsMaximumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobileOsMinimumVersion", m.GetMobileOsMinimumVersion())
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
        err = writer.WriteBoolValue("passwordBlockSimple", m.GetPasswordBlockSimple())
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
        err = writer.WriteInt32Value("passwordMinimumCharacterSetCount", m.GetPasswordMinimumCharacterSetCount())
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
    {
        err = writer.WriteBoolValue("passwordRequiredToUnlockFromIdle", m.GetPasswordRequiredToUnlockFromIdle())
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
        err = writer.WriteBoolValue("requireHealthyDeviceReport", m.GetRequireHealthyDeviceReport())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("rtpEnabled", m.GetRtpEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("secureBootEnabled", m.GetSecureBootEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("signatureOutOfDate", m.GetSignatureOutOfDate())
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
    {
        err = writer.WriteBoolValue("tpmRequired", m.GetTpmRequired())
        if err != nil {
            return err
        }
    }
    if m.GetValidOperatingSystemBuildRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValidOperatingSystemBuildRanges()))
        for i, v := range m.GetValidOperatingSystemBuildRanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("validOperatingSystemBuildRanges", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("virtualizationBasedSecurityEnabled", m.GetVirtualizationBasedSecurityEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetWslDistributions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWslDistributions()))
        for i, v := range m.GetWslDistributions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("wslDistributions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveFirewallRequired sets the activeFirewallRequired property value. Require active firewall on Windows devices.
func (m *Windows10CompliancePolicy) SetActiveFirewallRequired(value *bool)() {
    err := m.GetBackingStore().Set("activeFirewallRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetAntiSpywareRequired sets the antiSpywareRequired property value. Require any AntiSpyware solution registered with Windows Decurity Center to be on and monitoring (e.g. Symantec, Windows Defender).
func (m *Windows10CompliancePolicy) SetAntiSpywareRequired(value *bool)() {
    err := m.GetBackingStore().Set("antiSpywareRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetAntivirusRequired sets the antivirusRequired property value. Require any Antivirus solution registered with Windows Decurity Center to be on and monitoring (e.g. Symantec, Windows Defender).
func (m *Windows10CompliancePolicy) SetAntivirusRequired(value *bool)() {
    err := m.GetBackingStore().Set("antivirusRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetBitLockerEnabled sets the bitLockerEnabled property value. Require devices to be reported healthy by Windows Device Health Attestation - bit locker is enabled
func (m *Windows10CompliancePolicy) SetBitLockerEnabled(value *bool)() {
    err := m.GetBackingStore().Set("bitLockerEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetCodeIntegrityEnabled sets the codeIntegrityEnabled property value. Require devices to be reported as healthy by Windows Device Health Attestation.
func (m *Windows10CompliancePolicy) SetCodeIntegrityEnabled(value *bool)() {
    err := m.GetBackingStore().Set("codeIntegrityEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationManagerComplianceRequired sets the configurationManagerComplianceRequired property value. Require to consider SCCM Compliance state into consideration for Intune Compliance State.
func (m *Windows10CompliancePolicy) SetConfigurationManagerComplianceRequired(value *bool)() {
    err := m.GetBackingStore().Set("configurationManagerComplianceRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderEnabled sets the defenderEnabled property value. Require Windows Defender Antimalware on Windows devices.
func (m *Windows10CompliancePolicy) SetDefenderEnabled(value *bool)() {
    err := m.GetBackingStore().Set("defenderEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderVersion sets the defenderVersion property value. Require Windows Defender Antimalware minimum version on Windows devices.
func (m *Windows10CompliancePolicy) SetDefenderVersion(value *string)() {
    err := m.GetBackingStore().Set("defenderVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCompliancePolicyScript sets the deviceCompliancePolicyScript property value. The deviceCompliancePolicyScript property
func (m *Windows10CompliancePolicy) SetDeviceCompliancePolicyScript(value DeviceCompliancePolicyScriptable)() {
    err := m.GetBackingStore().Set("deviceCompliancePolicyScript", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceThreatProtectionEnabled sets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *Windows10CompliancePolicy) SetDeviceThreatProtectionEnabled(value *bool)() {
    err := m.GetBackingStore().Set("deviceThreatProtectionEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceThreatProtectionRequiredSecurityLevel sets the deviceThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *Windows10CompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    err := m.GetBackingStore().Set("deviceThreatProtectionRequiredSecurityLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetEarlyLaunchAntiMalwareDriverEnabled sets the earlyLaunchAntiMalwareDriverEnabled property value. Require devices to be reported as healthy by Windows Device Health Attestation - early launch antimalware driver is enabled.
func (m *Windows10CompliancePolicy) SetEarlyLaunchAntiMalwareDriverEnabled(value *bool)() {
    err := m.GetBackingStore().Set("earlyLaunchAntiMalwareDriverEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetFirmwareProtectionEnabled sets the firmwareProtectionEnabled property value. When TRUE, indicates that Firmware protection is required to be reported as healthy by Microsoft Azure Attestion. When FALSE, indicates that Firmware protection is not required to be reported as healthy. Devices that support either Dynamic Root of Trust for Measurement (DRTM) or Firmware Attack Surface Reduction (FASR) will report compliant for this setting. Default value is FALSE.
func (m *Windows10CompliancePolicy) SetFirmwareProtectionEnabled(value *bool)() {
    err := m.GetBackingStore().Set("firmwareProtectionEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetKernelDmaProtectionEnabled sets the kernelDmaProtectionEnabled property value. When TRUE, indicates that Kernel Direct Memory Access (DMA) protection is required to be reported as healthy by Microsoft Azure Attestion. When FALSE, indicates that Kernel DMA Protection is not required to be reported as healthy. Default value is FALSE.
func (m *Windows10CompliancePolicy) SetKernelDmaProtectionEnabled(value *bool)() {
    err := m.GetBackingStore().Set("kernelDmaProtectionEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMemoryIntegrityEnabled sets the memoryIntegrityEnabled property value. When TRUE, indicates that Memory Integrity as known as Hypervisor-protected Code Integrity (HVCI) or Hypervisor Enforced Code Integrity protection is required to be reported as healthy by Microsoft Azure Attestion. When FALSE, indicates that Memory Integrity Protection is not required to be reported as healthy. Default value is FALSE.
func (m *Windows10CompliancePolicy) SetMemoryIntegrityEnabled(value *bool)() {
    err := m.GetBackingStore().Set("memoryIntegrityEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileOsMaximumVersion sets the mobileOsMaximumVersion property value. Maximum Windows Phone version.
func (m *Windows10CompliancePolicy) SetMobileOsMaximumVersion(value *string)() {
    err := m.GetBackingStore().Set("mobileOsMaximumVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileOsMinimumVersion sets the mobileOsMinimumVersion property value. Minimum Windows Phone version.
func (m *Windows10CompliancePolicy) SetMobileOsMinimumVersion(value *string)() {
    err := m.GetBackingStore().Set("mobileOsMinimumVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum Windows 10 version.
func (m *Windows10CompliancePolicy) SetOsMaximumVersion(value *string)() {
    err := m.GetBackingStore().Set("osMaximumVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum Windows 10 version.
func (m *Windows10CompliancePolicy) SetOsMinimumVersion(value *string)() {
    err := m.GetBackingStore().Set("osMinimumVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockSimple sets the passwordBlockSimple property value. Indicates whether or not to block simple password.
func (m *Windows10CompliancePolicy) SetPasswordBlockSimple(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockSimple", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. The password expiration in days.
func (m *Windows10CompliancePolicy) SetPasswordExpirationDays(value *int32)() {
    err := m.GetBackingStore().Set("passwordExpirationDays", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumCharacterSetCount sets the passwordMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *Windows10CompliancePolicy) SetPasswordMinimumCharacterSetCount(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumCharacterSetCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. The minimum password length.
func (m *Windows10CompliancePolicy) SetPasswordMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *Windows10CompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinutesOfInactivityBeforeLock", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. The number of previous passwords to prevent re-use of.
func (m *Windows10CompliancePolicy) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    err := m.GetBackingStore().Set("passwordPreviousPasswordBlockCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequired sets the passwordRequired property value. Require a password to unlock Windows device.
func (m *Windows10CompliancePolicy) SetPasswordRequired(value *bool)() {
    err := m.GetBackingStore().Set("passwordRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequiredToUnlockFromIdle sets the passwordRequiredToUnlockFromIdle property value. Require a password to unlock an idle device.
func (m *Windows10CompliancePolicy) SetPasswordRequiredToUnlockFromIdle(value *bool)() {
    err := m.GetBackingStore().Set("passwordRequiredToUnlockFromIdle", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Possible values of required passwords.
func (m *Windows10CompliancePolicy) SetPasswordRequiredType(value *RequiredPasswordType)() {
    err := m.GetBackingStore().Set("passwordRequiredType", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireHealthyDeviceReport sets the requireHealthyDeviceReport property value. Require devices to be reported as healthy by Windows Device Health Attestation.
func (m *Windows10CompliancePolicy) SetRequireHealthyDeviceReport(value *bool)() {
    err := m.GetBackingStore().Set("requireHealthyDeviceReport", value)
    if err != nil {
        panic(err)
    }
}
// SetRtpEnabled sets the rtpEnabled property value. Require Windows Defender Antimalware Real-Time Protection on Windows devices.
func (m *Windows10CompliancePolicy) SetRtpEnabled(value *bool)() {
    err := m.GetBackingStore().Set("rtpEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSecureBootEnabled sets the secureBootEnabled property value. Require devices to be reported as healthy by Windows Device Health Attestation - secure boot is enabled.
func (m *Windows10CompliancePolicy) SetSecureBootEnabled(value *bool)() {
    err := m.GetBackingStore().Set("secureBootEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSignatureOutOfDate sets the signatureOutOfDate property value. Require Windows Defender Antimalware Signature to be up to date on Windows devices.
func (m *Windows10CompliancePolicy) SetSignatureOutOfDate(value *bool)() {
    err := m.GetBackingStore().Set("signatureOutOfDate", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageRequireEncryption sets the storageRequireEncryption property value. Require encryption on windows devices.
func (m *Windows10CompliancePolicy) SetStorageRequireEncryption(value *bool)() {
    err := m.GetBackingStore().Set("storageRequireEncryption", value)
    if err != nil {
        panic(err)
    }
}
// SetTpmRequired sets the tpmRequired property value. Require Trusted Platform Module(TPM) to be present.
func (m *Windows10CompliancePolicy) SetTpmRequired(value *bool)() {
    err := m.GetBackingStore().Set("tpmRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetValidOperatingSystemBuildRanges sets the validOperatingSystemBuildRanges property value. The valid operating system build ranges on Windows devices. This collection can contain a maximum of 10000 elements.
func (m *Windows10CompliancePolicy) SetValidOperatingSystemBuildRanges(value []OperatingSystemVersionRangeable)() {
    err := m.GetBackingStore().Set("validOperatingSystemBuildRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetVirtualizationBasedSecurityEnabled sets the virtualizationBasedSecurityEnabled property value. When TRUE, indicates that Virtualization-based Security is required to be reported as healthy by Microsoft Azure Attestion. When FALSE, indicates that Virtualization-based Security is not required to be reported as healthy. Default value is FALSE.
func (m *Windows10CompliancePolicy) SetVirtualizationBasedSecurityEnabled(value *bool)() {
    err := m.GetBackingStore().Set("virtualizationBasedSecurityEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetWslDistributions sets the wslDistributions property value. The wslDistributions property
func (m *Windows10CompliancePolicy) SetWslDistributions(value []WslDistributionConfigurationable)() {
    err := m.GetBackingStore().Set("wslDistributions", value)
    if err != nil {
        panic(err)
    }
}
type Windows10CompliancePolicyable interface {
    DeviceCompliancePolicyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveFirewallRequired()(*bool)
    GetAntiSpywareRequired()(*bool)
    GetAntivirusRequired()(*bool)
    GetBitLockerEnabled()(*bool)
    GetCodeIntegrityEnabled()(*bool)
    GetConfigurationManagerComplianceRequired()(*bool)
    GetDefenderEnabled()(*bool)
    GetDefenderVersion()(*string)
    GetDeviceCompliancePolicyScript()(DeviceCompliancePolicyScriptable)
    GetDeviceThreatProtectionEnabled()(*bool)
    GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel)
    GetEarlyLaunchAntiMalwareDriverEnabled()(*bool)
    GetFirmwareProtectionEnabled()(*bool)
    GetKernelDmaProtectionEnabled()(*bool)
    GetMemoryIntegrityEnabled()(*bool)
    GetMobileOsMaximumVersion()(*string)
    GetMobileOsMinimumVersion()(*string)
    GetOsMaximumVersion()(*string)
    GetOsMinimumVersion()(*string)
    GetPasswordBlockSimple()(*bool)
    GetPasswordExpirationDays()(*int32)
    GetPasswordMinimumCharacterSetCount()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinutesOfInactivityBeforeLock()(*int32)
    GetPasswordPreviousPasswordBlockCount()(*int32)
    GetPasswordRequired()(*bool)
    GetPasswordRequiredToUnlockFromIdle()(*bool)
    GetPasswordRequiredType()(*RequiredPasswordType)
    GetRequireHealthyDeviceReport()(*bool)
    GetRtpEnabled()(*bool)
    GetSecureBootEnabled()(*bool)
    GetSignatureOutOfDate()(*bool)
    GetStorageRequireEncryption()(*bool)
    GetTpmRequired()(*bool)
    GetValidOperatingSystemBuildRanges()([]OperatingSystemVersionRangeable)
    GetVirtualizationBasedSecurityEnabled()(*bool)
    GetWslDistributions()([]WslDistributionConfigurationable)
    SetActiveFirewallRequired(value *bool)()
    SetAntiSpywareRequired(value *bool)()
    SetAntivirusRequired(value *bool)()
    SetBitLockerEnabled(value *bool)()
    SetCodeIntegrityEnabled(value *bool)()
    SetConfigurationManagerComplianceRequired(value *bool)()
    SetDefenderEnabled(value *bool)()
    SetDefenderVersion(value *string)()
    SetDeviceCompliancePolicyScript(value DeviceCompliancePolicyScriptable)()
    SetDeviceThreatProtectionEnabled(value *bool)()
    SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)()
    SetEarlyLaunchAntiMalwareDriverEnabled(value *bool)()
    SetFirmwareProtectionEnabled(value *bool)()
    SetKernelDmaProtectionEnabled(value *bool)()
    SetMemoryIntegrityEnabled(value *bool)()
    SetMobileOsMaximumVersion(value *string)()
    SetMobileOsMinimumVersion(value *string)()
    SetOsMaximumVersion(value *string)()
    SetOsMinimumVersion(value *string)()
    SetPasswordBlockSimple(value *bool)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordMinimumCharacterSetCount(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinutesOfInactivityBeforeLock(value *int32)()
    SetPasswordPreviousPasswordBlockCount(value *int32)()
    SetPasswordRequired(value *bool)()
    SetPasswordRequiredToUnlockFromIdle(value *bool)()
    SetPasswordRequiredType(value *RequiredPasswordType)()
    SetRequireHealthyDeviceReport(value *bool)()
    SetRtpEnabled(value *bool)()
    SetSecureBootEnabled(value *bool)()
    SetSignatureOutOfDate(value *bool)()
    SetStorageRequireEncryption(value *bool)()
    SetTpmRequired(value *bool)()
    SetValidOperatingSystemBuildRanges(value []OperatingSystemVersionRangeable)()
    SetVirtualizationBasedSecurityEnabled(value *bool)()
    SetWslDistributions(value []WslDistributionConfigurationable)()
}
