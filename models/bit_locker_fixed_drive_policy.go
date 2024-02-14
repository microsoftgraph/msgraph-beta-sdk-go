package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// BitLockerFixedDrivePolicy bitLocker Fixed Drive Policies.
type BitLockerFixedDrivePolicy struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewBitLockerFixedDrivePolicy instantiates a new BitLockerFixedDrivePolicy and sets the default values.
func NewBitLockerFixedDrivePolicy()(*BitLockerFixedDrivePolicy) {
    m := &BitLockerFixedDrivePolicy{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBitLockerFixedDrivePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBitLockerFixedDrivePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBitLockerFixedDrivePolicy(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BitLockerFixedDrivePolicy) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *BitLockerFixedDrivePolicy) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEncryptionMethod gets the encryptionMethod property value. Select the encryption method for fixed drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
// returns a *BitLockerEncryptionMethod when successful
func (m *BitLockerFixedDrivePolicy) GetEncryptionMethod()(*BitLockerEncryptionMethod) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BitLockerFixedDrivePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["requireEncryptionForWriteAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireEncryptionForWriteAccess(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *BitLockerFixedDrivePolicy) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecoveryOptions gets the recoveryOptions property value. This policy setting allows you to control how BitLocker-protected fixed data drives are recovered in the absence of the required credentials. This policy setting is applied when you turn on BitLocker.
// returns a BitLockerRecoveryOptionsable when successful
func (m *BitLockerFixedDrivePolicy) GetRecoveryOptions()(BitLockerRecoveryOptionsable) {
    val, err := m.GetBackingStore().Get("recoveryOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(BitLockerRecoveryOptionsable)
    }
    return nil
}
// GetRequireEncryptionForWriteAccess gets the requireEncryptionForWriteAccess property value. This policy setting determines whether BitLocker protection is required for fixed data drives to be writable on a computer.
// returns a *bool when successful
func (m *BitLockerFixedDrivePolicy) GetRequireEncryptionForWriteAccess()(*bool) {
    val, err := m.GetBackingStore().Get("requireEncryptionForWriteAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BitLockerFixedDrivePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEncryptionMethod() != nil {
        cast := (*m.GetEncryptionMethod()).String()
        err := writer.WriteStringValue("encryptionMethod", &cast)
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
        err := writer.WriteObjectValue("recoveryOptions", m.GetRecoveryOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("requireEncryptionForWriteAccess", m.GetRequireEncryptionForWriteAccess())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BitLockerFixedDrivePolicy) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *BitLockerFixedDrivePolicy) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEncryptionMethod sets the encryptionMethod property value. Select the encryption method for fixed drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
func (m *BitLockerFixedDrivePolicy) SetEncryptionMethod(value *BitLockerEncryptionMethod)() {
    err := m.GetBackingStore().Set("encryptionMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BitLockerFixedDrivePolicy) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRecoveryOptions sets the recoveryOptions property value. This policy setting allows you to control how BitLocker-protected fixed data drives are recovered in the absence of the required credentials. This policy setting is applied when you turn on BitLocker.
func (m *BitLockerFixedDrivePolicy) SetRecoveryOptions(value BitLockerRecoveryOptionsable)() {
    err := m.GetBackingStore().Set("recoveryOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireEncryptionForWriteAccess sets the requireEncryptionForWriteAccess property value. This policy setting determines whether BitLocker protection is required for fixed data drives to be writable on a computer.
func (m *BitLockerFixedDrivePolicy) SetRequireEncryptionForWriteAccess(value *bool)() {
    err := m.GetBackingStore().Set("requireEncryptionForWriteAccess", value)
    if err != nil {
        panic(err)
    }
}
type BitLockerFixedDrivePolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEncryptionMethod()(*BitLockerEncryptionMethod)
    GetOdataType()(*string)
    GetRecoveryOptions()(BitLockerRecoveryOptionsable)
    GetRequireEncryptionForWriteAccess()(*bool)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEncryptionMethod(value *BitLockerEncryptionMethod)()
    SetOdataType(value *string)()
    SetRecoveryOptions(value BitLockerRecoveryOptionsable)()
    SetRequireEncryptionForWriteAccess(value *bool)()
}
