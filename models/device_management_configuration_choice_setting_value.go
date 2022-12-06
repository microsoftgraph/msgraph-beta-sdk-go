package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingValue 
type DeviceManagementConfigurationChoiceSettingValue struct {
    DeviceManagementConfigurationSettingValue
    // Child settings.
    children []DeviceManagementConfigurationSettingInstanceable
    // Choice setting value: an OptionDefinition ItemId.
    value *string
}
// NewDeviceManagementConfigurationChoiceSettingValue instantiates a new DeviceManagementConfigurationChoiceSettingValue and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingValue()(*DeviceManagementConfigurationChoiceSettingValue) {
    m := &DeviceManagementConfigurationChoiceSettingValue{
        DeviceManagementConfigurationSettingValue: *NewDeviceManagementConfigurationSettingValue(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationChoiceSettingValue";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingValue(), nil
}
// GetChildren gets the children property value. Child settings.
func (m *DeviceManagementConfigurationChoiceSettingValue) GetChildren()([]DeviceManagementConfigurationSettingInstanceable) {
    return m.children
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingValue.GetFieldDeserializers()
    res["children"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingInstanceFromDiscriminatorValue , m.SetChildren)
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetValue)
    return res
}
// GetValue gets the value property value. Choice setting value: an OptionDefinition ItemId.
func (m *DeviceManagementConfigurationChoiceSettingValue) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingValue.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildren() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetChildren())
        err = writer.WriteCollectionOfObjectValues("children", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildren sets the children property value. Child settings.
func (m *DeviceManagementConfigurationChoiceSettingValue) SetChildren(value []DeviceManagementConfigurationSettingInstanceable)() {
    m.children = value
}
// SetValue sets the value property value. Choice setting value: an OptionDefinition ItemId.
func (m *DeviceManagementConfigurationChoiceSettingValue) SetValue(value *string)() {
    m.value = value
}
