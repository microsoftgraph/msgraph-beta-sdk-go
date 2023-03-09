package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate 
type DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate struct {
    DeviceManagementConfigurationChoiceSettingValueDefaultTemplate
}
// NewDeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate instantiates a new DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate()(*DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) {
    m := &DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate{
        DeviceManagementConfigurationChoiceSettingValueDefaultTemplate: *NewDeviceManagementConfigurationChoiceSettingValueDefaultTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate(), nil
}
// GetChildren gets the children property value. Option Children
func (m *DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) GetChildren()([]DeviceManagementConfigurationSettingInstanceTemplateable) {
    val, err := m.GetBackingStore().Get("children")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationSettingInstanceTemplateable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationChoiceSettingValueDefaultTemplate.GetFieldDeserializers()
    res["children"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingInstanceTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingInstanceTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSettingInstanceTemplateable)
            }
            m.SetChildren(res)
        }
        return nil
    }
    res["settingDefinitionOptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDefinitionOptionId(val)
        }
        return nil
    }
    return res
}
// GetSettingDefinitionOptionId gets the settingDefinitionOptionId property value. Default Constant Value
func (m *DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) GetSettingDefinitionOptionId()(*string) {
    val, err := m.GetBackingStore().Get("settingDefinitionOptionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationChoiceSettingValueDefaultTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildren() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildren()))
        for i, v := range m.GetChildren() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    err := m.GetBackingStore().Set("children", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingDefinitionOptionId sets the settingDefinitionOptionId property value. Default Constant Value
func (m *DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) SetSettingDefinitionOptionId(value *string)() {
    err := m.GetBackingStore().Set("settingDefinitionOptionId", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplateable 
type DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplateable interface {
    DeviceManagementConfigurationChoiceSettingValueDefaultTemplateable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildren()([]DeviceManagementConfigurationSettingInstanceTemplateable)
    GetSettingDefinitionOptionId()(*string)
    SetChildren(value []DeviceManagementConfigurationSettingInstanceTemplateable)()
    SetSettingDefinitionOptionId(value *string)()
}
