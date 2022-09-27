package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationIntegerSettingValue 
type DeviceManagementConfigurationIntegerSettingValue struct {
    DeviceManagementConfigurationSimpleSettingValue
    // Value of the integer setting.
    value *int32
}
// NewDeviceManagementConfigurationIntegerSettingValue instantiates a new DeviceManagementConfigurationIntegerSettingValue and sets the default values.
func NewDeviceManagementConfigurationIntegerSettingValue()(*DeviceManagementConfigurationIntegerSettingValue) {
    m := &DeviceManagementConfigurationIntegerSettingValue{
        DeviceManagementConfigurationSimpleSettingValue: *NewDeviceManagementConfigurationSimpleSettingValue(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationIntegerSettingValue";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationIntegerSettingValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationIntegerSettingValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationIntegerSettingValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationIntegerSettingValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSimpleSettingValue.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetValue)
    return res
}
// GetValue gets the value property value. Value of the integer setting.
func (m *DeviceManagementConfigurationIntegerSettingValue) GetValue()(*int32) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationIntegerSettingValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSimpleSettingValue.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. Value of the integer setting.
func (m *DeviceManagementConfigurationIntegerSettingValue) SetValue(value *int32)() {
    m.value = value
}
