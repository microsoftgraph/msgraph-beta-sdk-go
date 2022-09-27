package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingCollectionInstance 
type DeviceManagementConfigurationChoiceSettingCollectionInstance struct {
    DeviceManagementConfigurationSettingInstance
    // Choice setting collection value
    choiceSettingCollectionValue []DeviceManagementConfigurationChoiceSettingValueable
}
// NewDeviceManagementConfigurationChoiceSettingCollectionInstance instantiates a new DeviceManagementConfigurationChoiceSettingCollectionInstance and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingCollectionInstance()(*DeviceManagementConfigurationChoiceSettingCollectionInstance) {
    m := &DeviceManagementConfigurationChoiceSettingCollectionInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationChoiceSettingCollectionInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingCollectionInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingCollectionInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingCollectionInstance(), nil
}
// GetChoiceSettingCollectionValue gets the choiceSettingCollectionValue property value. Choice setting collection value
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstance) GetChoiceSettingCollectionValue()([]DeviceManagementConfigurationChoiceSettingValueable) {
    return m.choiceSettingCollectionValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
    res["choiceSettingCollectionValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationChoiceSettingValueFromDiscriminatorValue , m.SetChoiceSettingCollectionValue)
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChoiceSettingCollectionValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetChoiceSettingCollectionValue())
        err = writer.WriteCollectionOfObjectValues("choiceSettingCollectionValue", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChoiceSettingCollectionValue sets the choiceSettingCollectionValue property value. Choice setting collection value
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstance) SetChoiceSettingCollectionValue(value []DeviceManagementConfigurationChoiceSettingValueable)() {
    m.choiceSettingCollectionValue = value
}
