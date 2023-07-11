package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSEndpointProtectionConfiguration macOS endpoint protection configuration profile.
type MacOSEndpointProtectionConfiguration struct {
    DeviceConfiguration
    // The OdataType property
    OdataType *string
}
// NewMacOSEndpointProtectionConfiguration instantiates a new macOSEndpointProtectionConfiguration and sets the default values.
func NewMacOSEndpointProtectionConfiguration()(*MacOSEndpointProtectionConfiguration) {
    m := &MacOSEndpointProtectionConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSEndpointProtectionConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSEndpointProtectionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSEndpointProtectionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSEndpointProtectionConfiguration(), nil
}
// GetAdvancedThreatProtectionAutomaticSampleSubmission gets the advancedThreatProtectionAutomaticSampleSubmission property value. Possible values of a property
func (m *MacOSEndpointProtectionConfiguration) GetAdvancedThreatProtectionAutomaticSampleSubmission()(*Enablement) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionAutomaticSampleSubmission")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetAdvancedThreatProtectionCloudDelivered gets the advancedThreatProtectionCloudDelivered property value. Possible values of a property
func (m *MacOSEndpointProtectionConfiguration) GetAdvancedThreatProtectionCloudDelivered()(*Enablement) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionCloudDelivered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetAdvancedThreatProtectionDiagnosticDataCollection gets the advancedThreatProtectionDiagnosticDataCollection property value. Possible values of a property
func (m *MacOSEndpointProtectionConfiguration) GetAdvancedThreatProtectionDiagnosticDataCollection()(*Enablement) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionDiagnosticDataCollection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetAdvancedThreatProtectionExcludedExtensions gets the advancedThreatProtectionExcludedExtensions property value. A list of file extensions to exclude from antivirus scanning for Microsoft Defender Advanced Threat Protection on macOS.
func (m *MacOSEndpointProtectionConfiguration) GetAdvancedThreatProtectionExcludedExtensions()([]string) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionExcludedExtensions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAdvancedThreatProtectionExcludedFiles gets the advancedThreatProtectionExcludedFiles property value. A list of paths to files to exclude from antivirus scanning for Microsoft Defender Advanced Threat Protection on macOS.
func (m *MacOSEndpointProtectionConfiguration) GetAdvancedThreatProtectionExcludedFiles()([]string) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionExcludedFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAdvancedThreatProtectionExcludedFolders gets the advancedThreatProtectionExcludedFolders property value. A list of paths to folders to exclude from antivirus scanning for Microsoft Defender Advanced Threat Protection on macOS.
func (m *MacOSEndpointProtectionConfiguration) GetAdvancedThreatProtectionExcludedFolders()([]string) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionExcludedFolders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAdvancedThreatProtectionExcludedProcesses gets the advancedThreatProtectionExcludedProcesses property value. A list of process names to exclude from antivirus scanning for Microsoft Defender Advanced Threat Protection on macOS.
func (m *MacOSEndpointProtectionConfiguration) GetAdvancedThreatProtectionExcludedProcesses()([]string) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionExcludedProcesses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAdvancedThreatProtectionRealTime gets the advancedThreatProtectionRealTime property value. Possible values of a property
func (m *MacOSEndpointProtectionConfiguration) GetAdvancedThreatProtectionRealTime()(*Enablement) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionRealTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSEndpointProtectionConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["advancedThreatProtectionAutomaticSampleSubmission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionAutomaticSampleSubmission(val.(*Enablement))
        }
        return nil
    }
    res["advancedThreatProtectionCloudDelivered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionCloudDelivered(val.(*Enablement))
        }
        return nil
    }
    res["advancedThreatProtectionDiagnosticDataCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionDiagnosticDataCollection(val.(*Enablement))
        }
        return nil
    }
    res["advancedThreatProtectionExcludedExtensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAdvancedThreatProtectionExcludedExtensions(res)
        }
        return nil
    }
    res["advancedThreatProtectionExcludedFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAdvancedThreatProtectionExcludedFiles(res)
        }
        return nil
    }
    res["advancedThreatProtectionExcludedFolders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAdvancedThreatProtectionExcludedFolders(res)
        }
        return nil
    }
    res["advancedThreatProtectionExcludedProcesses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAdvancedThreatProtectionExcludedProcesses(res)
        }
        return nil
    }
    res["advancedThreatProtectionRealTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionRealTime(val.(*Enablement))
        }
        return nil
    }
    res["fileVaultAllowDeferralUntilSignOut"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultAllowDeferralUntilSignOut(val)
        }
        return nil
    }
    res["fileVaultDisablePromptAtSignOut"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultDisablePromptAtSignOut(val)
        }
        return nil
    }
    res["fileVaultEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultEnabled(val)
        }
        return nil
    }
    res["fileVaultHidePersonalRecoveryKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultHidePersonalRecoveryKey(val)
        }
        return nil
    }
    res["fileVaultInstitutionalRecoveryKeyCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultInstitutionalRecoveryKeyCertificate(val)
        }
        return nil
    }
    res["fileVaultInstitutionalRecoveryKeyCertificateFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultInstitutionalRecoveryKeyCertificateFileName(val)
        }
        return nil
    }
    res["fileVaultNumberOfTimesUserCanIgnore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultNumberOfTimesUserCanIgnore(val)
        }
        return nil
    }
    res["fileVaultPersonalRecoveryKeyHelpMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultPersonalRecoveryKeyHelpMessage(val)
        }
        return nil
    }
    res["fileVaultPersonalRecoveryKeyRotationInMonths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultPersonalRecoveryKeyRotationInMonths(val)
        }
        return nil
    }
    res["fileVaultSelectedRecoveryKeyTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSFileVaultRecoveryKeyTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultSelectedRecoveryKeyTypes(val.(*MacOSFileVaultRecoveryKeyTypes))
        }
        return nil
    }
    res["firewallApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSFirewallApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSFirewallApplicationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MacOSFirewallApplicationable)
                }
            }
            m.SetFirewallApplications(res)
        }
        return nil
    }
    res["firewallBlockAllIncoming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallBlockAllIncoming(val)
        }
        return nil
    }
    res["firewallEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallEnabled(val)
        }
        return nil
    }
    res["firewallEnableStealthMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallEnableStealthMode(val)
        }
        return nil
    }
    res["gatekeeperAllowedAppSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSGatekeeperAppSources)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGatekeeperAllowedAppSource(val.(*MacOSGatekeeperAppSources))
        }
        return nil
    }
    res["gatekeeperBlockOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGatekeeperBlockOverride(val)
        }
        return nil
    }
    return res
}
// GetFileVaultAllowDeferralUntilSignOut gets the fileVaultAllowDeferralUntilSignOut property value. Optional. If set to true, the user can defer the enabling of FileVault until they sign out.
func (m *MacOSEndpointProtectionConfiguration) GetFileVaultAllowDeferralUntilSignOut()(*bool) {
    val, err := m.GetBackingStore().Get("fileVaultAllowDeferralUntilSignOut")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFileVaultDisablePromptAtSignOut gets the fileVaultDisablePromptAtSignOut property value. Optional. When using the Defer option, if set to true, the user is not prompted to enable FileVault at sign-out.
func (m *MacOSEndpointProtectionConfiguration) GetFileVaultDisablePromptAtSignOut()(*bool) {
    val, err := m.GetBackingStore().Get("fileVaultDisablePromptAtSignOut")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFileVaultEnabled gets the fileVaultEnabled property value. Whether FileVault should be enabled or not.
func (m *MacOSEndpointProtectionConfiguration) GetFileVaultEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("fileVaultEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFileVaultHidePersonalRecoveryKey gets the fileVaultHidePersonalRecoveryKey property value. Optional. A hidden personal recovery key does not appear on the user's screen during FileVault encryption, reducing the risk of it ending up in the wrong hands.
func (m *MacOSEndpointProtectionConfiguration) GetFileVaultHidePersonalRecoveryKey()(*bool) {
    val, err := m.GetBackingStore().Get("fileVaultHidePersonalRecoveryKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFileVaultInstitutionalRecoveryKeyCertificate gets the fileVaultInstitutionalRecoveryKeyCertificate property value. Required if selected recovery key type(s) include InstitutionalRecoveryKey. The DER Encoded certificate file used to set an institutional recovery key.
func (m *MacOSEndpointProtectionConfiguration) GetFileVaultInstitutionalRecoveryKeyCertificate()([]byte) {
    val, err := m.GetBackingStore().Get("fileVaultInstitutionalRecoveryKeyCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetFileVaultInstitutionalRecoveryKeyCertificateFileName gets the fileVaultInstitutionalRecoveryKeyCertificateFileName property value. File name of the institutional recovery key certificate to display in UI. (.der).
func (m *MacOSEndpointProtectionConfiguration) GetFileVaultInstitutionalRecoveryKeyCertificateFileName()(*string) {
    val, err := m.GetBackingStore().Get("fileVaultInstitutionalRecoveryKeyCertificateFileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFileVaultNumberOfTimesUserCanIgnore gets the fileVaultNumberOfTimesUserCanIgnore property value. Optional. When using the Defer option, this is the maximum number of times the user can ignore prompts to enable FileVault before FileVault will be required for the user to sign in. If set to -1, it will always prompt to enable FileVault until FileVault is enabled, though it will allow the user to bypass enabling FileVault. Setting this to 0 will disable the feature.
func (m *MacOSEndpointProtectionConfiguration) GetFileVaultNumberOfTimesUserCanIgnore()(*int32) {
    val, err := m.GetBackingStore().Get("fileVaultNumberOfTimesUserCanIgnore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFileVaultPersonalRecoveryKeyHelpMessage gets the fileVaultPersonalRecoveryKeyHelpMessage property value. Required if selected recovery key type(s) include PersonalRecoveryKey. A short message displayed to the user that explains how they can retrieve their personal recovery key.
func (m *MacOSEndpointProtectionConfiguration) GetFileVaultPersonalRecoveryKeyHelpMessage()(*string) {
    val, err := m.GetBackingStore().Get("fileVaultPersonalRecoveryKeyHelpMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFileVaultPersonalRecoveryKeyRotationInMonths gets the fileVaultPersonalRecoveryKeyRotationInMonths property value. Optional. If selected recovery key type(s) include PersonalRecoveryKey, the frequency to rotate that key, in months.
func (m *MacOSEndpointProtectionConfiguration) GetFileVaultPersonalRecoveryKeyRotationInMonths()(*int32) {
    val, err := m.GetBackingStore().Get("fileVaultPersonalRecoveryKeyRotationInMonths")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFileVaultSelectedRecoveryKeyTypes gets the fileVaultSelectedRecoveryKeyTypes property value. Recovery key types for macOS FileVault
func (m *MacOSEndpointProtectionConfiguration) GetFileVaultSelectedRecoveryKeyTypes()(*MacOSFileVaultRecoveryKeyTypes) {
    val, err := m.GetBackingStore().Get("fileVaultSelectedRecoveryKeyTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSFileVaultRecoveryKeyTypes)
    }
    return nil
}
// GetFirewallApplications gets the firewallApplications property value. List of applications with firewall settings. Firewall settings for applications not on this list are determined by the user. This collection can contain a maximum of 500 elements.
func (m *MacOSEndpointProtectionConfiguration) GetFirewallApplications()([]MacOSFirewallApplicationable) {
    val, err := m.GetBackingStore().Get("firewallApplications")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSFirewallApplicationable)
    }
    return nil
}
// GetFirewallBlockAllIncoming gets the firewallBlockAllIncoming property value. Corresponds to the 'Block all incoming connections' option.
func (m *MacOSEndpointProtectionConfiguration) GetFirewallBlockAllIncoming()(*bool) {
    val, err := m.GetBackingStore().Get("firewallBlockAllIncoming")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFirewallEnabled gets the firewallEnabled property value. Whether the firewall should be enabled or not.
func (m *MacOSEndpointProtectionConfiguration) GetFirewallEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("firewallEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFirewallEnableStealthMode gets the firewallEnableStealthMode property value. Corresponds to 'Enable stealth mode.'
func (m *MacOSEndpointProtectionConfiguration) GetFirewallEnableStealthMode()(*bool) {
    val, err := m.GetBackingStore().Get("firewallEnableStealthMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetGatekeeperAllowedAppSource gets the gatekeeperAllowedAppSource property value. App source options for macOS Gatekeeper.
func (m *MacOSEndpointProtectionConfiguration) GetGatekeeperAllowedAppSource()(*MacOSGatekeeperAppSources) {
    val, err := m.GetBackingStore().Get("gatekeeperAllowedAppSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSGatekeeperAppSources)
    }
    return nil
}
// GetGatekeeperBlockOverride gets the gatekeeperBlockOverride property value. If set to true, the user override for Gatekeeper will be disabled.
func (m *MacOSEndpointProtectionConfiguration) GetGatekeeperBlockOverride()(*bool) {
    val, err := m.GetBackingStore().Get("gatekeeperBlockOverride")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSEndpointProtectionConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdvancedThreatProtectionAutomaticSampleSubmission() != nil {
        cast := (*m.GetAdvancedThreatProtectionAutomaticSampleSubmission()).String()
        err = writer.WriteStringValue("advancedThreatProtectionAutomaticSampleSubmission", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAdvancedThreatProtectionCloudDelivered() != nil {
        cast := (*m.GetAdvancedThreatProtectionCloudDelivered()).String()
        err = writer.WriteStringValue("advancedThreatProtectionCloudDelivered", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAdvancedThreatProtectionDiagnosticDataCollection() != nil {
        cast := (*m.GetAdvancedThreatProtectionDiagnosticDataCollection()).String()
        err = writer.WriteStringValue("advancedThreatProtectionDiagnosticDataCollection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAdvancedThreatProtectionExcludedExtensions() != nil {
        err = writer.WriteCollectionOfStringValues("advancedThreatProtectionExcludedExtensions", m.GetAdvancedThreatProtectionExcludedExtensions())
        if err != nil {
            return err
        }
    }
    if m.GetAdvancedThreatProtectionExcludedFiles() != nil {
        err = writer.WriteCollectionOfStringValues("advancedThreatProtectionExcludedFiles", m.GetAdvancedThreatProtectionExcludedFiles())
        if err != nil {
            return err
        }
    }
    if m.GetAdvancedThreatProtectionExcludedFolders() != nil {
        err = writer.WriteCollectionOfStringValues("advancedThreatProtectionExcludedFolders", m.GetAdvancedThreatProtectionExcludedFolders())
        if err != nil {
            return err
        }
    }
    if m.GetAdvancedThreatProtectionExcludedProcesses() != nil {
        err = writer.WriteCollectionOfStringValues("advancedThreatProtectionExcludedProcesses", m.GetAdvancedThreatProtectionExcludedProcesses())
        if err != nil {
            return err
        }
    }
    if m.GetAdvancedThreatProtectionRealTime() != nil {
        cast := (*m.GetAdvancedThreatProtectionRealTime()).String()
        err = writer.WriteStringValue("advancedThreatProtectionRealTime", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fileVaultAllowDeferralUntilSignOut", m.GetFileVaultAllowDeferralUntilSignOut())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fileVaultDisablePromptAtSignOut", m.GetFileVaultDisablePromptAtSignOut())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fileVaultEnabled", m.GetFileVaultEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fileVaultHidePersonalRecoveryKey", m.GetFileVaultHidePersonalRecoveryKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("fileVaultInstitutionalRecoveryKeyCertificate", m.GetFileVaultInstitutionalRecoveryKeyCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileVaultInstitutionalRecoveryKeyCertificateFileName", m.GetFileVaultInstitutionalRecoveryKeyCertificateFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("fileVaultNumberOfTimesUserCanIgnore", m.GetFileVaultNumberOfTimesUserCanIgnore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileVaultPersonalRecoveryKeyHelpMessage", m.GetFileVaultPersonalRecoveryKeyHelpMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("fileVaultPersonalRecoveryKeyRotationInMonths", m.GetFileVaultPersonalRecoveryKeyRotationInMonths())
        if err != nil {
            return err
        }
    }
    if m.GetFileVaultSelectedRecoveryKeyTypes() != nil {
        cast := (*m.GetFileVaultSelectedRecoveryKeyTypes()).String()
        err = writer.WriteStringValue("fileVaultSelectedRecoveryKeyTypes", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFirewallApplications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFirewallApplications()))
        for i, v := range m.GetFirewallApplications() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("firewallApplications", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallBlockAllIncoming", m.GetFirewallBlockAllIncoming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallEnabled", m.GetFirewallEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallEnableStealthMode", m.GetFirewallEnableStealthMode())
        if err != nil {
            return err
        }
    }
    if m.GetGatekeeperAllowedAppSource() != nil {
        cast := (*m.GetGatekeeperAllowedAppSource()).String()
        err = writer.WriteStringValue("gatekeeperAllowedAppSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("gatekeeperBlockOverride", m.GetGatekeeperBlockOverride())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdvancedThreatProtectionAutomaticSampleSubmission sets the advancedThreatProtectionAutomaticSampleSubmission property value. Possible values of a property
func (m *MacOSEndpointProtectionConfiguration) SetAdvancedThreatProtectionAutomaticSampleSubmission(value *Enablement)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionAutomaticSampleSubmission", value)
    if err != nil {
        panic(err)
    }
}
// SetAdvancedThreatProtectionCloudDelivered sets the advancedThreatProtectionCloudDelivered property value. Possible values of a property
func (m *MacOSEndpointProtectionConfiguration) SetAdvancedThreatProtectionCloudDelivered(value *Enablement)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionCloudDelivered", value)
    if err != nil {
        panic(err)
    }
}
// SetAdvancedThreatProtectionDiagnosticDataCollection sets the advancedThreatProtectionDiagnosticDataCollection property value. Possible values of a property
func (m *MacOSEndpointProtectionConfiguration) SetAdvancedThreatProtectionDiagnosticDataCollection(value *Enablement)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionDiagnosticDataCollection", value)
    if err != nil {
        panic(err)
    }
}
// SetAdvancedThreatProtectionExcludedExtensions sets the advancedThreatProtectionExcludedExtensions property value. A list of file extensions to exclude from antivirus scanning for Microsoft Defender Advanced Threat Protection on macOS.
func (m *MacOSEndpointProtectionConfiguration) SetAdvancedThreatProtectionExcludedExtensions(value []string)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionExcludedExtensions", value)
    if err != nil {
        panic(err)
    }
}
// SetAdvancedThreatProtectionExcludedFiles sets the advancedThreatProtectionExcludedFiles property value. A list of paths to files to exclude from antivirus scanning for Microsoft Defender Advanced Threat Protection on macOS.
func (m *MacOSEndpointProtectionConfiguration) SetAdvancedThreatProtectionExcludedFiles(value []string)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionExcludedFiles", value)
    if err != nil {
        panic(err)
    }
}
// SetAdvancedThreatProtectionExcludedFolders sets the advancedThreatProtectionExcludedFolders property value. A list of paths to folders to exclude from antivirus scanning for Microsoft Defender Advanced Threat Protection on macOS.
func (m *MacOSEndpointProtectionConfiguration) SetAdvancedThreatProtectionExcludedFolders(value []string)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionExcludedFolders", value)
    if err != nil {
        panic(err)
    }
}
// SetAdvancedThreatProtectionExcludedProcesses sets the advancedThreatProtectionExcludedProcesses property value. A list of process names to exclude from antivirus scanning for Microsoft Defender Advanced Threat Protection on macOS.
func (m *MacOSEndpointProtectionConfiguration) SetAdvancedThreatProtectionExcludedProcesses(value []string)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionExcludedProcesses", value)
    if err != nil {
        panic(err)
    }
}
// SetAdvancedThreatProtectionRealTime sets the advancedThreatProtectionRealTime property value. Possible values of a property
func (m *MacOSEndpointProtectionConfiguration) SetAdvancedThreatProtectionRealTime(value *Enablement)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionRealTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVaultAllowDeferralUntilSignOut sets the fileVaultAllowDeferralUntilSignOut property value. Optional. If set to true, the user can defer the enabling of FileVault until they sign out.
func (m *MacOSEndpointProtectionConfiguration) SetFileVaultAllowDeferralUntilSignOut(value *bool)() {
    err := m.GetBackingStore().Set("fileVaultAllowDeferralUntilSignOut", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVaultDisablePromptAtSignOut sets the fileVaultDisablePromptAtSignOut property value. Optional. When using the Defer option, if set to true, the user is not prompted to enable FileVault at sign-out.
func (m *MacOSEndpointProtectionConfiguration) SetFileVaultDisablePromptAtSignOut(value *bool)() {
    err := m.GetBackingStore().Set("fileVaultDisablePromptAtSignOut", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVaultEnabled sets the fileVaultEnabled property value. Whether FileVault should be enabled or not.
func (m *MacOSEndpointProtectionConfiguration) SetFileVaultEnabled(value *bool)() {
    err := m.GetBackingStore().Set("fileVaultEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVaultHidePersonalRecoveryKey sets the fileVaultHidePersonalRecoveryKey property value. Optional. A hidden personal recovery key does not appear on the user's screen during FileVault encryption, reducing the risk of it ending up in the wrong hands.
func (m *MacOSEndpointProtectionConfiguration) SetFileVaultHidePersonalRecoveryKey(value *bool)() {
    err := m.GetBackingStore().Set("fileVaultHidePersonalRecoveryKey", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVaultInstitutionalRecoveryKeyCertificate sets the fileVaultInstitutionalRecoveryKeyCertificate property value. Required if selected recovery key type(s) include InstitutionalRecoveryKey. The DER Encoded certificate file used to set an institutional recovery key.
func (m *MacOSEndpointProtectionConfiguration) SetFileVaultInstitutionalRecoveryKeyCertificate(value []byte)() {
    err := m.GetBackingStore().Set("fileVaultInstitutionalRecoveryKeyCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVaultInstitutionalRecoveryKeyCertificateFileName sets the fileVaultInstitutionalRecoveryKeyCertificateFileName property value. File name of the institutional recovery key certificate to display in UI. (.der).
func (m *MacOSEndpointProtectionConfiguration) SetFileVaultInstitutionalRecoveryKeyCertificateFileName(value *string)() {
    err := m.GetBackingStore().Set("fileVaultInstitutionalRecoveryKeyCertificateFileName", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVaultNumberOfTimesUserCanIgnore sets the fileVaultNumberOfTimesUserCanIgnore property value. Optional. When using the Defer option, this is the maximum number of times the user can ignore prompts to enable FileVault before FileVault will be required for the user to sign in. If set to -1, it will always prompt to enable FileVault until FileVault is enabled, though it will allow the user to bypass enabling FileVault. Setting this to 0 will disable the feature.
func (m *MacOSEndpointProtectionConfiguration) SetFileVaultNumberOfTimesUserCanIgnore(value *int32)() {
    err := m.GetBackingStore().Set("fileVaultNumberOfTimesUserCanIgnore", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVaultPersonalRecoveryKeyHelpMessage sets the fileVaultPersonalRecoveryKeyHelpMessage property value. Required if selected recovery key type(s) include PersonalRecoveryKey. A short message displayed to the user that explains how they can retrieve their personal recovery key.
func (m *MacOSEndpointProtectionConfiguration) SetFileVaultPersonalRecoveryKeyHelpMessage(value *string)() {
    err := m.GetBackingStore().Set("fileVaultPersonalRecoveryKeyHelpMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVaultPersonalRecoveryKeyRotationInMonths sets the fileVaultPersonalRecoveryKeyRotationInMonths property value. Optional. If selected recovery key type(s) include PersonalRecoveryKey, the frequency to rotate that key, in months.
func (m *MacOSEndpointProtectionConfiguration) SetFileVaultPersonalRecoveryKeyRotationInMonths(value *int32)() {
    err := m.GetBackingStore().Set("fileVaultPersonalRecoveryKeyRotationInMonths", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVaultSelectedRecoveryKeyTypes sets the fileVaultSelectedRecoveryKeyTypes property value. Recovery key types for macOS FileVault
func (m *MacOSEndpointProtectionConfiguration) SetFileVaultSelectedRecoveryKeyTypes(value *MacOSFileVaultRecoveryKeyTypes)() {
    err := m.GetBackingStore().Set("fileVaultSelectedRecoveryKeyTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallApplications sets the firewallApplications property value. List of applications with firewall settings. Firewall settings for applications not on this list are determined by the user. This collection can contain a maximum of 500 elements.
func (m *MacOSEndpointProtectionConfiguration) SetFirewallApplications(value []MacOSFirewallApplicationable)() {
    err := m.GetBackingStore().Set("firewallApplications", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallBlockAllIncoming sets the firewallBlockAllIncoming property value. Corresponds to the 'Block all incoming connections' option.
func (m *MacOSEndpointProtectionConfiguration) SetFirewallBlockAllIncoming(value *bool)() {
    err := m.GetBackingStore().Set("firewallBlockAllIncoming", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallEnabled sets the firewallEnabled property value. Whether the firewall should be enabled or not.
func (m *MacOSEndpointProtectionConfiguration) SetFirewallEnabled(value *bool)() {
    err := m.GetBackingStore().Set("firewallEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallEnableStealthMode sets the firewallEnableStealthMode property value. Corresponds to 'Enable stealth mode.'
func (m *MacOSEndpointProtectionConfiguration) SetFirewallEnableStealthMode(value *bool)() {
    err := m.GetBackingStore().Set("firewallEnableStealthMode", value)
    if err != nil {
        panic(err)
    }
}
// SetGatekeeperAllowedAppSource sets the gatekeeperAllowedAppSource property value. App source options for macOS Gatekeeper.
func (m *MacOSEndpointProtectionConfiguration) SetGatekeeperAllowedAppSource(value *MacOSGatekeeperAppSources)() {
    err := m.GetBackingStore().Set("gatekeeperAllowedAppSource", value)
    if err != nil {
        panic(err)
    }
}
// SetGatekeeperBlockOverride sets the gatekeeperBlockOverride property value. If set to true, the user override for Gatekeeper will be disabled.
func (m *MacOSEndpointProtectionConfiguration) SetGatekeeperBlockOverride(value *bool)() {
    err := m.GetBackingStore().Set("gatekeeperBlockOverride", value)
    if err != nil {
        panic(err)
    }
}
// MacOSEndpointProtectionConfigurationable 
type MacOSEndpointProtectionConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedThreatProtectionAutomaticSampleSubmission()(*Enablement)
    GetAdvancedThreatProtectionCloudDelivered()(*Enablement)
    GetAdvancedThreatProtectionDiagnosticDataCollection()(*Enablement)
    GetAdvancedThreatProtectionExcludedExtensions()([]string)
    GetAdvancedThreatProtectionExcludedFiles()([]string)
    GetAdvancedThreatProtectionExcludedFolders()([]string)
    GetAdvancedThreatProtectionExcludedProcesses()([]string)
    GetAdvancedThreatProtectionRealTime()(*Enablement)
    GetFileVaultAllowDeferralUntilSignOut()(*bool)
    GetFileVaultDisablePromptAtSignOut()(*bool)
    GetFileVaultEnabled()(*bool)
    GetFileVaultHidePersonalRecoveryKey()(*bool)
    GetFileVaultInstitutionalRecoveryKeyCertificate()([]byte)
    GetFileVaultInstitutionalRecoveryKeyCertificateFileName()(*string)
    GetFileVaultNumberOfTimesUserCanIgnore()(*int32)
    GetFileVaultPersonalRecoveryKeyHelpMessage()(*string)
    GetFileVaultPersonalRecoveryKeyRotationInMonths()(*int32)
    GetFileVaultSelectedRecoveryKeyTypes()(*MacOSFileVaultRecoveryKeyTypes)
    GetFirewallApplications()([]MacOSFirewallApplicationable)
    GetFirewallBlockAllIncoming()(*bool)
    GetFirewallEnabled()(*bool)
    GetFirewallEnableStealthMode()(*bool)
    GetGatekeeperAllowedAppSource()(*MacOSGatekeeperAppSources)
    GetGatekeeperBlockOverride()(*bool)
    SetAdvancedThreatProtectionAutomaticSampleSubmission(value *Enablement)()
    SetAdvancedThreatProtectionCloudDelivered(value *Enablement)()
    SetAdvancedThreatProtectionDiagnosticDataCollection(value *Enablement)()
    SetAdvancedThreatProtectionExcludedExtensions(value []string)()
    SetAdvancedThreatProtectionExcludedFiles(value []string)()
    SetAdvancedThreatProtectionExcludedFolders(value []string)()
    SetAdvancedThreatProtectionExcludedProcesses(value []string)()
    SetAdvancedThreatProtectionRealTime(value *Enablement)()
    SetFileVaultAllowDeferralUntilSignOut(value *bool)()
    SetFileVaultDisablePromptAtSignOut(value *bool)()
    SetFileVaultEnabled(value *bool)()
    SetFileVaultHidePersonalRecoveryKey(value *bool)()
    SetFileVaultInstitutionalRecoveryKeyCertificate(value []byte)()
    SetFileVaultInstitutionalRecoveryKeyCertificateFileName(value *string)()
    SetFileVaultNumberOfTimesUserCanIgnore(value *int32)()
    SetFileVaultPersonalRecoveryKeyHelpMessage(value *string)()
    SetFileVaultPersonalRecoveryKeyRotationInMonths(value *int32)()
    SetFileVaultSelectedRecoveryKeyTypes(value *MacOSFileVaultRecoveryKeyTypes)()
    SetFirewallApplications(value []MacOSFirewallApplicationable)()
    SetFirewallBlockAllIncoming(value *bool)()
    SetFirewallEnabled(value *bool)()
    SetFirewallEnableStealthMode(value *bool)()
    SetGatekeeperAllowedAppSource(value *MacOSGatekeeperAppSources)()
    SetGatekeeperBlockOverride(value *bool)()
}
