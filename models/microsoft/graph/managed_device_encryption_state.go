package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagedDeviceEncryptionState struct {
    Entity
    advancedBitLockerStates *AdvancedBitLockerState;
    deviceName *string;
    deviceType *DeviceTypes;
    encryptionPolicySettingState *ComplianceStatus;
    encryptionReadinessState *EncryptionReadinessState;
    encryptionState *EncryptionState;
    fileVaultStates *FileVaultState;
    osVersion *string;
    policyDetails []EncryptionReportPolicyDetails;
    tpmSpecificationVersion *string;
    userPrincipalName *string;
}
func NewManagedDeviceEncryptionState()(*ManagedDeviceEncryptionState) {
    m := &ManagedDeviceEncryptionState{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ManagedDeviceEncryptionState) GetAdvancedBitLockerStates()(*AdvancedBitLockerState) {
    if m == nil {
        return nil
    } else {
        return m.advancedBitLockerStates
    }
}
func (m *ManagedDeviceEncryptionState) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *ManagedDeviceEncryptionState) GetDeviceType()(*DeviceTypes) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
func (m *ManagedDeviceEncryptionState) GetEncryptionPolicySettingState()(*ComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.encryptionPolicySettingState
    }
}
func (m *ManagedDeviceEncryptionState) GetEncryptionReadinessState()(*EncryptionReadinessState) {
    if m == nil {
        return nil
    } else {
        return m.encryptionReadinessState
    }
}
func (m *ManagedDeviceEncryptionState) GetEncryptionState()(*EncryptionState) {
    if m == nil {
        return nil
    } else {
        return m.encryptionState
    }
}
func (m *ManagedDeviceEncryptionState) GetFileVaultStates()(*FileVaultState) {
    if m == nil {
        return nil
    } else {
        return m.fileVaultStates
    }
}
func (m *ManagedDeviceEncryptionState) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
func (m *ManagedDeviceEncryptionState) GetPolicyDetails()([]EncryptionReportPolicyDetails) {
    if m == nil {
        return nil
    } else {
        return m.policyDetails
    }
}
func (m *ManagedDeviceEncryptionState) GetTpmSpecificationVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tpmSpecificationVersion
    }
}
func (m *ManagedDeviceEncryptionState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
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
func (m *ManagedDeviceEncryptionState) SetAdvancedBitLockerStates(value *AdvancedBitLockerState)() {
    m.advancedBitLockerStates = value
}
func (m *ManagedDeviceEncryptionState) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *ManagedDeviceEncryptionState) SetDeviceType(value *DeviceTypes)() {
    m.deviceType = value
}
func (m *ManagedDeviceEncryptionState) SetEncryptionPolicySettingState(value *ComplianceStatus)() {
    m.encryptionPolicySettingState = value
}
func (m *ManagedDeviceEncryptionState) SetEncryptionReadinessState(value *EncryptionReadinessState)() {
    m.encryptionReadinessState = value
}
func (m *ManagedDeviceEncryptionState) SetEncryptionState(value *EncryptionState)() {
    m.encryptionState = value
}
func (m *ManagedDeviceEncryptionState) SetFileVaultStates(value *FileVaultState)() {
    m.fileVaultStates = value
}
func (m *ManagedDeviceEncryptionState) SetOsVersion(value *string)() {
    m.osVersion = value
}
func (m *ManagedDeviceEncryptionState) SetPolicyDetails(value []EncryptionReportPolicyDetails)() {
    m.policyDetails = value
}
func (m *ManagedDeviceEncryptionState) SetTpmSpecificationVersion(value *string)() {
    m.tpmSpecificationVersion = value
}
func (m *ManagedDeviceEncryptionState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
