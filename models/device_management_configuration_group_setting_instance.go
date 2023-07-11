package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationGroupSettingInstance instance of a GroupSetting
type DeviceManagementConfigurationGroupSettingInstance struct {
    DeviceManagementConfigurationSettingInstance
    // The OdataType property
    OdataType *string
}
// NewDeviceManagementConfigurationGroupSettingInstance instantiates a new deviceManagementConfigurationGroupSettingInstance and sets the default values.
func NewDeviceManagementConfigurationGroupSettingInstance()(*DeviceManagementConfigurationGroupSettingInstance) {
    m := &DeviceManagementConfigurationGroupSettingInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationGroupSettingInstance"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationGroupSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationGroupSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationGroupSettingInstance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationGroupSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
    res["groupSettingValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationGroupSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupSettingValue(val.(DeviceManagementConfigurationGroupSettingValueable))
        }
        return nil
    }
    return res
}
// GetGroupSettingValue gets the groupSettingValue property value. The groupSettingValue property
func (m *DeviceManagementConfigurationGroupSettingInstance) GetGroupSettingValue()(DeviceManagementConfigurationGroupSettingValueable) {
    val, err := m.GetBackingStore().Get("groupSettingValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationGroupSettingValueable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationGroupSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("groupSettingValue", m.GetGroupSettingValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupSettingValue sets the groupSettingValue property value. The groupSettingValue property
func (m *DeviceManagementConfigurationGroupSettingInstance) SetGroupSettingValue(value DeviceManagementConfigurationGroupSettingValueable)() {
    err := m.GetBackingStore().Set("groupSettingValue", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationGroupSettingInstanceable 
type DeviceManagementConfigurationGroupSettingInstanceable interface {
    DeviceManagementConfigurationSettingInstanceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupSettingValue()(DeviceManagementConfigurationGroupSettingValueable)
    SetGroupSettingValue(value DeviceManagementConfigurationGroupSettingValueable)()
}
