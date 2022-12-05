package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody provides operations to call the createEnrollmentNotificationConfiguration method.
type DeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceEnrollmentNotificationConfigurations property
    deviceEnrollmentNotificationConfigurations []DeviceEnrollmentConfigurationable
}
// NewDeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody instantiates a new DeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody and sets the default values.
func NewDeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody()(*DeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) {
    m := &DeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceEnrollmentNotificationConfigurations gets the deviceEnrollmentNotificationConfigurations property value. The deviceEnrollmentNotificationConfigurations property
func (m *DeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) GetDeviceEnrollmentNotificationConfigurations()([]DeviceEnrollmentConfigurationable) {
    return m.deviceEnrollmentNotificationConfigurations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceEnrollmentNotificationConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceEnrollmentConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceEnrollmentConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceEnrollmentConfigurationable)
            }
            m.SetDeviceEnrollmentNotificationConfigurations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceEnrollmentNotificationConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceEnrollmentNotificationConfigurations()))
        for i, v := range m.GetDeviceEnrollmentNotificationConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("deviceEnrollmentNotificationConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceEnrollmentNotificationConfigurations sets the deviceEnrollmentNotificationConfigurations property value. The deviceEnrollmentNotificationConfigurations property
func (m *DeviceManagementDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) SetDeviceEnrollmentNotificationConfigurations(value []DeviceEnrollmentConfigurationable)() {
    m.deviceEnrollmentNotificationConfigurations = value
}
