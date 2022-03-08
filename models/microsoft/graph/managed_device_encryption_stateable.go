package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedDeviceEncryptionStateable 
type ManagedDeviceEncryptionStateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
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
