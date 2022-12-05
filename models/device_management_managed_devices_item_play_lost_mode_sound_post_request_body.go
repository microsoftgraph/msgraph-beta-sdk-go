package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody provides operations to call the playLostModeSound method.
type DeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The durationInMinutes property
    durationInMinutes *string
}
// NewDeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody instantiates a new DeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody and sets the default values.
func NewDeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody()(*DeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody) {
    m := &DeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDurationInMinutes gets the durationInMinutes property value. The durationInMinutes property
func (m *DeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody) GetDurationInMinutes()(*string) {
    return m.durationInMinutes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["durationInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInMinutes(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("durationInMinutes", m.GetDurationInMinutes())
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
func (m *DeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDurationInMinutes sets the durationInMinutes property value. The durationInMinutes property
func (m *DeviceManagementManagedDevicesItemPlayLostModeSoundPostRequestBody) SetDurationInMinutes(value *string)() {
    m.durationInMinutes = value
}
