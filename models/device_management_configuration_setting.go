package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSetting setting instance within policy
type DeviceManagementConfigurationSetting struct {
    Entity
    // List of related Setting Definitions. This property is read-only.
    settingDefinitions []DeviceManagementConfigurationSettingDefinitionable
    // Setting instance within policy
    settingInstance DeviceManagementConfigurationSettingInstanceable
}
// NewDeviceManagementConfigurationSetting instantiates a new deviceManagementConfigurationSetting and sets the default values.
func NewDeviceManagementConfigurationSetting()(*DeviceManagementConfigurationSetting) {
    m := &DeviceManagementConfigurationSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementConfigurationSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSetting(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settingDefinitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue , m.SetSettingDefinitions)
    res["settingInstance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementConfigurationSettingInstanceFromDiscriminatorValue , m.SetSettingInstance)
    return res
}
// GetSettingDefinitions gets the settingDefinitions property value. List of related Setting Definitions. This property is read-only.
func (m *DeviceManagementConfigurationSetting) GetSettingDefinitions()([]DeviceManagementConfigurationSettingDefinitionable) {
    return m.settingDefinitions
}
// GetSettingInstance gets the settingInstance property value. Setting instance within policy
func (m *DeviceManagementConfigurationSetting) GetSettingInstance()(DeviceManagementConfigurationSettingInstanceable) {
    return m.settingInstance
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSettingDefinitions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSettingDefinitions())
        err = writer.WriteCollectionOfObjectValues("settingDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settingInstance", m.GetSettingInstance())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettingDefinitions sets the settingDefinitions property value. List of related Setting Definitions. This property is read-only.
func (m *DeviceManagementConfigurationSetting) SetSettingDefinitions(value []DeviceManagementConfigurationSettingDefinitionable)() {
    m.settingDefinitions = value
}
// SetSettingInstance sets the settingInstance property value. Setting instance within policy
func (m *DeviceManagementConfigurationSetting) SetSettingInstance(value DeviceManagementConfigurationSettingInstanceable)() {
    m.settingInstance = value
}
