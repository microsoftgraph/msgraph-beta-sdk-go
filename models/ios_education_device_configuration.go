package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosEducationDeviceConfiguration iOS Education configuration profile
type IosEducationDeviceConfiguration struct {
    DeviceConfiguration
}
// NewIosEducationDeviceConfiguration instantiates a new iosEducationDeviceConfiguration and sets the default values.
func NewIosEducationDeviceConfiguration()(*IosEducationDeviceConfiguration) {
    m := &IosEducationDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosEducationDeviceConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosEducationDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosEducationDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosEducationDeviceConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosEducationDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *IosEducationDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// IosEducationDeviceConfigurationable 
type IosEducationDeviceConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
