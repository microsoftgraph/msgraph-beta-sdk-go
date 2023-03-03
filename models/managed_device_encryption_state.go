package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceEncryptionState encryption report per device
type ManagedDeviceEncryptionState struct {
    Entity
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
    val, err := m.GetBackingStore().Get("advancedBitLockerStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AdvancedBitLockerState)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. Device name
func (m *ManagedDeviceEncryptionState) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceType gets the deviceType property value. Device type.
func (m *ManagedDeviceEncryptionState) GetDeviceType()(*DeviceTypes) {
    val, err := m.GetBackingStore().Get("deviceType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceTypes)
    }
    return nil
}
// GetEncryptionPolicySettingState gets the encryptionPolicySettingState property value. The encryptionPolicySettingState property
func (m *ManagedDeviceEncryptionState) GetEncryptionPolicySettingState()(*ComplianceStatus) {
    val, err := m.GetBackingStore().Get("encryptionPolicySettingState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ComplianceStatus)
    }
    return nil
}
// GetEncryptionReadinessState gets the encryptionReadinessState property value. Encryption readiness state
func (m *ManagedDeviceEncryptionState) GetEncryptionReadinessState()(*EncryptionReadinessState) {
    val, err := m.GetBackingStore().Get("encryptionReadinessState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EncryptionReadinessState)
    }
    return nil
}
// GetEncryptionState gets the encryptionState property value. Encryption state
func (m *ManagedDeviceEncryptionState) GetEncryptionState()(*EncryptionState) {
    val, err := m.GetBackingStore().Get("encryptionState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EncryptionState)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceEncryptionState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["advancedBitLockerStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdvancedBitLockerState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedBitLockerStates(val.(*AdvancedBitLockerState))
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val.(*DeviceTypes))
        }
        return nil
    }
    res["encryptionPolicySettingState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionPolicySettingState(val.(*ComplianceStatus))
        }
        return nil
    }
    res["encryptionReadinessState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEncryptionReadinessState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionReadinessState(val.(*EncryptionReadinessState))
        }
        return nil
    }
    res["encryptionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEncryptionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionState(val.(*EncryptionState))
        }
        return nil
    }
    res["fileVaultStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileVaultState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultStates(val.(*FileVaultState))
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["policyDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEncryptionReportPolicyDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EncryptionReportPolicyDetailsable, len(val))
            for i, v := range val {
                res[i] = v.(EncryptionReportPolicyDetailsable)
            }
            m.SetPolicyDetails(res)
        }
        return nil
    }
    res["tpmSpecificationVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmSpecificationVersion(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetFileVaultStates gets the fileVaultStates property value. FileVault State. Possible values are: success, driveEncryptedByUser, userDeferredEncryption, escrowNotEnabled.
func (m *ManagedDeviceEncryptionState) GetFileVaultStates()(*FileVaultState) {
    val, err := m.GetBackingStore().Get("fileVaultStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FileVaultState)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. Operating system version of the device
func (m *ManagedDeviceEncryptionState) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyDetails gets the policyDetails property value. Policy Details
func (m *ManagedDeviceEncryptionState) GetPolicyDetails()([]EncryptionReportPolicyDetailsable) {
    val, err := m.GetBackingStore().Get("policyDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EncryptionReportPolicyDetailsable)
    }
    return nil
}
// GetTpmSpecificationVersion gets the tpmSpecificationVersion property value. Device TPM Version
func (m *ManagedDeviceEncryptionState) GetTpmSpecificationVersion()(*string) {
    val, err := m.GetBackingStore().Get("tpmSpecificationVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. User name
func (m *ManagedDeviceEncryptionState) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPolicyDetails()))
        for i, v := range m.GetPolicyDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    err := m.GetBackingStore().Set("advancedBitLockerStates", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. Device name
func (m *ManagedDeviceEncryptionState) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceType sets the deviceType property value. Device type.
func (m *ManagedDeviceEncryptionState) SetDeviceType(value *DeviceTypes)() {
    err := m.GetBackingStore().Set("deviceType", value)
    if err != nil {
        panic(err)
    }
}
// SetEncryptionPolicySettingState sets the encryptionPolicySettingState property value. The encryptionPolicySettingState property
func (m *ManagedDeviceEncryptionState) SetEncryptionPolicySettingState(value *ComplianceStatus)() {
    err := m.GetBackingStore().Set("encryptionPolicySettingState", value)
    if err != nil {
        panic(err)
    }
}
// SetEncryptionReadinessState sets the encryptionReadinessState property value. Encryption readiness state
func (m *ManagedDeviceEncryptionState) SetEncryptionReadinessState(value *EncryptionReadinessState)() {
    err := m.GetBackingStore().Set("encryptionReadinessState", value)
    if err != nil {
        panic(err)
    }
}
// SetEncryptionState sets the encryptionState property value. Encryption state
func (m *ManagedDeviceEncryptionState) SetEncryptionState(value *EncryptionState)() {
    err := m.GetBackingStore().Set("encryptionState", value)
    if err != nil {
        panic(err)
    }
}
// SetFileVaultStates sets the fileVaultStates property value. FileVault State. Possible values are: success, driveEncryptedByUser, userDeferredEncryption, escrowNotEnabled.
func (m *ManagedDeviceEncryptionState) SetFileVaultStates(value *FileVaultState)() {
    err := m.GetBackingStore().Set("fileVaultStates", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. Operating system version of the device
func (m *ManagedDeviceEncryptionState) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyDetails sets the policyDetails property value. Policy Details
func (m *ManagedDeviceEncryptionState) SetPolicyDetails(value []EncryptionReportPolicyDetailsable)() {
    err := m.GetBackingStore().Set("policyDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetTpmSpecificationVersion sets the tpmSpecificationVersion property value. Device TPM Version
func (m *ManagedDeviceEncryptionState) SetTpmSpecificationVersion(value *string)() {
    err := m.GetBackingStore().Set("tpmSpecificationVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User name
func (m *ManagedDeviceEncryptionState) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// ManagedDeviceEncryptionStateable 
type ManagedDeviceEncryptionStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedBitLockerStates()(*AdvancedBitLockerState)
    GetDeviceName()(*string)
    GetDeviceType()(*DeviceTypes)
    GetEncryptionPolicySettingState()(*ComplianceStatus)
    GetEncryptionReadinessState()(*EncryptionReadinessState)
    GetEncryptionState()(*EncryptionState)
    GetFileVaultStates()(*FileVaultState)
    GetOsVersion()(*string)
    GetPolicyDetails()([]EncryptionReportPolicyDetailsable)
    GetTpmSpecificationVersion()(*string)
    GetUserPrincipalName()(*string)
    SetAdvancedBitLockerStates(value *AdvancedBitLockerState)()
    SetDeviceName(value *string)()
    SetDeviceType(value *DeviceTypes)()
    SetEncryptionPolicySettingState(value *ComplianceStatus)()
    SetEncryptionReadinessState(value *EncryptionReadinessState)()
    SetEncryptionState(value *EncryptionState)()
    SetFileVaultStates(value *FileVaultState)()
    SetOsVersion(value *string)()
    SetPolicyDetails(value []EncryptionReportPolicyDetailsable)()
    SetTpmSpecificationVersion(value *string)()
    SetUserPrincipalName(value *string)()
}
