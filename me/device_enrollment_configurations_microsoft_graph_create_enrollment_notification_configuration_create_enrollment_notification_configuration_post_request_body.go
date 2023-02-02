package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody 
type DeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deviceEnrollmentNotificationConfigurations property
    deviceEnrollmentNotificationConfigurations []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable
}
// NewDeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody instantiates a new DeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody and sets the default values.
func NewDeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody()(*DeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) {
    m := &DeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceEnrollmentNotificationConfigurations gets the deviceEnrollmentNotificationConfigurations property value. The deviceEnrollmentNotificationConfigurations property
func (m *DeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) GetDeviceEnrollmentNotificationConfigurations()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable) {
    return m.deviceEnrollmentNotificationConfigurations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceEnrollmentNotificationConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceEnrollmentConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable)
            }
            m.SetDeviceEnrollmentNotificationConfigurations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceEnrollmentNotificationConfigurations sets the deviceEnrollmentNotificationConfigurations property value. The deviceEnrollmentNotificationConfigurations property
func (m *DeviceEnrollmentConfigurationsMicrosoftGraphCreateEnrollmentNotificationConfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) SetDeviceEnrollmentNotificationConfigurations(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable)() {
    m.deviceEnrollmentNotificationConfigurations = value
}
