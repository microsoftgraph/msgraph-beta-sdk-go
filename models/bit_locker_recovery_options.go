package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BitLockerRecoveryOptions bitLocker Recovery Options.
type BitLockerRecoveryOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether to block certificate-based data recovery agent.
    blockDataRecoveryAgent *bool
    // Indicates whether or not to enable BitLocker until recovery information is stored in AD DS.
    enableBitLockerAfterRecoveryInformationToStore *bool
    // Indicates whether or not to allow BitLocker recovery information to store in AD DS.
    enableRecoveryInformationSaveToStore *bool
    // Indicates whether or not to allow showing recovery options in BitLocker Setup Wizard for fixed or system disk.
    hideRecoveryOptions *bool
    // The OdataType property
    odataType *string
    // BitLockerRecoveryInformationType types
    recoveryInformationToStore *BitLockerRecoveryInformationType
    // Possible values of the ConfigurationUsage list.
    recoveryKeyUsage *ConfigurationUsage
    // Possible values of the ConfigurationUsage list.
    recoveryPasswordUsage *ConfigurationUsage
}
// NewBitLockerRecoveryOptions instantiates a new bitLockerRecoveryOptions and sets the default values.
func NewBitLockerRecoveryOptions()(*BitLockerRecoveryOptions) {
    m := &BitLockerRecoveryOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBitLockerRecoveryOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBitLockerRecoveryOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBitLockerRecoveryOptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BitLockerRecoveryOptions) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBlockDataRecoveryAgent gets the blockDataRecoveryAgent property value. Indicates whether to block certificate-based data recovery agent.
func (m *BitLockerRecoveryOptions) GetBlockDataRecoveryAgent()(*bool) {
    return m.blockDataRecoveryAgent
}
// GetEnableBitLockerAfterRecoveryInformationToStore gets the enableBitLockerAfterRecoveryInformationToStore property value. Indicates whether or not to enable BitLocker until recovery information is stored in AD DS.
func (m *BitLockerRecoveryOptions) GetEnableBitLockerAfterRecoveryInformationToStore()(*bool) {
    return m.enableBitLockerAfterRecoveryInformationToStore
}
// GetEnableRecoveryInformationSaveToStore gets the enableRecoveryInformationSaveToStore property value. Indicates whether or not to allow BitLocker recovery information to store in AD DS.
func (m *BitLockerRecoveryOptions) GetEnableRecoveryInformationSaveToStore()(*bool) {
    return m.enableRecoveryInformationSaveToStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BitLockerRecoveryOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blockDataRecoveryAgent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBlockDataRecoveryAgent)
    res["enableBitLockerAfterRecoveryInformationToStore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnableBitLockerAfterRecoveryInformationToStore)
    res["enableRecoveryInformationSaveToStore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnableRecoveryInformationSaveToStore)
    res["hideRecoveryOptions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetHideRecoveryOptions)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["recoveryInformationToStore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseBitLockerRecoveryInformationType , m.SetRecoveryInformationToStore)
    res["recoveryKeyUsage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseConfigurationUsage , m.SetRecoveryKeyUsage)
    res["recoveryPasswordUsage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseConfigurationUsage , m.SetRecoveryPasswordUsage)
    return res
}
// GetHideRecoveryOptions gets the hideRecoveryOptions property value. Indicates whether or not to allow showing recovery options in BitLocker Setup Wizard for fixed or system disk.
func (m *BitLockerRecoveryOptions) GetHideRecoveryOptions()(*bool) {
    return m.hideRecoveryOptions
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BitLockerRecoveryOptions) GetOdataType()(*string) {
    return m.odataType
}
// GetRecoveryInformationToStore gets the recoveryInformationToStore property value. BitLockerRecoveryInformationType types
func (m *BitLockerRecoveryOptions) GetRecoveryInformationToStore()(*BitLockerRecoveryInformationType) {
    return m.recoveryInformationToStore
}
// GetRecoveryKeyUsage gets the recoveryKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerRecoveryOptions) GetRecoveryKeyUsage()(*ConfigurationUsage) {
    return m.recoveryKeyUsage
}
// GetRecoveryPasswordUsage gets the recoveryPasswordUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerRecoveryOptions) GetRecoveryPasswordUsage()(*ConfigurationUsage) {
    return m.recoveryPasswordUsage
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
func (m *BitLockerRecoveryOptions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBlockDataRecoveryAgent sets the blockDataRecoveryAgent property value. Indicates whether to block certificate-based data recovery agent.
func (m *BitLockerRecoveryOptions) SetBlockDataRecoveryAgent(value *bool)() {
    m.blockDataRecoveryAgent = value
}
// SetEnableBitLockerAfterRecoveryInformationToStore sets the enableBitLockerAfterRecoveryInformationToStore property value. Indicates whether or not to enable BitLocker until recovery information is stored in AD DS.
func (m *BitLockerRecoveryOptions) SetEnableBitLockerAfterRecoveryInformationToStore(value *bool)() {
    m.enableBitLockerAfterRecoveryInformationToStore = value
}
// SetEnableRecoveryInformationSaveToStore sets the enableRecoveryInformationSaveToStore property value. Indicates whether or not to allow BitLocker recovery information to store in AD DS.
func (m *BitLockerRecoveryOptions) SetEnableRecoveryInformationSaveToStore(value *bool)() {
    m.enableRecoveryInformationSaveToStore = value
}
// SetHideRecoveryOptions sets the hideRecoveryOptions property value. Indicates whether or not to allow showing recovery options in BitLocker Setup Wizard for fixed or system disk.
func (m *BitLockerRecoveryOptions) SetHideRecoveryOptions(value *bool)() {
    m.hideRecoveryOptions = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BitLockerRecoveryOptions) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecoveryInformationToStore sets the recoveryInformationToStore property value. BitLockerRecoveryInformationType types
func (m *BitLockerRecoveryOptions) SetRecoveryInformationToStore(value *BitLockerRecoveryInformationType)() {
    m.recoveryInformationToStore = value
}
// SetRecoveryKeyUsage sets the recoveryKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerRecoveryOptions) SetRecoveryKeyUsage(value *ConfigurationUsage)() {
    m.recoveryKeyUsage = value
}
// SetRecoveryPasswordUsage sets the recoveryPasswordUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerRecoveryOptions) SetRecoveryPasswordUsage(value *ConfigurationUsage)() {
    m.recoveryPasswordUsage = value
}
