package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate 
type DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate struct {
    DeviceManagementConfigurationChoiceSettingValueDefaultTemplate
    // Option Children
    children []DeviceManagementConfigurationSettingInstanceTemplateable
    // Default Constant Value
    settingDefinitionOptionId *string
}
// NewDeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate instantiates a new DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate()(*DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) {
    m := &DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate{
        DeviceManagementConfigurationChoiceSettingValueDefaultTemplate: *NewDeviceManagementConfigurationChoiceSettingValueDefaultTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate(), nil
}
// GetChildren gets the children property value. Option Children
func (m *DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) GetChildren()([]DeviceManagementConfigurationSettingInstanceTemplateable) {
    return m.children
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationChoiceSettingValueDefaultTemplate.GetFieldDeserializers()
    res["children"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingInstanceTemplateFromDiscriminatorValue , m.SetChildren)
    res["settingDefinitionOptionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSettingDefinitionOptionId)
    return res
}
// GetSettingDefinitionOptionId gets the settingDefinitionOptionId property value. Default Constant Value
func (m *DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) GetSettingDefinitionOptionId()(*string) {
    return m.settingDefinitionOptionId
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationChoiceSettingValueDefaultTemplate.Serialize(writer)
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
        err = writer.WriteStringValue("settingDefinitionOptionId", m.GetSettingDefinitionOptionId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildren sets the children property value. Option Children
func (m *DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) SetChildren(value []DeviceManagementConfigurationSettingInstanceTemplateable)() {
    m.children = value
}
// SetSettingDefinitionOptionId sets the settingDefinitionOptionId property value. Default Constant Value
func (m *DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) SetSettingDefinitionOptionId(value *string)() {
    m.settingDefinitionOptionId = value
}
