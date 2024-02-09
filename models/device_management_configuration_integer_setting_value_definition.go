package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DeviceManagementConfigurationIntegerSettingValueDefinition struct {
    DeviceManagementConfigurationSettingValueDefinition
}
// NewDeviceManagementConfigurationIntegerSettingValueDefinition instantiates a new DeviceManagementConfigurationIntegerSettingValueDefinition and sets the default values.
func NewDeviceManagementConfigurationIntegerSettingValueDefinition()(*DeviceManagementConfigurationIntegerSettingValueDefinition) {
    m := &DeviceManagementConfigurationIntegerSettingValueDefinition{
        DeviceManagementConfigurationSettingValueDefinition: *NewDeviceManagementConfigurationSettingValueDefinition(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationIntegerSettingValueDefinition"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationIntegerSettingValueDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementConfigurationIntegerSettingValueDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationIntegerSettingValueDefinition(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
// returns a *int64 when successful
func (m *DeviceManagementConfigurationIntegerSettingValueDefinition) GetMaximumValue()(*int64) {
    val, err := m.GetBackingStore().Get("maximumValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetMinimumValue gets the minimumValue property value. Minimum allowed value of the integer
// returns a *int64 when successful
func (m *DeviceManagementConfigurationIntegerSettingValueDefinition) GetMinimumValue()(*int64) {
    val, err := m.GetBackingStore().Get("minimumValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
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
    return nil
}
// SetMaximumValue sets the maximumValue property value. Maximum allowed value of the integer
func (m *DeviceManagementConfigurationIntegerSettingValueDefinition) SetMaximumValue(value *int64)() {
    err := m.GetBackingStore().Set("maximumValue", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumValue sets the minimumValue property value. Minimum allowed value of the integer
func (m *DeviceManagementConfigurationIntegerSettingValueDefinition) SetMinimumValue(value *int64)() {
    err := m.GetBackingStore().Set("minimumValue", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementConfigurationIntegerSettingValueDefinitionable interface {
    DeviceManagementConfigurationSettingValueDefinitionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumValue()(*int64)
    GetMinimumValue()(*int64)
    SetMaximumValue(value *int64)()
    SetMinimumValue(value *int64)()
}
