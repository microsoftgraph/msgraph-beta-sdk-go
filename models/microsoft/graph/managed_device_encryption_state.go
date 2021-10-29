package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new managedDeviceEncryptionState and sets the default values.
func NewManagedDeviceEncryptionState()(*ManagedDeviceEncryptionState) {
    m := &ManagedDeviceEncryptionState{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the advancedBitLockerStates property value. Advanced BitLocker State. Possible values are: success, noUserConsent, osVolumeUnprotected, osVolumeTpmRequired, osVolumeTpmOnlyRequired, osVolumeTpmPinRequired, osVolumeTpmStartupKeyRequired, osVolumeTpmPinStartupKeyRequired, osVolumeEncryptionMethodMismatch, recoveryKeyBackupFailed, fixedDriveNotEncrypted, fixedDriveEncryptionMethodMismatch, loggedOnUserNonAdmin, windowsRecoveryEnvironmentNotConfigured, tpmNotAvailable, tpmNotReady, networkError.
func (m *ManagedDeviceEncryptionState) GetAdvancedBitLockerStates()(*AdvancedBitLockerState) {
    if m == nil {
        return nil
    } else {
        return m.advancedBitLockerStates
    }
}
// Gets the deviceName property value. Device name
func (m *ManagedDeviceEncryptionState) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the deviceType property value. Platform of the device. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, blackberry, palm, unknown.
func (m *ManagedDeviceEncryptionState) GetDeviceType()(*DeviceTypes) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
// Gets the encryptionPolicySettingState property value. Encryption policy setting state. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
func (m *ManagedDeviceEncryptionState) GetEncryptionPolicySettingState()(*ComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.encryptionPolicySettingState
    }
}
// Gets the encryptionReadinessState property value. Encryption readiness state. Possible values are: notReady, ready.
func (m *ManagedDeviceEncryptionState) GetEncryptionReadinessState()(*EncryptionReadinessState) {
    if m == nil {
        return nil
    } else {
        return m.encryptionReadinessState
    }
}
// Gets the encryptionState property value. Device encryption state. Possible values are: notEncrypted, encrypted.
func (m *ManagedDeviceEncryptionState) GetEncryptionState()(*EncryptionState) {
    if m == nil {
        return nil
    } else {
        return m.encryptionState
    }
}
// Gets the fileVaultStates property value. FileVault State. Possible values are: success, driveEncryptedByUser, userDeferredEncryption, escrowNotEnabled.
func (m *ManagedDeviceEncryptionState) GetFileVaultStates()(*FileVaultState) {
    if m == nil {
        return nil
    } else {
        return m.fileVaultStates
    }
}
// Gets the osVersion property value. Operating system version of the device
func (m *ManagedDeviceEncryptionState) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// Gets the policyDetails property value. Policy Details
func (m *ManagedDeviceEncryptionState) GetPolicyDetails()([]EncryptionReportPolicyDetails) {
    if m == nil {
        return nil
    } else {
        return m.policyDetails
    }
}
// Gets the tpmSpecificationVersion property value. Device TPM Version
func (m *ManagedDeviceEncryptionState) GetTpmSpecificationVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tpmSpecificationVersion
    }
}
// Gets the userPrincipalName property value. User name
func (m *ManagedDeviceEncryptionState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *ManagedDeviceEncryptionState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["advancedBitLockerStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdvancedBitLockerState)
        if err != nil {
            return err
        }
        cast := val.(AdvancedBitLockerState)
        m.SetAdvancedBitLockerStates(&cast)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["deviceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceTypes)
        if err != nil {
            return err
        }
        cast := val.(DeviceTypes)
        m.SetDeviceType(&cast)
        return nil
    }
    res["encryptionPolicySettingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        cast := val.(ComplianceStatus)
        m.SetEncryptionPolicySettingState(&cast)
        return nil
    }
    res["encryptionReadinessState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEncryptionReadinessState)
        if err != nil {
            return err
        }
        cast := val.(EncryptionReadinessState)
        m.SetEncryptionReadinessState(&cast)
        return nil
    }
    res["encryptionState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEncryptionState)
        if err != nil {
            return err
        }
        cast := val.(EncryptionState)
        m.SetEncryptionState(&cast)
        return nil
    }
    res["fileVaultStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileVaultState)
        if err != nil {
            return err
        }
        cast := val.(FileVaultState)
        m.SetFileVaultStates(&cast)
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsVersion(val)
        return nil
    }
    res["policyDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEncryptionReportPolicyDetails() })
        if err != nil {
            return err
        }
        res := make([]EncryptionReportPolicyDetails, len(val))
        for i, v := range val {
            res[i] = *(v.(*EncryptionReportPolicyDetails))
        }
        m.SetPolicyDetails(res)
        return nil
    }
    res["tpmSpecificationVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTpmSpecificationVersion(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *ManagedDeviceEncryptionState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedDeviceEncryptionState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdvancedBitLockerStates() != nil {
        cast := m.GetAdvancedBitLockerStates().String()
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
        cast := m.GetDeviceType().String()
        err = writer.WriteStringValue("deviceType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEncryptionPolicySettingState() != nil {
        cast := m.GetEncryptionPolicySettingState().String()
        err = writer.WriteStringValue("encryptionPolicySettingState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEncryptionReadinessState() != nil {
        cast := m.GetEncryptionReadinessState().String()
        err = writer.WriteStringValue("encryptionReadinessState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEncryptionState() != nil {
        cast := m.GetEncryptionState().String()
        err = writer.WriteStringValue("encryptionState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFileVaultStates() != nil {
        cast := m.GetFileVaultStates().String()
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
    {
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
// Sets the advancedBitLockerStates property value. Advanced BitLocker State. Possible values are: success, noUserConsent, osVolumeUnprotected, osVolumeTpmRequired, osVolumeTpmOnlyRequired, osVolumeTpmPinRequired, osVolumeTpmStartupKeyRequired, osVolumeTpmPinStartupKeyRequired, osVolumeEncryptionMethodMismatch, recoveryKeyBackupFailed, fixedDriveNotEncrypted, fixedDriveEncryptionMethodMismatch, loggedOnUserNonAdmin, windowsRecoveryEnvironmentNotConfigured, tpmNotAvailable, tpmNotReady, networkError.
// Parameters:
//  - value : Value to set for the advancedBitLockerStates property.
func (m *ManagedDeviceEncryptionState) SetAdvancedBitLockerStates(value *AdvancedBitLockerState)() {
    m.advancedBitLockerStates = value
}
// Sets the deviceName property value. Device name
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *ManagedDeviceEncryptionState) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the deviceType property value. Platform of the device. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, blackberry, palm, unknown.
// Parameters:
//  - value : Value to set for the deviceType property.
func (m *ManagedDeviceEncryptionState) SetDeviceType(value *DeviceTypes)() {
    m.deviceType = value
}
// Sets the encryptionPolicySettingState property value. Encryption policy setting state. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
// Parameters:
//  - value : Value to set for the encryptionPolicySettingState property.
func (m *ManagedDeviceEncryptionState) SetEncryptionPolicySettingState(value *ComplianceStatus)() {
    m.encryptionPolicySettingState = value
}
// Sets the encryptionReadinessState property value. Encryption readiness state. Possible values are: notReady, ready.
// Parameters:
//  - value : Value to set for the encryptionReadinessState property.
func (m *ManagedDeviceEncryptionState) SetEncryptionReadinessState(value *EncryptionReadinessState)() {
    m.encryptionReadinessState = value
}
// Sets the encryptionState property value. Device encryption state. Possible values are: notEncrypted, encrypted.
// Parameters:
//  - value : Value to set for the encryptionState property.
func (m *ManagedDeviceEncryptionState) SetEncryptionState(value *EncryptionState)() {
    m.encryptionState = value
}
// Sets the fileVaultStates property value. FileVault State. Possible values are: success, driveEncryptedByUser, userDeferredEncryption, escrowNotEnabled.
// Parameters:
//  - value : Value to set for the fileVaultStates property.
func (m *ManagedDeviceEncryptionState) SetFileVaultStates(value *FileVaultState)() {
    m.fileVaultStates = value
}
// Sets the osVersion property value. Operating system version of the device
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *ManagedDeviceEncryptionState) SetOsVersion(value *string)() {
    m.osVersion = value
}
// Sets the policyDetails property value. Policy Details
// Parameters:
//  - value : Value to set for the policyDetails property.
func (m *ManagedDeviceEncryptionState) SetPolicyDetails(value []EncryptionReportPolicyDetails)() {
    m.policyDetails = value
}
// Sets the tpmSpecificationVersion property value. Device TPM Version
// Parameters:
//  - value : Value to set for the tpmSpecificationVersion property.
func (m *ManagedDeviceEncryptionState) SetTpmSpecificationVersion(value *string)() {
    m.tpmSpecificationVersion = value
}
// Sets the userPrincipalName property value. User name
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *ManagedDeviceEncryptionState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
