package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceEncryptionState encryption report per device
type ManagedDeviceEncryptionState struct {
    Entity
    // Advanced BitLocker State. Possible values are: success, noUserConsent, osVolumeUnprotected, osVolumeTpmRequired, osVolumeTpmOnlyRequired, osVolumeTpmPinRequired, osVolumeTpmStartupKeyRequired, osVolumeTpmPinStartupKeyRequired, osVolumeEncryptionMethodMismatch, recoveryKeyBackupFailed, fixedDriveNotEncrypted, fixedDriveEncryptionMethodMismatch, loggedOnUserNonAdmin, windowsRecoveryEnvironmentNotConfigured, tpmNotAvailable, tpmNotReady, networkError.
    advancedBitLockerStates *AdvancedBitLockerState
    // Device name
    deviceName *string
    // Device type.
    deviceType *DeviceTypes
    // The encryptionPolicySettingState property
    encryptionPolicySettingState *ComplianceStatus
    // Encryption readiness state
    encryptionReadinessState *EncryptionReadinessState
    // Encryption state
    encryptionState *EncryptionState
    // FileVault State. Possible values are: success, driveEncryptedByUser, userDeferredEncryption, escrowNotEnabled.
    fileVaultStates *FileVaultState
    // Operating system version of the device
    osVersion *string
    // Policy Details
    policyDetails []EncryptionReportPolicyDetailsable
    // Device TPM Version
    tpmSpecificationVersion *string
    // User name
    userPrincipalName *string
}
// NewManagedDeviceEncryptionState instantiates a new managedDeviceEncryptionState and sets the default values.
func NewManagedDeviceEncryptionState()(*ManagedDeviceEncryptionState) {
    m := &ManagedDeviceEncryptionState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedDeviceEncryptionStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceEncryptionStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceEncryptionState(), nil
}
// GetAdvancedBitLockerStates gets the advancedBitLockerStates property value. Advanced BitLocker State. Possible values are: success, noUserConsent, osVolumeUnprotected, osVolumeTpmRequired, osVolumeTpmOnlyRequired, osVolumeTpmPinRequired, osVolumeTpmStartupKeyRequired, osVolumeTpmPinStartupKeyRequired, osVolumeEncryptionMethodMismatch, recoveryKeyBackupFailed, fixedDriveNotEncrypted, fixedDriveEncryptionMethodMismatch, loggedOnUserNonAdmin, windowsRecoveryEnvironmentNotConfigured, tpmNotAvailable, tpmNotReady, networkError.
func (m *ManagedDeviceEncryptionState) GetAdvancedBitLockerStates()(*AdvancedBitLockerState) {
    return m.advancedBitLockerStates
}
// GetDeviceName gets the deviceName property value. Device name
func (m *ManagedDeviceEncryptionState) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDeviceType gets the deviceType property value. Device type.
func (m *ManagedDeviceEncryptionState) GetDeviceType()(*DeviceTypes) {
    return m.deviceType
}
// GetEncryptionPolicySettingState gets the encryptionPolicySettingState property value. The encryptionPolicySettingState property
func (m *ManagedDeviceEncryptionState) GetEncryptionPolicySettingState()(*ComplianceStatus) {
    return m.encryptionPolicySettingState
}
// GetEncryptionReadinessState gets the encryptionReadinessState property value. Encryption readiness state
func (m *ManagedDeviceEncryptionState) GetEncryptionReadinessState()(*EncryptionReadinessState) {
    return m.encryptionReadinessState
}
// GetEncryptionState gets the encryptionState property value. Encryption state
func (m *ManagedDeviceEncryptionState) GetEncryptionState()(*EncryptionState) {
    return m.encryptionState
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceEncryptionState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["advancedBitLockerStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAdvancedBitLockerState , m.SetAdvancedBitLockerStates)
    res["deviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceName)
    res["deviceType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceTypes , m.SetDeviceType)
    res["encryptionPolicySettingState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseComplianceStatus , m.SetEncryptionPolicySettingState)
    res["encryptionReadinessState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEncryptionReadinessState , m.SetEncryptionReadinessState)
    res["encryptionState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEncryptionState , m.SetEncryptionState)
    res["fileVaultStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseFileVaultState , m.SetFileVaultStates)
    res["osVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOsVersion)
    res["policyDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEncryptionReportPolicyDetailsFromDiscriminatorValue , m.SetPolicyDetails)
    res["tpmSpecificationVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTpmSpecificationVersion)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    return res
}
// GetFileVaultStates gets the fileVaultStates property value. FileVault State. Possible values are: success, driveEncryptedByUser, userDeferredEncryption, escrowNotEnabled.
func (m *ManagedDeviceEncryptionState) GetFileVaultStates()(*FileVaultState) {
    return m.fileVaultStates
}
// GetOsVersion gets the osVersion property value. Operating system version of the device
func (m *ManagedDeviceEncryptionState) GetOsVersion()(*string) {
    return m.osVersion
}
// GetPolicyDetails gets the policyDetails property value. Policy Details
func (m *ManagedDeviceEncryptionState) GetPolicyDetails()([]EncryptionReportPolicyDetailsable) {
    return m.policyDetails
}
// GetTpmSpecificationVersion gets the tpmSpecificationVersion property value. Device TPM Version
func (m *ManagedDeviceEncryptionState) GetTpmSpecificationVersion()(*string) {
    return m.tpmSpecificationVersion
}
// GetUserPrincipalName gets the userPrincipalName property value. User name
func (m *ManagedDeviceEncryptionState) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *ManagedDeviceEncryptionState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPolicyDetails())
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
    m.advancedBitLockerStates = value
}
// SetDeviceName sets the deviceName property value. Device name
func (m *ManagedDeviceEncryptionState) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDeviceType sets the deviceType property value. Device type.
func (m *ManagedDeviceEncryptionState) SetDeviceType(value *DeviceTypes)() {
    m.deviceType = value
}
// SetEncryptionPolicySettingState sets the encryptionPolicySettingState property value. The encryptionPolicySettingState property
func (m *ManagedDeviceEncryptionState) SetEncryptionPolicySettingState(value *ComplianceStatus)() {
    m.encryptionPolicySettingState = value
}
// SetEncryptionReadinessState sets the encryptionReadinessState property value. Encryption readiness state
func (m *ManagedDeviceEncryptionState) SetEncryptionReadinessState(value *EncryptionReadinessState)() {
    m.encryptionReadinessState = value
}
// SetEncryptionState sets the encryptionState property value. Encryption state
func (m *ManagedDeviceEncryptionState) SetEncryptionState(value *EncryptionState)() {
    m.encryptionState = value
}
// SetFileVaultStates sets the fileVaultStates property value. FileVault State. Possible values are: success, driveEncryptedByUser, userDeferredEncryption, escrowNotEnabled.
func (m *ManagedDeviceEncryptionState) SetFileVaultStates(value *FileVaultState)() {
    m.fileVaultStates = value
}
// SetOsVersion sets the osVersion property value. Operating system version of the device
func (m *ManagedDeviceEncryptionState) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetPolicyDetails sets the policyDetails property value. Policy Details
func (m *ManagedDeviceEncryptionState) SetPolicyDetails(value []EncryptionReportPolicyDetailsable)() {
    m.policyDetails = value
}
// SetTpmSpecificationVersion sets the tpmSpecificationVersion property value. Device TPM Version
func (m *ManagedDeviceEncryptionState) SetTpmSpecificationVersion(value *string)() {
    m.tpmSpecificationVersion = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User name
func (m *ManagedDeviceEncryptionState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
