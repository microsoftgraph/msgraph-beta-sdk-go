package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedDeviceEncryptionState 
type ManagedDeviceEncryptionState struct {
    Entity
    // Advanced BitLocker State. Possible values are: success, noUserConsent, osVolumeUnprotected, osVolumeTpmRequired, osVolumeTpmOnlyRequired, osVolumeTpmPinRequired, osVolumeTpmStartupKeyRequired, osVolumeTpmPinStartupKeyRequired, osVolumeEncryptionMethodMismatch, recoveryKeyBackupFailed, fixedDriveNotEncrypted, fixedDriveEncryptionMethodMismatch, loggedOnUserNonAdmin, windowsRecoveryEnvironmentNotConfigured, tpmNotAvailable, tpmNotReady, networkError.
    advancedBitLockerStates *AdvancedBitLockerState;
    // Device name
    deviceName *string;
    // Platform of the device. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, blackberry, palm, unknown.
    deviceType *DeviceTypes;
    // Encryption policy setting state. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
    encryptionPolicySettingState *ComplianceStatus;
    // Encryption readiness state. Possible values are: notReady, ready.
    encryptionReadinessState *EncryptionReadinessState;
    // Device encryption state. Possible values are: notEncrypted, encrypted.
    encryptionState *EncryptionState;
    // FileVault State. Possible values are: success, driveEncryptedByUser, userDeferredEncryption, escrowNotEnabled.
    fileVaultStates *FileVaultState;
    // Operating system version of the device
    osVersion *string;
    // Policy Details
    policyDetails []EncryptionReportPolicyDetails;
    // Device TPM Version
    tpmSpecificationVersion *string;
    // User name
    userPrincipalName *string;
}
// NewManagedDeviceEncryptionState instantiates a new managedDeviceEncryptionState and sets the default values.
func NewManagedDeviceEncryptionState()(*ManagedDeviceEncryptionState) {
    m := &ManagedDeviceEncryptionState{
        Entity: *NewEntity(),
    }
    return m
}
// GetAdvancedBitLockerStates gets the advancedBitLockerStates property value. Advanced BitLocker State. Possible values are: success, noUserConsent, osVolumeUnprotected, osVolumeTpmRequired, osVolumeTpmOnlyRequired, osVolumeTpmPinRequired, osVolumeTpmStartupKeyRequired, osVolumeTpmPinStartupKeyRequired, osVolumeEncryptionMethodMismatch, recoveryKeyBackupFailed, fixedDriveNotEncrypted, fixedDriveEncryptionMethodMismatch, loggedOnUserNonAdmin, windowsRecoveryEnvironmentNotConfigured, tpmNotAvailable, tpmNotReady, networkError.
func (m *ManagedDeviceEncryptionState) GetAdvancedBitLockerStates()(*AdvancedBitLockerState) {
    if m == nil {
        return nil
    } else {
        return m.advancedBitLockerStates
    }
}
// GetDeviceName gets the deviceName property value. Device name
func (m *ManagedDeviceEncryptionState) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetDeviceType gets the deviceType property value. Platform of the device. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, blackberry, palm, unknown.
func (m *ManagedDeviceEncryptionState) GetDeviceType()(*DeviceTypes) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
// GetEncryptionPolicySettingState gets the encryptionPolicySettingState property value. Encryption policy setting state. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
func (m *ManagedDeviceEncryptionState) GetEncryptionPolicySettingState()(*ComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.encryptionPolicySettingState
    }
}
// GetEncryptionReadinessState gets the encryptionReadinessState property value. Encryption readiness state. Possible values are: notReady, ready.
func (m *ManagedDeviceEncryptionState) GetEncryptionReadinessState()(*EncryptionReadinessState) {
    if m == nil {
        return nil
    } else {
        return m.encryptionReadinessState
    }
}
// GetEncryptionState gets the encryptionState property value. Device encryption state. Possible values are: notEncrypted, encrypted.
func (m *ManagedDeviceEncryptionState) GetEncryptionState()(*EncryptionState) {
    if m == nil {
        return nil
    } else {
        return m.encryptionState
    }
}
// GetFileVaultStates gets the fileVaultStates property value. FileVault State. Possible values are: success, driveEncryptedByUser, userDeferredEncryption, escrowNotEnabled.
func (m *ManagedDeviceEncryptionState) GetFileVaultStates()(*FileVaultState) {
    if m == nil {
        return nil
    } else {
        return m.fileVaultStates
    }
}
// GetOsVersion gets the osVersion property value. Operating system version of the device
func (m *ManagedDeviceEncryptionState) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetPolicyDetails gets the policyDetails property value. Policy Details
func (m *ManagedDeviceEncryptionState) GetPolicyDetails()([]EncryptionReportPolicyDetails) {
    if m == nil {
        return nil
    } else {
        return m.policyDetails
    }
}
// GetTpmSpecificationVersion gets the tpmSpecificationVersion property value. Device TPM Version
func (m *ManagedDeviceEncryptionState) GetTpmSpecificationVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tpmSpecificationVersion
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User name
func (m *ManagedDeviceEncryptionState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceEncryptionState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["advancedBitLockerStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdvancedBitLockerState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedBitLockerStates(val.(*AdvancedBitLockerState))
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val.(*DeviceTypes))
        }
        return nil
    }
    res["encryptionPolicySettingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionPolicySettingState(val.(*ComplianceStatus))
        }
        return nil
    }
    res["encryptionReadinessState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEncryptionReadinessState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionReadinessState(val.(*EncryptionReadinessState))
        }
        return nil
    }
    res["encryptionState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEncryptionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionState(val.(*EncryptionState))
        }
        return nil
    }
    res["fileVaultStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileVaultState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultStates(val.(*FileVaultState))
        }
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["policyDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEncryptionReportPolicyDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EncryptionReportPolicyDetails, len(val))
            for i, v := range val {
                res[i] = *(v.(*EncryptionReportPolicyDetails))
            }
            m.SetPolicyDetails(res)
        }
        return nil
    }
    res["tpmSpecificationVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmSpecificationVersion(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *ManagedDeviceEncryptionState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceEncryptionState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdvancedBitLockerStates() != nil {
        cast := (*m.GetAdvancedBitLockerStates()).String()
        err = writer.WriteStringValue("advancedBitLockerStates", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceType() != nil {
        cast := (*m.GetDeviceType()).String()
        err = writer.WriteStringValue("deviceType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEncryptionPolicySettingState() != nil {
        cast := (*m.GetEncryptionPolicySettingState()).String()
        err = writer.WriteStringValue("encryptionPolicySettingState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEncryptionReadinessState() != nil {
        cast := (*m.GetEncryptionReadinessState()).String()
        err = writer.WriteStringValue("encryptionReadinessState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEncryptionState() != nil {
        cast := (*m.GetEncryptionState()).String()
        err = writer.WriteStringValue("encryptionState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFileVaultStates() != nil {
        cast := (*m.GetFileVaultStates()).String()
        err = writer.WriteStringValue("fileVaultStates", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    if m.GetPolicyDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPolicyDetails()))
        for i, v := range m.GetPolicyDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("policyDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tpmSpecificationVersion", m.GetTpmSpecificationVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdvancedBitLockerStates sets the advancedBitLockerStates property value. Advanced BitLocker State. Possible values are: success, noUserConsent, osVolumeUnprotected, osVolumeTpmRequired, osVolumeTpmOnlyRequired, osVolumeTpmPinRequired, osVolumeTpmStartupKeyRequired, osVolumeTpmPinStartupKeyRequired, osVolumeEncryptionMethodMismatch, recoveryKeyBackupFailed, fixedDriveNotEncrypted, fixedDriveEncryptionMethodMismatch, loggedOnUserNonAdmin, windowsRecoveryEnvironmentNotConfigured, tpmNotAvailable, tpmNotReady, networkError.
func (m *ManagedDeviceEncryptionState) SetAdvancedBitLockerStates(value *AdvancedBitLockerState)() {
    if m != nil {
        m.advancedBitLockerStates = value
    }
}
// SetDeviceName sets the deviceName property value. Device name
func (m *ManagedDeviceEncryptionState) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetDeviceType sets the deviceType property value. Platform of the device. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, blackberry, palm, unknown.
func (m *ManagedDeviceEncryptionState) SetDeviceType(value *DeviceTypes)() {
    if m != nil {
        m.deviceType = value
    }
}
// SetEncryptionPolicySettingState sets the encryptionPolicySettingState property value. Encryption policy setting state. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
func (m *ManagedDeviceEncryptionState) SetEncryptionPolicySettingState(value *ComplianceStatus)() {
    if m != nil {
        m.encryptionPolicySettingState = value
    }
}
// SetEncryptionReadinessState sets the encryptionReadinessState property value. Encryption readiness state. Possible values are: notReady, ready.
func (m *ManagedDeviceEncryptionState) SetEncryptionReadinessState(value *EncryptionReadinessState)() {
    if m != nil {
        m.encryptionReadinessState = value
    }
}
// SetEncryptionState sets the encryptionState property value. Device encryption state. Possible values are: notEncrypted, encrypted.
func (m *ManagedDeviceEncryptionState) SetEncryptionState(value *EncryptionState)() {
    if m != nil {
        m.encryptionState = value
    }
}
// SetFileVaultStates sets the fileVaultStates property value. FileVault State. Possible values are: success, driveEncryptedByUser, userDeferredEncryption, escrowNotEnabled.
func (m *ManagedDeviceEncryptionState) SetFileVaultStates(value *FileVaultState)() {
    if m != nil {
        m.fileVaultStates = value
    }
}
// SetOsVersion sets the osVersion property value. Operating system version of the device
func (m *ManagedDeviceEncryptionState) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetPolicyDetails sets the policyDetails property value. Policy Details
func (m *ManagedDeviceEncryptionState) SetPolicyDetails(value []EncryptionReportPolicyDetails)() {
    if m != nil {
        m.policyDetails = value
    }
}
// SetTpmSpecificationVersion sets the tpmSpecificationVersion property value. Device TPM Version
func (m *ManagedDeviceEncryptionState) SetTpmSpecificationVersion(value *string)() {
    if m != nil {
        m.tpmSpecificationVersion = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User name
func (m *ManagedDeviceEncryptionState) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
