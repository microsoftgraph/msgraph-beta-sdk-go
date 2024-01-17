package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingCollectionDefinition 
type DeviceManagementConfigurationSimpleSettingCollectionDefinition struct {
    DeviceManagementConfigurationSimpleSettingDefinition
}
// NewDeviceManagementConfigurationSimpleSettingCollectionDefinition instantiates a new deviceManagementConfigurationSimpleSettingCollectionDefinition and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingCollectionDefinition()(*DeviceManagementConfigurationSimpleSettingCollectionDefinition) {
    m := &DeviceManagementConfigurationSimpleSettingCollectionDefinition{
        DeviceManagementConfigurationSimpleSettingDefinition: *NewDeviceManagementConfigurationSimpleSettingDefinition(),
    }
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingCollectionDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingCollectionDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSimpleSettingCollectionDefinition(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingCollectionDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSimpleSettingDefinition.GetFieldDeserializers()
    res["maximumCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumCount(val)
        }
        return nil
    }
    res["minimumCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumCount(val)
        }
        return nil
    }
    return res
}
// GetMaximumCount gets the maximumCount property value. Maximum number of simple settings in the collection
func (m *DeviceManagementConfigurationSimpleSettingCollectionDefinition) GetMaximumCount()(*int32) {
    val, err := m.GetBackingStore().Get("maximumCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinimumCount gets the minimumCount property value. Minimum number of simple settings in the collection
func (m *DeviceManagementConfigurationSimpleSettingCollectionDefinition) GetMinimumCount()(*int32) {
    val, err := m.GetBackingStore().Get("minimumCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingCollectionDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSimpleSettingDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("maximumCount", m.GetMaximumCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumCount", m.GetMinimumCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMaximumCount sets the maximumCount property value. Maximum number of simple settings in the collection
func (m *DeviceManagementConfigurationSimpleSettingCollectionDefinition) SetMaximumCount(value *int32)() {
    err := m.GetBackingStore().Set("maximumCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumCount sets the minimumCount property value. Minimum number of simple settings in the collection
func (m *DeviceManagementConfigurationSimpleSettingCollectionDefinition) SetMinimumCount(value *int32)() {
    err := m.GetBackingStore().Set("minimumCount", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationSimpleSettingCollectionDefinitionable 
type DeviceManagementConfigurationSimpleSettingCollectionDefinitionable interface {
    DeviceManagementConfigurationSimpleSettingDefinitionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumCount()(*int32)
    GetMinimumCount()(*int32)
    SetMaximumCount(value *int32)()
    SetMinimumCount(value *int32)()
}
