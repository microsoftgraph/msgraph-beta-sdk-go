package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSetting setting instance within policy
type DeviceManagementConfigurationSetting struct {
    Entity
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSetting) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingDefinitions gets the settingDefinitions property value. List of related Setting Definitions. This property is read-only.
func (m *DeviceManagementConfigurationSetting) GetSettingDefinitions()([]DeviceManagementConfigurationSettingDefinitionable) {
    val, err := m.GetBackingStore().Get("settingDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationSettingDefinitionable)
    }
    return nil
}
// GetSettingInstance gets the settingInstance property value. Setting instance within policy
func (m *DeviceManagementConfigurationSetting) GetSettingInstance()(DeviceManagementConfigurationSettingInstanceable) {
    val, err := m.GetBackingStore().Get("settingInstance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationSettingInstanceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
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
        err = writer.WriteObjectValue("settingInstance", m.GetSettingInstance())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSetting) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingDefinitions sets the settingDefinitions property value. List of related Setting Definitions. This property is read-only.
func (m *DeviceManagementConfigurationSetting) SetSettingDefinitions(value []DeviceManagementConfigurationSettingDefinitionable)() {
    err := m.GetBackingStore().Set("settingDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingInstance sets the settingInstance property value. Setting instance within policy
func (m *DeviceManagementConfigurationSetting) SetSettingInstance(value DeviceManagementConfigurationSettingInstanceable)() {
    err := m.GetBackingStore().Set("settingInstance", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationSettingable 
type DeviceManagementConfigurationSettingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetSettingDefinitions()([]DeviceManagementConfigurationSettingDefinitionable)
    GetSettingInstance()(DeviceManagementConfigurationSettingInstanceable)
    SetOdataType(value *string)()
    SetSettingDefinitions(value []DeviceManagementConfigurationSettingDefinitionable)()
    SetSettingInstance(value DeviceManagementConfigurationSettingInstanceable)()
}
