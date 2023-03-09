package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// BitLockerSystemDrivePolicy bitLocker Encryption Base Policies.
type BitLockerSystemDrivePolicy struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewBitLockerSystemDrivePolicy instantiates a new bitLockerSystemDrivePolicy and sets the default values.
func NewBitLockerSystemDrivePolicy()(*BitLockerSystemDrivePolicy) {
    m := &BitLockerSystemDrivePolicy{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBitLockerSystemDrivePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBitLockerSystemDrivePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBitLockerSystemDrivePolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BitLockerSystemDrivePolicy) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *BitLockerSystemDrivePolicy) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEncryptionMethod gets the encryptionMethod property value. Select the encryption method for operating system drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
func (m *BitLockerSystemDrivePolicy) GetEncryptionMethod()(*BitLockerEncryptionMethod) {
    val, err := m.GetBackingStore().Get("encryptionMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*BitLockerEncryptionMethod)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BitLockerSystemDrivePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["encryptionMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBitLockerEncryptionMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionMethod(val.(*BitLockerEncryptionMethod))
        }
        return nil
    }
    res["minimumPinLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumPinLength(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["prebootRecoveryEnableMessageAndUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrebootRecoveryEnableMessageAndUrl(val)
        }
        return nil
    }
    res["prebootRecoveryMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrebootRecoveryMessage(val)
        }
        return nil
    }
    res["prebootRecoveryUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrebootRecoveryUrl(val)
        }
        return nil
    }
    res["recoveryOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitLockerRecoveryOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecoveryOptions(val.(BitLockerRecoveryOptionsable))
        }
        return nil
    }
    res["startupAuthenticationBlockWithoutTpmChip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupAuthenticationBlockWithoutTpmChip(val)
        }
        return nil
    }
    res["startupAuthenticationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupAuthenticationRequired(val)
        }
        return nil
    }
    res["startupAuthenticationTpmKeyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupAuthenticationTpmKeyUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["startupAuthenticationTpmPinAndKeyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupAuthenticationTpmPinAndKeyUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["startupAuthenticationTpmPinUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupAuthenticationTpmPinUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["startupAuthenticationTpmUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupAuthenticationTpmUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    return res
}
// GetMinimumPinLength gets the minimumPinLength property value. Indicates the minimum length of startup pin. Valid values 4 to 20
func (m *BitLockerSystemDrivePolicy) GetMinimumPinLength()(*int32) {
    val, err := m.GetBackingStore().Get("minimumPinLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BitLockerSystemDrivePolicy) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrebootRecoveryEnableMessageAndUrl gets the prebootRecoveryEnableMessageAndUrl property value. Enable pre-boot recovery message and Url. If requireStartupAuthentication is false, this value does not affect.
func (m *BitLockerSystemDrivePolicy) GetPrebootRecoveryEnableMessageAndUrl()(*bool) {
    val, err := m.GetBackingStore().Get("prebootRecoveryEnableMessageAndUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPrebootRecoveryMessage gets the prebootRecoveryMessage property value. Defines a custom recovery message.
func (m *BitLockerSystemDrivePolicy) GetPrebootRecoveryMessage()(*string) {
    val, err := m.GetBackingStore().Get("prebootRecoveryMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrebootRecoveryUrl gets the prebootRecoveryUrl property value. Defines a custom recovery URL.
func (m *BitLockerSystemDrivePolicy) GetPrebootRecoveryUrl()(*string) {
    val, err := m.GetBackingStore().Get("prebootRecoveryUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecoveryOptions gets the recoveryOptions property value. Allows to recover BitLocker encrypted operating system drives in the absence of the required startup key information. This policy setting is applied when you turn on BitLocker.
func (m *BitLockerSystemDrivePolicy) GetRecoveryOptions()(BitLockerRecoveryOptionsable) {
    val, err := m.GetBackingStore().Get("recoveryOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(BitLockerRecoveryOptionsable)
    }
    return nil
}
// GetStartupAuthenticationBlockWithoutTpmChip gets the startupAuthenticationBlockWithoutTpmChip property value. Indicates whether to allow BitLocker without a compatible TPM (requires a password or a startup key on a USB flash drive).
func (m *BitLockerSystemDrivePolicy) GetStartupAuthenticationBlockWithoutTpmChip()(*bool) {
    val, err := m.GetBackingStore().Get("startupAuthenticationBlockWithoutTpmChip")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStartupAuthenticationRequired gets the startupAuthenticationRequired property value. Require additional authentication at startup.
func (m *BitLockerSystemDrivePolicy) GetStartupAuthenticationRequired()(*bool) {
    val, err := m.GetBackingStore().Get("startupAuthenticationRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStartupAuthenticationTpmKeyUsage gets the startupAuthenticationTpmKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) GetStartupAuthenticationTpmKeyUsage()(*ConfigurationUsage) {
    val, err := m.GetBackingStore().Get("startupAuthenticationTpmKeyUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConfigurationUsage)
    }
    return nil
}
// GetStartupAuthenticationTpmPinAndKeyUsage gets the startupAuthenticationTpmPinAndKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) GetStartupAuthenticationTpmPinAndKeyUsage()(*ConfigurationUsage) {
    val, err := m.GetBackingStore().Get("startupAuthenticationTpmPinAndKeyUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConfigurationUsage)
    }
    return nil
}
// GetStartupAuthenticationTpmPinUsage gets the startupAuthenticationTpmPinUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) GetStartupAuthenticationTpmPinUsage()(*ConfigurationUsage) {
    val, err := m.GetBackingStore().Get("startupAuthenticationTpmPinUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConfigurationUsage)
    }
    return nil
}
// GetStartupAuthenticationTpmUsage gets the startupAuthenticationTpmUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) GetStartupAuthenticationTpmUsage()(*ConfigurationUsage) {
    val, err := m.GetBackingStore().Get("startupAuthenticationTpmUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConfigurationUsage)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BitLockerSystemDrivePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEncryptionMethod() != nil {
        cast := (*m.GetEncryptionMethod()).String()
        err := writer.WriteStringValue("encryptionMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minimumPinLength", m.GetMinimumPinLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("prebootRecoveryEnableMessageAndUrl", m.GetPrebootRecoveryEnableMessageAndUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("prebootRecoveryMessage", m.GetPrebootRecoveryMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("prebootRecoveryUrl", m.GetPrebootRecoveryUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recoveryOptions", m.GetRecoveryOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("startupAuthenticationBlockWithoutTpmChip", m.GetStartupAuthenticationBlockWithoutTpmChip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("startupAuthenticationRequired", m.GetStartupAuthenticationRequired())
        if err != nil {
            return err
        }
    }
    if m.GetStartupAuthenticationTpmKeyUsage() != nil {
        cast := (*m.GetStartupAuthenticationTpmKeyUsage()).String()
        err := writer.WriteStringValue("startupAuthenticationTpmKeyUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartupAuthenticationTpmPinAndKeyUsage() != nil {
        cast := (*m.GetStartupAuthenticationTpmPinAndKeyUsage()).String()
        err := writer.WriteStringValue("startupAuthenticationTpmPinAndKeyUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartupAuthenticationTpmPinUsage() != nil {
        cast := (*m.GetStartupAuthenticationTpmPinUsage()).String()
        err := writer.WriteStringValue("startupAuthenticationTpmPinUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartupAuthenticationTpmUsage() != nil {
        cast := (*m.GetStartupAuthenticationTpmUsage()).String()
        err := writer.WriteStringValue("startupAuthenticationTpmUsage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BitLockerSystemDrivePolicy) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *BitLockerSystemDrivePolicy) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEncryptionMethod sets the encryptionMethod property value. Select the encryption method for operating system drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
func (m *BitLockerSystemDrivePolicy) SetEncryptionMethod(value *BitLockerEncryptionMethod)() {
    err := m.GetBackingStore().Set("encryptionMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumPinLength sets the minimumPinLength property value. Indicates the minimum length of startup pin. Valid values 4 to 20
func (m *BitLockerSystemDrivePolicy) SetMinimumPinLength(value *int32)() {
    err := m.GetBackingStore().Set("minimumPinLength", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BitLockerSystemDrivePolicy) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPrebootRecoveryEnableMessageAndUrl sets the prebootRecoveryEnableMessageAndUrl property value. Enable pre-boot recovery message and Url. If requireStartupAuthentication is false, this value does not affect.
func (m *BitLockerSystemDrivePolicy) SetPrebootRecoveryEnableMessageAndUrl(value *bool)() {
    err := m.GetBackingStore().Set("prebootRecoveryEnableMessageAndUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetPrebootRecoveryMessage sets the prebootRecoveryMessage property value. Defines a custom recovery message.
func (m *BitLockerSystemDrivePolicy) SetPrebootRecoveryMessage(value *string)() {
    err := m.GetBackingStore().Set("prebootRecoveryMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetPrebootRecoveryUrl sets the prebootRecoveryUrl property value. Defines a custom recovery URL.
func (m *BitLockerSystemDrivePolicy) SetPrebootRecoveryUrl(value *string)() {
    err := m.GetBackingStore().Set("prebootRecoveryUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetRecoveryOptions sets the recoveryOptions property value. Allows to recover BitLocker encrypted operating system drives in the absence of the required startup key information. This policy setting is applied when you turn on BitLocker.
func (m *BitLockerSystemDrivePolicy) SetRecoveryOptions(value BitLockerRecoveryOptionsable)() {
    err := m.GetBackingStore().Set("recoveryOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetStartupAuthenticationBlockWithoutTpmChip sets the startupAuthenticationBlockWithoutTpmChip property value. Indicates whether to allow BitLocker without a compatible TPM (requires a password or a startup key on a USB flash drive).
func (m *BitLockerSystemDrivePolicy) SetStartupAuthenticationBlockWithoutTpmChip(value *bool)() {
    err := m.GetBackingStore().Set("startupAuthenticationBlockWithoutTpmChip", value)
    if err != nil {
        panic(err)
    }
}
// SetStartupAuthenticationRequired sets the startupAuthenticationRequired property value. Require additional authentication at startup.
func (m *BitLockerSystemDrivePolicy) SetStartupAuthenticationRequired(value *bool)() {
    err := m.GetBackingStore().Set("startupAuthenticationRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetStartupAuthenticationTpmKeyUsage sets the startupAuthenticationTpmKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) SetStartupAuthenticationTpmKeyUsage(value *ConfigurationUsage)() {
    err := m.GetBackingStore().Set("startupAuthenticationTpmKeyUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetStartupAuthenticationTpmPinAndKeyUsage sets the startupAuthenticationTpmPinAndKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) SetStartupAuthenticationTpmPinAndKeyUsage(value *ConfigurationUsage)() {
    err := m.GetBackingStore().Set("startupAuthenticationTpmPinAndKeyUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetStartupAuthenticationTpmPinUsage sets the startupAuthenticationTpmPinUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) SetStartupAuthenticationTpmPinUsage(value *ConfigurationUsage)() {
    err := m.GetBackingStore().Set("startupAuthenticationTpmPinUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetStartupAuthenticationTpmUsage sets the startupAuthenticationTpmUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) SetStartupAuthenticationTpmUsage(value *ConfigurationUsage)() {
    err := m.GetBackingStore().Set("startupAuthenticationTpmUsage", value)
    if err != nil {
        panic(err)
    }
}
// BitLockerSystemDrivePolicyable 
type BitLockerSystemDrivePolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEncryptionMethod()(*BitLockerEncryptionMethod)
    GetMinimumPinLength()(*int32)
    GetOdataType()(*string)
    GetPrebootRecoveryEnableMessageAndUrl()(*bool)
    GetPrebootRecoveryMessage()(*string)
    GetPrebootRecoveryUrl()(*string)
    GetRecoveryOptions()(BitLockerRecoveryOptionsable)
    GetStartupAuthenticationBlockWithoutTpmChip()(*bool)
    GetStartupAuthenticationRequired()(*bool)
    GetStartupAuthenticationTpmKeyUsage()(*ConfigurationUsage)
    GetStartupAuthenticationTpmPinAndKeyUsage()(*ConfigurationUsage)
    GetStartupAuthenticationTpmPinUsage()(*ConfigurationUsage)
    GetStartupAuthenticationTpmUsage()(*ConfigurationUsage)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEncryptionMethod(value *BitLockerEncryptionMethod)()
    SetMinimumPinLength(value *int32)()
    SetOdataType(value *string)()
    SetPrebootRecoveryEnableMessageAndUrl(value *bool)()
    SetPrebootRecoveryMessage(value *string)()
    SetPrebootRecoveryUrl(value *string)()
    SetRecoveryOptions(value BitLockerRecoveryOptionsable)()
    SetStartupAuthenticationBlockWithoutTpmChip(value *bool)()
    SetStartupAuthenticationRequired(value *bool)()
    SetStartupAuthenticationTpmKeyUsage(value *ConfigurationUsage)()
    SetStartupAuthenticationTpmPinAndKeyUsage(value *ConfigurationUsage)()
    SetStartupAuthenticationTpmPinUsage(value *ConfigurationUsage)()
    SetStartupAuthenticationTpmUsage(value *ConfigurationUsage)()
}
