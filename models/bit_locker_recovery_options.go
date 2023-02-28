package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// BitLockerRecoveryOptions bitLocker Recovery Options.
type BitLockerRecoveryOptions struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewBitLockerRecoveryOptions instantiates a new bitLockerRecoveryOptions and sets the default values.
func NewBitLockerRecoveryOptions()(*BitLockerRecoveryOptions) {
    m := &BitLockerRecoveryOptions{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBitLockerRecoveryOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBitLockerRecoveryOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBitLockerRecoveryOptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BitLockerRecoveryOptions) GetAdditionalData()(map[string]any) {
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
func (m *BitLockerRecoveryOptions) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBlockDataRecoveryAgent gets the blockDataRecoveryAgent property value. Indicates whether to block certificate-based data recovery agent.
func (m *BitLockerRecoveryOptions) GetBlockDataRecoveryAgent()(*bool) {
    val, err := m.GetBackingStore().Get("blockDataRecoveryAgent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableBitLockerAfterRecoveryInformationToStore gets the enableBitLockerAfterRecoveryInformationToStore property value. Indicates whether or not to enable BitLocker until recovery information is stored in AD DS.
func (m *BitLockerRecoveryOptions) GetEnableBitLockerAfterRecoveryInformationToStore()(*bool) {
    val, err := m.GetBackingStore().Get("enableBitLockerAfterRecoveryInformationToStore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableRecoveryInformationSaveToStore gets the enableRecoveryInformationSaveToStore property value. Indicates whether or not to allow BitLocker recovery information to store in AD DS.
func (m *BitLockerRecoveryOptions) GetEnableRecoveryInformationSaveToStore()(*bool) {
    val, err := m.GetBackingStore().Get("enableRecoveryInformationSaveToStore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BitLockerRecoveryOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blockDataRecoveryAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockDataRecoveryAgent(val)
        }
        return nil
    }
    res["enableBitLockerAfterRecoveryInformationToStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableBitLockerAfterRecoveryInformationToStore(val)
        }
        return nil
    }
    res["enableRecoveryInformationSaveToStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableRecoveryInformationSaveToStore(val)
        }
        return nil
    }
    res["hideRecoveryOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideRecoveryOptions(val)
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
    res["recoveryInformationToStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBitLockerRecoveryInformationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecoveryInformationToStore(val.(*BitLockerRecoveryInformationType))
        }
        return nil
    }
    res["recoveryKeyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecoveryKeyUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["recoveryPasswordUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecoveryPasswordUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    return res
}
// GetHideRecoveryOptions gets the hideRecoveryOptions property value. Indicates whether or not to allow showing recovery options in BitLocker Setup Wizard for fixed or system disk.
func (m *BitLockerRecoveryOptions) GetHideRecoveryOptions()(*bool) {
    val, err := m.GetBackingStore().Get("hideRecoveryOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BitLockerRecoveryOptions) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecoveryInformationToStore gets the recoveryInformationToStore property value. BitLockerRecoveryInformationType types
func (m *BitLockerRecoveryOptions) GetRecoveryInformationToStore()(*BitLockerRecoveryInformationType) {
    val, err := m.GetBackingStore().Get("recoveryInformationToStore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*BitLockerRecoveryInformationType)
    }
    return nil
}
// GetRecoveryKeyUsage gets the recoveryKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerRecoveryOptions) GetRecoveryKeyUsage()(*ConfigurationUsage) {
    val, err := m.GetBackingStore().Get("recoveryKeyUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConfigurationUsage)
    }
    return nil
}
// GetRecoveryPasswordUsage gets the recoveryPasswordUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerRecoveryOptions) GetRecoveryPasswordUsage()(*ConfigurationUsage) {
    val, err := m.GetBackingStore().Get("recoveryPasswordUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConfigurationUsage)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BitLockerRecoveryOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("blockDataRecoveryAgent", m.GetBlockDataRecoveryAgent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableBitLockerAfterRecoveryInformationToStore", m.GetEnableBitLockerAfterRecoveryInformationToStore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableRecoveryInformationSaveToStore", m.GetEnableRecoveryInformationSaveToStore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideRecoveryOptions", m.GetHideRecoveryOptions())
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
    if m.GetRecoveryInformationToStore() != nil {
        cast := (*m.GetRecoveryInformationToStore()).String()
        err := writer.WriteStringValue("recoveryInformationToStore", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecoveryKeyUsage() != nil {
        cast := (*m.GetRecoveryKeyUsage()).String()
        err := writer.WriteStringValue("recoveryKeyUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecoveryPasswordUsage() != nil {
        cast := (*m.GetRecoveryPasswordUsage()).String()
        err := writer.WriteStringValue("recoveryPasswordUsage", &cast)
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
func (m *BitLockerRecoveryOptions) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *BitLockerRecoveryOptions) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBlockDataRecoveryAgent sets the blockDataRecoveryAgent property value. Indicates whether to block certificate-based data recovery agent.
func (m *BitLockerRecoveryOptions) SetBlockDataRecoveryAgent(value *bool)() {
    err := m.GetBackingStore().Set("blockDataRecoveryAgent", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableBitLockerAfterRecoveryInformationToStore sets the enableBitLockerAfterRecoveryInformationToStore property value. Indicates whether or not to enable BitLocker until recovery information is stored in AD DS.
func (m *BitLockerRecoveryOptions) SetEnableBitLockerAfterRecoveryInformationToStore(value *bool)() {
    err := m.GetBackingStore().Set("enableBitLockerAfterRecoveryInformationToStore", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableRecoveryInformationSaveToStore sets the enableRecoveryInformationSaveToStore property value. Indicates whether or not to allow BitLocker recovery information to store in AD DS.
func (m *BitLockerRecoveryOptions) SetEnableRecoveryInformationSaveToStore(value *bool)() {
    err := m.GetBackingStore().Set("enableRecoveryInformationSaveToStore", value)
    if err != nil {
        panic(err)
    }
}
// SetHideRecoveryOptions sets the hideRecoveryOptions property value. Indicates whether or not to allow showing recovery options in BitLocker Setup Wizard for fixed or system disk.
func (m *BitLockerRecoveryOptions) SetHideRecoveryOptions(value *bool)() {
    err := m.GetBackingStore().Set("hideRecoveryOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BitLockerRecoveryOptions) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRecoveryInformationToStore sets the recoveryInformationToStore property value. BitLockerRecoveryInformationType types
func (m *BitLockerRecoveryOptions) SetRecoveryInformationToStore(value *BitLockerRecoveryInformationType)() {
    err := m.GetBackingStore().Set("recoveryInformationToStore", value)
    if err != nil {
        panic(err)
    }
}
// SetRecoveryKeyUsage sets the recoveryKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerRecoveryOptions) SetRecoveryKeyUsage(value *ConfigurationUsage)() {
    err := m.GetBackingStore().Set("recoveryKeyUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetRecoveryPasswordUsage sets the recoveryPasswordUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerRecoveryOptions) SetRecoveryPasswordUsage(value *ConfigurationUsage)() {
    err := m.GetBackingStore().Set("recoveryPasswordUsage", value)
    if err != nil {
        panic(err)
    }
}
// BitLockerRecoveryOptionsable 
type BitLockerRecoveryOptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBlockDataRecoveryAgent()(*bool)
    GetEnableBitLockerAfterRecoveryInformationToStore()(*bool)
    GetEnableRecoveryInformationSaveToStore()(*bool)
    GetHideRecoveryOptions()(*bool)
    GetOdataType()(*string)
    GetRecoveryInformationToStore()(*BitLockerRecoveryInformationType)
    GetRecoveryKeyUsage()(*ConfigurationUsage)
    GetRecoveryPasswordUsage()(*ConfigurationUsage)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBlockDataRecoveryAgent(value *bool)()
    SetEnableBitLockerAfterRecoveryInformationToStore(value *bool)()
    SetEnableRecoveryInformationSaveToStore(value *bool)()
    SetHideRecoveryOptions(value *bool)()
    SetOdataType(value *string)()
    SetRecoveryInformationToStore(value *BitLockerRecoveryInformationType)()
    SetRecoveryKeyUsage(value *ConfigurationUsage)()
    SetRecoveryPasswordUsage(value *ConfigurationUsage)()
}
