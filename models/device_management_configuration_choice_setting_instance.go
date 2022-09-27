package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingInstance 
type DeviceManagementConfigurationChoiceSettingInstance struct {
    DeviceManagementConfigurationSettingInstance
    // The choiceSettingValue property
    choiceSettingValue DeviceManagementConfigurationChoiceSettingValueable
}
// NewDeviceManagementConfigurationChoiceSettingInstance instantiates a new DeviceManagementConfigurationChoiceSettingInstance and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingInstance()(*DeviceManagementConfigurationChoiceSettingInstance) {
    m := &DeviceManagementConfigurationChoiceSettingInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationChoiceSettingInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingInstance(), nil
}
// GetChoiceSettingValue gets the choiceSettingValue property value. The choiceSettingValue property
func (m *DeviceManagementConfigurationChoiceSettingInstance) GetChoiceSettingValue()(DeviceManagementConfigurationChoiceSettingValueable) {
    return m.choiceSettingValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
    res["choiceSettingValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementConfigurationChoiceSettingValueFromDiscriminatorValue , m.SetChoiceSettingValue)
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("choiceSettingValue", m.GetChoiceSettingValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChoiceSettingValue sets the choiceSettingValue property value. The choiceSettingValue property
func (m *DeviceManagementConfigurationChoiceSettingInstance) SetChoiceSettingValue(value DeviceManagementConfigurationChoiceSettingValueable)() {
    m.choiceSettingValue = value
}
