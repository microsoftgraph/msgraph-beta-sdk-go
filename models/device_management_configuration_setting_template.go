package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingTemplate setting Template
type DeviceManagementConfigurationSettingTemplate struct {
    Entity
    // List of related Setting Definitions
    settingDefinitions []DeviceManagementConfigurationSettingDefinitionable
    // Setting Instance Template
    settingInstanceTemplate DeviceManagementConfigurationSettingInstanceTemplateable
}
// NewDeviceManagementConfigurationSettingTemplate instantiates a new deviceManagementConfigurationSettingTemplate and sets the default values.
func NewDeviceManagementConfigurationSettingTemplate()(*DeviceManagementConfigurationSettingTemplate) {
    m := &DeviceManagementConfigurationSettingTemplate{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSettingTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSettingTemplate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settingDefinitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue , m.SetSettingDefinitions)
    res["settingInstanceTemplate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementConfigurationSettingInstanceTemplateFromDiscriminatorValue , m.SetSettingInstanceTemplate)
    return res
}
// GetSettingDefinitions gets the settingDefinitions property value. List of related Setting Definitions
func (m *DeviceManagementConfigurationSettingTemplate) GetSettingDefinitions()([]DeviceManagementConfigurationSettingDefinitionable) {
    return m.settingDefinitions
}
// GetSettingInstanceTemplate gets the settingInstanceTemplate property value. Setting Instance Template
func (m *DeviceManagementConfigurationSettingTemplate) GetSettingInstanceTemplate()(DeviceManagementConfigurationSettingInstanceTemplateable) {
    return m.settingInstanceTemplate
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("settingInstanceTemplate", m.GetSettingInstanceTemplate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettingDefinitions sets the settingDefinitions property value. List of related Setting Definitions
func (m *DeviceManagementConfigurationSettingTemplate) SetSettingDefinitions(value []DeviceManagementConfigurationSettingDefinitionable)() {
    m.settingDefinitions = value
}
// SetSettingInstanceTemplate sets the settingInstanceTemplate property value. Setting Instance Template
func (m *DeviceManagementConfigurationSettingTemplate) SetSettingInstanceTemplate(value DeviceManagementConfigurationSettingInstanceTemplateable)() {
    m.settingInstanceTemplate = value
}
