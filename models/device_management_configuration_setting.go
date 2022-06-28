package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSetting setting instance within policy
type DeviceManagementConfigurationSetting struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSetting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settingDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSettingDefinitionable)
            }
            m.SetSettingDefinitions(res)
        }
        return nil
    }
    res["settingInstance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstance(val.(DeviceManagementConfigurationSettingInstanceable))
        }
        return nil
    }
    return res
}
// GetSettingDefinitions gets the settingDefinitions property value. List of related Setting Definitions. This property is read-only.
func (m *DeviceManagementConfigurationSetting) GetSettingDefinitions()([]DeviceManagementConfigurationSettingDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitions
    }
}
// GetSettingInstance gets the settingInstance property value. Setting instance within policy
func (m *DeviceManagementConfigurationSetting) GetSettingInstance()(DeviceManagementConfigurationSettingInstanceable) {
    if m == nil {
        return nil
    } else {
        return m.settingInstance
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSettingDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettingDefinitions()))
        for i, v := range m.GetSettingDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSetting) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSettingDefinitions sets the settingDefinitions property value. List of related Setting Definitions. This property is read-only.
func (m *DeviceManagementConfigurationSetting) SetSettingDefinitions(value []DeviceManagementConfigurationSettingDefinitionable)() {
    if m != nil {
        m.settingDefinitions = value
    }
}
// SetSettingInstance sets the settingInstance property value. Setting instance within policy
func (m *DeviceManagementConfigurationSetting) SetSettingInstance(value DeviceManagementConfigurationSettingInstanceable)() {
    if m != nil {
        m.settingInstance = value
    }
}
