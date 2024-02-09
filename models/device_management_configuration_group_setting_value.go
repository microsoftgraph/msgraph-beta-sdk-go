package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationGroupSettingValue value of the GroupSetting
type DeviceManagementConfigurationGroupSettingValue struct {
    DeviceManagementConfigurationSettingValue
}
// NewDeviceManagementConfigurationGroupSettingValue instantiates a new DeviceManagementConfigurationGroupSettingValue and sets the default values.
func NewDeviceManagementConfigurationGroupSettingValue()(*DeviceManagementConfigurationGroupSettingValue) {
    m := &DeviceManagementConfigurationGroupSettingValue{
        DeviceManagementConfigurationSettingValue: *NewDeviceManagementConfigurationSettingValue(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationGroupSettingValue"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationGroupSettingValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementConfigurationGroupSettingValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationGroupSettingValue(), nil
}
// GetChildren gets the children property value. Collection of child setting instances contained within this GroupSetting
// returns a []DeviceManagementConfigurationSettingInstanceable when successful
func (m *DeviceManagementConfigurationGroupSettingValue) GetChildren()([]DeviceManagementConfigurationSettingInstanceable) {
    val, err := m.GetBackingStore().Get("children")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationSettingInstanceable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementConfigurationGroupSettingValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingValue.GetFieldDeserializers()
    res["children"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingInstanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationSettingInstanceable)
                }
            }
            m.SetChildren(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationGroupSettingValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingValue.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildren() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildren()))
        for i, v := range m.GetChildren() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("children", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildren sets the children property value. Collection of child setting instances contained within this GroupSetting
func (m *DeviceManagementConfigurationGroupSettingValue) SetChildren(value []DeviceManagementConfigurationSettingInstanceable)() {
    err := m.GetBackingStore().Set("children", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementConfigurationGroupSettingValueable interface {
    DeviceManagementConfigurationSettingValueable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildren()([]DeviceManagementConfigurationSettingInstanceable)
    SetChildren(value []DeviceManagementConfigurationSettingInstanceable)()
}
