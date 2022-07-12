package models

import (
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
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBlockDataRecoveryAgent gets the blockDataRecoveryAgent property value. Indicates whether to block certificate-based data recovery agent.
func (m *BitLockerRecoveryOptions) GetBlockDataRecoveryAgent()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockDataRecoveryAgent
    }
}
// GetEnableBitLockerAfterRecoveryInformationToStore gets the enableBitLockerAfterRecoveryInformationToStore property value. Indicates whether or not to enable BitLocker until recovery information is stored in AD DS.
func (m *BitLockerRecoveryOptions) GetEnableBitLockerAfterRecoveryInformationToStore()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableBitLockerAfterRecoveryInformationToStore
    }
}
// GetEnableRecoveryInformationSaveToStore gets the enableRecoveryInformationSaveToStore property value. Indicates whether or not to allow BitLocker recovery information to store in AD DS.
func (m *BitLockerRecoveryOptions) GetEnableRecoveryInformationSaveToStore()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableRecoveryInformationSaveToStore
    }
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
    if m == nil {
        return nil
    } else {
        return m.hideRecoveryOptions
    }
}
// GetRecoveryInformationToStore gets the recoveryInformationToStore property value. BitLockerRecoveryInformationType types
func (m *BitLockerRecoveryOptions) GetRecoveryInformationToStore()(*BitLockerRecoveryInformationType) {
    if m == nil {
        return nil
    } else {
        return m.recoveryInformationToStore
    }
}
// GetRecoveryKeyUsage gets the recoveryKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerRecoveryOptions) GetRecoveryKeyUsage()(*ConfigurationUsage) {
    if m == nil {
        return nil
    } else {
        return m.recoveryKeyUsage
    }
}
// GetRecoveryPasswordUsage gets the recoveryPasswordUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerRecoveryOptions) GetRecoveryPasswordUsage()(*ConfigurationUsage) {
    if m == nil {
        return nil
    } else {
        return m.recoveryPasswordUsage
    }
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
    if m != nil {
        m.additionalData = value
    }
}
// SetBlockDataRecoveryAgent sets the blockDataRecoveryAgent property value. Indicates whether to block certificate-based data recovery agent.
func (m *BitLockerRecoveryOptions) SetBlockDataRecoveryAgent(value *bool)() {
    if m != nil {
        m.blockDataRecoveryAgent = value
    }
}
// SetEnableBitLockerAfterRecoveryInformationToStore sets the enableBitLockerAfterRecoveryInformationToStore property value. Indicates whether or not to enable BitLocker until recovery information is stored in AD DS.
func (m *BitLockerRecoveryOptions) SetEnableBitLockerAfterRecoveryInformationToStore(value *bool)() {
    if m != nil {
        m.enableBitLockerAfterRecoveryInformationToStore = value
    }
}
// SetEnableRecoveryInformationSaveToStore sets the enableRecoveryInformationSaveToStore property value. Indicates whether or not to allow BitLocker recovery information to store in AD DS.
func (m *BitLockerRecoveryOptions) SetEnableRecoveryInformationSaveToStore(value *bool)() {
    if m != nil {
        m.enableRecoveryInformationSaveToStore = value
    }
}
// SetHideRecoveryOptions sets the hideRecoveryOptions property value. Indicates whether or not to allow showing recovery options in BitLocker Setup Wizard for fixed or system disk.
func (m *BitLockerRecoveryOptions) SetHideRecoveryOptions(value *bool)() {
    if m != nil {
        m.hideRecoveryOptions = value
    }
}
// SetRecoveryInformationToStore sets the recoveryInformationToStore property value. BitLockerRecoveryInformationType types
func (m *BitLockerRecoveryOptions) SetRecoveryInformationToStore(value *BitLockerRecoveryInformationType)() {
    if m != nil {
        m.recoveryInformationToStore = value
    }
}
// SetRecoveryKeyUsage sets the recoveryKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerRecoveryOptions) SetRecoveryKeyUsage(value *ConfigurationUsage)() {
    if m != nil {
        m.recoveryKeyUsage = value
    }
}
// SetRecoveryPasswordUsage sets the recoveryPasswordUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerRecoveryOptions) SetRecoveryPasswordUsage(value *ConfigurationUsage)() {
    if m != nil {
        m.recoveryPasswordUsage = value
    }
}
