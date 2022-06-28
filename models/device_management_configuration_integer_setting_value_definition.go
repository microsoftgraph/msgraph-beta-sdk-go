package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationIntegerSettingValueDefinition 
type DeviceManagementConfigurationIntegerSettingValueDefinition struct {
    DeviceManagementConfigurationSettingValueDefinition
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Maximum allowed value of the integer
    maximumValue *int64
    // Minimum allowed value of the integer
    minimumValue *int64
}
// NewDeviceManagementConfigurationIntegerSettingValueDefinition instantiates a new DeviceManagementConfigurationIntegerSettingValueDefinition and sets the default values.
func NewDeviceManagementConfigurationIntegerSettingValueDefinition()(*DeviceManagementConfigurationIntegerSettingValueDefinition) {
    m := &DeviceManagementConfigurationIntegerSettingValueDefinition{
        DeviceManagementConfigurationSettingValueDefinition: *NewDeviceManagementConfigurationSettingValueDefinition(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationIntegerSettingValueDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationIntegerSettingValueDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationIntegerSettingValueDefinition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationIntegerSettingValueDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationIntegerSettingValueDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingValueDefinition.GetFieldDeserializers()
    res["maximumValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumValue(val)
        }
        return nil
    }
    res["minimumValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumValue(val)
        }
        return nil
    }
    return res
}
// GetMaximumValue gets the maximumValue property value. Maximum allowed value of the integer
func (m *DeviceManagementConfigurationIntegerSettingValueDefinition) GetMaximumValue()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.maximumValue
    }
}
// GetMinimumValue gets the minimumValue property value. Minimum allowed value of the integer
func (m *DeviceManagementConfigurationIntegerSettingValueDefinition) GetMinimumValue()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.minimumValue
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationIntegerSettingValueDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingValueDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("maximumValue", m.GetMaximumValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("minimumValue", m.GetMinimumValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationIntegerSettingValueDefinition) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaximumValue sets the maximumValue property value. Maximum allowed value of the integer
func (m *DeviceManagementConfigurationIntegerSettingValueDefinition) SetMaximumValue(value *int64)() {
    if m != nil {
        m.maximumValue = value
    }
}
// SetMinimumValue sets the minimumValue property value. Minimum allowed value of the integer
func (m *DeviceManagementConfigurationIntegerSettingValueDefinition) SetMinimumValue(value *int64)() {
    if m != nil {
        m.minimumValue = value
    }
}
