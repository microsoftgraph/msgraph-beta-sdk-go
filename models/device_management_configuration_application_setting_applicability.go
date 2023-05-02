package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationApplicationSettingApplicability 
type DeviceManagementConfigurationApplicationSettingApplicability struct {
    DeviceManagementConfigurationSettingApplicability
}
// NewDeviceManagementConfigurationApplicationSettingApplicability instantiates a new DeviceManagementConfigurationApplicationSettingApplicability and sets the default values.
func NewDeviceManagementConfigurationApplicationSettingApplicability()(*DeviceManagementConfigurationApplicationSettingApplicability) {
    m := &DeviceManagementConfigurationApplicationSettingApplicability{
        DeviceManagementConfigurationSettingApplicability: *NewDeviceManagementConfigurationSettingApplicability(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationApplicationSettingApplicability"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationApplicationSettingApplicabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationApplicationSettingApplicabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationApplicationSettingApplicability(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationApplicationSettingApplicability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingApplicability.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationApplicationSettingApplicability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingApplicability.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// DeviceManagementConfigurationApplicationSettingApplicabilityable 
type DeviceManagementConfigurationApplicationSettingApplicabilityable interface {
    DeviceManagementConfigurationSettingApplicabilityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
