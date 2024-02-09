package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingTemplate setting Template
type DeviceManagementConfigurationSettingTemplate struct {
    Entity
}
// NewDeviceManagementConfigurationSettingTemplate instantiates a new DeviceManagementConfigurationSettingTemplate and sets the default values.
func NewDeviceManagementConfigurationSettingTemplate()(*DeviceManagementConfigurationSettingTemplate) {
    m := &DeviceManagementConfigurationSettingTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSettingTemplate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementConfigurationSettingTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settingDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationSettingDefinitionable)
                }
            }
            m.SetSettingDefinitions(res)
        }
        return nil
    }
    res["settingInstanceTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSettingInstanceTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstanceTemplate(val.(DeviceManagementConfigurationSettingInstanceTemplateable))
        }
        return nil
    }
    return res
}
// GetSettingDefinitions gets the settingDefinitions property value. List of related Setting Definitions
// returns a []DeviceManagementConfigurationSettingDefinitionable when successful
func (m *DeviceManagementConfigurationSettingTemplate) GetSettingDefinitions()([]DeviceManagementConfigurationSettingDefinitionable) {
    val, err := m.GetBackingStore().Get("settingDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationSettingDefinitionable)
    }
    return nil
}
// GetSettingInstanceTemplate gets the settingInstanceTemplate property value. Setting Instance Template
// returns a DeviceManagementConfigurationSettingInstanceTemplateable when successful
func (m *DeviceManagementConfigurationSettingTemplate) GetSettingInstanceTemplate()(DeviceManagementConfigurationSettingInstanceTemplateable) {
    val, err := m.GetBackingStore().Get("settingInstanceTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationSettingInstanceTemplateable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSettingDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettingDefinitions()))
        for i, v := range m.GetSettingDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    err := m.GetBackingStore().Set("settingDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingInstanceTemplate sets the settingInstanceTemplate property value. Setting Instance Template
func (m *DeviceManagementConfigurationSettingTemplate) SetSettingInstanceTemplate(value DeviceManagementConfigurationSettingInstanceTemplateable)() {
    err := m.GetBackingStore().Set("settingInstanceTemplate", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementConfigurationSettingTemplateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSettingDefinitions()([]DeviceManagementConfigurationSettingDefinitionable)
    GetSettingInstanceTemplate()(DeviceManagementConfigurationSettingInstanceTemplateable)
    SetSettingDefinitions(value []DeviceManagementConfigurationSettingDefinitionable)()
    SetSettingInstanceTemplate(value DeviceManagementConfigurationSettingInstanceTemplateable)()
}
